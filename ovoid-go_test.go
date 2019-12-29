package ovoid_go

import (
	"context"
	"fmt"
	"testing"
)

const (
	phoneNumber = "085643553402"
)

func Initiate() {
	Init(Config{
		AppID:        "C7UMRSMFRZ46D9GW9IK7",
		AppVersion:   "2.8.0",
		OSName:       "Android",
		OsVersion:    "8.1.0",
		MACAddress:   "d8:8e:35:4d:bd:88",
		BaseEndpoint: "https://api.ovo.id/",
		AWSEndpoint:  "https://apigw01.aws.ovo.id/",
		TransferOVO:  "trf_ovo",
		TransferBank: "trf_other_bank",
	})
}

func TestLogin2FA(t *testing.T)  {
	Initiate()
	ctx := context.Background()
	response, err := login2FA(ctx, phoneNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
