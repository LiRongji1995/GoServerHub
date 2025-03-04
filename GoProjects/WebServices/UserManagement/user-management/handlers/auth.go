package handlers

import (
	"encoding/json"
	"net/http"
	"user-management/models" // 引入 models 包
	"user-management/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User // 使用 models.User 类型
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 验证邮箱格式和密码强度
	if !utils.ValidateEmail(user.Email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}
	if !utils.ValidatePassword(user.Password) {
		http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// 保存用户到数据库（这里省略数据库操作）
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 查询用户（这里省略数据库操作）
	user := models.User{Email: credentials.Email, Password: "hashed_password_from_db"} // 使用 models.User 类型

	// 验证密码
	if !utils.VerifyPassword(user.Password, credentials.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 生成JWT Token
	token, err := utils.GenerateJWT(user.ID, user.RoleID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
