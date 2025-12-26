// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetUploadImageJobDetails gets UploadImageJobDetails by uuid
func (cli *ZSClient) GetUploadImageJobDetails(uuid string) (*view.GetUploadImageJobDetailsView, error) {
	var resp view.GetUploadImageJobDetailsView
	if err := cli.Get("v1/images/upload-job/details/{imageId}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
