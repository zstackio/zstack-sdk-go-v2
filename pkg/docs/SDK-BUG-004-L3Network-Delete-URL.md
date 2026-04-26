# SDK-BUG-004 — `DeleteL3Network` URL / 参数异常导致删除失败

| 字段 | 值 |
|---|---|
| 严重度 | **High** |
| 状态 | ✅ 已修（2026-04-26，按当前代码实现复核） |
| 发现日期 | 2026-04-24（QA 跟踪扫描） |
| 影响 SDK 文件 | `pkg/client/l3network_actions.go` + `pkg/client/http_client.go` |
| 影响 provider 资源数 | 1（`zstack_l3network`） |
| 跟踪关键字 | `SDK-BUG-004`、`BUG-059` |

---

## 1. 复核结论（当前代码）

当前代码中：

```go
func (cli *ZSClient) DeleteL3Network(uuid string, deleteMode param.DeleteMode) error {
    return cli.Delete("v1/l3-networks", uuid, string(deleteMode))
}
```

底层 `ZSHttpClient.Delete` 会把 `resourceId` 继续传给 `DeleteWithSpec`，而 `getDeleteURL()` / `getURL()` 在 `resourceId != ""` 时会把 UUID 追加到资源路径中。因此当前实现会构造：

```text
DELETE /v1/l3-networks/{uuid}?deleteMode=Permissive
```

从当前代码来看，文档里描述的“UUID 丢失”问题**已不再成立**。

---

## 2. 历史现象

下游 provider 在 `terraform destroy` 时调用 `cli.DeleteL3Network(uuid, mode)`，被报告"URL 缺 UUID 参数"，DELETE 请求打到错误 endpoint，资源无法删除（必须人工介入或后端清理）。

---

## 3. SDK 当前实现

```go
// pkg/client/l3network_actions.go:51-54
// DeleteL3Network deletes L3Network
func (cli *ZSClient) DeleteL3Network(uuid string, deleteMode param.DeleteMode) error {
    return cli.Delete("v1/l3-networks", uuid, string(deleteMode))
}
```

底层 `ZSHttpClient.Delete` 拼接：

```go
// pkg/client/http_client.go:408
func (cli *ZSHttpClient) Delete(ctx context.Context, resource, resourceId, deleteMode string) error {
    return cli.DeleteWithSpec(ctx, resource, resourceId, "", fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
```

`getDeleteURL` 应该把 `resourceId` 拼到 URL 里：

```
DELETE /v1/l3-networks/{uuid}?deleteMode=Permissive
```

但 BUG-059 的 QA 报告记录 URL 中**缺了 UUID**，请求实际发送到：

```
DELETE /v1/l3-networks?deleteMode=Permissive
```

ZStack 拒绝（或当作"列表删除"）。

---

## 4. 历史复现记录（保留归档）

QA tracker 标注此 bug 为 "SDK 侧问题，非 provider"，但未提交完整 trace。建议 SDK 团队按以下步骤复现：

### 复现步骤

```go
import (
    "github.com/zstackio/zstack-sdk-go-v2/pkg/client"
    "github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

cfg := client.NewZSConfig("172.24.x.x", 8080, "zstack").
    LoginAccount("admin", "password")
cli := client.NewZSClient(cfg)
_, _ = cli.Login(context.Background())

// 先创建一个 L3Network
zoneUuid := "..."
l2Uuid := "..."
l3Resp, err := cli.CreateL3Network(param.CreateL3NetworkParam{
    Params: param.CreateL3NetworkParamDetail{
        Name:          "test-l3",
        L2NetworkUuid: l2Uuid,
        Type:          "L3BasicNetwork",
        IpVersion:     4,
    },
})

// 立即删除 — 观察实际 HTTP URL
err = cli.DeleteL3Network(l3Resp.UUID, param.DeleteModePermissive)
// 抓包/打开 SDK debug 模式查看实际请求 URL
```

### 调试技巧

打开 SDK debug：

```go
cfg.Debug(true)
```

观察日志中的：

```
[DEBUG] DELETE http://zstack:8080/v1/l3-networks/<EXPECT_UUID_HERE>?deleteMode=Permissive
```

如果 UUID 缺失，则是 `getDeleteURL` 在 `resourceId == ""`（或某种空字符串语义混淆）下拼错。

### 可能的根因方向

1. **getDeleteURL 内部 spec 处理 bug**：`Delete()` 传了 `spec=""`，`getDeleteURL` 可能把空 spec 当成 "no resource id" 处理；
2. **下游 provider 传错 uuid**：需要 SDK 团队反向查 provider 实际传入的值（已在 provider commit log 看似传入正确 uuid，但需 debug 模式确认）；
3. **DeleteMode 解析问题**：某些 ZStack 版本对 deleteMode 大小写敏感。

---

## 5. Provider 当前状态

未绕过。`zstack_l3network` 的 Delete 路径目前仍直接调用 SDK 方法：

```go
// resource_zstack_l3network.go (Delete 函数)
err := r.client.DeleteL3Network(state.Uuid.ValueString(), param.DeleteModePermissive)
```

如果历史分支上再次出现同类问题，**短期 provider 仍可走 ZSHttpClient 直接拼 URL 绕过**：

```go
err := r.client.ZSHttpClient.DeleteWithSpec(
    ctx,
    "v1/l3-networks",
    state.Uuid.ValueString(),
    "",
    fmt.Sprintf("deleteMode=%s", param.DeleteModePermissive),
    nil,
)
```

但按当前仓库代码复核，该问题应视为**已修复的历史问题**。

---

## 6. 建议下一步

1. 下游 provider 升级到当前 SDK 后，优先验证 `zstack_l3network` Delete 是否正常；
2. 若验证通过，可关闭该项历史追踪；
3. 若在特定分支/特定版本仍能复现，再补充 trace 并按版本单独记录，而不是继续标记当前主线 SDK 为待修。

---

## 7. 原始报告

- `terraform-provider-zstack/_bmad-output/bug-tracker.md` § BUG-059（标 P0 Open）
- `terraform-provider-zstack/_bmad-output/bug-tracker.md` § "SDK 修复跟进列表" → `SDK-FIX-004`
