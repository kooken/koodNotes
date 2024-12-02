package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNotes(file *os.File) {
	colors := GetColors()

	file.Seek(0, 0)                   // setting cursor at the beginning of the file
	scanner := bufio.NewScanner(file) // creating a scanner to read file line by line
	if !scanner.Scan() {              // if we can't scan the file (notebook is empty)
		PrintColoredText("Your notebook is empty 🥺 Please select '2' to add a new note!", colors.Red)
		return
	}
	fmt.Println("Your notes 📝:")
	fmt.Println(scanner.Text()) // printing first line
	for scanner.Scan() {        // printing other lines of the notebook
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // if error occured
		PrintColoredText("Error occured! Please, try again 🥺", colors.Red) // print error
	}
}

// Function for reading the last note
// to know what's the last existing number of the note
func ReadLastNote(file *os.File) ([]string, int, error) {
	var notes []string
	lastNoteNum := 0 // number of the last written note

	file.Seek(0, 0)                   // setting cursor at the beginning of the file
	scanner := bufio.NewScanner(file) // creating a scanner to read file line by line
	for scanner.Scan() {              // loop for going line by line
		line := scanner.Text()                  // current line
		parts := strings.SplitN(line, " - ", 2) // dividing a line in two parts(before [0] and after [1] '-')
		if len(parts) == 2 {                    // if note has got two parts (number and the string)
			currentNum, err := strconv.Atoi(parts[0]) // converting 00x (number of the note - [0]) to digit (num)
			if err == nil && currentNum > lastNoteNum {
				lastNoteNum = currentNum // updating a new lastNoteNum
			}
			notes = append(notes, line) // adding a current line to the slice of notes
		}
	}
	// checking for errors
	if err := scanner.Err(); err != nil {
		return nil, 0, err
	}
	return notes, lastNoteNum, nil
}
