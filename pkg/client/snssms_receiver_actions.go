// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddSNSSmsReceiver adds SNSSmsReceiver
func (cli *ZSClient) AddSNSSmsReceiver(params param.AddSNSSmsReceiverParam) (*view.SNSSmsReceiverInventoryView, error) {
	resp := view.SNSSmsReceiverInventoryView{}
	if err := cli.Post("v1/sns/sms-endpoints/receivers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RemoveSNSSmsReceiver removes SNSSmsReceiver
func (cli *ZSClient) RemoveSNSSmsReceiver(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/sms-endpoints/{endpointUuid}/receivers/{phoneNumber}", uuid, string(deleteMode))
}
