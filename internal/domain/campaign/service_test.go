package campaign

import (
	"email/internal/contract"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	service = Service{}
	newCampaign = contract.NewCampaignDTO{
		Name: "My Campaign",
		Content: "My Content",
		Emails: []string{"teste1@test.com"},
	}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	
	
	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func Test_Create_SaveCampaign(t *testing.T) {
	assert := assert.New(t)
	
	repositoryMock := new(MockRepository)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		} else if campaign.Content != newCampaign.Content { 
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		} 
		return true
	})).Return(nil)
	service.Repository = repositoryMock
	
	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
	repositoryMock.AssertExpectations(t)
}