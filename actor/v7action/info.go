package v7action

import "fmt"

func (actor Actor) GetLogCacheEndpoint() (string, Warnings, error) {
	fmt.Printf("cli/actor/v7action/info.go GetLogCacheEndpoint 1\n")
	info, _, warnings, err := actor.CloudControllerClient.GetInfo()
	fmt.Printf("actor/info.go GetLogCacheEndpoint 2\n")
	if err != nil {
		return "", Warnings(warnings), err
	}
	fmt.Printf("actor/info.go GetLogCacheEndpoint 3\n")
	return info.LogCache(), Warnings(warnings), nil
}
