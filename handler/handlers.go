package handler

import (
	"TEKSystems/database"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Alerts(db *gorm.DB, w http.ResponseWriter, service_id string, start_ts string, end_ts string) {
	fmt.Println("GET alerts for ", service_id, " between ", start_ts, " and ", end_ts)
	alertList, err := database.GetAlertList(db, service_id, start_ts, end_ts)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(alertList)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonData))
}

func AddAlert(db *gorm.DB, w http.ResponseWriter, alert database.Alert) {
	alertList, err := database.AddAlert(db, alert)
	// err = errors.New("test my error logic")
	if err != nil {
		response := database.AlertResponse{AlertID: alertList.Alert_ID, Error: err.Error()}
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		errResponse := database.AlertResponse{AlertID: alertList.Alert_ID, Error: response.Error}
		jsonString, _ := json.Marshal(errResponse)
		fmt.Fprintln(w, string(jsonString))
		return
	}
	response := database.AlertResponse{AlertID: alertList.Alert_ID, Error: ""}
	jsonString, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonString))
}
