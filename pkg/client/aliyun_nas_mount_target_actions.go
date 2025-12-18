// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateAliyunNasMountTarget 创建AliyunNasMountTarget
func (cli *ZSClient) CreateAliyunNasMountTarget(params param.CreateAliyunNasMountTargetParam) (*view.CreateNasMountTargetEventView, error) {
	resp := view.CreateNasMountTargetEventView{}
	if err := cli.Post("v1/nas/aliyun/mount", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

