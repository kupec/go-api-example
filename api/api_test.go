package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"example.com/api/models"
	_ "github.com/mattn/go-sqlite3"
)

func TestHealthCheck(t *testing.T) {
	db, router := SetupRouter()
	defer db.Close()

	req, _ := http.NewRequest("GET", "/health_check", nil)
	resp := DoRequest(router, req)

	if resp.Code != 200 {
		t.Errorf("Expected an 200 response. Got %d", resp.Code)
	}
}

func TestListUsers(t *testing.T) {
	db, router := SetupRouter()
	defer db.Close()

	user := models.User{
		Id:     1,
		Name:   "name",
		Visits: 2,
	}
	db.Exec("INSERT INTO users(id, name, visits) VALUES(?, ?, ?)", user.Id, user.Name, user.Visits)

	req, _ := http.NewRequest("GET", "/users", nil)
	resp := DoRequest(router, req)

	if resp.Code != 200 {
		t.Errorf("Expected an 200 response. Got %d", resp.Code)
	}

	var responseList []models.User
	json.Unmarshal(resp.Body.Bytes(), &responseList)
	if len(responseList) != 1 {
		t.Errorf("Expected response with len=1. Got %d", len(responseList))
		t.Errorf("responseList=%v", responseList)
	}

	response := responseList[0]
	if response.Id != user.Id {
		t.Errorf("Expected id=%d. Got %d", user.Id, response.Id)
	}
	if response.Name != user.Name {
		t.Errorf("Expected id=%s. Got %s", user.Name, response.Name)
	}
	if response.Visits != user.Visits {
		t.Errorf("Expected id=%d. Got %d", user.Visits, response.Visits)
	}
}
