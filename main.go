package main

import (
	"flag"
	"notification-server/notificationserver"
	"os"

	"github.com/armosec/utils-k8s-go/probes"
	logger "github.com/dwertent/go-logger"
	"github.com/dwertent/go-logger/helpers"
)

func main() {
	flag.Parse()

	displayBuildTag()

	isReadinessReady := false
	go probes.InitReadinessV1(&isReadinessReady)

	ns := notificationserver.NewNotificationServer()
	isReadinessReady = true
	ns.SetupNotificationServer()
}

// DisplayBuildTag display on startup
func displayBuildTag() {
	logger.L().Info("Image version", helpers.String("release", os.Getenv("RELEASE")))
}
