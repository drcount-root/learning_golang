package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	reader := bufio.NewScanner(file)

	lines := []string{}

	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	err = reader.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}

func (fm *FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	time.Sleep(4 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("")
	}

	file.Close()
	return nil
}

// filemanager constructor function
func New (inputPath, OutputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: OutputFilePath,
	}
}