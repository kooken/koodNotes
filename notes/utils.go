package notes

import (
	"fmt"
	"os"
)

func InitializeCollection() (string, *os.File) {
	// Checking the number of arguments
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments! Try again 🥺")
		os.Exit(1) // exit with code 1 - exit with error
	}

	collectionName := os.Args[1] // first argument

	if collectionName == "help" { // if first argument is "help"
		fmt.Println("Create or manage a collection of notes by typing the collection name.") // short message
		fmt.Println("Usage: ./notestool [COLLECTION_NAME]")                                  // usage example - that user needs to type collection name
		os.Exit(0)                                                                           // exit with code 0 - exit with no errors
	}

	/*
	   Checking if file with given name is existed
	   	os.O_RDWR - read and write to the file if found
	   	os.O_CREATE - will create the file with the given name IF there isn't such a file
	   	0666 - access rights to the file - will be accessable for read and write to all users
	   	0 - prefix
	   	6 stands for read and write, 666 for owner, group & others
	*/

	file, err := os.OpenFile(collectionName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil { // if err == nil - operation was successful (err is nil - empty)
		fmt.Printf("Error opening or creating collection file: %v\n", err) // will print error
		os.Exit(1)                                                         // exit with code 1 - exit with error
	}

	return collectionName, file
}
