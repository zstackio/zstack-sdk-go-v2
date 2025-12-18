// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunNasPrimaryStorage 操作AddAliyunNasPrimaryStorage
func (cli *ZSClient) AddAliyunNasPrimaryStorage(params param.AddAliyunNasPrimaryStorageParam) (*view.AddPrimaryStorageEventView, error) {
	resp := view.AddPrimaryStorageEventView{}
	if err := cli.Post("v1/primary-storage/aliyun/nas", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

