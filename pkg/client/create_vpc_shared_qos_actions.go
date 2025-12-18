// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpcSharedQos creates VpcSharedQos
func (cli *ZSClient) CreateVpcSharedQos(params param.CreateVpcSharedQosParam) (*view.CreateVpcSharedQosEventView, error) {
	resp := view.CreateVpcSharedQosEventView{}
	if err := cli.Post("v1/vips/sharedqos", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
