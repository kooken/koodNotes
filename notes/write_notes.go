package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func AddNote(file *os.File) {

	colors := GetColors()

	_, lastNoteNum, err := ReadLastNote(file) // reading file to know the last note number
	if err != nil {                           // if error occured
		PrintColoredText("Error occured! Please, try again 🥺", colors.Red)
		return // exit
	}

	reader := bufio.NewReader(os.Stdin)
	var noteText string
	for { // Infinite loop till the user type something in a note
		fmt.Println("Enter a new note 📝: ")
		noteText, _ = reader.ReadString('\n')  // Read input until the user presses Enter
		noteText = strings.TrimSpace(noteText) // Trimming spaces from the input

		if noteText != "" {
			break // exit from the loop if the note is not empty
		}

		PrintColoredText("Note cannot be empty 🚫 Please, type something in a note 🥺", colors.Red)
	}

	newNum := lastNoteNum + 1 // number of the new note
	// creating a timestamp when creating a new note (using default go value for timestamp)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	newNote := fmt.Sprintf("%03d - %s (created on %s)\n", newNum, noteText, currentTime) // formatting
	_, err = file.WriteString(newNote)                                                   // adding a new note in the end of the file
	if err != nil {
		PrintColoredText("Error occured! Please, try again 🥺", colors.Red)
	} else {
		PrintColoredText("Note added successfully ✅", colors.Green)
	}
}
