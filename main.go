package main

import (
	"bufio"
	"fmt"
	"notes/notes"
	"os"
	"strings"
)

func main() {

	colors := notes.GetColors()

	collectionName, file := notes.InitializeCollection()
	defer file.Close() // close file after exit - delayed close of the file

	fmt.Printf("Collection '%s' is ready for use 🎉\n", collectionName)
	reader := bufio.NewReader(os.Stdin) // Create a reader for standard input.

	// Infinite loop for showing menu until user type "4" -> return
	for {
		// loop for selecting operation
		for {
			fmt.Printf("Select operation (1/4):\n1. Show notes.\n2. Add a note.\n3. Delete a note.\n4. Exit.\n")
			operationInput, _ := reader.ReadString('\n')       // Read input until the user presses Enter
			operationInput = strings.TrimSpace(operationInput) // Trimming spaces from the input

			// Loop for selecting the operation. This will continue until the user provides valid input.
			if operationInput == "1" {
				notes.ReadNotes(file)
				break // go back to menu
			} else if operationInput == "2" {
				notes.AddNote(file)
				break // go back to menu
			} else if operationInput == "3" {
				notes.DeleteNote(file)
				break // go back to menu
			} else if operationInput == "4" {
				notes.PrintColoredText("Exiting 🦕 \nThank you for using our notes!", colors.Green)
				return // end of function -> back to terminal
			} else {
				notes.PrintColoredText("Invalid input 🚫 Please enter a number from 1️⃣ to 4️⃣\n", colors.Red)
			}
		}
	}
}
