package category

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/MWT-proger/go-category-formatter/internal/models"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

func GenerateCategory(name string) *models.Category {

	newCategory := models.Category{
		ID:        uuid.New(),
		Name:      name,
		Slug:      slug.Make(name),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &newCategory
}
func SaveCategoryToFile(data []*models.Category, fileName string) {
	b, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("JOPA")
	}

	os.WriteFile("data/"+fileName+".json", b, 0644)
}
