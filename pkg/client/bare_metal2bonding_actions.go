// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2Bonding queries BareMetal2Bonding list
func (cli *ZSClient) QueryBareMetal2Bonding(params *param.QueryParam) ([]view.BareMetal2BondingInventoryView, error) {
	var resp []view.BareMetal2BondingInventoryView
	return resp, cli.List("v1/baremetal2/bonding", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Bonding(uuid string) (*view.BareMetal2BondingInventoryView, error) {
	var resp view.BareMetal2BondingInventoryView
	if err := cli.Get("v1/baremetal2/bonding", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Bonding Pagination
func (cli *ZSClient) PageBareMetal2Bonding(params *param.QueryParam) ([]view.BareMetal2BondingInventoryView, int, error) {
	var bareMetal2Bondings []view.BareMetal2BondingInventoryView
	total, err := cli.Page("v1/baremetal2/bonding", params, &bareMetal2Bondings)
	return bareMetal2Bondings, total, err
}
// CreateBareMetal2Bonding creates BareMetal2Bonding
func (cli *ZSClient) CreateBareMetal2Bonding(params param.CreateBareMetal2BondingParam) (*view.BareMetal2BondingInventoryView, error) {
	resp := view.BareMetal2BondingInventoryView{}
	if err := cli.Post("v1/baremetal2/chassis/bond", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachProvisionNicToBonding operates on ProvisionNicToBonding
func (cli *ZSClient) AttachProvisionNicToBonding(uuid string, bondingUuid string, params param.AttachProvisionNicToBondingParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/baremetal2/bm-instances/%s/bm2-bondings/%s", uuid, bondingUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpBareMetal2Bonding operates on UpBareMetal2Bonding
func (cli *ZSClient) CleanUpBareMetal2Bonding(chassisUuid string, params param.CleanUpBareMetal2BondingParam) (*view.CleanUpBaremetal2BondingEventView, error) {
	resp := view.CleanUpBaremetal2BondingEventView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis", chassisUuid, "", map[string]interface{}{
		"cleanUpBareMetal2Bonding": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachProvisionNicFromBonding operates on ProvisionNicFromBonding
func (cli *ZSClient) DetachProvisionNicFromBonding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/bm-instances/bm2-bondings", uuid, string(deleteMode))
}
