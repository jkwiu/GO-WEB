package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct {
}

// interface로 ServeHTTP()를 등록
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	// json을 받아서 user struct로 변환.
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()
	// 다시 json 형태로 바꾼다.
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json") // header에 content type을 json으로 명시
	w.WriteHeader(http.StatusCreated)
	// data는 byte이므로 string으로 변환
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	// function을 등록
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	// 인스턴스를 등록
	// json으로 받아서 json으로 변환해주는 코드
	mux.Handle("/foo", &fooHandler{})
	return mux
}
