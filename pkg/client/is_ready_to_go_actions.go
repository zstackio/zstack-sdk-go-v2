// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// IsReadyToGo operates on IsReadyToGo
func (cli *ZSClient) IsReadyToGo(params param.IsReadyToGoParam) (*view.IsReadyToGoView, error) {
	var resp view.IsReadyToGoView
	if err := cli.Get("v1/management-nodes/ready", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
