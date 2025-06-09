package tickets_test

import (
	"challenge/internal/tickets"
	contracts "challenge/pkg/interfaces"
	"math"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	data := []string{
		"1,Fulano de Tal,fdt0@gmail.com,Finland,17:11,785",
		"2,Fulanete Zete,fnzt@hotmail.com,China,20:19,537",
		"3,Margarida Florinda,margui@@outlook.com,China,18:11,579",
		"4,Jovana di Giovanette,jojo3@fulana.com,Mongolia,23:16,1238",
	}

	t.Run("Count for Finland", func(t *testing.T) {
		total, err := tickets.GetTotalTickets("Finland", data)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if total != 1 {
			t.Errorf("expected 1, got %d", total)
		}
	})

	t.Run("Count for China", func(t *testing.T) {
		total, err := tickets.GetTotalTickets("China", data)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if total != 2 {
			t.Errorf("expected 2, got %d", total)
		}
	})
}

func TestGetCountByPeriod(t *testing.T) {
	list := []contracts.Ticket{
		{Schedule: 3},  // earlyMorning
		{Schedule: 8},  // morning
		{Schedule: 15}, // afternoon
		{Schedule: 21}, // night
		{Schedule: 25}, // invalid, should be ignored
	}

	t.Run("earlyMorning", func(t *testing.T) {
		count, err := tickets.GetCountByPeriod("earlyMorning", list)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if count != 1 {
			t.Errorf("expected 1, got %d", count)
		}
	})

	t.Run("night", func(t *testing.T) {
		count, err := tickets.GetCountByPeriod("night", list)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if count != 1 {
			t.Errorf("expected 1, got %d", count)
		}
	})

	t.Run("Invalid period", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		tickets.GetCountByPeriod("invalid", list)
	})
}

func TestAverageDestination(t *testing.T) {
	list := []contracts.Ticket{
		{Country: "Poland"},
		{Country: "Poland"},
		{Country: "France"},
	}

	t.Run("Country exists multiple times", func(t *testing.T) {
		result, err := tickets.AverageDestination("Poland", list)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected := (2.0 / 3.0) * 100
		if !floatEquals(result, expected, 0.01) {
			t.Errorf("Esperado %.2f, obtido %.2f", expected, result)
		}
	})
}

func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
