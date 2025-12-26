// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachServiceToObservabilityServer operates on ServiceToObservabilityServer
func (cli *ZSClient) AttachServiceToObservabilityServer(params param.AttachServiceToObservabilityServerParam) (*view.AttachServiceToObservabilityServerEventView, error) {
	resp := view.AttachServiceToObservabilityServerEventView{}
	if err := cli.Post("v1/observability-server/{observabilityServerUuid}/service", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
