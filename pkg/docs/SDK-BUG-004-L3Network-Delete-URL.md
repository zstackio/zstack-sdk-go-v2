# SDK-BUG-004 — `DeleteL3Network` URL / 参数异常导致删除失败

| 字段 | 值 |
|---|---|
| 严重度 | **High** |
| 状态 | 🔲 待复现 / 待修 |
| 发现日期 | 2026-04-24（QA 跟踪扫描） |
| 影响 SDK 文件 | `pkg/client/l3network_actions.go` 第 51-54 行 |
| 影响 provider 资源数 | 1（`zstack_l3network`） |
| 跟踪关键字 | `SDK-BUG-004`、`BUG-059` |

---

## 1. 现象

下游 provider 在 `terraform destroy` 时调用 `cli.DeleteL3Network(uuid, mode)`，被报告"URL 缺 UUID 参数"，DELETE 请求打到错误 endpoint，资源无法删除（必须人工介入或后端清理）。

---

## 2. SDK 当前实现

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

## 3. 待 SDK 团队复现的开放问题

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

## 4. Provider 当前状态

未绕过。`zstack_l3network` 的 Delete 路径目前仍直接调用 SDK 方法：

```go
// resource_zstack_l3network.go (Delete 函数)
err := r.client.DeleteL3Network(state.Uuid.ValueString(), param.DeleteModePermissive)
```

如果 SDK 团队复现确认是 SDK bug，**短期 provider 可以走 ZSHttpClient 直接拼 URL 绕过**：

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

但在 SDK 复现/修复之前，建议先按"复现步骤"section 拿到准确 trace 再决定方案。

---

## 5. 建议下一步

1. **SDK 团队**：按本文 § 3 复现，捕获完整 HTTP trace；
2. 若复现成功 → 修 `getDeleteURL` 或 `DeleteWithSpec` 在 `resourceId != ""` 时强制保留 uuid；
3. 若复现失败 → 重新审视 provider 侧调用（也许 provider 在某分支误传 `state.Uuid` 为空）；
4. 修复后给下游 provider 升 SDK 即可，无 workaround 需回收。

---

## 6. 原始报告

- `terraform-provider-zstack/_bmad-output/bug-tracker.md` § BUG-059（标 P0 Open）
- `terraform-provider-zstack/_bmad-output/bug-tracker.md` § "SDK 修复跟进列表" → `SDK-FIX-004`
