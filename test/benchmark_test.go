package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func BenchmarkAPISAVE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		requestBody := map[string]interface{}{
			"referral_code": "kawasaki.meme",
			"domain":        "kawasaki696969.meme",
			"price":         1.2,
			"address":       "0x8f9d9aA7B313cf9360d4E61D1Ae809443f97aCad",
		}

		jsonData, _ := json.Marshal(requestBody)
		url := "http://127.0.0.1:5000/save"
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Lỗi khi gửi yêu cầu:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}
	}
}
