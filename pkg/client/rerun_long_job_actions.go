// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RerunLongJob 操作RerunLongJob
func (cli *ZSClient) RerunLongJob(uuid string, params param.RerunLongJobParam) (*view.RerunLongJobEventView, error) {
	resp := view.RerunLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

