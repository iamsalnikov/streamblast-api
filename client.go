package streamblast_api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
	"strconv"
)

type Client struct {
	BaseURI string
	DreamsContentID int64
}

func (c *Client) GetLinks(episodeID string) (map[string]string, error) {
	url := c.buildEpisodeURL(episodeID)

	resp, err := http.Get(url)
	if err != nil {
		return map[string]string{}, err
	}

	decoder := json.NewDecoder(resp.Body)
	var answer map[string]interface{}
	err = decoder.Decode(&answer)
	if err != nil {
		return map[string]string{}, err
	}

	if code, ok := answer["error_code"]; ok {
		return map[string]string{}, errors.New("error code: " + strconv.FormatFloat(code.(float64), 'f', 0, 64))
	}

	if message, ok := answer["error"]; ok {
		return map[string]string{}, errors.New("error: " + message.(string))
	}

	result := map[string]string{}
	for k, v := range answer {
		result[k] = v.(string)
	}

	return result, nil
}

func (c *Client) buildEpisodeURL(episodeID string) string {
	return fmt.Sprintf(
		"%s/video/%s/manifest_mp4.json?dreams_content_id=%d",
		c.BaseURI,
		episodeID,
		c.DreamsContentID,
	)
}
