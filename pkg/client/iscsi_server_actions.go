// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIscsiServer 更新IscsiServer
func (cli *ZSClient) UpdateIscsiServer(uuid string, params param.UpdateIscsiServerParam) (*view.UpdateIscsiServerEventView, error) {
	resp := view.UpdateIscsiServerEventView{}
	if err := cli.Put("v1/storage-devices/iscsi/servers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

