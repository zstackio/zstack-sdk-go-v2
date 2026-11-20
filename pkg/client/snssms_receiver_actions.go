// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"

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
// RemoveSNSSmsReceiver removes one SMS receiver.
func (cli *ZSClient) RemoveSNSSmsReceiver(endpointUuid string, phoneNumber string, deleteMode param.DeleteMode) error {
	return cli.RemoveSNSSmsReceiverWithParams(endpointUuid, phoneNumber, deleteMode, param.RemoveSNSSmsReceiverParamDetail{})
}

// RemoveSNSSmsReceiverWithParams removes the supplied list of numbers or the
// single path number when the list is empty. The route requires a path number.
func (cli *ZSClient) RemoveSNSSmsReceiverWithParams(endpointUuid, phoneNumber string, deleteMode param.DeleteMode, params param.RemoveSNSSmsReceiverParamDetail) error {
	query := url.Values{}
	var required []string
	if len(params.PhoneNumberList) > 0 {
		query["phoneNumberList"] = params.PhoneNumberList
		required = append(required, "phoneNumberList")
		if phoneNumber == "" {
			phoneNumber = params.PhoneNumberList[0]
		}
	}
	number, err := actionPathSegment("phoneNumber", phoneNumber)
	if err != nil {
		return err
	}
	return cli.deleteActionWithQuery("v1/sns/sms-endpoints", endpointUuid, "receivers/"+number,
		deleteModeOrDefault(deleteMode, params.DeleteMode), query, required...)
}
