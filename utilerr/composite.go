package utilerr

import (
	"strconv"
	"strings"
)

type CompositeError struct {
	ErrList []*error
}

func (err *CompositeError) GetRef(i int) *error {
	return err.ErrList[i]
}

func (err *CompositeError) GetError() error {
	for _, e := range err.ErrList {
		if *e != nil {
			return err
		}
	}
	return nil
}

func (err *CompositeError) Error() string {
	errStrings := []string{}
	for i, e := range err.ErrList {
		if *e != nil {
			errStrings = append(errStrings, "("+strconv.Itoa(i)+")"+
				(*e).Error())
		}
	}
	return "composite error: " + strings.Join(errStrings, "; ")
}

func NewCompositeError(length int) *CompositeError {
	errList := []*error{}
	for i := 0; i < length; i++ {
		var newError error
		errList = append(errList, &newError)
	}
	return &CompositeError{
		ErrList: errList,
	}
}
