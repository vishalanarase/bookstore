package test

import (
	"fmt"
	"os"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/gorm"
)

var testFixtures *testfixtures.Loader

func InitialFixtureLoad(this *gorm.DB) {
	db, err := this.DB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fixturePath := "../../test/fixtures"

	fmt.Println("Test Fixture PATH: ", fixturePath)
	fmt.Printf("DB: %+v\nSql DB %+v\n", this, db)

	testFixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(fixturePath),
	)
	if err != nil {
		panic(fmt.Sprintf("Unable to prepare testfixtures %s", err.Error()))
	}
}

func ResetDatabaseFixtures(this *gorm.DB) {
	if testFixtures == nil {
		InitialFixtureLoad(this)
	}

	err := testFixtures.Load()
	if err != nil {
		panic(fmt.Sprintf("Unable to load fixtures %s\n", err.Error()))
	}
}
