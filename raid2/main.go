package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
)

func main() {

	arguments := os.Args

	StringArgs := arguments[1:]

	if len(arguments) != 10 {

		fmt.Println("Error")
	} else {

		var Buffer [9][9]int

		for x := 0; x < 9; x++ {
			for n := 0; n < 9; n++ {
				if StringArgs[x][n] == ' ' {
					continue
				}
				if StringArgs[x][n] == '.' {
					Buffer[x][n] = 0
				} else {

					Buffer[x][n] = int(rune(StringArgs[x][n]) - 48)

				}
			}
		}

		Gotoreturn(&Buffer)

		check := false

		for i := 0; i < 9; i++ {

			for x := 0; x < 9; x++ {

				for c := x + 1; c < 9; c++ {

					if Buffer[i][x] == Buffer[i][c] {
						check = true
					}
				}
			}
		}

		for x := 1; x <= 9; x++ {
			if len(arguments[x]) != 9 {

				check = true
			}
		}

		if check == true {
			fmt.Println("Error")
		} else {

			for t := 0; t < 9; t++ {
				for r := 0; r < 9; r++ {

					z01.PrintRune(rune(Buffer[t][r] + 48))

					if r != 8 {
						z01.PrintRune(' ')
					}

				}

				z01.PrintRune(10)

			}

		}

	}

}

func Gotoreturn(Buffer *[9][9]int) bool {

	if !Findzero(Buffer) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if Buffer[i][j] == 0 {
				for Maybe := 9; Maybe >= 1; Maybe-- {
					Buffer[i][j] = Maybe
					if Cheksrules(Buffer) {
						if Gotoreturn(Buffer) {
							return true
						}
						Buffer[i][j] = 0
					} else {
						Buffer[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func Findzero(Buffer *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if Buffer[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func Cheksrules(Buffer *[9][9]int) bool {

	//check duplicates by row
	for Yos := 0; Yos < 9; Yos++ {
		counter := [10]int{}
		for Xos := 0; Xos < 9; Xos++ {
			counter[Buffer[Yos][Xos]]++
		}
		if ItsCopi(counter) {
			return false
		}
	}

	//check duplicates by column
	for Yos := 0; Yos < 9; Yos++ {
		counter := [10]int{}
		for Xos := 0; Xos < 9; Xos++ {
			counter[Buffer[Xos][Yos]]++
		}

		if ItsCopi(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[Buffer[row][col]]++
				}
				if ItsCopi(counter) {
					return false
				}
			}
		}
	}

	return true
}

func ItsCopi(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
