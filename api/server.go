package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Status struct {
	Running  bool   `json:"running"`
	Version  string `json:"version"`
	Uptime   string `json:"uptime"`
}

func StartServer() {
	http.HandleFunc("/status", func(w http.ResponseWriter, _ *http.Request) {
		resp := Status{
			Running: true,
			Version: "1.0",
			Uptime:  "n/a",
		}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("🌐 API REST escuchando en http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
