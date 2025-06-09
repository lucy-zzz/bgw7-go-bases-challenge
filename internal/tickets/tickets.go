package tickets

import (
	Contracts "challenge/pkg/interfaces"
	"errors"
	"strings"
)

func GetTotalTickets(destination string, data []string) (int, error) {
	var total []string

	for _, d := range data {
		if strings.Contains(d, destination) {
			total = append(total, d)
		}
	}

	return len(total), nil
}

func GetCountByPeriod(p string, list []Contracts.Ticket) (int, error) {
	earlyMorning := "earlyMorning"
	morning := "morning"
	afternoon := "afternoon"
	night := "night"

	var flights int
	var err error

	switch p {
	case earlyMorning:
		for _, i := range list {
			if i.Schedule >= 0 && i.Schedule <= 6 {
				flights++
			}
		}
	case morning:
		for _, i := range list {
			if i.Schedule >= 7 && i.Schedule <= 12 {
				flights++
			}
		}
	case afternoon:
		for _, i := range list {
			if i.Schedule >= 13 && i.Schedule <= 19 {
				flights++
			}
		}
	case night:
		for _, i := range list {
			if i.Schedule >= 20 && i.Schedule <= 23 {
				flights++
			}
		}
	default:
		err = errors.New("Schedule not available. Choose between early morning, morning, afternoon or night.")
		panic(err.Error())
	}

	return flights, nil
}

func AverageDestination(destination string, list []Contracts.Ticket) (float64, error) {
	if len(list) == 0 {
		return 0, errors.New("no such country on the list")
	}

	var countCountry int
	for _, t := range list {
		if t.Country == destination {
			countCountry++
		}
	}

	result := (float64(countCountry) / float64(len(list))) * 100

	return result, nil
}
