package meeting

// Invitees 邀请参与会议的成员
type Invitees struct {
	UserID []string `json:"userid"`
}

// Guest 会议嘉宾
type Guest struct {
	Area        string `json:"area"`
	PhoneNumber string `json:"phone_number"`
	GuestName   string `json:"guest_name,omitempty"`
}

// Hosts 会议主持人列表
type Hosts struct {
	UserID []string `json:"userid"`
}

// RingUsers 指定响铃的成员列表
type RingUsers struct {
	UserID []string `json:"userid"`
}

// Settings 会议配置
type Settings struct {
	Password                  string     `json:"password,omitempty"`
	EnableWaitingRoom         *bool      `json:"enable_waiting_room,omitempty"`
	AllowEnterBeforeHost      *bool      `json:"allow_enter_before_host,omitempty"`
	EnableEnterMute           uint32     `json:"enable_enter_mute,omitempty"`
	AllowUnmuteSelf           *bool      `json:"allow_unmute_self,omitempty"`
	MuteAll                   *bool      `json:"mute_all,omitempty"`
	AllowExternalUser         *bool      `json:"allow_external_user,omitempty"`
	EnableScreenWatermark     *bool      `json:"enable_screen_watermark,omitempty"`
	WatermarkType             uint32     `json:"watermark_type,omitempty"`
	AutoRecordType            string     `json:"auto_record_type,omitempty"`
	AttendeeJoinAutoRecord    *bool      `json:"attendee_join_auto_record,omitempty"`
	EnableHostPauseAutoRecord *bool      `json:"enable_host_pause_auto_record,omitempty"`
	EnableInterpreter         *bool      `json:"enable_interpreter,omitempty"`
	EnableEnroll              *bool      `json:"enable_enroll,omitempty"`
	EnableHostKey             *bool      `json:"enable_host_key,omitempty"`
	HostKey                   string     `json:"host_key,omitempty"`
	Hosts                     *Hosts     `json:"hosts,omitempty"`
	RemindScope               uint32     `json:"remind_scope,omitempty"`
	RingUsers                 *RingUsers `json:"ring_users,omitempty"`
	// 以下字段仅在获取会议详情响应中返回
	NeedPassword              *bool  `json:"need_password,omitempty"`
	EnableEnterMuteType       uint32 `json:"enable_enter_mute_type,omitempty"`
	EnableDocUploadPermission *bool  `json:"enable_doc_upload_permission,omitempty"`
	CurrentHosts              *Hosts `json:"current_hosts,omitempty"`
	CoHosts                   *Hosts `json:"co_hosts,omitempty"`
}

// Reminders 周期性会议配置
type Reminders struct {
	IsRepeat         uint32   `json:"is_repeat,omitempty"`
	RepeatType       uint32   `json:"repeat_type,omitempty"`
	IsCustomRepeat   uint32   `json:"is_custom_repeat,omitempty"`
	RepeatUntilType  uint32   `json:"repeat_until_type,omitempty"`
	RepeatUntilCount uint32   `json:"repeat_until_count,omitempty"`
	RepeatUntil      uint32   `json:"repeat_until,omitempty"`
	RepeatInterval   uint32   `json:"repeat_interval,omitempty"`
	RepeatDayOfWeek  []uint32 `json:"repeat_day_of_week,omitempty"`
	RepeatDayOfMonth []uint32 `json:"repeat_day_of_month,omitempty"`
	RemindBefore     []uint32 `json:"remind_before,omitempty"`
}

// Attendees 会议成员
type Attendees struct {
	Member          []*Member          `json:"member"`
	TmpExternalUser []*TmpExternalUser `json:"tmp_external_user"`
}

// Member 企业内部成员
type Member struct {
	UserID         string `json:"userid"`
	Status         uint32 `json:"status"`
	FirstJoinTime  uint32 `json:"first_join_time"`
	LastQuitTime   uint32 `json:"last_quit_time"`
	TotalJoinCount uint32 `json:"total_join_count"`
	CumulativeTime uint32 `json:"cumulative_time"`
}

// TmpExternalUser 会中参会的外部联系人
type TmpExternalUser struct {
	TmpExternalUserID string `json:"tmp_external_userid"`
	Status            uint32 `json:"status"`
	FirstJoinTime     uint32 `json:"first_join_time"`
	LastQuitTime      uint32 `json:"last_quit_time"`
	TotalJoinCount    uint32 `json:"total_join_count"`
	CumulativeTime    uint32 `json:"cumulative_time"`
}

// SubMeeting 周期性子会议
type SubMeeting struct {
	SubMeetingID string `json:"sub_meetingid"`
	Status       uint32 `json:"status"`
	StartTime    uint32 `json:"start_time"`
	EndTime      uint32 `json:"end_time"`
	Title        string `json:"title"`
	RepeatID     string `json:"repeat_id"`
}

// SubRepeatInfo 周期性会议分段信息
type SubRepeatInfo struct {
	RepeatID         string   `json:"repeat_id"`
	RepeatType       uint32   `json:"repeat_type"`
	IsCustomRepeat   uint32   `json:"is_custom_repeat"`
	RepeatInterval   uint32   `json:"repeat_interval"`
	RepeatDayOfWeek  []uint32 `json:"repeat_day_of_week"`
	RepeatDayOfMonth []uint32 `json:"repeat_day_of_month"`
	RepeatUntilType  uint32   `json:"repeat_until_type"`
	RepeatUntilCount uint32   `json:"repeat_until_count"`
	RepeatUntil      uint32   `json:"repeat_until"`
}
