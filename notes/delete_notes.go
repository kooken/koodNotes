package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DeleteNote(file *os.File) {
	colors := GetColors()

	reader := bufio.NewReader(os.Stdin) // Create a reader for standard input.

	// infinite loop until the correct number of note or 0 is provided by user
	for {
		fmt.Print("Enter the number of the note to delete or 0 to cancel:\n")
		input, _ := reader.ReadString('\n')  // Read input until the user presses Enter
		input = strings.TrimSpace(input)     // Trimming spaces from the input
		noteNum, err1 := strconv.Atoi(input) // converting input string to number
		notes, _, err2 := ReadLastNote(file) // reading all notes

		// checking for validity
		if err1 != nil || err2 != nil || noteNum < 0 || noteNum > len(notes) {
			PrintColoredText("Invalid note number 🚫 Please, try again 🥺", colors.Red)
			continue // loop starts again in case of invalid number
		}

		if noteNum == 0 { // cancellation
			PrintColoredText("Deletion of the note is canceled 😔", colors.Red)
			return // go back to menu
		}

		// deleting chosen note and updating numbers
		var updatedNotes []string
		for i, note := range notes {
			if i+1 == noteNum {
				continue // skipping deleted note
			}
			newNum := len(updatedNotes) + 1         // updating note number
			parts := strings.SplitN(note, " - ", 2) // split string into two parts (number and the text)
			if len(parts) == 2 {
				// creating a new string for notes with updated number (notes after deleted one)
				updatedNotes = append(updatedNotes, fmt.Sprintf("%03d - %s", newNum, parts[1]))
			}
		}

		file.Truncate(0)                    // deleting all in file
		file.Seek(0, 0)                     // setting cursor at the beginning of the file
		for _, note := range updatedNotes { // loop for every note in a string array
			_, err := file.WriteString(note + "\n") // writing each note on a newline
			if err != nil {
				PrintColoredText("Error occured! Please, try again 🥺", colors.Red)
				return
			}
		}

		PrintColoredText("Note deleted successfully ✅", colors.Green)
		break // exit the loop after successful delete
	}
}
