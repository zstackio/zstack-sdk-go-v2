// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateObservabilityServerOffering creates ObservabilityServerOffering
func (cli *ZSClient) CreateObservabilityServerOffering(params param.CreateObservabilityServerOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
