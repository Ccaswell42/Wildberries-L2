package server

import (
	"io"
	"log"
	"net/http"
	"time"
)

//func Handler2(w http.ResponseWriter, r *http.Request) {
//	time.Sleep(1 * time.Second)
//	fmt.Fprintln(w, "Hello, Dima, your URL:", r.URL)
//}

// дописать коды ошибок

func (s *Storage) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ошибка чтения тела запроса", err)
	}

	ok, reqBody := ValidatePostRequest(body, w)
	if ok != true {
		return
	}

	eventData := *reqBody
	s.RWMutex.RLock()
	s.Calendar[reqBody.Date] = eventData
	s.RWMutex.RUnlock()
	JsonResponse(true, w, "created", http.StatusCreated)
}

func (s *Storage) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("user_id")
	date := r.FormValue("date")
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("invalid date:", err)
		JsonResponse(false, w, "invalid date", 400)
		return
	}
	s.RLock()
	c := s.Calendar[date]
	s.RUnlock()
	if c.Id == id {
		JsonResponse(true, w, c, http.StatusOK)
	} else {
		JsonResponse(false, w, "no such event", http.StatusOK)
	}
}

func (s *Storage) EventsForWeek(w http.ResponseWriter, r *http.Request) {
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
	events := s.FindEventsForWeek(id, dateTime, dateTime.AddDate(0, 0, 7))

	if len(events) == 0 {
		JsonResponse(false, w, "no such events", http.StatusOK)
	} else {
		JsonResponse(true, w, events, http.StatusOK)
	}
}
