// gitclt
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/conn_cfg/resolve.go
// Original timestamp : 2026.10.18 12:30:00

package conn_cfg

import (
	"fmt"
	"net/url"
	"slices"
	"strings"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
)

// Returns the named connection, ready to be used
// This is where a hand-edited, invalid connection gets caught

func GetConnection(name string) (ConnectionStruct, *cerr.CustomError) {
	cf, err := LoadConnectionsFile()
	if err != nil {
		return ConnectionStruct{}, err
	}

	c, ok := cf.Auths[name]
	if !ok {
		return ConnectionStruct{}, &cerr.CustomError{Title: "Unknown connection", Message: fmt.Sprintf("Connection %s does not exist", name)}
	}
	if err = c.Validate(); err != nil {
		err.Message = fmt.Sprintf("Connection %s: %s; fix it with: stubber conn edit %s", name, err.Message, name)
		return ConnectionStruct{}, err
	}
	return c, nil
}

// Finds the connection whose server hosts repoURL (ex: https://myserver:3000/OWNER/REPO), and splits the URL's path into owner and repo name
// connName (from -c) is optional: it is only needed when several connections point to the server hosting repoURL
// When given, it must be one of these connections
func ConnectionForURL(repoURL, connName string) (name, owner, repo string, ce *cerr.CustomError) {
	invalid := &cerr.CustomError{Title: "Invalid repository URL",
		Message: fmt.Sprintf("'%s' is not of the form https://SERVER[:PORT]/OWNER/REPO", repoURL)}
	ru, perr := url.Parse(repoURL)
	if perr != nil || ru.Scheme == "" || ru.Host == "" {
		return "", "", "", invalid
	}
	repoPath := strings.TrimSuffix(strings.Trim(ru.Path, "/"), ".git")

	cf, err := LoadConnectionsFile()
	if err != nil {
		return "", "", "", err
	}
	if _, ok := cf.Auths[connName]; connName != "" && !ok {
		return "", "", "", &cerr.CustomError{Title: "Unknown connection", Message: fmt.Sprintf("Connection %s does not exist", connName)}
	}

	// The connections pointing to the server hosting repoURL; their server type is the forge's
	type match struct{ name, owner, repo string }
	var matches []match
	hostMatched := false
	for n, c := range cf.Auths {
		cu, perr := url.Parse(c.URL)
		if perr != nil || serverHost(cu) != serverHost(ru) {
			continue
		}
		hostMatched = true
		// Sub-path installs (ex: https://example.com/forgejo): the repo lives under the connection's path
		rest := repoPath
		if p := serverPath(cu); p != "" {
			var found bool
			if rest, found = strings.CutPrefix(repoPath, p+"/"); !found {
				continue
			}
		}
		o, r, found := strings.Cut(rest, "/")
		if !found || o == "" || r == "" || strings.Contains(r, "/") {
			continue
		}
		matches = append(matches, match{n, o, r})
	}

	if connName != "" {
		for _, m := range matches {
			if m.name == connName {
				return m.name, m.owner, m.repo, nil
			}
		}
		if hostMatched && len(matches) == 0 {
			return "", "", "", invalid
		}
		return "", "", "", &cerr.CustomError{Title: "Connection mismatch",
			Message: fmt.Sprintf("Connection %s points to %s, not to the server hosting %s", connName, cf.Auths[connName].URL, repoURL)}
	}

	switch {
	case len(matches) == 1:
		return matches[0].name, matches[0].owner, matches[0].repo, nil
	case len(matches) == 0 && hostMatched:
		return "", "", "", invalid
	case len(matches) == 0:
		return "", "", "", &cerr.CustomError{Title: "No matching connection",
			Message: fmt.Sprintf("No connection points to %s (a %s server); add one with: stubber conn add CONNECTION",
				serverHost(ru), guessServerType(ru))}
	}
	names := make([]string, 0, len(matches))
	for _, m := range matches {
		names = append(names, m.name)
	}
	slices.Sort(names)
	return "", "", "", &cerr.CustomError{Title: "Ambiguous connection",
		Message: fmt.Sprintf("Connections %s all point to the server hosting %s; choose one with -c CONNECTION",
			strings.Join(names, ", "), repoURL)}
}

// Only used in messages, when no connection hosts the URL (otherwise, the connection's own server type is authoritative):
// a GitHub Enterprise server, or a Bitbucket proxy, would be wrongly guessed as Forgejo
func guessServerType(u *url.URL) ServerType {
	switch serverHost(u) {
	case "github.com":
		return ServerTypeGithub
	case "bitbucket.org":
		return ServerTypeBitBucket
	}
	return ServerTypeForgejo
}

// The server's host[:port], as seen by its users: default ports are dropped, and the API hosts of
// github.com and bitbucket.org (api.github.com, api.bitbucket.org) are folded into their web hosts
func serverHost(u *url.URL) string {
	host := strings.TrimPrefix(strings.ToLower(u.Hostname()), "www.")
	switch host {
	case "api.github.com":
		host = "github.com"
	case "api.bitbucket.org":
		host = "bitbucket.org"
	}
	if p := u.Port(); p != "" && !(p == "443" && u.Scheme == "https") && !(p == "80" && u.Scheme == "http") {
		host += ":" + p
	}
	return host
}

// The connection's path, minus any API suffix (Forgejo's /api/v1, GitHub Enterprise's /api/v3, Bitbucket's /2.0)
func serverPath(u *url.URL) string {
	p := strings.Trim(u.Path, "/")
	for _, suffix := range []string{"api/v1", "api/v3", "2.0"} {
		if p == suffix {
			return ""
		}
		p = strings.TrimSuffix(p, "/"+suffix)
	}
	return p
}
