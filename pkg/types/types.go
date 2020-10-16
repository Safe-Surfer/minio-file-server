/*
	handle all types used by API
*/

package types

import (
	"net/http"
)

// Endpoints ...
// array for endpoint fields
type Endpoints []struct {
	EndpointPath string
	HandlerFunc  http.HandlerFunc
	HTTPMethods  []string
}

// HTTPMethod ...
// extra HTTP methods
type HTTPMethod string

// valid extra HTTPMethods
const (
	HTTPMethodList HTTPMethod = "LIST"
)
