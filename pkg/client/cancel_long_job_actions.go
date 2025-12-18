// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CancelLongJob 操作CancelLongJob
func (cli *ZSClient) CancelLongJob(uuid string, params param.CancelLongJobParam) (*view.CancelLongJobEventView, error) {
	resp := view.CancelLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

