// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAccessControlList creates AccessControlList
func (cli *ZSClient) CreateAccessControlList(params param.CreateAccessControlListParam) (*view.CreateAccessControlListEventView, error) {
	resp := view.CreateAccessControlListEventView{}
	if err := cli.Post("v1/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
