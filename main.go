package main

import "net/http"

// title: ユーザー情報取得
// steps:
//   - title: get
//     protocol: http
//     request:
//       method: GET
//       url: http://localhost:8080/users/1
//     expect:
//       code: OK
//       body:
//         id: 1
//         name: ユーザー名
//         age: 20

func main() {
	http.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id":1,"name":"ユーザー名","age":20}`))
	})
	http.ListenAndServe(":8080", nil)
}