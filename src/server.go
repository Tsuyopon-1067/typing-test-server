package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Server() {
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // CORS
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fmt.Fprintf(w, "hello world\n")

		if r.Method == http.MethodOptions {
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed.", http.StatusMethodNotAllowed)
			return
		}

		//ここでパース
		pathSegments := strings.Split(r.URL.Path, "/")
		urlOk := true
		id := ""
		userScore := RegisterScore{}
		if len(pathSegments) == 4 {
			if  pathSegments[2] != "" {
				id = pathSegments[2]
			} else {
				urlOk = false
			}
			if pathSegments[3] != "scores" {
				urlOk = false
			}
		} else {
			urlOk = false
		}

		fmt.Printf("urlOk: %t\n", urlOk)
		if err := json.NewDecoder(r.Body).Decode(&userScore); err != nil {
			fmt.Println("error:  47", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Printf("userScore\n%s\n", userScore.Text())

		if urlOk {
			fmt.Fprintf(w, "User ID: %s", id)
		} else {
			http.NotFound(w, r)
		}
	})


	currentTime := time.Now()
	fmt.Println("Formatted Time:", currentTime.Format("2006-01-02 15:04:05"))
	port := "8080"
	fmt.Printf("server: http://localhost:%s/\n", port);
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		fmt.Println(err)
	}
}
