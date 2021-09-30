# 企业微信

host: https://qyapi.weixin.qq.com/

## 微信客服

[官方文档](https://work.weixin.qq.com/api/doc/90000/90135/94638)

### 客服账号管理

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94684)

|       名称       | 请求方式 | URL                         | 是否已实现 | 使用方法                             |贡献者       |
| :--------------: | -------- | :-------------------------- | ---------- | -------------------------------|------------|
|   添加客服帐号   | POST     | /cgi-bin/kf/account/add     | YES        | (r *Client) AccountAdd           | NICEXAI    |
|   删除客服帐号   | POST     | /cgi-bin/kf/account/del     | YES        | (r *Client) AccountDel           | NICEXAI    |
|   修改客服帐号   | POST     | /cgi-bin/kf/account/update  | YES        | (r *Client) AccountUpdate        | NICEXAI    |
| 获取客服帐号列表 | GET      | /cgi-bin/kf/account/list     | YES        | (r *Client) AccountList          | NICEXAI    |
| 获取客服帐号链接 | GET      | /cgi-bin/kf/add_contact_way  | YES        | (r *Client) AddContactWay        | NICEXAI    |

### 接待人员列表

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94693)

|       名称       | 请求方式 | URL                         | 是否已实现 | 使用方法                             |贡献者       |
| :--------------: | -------- | :-------------------------- | ---------- | -------------------------------|------------|
|   添加接待人员     | POST     | /cgi-bin/kf/servicer/add    | YES        | (r *Client) ReceptionistAdd     | NICEXAI    |
|   删除接待人员     | POST     | /cgi-bin/kf/servicer/del    | YES        | (r *Client) ReceptionistDel     | NICEXAI    |
| 获取接待人员列表    | GET      | /cgi-bin/kf/servicer/list   | YES        | (r *Client) ReceptionistList    | NICEXAI    |

### 会话分配与消息收发

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94694)

|       名称       | 请求方式 | URL                               | 是否已实现   | 使用方法                          |贡献者       |
| :--------------: | -------- | :-------------------------------| ---------- | ------------------------------- |------------|
|   获取会话状态     | POST     | /cgi-bin/kf/service_state/get   | YES        | (r *Client) ServiceStateGet     | NICEXAI    |
|   变更会话状态     | POST     | /cgi-bin/kf/service_state/trans | YES        | (r *Client) ServiceStateTrans   | NICEXAI    |
|   读取消息        | POST     | /cgi-bin/kf/sync_msg            | YES        | (r *Client) SyncMsg             | NICEXAI    |
|   发送消息        | POST     | /cgi-bin/kf/send_msg            | YES        | (r *Client) SendMsg             | NICEXAI    |
|   发送事件响应消息 | POST     | /cgi-bin/kf/send_msg_on_event   | YES        | (r *Client) SendMsgOnEvent      | NICEXAI    |

###「升级服务」配置

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/94702)

|       名称                | 请求方式  | URL                                               | 是否已实现 | 使用方法                            |贡献者       |
| :--------------:         | -------- | :-------------------------------------------------| ---------- | -------------------------------  |------------|
| 获取配置的专员与客户群       | POST     | /cgi-bin/kf/customer/get_upgrade_service_config   | YES        | (r *Client) UpgradeServiceConfig | NICEXAI    |
| 为客户升级为专员或客户群服务  | POST     | /cgi-bin/kf/customer/upgrade_service              | YES        | (r *Client) UpgradeService       | NICEXAI    |
| 为客户取消推荐             | POST     | /cgi-bin/kf/customer/cancel_upgrade_service       | YES        | (r *Client) UpgradeServiceCancel  | NICEXAI    |

### 其他基础信息获取

[官方文档](https://open.work.weixin.qq.com/api/doc/90001/90143/95148)

|       名称            | 请求方式  | URL                                     | 是否已实现   | 使用方法                            | 贡献者       |
| :--------------:     | -------- | :---------------------------------------| ---------- | -------------------------------   |------------|
| 获取客户基础信息        | POST     | /cgi-bin/kf/customer/batchget           | YES        | (r *Client) CustomerBatchGet      | NICEXAI    |
| 获取视频号绑定状态      | POST     |  /cgi-bin/kf/get_corp_qualification      | YES        | (r *Client) GetCorpQualification  | NICEXAI    |

## 应用管理

TODO

