package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    "github.com/AdityaPacharne/task/proto"
    "github.com/AdityaPacharne/task/server"
    "github.com/AdityaPacharne/task/cronjob"
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

    cronjob.StartCron(serverObject);
    if err := grcpServer.Serve(listener); err != nil {
        log.Fatal("Failed to serve: ", err);
    }
}
