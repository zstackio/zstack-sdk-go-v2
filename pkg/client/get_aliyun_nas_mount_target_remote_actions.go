// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAliyunNasMountTargetRemote gets AliyunNasMountTargetRemote by uuid
func (cli *ZSClient) GetAliyunNasMountTargetRemote(uuid string) (*view.GetAliyunNasMountTargetRemoteView, error) {
	var resp view.GetAliyunNasMountTargetRemoteView
	if err := cli.Get("v1/nas/aliyun/mount/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
