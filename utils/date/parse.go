package date

import "time"

func ParseDate(date string) (*time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	return &parsedDate, nil
}
