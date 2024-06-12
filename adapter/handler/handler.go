package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-michi/michi"
	"github.com/google/wire"
	"github.com/pranc1ngpegasus/go-playground/domain/model"
	"github.com/pranc1ngpegasus/go-playground/domain/usecase"
)

var _ http.Handler = (*Handler)(nil)

var NewHandlerSet = wire.NewSet(
	wire.Bind(new(http.Handler), new(*Handler)),
	NewHandler,
)

type Handler struct {
	mux http.Handler

	usecase usecase.Usecase[model.User, *model.User]
}

func NewHandler(
	usecase usecase.Usecase[model.User, *model.User],
) (*Handler, error) {
	router := michi.NewRouter()
	router.Use(middleware.Logger)

	h := &Handler{
		mux:     router,
		usecase: usecase,
	}

	router.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	router.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.PathValue("id")

		if _, err := h.usecase.Handle(ctx, model.User{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(fmt.Sprintf(("User ID: %s"), id)))
	})

	return h, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}
