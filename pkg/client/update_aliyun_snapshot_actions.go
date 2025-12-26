// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunSnapshot updates AliyunSnapshot
func (cli *ZSClient) UpdateAliyunSnapshot(uuid string, params param.UpdateAliyunSnapshotParam) (*view.UpdateAliyunSnapshotEventView, error) {
	resp := view.UpdateAliyunSnapshotEventView{}
	if err := cli.Put("v1/hybrid/aliyun/snapshot/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
