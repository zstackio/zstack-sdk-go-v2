// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RequestConsoleAccess operates on RequestConsoleAccess
func (cli *ZSClient) RequestConsoleAccess(params param.RequestConsoleAccessParam) (*view.RequestConsoleAccessEventView, error) {
	resp := view.RequestConsoleAccessEventView{}
	if err := cli.Post("v1/consoles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
