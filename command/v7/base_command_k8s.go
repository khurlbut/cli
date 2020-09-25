package v7

import (
	"fmt"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/uaa"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/clock"
)

// K8BaseCommand structure
type K8BaseCommand struct {
	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       Actor

	cloudControllerClient *ccv3.Client
	uaaClient             *uaa.Client
}

// Setup the K8BaseCommand
func (cmd *K8BaseCommand) Setup(config command.Config, ui command.UI) error {
	fmt.Printf("base_command_k8s.go: Setup 1\n")
	cmd.UI = ui
	cmd.Config = config
	sharedActor := sharedaction.NewActor(config)
	cmd.SharedActor = sharedActor

	cmd.Actor = v7action.NewActor(nil, config, sharedActor, nil, nil, clock.NewClock())
	fmt.Printf("base_command_k8s.go: Setup 4\n")
	return nil
}
