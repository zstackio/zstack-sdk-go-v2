// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteTag deletes Tag
func (cli *ZSClient) DeleteTag(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags", uuid, string(deleteMode))
}

// DetachTagFromResources operates on TagFromResources
func (cli *ZSClient) DetachTagFromResources(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/tags", uuid, string(deleteMode))
}

// AttachTagToResources operates on TagToResources
func (cli *ZSClient) AttachTagToResources(tagUuid string, params param.AttachTagToResourcesParam) (*view.AttachTagToResourcesEventView, error) {
	resp := view.AttachTagToResourcesEventView{}
	if _, err := cli.PostWithAsync(fmt.Sprintf("v1/tags/%s/resources", tagUuid), "", params, &resp, false); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryTag queries Tag list
func (cli *ZSClient) QueryTag(params *param.QueryParam) ([]view.TagPatternInventoryView, error) {
	var resp []view.TagPatternInventoryView
	return resp, cli.List("v1/tags", params, &resp)
}

func (cli *ZSClient) GetTag(uuid string) (*view.TagPatternInventoryView, error) {
	var resp view.TagPatternInventoryView
	if err := cli.Get("v1/tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTag Pagination
func (cli *ZSClient) PageTag(params *param.QueryParam) ([]view.TagPatternInventoryView, int, error) {
	var tags []view.TagPatternInventoryView
	total, err := cli.Page("v1/tags", params, &tags)
	return tags, total, err
}

// CreateTag creates Tag
func (cli *ZSClient) CreateTag(params param.CreateTagParam) (*view.TagPatternInventoryView, error) {
	resp := view.TagPatternInventoryView{}
	if err := cli.Post("v1/tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTag updates Tag
func (cli *ZSClient) UpdateTag(uuid string, params param.UpdateTagParam) (*view.TagPatternInventoryView, error) {
	resp := view.TagPatternInventoryView{}
	if err := cli.PutWithRespKey("v1/tags", uuid, "", map[string]interface{}{
		"updateTag": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
