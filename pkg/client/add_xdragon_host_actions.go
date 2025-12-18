// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddXDragonHost 操作AddXDragonHost
func (cli *ZSClient) AddXDragonHost(params param.AddXDragonHostParam) (*view.AddHostEventView, error) {
	resp := view.AddHostEventView{}
	if err := cli.Post("v1/hosts/xdragon", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

