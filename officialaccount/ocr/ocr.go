package ocr

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
	"net/url"
)

const (
	ocrIDCardURL         = "https://api.weixin.qq.com/cv/ocr/idcard"
	ocrBankCardURL       = "https://api.weixin.qq.com/cv/ocr/bankcard"
	ocrDrivingURL        = "https://api.weixin.qq.com/cv/ocr/driving"
	ocrDrivingLicenseURL = "https://api.weixin.qq.com/cv/ocr/drivinglicense"
	ocrBizLicenseURL     = "https://api.weixin.qq.com/cv/ocr/bizlicense"
	ocrCommonURL         = "https://api.weixin.qq.com/cv/ocr/comm"
	ocrPlateNumberURL    = "https://api.weixin.qq.com/cv/ocr/platenum"
)

type OCR struct {
	*context.Context
}

type coordinate struct {
	X int64 `json:"x,omitempty"`
	Y int64 `json:"y,omitempty"`
}

type position struct {
	LeftTop     coordinate `json:"left_top"`
	RightTop    coordinate `json:"right_top"`
	RightBottom coordinate `json:"right_bottom"`
	LeftBottom  coordinate `json:"left_bottom"`
}

type imageSize struct {
	Width  int64 `json:"w,omitempty"`
	Height int64 `json:"h,omitempty"`
}

type resDriving struct {
	util.CommonError

	PlateNumber       string              `json:"plate_num,omitempty,omitempty"`
	VehicleType       string              `json:"vehicle_type,omitempty"`
	Owner             string              `json:"owner,omitempty"`
	Address           string              `json:"addr,omitempty"`
	UseCharacter      string              `json:"use_character,omitempty"`
	Model             string              `json:"model,omitempty"`
	Vin               string              `json:"vin,omitempty"`
	EngineNumber      string              `json:"engine_num,omitempty"`
	RegisterDate      string              `json:"register_date,omitempty"`
	IssueDate         string              `json:"issue_date,omitempty"`
	PlateNumberB      string              `json:"plate_num_b,omitempty"`
	Record            string              `json:"record,omitempty"`
	PassengersNumber  string              `json:"passengers_num,omitempty"`
	TotalQuality      string              `json:"total_quality,omitempty"`
	PrepareQuality    string              `json:"prepare_quality,omitempty"`
	OverallSize       string              `json:"overall_size,omitempty"`
	CardPositionFront map[string]position `json:"card_position_front,omitempty"`
	CardPositionBack  map[string]position `json:"card_position_back,omitempty"`
	ImageSize         imageSize           `json:"img_size,omitempty"`
}

