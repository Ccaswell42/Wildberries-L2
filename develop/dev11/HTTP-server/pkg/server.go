package pkg

import (
	"log"
	"net/http"
	"time"
)

func LaunchServer(conf *Config, server *http.Server) error {
	// создаем хранилище для событий
	storage := NewStorage()

	mux := http.NewServeMux() // мультиплексор (роутер)

	// обработчики методов
	mux.Handle("/create_event", MiddlewareLogger(storage.CreateEvent))
	mux.Handle("/events_for_day", MiddlewareLogger(storage.EventsForDay))
	mux.Handle("/events_for_week", MiddlewareLogger(storage.EventsForWeek))
	mux.Handle("/events_for_month", MiddlewareLogger(storage.EventsForMonth))
	mux.Handle("/delete_event", MiddlewareLogger(storage.DeleteEvent))
	mux.Handle("/update_event", MiddlewareLogger(storage.UpdateEvent))

	server.Addr = conf.Port
	server.Handler = mux
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second

	log.Println("starting server at", conf.Port)
	return server.ListenAndServe()
}
