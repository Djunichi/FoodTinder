package model

import "github.com/google/uuid"

type FetchResponse struct {
	APIVersion string `json:"apiVersion"`
	Status     string `json:"status"`
	Data       struct {
		ID              uuid.UUID        `json:"id"`
		Name            string           `json:"name"`
		Currency        string           `json:"currency"`
		MachineProducts []MachineProduct `json:"machineProducts"`
	} `json:"data"`
}
