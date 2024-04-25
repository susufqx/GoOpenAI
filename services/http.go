/*
 * @Author: rui.li
 * @Date: 2024-04-24 17:39:54
 * @LastEditors: rui.li
 * @LastEditTime: 2024-04-24 18:14:19
 * @FilePath: /GoOpenAI/services/http.go
 */
package services

import (
	"context"
	"errors"

	"github.com/susufqx/GoOpenAI/utils"

	"github.com/go-resty/resty/v2"
)

var HttpClient *resty.Client

func init() {
	HttpClient = resty.New()
}

func postRawData(ctx context.Context, url string, datas []byte, headers map[string]string) ([]byte, error) {
	headers["Content-Type"] = "application/json"
	req := HttpClient.R().SetContext(ctx).SetHeaders(headers)
	resp, err := req.SetBody(datas).Post(url)
	if err != nil {
		return nil, err
	}

	utils.Logger.Debugf("Remote algo server response status code is : %d", resp.StatusCode())
	if resp.StatusCode() != 200 {
		utils.Logger.Debugf("Remote algo server response body is : %s", resp.Body())
		return nil, errors.New("remote algo server has no response")
	}

	return resp.Body(), nil
}
