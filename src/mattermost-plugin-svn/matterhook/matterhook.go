package matterhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (m *Message) AddAttachment(attachment Attachment) *Message {
	m.Attachments = append(m.Attachments, attachment)
	return m
}

func (m *Message) AddAttachments(attachments []Attachment) *Message {
	m.Attachments = append(m.Attachments, attachments...)
	return m
}

func Send(url string, msg Message, accessToken string) error {
	payloadBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if accessToken != "" {
		req.Header.Set("Authorization", accessToken)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("response read err %s", err)
		return err
	}
	fmt.Printf("response data:%v\n", string(respBody))
	return nil
}
