// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddVCenter 操作AddVCenter
func (cli *ZSClient) AddVCenter(params param.AddVCenterParam) (*view.AddVCenterEventView, error) {
	resp := view.AddVCenterEventView{}
	if err := cli.Post("v1/vcenters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

