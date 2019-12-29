package ovoid_go

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type (
	LoginResponse struct {
		RefID string `json:"refId"`
	}
)

func login2FA(ctx context.Context, mobilePhone string) (interface{}, error)  {
	request := Requester{
		ExtraURL: "v2.0/api/auth/customer/login2FA",
		Body:     BodyRequest{
			DeviceID:    uuid.NewV4().String(),
			MobilePhone: mobilePhone,
		},
		Headers:  Headers{
			AppID:      cfg.AppID,
			AppVersion: cfg.AppVersion,
			OSName:     cfg.OSName,
		},
	}

	response := &LoginResponse{}
	err := request.Post(ctx, response)
	if err != nil{
		return nil, err
	}

	return response, nil

}
