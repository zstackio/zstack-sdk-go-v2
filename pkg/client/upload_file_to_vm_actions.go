// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UploadFileToVm operates on UploadFileToVm
func (cli *ZSClient) UploadFileToVm(params param.UploadFileToVmParam) (*view.UploadFileToVmEventView, error) {
	resp := view.UploadFileToVmEventView{}
	if err := cli.Post("v1/upload-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
