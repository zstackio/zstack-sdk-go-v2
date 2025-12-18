// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SubmitLongJob 操作SubmitLongJob
func (cli *ZSClient) SubmitLongJob(params param.SubmitLongJobParam) (*view.SubmitLongJobEventView, error) {
	resp := view.SubmitLongJobEventView{}
	if err := cli.Post("v1/longjobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

