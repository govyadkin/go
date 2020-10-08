package workWithFiles

import (
	"io"
	"log"
	"os"
)

func Write(output []string, fileName string) error {
	if fileName != "" {
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		for _, item := range output {
			_, err = file.WriteString(item + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		for _, item := range output {
			_, err := io.WriteString(os.Stdout, item+"\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
