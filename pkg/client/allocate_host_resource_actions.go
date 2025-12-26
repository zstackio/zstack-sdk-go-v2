// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AllocateHostResource operates on AllocateHostResource
func (cli *ZSClient) AllocateHostResource(params param.AllocateHostResourceParam) (*view.AllocateHostResourceEventView, error) {
	resp := view.AllocateHostResourceEventView{}
	if err := cli.Post("v1/hosts/{uuid}/allocate-resource", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
