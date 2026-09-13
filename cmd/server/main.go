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

func generateQuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var result strings.Builder

	if x == 1 {
		for i := 0; i < y; i++ {
			if i == 0 {
				result.WriteString("/")
			} else if i == y-1 {
				result.WriteString("\\")
			} else {
				result.WriteString("*")
			}
			result.WriteRune('\n')
		}
		return result.String()
	}

	for i := 0; i < y; i++ {
		if i == 0 {
			result.WriteString("/")

			for j := 0; j < x-2; j++ {
				result.WriteString("*")
			}

			result.WriteString("\\")
			result.WriteRune('\n')

		} else if i == y-1 {
			result.WriteString("\\")

			for j := 0; j < x-2; j++ {
				result.WriteString("*")
			}

			result.WriteString("/")
			result.WriteRune('\n')

		} else {
			result.WriteString("*")

			for j := 0; j < x-2; j++ {
				result.WriteString(" ")
			}

			result.WriteString("*")
			result.WriteRune('\n')
		}
	}

	return result.String()
}

func generateQuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var result strings.Builder

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				result.WriteRune('A')
			} else if row == 0 && col == x-1 {
				result.WriteRune('A')
			} else if row == y-1 && col == 0 {
				result.WriteRune('C')
			} else if row == y-1 && col == x-1 {
				result.WriteRune('C')
			} else if row == 0 {
				result.WriteRune('B')
			} else if row == 0 || row == y-1 {
				result.WriteRune('B')
			} else if col == 0 || col == x-1 {
				result.WriteRune('B')
			} else {
				result.WriteRune(' ')
			}
		}

		result.WriteRune('\n')
	}

	return result.String()
}

func generateQuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var result strings.Builder

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {

			if row == 0 {
				if col == 0 {
					result.WriteRune('A')
				} else if col == x-1 {
					result.WriteRune('C')
				} else {
					result.WriteRune('B')
				}

			} else if row == y-1 {
				if col == 0 {
					result.WriteRune('A')
				} else if col == x-1 {
					result.WriteRune('C')
				} else {
					result.WriteRune('B')
				}

			} else {
				if col == 0 || col == x-1 {
					result.WriteRune('B')
				} else {
					result.WriteRune(' ')
				}
			}
		}

		result.WriteRune('\n')
	}

	return result.String()
}

func generateQuadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var result strings.Builder

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 {
				if col == 0 {
					result.WriteRune('A')
				} else if col == x-1 {
					result.WriteRune('C')
				} else {
					result.WriteRune('B')
				}

			} else if row == y-1 {
				if col == 0 {
					result.WriteRune('C')
				} else if col == x-1 {
					result.WriteRune('A')
				} else {
					result.WriteRune('B')
				}

			} else {
				if col == 0 || col == x-1 {
					result.WriteRune('B')
				} else {
					result.WriteRune(' ')
				}
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
} else if quad == "B" {
	fmt.Fprint(w, generateQuadB(x, y))
} else if quad == "C" {
	fmt.Fprint(w, generateQuadC(x, y))
} else if quad == "D" {
	fmt.Fprint(w, generateQuadD(x, y))
} else if quad == "E" {
	fmt.Fprint(w, generateQuadE(x, y))
}
}

func main() {
	files := http.FileServer(http.Dir("web"))

	http.HandleFunc("/api/generate", generateHandler)
	http.Handle("/", files)

	log.Println("Quad Generator running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}