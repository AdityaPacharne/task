package main

import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
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

func main() {
    listencer, err := net.Listen("tcp", ":9090");
    if err != nil {
        log.Fatal("Failed to Listen: ", err);
    }

    grpcServer := grpc.NewServer();
    
    serverObject := &server {
        reports: make(map[string]string),
    }

    reportservice.RegisterReportServiceServer(grcpServer, serverObject)

    log.Println("Server runing on :9090");
    err := grcpServer.Serve(listener);
    if err := grcpServer.Serve(listener); err != nil {
        log.Fatal("Failed to serve: ", err);
    }
}
