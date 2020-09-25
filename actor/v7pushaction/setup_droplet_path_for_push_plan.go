package v7pushaction

import "fmt"

func SetupDropletPathForPushPlan(pushPlan PushPlan, overrides FlagOverrides) (PushPlan, error) {
	fmt.Printf("actor/v7pushaction/setup_droplet_path_for_push_plan.go SetupDroplentPathForPushPlan 1 setting pushPlan.DroplentPath to: %s", pushPlan.DropletPath)
	pushPlan.DropletPath = overrides.DropletPath

	return pushPlan, nil
}
