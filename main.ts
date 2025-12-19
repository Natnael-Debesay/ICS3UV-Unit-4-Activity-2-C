/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-12-18
 * @fileoverview This program uses a While loop, with strings.
 */

// variables
let someString: string = "";
let someStringUpper: string = "";

// input some string from the user
someString = prompt("Enter a string (type 'exit' to quit): ") || "Nothing entered";
someStringUpper = someString.toUpperCase();

// keep looping until the user enters 'exit' to quit
while (someStringUpper != "EXIT") {
  // output the string
  console.log("You entered: " + someString);

  // ask for the next string
  someString = prompt("Enter a string (type 'exit' to quit): ") || "Nothing entered";
  someStringUpper = someString.toUpperCase();
}

console.log("\nDone.");
