(AI was used for Readme generation)

This Go-based microservice exposes a gRPC API to generate reports for users and periodically triggers reports using a cron job.

* `GenerateReport(UserId)` → returns a unique `ReportId`
* `HealthCheck()` → returns the service status
* Scheduled job every 10 seconds to generate reports for predefined users
* In-memory report storage using a thread-safe map
* Logs every operation with timestamps

## Running the Service

### Prerequisites

* Go 1.20+ installed
* `protoc` (Protocol Buffers compiler)
* `protoc-gen-go` and `protoc-gen-go-grpc`

Install with:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 1. Clone the Repository

```bash
git clone https://github.com/AdityaPacharne/task.git
cd task
```

### 2. Generate gRPC Code (if not already done)

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/report.proto
```

### 3. Run the Server

```bash
go run main.go
```

You should see logs indicating the server is running and cron reports are being generated.

## Testing gRPC Endpoints

### Health Check

Install `grpcurl`:

```bash
brew install grpcurl  # macOS
```

Run:

```bash
grpcurl -plaintext localhost:9090 reportservice.ReportService/HealthCheck
```

Expected output:

```json
{
  "status": "Serving"
}
```

### Manual GenerateReport

```bash
grpcurl -plaintext -d '{"UserId": "999"}' localhost:9090 reportservice.ReportService/GenerateReport
```

Expected output:

```json
{
  "reportId": "reportid_999_XXXXXX"
}
```

## Dependencies Used

* `google.golang.org/grpc` – gRPC server
* `google.golang.org/protobuf` – Protobuf support
* `github.com/robfig/cron/v3` – Cron job scheduling

