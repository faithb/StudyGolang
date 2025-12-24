# StudyGolang

Kho skeleton Todo sử dụng Gin + GORM để thực hành những phần cốt lõi của Golang.

## Lộ trình nền tảng Golang (cốt lõi)
- Cú pháp căn bản: khai báo biến/hằng, hàm, struct, slice/map, pointer.
- Điều khiển luồng: if/else, switch, for (range), defer, panic & recover.
- Xử lý lỗi chuẩn: `error`, gói `errors`, wrapping với `fmt.Errorf`.
- Làm việc với module: `go mod init`, `go get`, `go mod tidy`, cấu trúc project.
- Concurrency: goroutine, channel (buffered/unbuffered), `select`, `context` để hủy & timeout, nhóm với `sync.WaitGroup`.
- Chuẩn hóa code: `go fmt`, `go vet`, phân tách interface nhỏ, viết comment dạng GoDoc.
- Standard lib quan trọng: `fmt`, `net/http`, `encoding/json`, `time`, `io`/`os`.
- Testing cơ bản: `go test ./...`, table-driven test, mock đơn giản bằng interface.

## Chạy nhanh dự án mẫu
```bash
DB_CONNECT="user:pass@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local" go run main.go
```

API mẫu chạy tại `:3001` với các endpoint CRUD `/v1/items` và `/ping`.
