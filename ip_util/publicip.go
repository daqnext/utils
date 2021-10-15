package ip_util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IP struct {
	Query string
}

var publicip string

func GetPubIp(forceUpdate bool) (string, error) {

	if !forceUpdate {
		if publicip != "" {
			return publicip, nil
		}
	}

	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	var ip IP
	jsonErr := json.Unmarshal(body, &ip)
	if jsonErr != nil {
		return "", jsonErr
	}
	publicip = ip.Query
	return publicip, nil
}
