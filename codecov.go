// SPDX-FileCopyrightText: 2021 The go-codecov Authors
// SPDX-License-Identifier: BSD-3-Clause

package codecov

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/certifi/gocertifi"
	gzip "github.com/klauspost/compress/gzip"

	"github.com/zchee/go-codecov/provider"
)

// Token is the Codecov token environment variable.
const Token = "CODECOV_TOKEN"

var (
	// ignore error because valid Codecov url.
	baseURL, _ = url.Parse("https://codecov.io")

	// gocertifi.CACerts doesn't return error.
	caCerts, _ = gocertifi.CACerts()
)

// DefaultClient returns the default http.Client for send to Codecov.
func DefaultClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{
		RootCAs: caCerts,
	}

	hc := &http.Client{
		Transport: transport,
		// TODO(zchee): set Timeout?
	}

	return hc
}

// Client represents a CodeCov client.
type Client struct {
	hc      *http.Client
	baseURL *url.URL
	svc     provider.Service

	// queries
	commit      string // commit is the destination commit sha for the report. Required.
	token       string // token is a UUID token used to identify the project. Required.
	branch      string // branch is the target branch for the report. This value may be overridden during the Codecov discovery process.
	build       string // build is the build number provided by your CI service.
	job         int    // job is the job number provided by your CI service.
	buildURL    string // buildURL is the http url to link back to your CI provider.
	name        string // name is a custom name for this specific upload.
	slug        string // slug is the owner/repo slug name of the project.
	yamlPath    string // yaml is the relative path to the codecov.yml in this project.
	serviceName string // service is the CI service name. See below for acceptable values.
	flags       string // flags used for Flags. Can be one or more flags. E.g., flags=unit or flags=unit,java
	pr          string // pr is the pull request number this commit is currently found in.
}

// Endpoint represents a Codecov upload endpoint url and s3 url.
type Endpoint struct {
	// URL is the url of Codecov endpoint.
	URL string

	// S3URL is the s3 url of Codecov endpoint.
	S3URL string
}

// NewClient returns the new Codecov client.
func NewClient(hc *http.Client) *Client {
	if hc == nil {
		hc = defaultClient()
	}

	return &Client{
		hc:      hc,
		baseURL: baseURL,
		svc:     provider.Detect(),
	}
}

// Service returns the detected service configurations.
//
// This method mainly for debug to detect services.
func (c *Client) Service() {
	fmt.Printf("c.svc: %#v\n", c.svc)
	fmt.Printf("c.svc.Name(): %#v\n", c.svc.Name())
	fmt.Printf("c.svc.Root(): %#v\n", c.svc.Root())
	fmt.Printf("c.svc.Branch(): %#v\n", c.svc.Branch())
	fmt.Printf("c.svc.Build(): %#v\n", c.svc.Build())
	fmt.Printf("c.svc.BuildURL(): %#v\n", c.svc.BuildURL())
	fmt.Printf("c.svc.Job(): %#v\n", c.svc.Job())
	fmt.Printf("c.svc.Tag(): %#v\n", c.svc.Tag())
	fmt.Printf("c.svc.Commit(): %#v\n", c.svc.Commit())
	fmt.Printf("c.svc.PullRequest(): %#v\n", c.svc.PullRequest())
	fmt.Printf("c.svc.Slug(): %#v\n", c.svc.Slug())
	fmt.Printf("c.svc.Envs(): %#v\n", c.svc.Envs())
}

// Prepare prepares Codecov endpoint.
func (c *Client) Prepare(ctx context.Context) (ep Endpoint, err error) {
	u := *c.baseURL // copy
	u.Path = path.Join(u.Path, "upload", "v4")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)
	if err != nil {
		return ep, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("X-Reduced-Redundancy", "false")
	req.Header.Set("X-Content-Type", "application/x-gzip")

	resp, err := c.hc.Do(req)
	if err != nil {
		return ep, fmt.Errorf("post request to Codecov: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ep, fmt.Errorf("read response: %w", err)
	}

	result := bytes.Split(body, []byte("\n"))
	if len(result) > 2 {
		return ep, fmt.Errorf("replied result is not valid: %s", string(body))
	}

	ep.URL = string(result[0])
	ep.S3URL = string(result[1])

	return ep, nil
}

// Compress compresses coverage reports in gzip format.
func (c *Client) Compress(reports []byte) (io.Reader, error) {
	var b bytes.Buffer

	// ignore error because set valid compress level
	w, _ := gzip.NewWriterLevel(&b, gzip.BestCompression)

	if _, err := w.Write(reports); err != nil {
		return nil, fmt.Errorf("write reports to gzip writer: %w", err)
	}

	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("close gzip writer: %w", err)
	}

	return &b, nil
}

// Upload uploads coverage reports to Codecov endpoint.
func (c *Client) Upload(ctx context.Context, ep Endpoint, reports io.Reader) error {
	u, err := url.Parse(ep.S3URL)
	if err != nil {
		return fmt.Errorf("parse s3 %s url: %w", ep.S3URL, err)
	}

	q := u.Query()
	q.Set("package", fmt.Sprintf("go-%s", "0.0.0"))
	q.Set("service", c.svc.Name())
	q.Set("branch", c.svc.Branch())
	q.Set("commit", c.svc.Commit())
	q.Set("name", "")
	q.Set("tag", c.svc.Tag())
	q.Set("slug", c.svc.Slug())
	q.Set("pr", c.svc.PullRequest())
	q.Set("job", c.svc.Job())
	if token := os.Getenv(Token); token != "" {
		q.Set("token", token)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u.String(), reports)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-gzip")
	req.Header.Set("Content-Encoding", "gzip")

	resp, err := c.hc.Do(req)
	if err != nil {
		return fmt.Errorf("put request to Codecov: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read response body: %w", err)
		}

		return fmt.Errorf("failed to upload reports: %s", string(body))
	}

	return nil
}
