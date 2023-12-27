package campaign

import (
	"email/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	service     = Service{}
	newCampaign = contract.NewCampaignDTO{
		Name:    "My Campaign",
		Content: "My Content",
		Emails:  []string{"teste1@test.com"},
	}
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}
func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	assert.Nil(err)
	assert.Equal(name, campaign.Name)

}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(MockRepository)
	service.Repository = repositoryMock
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepo(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(MockRepository)
	service.Repository = repositoryMock
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)
	assert.Equal(err.Error(), "error to save on database")
}

func Test_Create_ValidateDomain(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal(err.Error(), "name is required")
}
