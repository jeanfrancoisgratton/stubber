package conn_cfg

import (
	"os"
	"path/filepath"
	"testing"
)

func writeRegistry(t *testing.T, cf ConnectionsFileStruct) {
	t.Helper()
	home := t.TempDir()
	t.Setenv("HOME", home)
	if err := os.MkdirAll(filepath.Dir(connectionsFilePath()), 0755); err != nil {
		t.Fatal(err)
	}
	if err := cf.SaveConnectionsFile(); err != nil {
		t.Fatal(err)
	}
}

func TestConnectionForURL(t *testing.T) {
	// Two GitHub accounts: -c is mandatory for github.com URLs
	writeRegistry(t, ConnectionsFileStruct{Auths: map[string]ConnectionStruct{
		"forgejo":  {ServerType: ServerTypeForgejo, URL: "https://git.example.net:3000", Token: "t"},
		"subpath":  {ServerType: ServerTypeForgejo, URL: "https://example.com/forgejo/api/v1", Token: "t"},
		"gh-alice": {ServerType: ServerTypeGithub, URL: "https://api.github.com", Token: "t", Username: "alice"},
		"gh-bob":   {ServerType: ServerTypeGithub, URL: "https://github.com", Token: "t", Username: "bob"},
		"bb":       {ServerType: ServerTypeBitBucket, URL: "https://api.bitbucket.org/2.0", Token: "t"},
	}})

	tests := []struct {
		url, flag, conn, owner, repo string
		wantErr                      bool
	}{
		{url: "https://git.example.net:3000/jfg/myrepo", conn: "forgejo", owner: "jfg", repo: "myrepo"},
		{url: "https://git.example.net:3000/jfg/myrepo.git/", conn: "forgejo", owner: "jfg", repo: "myrepo"},
		{url: "https://example.com/forgejo/org/tool", conn: "subpath", owner: "org", repo: "tool"},
		{url: "https://github.com/someorg/thing", flag: "gh-bob", conn: "gh-bob", owner: "someorg", repo: "thing"},
		{url: "https://git.example.net:3000/mainline/stubber.git", flag: "forgejo", conn: "forgejo", owner: "mainline", repo: "stubber"},
		{url: "https://bitbucket.org/ws/thing", conn: "bb", owner: "ws", repo: "thing"},
		{url: "https://github.com/bob/thing", wantErr: true},                  // two GitHub connections, and no -c
		{url: "https://github.com/bob/thing", flag: "forgejo", wantErr: true}, // -c points to another server
		{url: "https://github.com/bob/thing", flag: "nope", wantErr: true},    // -c is not a connection
		{url: "https://git.example.net/jfg/myrepo", wantErr: true},            // port differs
		{url: "https://git.example.net:3000/jfg", wantErr: true},              // no repo name
		{url: "https://git.example.net:3000/a/b/c", wantErr: true},            // too deep
		{url: "git.example.net:3000/jfg/myrepo", wantErr: true},               // no scheme
		{url: "https://unknown.example.org/jfg/myrepo", wantErr: true},        // no connection
	}
	for _, tt := range tests {
		conn, owner, repo, err := ConnectionForURL(tt.url, tt.flag)
		if tt.wantErr {
			if err == nil {
				t.Errorf("%s: expected an error, got %s %s/%s", tt.url, conn, owner, repo)
			}
			continue
		}
		if err != nil {
			t.Errorf("%s: unexpected error: %s", tt.url, err.Error())
			continue
		}
		if conn != tt.conn || owner != tt.owner || repo != tt.repo {
			t.Errorf("%s: got %s %s/%s, want %s %s/%s", tt.url, conn, owner, repo, tt.conn, tt.owner, tt.repo)
		}
	}
}
