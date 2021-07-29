package coingecko

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/billzong/go-gecko/format"
	"github.com/billzong/go-gecko/v3/types"
)

// Exchanges list, paginated
func (c *Client) Exchanges(perPage int, page int) (*types.ExchangesDetail, error) {
	params := url.Values{}
	// per_page
	if perPage <= 0 || perPage > 250 {
		perPage = 100
	}
	params.Add("per_page", format.Int2String(perPage))
	params.Add("page", format.Int2String(page))

	url := fmt.Sprintf("%s/exchanges?%s", baseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.ExchangesDetail
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Exchanges list, no pagination required
func (c *Client) ExchangesList() (*types.ExchangesBase, error) {
	url := fmt.Sprintf("%s/exchanges/list", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.ExchangesBase
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
