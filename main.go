package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const backendURL = "test-back.iot-edge.svc.cluster.local/data" // 백엔드 URL

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 백엔드 호출
		resp, err := http.Get(backendURL)
		if err != nil {
			http.Error(w, "Failed to fetch data from backend", http.StatusInternalServerError)
			log.Printf("Error fetching backend data: %v", err)
			return
		}
		defer resp.Body.Close()

		// 응답 데이터 읽기
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read backend response", http.StatusInternalServerError)
			log.Printf("Error reading backend response: %v", err)
			return
		}

		// 값만 추출 (Value: 이후의 값)
		data := strings.TrimPrefix(string(body), "Value: ")

		// HTML 응답 생성
		html := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Frontend</title>
			</head>
			<body>
				<h1>Backend Value</h1>
				<p>The value is: <strong>%s</strong></p>
			</body>
			</html>
		`, data)

		// HTML 응답 전송
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(html))
	})

	// 프론트엔드 서버 실행
	port := "8081"
	log.Printf("Frontend server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
