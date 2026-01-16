// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accessKeyAuthCli.QueryBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBackupStorage error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryBackupStorage result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s\t%d", r.UUID, r.Name, r.State, r.Type, r.TotalCapacity)
	}
	golog.Infof("======================================")
}

func TestQueryBackupStorage2(t *testing.T) {
	// Query with conditions - Connected backup storages
	params := param.NewQueryParam()
	params.AddQ("state=Enabled")
	params.AddQ("status=Connected")
	params.Start(0).Limit(10).ReplyWithCount(true)
	result, err := accessKeyAuthCli.QueryBackupStorage(&params)
	if err != nil {
		t.Errorf("TestQueryBackupStorage2 error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("Found %d Enabled/Connected BackupStorages:", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%d GB total\t%d GB available", r.UUID, r.Name, r.Type, r.TotalCapacity/1024/1024/1024, r.AvailableCapacity/1024/1024/1024)
	}
	golog.Infof("======================================")
}

func TestPageBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestPageBackupStorage error: %v", err)
		return
	}
	golog.Infof("PageBackupStorage result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s\t%s", r.UUID, r.Name, r.State, r.Type)
	}
}

func TestGetBackupStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BackupStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accessKeyAuthCli.GetBackupStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBackupStorage error: %v", err)
		return
	}
	golog.Infof("GetBackupStorage result: %s, Name: %s, Type: %s", result.UUID, result.Name, result.Type)
}
