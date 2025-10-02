package services

import (
	"fmt"
	"testing"
)

func TestModpackCategories(t *testing.T) {
	categories, err := ModpackCategories()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(categories)
}

func TestModpacks(t *testing.T) {
	// modpacks, err := Modpacks()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Print(modpacks)
}
