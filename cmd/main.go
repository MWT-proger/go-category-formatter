package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MWT-proger/go-category-formatter/internal/category"
	"github.com/MWT-proger/go-category-formatter/internal/models"
)

func parseFlags(pathFile *string, pathOut *string) {
	flag.StringVar(pathFile, "file", "input.json", "путь к файлу с категориями")
	flag.StringVar(pathOut, "out", ".", "путь к месту куда выгрузить файлы")
	flag.Parse()
}

func main() {
	var result map[string][]string
	var pathFile string
	var pathOut string
	var ParentListCategory []*models.Category
	var ChildListCategory []*models.Category

	parseFlags(&pathFile, &pathOut)

	content, err := os.ReadFile(pathFile)

	if err != nil {
		fmt.Println("Error:", err)
		return

	}
	if err := json.Unmarshal(content, &result); err != nil {
		fmt.Println("Error:", err)
		return
	}
	// err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for parent, children := range result {
		p := category.GenerateCategory(parent)
		ParentListCategory = append(ParentListCategory, p)
		for _, child := range children {
			ch := category.GenerateCategory(child)
			ch.ParentID = p.ID.String()
			ChildListCategory = append(ChildListCategory, ch)
		}
	}
	category.SaveCategoryToFile(ParentListCategory, pathOut, "parent_category")
	category.SaveCategoryToFile(ChildListCategory, pathOut, "child_category")
	log.Println("Финиш генерации")

}
