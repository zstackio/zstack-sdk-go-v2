// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshIscsiServer 操作RefreshIscsiServer
func (cli *ZSClient) RefreshIscsiServer(params param.RefreshIscsiServerParam) (*view.RefreshIscsiServerEventView, error) {
	resp := view.RefreshIscsiServerEventView{}
	if err := cli.Post("v1/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

