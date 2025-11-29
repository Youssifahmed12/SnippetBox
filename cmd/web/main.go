package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
    
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    
    app := &application{
        logger: logger,
    }
    
    addr := flag.String("addr", ":4000", "HTTP network address")
    flag.Parse()
    
	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
