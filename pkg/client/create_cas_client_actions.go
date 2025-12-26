// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateCasClient creates CasClient
func (cli *ZSClient) CreateCasClient(params param.CreateCasClientParam) (*view.CreateCasClientEventView, error) {
	resp := view.CreateCasClientEventView{}
	if err := cli.Post("v1/create/cas/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
