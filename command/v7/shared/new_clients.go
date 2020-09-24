package shared

import (
	"fmt"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	ccWrapper "code.cloudfoundry.org/cli/api/cloudcontroller/wrapper"
	"code.cloudfoundry.org/cli/api/router"
	routingWrapper "code.cloudfoundry.org/cli/api/router/wrapper"
	"code.cloudfoundry.org/cli/api/uaa"
	uaaWrapper "code.cloudfoundry.org/cli/api/uaa/wrapper"
	"code.cloudfoundry.org/cli/command"
)

func GetNewClientsAndConnectToCF(config command.Config, ui command.UI, minVersionV3 string) (*ccv3.Client, *uaa.Client, *router.Client, error) {
	fmt.Printf("new_clients.go GetNewClientsAndConnecToCR 1\n")
	var err error

	ccClient, authWrapper := NewWrappedCloudControllerClient(config, ui)

	ccClient, err = connectToCF(config, ui, ccClient, minVersionV3)
	fmt.Printf("new_clients.go GetNewClientsAndConnecToCR 2\n")
	if err != nil {
		return nil, nil, nil, err
	}

	uaaClient, err := newWrappedUAAClient(config, ui, ccClient, authWrapper)
	fmt.Printf("new_clients.go GetNewClientsAndConnecToCR 3\n")
	if err != nil {
		return nil, nil, nil, err
	}

	routingClient, err := newWrappedRoutingClient(config, ui, uaaClient)

	fmt.Printf("new_clients.go GetNewClientsAndConnecToCR 4\n")
	return ccClient, uaaClient, routingClient, err
}

func NewWrappedCloudControllerClient(config command.Config, ui command.UI) (*ccv3.Client, *ccWrapper.UAAAuthentication) {
	ccWrappers := []ccv3.ConnectionWrapper{}

	verbose, location := config.Verbose()
	if verbose {
		ccWrappers = append(ccWrappers, ccWrapper.NewRequestLogger(ui.RequestLoggerTerminalDisplay()))
	}
	if location != nil {
		ccWrappers = append(ccWrappers, ccWrapper.NewRequestLogger(ui.RequestLoggerFileWriter(location)))
	}

	authWrapper := ccWrapper.NewUAAAuthentication(nil, config)

	ccWrappers = append(ccWrappers, authWrapper)
	ccWrappers = append(ccWrappers, ccWrapper.NewRetryRequest(config.RequestRetryCount()))

	ccClient := ccv3.NewClient(ccv3.Config{
		AppName:            config.BinaryName(),
		AppVersion:         config.BinaryVersion(),
		JobPollingTimeout:  config.OverallPollingTimeout(),
		JobPollingInterval: config.PollingInterval(),
		Wrappers:           ccWrappers,
	})
	return ccClient, authWrapper
}

func newWrappedUAAClient(config command.Config, ui command.UI, ccClient *ccv3.Client, authWrapper *ccWrapper.UAAAuthentication) (*uaa.Client, error) {
	fmt.Printf("new_clients.go: newWrappedUAAClient 1\n")
	// var err error
	verbose, location := config.Verbose()

	uaaClient := uaa.NewClient(config)
	if verbose {
		uaaClient.WrapConnection(uaaWrapper.NewRequestLogger(ui.RequestLoggerTerminalDisplay()))
	}
	if location != nil {
		uaaClient.WrapConnection(uaaWrapper.NewRequestLogger(ui.RequestLoggerFileWriter(location)))
	}

	uaaAuthWrapper := uaaWrapper.NewUAAAuthentication(uaaClient, config)
	uaaClient.WrapConnection(uaaAuthWrapper)
	uaaClient.WrapConnection(uaaWrapper.NewRetryRequest(config.RequestRetryCount()))

	// KDH: Maybe repurpose this to do Azure login?
	// err = uaaClient.SetupResources(ccClient.Login())
	// fmt.Printf("new_clients.go: newWrappedUAAClient 2\n")
	// if err != nil {
	// 	return nil, err
	// }

	uaaAuthWrapper.SetClient(uaaClient)
	authWrapper.SetClient(uaaClient)

	fmt.Printf("new_clients.go: newWrappedUAAClient 3")
	return uaaClient, nil
}

func newWrappedRoutingClient(config command.Config, ui command.UI, uaaClient *uaa.Client) (*router.Client, error) {
	routingConfig := router.Config{
		AppName:    config.BinaryName(),
		AppVersion: config.BinaryVersion(),
		ConnectionConfig: router.ConnectionConfig{
			DialTimeout:       config.DialTimeout(),
			SkipSSLValidation: config.SkipSSLValidation(),
		},
		RoutingEndpoint: config.RoutingEndpoint(),
	}

	routingWrappers := []router.ConnectionWrapper{routingWrapper.NewErrorWrapper()}

	verbose, location := config.Verbose()

	if verbose {
		routingWrappers = append(routingWrappers, routingWrapper.NewRequestLogger(ui.RequestLoggerTerminalDisplay()))
	}

	if location != nil {
		routingWrappers = append(routingWrappers, routingWrapper.NewRequestLogger(ui.RequestLoggerFileWriter(location)))
	}

	authWrapper := routingWrapper.NewUAAAuthentication(uaaClient, config)

	routingWrappers = append(routingWrappers, authWrapper)
	routingConfig.Wrappers = routingWrappers

	routingClient := router.NewClient(routingConfig)

	return routingClient, nil
}

func connectToCF(config command.Config, ui command.UI, ccClient *ccv3.Client, minVersionV3 string) (*ccv3.Client, error) {

	fmt.Printf("new_clients.go: connectToCF config.DialTimeout: %d\n", config.DialTimeout)
	ccClient.TargetCF(ccv3.TargetSettings{
		URL:               "kube-cluster-name",
		SkipSSLValidation: false,
		DialTimeout:       1000,
	})
	return ccClient, nil

	// if config.Target() == "" {
	// 	return nil, translatableerror.NoAPISetError{
	// 		BinaryName: config.BinaryName(),
	// 	}
	// }

	// _, _, err := ccClient.TargetCF(ccv3.TargetSettings{
	// 	URL:               config.Target(),
	// 	SkipSSLValidation: config.SkipSSLValidation(),
	// 	DialTimeout:       config.DialTimeout(),
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// if minVersionV3 != "" {
	// 	err = command.MinimumCCAPIVersionCheck(ccClient.CloudControllerAPIVersion(), minVersionV3)
	// 	if err != nil {
	// 		if _, ok := err.(translatableerror.MinimumCFAPIVersionNotMetError); ok {
	// 			return nil, translatableerror.V3V2SwitchError{}
	// 		}
	// 		return nil, err
	// 	}
	// }
	// return ccClient, nil
}
