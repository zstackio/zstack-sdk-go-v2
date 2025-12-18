// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetInterdependentL3NetworksImages 获取InterdependentL3NetworksImages详情
func (cli *ZSClient) GetInterdependentL3NetworksImages(uuid string) (*view.GetInterdependentL3NetworkImageView, error) {
	var resp view.GetInterdependentL3NetworkImageView
	if err := cli.Get("v1/images-l3networks/dependencies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

