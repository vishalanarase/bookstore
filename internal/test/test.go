package test

import (
	"fmt"
	"os"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/gorm"
)

var testFixtures *testfixtures.Loader

// InitialFixtureLoad returns a database object that loads the initial fixture
func InitialFixtureLoad(this *gorm.DB) {
	db, err := this.DB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fixturePath := "../../test/fixtures"

	testFixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(fixturePath),
	)
	if err != nil {
		panic(fmt.Sprintf("Unable to prepare testfixtures %s", err.Error()))
	}
}

// ResetDatabaseFixtures ensures that the database is reset
func ResetDatabaseFixtures(this *gorm.DB) {
	if testFixtures == nil {
		InitialFixtureLoad(this)
	}

	for i := 0; i < 10; i++ {
		err := testFixtures.Load()
		if err != nil {
			fmt.Printf("Unable to load fixtures %s\n", err.Error())
			continue
		}
		break
	}
}
