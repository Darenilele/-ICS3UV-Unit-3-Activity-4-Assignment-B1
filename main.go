/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will if statements strings
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Variables
  var color string = ""

	// INPUT
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please choose a sweater colour from the available choices: Blue, Black, Red, White: ")
	input, _ := reader.ReadString('\n')
	color = strings.TrimSpace(input) // removes newline

	// PROCESSING
	// Statement checks for black or white color
	if color == "black" || color == "Black" || color == "white" || color == "White" {
		fmt.Println("You have enough sweaters in this colour.")
		// Statement checks for red color
	} else if color == "red" || color == "Red" {
		fmt.Println("This colour will look good on you.")
		fmt.Println("How about a pair of jeans to go with the sweater?")
		// Statement checks for blue color
	} else if color == "blue" || color == "Blue" {
		fmt.Println("This colour doesnt go well with your eyes.")
		// Statement checks for invalid color
	} else {
		fmt.Println("This is an invalid option")
	}

	fmt.Println("\nDone.")
}
