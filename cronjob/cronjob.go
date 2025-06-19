package cronjob

import (
    "context"
    "log"
    "github.com/robfig/cron/v3"
    "github.com/AdityaPacharne/task/server"
    reportservice "github.com/AdityaPacharne/task/proto"
)

func StartCron(srv *server.ReportServer) {
    c := cron.New();
    c.AddFunc("@every 10s", func() {
        users := []string{"12", "23", "34", "45", "56", "67", "78", "89", "90"}
        for _, userid := range users {
            req := &reportservice.SendReportRequest(UserId: userid}
            _, err := srv.GenerateReport(context.Background(), req)
            if err != nil {
                log.Println("Cron Failed: ", userid, " Error: ", err);
            }
        }
    });
    c.Start();
}
