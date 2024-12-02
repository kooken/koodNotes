package notes

import "fmt"

type Colors struct {
	Red   string
	Green string
	Reset string
}

// function to return struct with colors
func GetColors() Colors {
	return Colors{
		Red:   "\033[31m",
		Green: "\033[32m",
		Reset: "\033[0m", // colors reset
	}
}

// printing text and then reseting color to default
func PrintColoredText(text string, colorCode string) {
	fmt.Printf("%s%s\033[0m\n", colorCode, text)
}
