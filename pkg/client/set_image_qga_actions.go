// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetImageQga 操作SetImageQga
func (cli *ZSClient) SetImageQga(uuid string, params param.SetImageQgaParam) (*view.SetImageQgaEventView, error) {
	resp := view.SetImageQgaEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

