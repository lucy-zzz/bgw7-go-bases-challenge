package main

import (
	"challenge/internal/tickets"
	contracts "challenge/pkg/interfaces"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var ticketsSlice []contracts.Ticket

func main() {
	// Para fins de aprendizado, realizei testei a leitura do arquivo de formas diferentes

	// exercício 1 - forma 1 - manipulação de array
	ticketsList, err := readTicketsFile()
	if err != nil {
		panic(err)
	}

	ticketsData := strings.Split(string(ticketsList), "\n")

	total, err := tickets.GetTotalTickets("Finland", ticketsData)

	if err != nil {
		panic(err)
	}

	fmt.Println(total)

	// exercício 2
	// utilizando pkg csv
	list := useCSV()
	countPeriod, err := tickets.GetCountByPeriod("morning", list)

	if err != nil {
		panic(err)
	}

	fmt.Println(countPeriod)

	// exercício 3
	percentage, err := tickets.AverageDestination("Poland", list)
	if err != nil {
		panic(err)
	}
	fmt.Println(percentage)

}

func readTicketsFile() (string, error) {
	data, err := os.ReadFile("./tickets.csv")

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func useCSV() []contracts.Ticket {
	data, err := os.Open("./tickets.csv")

	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(data)

	tickets, readerErr := reader.ReadAll()

	if readerErr != nil {
		panic(readerErr)
	}

	for _, line := range tickets {
		if len(line) != 6 {
			continue
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			continue
		}
		price, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			continue
		}

		hour := line[4]
		schedule, err := parseSchedule(hour)

		if err != nil {
			continue
		}

		t := contracts.Ticket{
			Id:       id,
			Name:     line[1],
			Email:    line[2],
			Country:  line[3],
			Schedule: schedule,
			Price:    price,
		}

		ticketsSlice = append(ticketsSlice, t)
	}
	return ticketsSlice
}

func parseSchedule(h string) (int, error) {
	layout := "15:04"
	t, err := time.Parse(layout, h)
	if err != nil {
		return -1, err
	}
	return t.Hour(), nil
}
