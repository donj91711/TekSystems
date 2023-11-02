package database

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Alert struct {
	ID           uint `gorm:"primaryKey"`
	Alert_ID     string
	Service_id   string
	Service_name string
	Model        string
	Alert_type   string
	Alert_ts     string
	Alert_ts_int int64
	Severity     string
	Team_slack   string
}

type AlertResponse struct {
	AlertID string `json:"alert_id"`
	Error   string `json:"error"`
}

type AlertReturnSingle struct {
	Alert_ID   string
	Model      string
	Alert_type string
	Alert_ts   string
	Severity   string
	Team_slack string
}

type AlertReturn struct {
	Service_id   string
	Service_name string
	Alerts       []AlertReturnSingle `gorm:"foreignkey:ServiceID"`
}

func OpenDB(username, password, host, port, dbname string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	// try to open the database several times because sometimes Docker doesn't have MySQL ready to go yet
	maxRetries := 10
	retryCount := 0
	retryInterval := 5 * time.Second
	for {
		fmt.Println("TRY: ", retryCount)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("DATABASE OPEN")
			break
		}

		retryCount++

		if retryCount >= maxRetries {
			fmt.Printf("Max retries reached. Unable to connect to the database: %v\n", err)
			return nil, err
		}

		fmt.Printf("Error connecting to the database (retry %d/%d): %v\n", retryCount, maxRetries, err)

		time.Sleep(retryInterval)
	}
	return db, nil
}

func (AlertReturn) TableName() string {
	return "alerts"
}

func GetAlertList(db *gorm.DB, service_id string, start_ts string, end_ts string) (AlertReturn, error) {
	var alerts []Alert
	condition := db.Where("service_id = ?", service_id)
	if start_ts != "" {
		condition = condition.Where("alert_ts >= ?", start_ts)
	}
	if end_ts != "" {
		condition = condition.Where("alert_ts <= ?", end_ts)
	}
	result := condition.Find(&alerts)
	fmt.Println(result)
	alertReturn := AlertReturn{
		Service_id:   alerts[0].Service_id,
		Service_name: alerts[0].Service_name,
		Alerts:       make([]AlertReturnSingle, len(alerts)),
	}
	for i, alert := range alerts {
		alertReturn.Alerts[i] = AlertReturnSingle{
			Alert_ID:   alert.Alert_ID,
			Model:      alert.Model,
			Alert_type: alert.Alert_type,
			Alert_ts:   alert.Alert_ts,
			Severity:   alert.Severity,
			Team_slack: alert.Team_slack,
		}
	}
	fmt.Println(alertReturn)
	return alertReturn, nil
}

func AddAlert(db *gorm.DB, alert Alert) (Alert, error) {
	unix_time, err := strconv.ParseInt(alert.Alert_ts, 10, 64)
	if err != nil {
		unix_time = 0
	}
	alert.Alert_ts_int = unix_time
	result := db.Create(&alert)
	if result.Error != nil {
		return Alert{}, result.Error
	}
	return alert, nil
}
