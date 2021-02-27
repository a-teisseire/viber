package viber

import (
	"context"
	"encoding/json"
	"errors"
)

//
//https://chatapi.viber.com/pa/set_webhook
// {
//    "url": "https://my.host.com",
//    "event_types": ["delivered", "seen", "failed", "subscribed", "unsubscribed", "conversation_started"]
// }

// WebhookReq request
type WebhookReq struct {
	URL        string   `json:"url"`
	EventTypes []string `json:"event_types"`
	SendName   bool     `json:"send_name"`
	SendPhoto  bool     `json:"send_photo"`
}

// {
//     "status": 0,
//     "status_message": "ok",
//     "event_types": ["delivered", "seen", "failed", "subscribed",  "unsubscribed", "conversation_started"]
// }

//WebhookResp response
type WebhookResp struct {
	Status        int      `json:"status"`
	StatusMessage string   `json:"status_message"`
	EventTypes    []string `json:"event_types,omitempty"`
}

// WebhookVerify response
type WebhookVerify struct {
	Event        string `json:"event"`
	Timestamp    uint64 `json:"timestamp"`
	MessageToken uint64 `json:"message_token"`
}

// SetWebhook for Viber callbacks
// if eventTypes is nil, all callbacks will be set to webhook
// if eventTypes is empty []string mandatory callbacks will be set
// Mandatory callbacks: "message", "subscribed", "unsubscribed"
// All possible callbacks: "message", "subscribed",  "unsubscribed", "delivered", "seen", "failed", "conversation_started"
func (v *Viber) SetWebhook(ctx context.Context, url string, eventTypes []string) (WebhookResp, error) {
	var resp WebhookResp

	req := WebhookReq{
		URL:        url,
		EventTypes: eventTypes,
		SendName:   true,
		SendPhoto:  true,
	}
	r, err := v.PostData(ctx, "https://chatapi.viber.com/pa/set_webhook", req)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(r, &resp)
	if err != nil {
		return resp, err
	}

	if resp.Status != 0 {
		return resp, errors.New("non-zero status: " + resp.StatusMessage)
	}

	return resp, nil
}
