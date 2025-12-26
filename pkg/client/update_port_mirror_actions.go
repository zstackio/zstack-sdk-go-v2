// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePortMirror updates PortMirror
func (cli *ZSClient) UpdatePortMirror(uuid string, params param.UpdatePortMirrorParam) (*view.UpdatePortMirrorEventView, error) {
	resp := view.UpdatePortMirrorEventView{}
	if err := cli.Put("v1/port-mirrors/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
