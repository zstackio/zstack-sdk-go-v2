// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAllEventMetadata 获取AllEventMetadata详情
func (cli *ZSClient) GetAllEventMetadata(uuid string) (*view.GetAllEventMetadataView, error) {
	var resp view.GetAllEventMetadataView
	if err := cli.Get("v1/zwatch/events/meta-data", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

