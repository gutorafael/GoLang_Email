package endpoints

import (
	"bytes"
	"email/internal/contract"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type serviceMock struct {
	mock.Mock
}

func (m *serviceMock) Create(newCampaign contract.NewCampaignDTO) (string, error) {
	args := m.Called(newCampaign)
	return args.String(0), args.Error(1)
}


func (m *serviceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	//args := m.Called(campaign)
	return nil, nil
}

func Test_PostCampaigns_Save(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignDTO{
		Name:    "My Campaign",
		Content: "My Content",
		Emails:  []string{"teste@teste.com.br"},
	}
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaignDTO) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("123", nil)
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignsPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)
}
