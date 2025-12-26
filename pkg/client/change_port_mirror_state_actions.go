// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangePortMirrorState changes PortMirrorState
func (cli *ZSClient) ChangePortMirrorState(uuid string, params param.ChangePortMirrorStateParam) (*view.ChangePortMirrorStateEventView, error) {
	resp := view.ChangePortMirrorStateEventView{}
	if err := cli.Put("v1/port-mirrors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
