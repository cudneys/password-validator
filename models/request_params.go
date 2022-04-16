package models

import (
	"github.com/cudneys/password-validator/config"
	"github.com/nbutton23/zxcvbn-go"
	"github.com/nbutton23/zxcvbn-go/scoring"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

type RequestParams struct {
	Password string `form:"password", json"password"`
}

func (r *RequestParams) Validate() (scoring.MinEntropyMatch, error) {
	//ret := make(map[string]string)
	var ret scoring.MinEntropyMatch
	_ = passwordvalidator.GetEntropy(config.GetEntropyPassword())

	const minEntropyBits = 60
	if err := passwordvalidator.Validate(r.Password, minEntropyBits); err != nil {
		return ret, err
	}

	ret = zxcvbn.PasswordStrength(r.Password, nil)
	return ret, nil
}
