package proxy

import (
	"io/ioutil"

	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

// GetIndex returns the contetent of API.
func GetIndex(ctx context.Context) (string, error) {
	client := urlfetch.Client(ctx)

	resp, err := client.Get("http://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
