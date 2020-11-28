package lib

import "github.com/imroc/req"

var httpRequester *req.Req

// GetHTTPRequester provide http request object for external calls
var GetHTTPRequester = func() *req.Req {
	if httpRequester == nil {
		httpRequester = req.New()
	}
	return httpRequester
}
