package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    reportservice "github.com/AdityaPacharne/task"
    "github.com/AdityaPacharne/task/server"
)

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
