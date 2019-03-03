package main

const (
	GET_CHAPTER_QUERY = `
	SELECT 
		c.company_name AS Company, 
		p.project_name AS Project, 
		ch.chapter_name AS Chapter
	FROM chapter ch 
	INNER JOIN project p ON ch.chapter_project_id = p.project_id
	INNER JOIN company c ON p.project_company_id = c.company_id
	WHERE chapter_id = $1
	LIMIT 1
	`

	GET_CHAPTER_VERSIONS_QUERY = `
	SELECT 
		cv.chapter_version_id AS Chapter_version_id, 
		p.person_username AS Created_by, 
		cv.chapter_version_number AS Chapter_version_number,
		cv.chapter_version_create_date AS Created,
		cv.chapter_version_appversion AS Appversion
	FROM chapter_version cv 
	INNER JOIN person p ON cv.chapter_version_person_id = p.person_id
	WHERE chapter_version_chapter_id = $1
	ORDER BY  chapter_version_number ASC
	`
)
