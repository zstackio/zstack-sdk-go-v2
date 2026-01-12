// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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
func (cli *ZSClient) RemoveSNSSmsReceiver(endpointUuid string, phoneNumber string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sns/sms-endpoints", endpointUuid, fmt.Sprintf("receivers/%s", phoneNumber), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
