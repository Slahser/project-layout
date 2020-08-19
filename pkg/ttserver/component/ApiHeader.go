package component

type ApiEnv string

const (
	PRODUCTION ApiEnv = "PRODUCTION"
	STAGING    ApiEnv = "STAGING"
	TEST       ApiEnv = "TEST"
	DAILY      ApiEnv = "DAILY"
)

type ApiHeader struct {
	Org     string `json:"org" validate:"required" binding:"required"`
	Project string `json:"project" validate:"required" binding:"required"`
	Env     string `json:"env" validate:"required,oneof='PRODUCTION' 'STAGING' 'TEST' 'DAILY'" binding:"required,oneof='PRODUCTION' 'STAGING' 'TEST' 'DAILY'"`
}
