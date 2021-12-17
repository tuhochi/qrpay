package main

import (
	"fmt"
	"time"

	payment "github.com/tuhochi/qrpay"
)

func main() {
	spaydPayment := payment.NewSpaydPayment()
	spaydPayment.SetIBAN("CZ5855000000001265098001")
	spaydPayment.SetAmount("100")
	spaydPayment.SetDate(time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC))
	spaydPayment.SetMessage("M")
	spaydPayment.SetRecipientName("go")
	spaydPayment.SetNofificationType('E')
	spaydPayment.SetNotificationValue("daniel@milde.cz")
	spaydPayment.SetExtendedAttribute("vs", "1234567890")

	if err := payment.SaveQRCodeImageToFile(spaydPayment, "qr-spaydPayment.png"); err != nil {
		fmt.Printf("could not generate payment QR code: %v", err)
	}

	epcPayment := payment.NewEpcPayment()
	epcPayment.SetRecipientName("Franz Mustermänn")
	epcPayment.SetIBAN("DE71110220330123456789")
	epcPayment.SetBIC("BHBLDEHHXXX")
	epcPayment.SetEurAmountCent(1080)
	epcPayment.SetSenderReference("blablablubb ddfdf dfdf")

	if err := payment.SaveQRCodeImageToFile(epcPayment, "qr-epcPayment.png"); err != nil {
		fmt.Printf("could not generate payment QR code: %v", err)
	}

}
