package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

type errChapterNotFound string

func (e errChapterNotFound) Error() string {
	return string(e)
}

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

func getChapter(chapter *Chapter, chapter_id int) error {
	row := db.QueryRow(GET_CHAPTER_QUERY, chapter_id)
	err := row.Scan(&chapter.Company, &chapter.Project, &chapter.Chapter)
	if err == sql.ErrNoRows {
		return errChapterNotFound("Chapter not found")
	}

	if err != nil {
		return err
	}

	return getChapterVersions(chapter, chapter_id)
}

func getChapterVersions(chapter *Chapter, chapter_id int) error {
	rows, err := db.Query(GET_CHAPTER_VERSIONS_QUERY, chapter_id)
	if err != nil {
		return err
	}

	for rows.Next() {
		version := ChapterVersion{}
		err = rows.Scan(
			&version.Chapter_version_id,
			&version.Created_by,
			&version.Chapter_version_number,
			&version.Created,
			&version.Appversion)

		if err != nil {
			return err
		}
		chapter.Versions = append(chapter.Versions, version)
	}

	return nil
}
