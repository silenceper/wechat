package suite

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// GetPermanentCodeURL 获取企业永久授权码
	GetPermanentCodeURL = "https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code?suite_access_token=%s"
)

type (
	// GetPermanentCodeRequest 获取企业永久授权码请求
	GetPermanentCodeRequest struct {
		AuthCode string `json:"auth_code"`
	}
	// GetPermanentCodeResponse 获取企业永久授权码响应
	GetPermanentCodeResponse struct {
		util.CommonError

		AccessToken    string `json:"access_token"`
		ExpiresIn      int    `json:"expires_in"`
		PermanentCode  string `json:"permanent_code"`
		DealerCorpInfo struct {
			Corpid   string `json:"corpid"`
			CorpName string `json:"corp_name"`
		} `json:"dealer_corp_info"`
		AuthCorpInfo struct {
			Corpid            string `json:"corpid"`
			CorpName          string `json:"corp_name"`
			CorpType          string `json:"corp_type"`
			CorpSquareLogoUrl string `json:"corp_square_logo_url"`
			CorpUserMax       int    `json:"corp_user_max"`
			CorpFullName      string `json:"corp_full_name"`
			VerifiedEndTime   int    `json:"verified_end_time"`
			SubjectType       int    `json:"subject_type"`
			CorpWxqrcode      string `json:"corp_wxqrcode"`
			CorpScale         string `json:"corp_scale"`
			CorpIndustry      string `json:"corp_industry"`
			CorpSubIndustry   string `json:"corp_sub_industry"`
		} `json:"auth_corp_info"`
		AuthInfo struct {
			Agent []struct {
				Agentid          int    `json:"agentid"`
				Name             string `json:"name"`
				RoundLogoUrl     string `json:"round_logo_url"`
				SquareLogoUrl    string `json:"square_logo_url"`
				Appid            int    `json:"appid"`
				AuthMode         int    `json:"auth_mode,omitempty"`
				IsCustomizedApp  bool   `json:"is_customized_app,omitempty"`
				AuthFromThirdapp bool   `json:"auth_from_thirdapp,omitempty"`
				Privilege        struct {
					Level      int      `json:"level"`
					AllowParty []int    `json:"allow_party"`
					AllowUser  []string `json:"allow_user"`
					AllowTag   []int    `json:"allow_tag"`
					ExtraParty []int    `json:"extra_party"`
					ExtraUser  []string `json:"extra_user"`
					ExtraTag   []int    `json:"extra_tag"`
				} `json:"privilege,omitempty"`
				SharedFrom struct {
					Corpid    string `json:"corpid"`
					ShareType int    `json:"share_type"`
				} `json:"shared_from"`
			} `json:"agent"`
		} `json:"auth_info"`
		AuthUserInfo struct {
			Userid     string `json:"userid"`
			OpenUserid string `json:"open_userid"`
			Name       string `json:"name"`
			Avatar     string `json:"avatar"`
		} `json:"auth_user_info"`
		RegisterCodeInfo struct {
			RegisterCode string `json:"register_code"`
			TemplateId   string `json:"template_id"`
			State        string `json:"state"`
		} `json:"register_code_info"`
		State string `json:"state"`
	}
)

func (r *Client) GetPermanentCode(request *GetPermanentCodeRequest) (*GetPermanentCodeResponse, error) {
	var (
		response []byte
		err      error
	)
	jsonData, _ := json.Marshal(request)
	response, err = util.HTTPPost(fmt.Sprintf(GetPermanentCodeURL, r.SuiteAccessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	var result *GetPermanentCodeResponse
	err = util.DecodeWithError(response, &result, "GetPermanentCode")
	if err != nil {
		return nil, err
	}
	return result, nil
}
