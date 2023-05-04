package Template

import (
	"fmt"
	"testing"
)

func TestGenAndTest(t *testing.T) {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	err := o.genAndSendOTP(4)
	if err != nil {
		return
	}

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	err = o.genAndSendOTP(4)
	if err != nil {
		return
	}
}
