package server

import (
	"context"

	"connectrpc.com/connect"
	"github.com/s992/logger/internal/generated/db"
	loggerv1 "github.com/s992/logger/internal/generated/proto/logger/v1"
)

type TagService struct {
	queries *db.Queries
}

func NewTagService(queries *db.Queries) *TagService {
	return &TagService{queries}
}

func (s *TagService) CreateTag(
	ctx context.Context,
	req *connect.Request[loggerv1.CreateTagRequest],
) (*connect.Response[loggerv1.CreateTagResponse], error) {
	tag, err := s.queries.CreateTag(ctx, req.Msg.Name)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&loggerv1.CreateTagResponse{
		Tag: &loggerv1.Tag{
			TagId: tag.ID,
			Name:  tag.Name,
		},
	})

	return res, nil
}

func (s *TagService) ListTags(
	ctx context.Context,
	req *connect.Request[loggerv1.ListTagsRequest],
) (*connect.Response[loggerv1.ListTagsResponse], error) {
	tags, err := s.queries.ListTags(ctx)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&loggerv1.ListTagsResponse{
		Tags: tagsToList(tags),
	})

	return res, nil
}

func (s *TagService) SearchTags(
	ctx context.Context,
	req *connect.Request[loggerv1.SearchTagsRequest],
) (*connect.Response[loggerv1.SearchTagsResponse], error) {
	tags, err := s.queries.SearchTags(ctx, req.Msg.Query)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&loggerv1.SearchTagsResponse{
		Tags: tagsToList(tags),
	})

	return res, nil
}

func tagsToList(tagResults []db.Tag) (tags []*loggerv1.Tag) {
	for _, tag := range tagResults {
		tags = append(tags, &loggerv1.Tag{
			TagId: tag.ID,
			Name:  tag.Name,
		})
	}

	return
}
