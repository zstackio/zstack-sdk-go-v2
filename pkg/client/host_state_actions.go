// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeHostState 操作HostState
func (cli *ZSClient) ChangeHostState(uuid string, params param.ChangeHostStateParam) (*view.ChangeHostStateEventView, error) {
	resp := view.ChangeHostStateEventView{}
	if err := cli.Put("v1/hosts/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

