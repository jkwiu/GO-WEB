package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	// file mode는 읽기 권한을 어떻게 줄것이냐
	os.MkdirAll(dirname, 0777)
	// file name은 헤더에 있다.
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	// file을 만들면 항상 닫아줘야 한다.(file을 만들기 위해서는 file의 handle을 사용하는데 os자원이므로 닫아주지 않으면 문제가 생긴다.)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	// 아래의 file은 비어있는 file이고 file handler이기 때문에 upload file을 file handle에 복사해줘야 한다.
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

// 가장 고전적인 파일서버
func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
