package notificationserver

import (
	"notification-server/cautils"
	"os"
	"strings"

	"github.com/golang/glog"
)

var (
	// MASTER_ATTRIBUTES attributes master is expecting
	MASTER_ATTRIBUTES []string
	// MASTER_HOST -
	MASTER_HOST string
)

func SetupMasterInfo() {
	att, k1 := os.LookupEnv("MASTER_NOTIFICATION_SERVER_ATTRIBUTES")
	host, k0 := os.LookupEnv("MASTER_NOTIFICATION_SERVER_HOST")
	if !k0 || !k1 {
		return
	}
	MASTER_HOST = host
	MASTER_ATTRIBUTES = strings.Split(att, ";")
	if MASTER_ATTRIBUTES[len(MASTER_ATTRIBUTES)-1] == "" {
		cautils.RemoveIndexFromStringSlice(&MASTER_ATTRIBUTES, len(MASTER_ATTRIBUTES)-1)
	}
	glog.Infof("master host: %s, masterattribites: %v", MASTER_HOST, MASTER_ATTRIBUTES)
}

// IsMaster is server master or edge
func IsMaster() bool {
	return (len(MASTER_ATTRIBUTES) == 0 || MASTER_HOST == "")
}
