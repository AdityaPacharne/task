package server

import(
    "context"
    "log"
    "sync"
    "time"
    reportservice "github.com/AdityaPacharne/task"
)

type server struct {
    reportservice.UnimplementedReportServiceServer
    mu sync.Mutex
    reports map[string]string
}

func (s *server) GenerateReport(ct context.Context, request *reportservice.SendReportRequest) (*reportservice.SendReportResponse, error){
    reportId := "reportid_" + request.UserId + "_" + time.Now();
    s.mu.Lock();
    s.reports[request.UserId] = reportId;
    s.mu.Unlock();
    log.Println("Generated report: ", request.UserID, " -> ", reportId);
    response := &reportservice.SendReportResponse{
        ReportId: reportId,
    }
    return response, nil;
}

