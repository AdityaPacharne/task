package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    reportservice "github.com/AdityaPacharne/task/proto"
    "github.com/AdityaPacharne/task/server"
    "github.com/AdityaPacharne/task/cronjob"
    "google.golang.org/grpc/reflection"
)

func main() {
    listener, err := net.Listen("tcp", ":9090");
    if err != nil {
        log.Fatal("Failed to Listen: ", err);
    }

    grpcServer := grpc.NewServer();
    
    serverObject := &server.ReportServer {
        Reports: make(map[string]string),
    }

    reportservice.RegisterReportServiceServer(grpcServer, serverObject);
    reflection.Register(grpcServer);

    log.Println("Server runing on :9090");

    cronjob.StartCron(serverObject);
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatal("Failed to serve: ", err);
    }
}
