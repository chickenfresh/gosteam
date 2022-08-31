package gosteam

import (
	"errors"
	"fmt"
)

var (
	ErrRSA             = errors.New("error generate rsakey")
	ErrBadCredentials  = errors.New("bad credentials")
	ErrNoAuthenticator = errors.New("account doesn't have authenticator")
	ErrNeedTwoFactor   = errors.New("invalid twofactor code")
	ErrUnauthorized    = errors.New("unauthorized")

	errorApiKey = errors.New("invalid apikey")
)

type CaptchaNeededError struct {
	Err        error
	CaptchaGid string
}

func (err CaptchaNeededError) Error() string {
	return err.Err.Error()
}

func newCaptchaNeededError(captchaGid string) CaptchaNeededError {
	return CaptchaNeededError{
		Err:        errors.New("need to solve captcha"),
		CaptchaGid: captchaGid,
	}
}

func errorStatusCode(functionname string, statuscode int) error {
	return errors.New(fmt.Sprintf("%s | %d", functionname, statuscode))
}

func errorText(text string) error {
	return errors.New(text)
}

func wrappedError(text string, err error) error {
	return fmt.Errorf("%s: %w", text, err)
}
