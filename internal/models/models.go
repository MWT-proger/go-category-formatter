package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ParentID    string    `json:"parent"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
}
type Price struct {
	Days  int `json:"days"`
	Price int `json:"price"`
}

type Product struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Description string      `json:"description"`
	Adress      string      `json:"adress"`
	Images      []uuid.UUID `json:"images"`
	Prices      []*Price    `json:"prices"`
	CategoryID  uuid.UUID   `json:"category_id"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
}

func (d *Category) MarshalJSON() ([]byte, error) {
	type Alias Category
	return json.Marshal(&struct {
		*Alias
		ID          uuid.UUID `json:"id"`
		CreatedAt   string    `json:"created_at"`
		UpdatedAt   string    `json:"updated_at"`
		ParentID    string    `json:"parent"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Slug        string    `json:"slug"`
	}{
		Alias:       (*Alias)(d),
		ID:          d.ID,
		CreatedAt:   d.CreatedAt.Format(time.DateTime),
		UpdatedAt:   d.UpdatedAt.Format(time.DateTime),
		ParentID:    d.ParentID,
		Name:        d.Name,
		Description: d.Description,
		Slug:        d.Slug,
	})
}
