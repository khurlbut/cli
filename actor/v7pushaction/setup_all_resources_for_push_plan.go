package v7pushaction

import (
	"errors"
	"fmt"
	"os"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
)

func (actor Actor) SetupAllResourcesForPushPlan(pushPlan PushPlan, overrides FlagOverrides) (PushPlan, error) {
	fmt.Printf("actor/v7pushaction/setup_all_resources_for_push_plan.go SetupAllResourcesForPushPlan 1\n")
	if pushPlan.DropletPath != "" {
		return pushPlan, nil
	}

	fmt.Printf("actor/v7pushaction/setup_all_resources_for_push_plan.go SetupAllResourcesForPushPlan 2\n")
	if pushPlan.Application.LifecycleType == constant.AppLifecycleTypeDocker {
		return pushPlan, nil
	}

	fmt.Printf("actor/v7pushaction/setup_all_resources_for_push_plan.go SetupAllResourcesForPushPlan 3\n")
	path := pushPlan.BitsPath
	if path == "" {
		return PushPlan{}, errors.New("developer error: Bits Path needs to be set prior to generating app resources")
	}

	fmt.Printf("actor/v7pushaction/setup_all_resources_for_push_plan.go SetupAllResourcesForPushPlan 4\n")
	info, err := os.Stat(path)
	if err != nil {
		return PushPlan{}, err
	}

	var archive bool
	var resources []sharedaction.Resource
	if info.IsDir() {
		fmt.Printf("actor/v7pushaction/setup_all_resources_for_push_plan.go SetupAllResourcesForPushPlan 5\n")
		resources, err = actor.SharedActor.GatherDirectoryResources(path)
	} else {
		fmt.Printf("actor/v7pushaction/setup_all_resources_for_push_plan.go SetupAllResourcesForPushPlan 6\n")
		archive = true
		resources, err = actor.SharedActor.GatherArchiveResources(path)
	}
	if err != nil {
		return PushPlan{}, err
	}

	var v3Resources []sharedaction.V3Resource
	for _, resource := range resources {
		v3Resources = append(v3Resources, resource.ToV3Resource())
	}

	pushPlan.Archive = archive
	pushPlan.AllResources = v3Resources
	fmt.Println("Archive is %t", pushPlan.Archive)

	if pushPlan.AllResources == nil {
		fmt.Println("AllResources is nil!")
	} else {
		fmt.Println("AllResources is NOT nil!")
		// fmt.Println("AllResources is %T", pushPlan.AllResources)
	}

	fmt.Printf("actor/v7pushaction/setup_all_resources_for_push_plan.go SetupAllResourcesForPushPlan returning\n")
	return pushPlan, nil
}
