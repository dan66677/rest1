package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"main.go/internal/handlers"
	"main.go/pkg/logging"
)

const (
	userUrl  = "/users/:uuid"
	usersUrl = "/users"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersUrl, h.GetList)
	router.GET(userUrl, h.GetUserId)
	router.PUT(usersUrl, h.UpdateUser)
	router.POST(userUrl, h.Createuser)
	router.PATCH(userUrl, h.PartialUpdUser)
	router.DELETE(userUrl, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("list users"))
}

func (h *handler) GetUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("get user by user"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("update user"))
}
func (h *handler) Createuser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("create user"))
}
func (h *handler) PartialUpdUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("Partial update user"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("delete user"))
}
