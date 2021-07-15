package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/lizuguang/spectrum-blockchain/application/blockchain"
	"github.com/lizuguang/spectrum-blockchain/application/pkg/app"
	"net/http"
	"strconv"
)

type RealEstateRequestBody struct {
	AccountId   string  `json:"accountId"`   //操作人ID
	Proprietor  string  `json:"proprietor"`  //所有者(用户)(用户AccountId)
	MinFrequency   float64 `json:"minfrequency"`   //频率最小值
	MaxFrequency float64 `json:"maxfrequency"` //频谱最大值
	StartDate  string   `json:"startdate"` //开始时间
	EndDate  string   `json:"enddate"` //结束时间
}

type RealEstateQueryRequestBody struct {
	Proprietor string `json:"proprietor"` //所有者(用户)(用户AccountId)
}

// @Summary 新建频谱(管理员)
// @Param realEstate body RealEstateRequestBody true "realEstate"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/createRealEstate [post]
func CreateRealEstate(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.MinFrequency <= 0 || body.MaxFrequency <= 0 || body.MaxFrequency <= body.MinFrequency {
		appG.Response(http.StatusBadRequest, "失败", "频率的最小值必须小于最大值")
		return
	}
	a1 := " "
	var formattedStartDate string
	var formattedEndDate string
	if body.StartDate == "" {
		fmt.Sprintf("开始时间为空")
	} else {
		formattedStartDate = body.StartDate[:10] + a1 + body.StartDate[11:19]
	}
	if body.EndDate == "" {
		fmt.Sprintf("结束时间为空")
	} else {
		formattedEndDate = body.EndDate[:10] + a1 + body.EndDate[11:19]
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.MinFrequency, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.MaxFrequency, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(formattedStartDate))
	bodyBytes = append(bodyBytes, []byte(formattedEndDate))
	//调用智能合约
	resp, err := bc.ChannelExecute("createRealEstate", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// @Summary 获取频谱信息(空json{}可以查询所有，指定proprietor可以查询指定用户的频谱资源)
// @Param realEstateQuery body RealEstateQueryRequestBody true "realEstateQuery"
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/queryRealEstateList [post]
func QueryRealEstateList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Proprietor != "" {
		bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryRealEstateList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
