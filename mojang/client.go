package mojang

import (
	"encoding/json"
	"fmt"
	"mc-server/config"
	"mc-server/utils"

	"github.com/valyala/fasthttp"
)

// Client represents an API mojang
type Client struct {
	http *fasthttp.Client
}

// New creates a new fasthttp mojang and wraps it into an API mojang
func New() *Client {
	return &Client{
		http: &fasthttp.Client{
			Name: "mojango",
		},
	}
}

// FetchStatus fetches the states of all Mojang services and wraps them into a single object
func (client *Client) FetchStatus() (*Status, error) {
	// Call the Mojang status endpoint
	code, body, err := client.http.Get(nil, fmt.Sprintf("%s/check", uriStatus))
	if err != nil {
		return nil, err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the result into a status object and return it
	return parseStatusFromBody(body)
}

// FetchUUID fetches the current UUID of the given username
func (client *Client) FetchUUID(username string) (string, error) {
	return client.FetchUUIDAtTime(username, -1)
}

// FetchUUIDAtTime fetches the UUID of the given username at a given timestamp
func (client *Client) FetchUUIDAtTime(username string, timestamp int64) (string, error) {
	// Call the Mojang profile endpoint
	atExtension := ""
	if timestamp >= 0 {
		atExtension = fmt.Sprintf("?at=%d", timestamp)
	}
	code, body, err := client.http.Get(nil, fmt.Sprintf("%s/users/profiles/minecraft/%s%s", uriApi, username, atExtension))
	if err != nil {
		return "", err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return "", errorFromCode(code)
	}

	// Parse the result into a map containing the profile data
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	// Return the UUID of the requested profile
	return result["id"].(string), nil
}

// FetchMultipleUUIDs fetches the UUIDs of the given usernames
func (client *Client) FetchMultipleUUIDs(usernames []string) (map[string]string, error) {
	// Define the request object
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)
	request.SetRequestURI(fmt.Sprintf("%s/profiles/minecraft", uriApi))
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")
	reqBody, err := json.Marshal(usernames)
	if err != nil {
		return nil, err
	}
	request.SetBody(reqBody)

	// Define the response object
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Call the Mojang profile endpoint
	err = client.http.Do(request, response)
	if err != nil {
		return nil, err
	}

	// Define the important response values
	code := response.StatusCode()
	body := response.Body()

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the response body into a list of results
	var rawResults []struct {
		UUID string `json:"id"`
		Name string `json:"name"`
	}
	err = json.Unmarshal(body, &rawResults)
	if err != nil {
		return nil, err
	}

	// Parse the list of results into a map and return it
	result := make(map[string]string)
	for _, rawResult := range rawResults {
		result[rawResult.Name] = rawResult.UUID
	}
	return result, nil
}

// FetchNameHistory fetches all names of the given UUID and their corresponding changing timestamps
func (client *Client) FetchNameHistory(uuid string) ([]NameHistoryEntry, error) {
	// Call the Mojang profile endpoint
	code, body, err := client.http.Get(nil, fmt.Sprintf("%s/user/profiles/%s/names", uriApi, uuid))
	if err != nil {
		return nil, err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the response body into a list of name history entries and return it
	var entries []NameHistoryEntry
	err = json.Unmarshal(body, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// FetchProfile fetches the profile of the given UUID
// unsigned 将不对mojang请求起作用，mojang的profile默认为unsigned=false，
func (client *Client) FetchProfile(uuid string, unsigned bool) (*Profile, error) {
	// Call the Mojang profile endpoint
	code, body, err := client.http.Get(nil, fmt.Sprintf("%s/session/minecraft/profile/%s", uriSession, uuid))
	if err != nil {
		return nil, err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the response body into a profile and return it
	profile := new(Profile)
	err = json.Unmarshal(body, profile)
	if err != nil {
		return nil, err
	}
	if !unsigned {
		//用本地文件签名
		newProfiles := make([]ProfileProperty, len(profile.Properties))
		for i, v := range profile.Properties {
			signature, _ := utils.EncryptOAEP(v.Value, []byte(config.PublicKey))
			newProfiles[i] = ProfileProperty{
				Name:      v.Name,
				Value:     v.Value,
				Signature: signature,
			}
		}
		profile.Properties = newProfiles
	}
	return profile, nil
}
