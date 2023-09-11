package schemas

type DataResponse struct {
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}

type Response struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}
