// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateHybridEip creates HybridEip
func (cli *ZSClient) CreateHybridEip(params param.CreateHybridEipParam) (*view.CreateHybridEipEventView, error) {
	resp := view.CreateHybridEipEventView{}
	if err := cli.Post("v1/hybrid/eip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
