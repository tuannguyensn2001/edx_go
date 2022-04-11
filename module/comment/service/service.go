package commentservice

import (
	"context"
	"edx_go/models"
	commentrepository "edx_go/module/comment/repository"
	commentstruct "edx_go/module/comment/struct"
)

type dataComment struct {
	Content  string `json:"content"`
	LessonId int    `json:"lessonId"`
}

type service struct {
	repository commentrepository.CommentRepository
}

func NewCommentService(repo commentrepository.CommentRepository) *service {
	return &service{repository: repo}
}

func (s *service) Create(ctx context.Context, input commentstruct.CommentInput) (*models.Comment, error) {

	comment := &models.Comment{
		Content:          input.Content,
		UserId:           input.UserId,
		CommentableModel: input.CommentableModel,
		CommentableId:    input.CommentableId,
	}

	err := s.repository.Create(ctx, comment)

	if err != nil {
		return nil, err
	}

	return comment, nil

}
