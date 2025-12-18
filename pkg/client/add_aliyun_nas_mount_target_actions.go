// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunNasMountTarget 操作AddAliyunNasMountTarget
func (cli *ZSClient) AddAliyunNasMountTarget(uuid string, params param.AddAliyunNasMountTargetParam) (*view.AddAliyunNasMountTargetEventView, error) {
	resp := view.AddAliyunNasMountTargetEventView{}
	if err := cli.Put("v1/nas/aliyun/mount", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

