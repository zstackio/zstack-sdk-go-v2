// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAliyunMountTarget updates AliyunMountTarget
func (cli *ZSClient) UpdateAliyunMountTarget(uuid string, params param.UpdateAliyunMountTargetParam) (*view.UpdateNasMountTargetEventView, error) {
	resp := view.UpdateNasMountTargetEventView{}
	if err := cli.Put("v1/nas/aliyun/mount", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
