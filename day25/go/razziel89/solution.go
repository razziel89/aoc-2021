package main

import (
	"fmt"
	"log"
)

// tag::solution[]

func pretty(image []rune, lineLen int) {
	var zeroVal rune
	for idx, char := range image {
		if char == zeroVal || char == kindEmpty {
			char = '.'
		}
		fmt.Printf("%c", char)
		if (idx+1)%lineLen == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func step(image []rune, lineLen int) ([]rune, bool) {
	var zeroVal rune
	length := len(image)
	next := make([]rune, len(image))
	moved := false
	// East first.
	for idx, char := range image {
		if char == kindEast {
			neigh := lineLen*(idx/lineLen) + (idx+1)%lineLen
			if image[neigh] == zeroVal {
				next[neigh] = char
				moved = true
			} else {
				next[idx] = char
			}
		}
	}
	// Then south.
	for idx, char := range image {
		if char == kindSouth {
			neigh := (idx + lineLen) % length
			if image[neigh] != kindSouth && next[neigh] != kindEast {
				next[neigh] = char
				moved = true
			} else {
				next[idx] = char
			}
		}
	}
	return next, moved
}

func main() {
	image, lineLen, err := ReadLinesAsImage()
	if err != nil {
		log.Fatal(err.Error())
	}

	count := 0
	fmt.Println(count)
	// pretty(image, lineLen)
	for moved := true; moved; {
		count++
		image, moved = step(image, lineLen)
		fmt.Println(count)
		// pretty(image, lineLen)
	}
	fmt.Println(count)
	_ = pretty
}

// end::solution[]
