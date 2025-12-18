// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckNetworkReachable 操作CheckNetworkReachable
func (cli *ZSClient) CheckNetworkReachable(params param.CheckNetworkReachableParam) (*view.CheckNetworkReachableView, error) {
	var resp view.CheckNetworkReachableView
	if err := cli.Get("v1/zops/check/network", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

