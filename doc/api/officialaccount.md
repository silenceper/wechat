# 微信公众号API列表

## 基础接口

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html)

|          名称           | 请求方式 | URL                        | 是否已实现 | 使用方法 |
| :---------------------: | -------- | :------------------------- | ---------- | -------- |
|    获取Access token     | GET      | /cgi-bin/token             | YES        |          |
|  获取微信服务器IP地址   | GET      | /cgi-bin/get_api_domain_ip | YES        |          |
| 获取微信callback IP地址 | GET      | /cgi-bin/getcallbackip     | YES        |          |
|    清理接口调用次数     | POST     | /cgi-bin/clear_quota       | YES        |          |

## 订阅通知

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html)

| 名称                 | 请求方式 | URL                                    | 是否已实现 | 使用方法                |
| -------------------- | -------- | -------------------------------------- | ---------- | ----------------------- |
| 选用模板             | POST     | /wxaapi/newtmpl/addtemplate            | YES        | (tpl *Subscribe) Add    |
| 删除模板             | POST     | /wxaapi/newtmpl/deltemplate            | YES        | (tpl *Subscribe) Delete |
| 获取公众号类目       | GET      | /wxaapi/newtmpl/getcategory            | NO         |                         |
| 获取模板中的关键词   | GET      | /wxaapi/newtmpl/getpubtemplatekeywords | NO         |                         |
| 获取类目下的公共模板 | GET      | /wxaapi/newtmpl/getpubtemplatetitles   | NO         |                         |
| 获取私有模板列表     | GET      | /wxaapi/newtmpl/gettemplate            | YES        | (tpl *Subscribe) List() |
| 发送订阅通知         | POST     | /cgi-bin/message/subscribe/bizsend     | YES        | (tpl *Subscribe) Send   |

## 客服消息

[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html)

|      |      |      |
| ---- | ---- | ---- |
|      |      |      |
|      |      |      |
|      |      |      |

