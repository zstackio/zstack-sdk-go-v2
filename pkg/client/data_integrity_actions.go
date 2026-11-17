// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// AddIntegrityResource adds IntegrityResource
func (cli *ZSClient) AddIntegrityResource(params param.AddIntegrityResourceParam) (*view.AddIntegrityResourceEventView, error) {
	resp := view.AddIntegrityResourceEventView{}
	if err := cli.Post("v1/integrity/resource/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckBatchDataIntegrity operates on BatchDataIntegrity
func (cli *ZSClient) CheckBatchDataIntegrity() (*view.CheckBatchDataIntegrityView, error) {
	var resp view.CheckBatchDataIntegrityView
	if err := cli.GetWithRespKey("v1/check/batch/data/integrity/", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
