package global

import "encoding/json"

type PagingCount struct {
	TOTAL      int `db:"total"`
	TOTALCHECK int `db:"totalcheck"`
}

type ALLLEARN struct {
	TOTAL      float64 `db:"total"`
	TOTALCHECK int     `db:"totalcheck"`
}

type LEARNEXPECTED struct {
	TOTALEXPECTED float64 `db:"totalexpected"`
}

type LEARNADJUST struct {
	TOTALADJUST float64 `db:"totaladjust"`
}

type PagingResult struct {
	TOTAL       int             `json:"total"`
	SUM         float64         `json:"sumtotal"`
	SUMEXPECTED float64         `json:"sumexpected"`
	SUMADJUST   float64         `json:"sumadjust"`
	DATA        json.RawMessage `json:"data"`
}

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APILoginResponse struct {
	Code    int    `json:"responseCode"`
	Message string `json:"message"`
}

type APIGetResponse struct {
	Code    int             `json:"responseCode"`
	Message string          `json:"responseDescription"`
	Data    json.RawMessage `json:"data"`
}
