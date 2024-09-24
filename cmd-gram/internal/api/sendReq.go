package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shynn2/cmd-gram/internal/models"
)

func sendReq(msg *models.MessageDTO, meth string, url string) (*http.Response, error) {
	var client = &http.Client{}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	if jsonData == nil && meth != http.MethodGet {
		return nil, fmt.Errorf("data is nil")
	}

	req, err := http.NewRequest(meth, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("somthing went wrong, check your connection %v", err)

	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can`t do req due to error: %v", err)
	}

	return resp, nil
}
