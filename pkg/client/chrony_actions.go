// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// UpdateChronyServers updates ChronyServers
func (cli *ZSClient) UpdateChronyServers(params param.UpdateChronyServersParam) (*view.UpdateChronyServersEventView, error) {
	resp := view.UpdateChronyServersEventView{}
	if err := cli.PutWithRespKey("v1/zops/chrony/actions", "", "", map[string]interface{}{
		"updateChronyServers": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetChronyServers gets ChronyServers by uuid
func (cli *ZSClient) GetChronyServers() (*view.GetChronyServersView, error) {
	var resp view.GetChronyServersView
	if err := cli.GetWithRespKey("v1/zops/chrony/servers", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncChronyServers operates on ChronyServers
func (cli *ZSClient) SyncChronyServers() (*view.SyncChronyServersEventView, error) {
	resp := view.SyncChronyServersEventView{}
	if err := cli.PutWithRespKey("v1/zops/chrony/actions", "", "", map[string]interface{}{
		"syncChronyServers": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
