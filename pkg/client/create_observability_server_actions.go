// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateObservabilityServer creates ObservabilityServer
func (cli *ZSClient) CreateObservabilityServer(params param.CreateObservabilityServerParam) (*view.CreateObservabilityServerEventView, error) {
	resp := view.CreateObservabilityServerEventView{}
	if err := cli.Post("v1/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
