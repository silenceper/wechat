package datacube

import (
	"github.com/silenceper/wechat/v2/officialaccount/context"
)

type reqDate struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type DataCube struct {
	*context.Context
}

func NewCube(context *context.Context) *DataCube {
	dataCube := new(DataCube)
	dataCube.Context = context
	return dataCube
}
