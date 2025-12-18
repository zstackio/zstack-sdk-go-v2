// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreatePortMirrorSession creates PortMirrorSession
func (cli *ZSClient) CreatePortMirrorSession(params param.CreatePortMirrorSessionParam) (*view.CreatePortMirrorSessionEventView, error) {
	resp := view.CreatePortMirrorSessionEventView{}
	if err := cli.Post("v1/port-mirrors/sessions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
