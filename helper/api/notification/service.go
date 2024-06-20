package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type serviceNotification struct {
	url    string
	debug  bool
	header struct {
		key string
	}
}

func NewNotification(url string, header string, debug bool) ServiceNotification {
	return &serviceNotification{
		url:   url,
		debug: debug,
		header: struct {
			key string
		}{
			key: header,
		},
	}
}

// SendEmail send email from notification-service
func (b *serviceNotification) SendEmail(param *EmailRequest, token string) (*EmailResponse, error) {
	var respPayload *EmailResponse
	
	urlStr := fmt.Sprintf("%v/%v", b.url, "notification/email")
	reqBody, err := json.Marshal(param)
	if err != nil {
		if b.debug {
			return nil, err
		}
		
		err = fmt.Errorf(fmt.Sprintf("Failed when send email from Notification API"))
		return nil, err
		
	}
	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		if b.debug {
			return nil, err
		}
		
		err = fmt.Errorf(fmt.Sprintf("Failed when send email from Notification API"))
		return nil, err
	}
	
	// req.Header.Set(b.header.key, fmt.Sprintf("%v %v", "Bearer", token))
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if b.debug {
			return nil, err
		}
		
		err = fmt.Errorf(fmt.Sprintf("Failed when send email from Notification API"))
		return nil, err
	}
	
	defer resp.Body.Close()
	
	err = json.NewDecoder(resp.Body).Decode(&respPayload)
	if err != nil {
		if b.debug {
			return nil, err
		}
		
		err = fmt.Errorf(fmt.Sprintf("Failed when decode"))
		return nil, err
	}
	
	if resp.StatusCode == http.StatusOK {
		return respPayload, nil
	}
	
	return nil, fmt.Errorf(fmt.Sprintf("Send email failed"))
}
