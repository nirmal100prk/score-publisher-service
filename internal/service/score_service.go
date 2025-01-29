package service

import (
	"context"
	"score-publisher-svc/internal/repository/postgres"
)

type ScoreService struct {
	scoreRepo postgres.PgxRepository
}

func NewScoreService(repo postgres.PgxRepository) *ScoreService {
	return &ScoreService{scoreRepo: repo}
}

func (s *ScoreService) InsertScore(ctx context.Context, val int64) error {
	return s.scoreRepo.InsertScore(ctx, val)
}
