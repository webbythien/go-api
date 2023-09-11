package test

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"github.com/google/uuid"
// )

// func BenchmarkAPIGEN(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		address := uuid.New().String()
// 		url := fmt.Sprintf("http://127.0.0.1:5000/get?address=%s", address)
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			b.Fatalf("Failed to send GET request: %v", err)
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode != http.StatusOK {
// 			b.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
// 		}

// 	}
// }
