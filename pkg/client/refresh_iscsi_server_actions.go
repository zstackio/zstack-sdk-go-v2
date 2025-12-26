// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshIscsiServer operates on RefreshIscsiServer
func (cli *ZSClient) RefreshIscsiServer(params param.RefreshIscsiServerParam) (*view.RefreshIscsiServerEventView, error) {
	resp := view.RefreshIscsiServerEventView{}
	if err := cli.Post("v1/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
