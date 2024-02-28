package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TimeResponse struct {
	FormattedTime string `json:"time"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			currentTime := time.Now()
			formattedTime := currentTime.Format(time.RFC3339)

			timeResponse := TimeResponse{FormattedTime: formattedTime}

			jsonData, err := json.Marshal(timeResponse)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.Write(jsonData)
		}
	})

	fmt.Println("Сервер запущено")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
