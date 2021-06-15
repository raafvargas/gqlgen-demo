package marvel

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func doGet(ctx context.Context, path string, id int) (*MarvelResponse, error) {
	uri := fmt.Sprintf("%s/%s/%d", MarvelEndpoint, path, id)

	req, _ := http.NewRequest("GET", uri, nil)

	return doRequest(ctx, req)
}

func doSearch(ctx context.Context, path string, keyValues ...string) (*MarvelResponse, error) {
	u, _ := url.Parse(fmt.Sprintf("%s/%s", MarvelEndpoint, path))
	q, _ := url.ParseQuery("")

	q.Set("offset", "0")
	q.Set("limit", "10")

	for i := 0; i < len(keyValues); i += 2 {
		q.Set(keyValues[i], keyValues[i+1])
	}

	u.RawQuery = q.Encode()

	req, _ := http.NewRequest("GET", u.String(), nil)

	return doRequest(ctx, req)
}

func doRequest(ctx context.Context, req *http.Request) (*MarvelResponse, error) {
	doRequestAuth(req)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data := &MarvelResponse{}

	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

func doRequestAuth(req *http.Request) {
	q := req.URL.Query()

	ts := fmt.Sprintf("%d", time.Now().Unix())
	hash := []byte(ts + MarvelPrivateKey + MarvelPublicKey)
	md := md5.Sum(hash)

	q.Add("ts", ts)
	q.Add("apikey", MarvelPublicKey)
	q.Add("hash", hex.EncodeToString(md[:]))

	req.URL.RawQuery = q.Encode()
}
