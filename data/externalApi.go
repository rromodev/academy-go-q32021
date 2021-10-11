package data

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/rromodev/academy-go-q32021/model"
)

type ExternalData struct {
	url string
}

func NewExternalData(url string) ExternalData {
	return ExternalData{url}
}

const USER = "user"

func (ed ExternalData) GetFromUrl() ([]string, error) {
	resp, err := http.Get(ed.url)
	if err != nil {
		return nil, errors.New("external service error")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("external response error")
	}

	var response model.Response
	json.Unmarshal([]byte(body), &response)

	user := []string{
		strconv.Itoa(response.Data.ID),
		USER,
		response.Data.Name,
		response.Data.LastName,
		response.Data.Email,
		response.Data.Avatar,
	}

	return user, nil
}
