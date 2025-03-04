package handlers

import (
	"encoding/json"
	"net/http"
)

func GetReport(w http.ResponseWriter, r *http.Request) {
	// 生成报表数据（这里省略数据库操作）
	report := Report{
		TotalUsers:    100,
		ActiveUsers:   80,
		NewUsersToday: 5,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
