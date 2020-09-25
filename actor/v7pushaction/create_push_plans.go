package v7pushaction

import (
	"fmt"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/util/manifestparser"
)

// CreatePushPlans returns a set of PushPlan objects based off the inputs
// provided. It's assumed that all flag and argument and manifest combinations
// have been validated prior to calling this function.
func (actor Actor) CreatePushPlans(
	spaceGUID string,
	orgGUID string,
	manifest manifestparser.Manifest,
	overrides FlagOverrides,
) ([]PushPlan, v7action.Warnings, error) {
	fmt.Printf("actor/v7pushaction/create_push_plans.go CreatePushPlans 1\n")
	var pushPlans []PushPlan

	return pushPlans, nil, nil

	apps, warnings, err := actor.V7Actor.GetApplicationsByNamesAndSpace(manifest.AppNames(), spaceGUID)
	if err != nil {
		return nil, warnings, err
	}
	fmt.Printf("actor/v7pushaction/create_push_plans.go CreatePushPlans 2\n")
	nameToApp := actor.generateAppNameToApplicationMapping(apps)

	fmt.Printf("actor/v7pushaction/create_push_plans.go CreatePushPlans 3 manifest.Applications:  %T %+v\n", manifest.Applications, manifest.Applications)
	for _, manifestApplication := range manifest.Applications {
		fmt.Printf("actor/v7pushaction/create_push_plans.go CreatePushPlans 4\n")
		plan := PushPlan{
			OrgGUID:     "",
			SpaceGUID:   "",
			Application: nameToApp[manifestApplication.Name],
			BitsPath:    manifestApplication.Path,
		}

		fmt.Printf("\nactor/v7pushaction/create_push_plans.go CreatePushPlans 5:\n--- plan ---\n%+v\n---\n\n", plan)

		if manifestApplication.Docker != nil {
			plan.DockerImageCredentials = v7action.DockerImageCredentials{
				Path:     manifestApplication.Docker.Image,
				Username: manifestApplication.Docker.Username,
				Password: overrides.DockerPassword,
			}
		}

		/* !!! KDH !!!
		See:

				actor/7vpushaction/action.go

		*/
		// List of PreparePushPlanSequence is defined in NewActor
		for _, updatePlan := range actor.PreparePushPlanSequence {
			var err error
			plan, err = updatePlan(plan, overrides)
			if err != nil {
				return nil, warnings, err
			}
		}

		fmt.Printf("\nactor/v7pushaction/create_push_plans.go CreatePushPlans 6:\n--- plan ---\n%+v\n---\n\n", plan)

		pushPlans = append(pushPlans, plan)
	}

	fmt.Printf("actor/v7pushaction/create_push_plans.go CreatePushPlans returning:\n--- pushPlans ---\n%+v\n---\n\n", pushPlans)
	return pushPlans, warnings, nil
}

func (actor Actor) generateAppNameToApplicationMapping(applications []resources.Application) map[string]resources.Application {
	fmt.Printf("actor/v7pushaction/create_push_plans.go generateAppNameToApplicationMapping 1\n")
	nameToApp := make(map[string]resources.Application, len(applications))
	for _, app := range applications {
		nameToApp[app.Name] = app
	}
	fmt.Printf("actor/v7pushaction/create_push_plans.go generateAppNameToApplicationMapping nameToApp: %+v\n", nameToApp)
	return nameToApp
}
