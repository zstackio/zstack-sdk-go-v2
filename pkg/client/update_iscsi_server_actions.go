// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateIscsiServer updates IscsiServer
func (cli *ZSClient) UpdateIscsiServer(uuid string, params param.UpdateIscsiServerParam) (*view.UpdateIscsiServerEventView, error) {
	resp := view.UpdateIscsiServerEventView{}
	if err := cli.Put("v1/storage-devices/iscsi/servers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
