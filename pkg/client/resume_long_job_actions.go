// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ResumeLongJob 操作ResumeLongJob
func (cli *ZSClient) ResumeLongJob(uuid string, params param.ResumeLongJobParam) (*view.ResumeLongJobEventView, error) {
	resp := view.ResumeLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

