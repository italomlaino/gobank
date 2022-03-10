package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Handle(c *gin.Context) {
	m.Called()
}

func (m *Mock) CreateAccountHandler() func(c *gin.Context) {
	args := m.Called()
	return args.Get(0).(func(c *gin.Context))
}

func (m *Mock) FetchAccountHandler() func(c *gin.Context) {
	args := m.Called()
	return args.Get(0).(func(c *gin.Context))
}

func (m *Mock) CreateTransationHandler() func(c *gin.Context) {
	args := m.Called()
	return args.Get(0).(func(c *gin.Context))
}

func (m *Mock) ListTransationHandler() func(c *gin.Context) {
	args := m.Called()
	return args.Get(0).(func(c *gin.Context))
}

func TestStart(t *testing.T) {
	mock := new(Mock)
	mock.On("Handle").Return().Times(4)
	mock.On("CreateAccountHandler").Return(mock.Handle)
	mock.On("FetchAccountHandler").Return(mock.Handle)
	mock.On("CreateTransationHandler").Return(mock.Handle)
	mock.On("ListTransationHandler").Return(mock.Handle)

	server := NewServer("8080", mock, mock)
	go func() {
		server.Start()
	}()

	time.Sleep(5 * time.Second)

	assert := assert.New(t)
	json, _ := json.Marshal(map[string]string{})
	resp, err := http.Post("http://localhost:8080/accounts", "application/json", bytes.NewBuffer(json))
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	resp, err = http.Get("http://localhost:8080/accounts/1")
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	resp, err = http.Post("http://localhost:8080/transactions", "application/json", bytes.NewBuffer(json))
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	resp, err = http.Get("http://localhost:8080/transactions")
	assert.Nil(err)
	assert.Equal(200, resp.StatusCode)

	mock.AssertExpectations(t)
}
