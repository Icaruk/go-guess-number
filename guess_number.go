package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func tryGuess(num int, guess int) bool {

	// Check if guess is correct
	if guess == num {
		fmt.Println("✅ Correcto")
		return true
	}

	// Get difference between guess and num
	diffA := guess - num
	diffB := num - guess

	// Get min value
	diff := math.Min(float64(diffA), float64(diffB))
	absoluteDiff := math.Abs(float64(diff))

	directionArrow := "↓"
	if guess < num {
		directionArrow = "↑"
	}

	// switch with 4 options
	directionArrows := ""
	switch {
	case absoluteDiff < 5:
		directionArrows = directionArrow
	case absoluteDiff >= 5:
		directionArrows = strings.Repeat(directionArrow, 2)
	case absoluteDiff >= 20:
		directionArrows = strings.Repeat(directionArrow, 3)
	}

	fmt.Println("❌ Incorrecto, inténtalo de nuevo", directionArrows)
	return false
}

func countdown() {
	rem := 5

	for {
		fmt.Printf("\rCerrando en... %d", rem)
		time.Sleep(time.Second * 1)
		rem--
	}
}

func main() {

	fmt.Println("Adivina un número entre el 1 y el 100")

	// Generate a random number between 1 and 1000
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1

	for {
		fmt.Println("Escribe el número: ")
		var guess int
		fmt.Scan(&guess)

		isValid := tryGuess(num, guess)

		if isValid {
			break
		}
	}

	go countdown()
	time.Sleep(time.Second * 5)

}
