package campaign

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/jaswdr/faker"
)

var (
	fake = faker.New()
	name = "My Campaign"
	content = "My Content"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_ID(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOn(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_ValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal(err.Error(), "Name cannot be less than 5 characters")
}

func Test_NewCampaign_ValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal(err.Error(), "Name cannot be greater than 24 characters")
}

func Test_NewCampaign_ValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal(err.Error(), "Content cannot be less than 5 characters")
}

func Test_NewCampaign_ValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)

	assert.Equal(err.Error(), "Content cannot be greater than 1024 characters")
}

func Test_NewCampaign_ValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"invalid"})

	assert.Equal(err.Error(), "Email must be a valid email")
}

func Test_NewCampaign_ValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal(err.Error(), "Contacts must have at least one contact")
}

func Test_NewCampaign_StatusPending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(Pending, campaign.Status)
}