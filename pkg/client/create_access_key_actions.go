// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAccessKey creates AccessKey
func (cli *ZSClient) CreateAccessKey(params param.CreateAccessKeyParam) (*view.CreateAccessKeyEventView, error) {
	resp := view.CreateAccessKeyEventView{}
	if err := cli.Post("v1/accesskeys", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
