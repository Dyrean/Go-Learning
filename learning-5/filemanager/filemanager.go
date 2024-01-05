package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm *FileManager) Read() ([]string, error) {
	fmt.Println("Reading file starting:", fm.InputFilePath)

	readFile, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Printf("Failed to open %v file\n", fm.InputFilePath)
		return nil, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		floatInput := fileScanner.Text()

		fileLines = append(fileLines, floatInput)
	}

	fmt.Printf("Reading file done: %v\n", fm.InputFilePath)
	return fileLines, nil
}

func (fm *FileManager) Write(data any) error {
	fmt.Println("Writing file starting:", fm.OutputFilePath)

	writeFile, err := os.Create(fm.OutputFilePath)

	if err != nil {
		fmt.Printf("Failed to create %v file\n", fm.OutputFilePath)
		return err
	}

	defer writeFile.Close()

	encoder := json.NewEncoder(writeFile)
	err = encoder.Encode(data)

	if err != nil {
		fmt.Println("Failed to convert data to JSON.")
		return err
	}

	fmt.Printf("Writing file done: %v\n", fm.OutputFilePath)
	return nil
}
