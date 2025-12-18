// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangePortMirrorState 操作PortMirrorState
func (cli *ZSClient) ChangePortMirrorState(uuid string, params param.ChangePortMirrorStateParam) (*view.ChangePortMirrorStateEventView, error) {
	resp := view.ChangePortMirrorStateEventView{}
	if err := cli.Put("v1/port-mirrors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

