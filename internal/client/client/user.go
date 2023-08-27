package client

import (
	"bytes"
	"encoding/json"
	"math/big"
)

type User struct {
	TelegramID big.Int `json:"telegram_id"`
}

type UserRequest struct {
	TelegramID int64 `json:"telegram_id"`
}

func (c *Client) CreateUser(telegramID int64) (int, error) {
	const endpoint = "user/"

	req := UserRequest{
		TelegramID: telegramID,
	}

	requestBodyBytes, err := json.Marshal(req)
	if err != nil {
		return 0, err
	}

	resp, err := c.generalRequest("POST", endpoint, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
