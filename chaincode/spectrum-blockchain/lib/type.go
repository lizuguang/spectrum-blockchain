package lib

// Account 账户，虚拟管理员和若干用户账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

// RealEstate 频谱资源作为担保出租、转让或质押时Encumbrance为true，默认状态false。
// 仅当Encumbrance为false时，才可发起出租、转让或质押
// Proprietor和RealEstateID一起作为复合键,保证可以通过Proprietor查询到名下所有的频谱信息
type RealEstate struct {
	RealEstateID string  `json:"realEstateId"` //频谱资源ID
	Proprietor   string  `json:"proprietor"`   //所有者(用户)(用户AccountId)
	Encumbrance  bool    `json:"encumbrance"`  //是否作为担保
	MinFrequency  float64  `json:"minfrequency"`    //频率最小值
	MaxFrequency  float64  `json:"maxfrequency"`    //频率最大值
	StartDate   string   `json:"startdate"`  //开始时间
	EndDate   string   `json:"enddate"`  //结束时间
}

// Selling 出租要约
// 需要确定ObjectOfSale是否属于Seller
// 买家初始为空
// Seller和ObjectOfSale一起作为复合键,保证可以通过seller查询到名下所有发起的出租
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //出租对象(正在出租的频谱资源RealEstateID)
	Seller        string  `json:"seller"`        //发起出租人、卖家(卖家AccountId)
	Buyer         string  `json:"buyer"`         //参与出租人、买家(买家AccountId)
	Price         float64 `json:"price"`         //价格
	CreateTime    string  `json:"createTime"`    //创建时间
	SalePeriod    int     `json:"salePeriod"`    //智能合约的有效期(单位为天)
	SellingStatus string  `json:"sellingStatus"` //出租状态
}

// SellingStatusConstant 出租状态
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "出租中", //正在出租状态,等待买家光顾
		"cancelled": "已取消", //被卖家取消出租或买家退款操作导致取消
		"expired":   "已过期", //出租期限到期
		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
		"done":      "完成",  //卖家确认接收资金，交易完成
	}
}

// SellingBuy 买家参与出租
// 出租对象不能是买家发起的
// Buyer和CreateTime作为复合键,保证可以通过buyer查询到名下所有参与的出租
type SellingBuy struct {
	Buyer      string  `json:"buyer"`      //参与出租人、买家(买家AccountId)
	CreateTime string  `json:"createTime"` //创建时间
	Selling    Selling `json:"selling"`    //出租对象
}

// Donating 转让要约
// 需要确定ObjectOfDonating是否属于Donor
// 需要指定接收人Grantee，并等待接收人同意接收
type Donating struct {
	ObjectOfDonating string `json:"objectOfDonating"` //转让对象(正在转让的频谱资源RealEstateID)
	Donor            string `json:"donor"`            //转让人(转让人AccountId)
	Grantee          string `json:"grantee"`          //接收人(接收人AccountId)
	CreateTime       string `json:"createTime"`       //创建时间
	DonatingStatus   string `json:"donatingStatus"`   //转让状态
}

// DonatingStatusConstant 转让状态
var DonatingStatusConstant = func() map[string]string {
	return map[string]string{
		"donatingStart": "转让中", //转让人发起转让合约，等待接收人确认接收
		"cancelled":     "已取消", //转让人在接收人确认接收之前取消转让或接收人取消接收接收
		"done":          "完成",  //接收人确认接收，交易完成
	}
}

// DonatingGrantee 供接收人查询的
type DonatingGrantee struct {
	Grantee    string   `json:"grantee"`    //接收人(接收人AccountId)
	CreateTime string   `json:"createTime"` //创建时间
	Donating   Donating `json:"donating"`   //转让对象
}

const (
	AccountKey         = "account-key"
	RealEstateKey      = "real-estate-key"
	SellingKey         = "selling-key"
	SellingBuyKey      = "selling-buy-key"
	DonatingKey        = "donating-key"
	DonatingGranteeKey = "donating-grantee-key"
)
