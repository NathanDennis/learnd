package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Building struct {
	Name     string `json:"name"`
	Customer string `json:"customer"`
	SerialID string `json:"serialID"`
}

type PowerMeter struct {
	Building         Building  `json:"building"`
	Consumes         float64   `json:"consumes"` // kwh per day
	InstallationDate time.Time `json:"installationDate"`
}

var powerMeters = map[string]PowerMeter{
	"1111-1111-1111": {
		Building: Building{
			Name:     "Treatment Plant A",
			Customer: "Aquaflow",
			SerialID: "1111-1111-1111",
		},
		Consumes:         20,
		InstallationDate: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)},
	"1111-1111-2222": {
		Building: Building{
			Name:     "Treatment Plant B",
			Customer: "Aquaflow",
			SerialID: "1111-1111-2222",
		},
		Consumes:         30,
		InstallationDate: time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)},
	"1111-1111-3333": {
		Building: Building{
			Name:     "Student Halls",
			Customer: "Albers Facilities Management",
			SerialID: "1111-1111-3333",
		}, Consumes: 40,
		InstallationDate: time.Date(2023, time.March, 1, 0, 0, 0, 0, time.UTC)},
}

func getMetersForCustomer(w http.ResponseWriter, r *http.Request) {
	customer := r.URL.Query().Get("customer")

	var customerMeters []Building
	for _, meter := range powerMeters {
		if meter.Building.Customer == customer {
			customerMeters = append(customerMeters, meter.Building)
		}
	}

	jsonResponse, err := json.Marshal(customerMeters)
	if err != nil {
		http.Error(w, "error encoding customer meters JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getKWHReadingForMeter(w http.ResponseWriter, r *http.Request) {
	serialID := r.URL.Query().Get("serialID")

	meter, exists := powerMeters[serialID]
	if !exists {
		http.Error(w, "requested meter not found", http.StatusNotFound)
		return
	}

	// assume kWh reading is cumulative and goes by consumes * days
	// probably too simple for actual Learnd production logic
	daysSinceInstallation := time.Since(meter.InstallationDate).Hours() / 24
	kwhReading := int(meter.Consumes * daysSinceInstallation)

	jsonResponse, err := json.Marshal(map[string]int{"kWh reading": kwhReading})
	if err != nil {
		http.Error(w, "error encoding meter reading JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/getMetersForCustomer", getMetersForCustomer)
	http.HandleFunc("/getMeterReading", getKWHReadingForMeter)

	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
