package models

import (
	"courier/services/logistic/database"
	"net/http"
)

type Logistic struct {
	LogisticName    string `json:"logistic_name"`
	Amount          string `json:"amount"`
	DestinationName string `json:"destination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}

func GetAllLogistics(logistics *[]Logistic) (int, error) {
	db := database.DB

	query := `SELECT logistic_name, amount, destination_name, origin_name, duration
			  FROM couriers`

	rows, err := db.Query(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		var logistic Logistic
		err := rows.Scan(&logistic.LogisticName, &logistic.Amount, &logistic.DestinationName, &logistic.OriginName, &logistic.Duration)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		*logistics = append(*logistics, logistic)
	}

	return http.StatusOK, nil
}

func FindLogistics(origin_name string, destination_name string, logistics *[]Logistic) (int, error) {
	db := database.DB

	query := `SELECT logistic_name, amount, destination_name, origin_name, duration
			  FROM couriers
			  WHERE couriers.origin_name = ? 
			  AND couriers.destination_name = ?`

	rows, err := db.Query(query, origin_name, destination_name)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		var logistic Logistic
		err := rows.Scan(&logistic.LogisticName, &logistic.Amount, &logistic.DestinationName, &logistic.OriginName, &logistic.Duration)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		*logistics = append(*logistics, logistic)
	}

	return http.StatusOK, nil
}
