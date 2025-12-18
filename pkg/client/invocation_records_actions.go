// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetInvocationRecords 获取InvocationRecords详情
func (cli *ZSClient) GetInvocationRecords(uuid string) (*view.GetInvocationRecordsView, error) {
	var resp view.GetInvocationRecordsView
	if err := cli.Get("v1/scripts/aliyun-invocations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

