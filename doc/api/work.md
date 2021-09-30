# 企业微信

host: https://qyapi.weixin.qq.com/

## 微信客服

[官方文档](https://work.weixin.qq.com/api/doc/90000/90135/94638)

### 客服账号管理

|       名称       | 请求方式 | URL                         | 是否已实现 | 使用方法                  |
| :--------------: | -------- | :-------------------------- | ---------- | ------------------------- |
|   添加客服帐号   | POST     | /cgi-bin/kf/account/add     | YES        | (r *Client) AccountAdd    |
|   删除客服帐号   | POST     | /cgi-bin/kf/account/del     | YES        | (r *Client) AccountDel    |
|   修改客服帐号   | POST     | /cgi-bin/kf/account/update  | YES        | (r *Client) AccountUpdate |
| 获取客服帐号列表 | GET      | /cgi-bin/kf/account/list    | YES        | (r *Client) AccountList() |
| 获取客服帐号链接 | GET      | /cgi-bin/kf/add_contact_way | YES        | (r *Client) AddContactWay |

### 接待人员列表

TODO

### 会话分配与消息收发

TODO

### 其他基础信息获取

TODO

## 应用管理

TODO

