// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunMountTarget updates AliyunMountTarget
func (cli *ZSClient) UpdateAliyunMountTarget(uuid string, params param.UpdateAliyunMountTargetParam) (*view.UpdateNasMountTargetEventView, error) {
	resp := view.UpdateNasMountTargetEventView{}
	if err := cli.Put("v1/nas/aliyun/mount", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
