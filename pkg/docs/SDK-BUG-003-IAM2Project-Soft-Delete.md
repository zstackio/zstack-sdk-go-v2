# SDK-BUG-003 — `DeleteIAM2Project` 是软删除，不释放 name

| 字段 | 值 |
|---|---|
| 严重度 | Medium |
| 状态 | ✅ 已修（2026-04-26，本仓库当前工作树） |
| 发现日期 | 2026-04-24 |
| 影响 SDK 文件 | `pkg/client/iam2project_actions.go` |
| 影响 provider 资源数 | 1（`zstack_iam2_project`） |
| 跟踪关键字 | `SDK-BUG-003`、`SDK-WA-003`、`ExpungeIAM2Project` |

---

## 1. 现象

调用 `cli.DeleteIAM2Project(uuid, mode)` 后，project 进入回收站（soft-deleted），但**该 project name 仍被占用**。

下游再尝试 `CreateIAM2Project` 创建同名 project 时报错：

```
DUPLICATE_NAME: IAM2Project 'acc-test-project' already exists.
```

实际 project 已经被 Terraform "destroy" 了，但 ZStack 后台还有残留。

---

## 2. 根因

SDK 当前两个方法分离：

```go
// pkg/client/iam2project_actions.go:21
func (cli *ZSClient) DeleteIAM2Project(uuid string, deleteMode param.DeleteMode) error {
    return cli.Delete("v1/iam2/projects", uuid, string(deleteMode))
}

// pkg/client/iam2project_actions.go:56
func (cli *ZSClient) ExpungeIAM2Project(uuid string) error {
    params := map[string]interface{}{
        "expungeIAM2Project": map[string]interface{}{},
    }
    return cli.Put("v1/iam2/projects", uuid, params, nil)
}
```

`DeleteIAM2Project` 走 ZStack 的"软删除→回收站"路径；`ExpungeIAM2Project` 是回收站清空动作。**两步必须连续调用才能释放 name**。

ZStack 其他可回收资源（VM 实例、卷、镜像）的 `DeleteXxx` 也是软删，但只有 IAM2Project 的"name 唯一性约束作用于回收站资源"这一行为让用户感受最强烈。

---

## 3. 复现

```go
cli := client.NewZSClient(cfg)

// 第一次创建 + 删除
proj, _ := cli.CreateIAM2Project(param.CreateIAM2ProjectParam{Params: ...{Name: "test"}})
_ = cli.DeleteIAM2Project(proj.UUID, param.DeleteModePermissive)

// 立即重建同名 project — 失败
_, err := cli.CreateIAM2Project(param.CreateIAM2ProjectParam{Params: ...{Name: "test"}})
// err: DUPLICATE_NAME
```

加上 expunge 就能成功：

```go
_ = cli.DeleteIAM2Project(proj.UUID, param.DeleteModePermissive)
_ = cli.ExpungeIAM2Project(proj.UUID)
_, err := cli.CreateIAM2Project(...)  // OK
```

---

## 4. 修复方案

### 方案 A — 给 `DeleteIAM2Project` 加 `purge bool` 参数（推荐）

```go
func (cli *ZSClient) DeleteIAM2Project(uuid string, deleteMode param.DeleteMode, purge bool) error {
    if err := cli.Delete("v1/iam2/projects", uuid, string(deleteMode)); err != nil {
        return err
    }
    if purge {
        return cli.ExpungeIAM2Project(uuid)
    }
    return nil
}
```

**注意**：这是 break change，需要 SDK major / minor bump，并提示下游迁移：

```diff
- err := cli.DeleteIAM2Project(uuid, param.DeleteModePermissive)
+ err := cli.DeleteIAM2Project(uuid, param.DeleteModePermissive, true)
```

### 方案 B — 新增 `DeleteAndExpungeIAM2Project` 复合方法（已采用，无 break）

```go
func (cli *ZSClient) DeleteAndExpungeIAM2Project(uuid string, deleteMode param.DeleteMode) error {
    if err := cli.DeleteIAM2Project(uuid, deleteMode); err != nil {
        return err
    }
    return cli.ExpungeIAM2Project(uuid)
}
```

当前仓库已在 `pkg/client/iam2project_actions.go` 新增该方法，下游可直接切换，无 break。

### 方案 C — 不动 SDK，永远由调用方手动两步

provider 已经这么做了。但其他 SDK 用户会重复踩坑。

---

## 5. Provider 当前 workaround

`terraform-provider-zstack/zstack/provider/resource_zstack_iam2_project.go:226-249`：

```go
func (r *iam2ProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var state iam2ProjectModel
    ...
    if err := r.client.DeleteIAM2Project(state.Uuid.ValueString(), param.DeleteModePermissive); err != nil {
        resp.Diagnostics.AddError("Failed to delete IAM2 project", err.Error())
        return
    }
    // SDK-BUG-003: DeleteIAM2Project soft-deletes; expunge to release the name.
    if err := r.client.ExpungeIAM2Project(state.Uuid.ValueString()); err != nil {
        resp.Diagnostics.AddWarning(
            "IAM2 project soft-deleted but expunge failed",
            "The project will remain in the recycle bin and may block re-creation by name. "+err.Error(),
        )
    }
}
```

> 注意：这是 SDK 修复前的 provider workaround。升级到包含 `DeleteAndExpungeIAM2Project` 的 SDK 后，可收敛为单次 SDK 调用。

---

## 6. SDK 修复后的 Provider 回收清单

### 当前状态

- SDK 侧已新增 `DeleteAndExpungeIAM2Project(uuid, deleteMode)`。
- 本仓库已补回归测试：`pkg/client/client_behavior_test.go`
- 下游 provider 可以把“Delete + Expunge 两步调用”替换成新的复合方法。

### 如果 SDK 走方案 A（加 `purge` 参数）

```diff
- if err := r.client.DeleteIAM2Project(state.Uuid.ValueString(), param.DeleteModePermissive); err != nil {
-     resp.Diagnostics.AddError(...)
-     return
- }
- if err := r.client.ExpungeIAM2Project(state.Uuid.ValueString()); err != nil {
-     resp.Diagnostics.AddWarning(...)
- }
+ if err := r.client.DeleteIAM2Project(state.Uuid.ValueString(), param.DeleteModePermissive, true); err != nil {
+     resp.Diagnostics.AddError(...)
+     return
+ }
```

### 如果 SDK 走方案 B（新方法）

```diff
- if err := r.client.DeleteIAM2Project(...); err != nil { ... }
- if err := r.client.ExpungeIAM2Project(...); err != nil { ... }
+ if err := r.client.DeleteAndExpungeIAM2Project(state.Uuid.ValueString(), param.DeleteModePermissive); err != nil {
+     resp.Diagnostics.AddError(...)
+     return
+ }
```

无论方案 A 或 B，都要：

```bash
grep -rn "ExpungeIAM2Project" terraform-provider-zstack/zstack/provider/
# 删除每处的两步调用 + // SDK-BUG-003 注释
```

---

## 7. 原始报告

- `terraform-provider-zstack/_bmad-output/bug-tracker.md` § BUG-053（功能性 critical），§ SDK-WA-003（workaround 登记）
