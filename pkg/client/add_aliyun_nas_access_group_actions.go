// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAliyunNasAccessGroup adds AliyunNasAccessGroup
func (cli *ZSClient) AddAliyunNasAccessGroup(params param.AddAliyunNasAccessGroupParam) (*view.AddAliyunNasAccessGroupEventView, error) {
	resp := view.AddAliyunNasAccessGroupEventView{}
	if err := cli.Post("v1/nas/aliyun/access", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
