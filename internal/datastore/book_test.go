package datastore

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	setup "github.com/vishalanarase/bookstore/internal/config"
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

	config := setup.Config("../../")
	db, err = setup.DatabaseConnection(config)
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

	g.Expect(len(books)).To(Equal(2))
}

func TestBookDelete(t *testing.T) {
	g := NewWithT(t)
	test.ResetDatabaseFixtures(db)

	book, err := store.Book.Get(&gin.Context{}, "5317ab5c-3480-451d-ad0a-adee2ba07ca9")
	g.Expect(err).To(BeNil())

	g.Expect(book.ID).To(Equal("5317ab5c-3480-451d-ad0a-adee2ba07ca9"))
	g.Expect(book.Name).To(Equal("My Book"))
	g.Expect(book.Authorname).To(Equal("Vishal Anarase"))
	g.Expect(book.Rating).To(Equal(2))
}
