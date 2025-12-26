// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAccessControlList creates AccessControlList
func (cli *ZSClient) CreateAccessControlList(params param.CreateAccessControlListParam) (*view.CreateAccessControlListEventView, error) {
	resp := view.CreateAccessControlListEventView{}
	if err := cli.Post("v1/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
