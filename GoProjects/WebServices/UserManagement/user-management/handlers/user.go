package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-management/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{ID: 1, Username: "admin", Email: "admin@example.com", RoleID: 1},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	_, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	// 更新用户信息（这里省略数据库操作）
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}
