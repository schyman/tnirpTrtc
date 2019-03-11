package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database connection Object
var db *sql.DB

// Custom error to inform that chapter requested is not found
type errChapterNotFound string

func (e errChapterNotFound) Error() string {
	return string(e)
}

// App version to marketing name mapping
var appVersionMap = map[string]string{
	"8.0":     "CS6",
	"11.0":    "CC 2015",
	"12.0":    "CC 2017",
	"MISSING": "-",
}

// Initialize the database connection
func startDatabase(config DatabaseConfig) {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		config.User, config.Password, config.Name, config.Host, config.Port))

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to postgres database: " + config.Name)
}

// Fetch details for a chapter with given id
func getChapter(chapterId int) (*Chapter, error) {
	row := db.QueryRow(GetChapterQuery, chapterId)
	var chapter Chapter
	err := row.Scan(&chapter.Company, &chapter.Project, &chapter.Chapter)

	// If database returns nothing, we return chapter not found error
	if err == sql.ErrNoRows {
		return nil, errChapterNotFound("Chapter not found")
	} else if err != nil {
		return nil, err
	}

	// We also fetch the list of chapter versions from database
	chapter.Versions, err = getChapterVersions(chapterId)
	if err != nil {
		return nil, err
	}

	return &chapter, nil
}

// Fetch versions of chapter from database and add it to the chapter object passed as input
func getChapterVersions(chapterId int) ([]ChapterVersion, error) {
	rows, err := db.Query(ChapterVersionsQuery, chapterId)
	if err != nil {
		return nil, err
	}

	var chapterVersions []ChapterVersion

	for rows.Next() {
		version := ChapterVersion{}
		var appVersion sql.NullString
		err = rows.Scan(
			&version.ChapterVersionId,
			&version.CreatedBy,
			&version.ChapterVersionNumber,
			&version.Created,
			&appVersion)

		if err != nil {
			return nil, err
		}

		// Convert appversion value into marketing friendly label
		version.Appversion = appVersionMap["MISSING"]
		if appVersion.Valid {
			if marketingName, ok := appVersionMap[appVersion.String]; ok {
				version.Appversion = marketingName
			}
		}

		chapterVersions = append(chapterVersions, version)
	}

	return chapterVersions, nil
}
