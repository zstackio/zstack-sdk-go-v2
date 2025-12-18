// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunNasAccessGroup 更新AliyunNasAccessGroup
func (cli *ZSClient) UpdateAliyunNasAccessGroup(uuid string, params param.UpdateAliyunNasAccessGroupParam) (*view.UpdateAliyunNasAccessGroupEventView, error) {
	resp := view.UpdateAliyunNasAccessGroupEventView{}
	if err := cli.Put("v1/nas/aliyun/access", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

