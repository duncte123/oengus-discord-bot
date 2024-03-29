package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"oenugs-bot/utils"
)

var httpClient = http.DefaultClient

var baseUrl = utils.GetEnv("OENGUS_API_BASE", "https://oengus.io/api/v1")

func GetMarathonStats(marathonId string) (*MarathonStats, error) {
	url := fmt.Sprintf("%s/marathons/%s/stats", baseUrl, marathonId)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("User-Agent", "oengus.io/bot")

	res, httpErr := httpClient.Do(req)
	if httpErr != nil {
		log.Println(httpErr)
		return nil, httpErr
	}

	if res.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("you are not authorised to do this action")
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("marathon not found")
	}

	// defer calls are not executed until the function returns
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Print(readErr)
		return nil, readErr
	}

	//log.Println(string(body))

	var response *MarathonStats

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Println(jsonErr)
		return response, jsonErr
	}

	return response, nil
}
