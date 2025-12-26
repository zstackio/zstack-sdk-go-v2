// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReloadElaboration operates on ReloadElaboration
func (cli *ZSClient) ReloadElaboration(uuid string, params param.ReloadElaborationParam) (*view.ReloadElaborationEventView, error) {
	resp := view.ReloadElaborationEventView{}
	if err := cli.Put("v1/errorcode/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
