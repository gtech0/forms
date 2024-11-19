package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/model/form/published"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := ConvertToDSN()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
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
		&published.FormPublished{},
		&published.MarkConfiguration{},
		&published.FormPublishedGroup{},
		&published.FormPublishedUser{},

		&generated.FormGenerated{},

		&pattern.FormPattern{},

		&section.Section{},

		&block.Block{},
		&block.StaticBlock{},
		&block.DynamicBlock{},
		&block.Variant{},

		&question.Attachment{},
		&model.File{},

		&question.Question{},

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
