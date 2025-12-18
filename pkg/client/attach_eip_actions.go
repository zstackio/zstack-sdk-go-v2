// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachEip operates on Eip
func (cli *ZSClient) AttachEip(params param.AttachEipParam) (*view.AttachEipEventView, error) {
	resp := view.AttachEipEventView{}
	if err := cli.Post("v1/eips/{eipUuid}/vm-instances/nics/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
