package country

import (
	"bufio"
	"log"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/entities"
	"os"
)

func LoadData(fp string) *[]entities.Country {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(msgconst.MsgFailedOpenFile)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	countries := make([]entities.Country, 0, 250)
	i := uint(1)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		countries = append(
			countries,
			entities.Country{
				ID:   i,
				Name: string(line),
			},
		)
		i++
	}

	return &countries
}
