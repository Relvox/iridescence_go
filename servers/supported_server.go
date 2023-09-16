package servers

import (
	"github.com/gorilla/mux"
)

type SupportedServer interface {
	RegisterRoutes(r *mux.Router)
}
