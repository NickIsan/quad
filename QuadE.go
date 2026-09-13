package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {

			if row == 0 {
				if col == 0 {
					z01.PrintRune('A')
				} else if col == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}

			} else if row == y-1 {
				if col == 0 {
					z01.PrintRune('C')
				} else if col == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}

			} else {
				if col == 0 || col == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}

		z01.PrintRune('\n')
	}
}