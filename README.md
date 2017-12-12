# kavenegar-examples-go

## Installation
```
go get github.com/kavenegar/kavenegar-go
```
## Usage
### Send
```golang
package main
import (
	"fmt"
	"net/url"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
	api := kavenegar.New(" your apikey ")
	sender := ""                 
	receptor := []string{"", ""}
	message := "Hello Go!" 
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
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
```
### OTP
```golang
package main
import (
	"fmt"
	"net/url"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
	api := kavenegar.New(" your apikey ")
	receptor := ""
	template := ""
	token := ""
	params := &kavenegar.VerifyLookupParam{
	}
	if res, err := api.Verify.Lookup(receptor, template, token, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("MessageID 	= ", res.MessageID)
		fmt.Println("Status    	= ", res.Status)
		//...
	}

}
```
### Send Bulk
```golang
package main
import (
	"fmt"
	"net/url"
	"github.com/kavenegar/kavenegar-go"
)
func main() {
	api := kavenegar.New(" your apikey here ")	
	res, err := api.Message.SendArray(url.Values{
		"receptor": {"",""},
		"message": {"Hello Go!","Hello Go!"},
		"sender": {"",""},
	})
	if err != nil {
         switch err := err.(type) {
			case *kavenegar.APIError:
				fmt.Println(err.Error())
			case *kavenegar.HTTPError:
				fmt.Println(err.Error())
			default:
				fmt.Println(err.Error())
         }
	}else{
		fmt.Println(res)
	}
}
```
## Contribution

Bug fixes, docs, and enhancements welcome! Please let us know <a href="mailto:support@kavenegar.com?Subject=SDK" target="_top">support@kavenegar.com</a>

<hr> 
<div dir='rtl'>

	
#### راهنما

##### معرفی سرویس کاوه نگار

کاوه نگار یک وب سرویس ارسال و دریافت پیامک و تماس صوتی است که به راحتی میتوانید از آن استفاده نمایید.

##### ساخت حساب کاربری

اگر در وب سرویس کاوه نگار عضو نیستید میتوانید از [لینک عضویت](http://panel.kavenegar.com/client/membership/register) ثبت نام  و اکانت آزمایشی برای تست API دریافت نمایید.

##### مستندات

برای مشاهده اطلاعات کامل مستندات [وب سرویس پیامک](http://kavenegar.com/وب-سرویس-پیامک.html)  به صفحه [مستندات وب سرویس](http://kavenegar.com/rest.html) مراجعه نمایید.

##### راهنمای فارسی

در صورتی که مایل هستید راهنمای فارسی کیت توسعه کاوه نگار را مطالعه کنید به صفحه [کد ارسال پیامک](http://kavenegar.com/sdk.html) مراجعه نمایید.

##### اطالاعات بیشتر
برای مطالعه بیشتر به صفحه معرفی [وب سرویس کاوه نگار](http://kavenegar.com/%D9%88%D8%A8%D8%B3%D8%B1%D9%88%DB%8C%D8%B3-%D9%BE%DB%8C%D8%A7%D9%85%DA%A9.html) مراجعه نمایید .

 اگر در استفاده از کیت های سرویس کاوه نگار مشکلی یا پیشنهادی  داشتید ما را با یک Pull Request  یا  ارسال ایمیل به support@kavenegar.com  خوشحال کنید.
 

</div>
