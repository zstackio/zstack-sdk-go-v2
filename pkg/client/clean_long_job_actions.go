// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CleanLongJob operates on CleanLongJob
func (cli *ZSClient) CleanLongJob(uuid string, params param.CleanLongJobParam) (*view.CleanLongJobEventView, error) {
	resp := view.CleanLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
