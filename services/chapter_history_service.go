package services

import (
	"github.com/kaveesh680/ctrl-print/internal/domain"
	"github.com/kaveesh680/ctrl-print/internal/errors"
	"github.com/kaveesh680/ctrl-print/repositories"
	"net/http"
)

type ChapterHistoryService struct {
	ChapterHistoryByChapterIdRepository repositories.ChapterHistoryRepoInterface
}

func NewChapterHistoryService(repo repositories.ChapterHistoryRepoInterface) ChapterHistoryServiceInterface {
	return &ChapterHistoryService{
		ChapterHistoryByChapterIdRepository: repo,
	}
}

type ChapterHistoryServiceInterface interface {
	ChapterHistoryService(chapterId int) (interface{}, error)
}

func (i *ChapterHistoryService) ChapterHistoryService(chapterId int) (interface{}, error) {

	data, err := i.ChapterHistoryByChapterIdRepository.GetChapterHistoryByChapterId(chapterId)
	if err != nil {
		return nil, err
	}

	data, ok := data.(domain.ChapterHistory)
	if !ok {
		err := errors.NewGeneralError(
			http.StatusInternalServerError,
			"Oops Something Went Wrong",
			"ChapterHistory casting error",
		)
		return nil, err
	}

	return data, nil

}
