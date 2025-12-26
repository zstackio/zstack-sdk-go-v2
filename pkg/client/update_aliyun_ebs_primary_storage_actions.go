// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunEbsPrimaryStorage updates AliyunEbsPrimaryStorage
func (cli *ZSClient) UpdateAliyunEbsPrimaryStorage(uuid string, params param.UpdateAliyunEbsPrimaryStorageParam) (*view.UpdatePrimaryStorageEventView, error) {
	resp := view.UpdatePrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/aliyun/ebs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
