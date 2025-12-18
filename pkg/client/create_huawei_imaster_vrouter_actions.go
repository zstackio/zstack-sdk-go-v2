// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateHuaweiIMasterVRouter creates HuaweiIMasterVRouter
func (cli *ZSClient) CreateHuaweiIMasterVRouter(params param.CreateHuaweiIMasterVRouterParam) (*view.CreateHuaweiIMasterVRouterEventView, error) {
	resp := view.CreateHuaweiIMasterVRouterEventView{}
	if err := cli.Post("v1/sdn-controller/huawei-imaster/vrouters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
