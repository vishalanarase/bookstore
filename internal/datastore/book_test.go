package datastore

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	setup "github.com/vishalanarase/bookstore/internal/config"
	"github.com/vishalanarase/bookstore/internal/test"
	"gorm.io/gorm"
)

var (
	modes = &Models{}
	db    = &gorm.DB{}
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	err := os.Setenv("API_ENV", "test")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to api env to test")
	}

	config := setup.Config("../../")
	db, err = setup.DatabaseConnection(config)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to database")
	}

	modes = NewModels(db)
	test.ResetDatabaseFixtures(db)

	os.Exit(m.Run())
}

func TestNewBookModel(t *testing.T) {
	g := NewWithT(t)
	bm := NewBookModel(db)

	g.Expect(bm.GetDatabaseObject()).NotTo(BeNil())
}

func TestBookList(t *testing.T) {
	g := NewWithT(t)
	test.ResetDatabaseFixtures(db)

	books, err := modes.Book.List(&gin.Context{})
	g.Expect(err).To(BeNil())

	g.Expect(len(books)).To(Equal(2))
}

func TestBookDelete(t *testing.T) {
	g := NewWithT(t)
	test.ResetDatabaseFixtures(db)

	book, err := modes.Book.Get(&gin.Context{}, "5317ab5c-3480-451d-ad0a-adee2ba07ca9")
	g.Expect(err).To(BeNil())

	g.Expect(book.ID).To(Equal("5317ab5c-3480-451d-ad0a-adee2ba07ca9"))
	g.Expect(book.Name).To(Equal("My Book"))
	g.Expect(book.Authorname).To(Equal("Vishal Anarase"))
	g.Expect(book.Rating).To(Equal(2))
}
