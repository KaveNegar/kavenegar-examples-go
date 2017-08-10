package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.New(" your apikey ")

	//Message.Send
	sender := ""                 //Sender Line Number(optional)
	receptor := []string{"", ""} //Recipient numbers
	message := "Hello Go!"       //Text message
	params := &kavenegar.MessageSendParam{
		Date:    time.Now().Add(time.Duration(10) * time.Minute),
		LocalID: []string{"1000", "1001"},
		Type:    []kavenegar.MessageSendType{kavenegar.Type_MessageSend_AppMemory,kavenegar.Type_MessageSend_AppMemory},
	}
	if res, err := api.Message.Send(sender, receptor, message, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}

	//Message.CreateSend
	v := url.Values{}
	//v.Set("sender", "")
	v.Set("message", "Hello Go!")
	v.Add("receptor", "")
	v.Add("receptor", "")
	//v.Add("type",kavenegar.Type_MessageSend_AppMemory.String())
	//v.Add("type",kavenegar.Type_MessageSend_AppMemory.String())
	//v.Add("localid","1000")
	//v.Add("localid","1001")
	//t := time.Now().Add(time.Duration(10) * time.Minute)
	//v.Set("date", kavenegar.TimeToUnix(t))
	if res, err := api.Message.CreateSend(v); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}
}
