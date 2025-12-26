// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetL2NetworkTypes gets L2NetworkTypes by uuid
func (cli *ZSClient) GetL2NetworkTypes(uuid string) (*view.GetL2NetworkTypesView, error) {
	var resp view.GetL2NetworkTypesView
	if err := cli.Get("v1/l2-networks/types", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
