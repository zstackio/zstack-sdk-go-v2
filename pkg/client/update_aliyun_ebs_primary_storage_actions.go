// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunEbsPrimaryStorage updates AliyunEbsPrimaryStorage
func (cli *ZSClient) UpdateAliyunEbsPrimaryStorage(uuid string, params param.UpdateAliyunEbsPrimaryStorageParam) (*view.UpdatePrimaryStorageEventView, error) {
	resp := view.UpdatePrimaryStorageEventView{}
	if err := cli.Put("v1/primary-storage/aliyun/ebs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
