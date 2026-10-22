// stubber
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/git/createrepo.go
// Original timestamp : 2026.10.22 03:27:41

package git

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"stubber/conn_cfg"
	"stubber/helpers"

	cerr "github.com/jeanfrancoisgratton/customError/v3"
	gol "github.com/jeanfrancoisgratton/gitoperationslib"
	hftx "github.com/jeanfrancoisgratton/helperFunctions/v5/terminalfx"
)

// Everything needed to create the repository, gathered and checked by PrepareRepo() before the project is scaffolded
type RepoPlan struct {
	connName    string
	conn        conn_cfg.ActiveConnection
	client      gol.Client
	owner, name string
	dir         string // the project root, where the scaffolded code lands
}

// PrepareRepo checks everything that can be checked without side effects, before the project gets scaffolded:
// the --url and the connection to use, the forge's credentials, that the repository does not already exist,
// and that git can commit in the project directory. What can still fail after the scaffolding is then limited to
// the repository creation itself (ex: no rights in the owner organization), and the push
//
// The connection (and thus the backend: forgejo, github, bitbucket) is the one whose server hosts the URL; -c picks one
// when several do
func PrepareRepo(softname string) (*RepoPlan, *cerr.CustomError) {
	ctx := context.Background()
	p := &RepoPlan{}
	var ce *cerr.CustomError

	if p.connName, p.owner, p.name, ce = conn_cfg.ConnectionForURL(helpers.Url, helpers.ConnectionName); ce != nil {
		return nil, ce
	}
	if p.conn, ce = conn_cfg.ResolveConnection(p.connName); ce != nil {
		return nil, ce
	}

	// Same rule as createAssets.CreateStub(): the default "." means ./SOFTWARENAME
	dir := helpers.RootDir
	if dir == "." {
		dir = softname
	}
	var err error
	if p.dir, err = filepath.Abs(dir); err != nil {
		return nil, &cerr.CustomError{Title: "Unable to locate the project directory", Message: err.Error()}
	}
	if ce = checkLocalPrerequisites(p.dir); ce != nil {
		return nil, ce
	}

	// Contacting the forge validates the credentials (and on Forgejo, the server version)
	if p.client, ce = connect(ctx, p.conn); ce != nil {
		return nil, ce
	}
	_, gerr := p.client.Repos().Get(ctx, p.owner, p.name)
	switch {
	case gerr == nil:
		return nil, &cerr.CustomError{Title: "Repository already exists",
			Message: fmt.Sprintf("%s/%s already exists on %s; nothing was created", p.owner, p.name, p.connName)}
	case !gol.IsNotFound(gerr):
		return nil, forgeError(fmt.Sprintf("Could not check whether %s/%s exists on %s", p.owner, p.name, p.connName), gerr)
	}
	return p, nil
}

// Create creates the repository on the forge, then commits the scaffolded project and pushes it on the develop branch
//
// The forge repo is created empty: the scaffold already has its README.md and .gitignore, and an
// initial commit on the forge would only conflict with ours. Our push then creates develop, which
// becomes the default branch
func (p *RepoPlan) Create() *cerr.CustomError {
	ctx := context.Background()

	// Always private: making the repo public is left to the user (principle of precaution)
	repo, gerr := p.client.Repos().Create(ctx, gol.CreateRepoOptions{Owner: p.owner, Name: p.name, Description: helpers.Description,
		Visibility: gol.VisibilityPrivate})
	if gerr != nil {
		ce := forgeError(fmt.Sprintf("Could not create repository %s/%s on %s", p.owner, p.name, p.connName), gerr)
		ce.Message += fmt.Sprintf("\nThe project is scaffolded in %s; create the repository by hand, then push it (see README.md, \"Recovering from a failed -G\")", p.dir)
		return ce
	}
	if !helpers.Quiet {
		fmt.Printf("%s %s has been %s on %s (%s)\n", hftx.EnabledSign("Repository"), hftx.Brown(repo.FullName),
			hftx.Green("created"), hftx.Green(p.connName), p.client.Backend())
	}

	remoteURL := repo.CloneURL
	if remoteURL == "" {
		remoteURL = helpers.Url
	}
	if ce := commitAndPush(p.dir, remoteURL, p.conn); ce != nil {
		ce.Message += fmt.Sprintf("\nThe repository %s exists on the forge; finish from %s by hand (see README.md, \"Recovering from a failed -G\")",
			repo.FullName, p.dir)
		return ce
	}

	// The first branch pushed to an empty repo normally becomes its default branch; we make sure of it
	if repo.DefaultBranch != initialBranch {
		if _, gerr = p.client.Repos().Edit(ctx, p.owner, repo.Name, gol.EditRepoOptions{DefaultBranch: gol.Ptr(initialBranch)}); gerr != nil && !helpers.Quiet {
			fmt.Fprintln(os.Stderr, hftx.WarningSign(fmt.Sprintf("Could not set %s as the default branch: %s", initialBranch, gerr)))
		}
	}

	if !helpers.Quiet {
		fmt.Printf("%s pushed to %s, branch %s\n", hftx.EnabledSign("Scaffolded code"), hftx.Brown(remoteURL), hftx.Green(initialBranch))
	}
	return nil
}
