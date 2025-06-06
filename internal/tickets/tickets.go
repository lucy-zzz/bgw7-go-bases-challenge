package tickets

import (
	"errors"
	"strings"
	"time"
)

type Ticket struct {
	id       int
	name     string
	email    string
	country  string
	schedule time.Time
	price    float64
}

// TODO separate tickets by country
func GetTotalTickets(destination string, data []string) (int, error) {
	var total []string

	for _, d := range data {
		if strings.Contains(d, destination) {
			total = append(total, d)
		}
	}

	return len(total), nil
}

func GetCountByPeriod(p string, data []string) (int, error) {
	earlyMorning := "earlyMorning"
	morning := "morning"
	afternoon := "afternoon"
	night := "night"

	getSchedules(data)

	var flights int
	var err error

	switch p {
	case earlyMorning:
		getEarlyMorningFlights()
	case morning:
		getMorningFlights()
	case afternoon:
		getAfternoonFlights()
	case night:
		getNightFlights()
	default:
		err = errors.New("Schedule not available. Choose between early morning, morning, afternoon or night.")
		panic(err.Error())
	}

	return flights, nil
}

func getSchedules(data []string) []string {
	var s []string
	// for _, d := range data {

	// }
	return s
}

func getEarlyMorningFlights() {

}

func getMorningFlights() {

}

func getAfternoonFlights() {

}

func getNightFlights() {

}

// // ejemplo 2
// func GetMornings(time string) (int error) {}

// // ejemplo 3
// func AverageDestination(destination string, total int) (int error) {}
