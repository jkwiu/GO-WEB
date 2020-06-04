package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
