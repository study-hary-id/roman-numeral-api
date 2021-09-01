package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		if urlPathElements[1] == "roman-numbers" {
			if len(urlPathElements) < 3 {
				errJson := errors{
					Status: http.StatusBadRequest,
					Title:  "Incomplete Endpoint",
					Detail: "'/roman-numbers' requires numeric parameter.",
				}

				js, err := json.Marshal(errorPayload{errJson})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				w.Write(js)

			} else {
				num, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
				w.Write([]byte(convertToRoman(num)))
			}

		} else {
			// Response failure if using random endpoints.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})

	// Create a server and run it on 8000 port.
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