type resIDCard struct {
	util.CommonError

	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
	Address     string `json:"addr,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	ValidDate   string `json:"valid_date,omitempty"`
}

type resBankCard struct {
	util.CommonError

	Number string `json:"number,omitempty"`
}

type resDrivingLicense struct {
	util.CommonError

	IdNumber     string `json:"id_num,omitempty"`
	Name         string `json:"name,omitempty"`
	Sex          string `json:"sex,omitempty"`
	Nationality  string `json:"nationality,omitempty"`
	Address      string `json:"address,omitempty"`
	Birthday     string `json:"birth_date,omitempty"`
	IssueDate    string `json:"issue_date,omitempty"`
	CarClass     string `json:"car_class,omitempty"`
	ValidFrom    string `json:"valid_from,omitempty"`
	ValidTo      string `json:"valid_to,omitempty"`
	OfficialSeal string `json:"official_seal,omitempty"`
}

type resBizLicense struct {
	util.CommonError

	RegisterNumber      string              `json:"reg_num,omitempty"`
	Serial              string              `json:"serial,omitempty"`
	LegalRepresentative string              `json:"legal_representative,omitempty"`
	EnterpriseName      string              `json:"enterprise_name,omitempty"`
	TypeOfOrganization  string              `json:"type_of_organization,omitempty"`
	Address             string              `json:"address,omitempty"`
	TypeOfEnterprise    string              `json:"type_of_enterprise,omitempty"`
	BusinessScope       string              `json:"business_scope,omitempty"`
	RegisteredCapital   string              `json:"registered_capital,omitempty"`
	PaidInCapital       string              `json:"paid_in_capital,omitempty"`
	ValidPeriod         string              `json:"valid_period,omitempty"`
	RegisterDate        string              `json:"registered_date,omitempty"`
	CertPosition        map[string]position `json:"cert_position,omitempty"`
	ImageSize           imageSize           `json:"img_size,omitempty"`
}

type resCommon struct {
	util.CommonError

	Items     []commonItem `json:"items,omitempty"`
	ImageSize imageSize    `json:"img_size,omitempty"`
}

type commonItem struct {
	Position position `json:"pos"`
	Text     string   `json:"text"`
}

type resPlateNumber struct {
	util.CommonError

	Number string `json:"number"`
}

//NewOCR 实例
func NewOCR(c *context.Context) *OCR {
	ocr := new(OCR)
	ocr.Context = c
	return ocr
}

//身份证OCR识别接口
func (ocr *OCR) IdCard(path string) (resIDCard resIDCard, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrIDCardURL, url.QueryEscape(path), accessToken)

	response, err := util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resIDCard)
	if err != nil {
		return
	}

	if resIDCard.ErrCode != 0 {
		err = fmt.Errorf("OCRIdCard Error , errcode=%d , errmsg=%s", resIDCard.ErrCode, resIDCard.ErrMsg)
		return
	}

	return
}

//银行卡OCR识别接口
func (ocr *OCR) BankCard(path string) (resBankCard resBankCard, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrBankCardURL, url.QueryEscape(path), accessToken)

	response, err := util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resBankCard)
	if err != nil {
		return
	}

	if resBankCard.ErrCode != 0 {
		err = fmt.Errorf("OCRBankCard Error , errcode=%d , errmsg=%s", resBankCard.ErrCode, resBankCard.ErrMsg)
		return
	}

	return
}

//行驶证OCR识别接口
func (ocr *OCR) Driving(path string) (resDriving resDriving, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrDrivingURL, url.QueryEscape(path), accessToken)

	response, err := util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resDriving)
	if err != nil {
		return
	}

	if resDriving.ErrCode != 0 {
		err = fmt.Errorf("OCRBankCard Error , errcode=%d , errmsg=%s", resDriving.ErrCode, resDriving.ErrMsg)
		return
	}

	return
}

//驾驶证OCR识别接口
func (ocr *OCR) DrivingLicense(path string) (resDrivingLicense resDrivingLicense, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrDrivingLicenseURL, url.QueryEscape(path), accessToken)

	response, err := util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resDrivingLicense)
	if err != nil {
		return
	}

	if resDrivingLicense.ErrCode != 0 {
		err = fmt.Errorf("OCRBankCard Error , errcode=%d , errmsg=%s", resDrivingLicense.ErrCode, resDrivingLicense.ErrMsg)
		return
	}

	return
}

//营业执照OCR识别接口
func (ocr *OCR) BizLicense(path string) (resBizLicense resBizLicense, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrBizLicenseURL, url.QueryEscape(path), accessToken)

	response, err := util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resBizLicense)
	if err != nil {
		return
	}

	if resBizLicense.ErrCode != 0 {
		err = fmt.Errorf("OCRBankCard Error , errcode=%d , errmsg=%s", resBizLicense.ErrCode, resBizLicense.ErrMsg)
		return
	}

	return
}

//通用印刷体OCR识别接口
func (ocr *OCR) Common(path string) (resCommon resCommon, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrCommonURL, url.QueryEscape(path), accessToken)

	response, err := util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resCommon)
	if err != nil {
		return
	}

	if resCommon.ErrCode != 0 {
		err = fmt.Errorf("OCRBankCard Error , errcode=%d , errmsg=%s", resCommon.ErrCode, resCommon.ErrMsg)
		return
	}

	return
}

//车牌OCR识别接口
func (ocr *OCR) PlateNumber(path string) (resPlateNumber resPlateNumber, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?img_url=%s&access_token=%s", ocrPlateNumberURL, url.QueryEscape(path), accessToken)

	response, err := util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &resPlateNumber)
	if err != nil {
		return
	}

	if resPlateNumber.ErrCode != 0 {
		err = fmt.Errorf("OCRBankCard Error , errcode=%d , errmsg=%s", resPlateNumber.ErrCode, resPlateNumber.ErrMsg)
		return
	}

	return
}
