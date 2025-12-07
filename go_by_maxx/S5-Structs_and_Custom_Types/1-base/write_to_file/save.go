package writetofile

import (
	"errors"
	"fmt"
	"os"
)

func WriteUserDataToFile(filename string, data string) error {
	if data == "" {
		return errors.New("data is empty")
	}

	dataToWrite := fmt.Sprint(data)
	os.WriteFile(filename, []byte(dataToWrite), 0644)
	return nil
}
