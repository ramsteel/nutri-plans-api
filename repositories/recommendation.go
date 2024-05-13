package repositories

import (
	"context"
	recconst "nutri-plans-api/constants/recommendation"
	"nutri-plans-api/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RecommendationRepository interface {
	GetRecommendations(ctx context.Context, uid uuid.UUID) (*[]entities.Recommendation, error)
	CreateRecommendations(ctx context.Context, recommendation *[]entities.Recommendation) error
}

type recommendationRepository struct {
	db *gorm.DB
}

func NewRecommendationRepository(db *gorm.DB) *recommendationRepository {
	return &recommendationRepository{
		db: db,
	}
}

func (r *recommendationRepository) GetRecommendations(
	ctx context.Context,
	uid uuid.UUID,
) (*[]entities.Recommendation, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	recommendations := new([]entities.Recommendation)
	res := r.db.Order("id desc").
		Limit(recconst.RecommendationLimit).
		Where("user_preference_id = ?", uid).
		Find(recommendations)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return recommendations, nil
}

func (r *recommendationRepository) CreateRecommendations(
	ctx context.Context,
	recommendation *[]entities.Recommendation,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	return r.db.Create(recommendation).Error
}
