// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateL3Network 更新L3Network
func (cli *ZSClient) UpdateL3Network(uuid string, params param.UpdateL3NetworkParam) (*view.UpdateL3NetworkEventView, error) {
	resp := view.UpdateL3NetworkEventView{}
	if err := cli.Put("v1/l3-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

