package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"stubber/conn_cfg"
	"testing"
)

func TestCommitAndPush(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not installed")
	}
	// Isolate from the user's git configuration, and provide an identity for the commit
	t.Setenv("HOME", t.TempDir())
	t.Setenv("GIT_CONFIG_NOSYSTEM", "1")
	t.Setenv("GIT_AUTHOR_NAME", "stubber test")
	t.Setenv("GIT_AUTHOR_EMAIL", "test@example.net")
	t.Setenv("GIT_COMMITTER_NAME", "stubber test")
	t.Setenv("GIT_COMMITTER_EMAIL", "test@example.net")

	remote := filepath.Join(t.TempDir(), "remote.git")
	if out, err := exec.Command("git", "init", "--bare", "--quiet", remote).CombinedOutput(); err != nil {
		t.Fatalf("git init --bare: %s", out)
	}
	project := t.TempDir()
	if err := os.WriteFile(filepath.Join(project, "README.md"), []byte("# test\n"), 0644); err != nil {
		t.Fatal(err)
	}

	if ce := checkLocalPrerequisites(project); ce != nil {
		t.Fatal(ce.Error())
	}
	ac := conn_cfg.ActiveConnection{ServerType: conn_cfg.ServerTypeForgejo, Auth: conn_cfg.AuthToken, Token: "t"}
	if ce := commitAndPush(project, remote, ac); ce != nil {
		t.Fatal(ce.Error())
	}

	out, err := exec.Command("git", "-C", remote, "log", "--format=%s", initialBranch).CombinedOutput()
	if err != nil {
		t.Fatalf("git log on the remote: %s", out)
	}
	if got := strings.TrimSpace(string(out)); got != initialCommitMessage {
		t.Errorf("remote %s: got commit %q, want %q", initialBranch, got, initialCommitMessage)
	}
	// The credentials must never end up in the project's git config
	cfg, _ := os.ReadFile(filepath.Join(project, ".git", "config"))
	if strings.Contains(string(cfg), "Authorization") {
		t.Error("the push credentials were written in .git/config")
	}
	// A second run must refuse the existing repo
	if ce := checkLocalPrerequisites(project); ce == nil {
		t.Error("expected an error on an existing git repository")
	}
}

func TestGitCredentials(t *testing.T) {
	tests := []struct {
		ac         conn_cfg.ActiveConnection
		user, pass string
	}{
		{conn_cfg.ActiveConnection{ServerType: conn_cfg.ServerTypeGithub, Auth: conn_cfg.AuthToken, Token: "ghp"}, "x-access-token", "ghp"},
		{conn_cfg.ActiveConnection{ServerType: conn_cfg.ServerTypeBitBucket, Auth: conn_cfg.AuthToken, Token: "ATATTx"}, "x-bitbucket-api-token-auth", "ATATTx"},
		{conn_cfg.ActiveConnection{ServerType: conn_cfg.ServerTypeBitBucket, Auth: conn_cfg.AuthToken, Token: "ATCTTx"}, "x-token-auth", "ATCTTx"},
		{conn_cfg.ActiveConnection{ServerType: conn_cfg.ServerTypeForgejo, Auth: conn_cfg.AuthToken, Token: "tk", Username: "jfg"}, "jfg", "tk"},
		{conn_cfg.ActiveConnection{ServerType: conn_cfg.ServerTypeForgejo, Auth: conn_cfg.AuthBasic, Username: "jfg", Password: "pw"}, "jfg", "pw"},
	}
	for _, tt := range tests {
		if u, p := gitCredentials(tt.ac); u != tt.user || p != tt.pass {
			t.Errorf("%s/%v: got %s:%s, want %s:%s", tt.ac.ServerType, tt.ac.Auth, u, p, tt.user, tt.pass)
		}
	}
}
