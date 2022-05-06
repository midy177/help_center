package biz

type JsonFormat struct {
	Code       int         `json:"code"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	PageNum    int         `json:"page_num"`
	ArticleNum int         `json:"article_num"`
	Data       interface{} `json:"data"`
}
type BaseJson struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
type LoginData struct {
	Name     string `json:"name"`
	Password string `json:"passwd"`
}
