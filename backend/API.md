# Volunteer-Team 后端 API 对接文档

面向前端的接口说明，包含路由、鉴权、请求/响应结构，以及 JSON 序列化约定。请特别注意：返回中的所有 `int64` 类型 ID 字段统一以字符串序列化，避免 JS 精度问题。

## 概览

- 基础路径：`/api`
- 静态资源：`/uploads`（用于访问已上传文件，如用户头像）
- 统一响应包：
  - `code`：整数，`0` 表示成功，`-1` 表示业务失败
  - `message`：字符串，错误描述或 `"OK"`
  - `data`：业务数据对象或 `null`
- 服务地址：`http://{host}:{port}`（来自后端配置，实际以部署为准）

## 鉴权

- 携带 `Authorization` 请求头：`Bearer <token>`
- 角色等级：`COMMON_USER` < `INTERNAL_USER` < `ADMIN`
- 鉴权失败：
  - 权限不足：`200` 且统一包返回 `code = -1, message = "权限不足"`
  - token 为空/过期/非法：返回 `401`，响应体为纯字符串（如 `"token已过期"`），非统一 JSON 包。前端需兼容该场景。

## JSON 序列化约定

- 返回中的 `int64` 类型 ID 字段一律以字符串序列化（`json:",string"`）：
  - 示例：`order_id`、`summary_id`
- 请求中的 ID 字段（如 `order_id`、`summary_id`）统一使用字符串传递
- 时间字段 `ctime`、`utime` 为整型时间戳（`int64`），前端按整数处理即可（常见为秒；如需严格单位请以实际数据为准）
- 统一响应结构固定为 `code/message/data`，除 `401` 场景

## 统一响应示例

- 成功：

```json
{
  "code": 0,
  "message": "OK",
  "data": { "..." : "..." }
}
```

- 失败（业务错误）：

```json
{
  "code": -1,
  "message": "错误信息",
  "data": null
}
```

- 鉴权错误示例（401）：响应体为纯字符串

## 静态资源访问

- 上传接口会返回文件名（如 `md5.ext`）
- 前端展示 URL 模板：`{scheme}://{host}:{port}/uploads/{fileName}`

---

## 路由与接口

### 通用

- 路由组：`/api/common`

#### GET `/api/common/ping`
- 鉴权：否
- 请求：无
- 响应 `data`：
  - `"pong"`

---

### 用户相关

- 路由组：`/api/user`

#### POST `/api/user/login`
- 鉴权：否
- Content-Type：`application/json`
- 请求体：
  - `account`: string（账号密码登录时必填）
  - `password`: string（账号密码登录时必填）
  - `email`: string（邮箱验证码登录时必填）
  - `code`: string（邮箱验证码登录时必填）
  - `login_type`: string，枚举：`"account"` | `"email"`
- 响应 `data`：
  - `token`: string（JWT）
  - `username`: string
  - `avatar`: string（文件名；拼接 `/uploads/<avatar>` 访问）
- 失败 `message` 可能值：
  - `"账号或密码错误"`、`"验证码错误"`、`"邮箱错误"`、`"暂不支持这种登录方式"`、`"jwt默认错误"`

#### POST `/api/user/register`
- 鉴权：否
- Content-Type：`application/json`
- 请求体：
  - `account`: string
  - `password`: string
  - `username`: string
  - `email`: string
  - `code`: string（邮箱验证码）
  - `is_internal`: bool（是否内部人员；当前未用于角色分配）
- 响应 `data`：
  - `message`: string（如 `"用户注册成功！"`）
- 失败 `message` 可能值：
  - `"验证码错误"`、`"邮箱已经被使用"`、`"账号已经被使用"`、`"默认错误"`

#### POST `/api/user/code/login`
- 用途：获取登录验证码
- 鉴权：否
- Content-Type：`application/json`
- 请求体：
  - `email`: string
- 响应 `data`：
  - `code`: string

#### POST `/api/user/code/register`
- 用途：获取注册验证码
- 鉴权：否
- Content-Type：`application/json`
- 请求体：
  - `email`: string
- 响应 `data`：
  - `code`: string

#### POST `/api/user/code/reset`
- 用途：获取重置密码验证码
- 鉴权：否
- Content-Type：`application/json`
- 请求体：
  - `email`: string
- 响应 `data`：
  - `code`: string

#### POST `/api/user/resetPassword`
- 鉴权：否
- Content-Type：`application/json`
- 请求体：
  - `email`: string
  - `code`: string
  - `new_password`: string
- 响应 `data`：
  - `message`: string（如 `"重置密码成功！"`）
- 失败 `message` 可能值：
  - `"验证码错误"`、`"用户不存在"`、`"默认错误"`

