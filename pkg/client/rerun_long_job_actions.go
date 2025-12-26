// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RerunLongJob operates on RerunLongJob
func (cli *ZSClient) RerunLongJob(uuid string, params param.RerunLongJobParam) (*view.RerunLongJobEventView, error) {
	resp := view.RerunLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
