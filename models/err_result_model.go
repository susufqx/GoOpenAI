package models

type ErrResultModel struct {
	Message string  `json:"message"`
	Type    string  `json:"type"`
	Param   *string `json:"param,omitempty"`
	Code    any     `json:"code,omitempty"`
}
