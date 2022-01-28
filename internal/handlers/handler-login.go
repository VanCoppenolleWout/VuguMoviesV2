package handlers

import (
	"github.com/vugu-examples/taco-store/internal/memstore"
	"github.com/julienschmidt/httprouter"
)

type LoginHandler struct {
	Store *memstore.MemStore
	*httprouter.Router
}

func NewLoginHandler() {

}

func PostLogin() {
	
}