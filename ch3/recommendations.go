package main

import (
	"sort"
)

type set map[Book]struct{}
type bookCollection map[Book]struct{}
type bookRecommendations map[Book]bookCollection

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}

// newCollection initialises a new bookCollection.
func newCollection() bookCollection {
	return make(bookCollection)
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendations)
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBookOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBookOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations
}

// listOtherBooksOnShelves returns the list of my books except the one at the given index.
func listOtherBooksOnShelves(bookIndex int, myBooks []Book) []Book {
	otherBooks := make([]Book, bookIndex, len(myBooks)-1)

	copy(otherBooks, myBooks[:bookIndex])
	otherBooks = append(otherBooks, myBooks[bookIndex+1:]...)

	return otherBooks
}

func registerBookRecommendations(recommendations bookRecommendations, reference Book, otherBooks []Book) {
	for _, book := range otherBooks {
		collection, ok := recommendations[reference]
		if !ok {
			collection = newCollection()
			recommendations[reference] = collection
		}
		collection[book] = struct{}{}
	}
}

// recommendBooks returns the list of recommended books for a reference book.
func recommendBooks(recommendations bookRecommendations, myBooks []Book) []Book {
	bc := make(bookCollection)

	myShelf := make(map[Book]bool)
	for _, myBook := range myBooks {
		myShelf[myBook] = true
	}

	for _, myBook := range myBooks {
		for recommendation := range recommendations[myBook] {
			if myShelf[recommendation] {
				continue
			}
			bc[recommendation] = struct{}{}
		}
	}

	recommendationsForABook := bookCollectionToListOfBooks(bc)
	return recommendationsForABook
}

// bookCollectionToListOfBooks transforms a bookCollection entities into a list of Book.
func bookCollectionToListOfBooks(bc bookCollection) []Book {
	bookList := make([]Book, 0, len(bc))
	for book := range bc {
		bookList = append(bookList, book)
	}

	sort.Slice(bookList, func(i, j int) bool {
		if bookList[i].Author != bookList[j].Author {
			return bookList[i].Author < bookList[j].Author
		}
		return bookList[i].Title < bookList[j].Title
	})
	return bookList
}
