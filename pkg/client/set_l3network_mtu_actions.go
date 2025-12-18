// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetL3NetworkMtu 操作SetL3NetworkMtu
func (cli *ZSClient) SetL3NetworkMtu(params param.SetL3NetworkMtuParam) (*view.SetL3NetworkMtuEventView, error) {
	resp := view.SetL3NetworkMtuEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/mtu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

