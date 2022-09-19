package server

import (
	"log"
	"sync"
	"time"
)

type Event struct {
	Id          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type Storage struct {
	Calendar map[string]Event
	*sync.RWMutex
}

func NewStorage() *Storage {
	calendar := make(map[string]Event)
	return &Storage{calendar, &sync.RWMutex{}}
}

// разобраться с бефор и афтер и все потестить. решить с ключом для мапы uuid google
func (s *Storage) FindEventsForWeek(id string, start, end time.Time) []Event {
	var arr []Event
	s.RLock()
	for _, val := range s.Calendar {
		dateTime, err := time.Parse("2006-01-02", val.Date)
		if err != nil {
			log.Println(err)
		}
		if (dateTime == start || dateTime.After(start)) && dateTime.Before(end) {
			if val.Id == id {
				arr = append(arr, val)
			}
		}
	}

	s.RUnlock()
	return arr
}
