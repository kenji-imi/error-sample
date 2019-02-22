package error

import (
	"golang.org/x/xerrors"
)

func Wrap1() error {
	return xerrors.Errorf("[wrap1] error on wrap1: %w", wrap2())
}

func wrap2() error {
	return xerrors.Errorf("[wrap2] error on wrap2: %w", (permanentError))
}

func permanentError() error {
	return xerrors.Errorf("permanent error!")
}
