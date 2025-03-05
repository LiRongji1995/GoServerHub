package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func AssignRole(w http.ResponseWriter, r *http.Request) {
	var assignment struct {
		UserID int `json:"user_id"`
		RoleID int `json:"role_id"`
	}

	// 解析请求体
	if err := json.NewDecoder(r.Body).Decode(&assignment); err != nil {
		log.Printf("Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 验证输入数据
	if assignment.UserID <= 0 || assignment.RoleID <= 0 {
		log.Printf("Invalid UserID or RoleID: UserID=%d, RoleID=%d", assignment.UserID, assignment.RoleID)
		http.Error(w, "Invalid UserID or RoleID", http.StatusBadRequest)
		return
	}

	// 分配角色（这里省略数据库操作）
	// 例如：将角色ID分配给用户ID，并更新数据库中的用户角色信息
	log.Printf("Assigning RoleID=%d to UserID=%d", assignment.RoleID, assignment.UserID)

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Role assigned successfully",
		"data":    assignment,
	})
}
