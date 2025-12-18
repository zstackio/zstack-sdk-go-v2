// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectHost 操作ReconnectHost
func (cli *ZSClient) ReconnectHost(uuid string, params param.ReconnectHostParam) (*view.ReconnectHostEventView, error) {
	resp := view.ReconnectHostEventView{}
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

