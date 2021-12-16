<img src="./qr-payment.png" alt="QR code for payment" align="right">

# Payment QR code for Go

[![Build Status](https://travis-ci.com/tuhochi/qrpay.svg?branch=master)](https://travis-ci.com/tuhochi/qrpay)
[![codecov](https://codecov.io/gh/tuhochi/qrpay/branch/master/graph/badge.svg)](https://codecov.io/gh/tuhochi/qrpay)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuhochi/qrpay)](https://goreportcard.com/report/github.com/tuhochi/qrpay)
[![Maintainability](https://api.codeclimate.com/v1/badges/6b488832f724e32e0c6e/maintainability)](https://codeclimate.com/github/tuhochi/qrpay/maintainability)
[![CodeScene Code Health](https://codescene.io/projects/14391/status-badges/code-health)](https://codescene.io/projects/14391)

Golang library for creating QR codes for payments.

[Short Payment Descriptor](https://en.wikipedia.org/wiki/Short_Payment_Descriptor) format and
[EPC QR Code](https://en.wikipedia.org/wiki/EPC_QR_code) (SEPA) format is supported.

## Installation

    go get -u github.com/tuhochi/qrpay

## Usage

### Generating QR code image for Short Payment Descriptor format

```Go
import "github.com/tuhochi/qrpay"

p := qrpay.NewSpaydPayment()
p.SetIBAN("CZ5855000000001265098001")
p.SetAmount("10.8")
p.SetDate(time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC))
p.SetMessage("M")
p.SetRecipientName("go")
p.SetNofificationType('E')
p.SetNotificationValue("daniel@milde.cz")
p.SetExtendedAttribute("vs", "1234567890")

qrpay.SaveQRCodeImageToFile(p, "qr-payment.png")
```

### Generating QR code image for EPC QR Code

```Go
import "github.com/tuhochi/qrpay"

p := qrpay.NewEpcPayment()
p.SetIBAN("CZ5855000000001265098001")
p.SetAmount("10.8")
p.SetMessage("M")
p.SetRecipientName("go")

qrpay.SaveQRCodeImageToFile(p, "qr-payment.png")
```

QR code image encoding uses [skip2/go-qrcode](https://github.com/skip2/go-qrcode).

### Getting QR code content for Short Payment Descriptor format

```Go
import "github.com/tuhochi/qrpay"

p := qrpay.NewSpaydPayment()
p.SetIBAN("CZ5855000000001265098001")
p.SetAmount("108")

fmt.Println(qrpay.GenerateString())
// Output: SPD*1.0*ACC:CZ5855000000001265098001*AM:108*
```