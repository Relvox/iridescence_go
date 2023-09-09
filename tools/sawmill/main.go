package main

import (
	"flag"
	"log/slog"

	"github.com/relvox/iridescence_go/logging"
	"github.com/relvox/iridescence_go/middleware"
	"github.com/relvox/iridescence_go/servers"
	sm_serv "github.com/relvox/iridescence_go/tools/sawmill/servers"
)

func main() {
	addrFlag := flag.String("addr", ":8080", "Address including port to listen on")
	dirFlag := flag.String("dir", ".", "Directory to serve logs from")
	logFlag := flag.String("log", "self.log", "name of own log")
	flag.Parse()
	log := logging.NewLogger(*logFlag, slog.LevelInfo, slog.LevelDebug)

	sawmillServ := sm_serv.NewSawmillServer(*dirFlag, "./templates/", log)
	sfServ := sm_serv.StaticFileServer("./static/")

	servers.ConfigureAndListen(*addrFlag,
		servers.DefaultHeaders, []string{"*"}, servers.DefaultMethods,
		log, middleware.AllOptions-middleware.UserAgent-middleware.Response-middleware.RequestID,
		sawmillServ,
		&sfServ,
	)
}
