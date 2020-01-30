package eventStore

import (
	"log"
	"os"
)

const eventStoreFile = "eventStore.txt"

func WriteEvent(msg string) {
	f, err := os.OpenFile(eventStoreFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(msg + "\n"); err != nil {
		log.Println(err)
	}
}
