// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CheckElaborationContent operates on ElaborationContent
func (cli *ZSClient) CheckElaborationContent(params param.CheckElaborationContentParam) (*view.CheckElaborationContentView, error) {
	resp := view.CheckElaborationContentView{}
	if err := cli.Post("v1/errorcode/elaborations/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReloadElaboration operates on ReloadElaboration
func (cli *ZSClient) ReloadElaboration() (*view.ReloadElaborationEventView, error) {
	resp := view.ReloadElaborationEventView{}
	if err := cli.PutWithRespKey("v1/errorcode/actions", "", "", map[string]interface{}{
		"reloadElaboration": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborationCategories gets ElaborationCategories by uuid
func (cli *ZSClient) GetElaborationCategories() (*view.GetElaborationCategoriesView, error) {
	var resp view.GetElaborationCategoriesView
	if err := cli.GetWithRespKey("v1/errorcode/elaborations/categories", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetElaborations gets Elaborations by uuid
func (cli *ZSClient) GetElaborations() (*view.GetElaborationsView, error) {
	var resp view.GetElaborationsView
	if err := cli.GetWithRespKey("v1/errorcode/elaborations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
