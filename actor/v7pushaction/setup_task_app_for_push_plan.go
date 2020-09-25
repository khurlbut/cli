package v7pushaction

import "fmt"

func SetupTaskAppForPushPlan(pushPlan PushPlan, overrides FlagOverrides) (PushPlan, error) {
	fmt.Printf("actor/v7pushaction/setup_task_app_for_push_plan.go SetupTaskAppForPushPlan overrices.Task: %t\n", overrides.Task)
	pushPlan.TaskTypeApplication = overrides.Task

	return pushPlan, nil
}