#### POST `/api/user/updateAvatar`
- 鉴权：是（`COMMON_USER` 及以上）
- Content-Type：`multipart/form-data`
- form-data：
  - `file`: 文件（图片），大小限制约 `10MB`，后端校验后缀
- 响应 `data`：
  - 字符串：文件名（如 `md5.ext`）
- 失败 `message` 可能值：
  - `"文件大小不能超过10MB"`、`"文件读取失败"`、`"用户不存在"`、`"默认错误"`

---

### 订单相关

- 路由组：`/api/order`

#### POST `/api/order/create`
- 鉴权：是（`COMMON_USER` 及以上）
- Content-Type：`application/json`
- 请求体：
  - `username`: string
  - `student_id`: string
  - `campus_location`: string
  - `department`: string
  - `phone_number`: string
  - `wechat_id`: string
  - `address`: string
  - `device_model`: string
  - `os_version`: string
  - `problem_description`: string
  - `notes`: string
- 响应 `data`：
  - `"创建订单成功"`

#### GET `/api/order/list`
- 鉴权：是（`COMMON_USER` 及以上）
- Query：无
- 响应 `data`：
  - `orders`: 数组
    - `order_id`: string
    - `ctime`: int64
    - `problem_description`: string
- 权限差异：
  - 普通用户：仅返回本人未处理的订单
  - 内部人员/管理员：返回全部未处理订单

#### GET `/api/order/detail`
- 鉴权：是（`COMMON_USER` 及以上）
- Content-Type：`application/x-www-form-urlencoded`（Query 绑定）
- Query：
  - `order_id`: string
- 响应 `data`：
  - `username`: string
  - `student_id`: string
  - `campus_location`: string
  - `department`: string
  - `phone_number`: string
  - `wechat_id`: string
  - `address`: string
  - `device_model`: string
  - `os_version`: string
  - `problem_description`: string
  - `notes`: string
  - `ctime`: int64
- 校验与限制：
  - 非内部人员仅能查看自己的订单，否则 `message = "禁止操作别人的订单"`

#### PUT `/api/order/finish`
- 鉴权：是（`INTERNAL_USER` 及以上）
- Content-Type：`application/x-www-form-urlencoded`（Query 绑定）
- Query：
  - `order_id`: string
- 响应 `data`：
  - `"订单已完成"`
- 失败 `message` 可能值：
  - `"订单不存在"`、`"参数格式错误"`、`"默认错误"`

---

### 修机总结相关

- 路由组：`/api/summary`

#### POST `/api/summary/create`
- 鉴权：是（`INTERNAL_USER` 及以上）
- Content-Type：`application/json`
- 请求体（必填）：
  - `order_id`: string
  - `problem_type`: string
  - `problem_description`: string
  - `repair_summary`: string
  - `receiver_name`: string
- 响应 `data`：
  - `"创建修机总结成功"`
- 备注：
  - 创建总结后，后端会异步将对应订单状态置为完成

#### GET `/api/summary/list`
- 鉴权：是（`INTERNAL_USER` 及以上）
- Query：无
- 响应 `data`：
  - `summaries`: 数组
    - `summary_id`: string
    - `order_id`: string
    - `problem_type`: string
    - `problem_description`: string
    - `repair_summary`: string
    - `receiver_name`: string
    - `utime`: int64

#### GET `/api/summary/detail`
- 鉴权：是（`INTERNAL_USER` 及以上）
- Content-Type：`application/x-www-form-urlencoded`（Query 绑定）
- Query：
  - `summary_id`: string（必填）
- 响应 `data`：
  - `summary_id`: string
  - `order_id`: string
  - `problem_type`: string
  - `problem_description`: string
  - `repair_summary`: string
  - `receiver_name`: string
  - `utime`: int64

#### PUT `/api/summary/update`
- 鉴权：是（`INTERNAL_USER` 及以上）
- Content-Type：`application/json`
- 请求体：
  - `summary_id`: string（必填）
  - `problem_type`: string（可选）
  - `problem_description`: string（可选）
  - `repair_summary`: string（可选）
  - `receiver_name`: string（可选）
- 响应 `data`：
  - `"更新修机总结成功"`

---

## 常见错误文案（示例）

- 参数类：`"参数格式错误"`
- 通用：`"默认错误"`
- 登录：`"账号或密码错误"`、`"邮箱错误"`、`"验证码错误"`、`"暂不支持这种登录方式"`
- 用户：`"邮箱已经被使用"`、`"账号已经被使用"`、`"用户不存在"`
- 订单：`"订单不存在"`、`"禁止操作别人的订单"`
- 总结：`"修机总结不存在"`
- 文件：`"文件大小不能超过10MB"`、`"文件读取失败"`

> 注：以上错误文案以后端实际返回为准；前端统一读取 `message` 字段展示。