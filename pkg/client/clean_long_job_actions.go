// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CleanLongJob 操作CleanLongJob
func (cli *ZSClient) CleanLongJob(uuid string, params param.CleanLongJobParam) (*view.CleanLongJobEventView, error) {
	resp := view.CleanLongJobEventView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

