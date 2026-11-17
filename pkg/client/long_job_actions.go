// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CleanLongJob operates on LongJob
func (cli *ZSClient) CleanLongJob(uuid string, params param.CleanLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey("v1/longjobs", uuid, "", map[string]interface{}{
		"cleanLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeLongJob operates on LongJob
func (cli *ZSClient) ResumeLongJob(uuid string, params param.ResumeLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey("v1/longjobs", uuid, "", map[string]interface{}{
		"resumeLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLongJob deletes LongJob
func (cli *ZSClient) DeleteLongJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/longjobs", uuid, string(deleteMode))
}
// UpdateLongJob updates LongJob
func (cli *ZSClient) UpdateLongJob(uuid string, params param.UpdateLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey("v1/longjobs", uuid, "", map[string]interface{}{
		"updateLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryLongJob queries LongJob list
func (cli *ZSClient) QueryLongJob(params *param.QueryParam) ([]view.LongJobInventoryView, error) {
	var resp []view.LongJobInventoryView
	return resp, cli.List("v1/longjobs", params, &resp)
}

func (cli *ZSClient) GetLongJob(uuid string) (*view.LongJobInventoryView, error) {
	var resp view.LongJobInventoryView
	if err := cli.Get("v1/longjobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLongJob Pagination
func (cli *ZSClient) PageLongJob(params *param.QueryParam) ([]view.LongJobInventoryView, int, error) {
	var longJobs []view.LongJobInventoryView
	total, err := cli.Page("v1/longjobs", params, &longJobs)
	return longJobs, total, err
}

// SuspendLongJob operates on SuspendLongJob
func (cli *ZSClient) SuspendLongJob(uuid string, params param.SuspendLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey("v1/longjobs", uuid, "", map[string]interface{}{
		"suspendLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelLongJob operates on CancelLongJob
func (cli *ZSClient) CancelLongJob(uuid string, params param.CancelLongJobParam) (*view.CancelLongJobEventView, error) {
	resp := view.CancelLongJobEventView{}
	if err := cli.PutWithRespKey("v1/longjobs", uuid, "", map[string]interface{}{
		"cancelLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunIAM2Script operates on RunIAM2Script
func (cli *ZSClient) RunIAM2Script(params param.RunIAM2ScriptParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post("v1/iam2/iam2-script/run", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunIAM2ScriptAsync Async
func (cli *ZSClient) RunIAM2ScriptAsync(params param.RunIAM2ScriptParam) (string, error) {

	resource := "v1/iam2/iam2-script/run"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// SubmitLongJob operates on SubmitLongJob
func (cli *ZSClient) SubmitLongJob(params param.SubmitLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post("v1/longjobs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RerunLongJob operates on RerunLongJob
func (cli *ZSClient) RerunLongJob(uuid string, params param.RerunLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey("v1/longjobs", uuid, "", map[string]interface{}{
		"rerunLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
