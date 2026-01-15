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
	resp := view.CertificateInventoryView{}
	if err := cli.Put("v1/certificates", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateCertificate creates Certificate
func (cli *ZSClient) CreateCertificate(params param.CreateCertificateParam) (*view.CertificateInventoryView, error) {
	resp := view.CertificateInventoryView{}
	if err := cli.Post("v1/certificates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageCertificate Pagination
func (cli *ZSClient) PageCertificate(params *param.QueryParam) ([]view.CertificateInventoryView, int, error) {
	var certificates []view.CertificateInventoryView
	total, err := cli.Page("v1/certificates", params, &certificates)
	return certificates, total, err
}
