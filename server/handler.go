package server

import(
    "context"
    "log"
    "sync"
    "time"
    reportservice "github.com/AdityaPacharne/task/proto"
)

type ReportServer struct {
    reportservice.UnimplementedReportServiceServer
    mu sync.Mutex
    Reports map[string]string
}

func (s *ReportServer) GenerateReport(ct context.Context, request *reportservice.SendReportRequest) (*reportservice.SendReportResponse, error){
    reportId := "reportid_" + request.UserId + "_" + time.Now().Format("150505");
    s.mu.Lock();
    s.Reports[request.UserId] = reportId;
    s.mu.Unlock();
    log.Println("Generated report: ", request.UserId, " -> ", reportId);
    response := &reportservice.SendReportResponse{
        ReportId: reportId,
    }
    return response, nil;
}

