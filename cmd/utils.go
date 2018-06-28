package cmd

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func doRequest(method, path string, body interface{}) error {

	r, err := newRequest(method, path, body)
	if err != nil {
		return err
	}

	res, err := client.Do(r)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBody))
	defer res.Body.Close()
	return nil
}

func newRequest(method, path string, body interface{}) (*http.Request, error) {
	var r *http.Request
	var err error
	if body != nil {
		br, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		r, err = http.NewRequest(method, "https://api.dnsmadeeasy.com/V2.0"+path, bytes.NewBuffer(br))
	} else {
		r, err = http.NewRequest(method, "https://api.dnsmadeeasy.com/V2.0"+path, nil)
	}
	if err != nil {
		return nil, err
	}
	addAuthHeaders(r)
	return r, nil
}

func addAuthHeaders(r *http.Request) {
	ts := time.Now().UTC().Format("Mon, 2 Jan 2006 15:04:05 MST")
	r.Header.Add("x-dnsme-apiKey", apiKey)
	r.Header.Add("x-dnsme-requestDate", ts)
	r.Header.Add("Content-Type", "application/json")
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(ts))
	sk := hex.EncodeToString(mac.Sum(nil))
	r.Header.Add("x-dnsme-hmac", sk)
}
