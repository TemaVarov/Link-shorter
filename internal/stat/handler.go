package stat

import (
	"go/adv-demo/configs"
	"go/adv-demo/pkg/middleware"
	"go/adv-demo/pkg/res"
	"net/http"
	"time"
)

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Configs        *configs.Config
}

type StatHandler struct {
	StatRepository *StatRepository
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps) {
	handler := StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetAll(), deps.Configs))
}

func (handler *StatHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != "month" && by != "day" {
			http.Error(w, "Invalid by param", http.StatusBadRequest)
			return
		}
		stats := handler.StatRepository.GetStats(by, from, to)
		res.Json(w, 200, stats)
	}
}
