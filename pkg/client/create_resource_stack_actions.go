// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateResourceStack creates ResourceStack
func (cli *ZSClient) CreateResourceStack(params param.CreateResourceStackParam) (*view.CreateResourceStackEventView, error) {
	resp := view.CreateResourceStackEventView{}
	if err := cli.Post("v1/cloudformation/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
