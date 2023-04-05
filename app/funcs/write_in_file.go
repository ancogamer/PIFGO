package funcs

import (
	"log"
	"os"
)

// WriteLocalFile writes a file to the local folder of storage
func WriteLocalFile(filePath string, data []byte) (err error) {
	err = os.WriteFile(filePath, data, 0777) // TODO check if is the best aproach, by now it works
	if err != nil {
		log.Println(err)
	}
	return
}

// ReadLocalFile writes a file to the local folder of storage
func ReadLocalFile(filePath string) (b []byte, err error) {
	b, err = os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
