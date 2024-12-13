// Package services provides application business logic.
package services

import (
	"context"
	"fmt"
	"log/slog"
)

type RecsService struct {
	log  *slog.Logger
	recs RecsProvider
}

type RecsProvider interface {
	Recs(ctx context.Context) (string, error)
}

// NewRecsService is a constructor for RecsService.
func NewRecsService(
	log *slog.Logger,
	recs RecsProvider,
) *RecsService {
	return &RecsService{
		log:  log,
		recs: recs,
	}
}

func (s *RecsService) Recommend(
	ctx context.Context,
) (string, error) {
	const op = "services.RecsService.Recommend"

	log := s.log.With(slog.String("op", op))

	log.Info("trying to retrieve recommendations")

	res, err := s.recs.Recs(ctx)
	if err != nil {
		log.Error("failed to retrieve recommendations", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("recommendations retrieved")

	return res, nil
}
