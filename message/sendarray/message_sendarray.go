package main

import (
	"net/url"
	"fmt"
	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.New(" your apikey here ")	
	//Message.SendArray
	sender := []string{"",""}   //Sender Line Numbers(optional)
	receptor := []string{"", ""} //Recipient numbers
	message := []string{"Hello Go!", "Hello Go!"}  //Text messages
	localID := []string{"", ""}
	Type := []kavenegar.MessageSendType{kavenegar.Type_MessageSend_AppMemory,kavenegar.Type_MessageSend_AppMemory}
	params := &kavenegar.MessageSendParam{
		LocalID:localID,
		Type:Type,
		//...
	}
	if res, err := api.Message.SendArray(sender, receptor, message, params); err != nil {
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

	//Message.CreateSendArray
	v := url.Values{}
	v.Set("sender", kavenegar.ToJson(sender))
	v.Set("receptor", kavenegar.ToJson(receptor))
	v.Set("message", kavenegar.ToJson(message))

	// v.Set("localid", kavenegar.ToJson(params.LocalID))
	// v.Set("type", kavenegar.ToJson(params.Type))
	//...

	if res, err := api.Message.CreateSendArray(v); err != nil {
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
