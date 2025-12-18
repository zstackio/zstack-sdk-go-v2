// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetUploadImageJobDetails 获取UploadImageJobDetails详情
func (cli *ZSClient) GetUploadImageJobDetails(uuid string) (*view.GetUploadImageJobDetailsView, error) {
	var resp view.GetUploadImageJobDetailsView
	if err := cli.Get("v1/images/upload-job/details/{imageId}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

