package chaojiying_sdk

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// IChaoJiYing ChaoJiYing sdk open interface
type ChaoJiYing interface {
	UserInfo() (userInfo *userInfoResp, err error) // Get User info
	IdentifyPic(
		codeType int,
		minLen int,
		imgBase64 string,
	) (
		identifyPic *identifyPicResp,
		err error,
	) // Identify pictures
	ReportError(picId string) (
		reportError *reportErrorResp,
		err error,
	) // Error report and return to score
	SetHttpsProxy(u string)     // Set Https Proxy
	SetTimeout(t time.Duration) // Set timeout (seconds)
	SetUser(user string)        // Set User account
	SetPass(pass string)        // Set User password
	SetPass2(pass2 string)      // Set DM5 value of user password (32-bit lower case)
	SetSoftId(softId string)    // Set Software ID: in the user center, the software ID can be generated
}

// chaoJiYing
type chaoJiYing struct {
	c  *http.Client
	tr *http.Transport

	httpsProxy string        // HTTPS proxy
	timeout    time.Duration // Timeout (seconds)

	user   string // User account
	pass   string // User password
	pass2  string // MD5 value of user password (32-bit lower case)
	softId string // Software ID: in the user center, the software ID can be generated
}

// NewChaoJiYing
func NewChaoJiYing() ChaoJiYing {
	// When using HTTPS, set no authentication
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	var cjy ChaoJiYing = &chaoJiYing{
		c:      &http.Client{Transport: tr},
		tr:     tr,
		user:   os.Getenv("CHAOJIYING_USER"),
		pass:   os.Getenv("CHAOJIYING_PASS"),
		pass2:  os.Getenv("CHAOJIYING_PASS2"),
		softId: os.Getenv("CHAOJIYING_SOFT_ID"),
	}
	cjy.SetTimeout(60) // Default 60 second timeout
	return cjy
}

// SetHttpsProxy Set Https Proxy
func (c *chaoJiYing) SetHttpsProxy(u string) {
	proxy, _ := url.Parse(u)
	c.tr.Proxy = http.ProxyURL(proxy)
	c.c.Transport = c.tr
}

// SetTimeout Set timeout (seconds)
func (c *chaoJiYing) SetTimeout(t time.Duration) {
	s := t * time.Second
	c.timeout = s
	c.c.Timeout = s
}

// SetUser Set User account
func (c *chaoJiYing) SetUser(user string) {
	c.user = user
}

// SetPass Set User password
func (c *chaoJiYing) SetPass(pass string) {
	c.pass = pass
}

// SetPass2 Set DM5 value of user password (32-bit lower case)
func (c *chaoJiYing) SetPass2(pass2 string) {
	c.pass2 = pass2
}

// SetSoftId Set Software ID: in the user center, the software ID can be generated
func (c *chaoJiYing) SetSoftId(softId string) {
	c.softId = softId
}

// UserInfo Get User info
func (c *chaoJiYing) UserInfo() (userInfo *userInfoResp, err error) {
	var req *http.Request
	var resp *http.Response
	var body []byte
	params := url.Values{}
	params.Add("user", c.user)
	if c.pass != "" {
		params.Add("pass", c.pass)
	} else {
		params.Add("pass2", c.pass2)
	}
	req, err = http.NewRequest(
		"POST",
		"http://upload.chaojiying.net/Upload/GetScore.php",
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0)")
	req.Header.Set("Connection", "Keep-Alive")

	resp, err = c.c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &userInfo); err != nil {
		return
	}
	if userInfo.ErrNo != 0 {
		err = errors.New(userInfo.ErrStr)
		return
	}
	return userInfo, nil
}

// IdentifyPic Identify pictures
func (c *chaoJiYing) IdentifyPic(codeType int, minLen int, imgBase64 string) (identifyPic *identifyPicResp, err error) {
	var req *http.Request
	var resp *http.Response
	var body []byte
	params := url.Values{}
	params.Add("user", c.user)
	if c.pass != "" {
		params.Add("pass", c.pass)
	} else {
		params.Add("pass2", c.pass2)
	}
	params.Add("softid", c.softId)
	params.Add("codetype", strconv.Itoa(codeType))
	params.Add("len_min", strconv.Itoa(minLen))
	params.Add("file_base64", imgBase64)

	req, err = http.NewRequest(
		"POST",
		"http://upload.chaojiying.net/Upload/Processing.php",
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0)")
	req.Header.Set("Connection", "Keep-Alive")

	resp, err = c.c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &identifyPic); err != nil {
		return
	}
	if identifyPic.ErrNo != 0 {
		err = errors.New(identifyPic.ErrStr)
		return
	}
	return identifyPic, nil
}

// ReportError Error report and return to score
func (c *chaoJiYing) ReportError(picId string) (reportError *reportErrorResp, err error) {
	var req *http.Request
	var resp *http.Response
	var body []byte
	params := url.Values{}
	params.Add("user", c.user)
	if c.pass != "" {
		params.Add("pass", c.pass)
	} else {
		params.Add("pass2", c.pass2)
	}
	params.Add("softid", c.softId)
	params.Add("id", picId)

	req, err = http.NewRequest(
		"POST",
		"http://upload.chaojiying.net/Upload/ReportError.php",
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0)")
	req.Header.Set("Connection", "Keep-Alive")

	resp, err = c.c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &reportError); err != nil {
		return
	}
	if reportError.ErrNo != 0 {
		err = errors.New(reportError.ErrStr)
		return
	}
	return reportError, nil
}

// identifyPicResp
type identifyPicResp struct {
	ErrNo  int    `json:"err_no"`
	ErrStr string `json:"err_str"`
	PicId  string `json:"pic_id"`
	PicStr string `json:"pic_str"`
	Md5    string `json:"md5"`
}

// reportErrorResp
type reportErrorResp struct {
	ErrNo  int    `json:"err_no"`
	ErrStr string `json:"err_str"`
}

// scoreResp
type userInfoResp struct {
	ErrNo     int    `json:"err_no"`
	ErrStr    string `json:"err_str"`
	TiFen     int    `json:"tifen"`
	TiFenLock int    `json:"tifen_lock"`
}
