package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Error - стркутура для отправки сообщения об ошибке
type Error struct {
	Data interface{} `json:"error"`
}

// Result - структура для отправки результата
type Result struct {
	Data interface{} `json:"result"`
}

// JsonResponse - отправляет ответ пользователю
func JsonResponse(ok bool, w http.ResponseWriter, data interface{}, statusCode int) {
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if !ok {
		errResponse := Error{data}
		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			log.Println("encode error", err)
			http.Error(w, "encode error", 500)
		}
	} else {
		resultResponse := Result{data}
		err := json.NewEncoder(w).Encode(resultResponse)
		if err != nil {
			log.Println("encode error", err)
			http.Error(w, "encode error", 500)
		}
	}
}

// ValidatePostRequest - валидатор параметров Post запросов
func ValidatePostRequest(body []byte, w http.ResponseWriter) (bool, *Event) {

	var reqBody Event

	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Println("ошибка сериализации JSON:", err)
		JsonResponse(false, w, "invalid parameters", 400)
		return false, nil
	}

	_, err = time.Parse("2006-01-02", reqBody.Date)
	if err != nil {
		log.Println("invalid date:", err)
		JsonResponse(false, w, "invalid date", 400)
		return false, nil
	}
	return true, &reqBody

}
