// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAliyunNasMountTarget creates AliyunNasMountTarget
func (cli *ZSClient) CreateAliyunNasMountTarget(params param.CreateAliyunNasMountTargetParam) (*view.CreateNasMountTargetEventView, error) {
	resp := view.CreateNasMountTargetEventView{}
	if err := cli.Post("v1/nas/aliyun/mount", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
