package notification_library

import (
	"bytes"
	"net/http"
)

func httpPost(url string, headers map[string]string, msg []byte) error {
	postReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(msg))
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for k, v := range headers {
			postReq.Header.Add(k, v)
		}
	}
	httpClient := &http.Client{}
	_, errPosting := httpClient.Do(postReq)
	return errPosting
}
