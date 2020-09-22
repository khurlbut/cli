package v7pushaction

import (
	"fmt"

	"code.cloudfoundry.org/cli/util/manifestparser"
)

func (actor Actor) HandleFlagOverrides(
	baseManifest manifestparser.Manifest,
	flagOverrides FlagOverrides,
) (manifestparser.Manifest, error) {
	fmt.Printf("handle_flag_overrides.go 1 actor.TransformManifestSequence %v\n", actor.TransformManifestSequence)
	newManifest := baseManifest

	for _, transformPlan := range actor.TransformManifestSequence {
		var err error
		newManifest, err = transformPlan(newManifest, flagOverrides)
		if err != nil {
			return manifestparser.Manifest{}, err
		}
	}

	return newManifest, nil
}
