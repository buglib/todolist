package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setUpAll()
	m.Run()
	tearDownAll()
}

func setUpAll() {
	// 连接数据库，建库建表
	dsn := "buglib:123456@tcp(localhost:3306)/todolist_test?charset=utf8mb4&parseTime=True&loc=Local"
	initMysql(dsn)
}

func tearDownAll() {
	// 删库删表
	db.Exec("drop table tasks")
}

func TestPostTodolistReturn200(t *testing.T) {
	router := initRouter()
	w := httptest.NewRecorder()

	task := Task{
		ID:       1,
		TaskInfo: "实现'POST /todolist'的单元测试",
		State:    Todo,
	}
	reqBody, _ := json.Marshal(task)
	req, _ := http.NewRequest(
		"POST",
		"/v1/todolist",
		bytes.NewBuffer(reqBody),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestPostTodolistReturn500(t *testing.T) {
	r := initRouter()
	w := httptest.NewRecorder()

	task := Task{
		ID:       1,
		TaskInfo: "实现'POST /todolist'的单元测试",
		State:    Todo,
	}
	reqBody, _ := json.Marshal(task)
	req, _ := http.NewRequest(
		"POST",
		"/v1/todolist",
		bytes.NewBuffer(reqBody),
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}
