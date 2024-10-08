package main

import "fmt"

func mean(values []float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}

func linearRegression(clients, revenue []float64) (float64, float64) {
	xMean := mean(clients)
	yMean := mean(revenue)

	var covariance float64
	var variance float64
	for i := 0; i < len(clients); i++ {
		covariance += (clients[i] - xMean) * (revenue[i] - yMean)
		variance += (clients[i] - xMean) * (clients[i] - xMean)
	}
	slope := covariance / variance

	intercept := yMean - slope*xMean
	fmt.Println(intercept)

	return slope, intercept
}

func predict(clients float64, slope float64, intercept float64) float64 {
	return slope*clients + intercept
}

func main() {
	clients := []float64{10, 15, 20, 25, 30}
	revenue := []float64{100, 150, 200, 250, 300}

	slope, intercept := linearRegression(clients, revenue)
	clientsPrediction := 35.0
	revenuePrediction := predict(clientsPrediction, slope, intercept)
	fmt.Printf("A previsão de faturamento para %.0f clientes é: %.2f\n", clientsPrediction, revenuePrediction)
}
