// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ResumeLongJob operates on ResumeLongJob
func (cli *ZSClient) ResumeLongJob(uuid string, params param.ResumeLongJobParam) (*view.ResumeLongJobEventView, error) {
	resp := view.ResumeLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
