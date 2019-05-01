package models

import "time"

type (
	DisplayStatus struct {
		ItemID string				`json:"item_id"`
		ItemLabel string			`json:"item_label"`
		ItemType DisplayStatusType	`json:"item_type"`
		TeamID Team					`json:"team_id"`
		X float64					`json:"x"`
		Y float64					`json:"y"`
		Z float64					`json:"z"`
		LastUpdated	time.Time		`json:"last_updated"`
	}

	DisplayStatusProvider interface {
		GetDisplayStatus() *DisplayStatus
	}

	DisplayStatusType string
)


const (
	TypeBall DisplayStatusType = "ball"
	TypePlayer DisplayStatusType = "player"

)
