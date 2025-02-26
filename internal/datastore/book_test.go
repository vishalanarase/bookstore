package datastore

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/vishalanarase/bookstore/internal/configs"
	"github.com/vishalanarase/bookstore/internal/test"
	"gorm.io/gorm"
)

var (
	store = &Store{}
	db    = &gorm.DB{}
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	err := os.Setenv("API_ENV", "test")
	if err != nil {
		log.Fatal(err, "Failed to api env to test")
	}

	config := configs.Config("../../")
	db, err = configs.DatabaseConnection(config)
	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	store = NewStore(db)
	test.ResetDatabaseFixtures(db)

	os.Exit(m.Run())
}

func TestNewBookModel(t *testing.T) {
	g := NewWithT(t)
	bm := NewBookStore(db)

	g.Expect(bm.GetDatabaseObject()).NotTo(BeNil())
}

func TestBookList(t *testing.T) {
	g := NewWithT(t)
	test.ResetDatabaseFixtures(db)

	books, err := store.Book.List(&gin.Context{})
	g.Expect(err).To(BeNil())

	g.Expect(len(books)).To(Equal(5))
}

func TestBookGet(t *testing.T) {
	g := NewWithT(t)
	test.ResetDatabaseFixtures(db)

	expected := Book{
		ID:        "550e8400-e29b-41d4-a716-446655440003",
		Title:     "The Catcher in the Rye",
		Author:    "J.D. Salinger",
		Publisher: "Little, Brown and Company",
		ISBN:      "9780316769488",
		Year:      1951,
		Edition:   1,
	}

	book, err := store.Book.Get(&gin.Context{}, "550e8400-e29b-41d4-a716-446655440003")
	g.Expect(err).To(BeNil())

	g.Expect(book.ID).To(Equal("550e8400-e29b-41d4-a716-446655440003"))
	g.Expect(expected.ID).To(Equal(book.ID))
	g.Expect(expected.Title).To(Equal(book.Title))
	g.Expect(expected.Author).To(Equal(book.Author))
	g.Expect(expected.Publisher).To(Equal(book.Publisher))
	g.Expect(expected.ISBN).To(Equal(book.ISBN))
	g.Expect(expected.Year).To(Equal(book.Year))
	g.Expect(expected.Edition).To(Equal(book.Edition))
}
