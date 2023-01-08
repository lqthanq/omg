package email

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"os"

	"encoding/json"

	"github.com/lqthanq/omg-simpler/helper"
)

var Sender = "leecode2512@gmail.com"

var VerifyAccountTemplateHTML = `
	<h3>Dear {{.username}},</h3>
	<a href="{{.URL}}">Click here to active your account</a>
	<p>Thank you</p>
	<h4>OMG Team</h4>
`

type ContentRequestEmail struct {
	From     string `json:"From"`
	To       string `json:"To"`
	Subject  string `json:"Subject"`
	Tag      string `json:"Tag"`
	HTMLBody string `json:"HtmlBody"`
}

func CreateRequestToPostMark(method, endpoint string, data interface{}) (*http.Request, error) {
	var requestBody bytes.Buffer
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		requestBody = *bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, "https://api.postmarkapp.com"+endpoint, &requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Postmark-Server-Token", os.Getenv("EMAIL_API"))
	return req, nil
}

func SendMail(receiver string, subject string, body string, tokens ...map[string]string) (bool, error) {
	if len(receiver) == 0 {
		return false, errors.New("receiver not exist")
	}

	// get body mail
	bodyBuffer := new(bytes.Buffer)
	if len(tokens) == 0 {
		return false, errors.New("token is null")
	}

	b, err := template.New("email_body").Parse(body)
	if err != nil {
		return false, err
	}

	b.Execute(bodyBuffer, tokens[0])
	emailBody := bodyBuffer.String()
	content := ContentRequestEmail{
		From:     Sender,
		To:       receiver,
		Subject:  subject,
		Tag:      "OMG-Active",
		HTMLBody: emailBody,
	}

	// create request
	req, err := CreateRequestToPostMark("POST", "/email", content)
	if err != nil {
		return false, err
	}

	// send mail
	if res, err := helper.NetClient.Do(req); err != nil {
		return false, err
	} else {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			return false, errors.New("send_mail request error")
		}
		return true, nil
	}
}
