package initdb

import (
	"os"

	"../db/graph"
)

func initGraphDB() {
	if !graph.Connect() {
		os.Exit(1)
	}
	l.Info("Connected to DGraph database!")
	if graph.VerifySchema(true) {
		l.Warn("DGraph database is already initialized!")
		return
	}
	l.Info("Creating DGraph database...")
	if graph.ApplySchema() {
		l.Info("DGraph database created.")
	} else {
		l.Error("Unable to create DGraph database.")
	}
}
