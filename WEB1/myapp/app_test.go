package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t) // 이 라이브러리를 이용해 쉽게 테스트 가능

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler() // mux를 이용한 동적 라우팅을 적용한다.
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!! ", res.Code)
	// }
	data, _ := ioutil.ReadAll(res.Body) // buffer를 사용해 body를 읽어온다.
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t) // 이 라이브러리를 이용해 쉽게 테스트 가능

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler() // mux를 이용한 동적 라우팅을 적용한다.
	mux.ServeHTTP(res, req)

	// barHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!! ", res.Code)
	// }
	data, _ := ioutil.ReadAll(res.Body) // buffer를 사용해 body를 읽어온다.
	assert.Equal("Hello World!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t) // 이 라이브러리를 이용해 쉽게 테스트 가능

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=jk", nil)

	mux := NewHttpHandler() // mux를 이용한 동적 라우팅을 적용한다.
	mux.ServeHTTP(res, req)

	// barHandler(res, req)

	assert.Equal(http.StatusOK, res.Code)
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!! ", res.Code)
	// }
	data, _ := ioutil.ReadAll(res.Body) // buffer를 사용해 body를 읽어온다.
	assert.Equal("Hello jk!", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"jk", "last_name":"choi", "email":"sizzflyer@gmail.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	// decode
	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("jk", user.FirstName)
	assert.Equal("choi", user.LastName)
}
