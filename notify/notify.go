package notify

import "cloudtask/libtools/gounits/container"
import "cloudtask/libtools/gounits/httpx"
import "cloudtask/libtools/gounits/logger"

import (
	"context"
	"net"
	"net/http"
	"time"
)

//NotifySender is exported
type NotifySender struct {
	Runtime    string
	Key        string
	IPAddr     string
	CenterHost string
	client     *httpx.HttpClient
	syncQueue  *container.SyncQueue
}

//NewNotifySender is exported
func NewNotifySender(centerHost string, runtime string, key string, ipAddr string) *NotifySender {
	client := httpx.NewClient().
		SetTransport(&http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 90 * time.Second,
			}).DialContext,
			DisableKeepAlives:     false,
			MaxIdleConns:          50,
			MaxIdleConnsPerHost:   50,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   http.DefaultTransport.(*http.Transport).TLSHandshakeTimeout,
			ExpectContinueTimeout: http.DefaultTransport.(*http.Transport).ExpectContinueTimeout,
		})

	notifySender := &NotifySender{
		Runtime:    runtime,
		Key:        key,
		IPAddr:     ipAddr,
		CenterHost: centerHost,
		client:     client,
		syncQueue:  container.NewSyncQueue(),
	}
	go notifySender.doPopLoop()
	return notifySender
}

func (sender *NotifySender) doPopLoop() {
	for {
		value := sender.syncQueue.Pop()
		if value != nil {
			entry := value.(*NotifyEntry)
			switch entry.NotifyType {
			case NOTIFY_MESSAGE:
				go sender.sendMessage(entry.MsgID, entry.Data)
			case NOTIFY_LOG:
				go sender.sendLog(entry.MsgID, entry.Data)
			}
		}
	}
}

func (sender *NotifySender) sendMessage(msgid string, data interface{}) {
	resp, err := sender.client.PostJSON(context.Background(), sender.CenterHost+"/cloudtask/v2/messages", nil, data, nil)
	if err != nil {
		logger.ERROR("[#notify#] message request %s error, %s", msgid, err.Error())
		return
	}

	defer resp.Close()
	statusCode := resp.StatusCode()
	if statusCode >= http.StatusBadRequest {
		logger.ERROR("[#notify#] message request %s failure, %d", msgid, statusCode)
	}
}

func (sender *NotifySender) sendLog(msgid string, data interface{}) {
	resp, err := sender.client.PostJSON(context.Background(), sender.CenterHost+"/cloudtask/v2/logs", nil, data, nil)
	if err != nil {
		logger.ERROR("[#notify#] logs request %s error, %s", msgid, err.Error())
		return
	}

	defer resp.Close()
	statusCode := resp.StatusCode()
	if statusCode >= http.StatusBadRequest {
		logger.ERROR("[#notify#] logs request %s failure, %d", msgid, statusCode)
	}
}
