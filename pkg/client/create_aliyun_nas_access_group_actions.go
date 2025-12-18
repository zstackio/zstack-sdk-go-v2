// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunNasAccessGroup creates AliyunNasAccessGroup
func (cli *ZSClient) CreateAliyunNasAccessGroup(params param.CreateAliyunNasAccessGroupParam) (*view.CreateAliyunNasAccessGroupEventView, error) {
	resp := view.CreateAliyunNasAccessGroupEventView{}
	if err := cli.Post("v1/nas/aliyun/access", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
