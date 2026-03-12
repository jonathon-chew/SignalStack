package app

import (
	"log"
	"net/http"
	"os"
)

var commandFlags CommandFlags

// SetCommandFlags sets the flags returned by CLI so handlers can use them.
func SetCommandFlags(flags CommandFlags) {
	commandFlags = flags
}

func fourOhFour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, "./template/404.html")
}

func Log_message(message string) {
	// Back end log
	log.Print(message)

	if commandFlags.WriteToLog {
		os.WriteFile(commandFlags.FilePath, []byte(message), os.ModeAppend)
	}

	// Front end print
	// io.WriteString(w, message)
}
