// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunNasAccessGroup 操作AddAliyunNasAccessGroup
func (cli *ZSClient) AddAliyunNasAccessGroup(uuid string, params param.AddAliyunNasAccessGroupParam) (*view.AddAliyunNasAccessGroupEventView, error) {
	resp := view.AddAliyunNasAccessGroupEventView{}
	if err := cli.Put("v1/nas/aliyun/access", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

