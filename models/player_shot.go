package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// PlayerShot is used by pop to map your player_shots database table to your go code.
type PlayerShot struct {
	ID         uuid.UUID `json:"id" db:"id"`
	GameID     string    `json:"game_id" db:"game_id"`
	PlayerName string    `json:"player_name" db:"player_name"`
	Collision  string    `json:"collision" db:"collision"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p PlayerShot) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// PlayerShots is not required by pop and may be deleted
type PlayerShots []PlayerShot

// String is not required by pop and may be deleted
func (p PlayerShots) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PlayerShot) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.GameID, Name: "GameID"},
		&validators.StringIsPresent{Field: p.PlayerName, Name: "PlayerName"},
		&validators.StringIsPresent{Field: p.Collision, Name: "Collision"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PlayerShot) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PlayerShot) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
