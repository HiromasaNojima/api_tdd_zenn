package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		sub := strings.TrimPrefix(r.URL.Path, "/users/")
		_, id := filepath.Split(sub)
		// id to int
		userID, _ := strconv.Atoi(id)
		user, _ := GetUser(userID)
		response, _ := json.Marshal(user)
		w.Write(response)
	})
	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUser(id int) (User, error) {
	user, ok := database[id]
	if !ok {
		return User{}, fmt.Errorf("id %d is not found", id)
	}
	return user, nil
}

// 簡易的なデータベース
var database map[int]User = map[int]User{
	1: {ID: 1, Name: "ユーザー名", Age: 20},
	2: {ID: 2, Name: "Taro", Age: 21},
	3: {ID: 3, Name: "Jiro", Age: 22},
	// ...
}
