// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIscsiServer 操作AddIscsiServer
func (cli *ZSClient) AddIscsiServer(params param.AddIscsiServerParam) (*view.AddIscsiServerEventView, error) {
	resp := view.AddIscsiServerEventView{}
	if err := cli.Post("v1/storage-devices/iscsi/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

