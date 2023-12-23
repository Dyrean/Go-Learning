package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n", n.Title, n.Content)
}

func (n *Note) Save() error {
	fileName := fmt.Sprintf("save\\%v.json", n.ID)

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid input")
	}

	return &Note{
		ID:        uuid.NewString(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
