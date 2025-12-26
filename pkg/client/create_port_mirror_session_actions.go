// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePortMirrorSession creates PortMirrorSession
func (cli *ZSClient) CreatePortMirrorSession(params param.CreatePortMirrorSessionParam) (*view.CreatePortMirrorSessionEventView, error) {
	resp := view.CreatePortMirrorSessionEventView{}
	if err := cli.Post("v1/port-mirrors/sessions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
