package v7pushaction

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func SetDefaultBitsPathForPushPlan(pushPlan PushPlan, overrides FlagOverrides) (PushPlan, error) {
	fmt.Printf("actor/v7pushaction/set_default_bits_path_for_push_plan.go SetDefaultBitsPathForPushPlan 1\n")
	if pushPlan.BitsPath == "" && pushPlan.DropletPath == "" && pushPlan.DockerImageCredentials.Path == "" {
		var err error
		pushPlan.BitsPath, err = os.Getwd()
		fmt.Printf("actor/v7pushaction/set_default_bits_path_for_push_plan.go SetDefaultBitsPathForPushPlan 2 BitsPath set to: %s\n", pushPlan.BitsPath)
		log.WithField("path", pushPlan.BitsPath).Debug("using current directory for bits path")
		if err != nil {
			return pushPlan, err
		}
	}
	return pushPlan, nil
}
