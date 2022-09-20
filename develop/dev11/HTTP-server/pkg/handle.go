package pkg

import (
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
)

// CreateEvent - обаботчик запроса "/create_event"
func (s *server.Storage) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка чтения тела запроса", err)
		JsonResponse(false, w, "request body read error", http.StatusInternalServerError)
		return
	}

	ok, reqBody := ValidatePostRequest(body, w)
	if ok != true {
		return
	}
	key := uuid.New().String()
	eventData := *reqBody
	s.RWMutex.RLock()
	s.Calendar[key] = eventData
	s.RWMutex.RUnlock()
	JsonResponse(true, w, "created", http.StatusCreated)
}

// DeleteEvent - обаботчик запроса "/Delete_event"
func (s *server.Storage) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка чтения тела запроса", err)
		JsonResponse(false, w, "request body read error", http.StatusInternalServerError)
		return
	}
	ok, reqBody := ValidatePostRequest(body, w)
	if ok != true {
		return
	}
	eventData := *reqBody

	ok = s.FindEventForDelete(eventData)
	if ok != true {
		JsonResponse(false, w, "no such events", http.StatusOK)
	} else {
		JsonResponse(true, w, "deleted", http.StatusOK)
	}

}

// UpdateEvent - обаботчик запроса "/update_event"
func (s *server.Storage) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка чтения тела запроса", err)
		JsonResponse(false, w, "request body read error", http.StatusInternalServerError)
		return
	}
	ok, reqBody := ValidatePostRequest(body, w)
	if ok != true {
		return
	}
	eventData := *reqBody

	ok = s.FindEventForUpdate(eventData)
	if ok != true {
		JsonResponse(false, w, "no such events", http.StatusOK)
	} else {
		JsonResponse(true, w, "updated", http.StatusOK)
	}

}

// EventsForDay - обаботчик запроса "/events_for_day"
func (s *server.Storage) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("invalid date:", err)
		JsonResponse(false, w, "invalid date", 400)
		return
	}
	events := s.FindEvents(id, dateTime, dateTime.AddDate(0, 0, 1))
	if len(events) == 0 {
		JsonResponse(false, w, "no such events", http.StatusOK)
	} else {
		JsonResponse(true, w, events, http.StatusOK)
	}
}

// EventsForWeek - обаботчик запроса "/events_for_week"
func (s *server.Storage) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("user_id")
	date := r.FormValue("date")
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("invalid date:", err)
		JsonResponse(false, w, "invalid date", 400)
		return
	}
	events := s.FindEvents(id, dateTime, dateTime.AddDate(0, 0, 7))

	if len(events) == 0 {
		JsonResponse(false, w, "no such events", http.StatusOK)
	} else {
		JsonResponse(true, w, events, http.StatusOK)
	}
}

// EventsForMonth - обаботчик запроса "/events_for_month"
func (s *server.Storage) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("user_id")
	date := r.FormValue("date")
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("invalid date:", err)
		JsonResponse(false, w, "invalid date", 400)
		return
	}
	events := s.FindEvents(id, dateTime, dateTime.AddDate(0, 1, 0))

	if len(events) == 0 {
		JsonResponse(false, w, "no such events", http.StatusOK)
	} else {
		JsonResponse(true, w, events, http.StatusOK)
	}
}
