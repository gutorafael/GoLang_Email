package endpoints

import (
	internalerrors "email/internal/internalErrors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_InternalErr(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}

	handlerErr := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerErr.ServeHTTP(res, req)
	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}

func Test_HandlerError_DomainErr(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("Domain Error")
	}

	handlerErr := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerErr.ServeHTTP(res, req)
	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "Domain Error")
}

func Test_HandlerError_Object(t *testing.T) {
	assert := assert.New(t)
	type testBody struct {
		Id int
	}
	objExpected := testBody{Id: 4}
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, 201, nil
	}

	handlerErr := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerErr.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)
	returnedObj := testBody{}
	json.Unmarshal(res.Body.Bytes(), &returnedObj)
	assert.Equal(objExpected, returnedObj)
}
