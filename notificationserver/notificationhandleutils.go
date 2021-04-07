package notificationserver

import (
	"os"

	"asterix.cyberarmor.io/cyberarmor/capacketsgo/notificationserver"
	"github.com/golang/glog"
)

var (
	// MASTER_ATTRIBUTES attributes master is expecting
	MASTER_ATTRIBUTES []string
	// MASTER_HOST -
	MASTER_HOST string
)

func SetupMasterInfo() {
	// att, k1 := os.LookupEnv("MASTER_NOTIFICATION_SERVER_ATTRIBUTES")
	host, k0 := os.LookupEnv("MASTER_NOTIFICATION_SERVER_HOST")
	if !k0 {
		glog.Infof("Running notification server as master")
		return
	}
	MASTER_HOST = host
	// if MASTER_ATTRIBUTES[len(MASTER_ATTRIBUTES)-1] == "" {
	// 	cautils.RemoveIndexFromStringSlice(&MASTER_ATTRIBUTES, len(MASTER_ATTRIBUTES)-1)
	// }
	// if len(MASTER_ATTRIBUTES) == 0 {
	MASTER_ATTRIBUTES = []string{notificationserver.TargetCustomer, "customer"} // agent uses customer
	// }
	glog.Infof("master host: %s, master attributes: %v", MASTER_HOST, MASTER_ATTRIBUTES)
}

// IsMaster is server master or edge
func IsMaster() bool {
	return MASTER_HOST == ""
}
