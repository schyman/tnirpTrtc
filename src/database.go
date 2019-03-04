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
var APP_VERSION_MAP = map[string]string{
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
func getChapter(chapter *Chapter, chapter_id int) error {
	row := db.QueryRow(GET_CHAPTER_QUERY, chapter_id)
	err := row.Scan(&chapter.Company, &chapter.Project, &chapter.Chapter)

	// If database returns nothing, we return chapter not found error
	if err == sql.ErrNoRows {
		return errChapterNotFound("Chapter not found")
	}

	if err != nil {
		return err
	}

	// We also fetch the list of chapter versions from database
	return getChapterVersions(chapter, chapter_id)
}

// Fetch versions of chapter from database and add it to the chapter object passed as input
func getChapterVersions(chapter *Chapter, chapter_id int) error {
	rows, err := db.Query(GET_CHAPTER_VERSIONS_QUERY, chapter_id)
	if err != nil {
		return err
	}

	for rows.Next() {
		version := ChapterVersion{}
		var appVersion string
		err = rows.Scan(
			&version.Chapter_version_id,
			&version.Created_by,
			&version.Chapter_version_number,
			&version.Created,
			&appVersion)

		if err != nil {
			return err
		}

		// Convert appversion value into marketing friendly label
		if marketing_name, ok := APP_VERSION_MAP[appVersion]; ok {
			version.Appversion = marketing_name
		} else {
			version.Appversion = APP_VERSION_MAP["MISSING"]
		}

		chapter.Versions = append(chapter.Versions, version)
	}

	return nil
}
