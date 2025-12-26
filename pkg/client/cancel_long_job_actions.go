// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CancelLongJob operates on CancelLongJob
func (cli *ZSClient) CancelLongJob(uuid string, params param.CancelLongJobParam) (*view.CancelLongJobEventView, error) {
	resp := view.CancelLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
