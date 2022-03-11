package controller_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	"github.com/italomlaino/gobank/controller"
	"github.com/italomlaino/gobank/domain"
 	"github.com/italomlaino/gobank/mocks/domain"
)

func TestCreateAccountHandler(t *testing.T) {
	account := &domain.Account{ID: 1, DocumentNumber: 12345678}

	mock := new(mocks.AccountService)
	mock.On("Create", int64(12345678)).Return(account, nil)

	subject := controller.NewAccountController(mock)
	handler := subject.CreateAccountHandler()

	expectedBody, _ := json.Marshal(account)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "", bytes.NewBuffer(expectedBody))
	handler(c)

	assert := assert.New(t)
	assert.Equal(200, w.Code)
	actualBody, _ := ioutil.ReadAll(w.Body)

	assert.Equal(string(expectedBody), string(actualBody))
}

func TestFetchAccountHandler(t *testing.T) {
	account := &domain.Account{
		ID:             1,
		DocumentNumber: 12345678,
	}

	mock := new(mocks.AccountService)
	mock.On("FetchByID", account.ID).Return(account, nil)

	subject := controller.NewAccountController(mock)
	handler := subject.FetchAccountHandler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key:   "accountId",
			Value: "1",
		},
	}
	handler(c)

	assert := assert.New(t)
	assert.Equal(200, w.Code)
	actualBody, _ := ioutil.ReadAll(w.Body)

	expectedBody, _ := json.Marshal(account)
	assert.Equal(string(expectedBody), string(actualBody))
}
