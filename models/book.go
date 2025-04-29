
package models

import (
	"encoding/json"
)

// Book represents a book entry

type Book struct {
	ID   string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}