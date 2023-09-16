package servers

import (
	"fmt"
)

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	DELETE  HttpMethod = "DELETE"
	PATCH   HttpMethod = "PATCH"
	HEAD    HttpMethod = "HEAD"
	OPTIONS HttpMethod = "OPTIONS"
	TRACE   HttpMethod = "TRACE"
	CONNECT HttpMethod = "CONNECT"
)

func (m HttpMethod) validateBody(requireBody bool) error {
	switch m {
	case GET, DELETE, HEAD, OPTIONS, TRACE, CONNECT:
		if requireBody {
			return fmt.Errorf("%s method cannot require a body", m)
		}
		return nil
	case POST, PUT, PATCH:
		if !requireBody {
			return fmt.Errorf("%s method requires a body", m)
		}
		return nil
	default:
		return fmt.Errorf("bad method: %s", m)
	}
}
