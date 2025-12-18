// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetElaborations 获取Elaborations详情
func (cli *ZSClient) GetElaborations(uuid string) (*view.GetElaborationsView, error) {
	var resp view.GetElaborationsView
	if err := cli.Get("v1/errorcode/elaborations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

