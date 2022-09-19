package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"serv/pkg/server"
	"syscall"
	"time"
)

func main() {
	config, err := server.GetConfig()
	if err != nil {
		log.Fatal("error during config downloading: ", err)
	}

	serv := &http.Server{}

	go func() {
		err := server.LaunchServer(config, serv)
		if err != http.ErrServerClosed {
			log.Println(err)
		}
		fmt.Println("сервер успешно завершился")
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	err = serv.Shutdown(closeCtx)
	if err != nil {
		log.Println("fail shutdown:", err)
	}
	log.Println("Shutdown completed successfully")
}

/*	коды ошибок:
	400: bad request (неправильный запрос дата или id)
	503: Service Unavailable (ошибка бизнес логики, сервер не готов обработать данный запрос.)
	500: иные ошибки (внутренняя проблема сервера.)
*/
