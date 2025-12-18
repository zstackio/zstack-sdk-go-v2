// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunSnapshot 更新AliyunSnapshot
func (cli *ZSClient) UpdateAliyunSnapshot(uuid string, params param.UpdateAliyunSnapshotParam) (*view.UpdateAliyunSnapshotEventView, error) {
	resp := view.UpdateAliyunSnapshotEventView{}
	if err := cli.Put("v1/hybrid/aliyun/snapshot/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

