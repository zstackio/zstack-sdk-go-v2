// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachEip operates on Eip
func (cli *ZSClient) AttachEip(params param.AttachEipParam) (*view.AttachEipEventView, error) {
	resp := view.AttachEipEventView{}
	if err := cli.Post("v1/eips/{eipUuid}/vm-instances/nics/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
