package servers

import (
	"fmt"
	"net/http"
)

type handleFunc[TIn, TOut any] struct {
	handler       func() (TOut, error)
	handlerV      func(vars map[string]string) (TOut, error)
	handlerR      func(request TIn) (TOut, error)
	handlerRV     func(request TIn, vars map[string]string) (TOut, error)
	decodeRequest func(r *http.Request) (TIn, error)
}

func (h *handleFunc[TIn, TOut]) validate() error {
	handlers := 0
	if h.handler != nil {
		handlers++
	}
	if h.handlerV != nil {
		handlers++
	}
	if h.handlerR != nil {
		handlers++
		if h.decodeRequest == nil {
			return fmt.Errorf("request handler must set request decoder")
		}
	}
	if h.handlerRV != nil {
		handlers++
		if h.decodeRequest == nil {
			return fmt.Errorf("request handler must set request decoder")
		}
	}
	if handlers == 1 {
		return nil
	}
	return fmt.Errorf("exactly one handler must be set")
}
