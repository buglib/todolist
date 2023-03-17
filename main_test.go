package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"todolist/domain"
	"todolist/infra/db"

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
	db.InitMysql(dsn)
	db.AutoMigrate(&domain.Task{})
}

func tearDownAll() {
	// 删库删表
	db.Db.Exec("drop table tasks")
}

func TestPostTodolistReturn200(t *testing.T) {
	router := initRouter()
	w := httptest.NewRecorder()

	task := domain.Task{
		ID:       1,
		TaskInfo: "实现'POST /todolist'的单元测试",
		State:    domain.Todo,
	}
	reqBody, _ := json.Marshal(task)
	req, _ := http.NewRequest(
		"POST",
		"/v1/todolist",
		bytes.NewBuffer(reqBody),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	// respBody := make([]byte, w.Body.Cap())
	// w.Body.Read(respBody)
	// fmt.Println(string(respBody))
}

func TestPostTodolistReturn500(t *testing.T) {
	db.Db.Exec("insert into tasks (id, task_info, state) values (1, 'test', 0)")

	r := initRouter()
	w := httptest.NewRecorder()

	task := domain.Task{
		ID:       1,
		TaskInfo: "test",
		State:    domain.Todo,
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

func TestGetTodolistReturn200(t *testing.T) {
	insertTask()

	r := initRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"GET",
		"/v1/todolist",
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestPutTodoItemReturn200(t *testing.T) {
	id := insertTask()
	task := domain.Task{ // 将指定的任务标记为已完成
		ID:    uint(id),
		State: 0,
	}

	r := initRouter()
	w := httptest.NewRecorder()

	reqBody, _ := json.Marshal(&task)
	req, _ := http.NewRequest(
		"PUT",
		fmt.Sprintf("/v1/todolist/%d", id),
		bytes.NewBuffer(reqBody),
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestPutTodoItemReturn404(t *testing.T) {
	id := rand.Intn(1000)
	task := domain.Task{ // 将指定的任务标记为已完成
		ID:    uint(id),
		State: 0,
	}

	r := initRouter()
	w := httptest.NewRecorder()

	reqBody, _ := json.Marshal(&task)
	req, _ := http.NewRequest(
		"PUT",
		fmt.Sprintf("/v1/todolist/%d", id),
		bytes.NewBuffer(reqBody),
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

func TestPutTodoItemReturn500(t *testing.T) {

}

func TestDeleteTodoItemReturn200(t *testing.T) {
	id := insertTask()
	r := initRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/v1/todolist/%d", id),
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteTodoItemReturn404(t *testing.T) {
	id := rand.Intn(1000)
	r := initRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/v1/todolist/%d", id),
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 404, w.Code)
}

func insertTask() int {
	n := rand.Intn(1000)
	db.Db.Exec("insert into tasks (id, task_info, state) values (?, 'just test', 1)", n)
	return n
}
