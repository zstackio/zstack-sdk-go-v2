// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNasMountTarget 查询NasMountTarget列表
func (cli *ZSClient) QueryNasMountTarget(params param.QueryParam) ([]view.QueryNasMountTargetView, error) {
	var resp []view.QueryNasMountTargetView
	return resp, cli.List("v1/primary-storage/nas/mount", &params, &resp)
}

