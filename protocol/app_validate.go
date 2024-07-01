package protocol

import (
	"errors"
	"strings"
)

func (l *Login) Validate() error {
	if strings.EqualFold(l.Username, "") || strings.EqualFold(l.Email, "") {
		return errors.New("parameter is null")
	}
	return nil
}
