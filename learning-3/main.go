package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go-note-json/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(*title, *content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Note saved successfully.")
}

func getNoteData() (*string, *string) {
	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) *string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return nil
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return &text
}
