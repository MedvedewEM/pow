package powclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MedvedewEM/pow/pkg/api"
)

// Interface defines methods of pow-server.
// Every custom client must implement this interface for compatibility
type Interface interface {
	// Please returns challenge-to-solve's ID and suffix
	Please() (string, string, error)

	// WisdomWord returns wisdom's word by id requested from
	// "Please" method and token calculated from client side
	WisdomWord(string, string) (string, error)
}

// New returns official implementation of pow-server interface
func New(url string) Interface {
	return &client{
		url: url,
	}
}

type client struct {
	url string
}

func (c *client) methodUrl(method string) string {
	return fmt.Sprintf("%s/%s", c.url, method)
}

// Please method. See Interface above
func (c *client) Please() (string, string, error) {
	resp, err := http.Get(c.methodUrl("api/v1/please"))
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	r := api.ChallengePleaseResponse{}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", "", err
	}

	return r.ID, r.Suffix, nil
}

// WisdomWord method. See Interface above
func (c *client) WisdomWord(authid, authtoken string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.methodUrl("api/v1/wisdom/word"), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("POWID", authid)
	req.Header.Set("POWToken", authtoken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	r := api.WisdomWordResponse{}
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", err
	}

	return r.Word, nil
}
