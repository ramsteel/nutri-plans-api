package recommendation

import (
	"fmt"
	recconst "nutri-plans-api/constants/recommendation"
	"nutri-plans-api/entities"
)

func ToString(recommendations *[]entities.Recommendation, isCron bool) []string {
	var res []string
	var temp string
	for _, rec := range *recommendations {
		temp += fmt.Sprintf("%s\n", rec.Name)
		if rec.ID%recconst.RecommendationLimit == 1 {
			res = append(res, temp[:len(temp)-1])
			if !isCron {
				break
			}
			temp = ""
		}
	}
	fmt.Println(res)
	return res
}
