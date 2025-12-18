// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateTag creates Tag
func (cli *ZSClient) CreateTag(params param.CreateTagParam) (*view.CreateTagEventView, error) {
	resp := view.CreateTagEventView{}
	if err := cli.Post("v1/tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
