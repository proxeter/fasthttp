// +build windows

package fasthttp

import "testing"

func TestURIPathNormalizeIssue86(t *testing.T) {
<<<<<<< HEAD
	t.Parallel()

	// see https://github.com/valyala/fasthttp/issues/86
=======
	// see https://github.com/trafficstars/fasthttp/issues/86
>>>>>>> rename go mod pkg name
	var u URI

	testURIPathNormalize(t, &u, `C:\a\b\c\fs.go`, `C:\a\b\c\fs.go`)
}
