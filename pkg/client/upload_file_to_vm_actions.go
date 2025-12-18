// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UploadFileToVm 操作UploadFileToVm
func (cli *ZSClient) UploadFileToVm(params param.UploadFileToVmParam) (*view.UploadFileToVmEventView, error) {
	resp := view.UploadFileToVmEventView{}
	if err := cli.Post("v1/upload-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

