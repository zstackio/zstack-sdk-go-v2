// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAccessControlList updates AccessControlList
func (cli *ZSClient) UpdateAccessControlList(uuid string, params param.UpdateAccessControlListParam) (*view.UpdateAccessControlListEventView, error) {
	resp := view.UpdateAccessControlListEventView{}
	if err := cli.Put("v1/access-control-lists/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
