// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAliyunNasMountTarget adds AliyunNasMountTarget
func (cli *ZSClient) AddAliyunNasMountTarget(params param.AddAliyunNasMountTargetParam) (*view.AddAliyunNasMountTargetEventView, error) {
	resp := view.AddAliyunNasMountTargetEventView{}
	if err := cli.Post("v1/nas/aliyun/mount", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
