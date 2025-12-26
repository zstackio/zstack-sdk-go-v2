// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SubmitLongJob operates on SubmitLongJob
func (cli *ZSClient) SubmitLongJob(params param.SubmitLongJobParam) (*view.SubmitLongJobEventView, error) {
	resp := view.SubmitLongJobEventView{}
	if err := cli.Post("v1/longjobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
