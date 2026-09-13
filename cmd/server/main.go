package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func generateQuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var result strings.Builder

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if (row == 0 || row == y-1) && (col == 0 || col == x-1) {
				result.WriteRune('o')
			} else if row == 0 || row == y-1 {
				result.WriteRune('-')
			} else if col == 0 || col == x-1 {
				result.WriteRune('|')
			} else {
				result.WriteRune(' ')
			}
		}

		result.WriteRune('\n')
	}

	return result.String()
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))
	quad := r.URL.Query().Get("quad")

	if errX != nil || errY != nil {
		http.Error(w, "Invalid width or height", http.StatusBadRequest)
		return
	}

	if quad == "A" {
		fmt.Fprint(w, generateQuadA(x, y))
	}
}

func main() {
	files := http.FileServer(http.Dir("web"))

	http.HandleFunc("/api/generate", generateHandler)
	http.Handle("/", files)

	log.Println("Quad Generator running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}