package pkg

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Event - структура события
type Event struct {
	Id          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

// Storage - хранилище событий
type Storage struct {
	Calendar map[string]Event
	*sync.RWMutex
}

// NewStorage - конструктор хранилища
func NewStorage() *Storage {
	calendar := make(map[string]Event)
	return &Storage{calendar, &sync.RWMutex{}}
}

// FindEvents - метод поиска событий по дате и id
func (s *Storage) FindEvents(id string, start, end time.Time) []Event {
	var arr []Event
	s.RLock()
	for _, val := range s.Calendar {
		dateTime, err := time.Parse("2006-01-02", val.Date)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(dateTime, start, end)
		if (dateTime == start || dateTime.After(start)) && dateTime.Before(end) {
			fmt.Println(dateTime, start, end)
			if val.Id == id {
				arr = append(arr, val)
			}
		}
	}
	s.RUnlock()
	return arr
}

// FindEventForDelete - удаляет событие из хранилиша
func (s *Storage) FindEventForDelete(event Event) bool {
	s.RLock()
	for key, val := range s.Calendar {
		if event == val {
			delete(s.Calendar, key)
			return true
		}
	}
	s.RUnlock()
	return false
}

// FindEventForUpdate - обновляет событие в хранилище
func (s *Storage) FindEventForUpdate(event Event) bool {
	s.RLock()
	for key, val := range s.Calendar {
		if val.Id == event.Id && val.Date == event.Date {
			s.Calendar[key] = event
			return true
		}
	}
	s.RUnlock()
	return false
}
