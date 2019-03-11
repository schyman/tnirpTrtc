package main

import "time"

type Chapter struct {
	Company  string           `json:"company"`
	Project  string           `json:"project"`
	Chapter  string           `json:"chapter"`
	Versions []ChapterVersion `json:"versions"`
}

type ChapterVersion struct {
	CreatedBy            string    `json:"created_by"`
	ChapterVersionId     int       `json:"chapter_version_id"`
	ChapterVersionNumber int       `json:"version_number"`
	Created              time.Time ` json:"created"`
	Appversion           string    `json:"appversion"`
}

type Configuration struct {
	Database DatabaseConfig `json:"database"`
	Rest     RestConfig     `json:"rest"`
}

type DatabaseConfig struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type RestConfig struct {
	Port int `json:"port"`
}
