# 企业微信

[官方文档](https://work.weixin.qq.com/api/doc)

## 快速入门

## 会议-预约会议高级管理

[官方文档](https://developer.work.weixin.qq.com/document/path/98148)

### 使用示例

```go
//初始化
wc := work.NewWork(&config.Config{
    CorpID:     "corp_id",
    CorpSecret: "corp_secret",
})

//获取会议接口实例
meetingClient := wc.GetMeeting()

//创建预约会议
resp, err := meetingClient.MeetingCreate(&meeting.MeetingCreateRequest{
    AdminUserID:     "zhangsan",
    Title:           "新建会议",
    MeetingStart:    1600000000,
    MeetingDuration: 3600,
    Description:     "新建会议描述",
    Location:        "广州媒体港",
    Invitees: &meeting.Invitees{
        UserID: []string{"lisi", "wangwu"},
    },
    Settings: &meeting.Settings{
        Password:          "1234",
        EnableWaitingRoom: &[]bool{false}[0],
        AllowEnterBeforeHost: &[]bool{true}[0],
        EnableEnterMute:   1,
    },
})

//获取会议详情
info, err := meetingClient.MeetingGetInfo(&meeting.MeetingGetInfoRequest{
    MeetingID: "XXXXXXXXX",
})

//取消预约会议
err = meetingClient.MeetingCancel(&meeting.MeetingCancelRequest{
    MeetingID: "XXXXXXXXX",
})
```

### 接口列表

| 接口 | 方法 | 文档 |
| --- | --- | --- |
| 创建预约会议 | `MeetingCreate` | [98148](https://developer.work.weixin.qq.com/document/path/98148) |
| 修改预约会议 | `MeetingUpdate` | [98154](https://developer.work.weixin.qq.com/document/path/98154) |
| 取消预约会议 | `MeetingCancel` | [98153](https://developer.work.weixin.qq.com/document/path/98153) |
| 获取会议详情 | `MeetingGetInfo` | [98149](https://developer.work.weixin.qq.com/document/path/98149) |
| 获取会议受邀成员列表 | `GetInvitees` | [98160](https://developer.work.weixin.qq.com/document/path/98160) |
| 更新会议受邀成员列表 | `SetInvitees` | [98162](https://developer.work.weixin.qq.com/document/path/98162) |
| 获取成员会议ID列表 | `GetUserMeetingID` | [98714](https://developer.work.weixin.qq.com/document/path/98714) |
| 创建用户专属参会链接 | `CreateCustomerShortURL` | [98818](https://developer.work.weixin.qq.com/document/path/98818) |
| 获取用户专属参会链接 | `GetCustomerShortURL` | [98819](https://developer.work.weixin.qq.com/document/path/98819) |
| 获取实时会中成员列表 | `GetRealtimeAttendeeList` | [98157](https://developer.work.weixin.qq.com/document/path/98157) |
| 获取已参会成员列表 | `GetAttendeeList` | [98156](https://developer.work.weixin.qq.com/document/path/98156) |
| 获取实时等候室成员列表 | `WaitingRoomGetCurrentUserList` | [98163](https://developer.work.weixin.qq.com/document/path/98163) |
| 获取等候室成员记录 | `WaitingRoomGetUserList` | [98164](https://developer.work.weixin.qq.com/document/path/98164) |
| 获取成员设备是否入会 | `CheckDeviceInMeeting` | [98165](https://developer.work.weixin.qq.com/document/path/98165) |
| 获取会议嘉宾列表 | `GetGuests` | [99039](https://developer.work.weixin.qq.com/document/path/99039) |
| 更新会议嘉宾列表 | `SetGuests` | [99040](https://developer.work.weixin.qq.com/document/path/99040) |
| 获取会议健康度 | `GetQuality` | [98821](https://developer.work.weixin.qq.com/document/path/98821) |
| 修改会议报名配置 | `EnrollSetConfig` | [98797](https://developer.work.weixin.qq.com/document/path/98797) |
| 获取会议报名配置 | `EnrollGetConfig` | [98800](https://developer.work.weixin.qq.com/document/path/98800) |
| 获取会议成员报名ID | `EnrollQueryByTmpOpenID` | [98794](https://developer.work.weixin.qq.com/document/path/98794) |
| 获取会议报名信息 | `EnrollList` | [98810](https://developer.work.weixin.qq.com/document/path/98810) |
| 审批会议报名信息 | `EnrollApprove` | [98807](https://developer.work.weixin.qq.com/document/path/98807) |
| 导入会议报名信息 | `EnrollImport` | [98816](https://developer.work.weixin.qq.com/document/path/98816) |
| 删除会议报名信息 | `EnrollDelete` | [98817](https://developer.work.weixin.qq.com/document/path/98817) |