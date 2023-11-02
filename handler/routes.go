package handler

import (
	"TEKSystems/database"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetRoutes(router *mux.Router, db gorm.DB) {
	router.HandleFunc("/alerts/", func(w http.ResponseWriter, r *http.Request) {
		service_id := r.URL.Query().Get("service_id")
		start_ts := r.URL.Query().Get("start_ts")
		end_ts := r.URL.Query().Get("end_ts")
		fmt.Println(service_id, start_ts, end_ts)
		Alerts(&db, w, service_id, start_ts, end_ts)
	}).Methods("GET")

	router.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		var alert database.Alert
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&alert); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		AddAlert(&db, w, alert)
	}).Methods("POST")

}
