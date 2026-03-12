package app

import (
	"fmt"
	"log"
	"os"
)

type CommandFlags struct {
	WriteToLog bool
	FilePath   string
}

func Cli(arguments []string) CommandFlags {
	fmt.Println("Hello from CLI")

	var chosenFlags CommandFlags

	for i := 0; i < len(arguments); i++ {
		argument := arguments[i]

		switch argument {
		default:
			log.Print("Did not recognise: ", argument)
		case "--output", "-o":
			if len(arguments[i:]) > 1 {
				chosenFlags.WriteToLog = true

				// Check if file already exists
				_, ErrFindingFile := os.Lstat(arguments[i+1])
				if ErrFindingFile != nil {

					// If it doesn't - create it
					_, ErrCreatingFile := os.Create(argument)
					if ErrCreatingFile != nil {
						log.Panic(ErrCreatingFile)
					}
				}

				chosenFlags.FilePath = arguments[i+1]
				i++
			} else {
				log.Panic("No output file specified")
			}
		}
	}

	return chosenFlags
}
