package v7pushaction

import "fmt"

func SetupDeploymentStrategyForPushPlan(pushPlan PushPlan, overrides FlagOverrides) (PushPlan, error) {
	fmt.Printf("actor/v7pushaction/setup_deployment_strategy_for_push_plan SetupDeploymentStrategyForPushPlan overrides.Strategy: %s\n", overrides.Strategy)
	pushPlan.Strategy = overrides.Strategy

	return pushPlan, nil
}
