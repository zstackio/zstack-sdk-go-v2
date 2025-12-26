// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetElaborations gets Elaborations by uuid
func (cli *ZSClient) GetElaborations(uuid string) (*view.GetElaborationsView, error) {
	var resp view.GetElaborationsView
	if err := cli.Get("v1/errorcode/elaborations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
