package campaign

import (
	internalerrors "email/internal/internalErrors"
	"time"
	"github.com/rs/xid"
)

const (
	Pending = "Pending"
	Sent 	= "Sent"
	Failed	= "Failed"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"gte=1,dive"`
	Status	  string
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for index, emailValue := range emails {
		contacts[index].Email = emailValue
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}
	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	} else {
		return nil, err
	}
}
