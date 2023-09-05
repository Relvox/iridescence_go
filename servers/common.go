package servers

import (
	"fmt"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
	PATCH  HttpMethod = "PATCH"
)

func (m HttpMethod) validateBody(requireBody bool) error {
	switch m {
	case GET, DELETE:
		if requireBody {
			return fmt.Errorf("GET|DELETE cannot require body")
		}
		return nil
	case POST, PUT, PATCH:
		if !requireBody {
			return fmt.Errorf("POST|PUT|PATCH require body")
		}
		return nil
	}
	return fmt.Errorf("bad method: %s", m)
}
