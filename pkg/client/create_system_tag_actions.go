// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSystemTag creates SystemTag
func (cli *ZSClient) CreateSystemTag(params param.CreateSystemTagParam) (*view.CreateSystemTagEventView, error) {
	resp := view.CreateSystemTagEventView{}
	if err := cli.Post("v1/system-tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
