// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAliyunNasPrimaryStorage adds AliyunNasPrimaryStorage
func (cli *ZSClient) AddAliyunNasPrimaryStorage(params param.AddAliyunNasPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/aliyun/nas", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
