// stubber
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/git/localrepo.go
// Original timestamp : 2026.10.22 03:27:41

package git

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"stubber/conn_cfg"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
)

// Checked before anything gets created on the forge: we need the git binary, and a directory that is not already a repo
func checkLocalPrerequisites(dir string) *cerr.CustomError {
	if _, err := exec.LookPath("git"); err != nil {
		return &cerr.CustomError{Title: "git not found", Message: "The git binary is needed to commit and push the scaffolded code"}
	}
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		return &cerr.CustomError{Title: "Already a git repository", Message: fmt.Sprintf("%s is already a git repository; not creating a new one", dir)}
	} else if !errors.Is(err, fs.ErrNotExist) {
		return &cerr.CustomError{Title: "Unable to check the project directory", Message: err.Error()}
	}
	// The commit needs an identity; git var fails the same way git commit would
	if out, err := exec.Command("git", "var", "GIT_COMMITTER_IDENT").CombinedOutput(); err != nil {
		return &cerr.CustomError{Title: "No git identity",
			Message: fmt.Sprintf("git cannot commit without an identity; set it with: git config --global user.name \"Your Name\" && git config --global user.email you@example.com\n%s",
				strings.TrimSpace(string(out)))}
	}
	return nil
}

// git init on the initial branch, commit everything, then push to the freshly created (empty) remote
func commitAndPush(dir, remoteURL string, ac conn_cfg.ActiveConnection) *cerr.CustomError {
	steps := [][]string{
		{"init", "--initial-branch=" + initialBranch},
		{"add", "--all"},
		{"commit", "--quiet", "--message", initialCommitMessage},
		{"remote", "add", "origin", remoteURL},
	}
	for _, args := range steps {
		if err := runGit(dir, nil, args...); err != nil {
			return err
		}
	}
	return runGit(dir, pushAuthEnv(ac), "push", "--quiet", "--set-upstream", "origin", initialBranch)
}

func runGit(dir string, env []string, args ...string) *cerr.CustomError {
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	// Never hang on a credentials prompt: if our credentials are refused, we want an error
	cmd.Env = append(append(os.Environ(), "GIT_TERMINAL_PROMPT=0"), env...)
	if out, err := cmd.CombinedOutput(); err != nil {
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			msg = err.Error()
		}
		return &cerr.CustomError{Title: "git " + args[0] + " failed", Message: msg}
	}
	return nil
}

// The push is authenticated with the connection's credentials, sent as an HTTP header
// They go through the environment (GIT_CONFIG_*): never on the command line, and never written in .git/config
func pushAuthEnv(ac conn_cfg.ActiveConnection) []string {
	user, secret := gitCredentials(ac)
	header := "Authorization: Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+secret))
	return []string{"GIT_CONFIG_COUNT=1", "GIT_CONFIG_KEY_0=http.extraHeader", "GIT_CONFIG_VALUE_0=" + header}
}

// What each forge expects as the HTTPS username when a token is used as password
func gitCredentials(ac conn_cfg.ActiveConnection) (user, secret string) {
	if ac.Auth == conn_cfg.AuthBasic {
		return ac.Username, ac.Password
	}
	switch ac.ServerType {
	case conn_cfg.ServerTypeGithub:
		return "x-access-token", ac.Token
	case conn_cfg.ServerTypeBitBucket:
		if isAtlassianAPIToken(ac.Token) {
			return "x-bitbucket-api-token-auth", ac.Token
		}
		return "x-token-auth", ac.Token
	default:
		// Forgejo authenticates on the token alone; the username is not checked
		if ac.Username != "" {
			return ac.Username, ac.Token
		}
		return "oauth2", ac.Token
	}
}
