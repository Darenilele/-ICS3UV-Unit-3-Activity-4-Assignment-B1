/**
 * @author Daren Ileleji
 * @versopn 1.0.0
 * @date 2025-11-20
 * @fileoverview This program will if statements strings
 */


// Variables
let color: string = "";

// INPUT 
color = prompt("Please choose a sweater colour from the available choices: Blue, Black, Red, White.") || "none";

// PROCESSING
// Statement checks for black or white color
if (color == "black" || color == "Black" || color == "white" || color == "White") {
  console.log("You have enough sweaters in this colour.")
  // Statement checks for red color
} else if (color == "red" || color == "Red") {
  console.log("This colour will look good on you.")
  console.log("How about a pair of jeans to go with the sweater?")
  // Statement checks for blue color
} else if (color == "blue" || color == "Blue") {
  console.log("This colour doesnt go well with your eyes.")
  // Statement checks for invalid color
} else { console.log("This is an invalid option")
}

console.log("\nDone.")