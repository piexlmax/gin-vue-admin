package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type UserMFAInfo struct {
	Isopen bool    `json:"isopen"`
	QrCode    string `json:"qrcode"`
	AccountName string `json:"accountname"`
	Secret  string `json:"secret"`
}

// Find by id structure
type UserBindMfainfo struct {
	Id float64 `json:"id" form:"id"`
	AuthKey string `json:"authkey" form:"authkey"`
}