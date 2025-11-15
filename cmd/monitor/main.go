package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AngelUNC/GhostWA/core"
	"github.com/AngelUNC/GhostWA/db"
	"github.com/AngelUNC/GhostWA/cli"
)

func main() {
	log.Println("🚀 Iniciando GhostWA v1")
	msgDB, err := db.OpenMsgDB(db.MsgDBPath)
	if err != nil {
	log.Fatalf("❌ Hubo un problema al abrir msgtore DB: %v", err)
	}
	defer msgDB.Close()

	waDB, err := db.OpenWaDB(db.WaDBPath)
	if err != nil {
	log.Printf("⚠️ No se pudo abrir wa.db: %v — Continuando sin contactos.", err)
	}
	if waDB != nil {
	defer waDB.Close()
	}

	db.PreloadAll(waDB)

	core.InitializeSnapshot(msgDB)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
	<-sigChan
	log.Println("Finalizando...")
	msgDB.Close()
	if waDB != nil { waDB.Close() }
	os.Exit(0)
	}()

	go core.WatchMessages(msgDB, waDB)

	cli.Start()
	}