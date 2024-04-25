/*
 * @Author: rui.li
 * @Date: 2024-04-24 17:22:25
 * @LastEditors: rui.li
 * @LastEditTime: 2024-04-24 17:25:10
 * @FilePath: /GoOpenAI/models/err_response_model.go
 */
package models

type ErrResponseModel struct {
	Error ErrResultModel `json:"error"`
}
