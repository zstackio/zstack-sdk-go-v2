// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddCCSCertificate adds CCSCertificate
func (cli *ZSClient) AddCCSCertificate(params param.AddCCSCertificateParam) (*view.CCSCertificateInventoryView, error) {
	var resp view.AddCCSCertificateEventView
	if err := cli.Post("v1/crypto/ccs-certificate/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteCCSCertificate deletes CCSCertificate
func (cli *ZSClient) DeleteCCSCertificate(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/crypto/ccs-certificate/delete", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryCCSCertificate queries CCSCertificate list
func (cli *ZSClient) QueryCCSCertificate(params *param.QueryParam) ([]view.CCSCertificateInventoryView, error) {
	var resp []view.CCSCertificateInventoryView
	return resp, cli.List("v1/crypto/ccs-certificate/certificates/", params, &resp)
}

func (cli *ZSClient) GetCCSCertificate(uuid string) (*view.CCSCertificateInventoryView, error) {
	var resp view.CCSCertificateInventoryView
	if err := cli.Get("v1/crypto/ccs-certificate/certificates/", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
