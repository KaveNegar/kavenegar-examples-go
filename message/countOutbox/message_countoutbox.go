package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.New(" your apikey ")

	//Message.CountOutbox
	startdate := time.Now().Add(time.Duration(24) * time.Hour)
	enddate := time.Now()
	status := kavenegar.Type_MessageStatus_Canceled
	if res, err := api.Message.CountOutbox(startdate, enddate, status); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Sumcount 	= ", res.Sumcount)
		fmt.Println("Startdate 	= ", res.Startdate)
		fmt.Println("Enddate 	= ", res.Enddate)
		//..
	}

	//Message.CreateCountOutbox
	v := url.Values{}
	v.Set("startdate", kavenegar.ToUnix(time.Now().Add(time.Duration(24)*time.Hour)))
	v.Set("enddate", kavenegar.ToUnix(time.Now()))
	//v.Set("status",MessageStatusType.Canceled.String)
	if res, err := api.Message.CreateCountOutbox(v); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Sumcount 	= ", res.Sumcount)
		fmt.Println("Startdate 	= ", res.Startdate)
		fmt.Println("Enddate 	= ", res.Enddate)
		//..
	}
}
