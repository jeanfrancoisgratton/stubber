// stubber
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/git/types.go
// Original timestamp : 2026.10.22 03:27:41

package git

import "time"

// Fast-fail: protects the connect / TLS handshake / response headers waits
var FastFailTimeout = 30 * time.Second

const userAgent = "stubber"

// The branch the scaffolded code is pushed to; it becomes the repo's default branch
const initialBranch = "develop"

const initialCommitMessage = "Initial commit: project scaffolded by stubber"
