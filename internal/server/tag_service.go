package server

import (
	"context"

	"connectrpc.com/connect"
	"github.com/s992/lifelines/internal/generated/db"
	lifelinesv1 "github.com/s992/lifelines/internal/generated/proto/lifelines/v1"
)

type TagService struct {
	queries *db.Queries
}

func NewTagService(queries *db.Queries) *TagService {
	return &TagService{queries}
}

func (s *TagService) CreateTag(
	ctx context.Context,
	req *connect.Request[lifelinesv1.CreateTagRequest],
) (*connect.Response[lifelinesv1.CreateTagResponse], error) {
	tag, err := s.queries.CreateTag(ctx, req.Msg.Name)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&lifelinesv1.CreateTagResponse{
		Tag: &lifelinesv1.Tag{
			TagId: tag.ID,
			Name:  tag.Name,
		},
	})

	return res, nil
}

func (s *TagService) ListTags(
	ctx context.Context,
	req *connect.Request[lifelinesv1.ListTagsRequest],
) (*connect.Response[lifelinesv1.ListTagsResponse], error) {
	tags, err := s.queries.ListTags(ctx)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&lifelinesv1.ListTagsResponse{
		Tags: tagsToList(tags),
	})

	return res, nil
}

func (s *TagService) SearchTags(
	ctx context.Context,
	req *connect.Request[lifelinesv1.SearchTagsRequest],
) (*connect.Response[lifelinesv1.SearchTagsResponse], error) {
	tags, err := s.queries.SearchTags(ctx, req.Msg.Query)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&lifelinesv1.SearchTagsResponse{
		Tags: tagsToList(tags),
	})

	return res, nil
}

func tagsToList(tagResults []db.Tag) (tags []*lifelinesv1.Tag) {
	for _, tag := range tagResults {
		tags = append(tags, &lifelinesv1.Tag{
			TagId: tag.ID,
			Name:  tag.Name,
		})
	}

	return
}
