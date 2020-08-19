package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"github.com/dgryski/dgoogauth"
	"github.com/skip2/go-qrcode"
	"image/png"
	"net/url"
)

/*
  generate base32 secret string
*/
func GenerateSecret() (usersecret string, err error) {
	secret := make([]byte, 10)
	_, errs := rand.Read(secret)
	if errs != nil {
		return usersecret, errs
	}
	secretBase32 := base32.StdEncoding.EncodeToString(secret)
	return secretBase32, nil
}

/*
  generate base64 google mfa qrcode string
*/
func GenerateMFAQrcode(secret string, accountname string) (err error, qrcodebase64 string) {
	issuer := "gva"
	URL, err := url.Parse("otpauth://totp")
	if err != nil {
		return err, ""
	}
	URL.Path += "/" + url.PathEscape(issuer) + ":" + url.PathEscape(accountname)
	params := url.Values{}
	params.Add("secret", secret)
	params.Add("issuer", issuer)
	URL.RawQuery = params.Encode()
	p, errs := qrcode.New(URL.String(), qrcode.Medium)
	img := p.Image(256)
	if errs != nil {
		return errs, ""
	}
	out := new(bytes.Buffer)
	errx := png.Encode(out, img)
	if errx != nil {
		return errx, ""
	}
	return nil, "data:image/png;base64," + base64.StdEncoding.EncodeToString(out.Bytes())
}

/*
  authenticate mfa
*/
func AuthenticateMFA(secret string, authkey string) (err error, auth bool) {
	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
		// UTC:         true,
	}
	val, err := otpc.Authenticate(authkey)
	return err, val
}
