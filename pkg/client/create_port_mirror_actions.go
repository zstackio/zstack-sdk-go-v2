// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePortMirror creates PortMirror
func (cli *ZSClient) CreatePortMirror(params param.CreatePortMirrorParam) (*view.CreatePortMirrorEventView, error) {
	resp := view.CreatePortMirrorEventView{}
	if err := cli.Post("v1/port-mirrors", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
