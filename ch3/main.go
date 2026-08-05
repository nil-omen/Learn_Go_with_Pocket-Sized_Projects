package main

import (
	"fmt"
	"os"
)

func main() {
	bookworms, err := loadBookworms("./testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Some books in common are:")
	commonBooks := findCommonBooks(bookworms)
	displayBooks(commonBooks)
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
