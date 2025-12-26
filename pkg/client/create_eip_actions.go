// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEip creates Eip
func (cli *ZSClient) CreateEip(params param.CreateEipParam) (*view.CreateEipEventView, error) {
	resp := view.CreateEipEventView{}
	if err := cli.Post("v1/eips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
