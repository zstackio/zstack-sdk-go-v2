// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunNasAccessGroup adds AliyunNasAccessGroup
func (cli *ZSClient) AddAliyunNasAccessGroup(params param.AddAliyunNasAccessGroupParam) (*view.AddAliyunNasAccessGroupEventView, error) {
	resp := view.AddAliyunNasAccessGroupEventView{}
	if err := cli.Post("v1/nas/aliyun/access", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
