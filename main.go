package main

import (
	"challenge/internal/tickets"
	"fmt"
	"os"
	"strings"
)

func main() {
	ticketsList, err := readTicketsFile()
	if err != nil {
		panic(err)
	}

	ticketsData := strings.Split(string(ticketsList), "\n")

	tickets.GetCountByPeriod("17:11", ticketsData)

	total, err := tickets.GetTotalTickets("Finland", ticketsData)

	if err != nil {
		panic(err) // TODO verificar outras formas de tratar o erro
	}

	fmt.Println(total)
}

func readTicketsFile() (string, error) {
	data, err := os.ReadFile("./tickets.csv")

	if err != nil {
		return "", err
	}

	return string(data), nil
}
