package viber

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// PostData to viber API
func (v *Viber) PostData(ctx context.Context, url string, i interface{}) ([]byte, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	v.Logger.Debug("POST data: %s", string(b))

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(b))
	req.Header.Add("X-Viber-Auth-Token", v.AppKey)
	req.Close = true

	if v.client == nil {
		v.client = &http.Client{}
	}

	resp, err := v.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
