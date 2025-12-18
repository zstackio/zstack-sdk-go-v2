// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateObservabilityServer 创建ObservabilityServer
func (cli *ZSClient) CreateObservabilityServer(params param.CreateObservabilityServerParam) (*view.CreateObservabilityServerEventView, error) {
	resp := view.CreateObservabilityServerEventView{}
	if err := cli.Post("v1/observability-servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

