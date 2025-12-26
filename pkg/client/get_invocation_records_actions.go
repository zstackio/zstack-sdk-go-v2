// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetInvocationRecords gets InvocationRecords by uuid
func (cli *ZSClient) GetInvocationRecords(uuid string) (*view.GetInvocationRecordsView, error) {
	var resp view.GetInvocationRecordsView
	if err := cli.Get("v1/scripts/aliyun-invocations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
