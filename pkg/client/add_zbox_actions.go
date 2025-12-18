// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddZBox adds ZBox
func (cli *ZSClient) AddZBox(params param.AddZBoxParam) (*view.AddZBoxEventView, error) {
	resp := view.AddZBoxEventView{}
	if err := cli.Post("v1/zbox", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
