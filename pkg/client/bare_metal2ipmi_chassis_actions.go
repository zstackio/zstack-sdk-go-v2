// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddBareMetal2IpmiChassis adds BareMetal2IpmiChassis
func (cli *ZSClient) AddBareMetal2IpmiChassis(params param.AddBareMetal2IpmiChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2IpmiChassis updates BareMetal2IpmiChassis
func (cli *ZSClient) UpdateBareMetal2IpmiChassis(uuid string, params param.UpdateBareMetal2IpmiChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis/ipmi", uuid, "", map[string]interface{}{
		"updateBareMetal2IpmiChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBareMetal2IpmiChassisConfigFile operates on BareMetal2IpmiChassisConfigFile
func (cli *ZSClient) CheckBareMetal2IpmiChassisConfigFile() (*view.CheckBareMetal2ChassisConfigFileView, error) {
	resp := view.CheckBareMetal2ChassisConfigFileView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file/check", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchAddBareMetal2IpmiChassis operates on AddBareMetal2IpmiChassis
func (cli *ZSClient) BatchAddBareMetal2IpmiChassis() (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/from-file", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBareMetal2IpmiChassisHardwareInfo creates BareMetal2IpmiChassisHardwareInfo
func (cli *ZSClient) CreateBareMetal2IpmiChassisHardwareInfo(params param.CreateBareMetal2IpmiChassisHardwareInfoParam) (*view.CreateBareMetal2ChassisHardwareView, error) {
	resp := view.CreateBareMetal2ChassisHardwareView{}
	if err := cli.Post("v1/baremetal2/chassis/ipmi/hardwareinfos", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
