package utils

/*
export TWILIO_ACCOUNT_SID= <<your SID>> && \
export TWILIO_AUTH_TOKEN= <<your token>>
*/

/*
Chad developers of the day
Shivam Mishra
*/
// check credit limit of a user
// bonus points for using credit card that could be redeem
// paying on time => increase the limit
// when a user spends upto a particular amount in a month, he/she gets a prize
// refer and earn points
// kind of a lottery of three options
import (
	// randomCrypto "crypto/rand"
	"fmt"
	"log"
	randomMath "math/rand"
	"os"
	"strings"
	"udhaar/constants"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/cast"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendmailBySendGrid() bool {
	sendgridKey := os.Getenv("SendgridKey")
	from := mail.NewEmail(os.Getenv("JeremyName"), os.Getenv("JeremyEmail"))
	to := mail.NewEmail("tarun", "thakur.cs.tarun@gmail.com")
	subject := "Sending with Twilio SendGrid is Fun"
	plainTextContent := "and easy to do anywhere, even with Go"
	otp := GenRandom()
	htmlContent := fmt.Sprintf("<strong>OPT: %v</strong>", otp)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(constants.Green + "plz enter the otp" + constants.White)

	var usrInput string
	fmt.Scanln(&usrInput)

	if usrInput == otp {
		return true
	} else {
		return false
	}
}

func IsValidPhoneNumber(input string) (string, error) {
	phoneNumber := strings.ReplaceAll(input, " ", "")
	if len(phoneNumber) == 0 {
		return "", constants.ErrEmptyPhoneNumber
	}

	// +91-8xxxxxxxxxx
	// +1-8xxxxxxxxx
	phoneNumberElements := strings.Split(phoneNumber, "-")

	if len(phoneNumberElements) != 2 {
		fmt.Println("line 70")
		return "", constants.ErrInvalidPhoneNumber
	}

	countryCode := phoneNumberElements[0]
	restPhoneNum := phoneNumberElements[1]

	if rune(countryCode[0]) != rune('+') {
		fmt.Println("line 78")

		return "", constants.ErrInvalidPhoneNumber
	}

	if len(restPhoneNum) != 10 {
		fmt.Println("line 84")

		return "", constants.ErrInvalidPhoneNumber
	}

	phoneNumber = countryCode + restPhoneNum

	return phoneNumber, nil
}

func IsValidName(name string) (string, error) {
	name = strings.ReplaceAll(name, " ", "")
	if len(name) == 0 {
		return "", constants.ErrEmptyName
	}
	return name, nil
}

func IsValidCreditLimit(creditLimitString string) (float64, error) {
	creditLimitString = strings.ReplaceAll(creditLimitString, " ", "")
	if len(creditLimitString) == 0 {
		return 0.0, constants.ErrEmptyCreditLimit
	}

	creditLimitFloat := cast.ToFloat64(creditLimitString)

	if creditLimitFloat <= 0.0 {
		return 0.0, constants.ErrInvalidCreditLimit
	}

	return creditLimitFloat, nil
}

func IsValidDiscount(discountString string) (float64, error) {
	// discount is given in percentage and in numbers, not alphabets
	// it is a valid discount

	if discountString[len(discountString)-1] != '%' {
		return 0.0, constants.ErrInvalidDiscount
	}

	discount := cast.ToFloat64(discountString)

	if discount <= 0 && discount > 100 {
		return 0.0, constants.ErrInvalidDiscount
	}

	return discount, nil
}

func ValidateViaWhatsapp(phoneNumber string) error {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:%v", phoneNumber))
	params.SetFrom("whatsapp:+14155238886")
	otp := cast.ToString(100 + randomMath.Intn(900))

	params.SetBody(fmt.Sprintf("OTP: %v", otp))

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		fmt.Println("Message sent successfully!")
	}
	fmt.Println(constants.Green + "plz enter the otp" + constants.White)

	var usrInput string
	fmt.Scanln(&usrInput)

	if usrInput == otp {
		return nil
	} else {
		return constants.ErrInvalidOTP
	}
}
