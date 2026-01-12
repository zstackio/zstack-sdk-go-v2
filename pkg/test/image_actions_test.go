// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImage error: %v", err)
		return
	}
	golog.Infof("QueryImage result count: %d", len(result))
}

func TestGetImage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetImage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetImage error: %v", err)
		return
	}
	golog.Infof("GetImage result: %s", result.UUID)
}

func TestUpdateImage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateImageParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateImageParamDetail{
			Name: "centos-test",
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateImage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateImage error: %v", err)
		return
	}
	golog.Infof("UpdateImage result: %s", result.UUID)
}

func TestDeleteImage(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	//t.Skip("TestDeleteImage is skipped by default to prevent accidental deletion. Remove this line to enable.")

	err := accountLoginCli.DeleteImage("a2fe8439606847aeb0c0fde516e43654", param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteImage error: %v", err)
		return
	}
	golog.Infof("DeleteImage succeeded for UUID: %s", "a2fe8439606847aeb0c0fde516e43654")
}

func TestAddImage(t *testing.T) {
	// Add operation - similar to Create
	storage, err := accountLoginCli.QueryBackupStorage(&param.QueryParam{})
	if err != nil {
		t.Errorf("QueryBackupStorage error: %v", err)
		return
	}
	if len(storage) == 0 {
		t.Skip("No BackupStorage found to add image")
		return
	}

	imageParam := param.AddImageParam{
		BaseParam: param.BaseParam{
			SystemTags: []string{"bootMode::Legacy"},
		},
		Params: param.AddImageParamDetail{
			Name:               "CentOS-6.8-i386-LiveCD",
			Description:        "接口image",
			Url:                "http://172.20.15.213/rds/V3.14.1-p2/zstack-rds-3.14.1-p2_x86.qcow2",
			MediaType:          "RootVolumeTemplate",
			GuestOsType:        "Linux",
			System:             false,
			Format:             "qcow2",
			Platform:           "Linux",
			BackupStorageUuids: []string{storage[0].UUID},
			Type:               "",
			ResourceUuid:       "",
			Architecture:       "x86_64",
			Virtio:             false,
		},
	}

	result, err := accountLoginCli.AddImage(imageParam)
	if err != nil {
		t.Errorf("AddImage error: %v", err)
		return
	}

	golog.Infof("AddImage succeeded, UUID: %s, Name: %s", result.UUID, result.Name)
	//创建失败情况
	/*
		imageParam = param.AddImageParam{
			BaseParam: param.BaseParam{
				SystemTags: []string{"bootMode::Legacy"},
			},
			Params: param.AddImageDetailParam{
				Name:               "image-4",
				Description:        "接口image",
				Url:                "https://image.baidu.com/search/detail?tn=baiduimagedetail&word=%E6%B8%90%E5%8F%98%E9%A3%8E%E6%A0%BC%E6%8F%92%E7%94%BB&album_tab=%E8%AE%BE%E8%AE%A1%E7%B4%A0%E6%9D%90&album_id=409&ie=utf-8&fr=albumsdetail&cs=4036010509,3445021118&pi=144521&pn=1&ic=0&objurl=https%3A%2F%2Ft7.baidu.com%2Fit%2Fu%3D4036010509%2C3445021118%26fm%3D193%26f%3DGIF",
				MediaType:          param.RootVolumeTemplate,
				GuestOsType:        "Windows 10",
				System:             false,
				Format:             param.Qcow2,
				Platform:           "Windows",
				BackupStorageUuids: []string{"26684790e4734a0bbb506f40907f57da"},
				Type:               "",
				ResourceUuid:       "",
				Architecture:       "x86_64",
				Virtio:             false,
			},
		}

		_, err = accountLoginCli.AddImage(imageParam)
		if err != nil {
			golog.Errorf("ZSClient.CreateImage error:%v", err)
		}

		golog.Infof("======================================")
		//删除
		err = accountLoginCli.DeleteImage(r.UUID, param.DeleteModeEnforcing)
		if err != nil {
			golog.Errorf("ZSClient.DeleteImage error:%v", err)
		}

		//彻底删除
		err = accountLoginCli.ExpungeImage(r.UUID)
		if err != nil {
			golog.Errorf("ZSClient.ExpungeImage error:%v", err)
		}*/

}

func TestAddImageAsync(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddImageAsync requires valid creation parameters")

}

func TestSyncImage(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncImage requires a valid resource to sync")

}

func TestRecoverImage(t *testing.T) {
	// Recover operation - requires a deleted resource
	t.Skip("TestRecoverImage requires a deleted resource UUID")

}

func TestCloneImage(t *testing.T) {
	// Clone operation
	t.Skip("TestCloneImage requires a valid resource to clone")

}

func TestExpungeImage(t *testing.T) {
	// Expunge operation - permanently deletes
	t.Skip("TestExpungeImage is dangerous - permanently deletes resource")

}
