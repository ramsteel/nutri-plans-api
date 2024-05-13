package recommendation

import (
	"nutri-plans-api/entities"
	"strings"

	"github.com/google/uuid"
)

func ToStruct(res string, uid uuid.UUID) *[]entities.Recommendation {
	var recommendations []entities.Recommendation
	splittedRes := strings.Split(res, "\n")

	for _, line := range splittedRes {
		trim := strings.TrimLeft(line, "-1234567. ")
		if trim != "" {
			recommendations = append(recommendations, entities.Recommendation{
				UserPreferenceID: uid,
				Name:             trim,
			})
		}
	}

	return &recommendations
}
