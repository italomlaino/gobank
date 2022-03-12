package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/italomlaino/gobank/application/router"
	mocks "github.com/italomlaino/gobank/mocks/application/controller"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Handle(*gin.Context) {
	m.Called()
}

func (m *Mock) CreateHandler() func(c *gin.Context) {
	args := m.Called()
	return args.Get(0).(func(c *gin.Context))
}

func TestStart(t *testing.T) {
	handler := new(Mock)
	accountController := new(mocks.AccountController)
	transactionController := new(mocks.TransactionController)

	handler.On("Handle").Return().Times(4)
	accountController.On("CreateHandler").Return(handler.Handle)
	accountController.On("FetchByIDHandler").Return(handler.Handle)
	transactionController.On("CreateHandler").Return(handler.Handle)
	transactionController.On("FetchByAccountIDHandler").Return(handler.Handle)

	port := "9050"
	host := "http://localhost:" + port

	subject := router.NewRouter(port, accountController, transactionController)
	go func() {
		subject.Start()
	}()

	time.Sleep(5 * time.Second)

	assert := assert.New(t)
	json, _ := json.Marshal(map[string]string{})
	body := bytes.NewBuffer(json)
	resp, err := http.Post(host+"/accounts", "application/json", body)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	resp, err = http.Get(host + "/accounts/1")
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	resp, err = http.Get(host + "/accounts/1/transactions")
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	resp, err = http.Post(host+"/transactions", "application/json", body)
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	accountController.AssertExpectations(t)
}
