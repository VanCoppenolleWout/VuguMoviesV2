package handlers

import (
	"github.com/vugu-examples/taco-store/internal/memstore"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

type MovieAPIHandler struct {
	Store *memstore.MemStore
	*httprouter.Router
}

func NewMovieAPIHandler(mem *memstore.MemStore) *MovieAPIHandler {
	h := &MovieAPIHandler{
		Store:  mem,
		Router: httprouter.New(),
	}
	h.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	h.Router.GET("/query", h.GetMovieList)

	return h
}

// GetMovieList gets taco list
func (h *MovieAPIHandler) GetMovieList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	list := h.Store.SelectTacoList()
	json.NewEncoder(w).Encode(list)
}
