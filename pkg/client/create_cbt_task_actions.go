// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateCbtTask creates CbtTask
func (cli *ZSClient) CreateCbtTask(params param.CreateCbtTaskParam) (*view.CreateCbtTaskEventView, error) {
	resp := view.CreateCbtTaskEventView{}
	if err := cli.Post("v1/cbt-task/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
