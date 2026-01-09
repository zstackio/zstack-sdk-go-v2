// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateCertificate updates Certificate
func (cli *ZSClient) UpdateCertificate(uuid string, params param.UpdateCertificateParam) (*view.CertificateInventoryView, error) {
	var resp view.UpdateCertificateEventView
	if err := cli.Put("v1/certificates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateCertificate creates Certificate
func (cli *ZSClient) CreateCertificate(params param.CreateCertificateParam) (*view.CertificateInventoryView, error) {
	var resp view.CreateCertificateEventView
	if err := cli.Post("v1/certificates", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteCertificate deletes Certificate
func (cli *ZSClient) DeleteCertificate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/certificates", uuid, string(deleteMode))
}
// QueryCertificate queries Certificate list
func (cli *ZSClient) QueryCertificate(params *param.QueryParam) ([]view.CertificateInventoryView, error) {
	var resp []view.CertificateInventoryView
	return resp, cli.List("v1/certificates", params, &resp)
}

func (cli *ZSClient) GetCertificate(uuid string) (*view.CertificateInventoryView, error) {
	var resp view.CertificateInventoryView
	if err := cli.Get("v1/certificates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
