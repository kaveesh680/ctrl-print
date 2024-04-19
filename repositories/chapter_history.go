package repositories

import (
	"database/sql"
	"encoding/json"
	"github.com/kaveesh680/ctrl-print/internal/domain"
	"github.com/kaveesh680/ctrl-print/internal/errors"
	"net/http"
)

type ChapterHistoryRepo struct {
	DbConnection *sql.DB
}

type ChapterHistoryRepoInterface interface {
	GetChapterHistoryByChapterId(chapterId int) (interface{}, error)
}

func NewChapterHistoryRepo(DbConnection *sql.DB) ChapterHistoryRepoInterface {
	return &ChapterHistoryRepo{
		DbConnection: DbConnection,
	}
}

func (i *ChapterHistoryRepo) GetChapterHistoryByChapterId(chapterId int) (interface{}, error) {

	if chapterId <= 0 {
		return nil, errors.NewGeneralError(
			http.StatusBadRequest,
			"Invalid chapter ID",
			"Invalid chapter ID",
		)
	}

	query := `SELECT
    chapter.chapter_name AS chapter_name,
    project.project_name AS project_name,
    company.company_name AS company_name,
    json_agg(json_build_object(
        'chapter_version_id', chapter_version.chapter_version_id,
        'created_date', chapter_version.chapter_version_create_date,
        'version_number', chapter_version.chapter_version_number,
        'username', person.person_username,
        'application_version', 
            CASE chapter_version.chapter_version_appversion
                WHEN '11.0' THEN 'CC 2015'
                WHEN '12.0' THEN 'CC 2017'
                ELSE chapter_version.chapter_version_appversion
            END
    ) ORDER BY chapter_version.chapter_version_number) AS version_history
FROM
    chapter
JOIN project ON chapter.chapter_project_id = project.project_id
JOIN company ON project.project_company_id = company.company_id
JOIN chapter_version ON chapter.chapter_id = chapter_version.chapter_version_chapter_id
JOIN person ON chapter_version.chapter_version_person_id = person.person_id
WHERE
    chapter.chapter_id = $1
GROUP BY
    chapter.chapter_name,
    project.project_name,
    company.company_name;
`
	rows, err := i.DbConnection.Query(query, chapterId)

	if err != nil {
		err := errors.NewGeneralError(
			http.StatusInternalServerError,
			"Oops Something Went Wrong",
			"Database Query Error",
		)
		return nil, err
	}
	defer rows.Close()

	var chapterHistory domain.ChapterHistory

	for rows.Next() {
		var versionHistory string
		err := rows.Scan(
			&chapterHistory.ChapterName,
			&chapterHistory.ProjectName,
			&chapterHistory.CompanyName,
			&versionHistory)
		if err != nil {
			err := errors.NewGeneralError(
				http.StatusInternalServerError,
				"Oops Something Went Wrong",
				"Database rows.Scan Error",
			)
			return nil, err
		}

		err = json.Unmarshal([]byte(versionHistory), &chapterHistory.VersionHistory)
		if err != nil {
			err := errors.NewGeneralError(
				http.StatusInternalServerError,
				"Oops Something Went Wrong",
				"Unmarshal error",
			)
			return nil, err
		}
	}

	return chapterHistory, nil

}
