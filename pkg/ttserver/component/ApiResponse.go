package component

type ApiResponse struct {
	GmtCreate string `json:"gmt_create" validate:"required,datetime=2006-01-02 22:30:01"`
	Endpoint  string `json:"endpoint" validate:"required,url"`
}
