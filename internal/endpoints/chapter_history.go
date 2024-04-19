package endpoints

import (
	"encoding/json"
	"github.com/kaveesh680/ctrl-print/internal/db_connection"
	"github.com/kaveesh680/ctrl-print/internal/domain"
	"github.com/kaveesh680/ctrl-print/internal/errors"
	"github.com/kaveesh680/ctrl-print/repositories"
	"github.com/kaveesh680/ctrl-print/services"
	"net/http"
	"strconv"
)

func ChapterHistoryEndpoint(w http.ResponseWriter, r *http.Request) error {

	chapterId, err := DecodeChapterHistoryRequest(r)
	if err != nil {
		return err
	}

	chapterIdInt, ok := chapterId.(int)

	if !ok {
		err := errors.NewGeneralError(
			http.StatusInternalServerError,
			"Oops Something Went Wrong",
			"Int conversion error",
		)
		return err
	}

	if chapterIdInt <= 0 {
		err := errors.NewGeneralError(
			http.StatusBadRequest,
			"Invalid chapter ID",
			"Invalid chapter ID",
		)
		return err
	}

	data, err := services.NewChapterHistoryService(
		repositories.NewChapterHistoryRepo(
			db_connection.DBConnection,
		),
	).ChapterHistoryService(chapterIdInt)

	ChapterHistoryData, ok := data.(domain.ChapterHistory)

	if !ok {
		err := errors.NewGeneralError(
			http.StatusInternalServerError,
			"Oops Something Went Wrong",
			"ChapterHistory casting error",
		)
		return err
	}

	err = EncodeChapterHistoryResponse(w, ChapterHistoryData)
	if err != nil {
		return err
	}
	return nil
}

func DecodeChapterHistoryRequest(r *http.Request) (interface{}, error) {

	id := r.URL.Path[len("/chapter_versions/"):]
	defer r.Body.Close()

	chapterId, err := strconv.Atoi(id)

	if err != nil {
		err := errors.NewGeneralError(
			http.StatusBadRequest,
			"Oops Something Went Wrong",
			"Int conversion error",
		)
		return nil, err
	}

	return chapterId, nil
}

func EncodeChapterHistoryResponse(w http.ResponseWriter, response interface{}) error {

	r, ok := response.(domain.ChapterHistory)

	if !ok {
		err := errors.NewGeneralError(
			http.StatusInternalServerError,
			"Oops Something Went Wrong",
			"ChapterHistory casting error",
		)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	return json.NewEncoder(w).Encode(r)
}
