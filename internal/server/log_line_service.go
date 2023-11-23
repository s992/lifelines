package server

import (
	"context"
	"database/sql"

	"connectrpc.com/connect"
	"github.com/s992/lifelines/internal/generated/db"
	lifelinesv1 "github.com/s992/lifelines/internal/generated/proto/lifelines/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LogLineService struct {
	queries *db.Queries
}

func NewLogLineService(queries *db.Queries) *LogLineService {
	return &LogLineService{queries}
}

func (s *LogLineService) CreateLogLine(
	ctx context.Context,
	req *connect.Request[lifelinesv1.CreateLogLineRequest],
) (*connect.Response[lifelinesv1.CreateLogLineResponse], error) {
	tag, err := s.queries.GetTag(ctx, req.Msg.TagId)
	if err != nil {
		return nil, err
	}

	logLine, err := s.queries.CreateLogLine(ctx, db.CreateLogLineParams{
		TagID:       req.Msg.TagId,
		Value:       req.Msg.Value,
		Description: req.Msg.Description,
	})
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&lifelinesv1.CreateLogLineResponse{
		LogLine: &lifelinesv1.LogLine{
			LogLineId:   logLine.ID,
			Value:       logLine.Value,
			Description: logLine.Description,
			Tag: &lifelinesv1.Tag{
				TagId: tag.ID,
				Name:  tag.Name,
			},
		},
	})

	return res, nil
}

func (s *LogLineService) ListLogLines(
	ctx context.Context,
	req *connect.Request[lifelinesv1.ListLogLinesRequest],
) (*connect.Response[lifelinesv1.ListLogLinesResponse], error) {
	logLines, err := s.queries.ListLogLines(ctx, db.ListLogLinesParams{
		TagID:         sql.NullInt64{Int64: req.Msg.GetTagId(), Valid: req.Msg.TagId != nil},
		EndDateTime:   sql.NullTime{Time: req.Msg.End.AsTime(), Valid: req.Msg.End != nil},
		StartDateTime: sql.NullTime{Time: req.Msg.Start.AsTime(), Valid: req.Msg.Start != nil},
	})
	if err != nil {
		return nil, err
	}

	lines := []*lifelinesv1.LogLine{}
	for _, line := range logLines {
		lines = append(lines, &lifelinesv1.LogLine{
			LogLineId:   line.ID,
			Value:       line.Value,
			Description: line.Description,
			CreatedAt:   timestamppb.New(line.CreatedAt),
			Tag: &lifelinesv1.Tag{
				TagId: line.TagID,
				Name:  line.TagName,
			},
		})
	}

	res := connect.NewResponse(&lifelinesv1.ListLogLinesResponse{LogLines: lines})

	return res, nil
}
