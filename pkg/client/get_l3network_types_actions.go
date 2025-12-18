// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetL3NetworkTypes gets L3NetworkTypes by uuid
func (cli *ZSClient) GetL3NetworkTypes(uuid string) (*view.GetL3NetworkTypesView, error) {
	var resp view.GetL3NetworkTypesView
	if err := cli.Get("v1/l3-networks/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
