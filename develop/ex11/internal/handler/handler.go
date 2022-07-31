package handler

import (
	"github.com/julienschmidt/httprouter"
)

type HandlerI interface {
	Register(route *httprouter.Router)
}
