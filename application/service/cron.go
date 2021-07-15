package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	bc "github.com/lizuguang/spectrum-blockchain/application/blockchain"
	"github.com/lizuguang/spectrum-blockchain/application/lib"
	"log"
	"time"
)

const spec = "0 0 0 * * ?" // 每天0点执行
//const spec = "*/10 * * * * ?" //10秒执行一次，用于测试

func Init() {
	c := cron.New(cron.WithSeconds()) //支持到秒级别
	_, err := c.AddFunc(spec, GoRun)
	if err != nil {
		log.Printf("定时任务开启失败 %s", err)
	}
	c.Start()
	log.Printf("定时任务已开启")
	select {}
}

func GoRun() {
	log.Printf("定时任务已启动")
	//先把所有出租查询出来
	resp, err := bc.ChannelQuery("querySellingList", [][]byte{}) //调用智能合约
	if err != nil {
		log.Printf("定时任务-querySellingList失败%s", err.Error())
		return
	}
	// 反序列化json
	var data []lib.Selling
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		log.Printf("定时任务-反序列化json失败%s", err.Error())
		return
	}
	for _, v := range data {
		//把状态为出租中和交付中的筛选出来
		if v.SellingStatus == lib.SellingStatusConstant()["saleStart"] ||
			v.SellingStatus == lib.SellingStatusConstant()["delivery"] {
			//有效期天数
			day, _ := time.ParseDuration(fmt.Sprintf("%dh", v.SalePeriod*24))
			local, _ := time.LoadLocation("Local")
			t, _ := time.ParseInLocation("2006-01-02 15:04:05", v.CreateTime, local)
			vTime := t.Add(day)
			//如果 time.Now()大于 vTime 说明过期
			if time.Now().Local().After(vTime) {
				//将状态更改为已过期
				var bodyBytes [][]byte
				bodyBytes = append(bodyBytes, []byte(v.ObjectOfSale))
				bodyBytes = append(bodyBytes, []byte(v.Seller))
				bodyBytes = append(bodyBytes, []byte(v.Buyer))
				bodyBytes = append(bodyBytes, []byte("expired"))
				//调用智能合约
				resp, err := bc.ChannelExecute("updateSelling", bodyBytes)
				if err != nil {
					return
				}
				var data map[string]interface{}
				if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
					return
				}
				fmt.Println(data)
			}
		}
	}
}
