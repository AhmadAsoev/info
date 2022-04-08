package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Info struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAT *string   `json:"created_at"`
	UpdatedAT *string   `json:"updated_at,omitempty"`
	DeleteAT  *string   `json:"delete_at,omitempty"`
}

func (i Info) Validate() error {

	if i.ID == uuid.Nil {
		return errors.New("id must be not empty")
	}

	if len(i.Name) < 5 || len(i.Name) > 10 {
		return errors.New("is not correct")
	}
	if i.CreatedAT != nil {
		if _, err := time.Parse("02.01.2006T15:04", *i.CreatedAT); err != nil {
			return errors.New("time is not true")
		}
	}

	if i.UpdatedAT != nil {
		if _, err := time.Parse("02.01.2006T.15:04", *i.CreatedAT); err != nil {
			return errors.New("time is not true")
		}
	}

	if i.DeleteAT != nil {
		if _, err := time.Parse("02.01.06T.15:04", *i.CreatedAT); err != nil {
			return errors.New("time is not true")
		}
	}
	return nil
}
