package tt_practice

import (
	"github.com/go-errors/errors"
	"github.com/hashicorp/go-multierror"
)

func TtErrorHandle() error {

	var result error

	if err := step1(); err != nil {
		result = multierror.Append(result, err)
	}
	if err := step2(); err != nil {
		result = multierror.Append(result, err)
	}

	return result
}



func step1() error{
	var Crashed = errors.Errorf("step1")
	return errors.New(Crashed)

}

func step2() error{
	var Crashed = errors.Errorf("step2")
	return errors.New(Crashed)

}
