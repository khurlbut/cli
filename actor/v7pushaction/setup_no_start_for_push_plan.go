package v7pushaction

import "fmt"

func SetupNoStartForPushPlan(pushPlan PushPlan, overrides FlagOverrides) (PushPlan, error) {
	fmt.Printf("actor/v7pushaction/setup_no_start_for_push_plan.go SetupNoStartForPushPlan overrices.NoStart: %t\n", overrides.NoStart)
	pushPlan.NoStart = overrides.NoStart

	return pushPlan, nil
}
