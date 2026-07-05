// func grade(score int) string using switch with no condition: 90+ → "A", 80+ → "B", 70+ → "C", 60+ → "D", else "F". Return "invalid" if score < 0 or > 100. Done when: works for boundary values (89, 90, 100, -1, 101).

package main

import "fmt"

func grade(s int) string {
	switch {
		case s > 100:
			return "invalid"
		case s >= 90:
			return "A"
		case s >= 80:
			return "B"
		case s >= 70:
			return "C"
		case s >= 60:
			return "D"
		case s >= 0:
			return "F"
		default:
			return "invalid"
	}
}

func main() {
	fmt.Println(grade(-1))	
	fmt.Println(grade(89))	
	fmt.Println(grade(90))	
	fmt.Println(grade(100))	
	fmt.Println(grade(101))	

}


