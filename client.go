package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	//Basic OAuth related URLs
	OAUTH_REQUES_TOKEN string = "https://api.twitter.com/oauth/request_token"
	OAUTH_AUTH_TOKEN   string = "https://api.twitter.com/oauth/authorize"
	OAUTH_ACCESS_TOKEN string = "https://api.twitter.com/oauth/access_token"

	//List API URLs
	API_BASE              string = "https://api.twitter.com/1.1/"
	API_TIMELINE          string = API_BASE + "statuses/home_timeline.json"
	API_MENTIONS_TIMELINE string = API_BASE + "statuses/mentions_timeline.json"
	API_USER_TIMELINE     string = API_BASE + "statuses/user_timeline.json"
	API_FOLLOWERS_IDS     string = API_BASE + "followers/ids.json"
	API_FOLLOWERS_LIST    string = API_BASE + "followers/list.json"
	API_FOLLOWER_INFO     string = API_BASE + "users/show.json"
)

var (
	ErrNotAuth = errors.New("no client OAuth")
)
type Client struct {
	HttpConn *http.Client
}

func (c *Client) HasAuth() bool {
	return c.HttpConn != nil
}

func (c *Client) BasicQuery(queryString string) ([]byte, error) {
	if c.HttpConn == nil {
		return nil, ErrNotAuth
	}

	response, err := c.HttpConn.Get(queryString)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	return bits, err
}

// User Timeline by UserID
func (c *Client) QueryUserTimelineByUserID(user_id string) (UserTimeline, []byte, error) {
	requesURL := fmt.Sprintf("%s?user_id=%s", API_USER_TIMELINE, user_id)
	data, err := c.BasicQuery(requesURL)
	ret := UserTimeline{}
	if err != nil {
		return ret, nil, err
	}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// User Timeline by ScreenName
func (c *Client) QueryUserTimelineByScreenName(ScreeName string) (UserTimeline, []byte, error) {
	requesURL := fmt.Sprintf("%s?screen_name=%s", API_USER_TIMELINE, ScreeName)
	data, err := c.BasicQuery(requesURL)
	ret := UserTimeline{}
	if err != nil {
		return ret, nil, err
	}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// Mentions Timeline.
func (c *Client) QueryMentionsTimeline(count int) (TimelineTweets, []byte, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_MENTIONS_TIMELINE, count)
	data, err := c.BasicQuery(requesURL)
	ret := TimelineTweets{}
	if err != nil {
		return ret, nil, err
	}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// Query timeline by count
func (c *Client) QueryTimeLine(count int) (MentionsTimeline, []byte, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_TIMELINE, count)
	data, err := c.BasicQuery(requesURL)
	ret := MentionsTimeline{}
	if err != nil {
		return ret, nil, err
	}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// Query follower timeline by count
func (c *Client) QueryFollower(count int) (Followers, []byte, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_FOLLOWERS_LIST, count)
	data, err := c.BasicQuery(requesURL)
	ret := Followers{}
	if err != nil {
		return ret, nil, err
	}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// Query FollowerID by count
func (c *Client) QueryFollowerIDs(count int) (FollowerIDs, []byte, error) {
	requesURL := fmt.Sprintf("%s?count=%d", API_FOLLOWERS_IDS, count)
	data, err := c.BasicQuery(requesURL)
	var ret FollowerIDs
	if err != nil {
		return ret, nil, err
	}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}

// Query FollowerID by ID.
func (c *Client) QueryFollowerById(id int) (UserDetail, []byte, error) {
	requesURL := fmt.Sprintf("%s?user_id=%d", API_FOLLOWER_INFO, id)
	data, err := c.BasicQuery(requesURL)
	var ret UserDetail
	if err != nil {
		return ret, nil, err
	}
	err = json.Unmarshal(data, &ret)
	return ret, data, err
}
