package handlers

import (
	"encoding/json"
	"net/http"
	"user-management/models"
)

func GetReport(w http.ResponseWriter, r *http.Request) {
	// 生成报表数据（这里省略数据库操作）
	report := models.Report{
		TotalUsers:    100,
		ActiveUsers:   80,
		NewUsersToday: 5,
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 将报告对象编码为 JSON 并返回
	json.NewEncoder(w).Encode(report)
}
