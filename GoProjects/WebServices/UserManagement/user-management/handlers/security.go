package handlers

import (
	"net/http"
)

func EnableHTTPS(w http.ResponseWriter, r *http.Request) {
	// 启用HTTPS（这里省略具体实现）
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HTTPS enabled"))
}
