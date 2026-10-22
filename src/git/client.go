// stubber
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/git/client.go
// Original timestamp : 2026.10.22 03:27:41

package git

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"stubber/conn_cfg"
	"time"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	gol "github.com/jeanfrancoisgratton/gitoperationslib"
	"github.com/jeanfrancoisgratton/gitoperationslib/bitbucket"
	"github.com/jeanfrancoisgratton/gitoperationslib/forgejo"
	"github.com/jeanfrancoisgratton/gitoperationslib/github"
)

// Opens a gitoperationslib client on the resolved connection; the connection's server type picks the backend
// On Forgejo, this already contacts the server: it checks that its version is supported (GitHub and Bitbucket make no request here)
func connect(ctx context.Context, ac conn_cfg.ActiveConnection) (gol.Client, *cerr.CustomError) {
	var gc gol.Client
	var gerr error
	switch ac.ServerType {
	case conn_cfg.ServerTypeForgejo:
		creds := forgejo.Token(ac.Token)
		if ac.Auth == conn_cfg.AuthBasic {
			creds = forgejo.BasicAuth(ac.Username, ac.Password)
		}
		gc, gerr = forgejo.New(ctx, ac.URL, creds, forgejo.WithHTTPClient(httpClient()), forgejo.WithUserAgent(userAgent))
	case conn_cfg.ServerTypeGithub:
		// ResolveConnection() already refused GitHub connections without a token
		opts := []github.Option{github.WithHTTPClient(httpClient()), github.WithUserAgent(userAgent)}
		if !isGithubDotCom(ac.URL) {
			// GitHub Enterprise Server; the library appends /api/v3/
			opts = append(opts, github.WithEnterpriseURL(ac.URL, ""))
		}
		gc, gerr = github.New(ac.Token, opts...)
	case conn_cfg.ServerTypeBitBucket:
		// Bitbucket Cloud only (no Data Center)
		// Access tokens (workspace / project / repository) are bearer tokens, but Atlassian user API tokens (ATATT...)
		// need the Atlassian account e-mail, which we take from the connection's username
		creds := bitbucket.Token(ac.Token)
		switch {
		case ac.Auth == conn_cfg.AuthBasic:
			creds = bitbucket.BasicAuth(ac.Username, ac.Password)
		case isAtlassianAPIToken(ac.Token):
			if ac.Username == "" {
				return nil, &cerr.CustomError{Title: "Invalid connection",
					Message: fmt.Sprintf("Connection %s: an Atlassian API token needs your Atlassian account e-mail as username; fix it with: stubber conn edit %s", ac.Name, ac.Name)}
			}
			creds = bitbucket.APIToken(ac.Username, ac.Token)
		}
		opts := []bitbucket.Option{bitbucket.WithHTTPClient(httpClient()), bitbucket.WithUserAgent(userAgent)}
		if !isBitbucketDotOrg(ac.URL) {
			// Any other host is taken as the API root itself (ex: a proxy, or https://host/2.0)
			opts = append(opts, bitbucket.WithBaseURL(ac.URL))
		}
		gc, gerr = bitbucket.New(creds, opts...)
	default:
		return nil, &cerr.CustomError{Title: "Invalid connection", Message: fmt.Sprintf("Connection %s: unsupported server type '%s'", ac.Name, ac.ServerType)}
	}
	if gerr != nil {
		return nil, forgeError(fmt.Sprintf("Unable to connect to %s", ac.Name), gerr)
	}
	return gc, nil
}

// Converts a gitoperationslib error into our own error type; the HTTP status, if any, ends up in Code
func forgeError(title string, err error) *cerr.CustomError {
	if err == nil {
		return nil
	}
	return &cerr.CustomError{Title: title, Message: err.Error(), Code: gol.StatusCode(err)}
}

// github.com has its API on a separate host (api.github.com), which the library uses by default
func isGithubDotCom(serverURL string) bool {
	u, err := url.Parse(serverURL)
	if err != nil {
		return false
	}
	switch strings.ToLower(u.Hostname()) {
	case "github.com", "www.github.com", "api.github.com":
		return true
	}
	return false
}

// bitbucket.org has its API on a separate host (api.bitbucket.org/2.0/), which the library uses by default
func isBitbucketDotOrg(serverURL string) bool {
	u, err := url.Parse(serverURL)
	if err != nil {
		return false
	}
	switch strings.ToLower(u.Hostname()) {
	case "bitbucket.org", "www.bitbucket.org", "api.bitbucket.org":
		return true
	}
	return false
}

// Atlassian user API tokens start with ATATT; access tokens start with ATCTT
func isAtlassianAPIToken(token string) bool {
	return strings.HasPrefix(token, "ATATT")
}

// Fast-fail timeouts on connect / TLS handshake / response headers
func httpClient() *http.Client {
	dialer := &net.Dialer{Timeout: FastFailTimeout, KeepAlive: 30 * time.Second}
	return &http.Client{Transport: &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		TLSHandshakeTimeout:   FastFailTimeout,
		ResponseHeaderTimeout: FastFailTimeout,
		IdleConnTimeout:       90 * time.Second,
	}}
}
