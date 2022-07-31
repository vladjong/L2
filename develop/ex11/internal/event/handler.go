package event

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/vladjong/L2/develop/ex11/internal/handler"
	"github.com/vladjong/L2/develop/ex11/pkg/logging"
)

const (
	createEventURL   = "/create_event"
	updateEventURL   = "/update_event"
	deleteEventURL   = "/delete_event"
	eventForDayURL   = "/events_for_day"
	eventForWeekURL  = "/events_for_week"
	eventForMonthURL = "/events_for_month"
	dateExample      = "2006-01-02"
)

type Handler struct {
	logger *logging.Logger
	cache  EventCache
}

func NewHandler(logger *logging.Logger) handler.HandlerI {
	return &Handler{
		logger: logger,
		cache:  *NewEventCache(),
	}
}

func (h *Handler) Register(router *httprouter.Router) {
	router.POST(createEventURL, h.CreateEvent)
	router.POST(updateEventURL, h.UpdateEvent)
	router.POST(deleteEventURL, h.DeleteEvent)
	router.GET(eventForDayURL, h.GetEventForDay)
	router.GET(eventForWeekURL, h.GetEventForWeek)
	router.GET(eventForMonthURL, h.GetEventForMonth)
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user_id := r.Form.Get("user_id")
	date := r.Form.Get("date")
	time, _ := time.Parse(dateExample, date)
	if user_id == "" || date == "" {
		h.logger.Infof("element incorrect: %s, date: %s", user_id, date)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	h.cache.Create(Event{
		ID:   user_id,
		Date: time,
	})
	h.logger.Info("element added id: %s, date: %s", user_id, date)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user_id := r.Form.Get("user_id")
	date := r.Form.Get("date")
	time, _ := time.Parse(dateExample, date)
	if user_id == "" || date == "" {
		h.logger.Infof("element incorrect: %s, date: %s", user_id, date)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.cache.Update(Event{
		ID:   user_id,
		Date: time,
	})
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	h.logger.Info("element update: %s, date: %s", user_id, date)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user_id := r.Form.Get("user_id")
	date := r.Form.Get("date")
	time, _ := time.Parse(dateExample, date)
	if user_id == "" || date == "" {
		h.logger.Info("element incorrect: %s, date: %s", user_id, date)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.cache.Delete(Event{
		ID:   user_id,
		Date: time,
	})
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	h.logger.Infof("element delete: %s, date: %s", user_id, date)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetEventForDay(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dateStr := r.Form.Get("date")
	date, err := time.Parse(dateExample, dateStr)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	day := date.Day()
	event, err := h.cache.GetEventByDay(day)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	h.logger.Infof("event get for day: %s", day)
	w.Write(h.MakeJson(event))
}

func (h *Handler) GetEventForWeek(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dateStr := r.Form.Get("date")
	date, err := time.Parse(dateExample, dateStr)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	event, err := h.cache.GetEventByWeek(date)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	h.logger.Infof("event get for week: %s", date)
	w.Write(h.MakeJson(event))
}

func (h *Handler) GetEventForMonth(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dateStr := r.Form.Get("date")
	date, err := time.Parse(dateExample, dateStr)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	month := date.Month()
	event, err := h.cache.GetEventByMonth(month)
	if err != nil {
		h.logger.Info(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	h.logger.Infof("event get for month: %s", month)
	w.Write(h.MakeJson(event))
}

func (h *Handler) MakeJson(event []Event) []byte {
	json, _ := json.Marshal(event)
	return json
}
