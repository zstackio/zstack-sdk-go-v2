// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMiniStorageHostRef queries MiniStorageHostRef list
func (cli *ZSClient) QueryMiniStorageHostRef(params *param.QueryParam) ([]view.MiniStorageHostRefInventoryView, error) {
	var resp []view.MiniStorageHostRefInventoryView
	return resp, cli.List("v1/primary-storage/mini/host-refs", params, &resp)
}

func (cli *ZSClient) GetMiniStorageHostRef(uuid string) (*view.MiniStorageHostRefInventoryView, error) {
	var resp view.MiniStorageHostRefInventoryView
	if err := cli.Get("v1/primary-storage/mini/host-refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMiniStorageHostRef Pagination
func (cli *ZSClient) PageMiniStorageHostRef(params *param.QueryParam) ([]view.MiniStorageHostRefInventoryView, int, error) {
	var miniStorageHostRefs []view.MiniStorageHostRefInventoryView
	total, err := cli.Page("v1/primary-storage/mini/host-refs", params, &miniStorageHostRefs)
	return miniStorageHostRefs, total, err
}

// BootstrapMiniHost operates on MiniHost
func (cli *ZSClient) BootstrapMiniHost(params param.BootstrapMiniHostParam) (*view.BootstrapMiniHostEventView, error) {
	resp := view.BootstrapMiniHostEventView{}
	if err := cli.Post("v1/mini-clusters/hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateMiniHosts gets CandidateMiniHosts by uuid
func (cli *ZSClient) GetCandidateMiniHosts() (*view.GetCandidateMiniHostsView, error) {
	var resp view.GetCandidateMiniHostsView
	if err := cli.GetWithRespKey("v1/mini-clusters/candidate-hosts", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
