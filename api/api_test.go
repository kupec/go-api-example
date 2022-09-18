package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"example.com/api/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	db, router := SetupRouter()
	defer db.Close()

	req, _ := http.NewRequest("GET", "/health_check", nil)
	resp := DoRequest(router, req)

	assert.Equal(t, 200, resp.Code)
}

func TestListUsers(t *testing.T) {
	db, router := SetupRouter()
	defer db.Close()

	user := NewUser()
	db.MustExec("INSERT INTO users(id, name, visits) VALUES(?, ?, ?)", user.Id, user.Name, user.Visits)

	req, _ := http.NewRequest("GET", "/users", nil)
	resp := DoRequest(router, req)

	assert.Equal(t, 200, resp.Code)

	var responseList []models.User
	json.Unmarshal(resp.Body.Bytes(), &responseList)
	assert.Equal(t, []models.User{user}, responseList)
}
