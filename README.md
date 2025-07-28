# Go API Project Structure

Đây là cấu trúc thư mục chuẩn cho một dự án Go API. Dưới đây là giải thích chi tiết về mục đích và nội dung của từng thư mục:

## Cấu trúc thư mục

```
go-api/
├── api/        # Chứa API definitions (Swagger/OpenAPI specs, protocol definitions)
├── cmd/        # Chứa mã nguồn cho các điểm vào chính của ứng dụng
├── configs/    # Chứa các file cấu hình
├── internal/   # Chứa mã nguồn private của ứng dụng
├── pkg/        # Chứa mã nguồn có thể được sử dụng bởi các ứng dụng bên ngoài
└── test/       # Chứa các file test bổ sung và dữ liệu test
```

### Chi tiết từng thư mục

#### `/api`

- Chứa các file định nghĩa API như Swagger/OpenAPI specs
- Chứa các protocol definition files (proto files nếu sử dụng gRPC)
- File schema, JSON definitions

#### `/cmd`

- Thư mục chính chứa mã nguồn cho điểm vào của ứng dụng
- Tên của thư mục con trong cmd nên khớp với tên của executable muốn build
- Ví dụ: `cmd/myapp/main.go`

#### `/configs`

- Chứa các template file cấu hình hoặc các file cấu hình mặc định
- Ví dụ: `config.yaml`, `config.json`

#### `/internal`

- Chứa mã nguồn private của ứng dụng
- Code trong thư mục này không thể được import bởi các ứng dụng hoặc thư viện khác
- Thường chứa:
  - `/internal/app/`: logic ứng dụng
  - `/internal/pkg/`: các package dùng nội bộ
  - `/internal/types/`: các type và interface

#### `/pkg`

- Chứa code có thể được sử dụng bởi các ứng dụng bên ngoài
- Các package trong thư mục này nên được thiết kế để có thể được imported và sử dụng bởi các dự án khác
- Ví dụ: client libraries, utility functions

#### `/test`

- Chứa các file test bổ sung và dữ liệu test
- Có thể chứa các large test suites
- Ví dụ: `/test/integration/`, `/test/performance/`

## Quy ước đặt tên

- Sử dụng snake_case cho tên file
- Sử dụng camelCase cho tên hàm và biến trong Go
- Sử dụng PascalCase cho exported names (public)

## Tổ chức code

- Mỗi package nên có một file riêng cho tests (`xxx_test.go`)
- Các interface nên được định nghĩa gần nơi sử dụng
- Package main càng nhỏ càng tốt
# go
