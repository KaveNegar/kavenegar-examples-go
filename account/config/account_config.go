package main

import (
	"fmt"
	"net/url"

	"github.com/kavenegar/kavenegar-go"
)

func main() {
	api := kavenegar.New(" your apikey ")
	params := &kavenegar.AccountConfigParam{
		Apilogs:     kavenegar.Type_AccountAPILog_Enabled,
		Dailyreport: kavenegar.Type_AccountDailyReport_Enabled,
		Debugmode:   kavenegar.Type_AccountDebugMode_Enabled,
		//Defaultsender : "",
		//Mincreditalarm :" ",
		Resendfailed: kavenegar.Type_AccountResendFailed_Enabled,
	}

	if res, err := api.Account.Config(params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Apilogs 	= ", res.Apilogs)
		fmt.Println("Dailyreport 	= ", res.Dailyreport)
		fmt.Println("Debugmode 	= ", res.Debugmode)
		fmt.Println("Defaultsender 	= ", res.Defaultsender)
		fmt.Println("Mincreditalarm 	= ", res.Mincreditalarm)
		fmt.Println("Resendfailed 	= ", res.Resendfailed)
		//...

	}

	//Account.CreateAccountConfig
	v := url.Values{}
	v.Set("apilogs", kavenegar.Type_AccountAPILog_Enabled.String())
	v.Set("dailyreport", kavenegar.Type_AccountDailyReport_Enabled.String())
	v.Set("debugmode", kavenegar.Type_AccountDebugMode_Enabled.String())
	//v.Set("defaultsender", "")
	//v.Set("mincreditalarm", "")
	v.Set("resendfailed", kavenegar.Type_AccountResendFailed_Enabled.String())
	if res, err := api.Account.CreateConfig(v); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Apilogs 	= ", res.Apilogs)
		fmt.Println("Dailyreport 	= ", res.Dailyreport)
		fmt.Println("Debugmode 	= ", res.Debugmode)
		fmt.Println("Defaultsender 	= ", res.Defaultsender)
		fmt.Println("Mincreditalarm 	= ", res.Mincreditalarm)
		fmt.Println("Resendfailed 	= ", res.Resendfailed)
		//...
	}

}
