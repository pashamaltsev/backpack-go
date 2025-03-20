package models

type SystemStatus string

const (
	StatusOk          SystemStatus = "Ok"
	StatusMaintenance SystemStatus = "Maintenance"
)

type Status struct {
	Status  SystemStatus `json:"status"`
	Message string       `json:"message"`
}
