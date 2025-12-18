// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateEip creates Eip
func (cli *ZSClient) CreateEip(params param.CreateEipParam) (*view.CreateEipEventView, error) {
	resp := view.CreateEipEventView{}
	if err := cli.Post("v1/eips", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
