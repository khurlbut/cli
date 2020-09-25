package v7pushaction

import "fmt"

func SetupNoWaitForPushPlan(pushPlan PushPlan, overrides FlagOverrides) (PushPlan, error) {
	fmt.Printf("actor/v7pushaction/setup_no_wait_for_push_plan.go SetupNoWaitForPushPlan overrices.NoWait: %t\n", overrides.NoWait)
	pushPlan.NoWait = overrides.NoWait

	return pushPlan, nil
}
