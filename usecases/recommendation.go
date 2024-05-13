package usecases

import (
	"context"
	"log"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/dto"
	exnutri "nutri-plans-api/externals/nutrition"
	exoai "nutri-plans-api/externals/openai"
	"nutri-plans-api/repositories"
	"nutri-plans-api/utils/prompt"
	recutil "nutri-plans-api/utils/recommendation"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron"
)

type RecommendationUsecase interface {
	GetRecommendation(c echo.Context, uid uuid.UUID) (*[]exnutri.ItemNutrition, error)
}

type recommendationUsecase struct {
	recommendationRepo repositories.RecommendationRepository
	userPrefRepo       repositories.UserPreferenceRepository

	nutritionExternal exnutri.NutritionClient
	oaiExternal       exoai.OpenAIClient
}

func NewRecommendationUsecase(
	recommendationRepo repositories.RecommendationRepository,
	userPrefRepo repositories.UserPreferenceRepository,
	nutritionExternal exnutri.NutritionClient,
	recExternal exoai.OpenAIClient,
) *recommendationUsecase {
	return &recommendationUsecase{
		recommendationRepo: recommendationRepo,
		userPrefRepo:       userPrefRepo,
		nutritionExternal:  nutritionExternal,
		oaiExternal:        recExternal,
	}
}

func (r *recommendationUsecase) StartRecommendationCron() {
	c := cron.New()
	if err := c.AddFunc("@weekly", r.FetchOpenAIRecommendation); err != nil {
		log.Fatal(msgconst.MsgFailedAddRecommendationCron)
	}

	go func() {
		c.Start()
		defer c.Stop()

		select {}
	}()
}

func (r *recommendationUsecase) GetRecommendation(
	c echo.Context,
	uid uuid.UUID,
) (*[]exnutri.ItemNutrition, error) {
	ctx, cancel := context.WithCancel(c.Request().Context())
	defer cancel()

	recommendations, err := r.recommendationRepo.GetRecommendations(ctx, uid)
	if err != nil {
		return nil, err
	}

	query := recutil.ToString(recommendations, false)
	req := &dto.ItemNutritionRequest{
		Query: query[0],
	}
	itemNutritions, err := r.nutritionExternal.GetMultipleItemNutritions(ctx, req)
	if err != nil {
		return nil, err
	}

	return itemNutritions, nil
}

func (r *recommendationUsecase) FetchOpenAIRecommendation() {
	ctx := context.Background()
	preferences, _ := r.userPrefRepo.GetAllUserPreferences(ctx)
	for _, pref := range *preferences {
		prompt := prompt.GetRecommendationPrompt(&pref)
		res, err := r.oaiExternal.GetRecommendation(
			prompt, recutil.ToString(pref.Recommendations, true))
		if err == nil {
			recommendations := recutil.ToStruct(res, pref.UserID)
			if err := r.recommendationRepo.CreateRecommendations(ctx, recommendations); err != nil {
				log.Println(msgconst.MsgFailedCreateRecommendation)
			}
		}
	}
}
