// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetL2NetworkTypes 获取L2NetworkTypes详情
func (cli *ZSClient) GetL2NetworkTypes(uuid string) (*view.GetL2NetworkTypesView, error) {
	var resp view.GetL2NetworkTypesView
	if err := cli.Get("v1/l2-networks/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

