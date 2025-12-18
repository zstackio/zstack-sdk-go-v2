// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetL3NetworkTypes 获取L3NetworkTypes详情
func (cli *ZSClient) GetL3NetworkTypes(uuid string) (*view.GetL3NetworkTypesView, error) {
	var resp view.GetL3NetworkTypesView
	if err := cli.Get("v1/l3-networks/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

