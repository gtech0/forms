package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/section"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/model/form/section/block/question"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := ConvertToDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}
}

func ConvertToDSN() string {
	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s",
		hostname, port, user, pass, db, "sslmode=disable")

	return dsn
}

func Sync() {
	err := DB.AutoMigrate(
		&form.FormPattern{},
		&form.FormPublished{},
		&form.MarkConfiguration{},
		&form.FormGenerated{},
		&form.FormPublishedGroup{},
		&form.FormPublishedUser{},

		&section.Section{},

		//&block.Block{},
		&block.StaticBlock{},
		&block.DynamicBlock{},
		&block.Variant{},

		//&question.Question{},
		&question.Attachment{},

		&question.MultipleChoice{},
		&question.MultipleChoicePoints{},
		&question.MultipleChoiceOption{},

		&question.SingleChoice{},
		&question.SingleChoiceOption{},

		&question.TextInput{},
		&question.TextInputAnswer{},

		&question.Matching{},
		&question.MatchingPoint{},
		&question.MatchingTerm{},
		&question.MatchingDefinition{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
