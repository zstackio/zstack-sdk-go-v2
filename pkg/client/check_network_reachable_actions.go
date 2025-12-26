// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckNetworkReachable operates on CheckNetworkReachable
func (cli *ZSClient) CheckNetworkReachable(params param.CheckNetworkReachableParam) (*view.CheckNetworkReachableView, error) {
	var resp view.CheckNetworkReachableView
	if err := cli.Get("v1/zops/check/network", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
