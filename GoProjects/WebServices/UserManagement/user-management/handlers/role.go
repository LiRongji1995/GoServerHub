package handlers

import (
	"encoding/json"
	"net/http"
)

func AssignRole(w http.ResponseWriter, r *http.Request) {
	var assignment struct {
		UserID int `json:"user_id"`
		RoleID int `json:"role_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 分配角色（这里省略数据库操作）
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(assignment)
}
