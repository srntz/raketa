package rkstr

import "errors"

type Metadata struct {
	maxLength *int
	minLength *int
}

func NewMetadata() *Metadata {
	return &Metadata{
		maxLength: nil,
		minLength: nil,
	}
}

func (m *Metadata) SetMaxLength(value int) (int, error) {
	if value < 1 {
		return 0, errors.New("Invalid option assignment: MAXLENGTH can not be lower than 1")
	}

	if value < *m.minLength {
		return 0, errors.New("Invalid option assignment: MAXLENGTH of type STR can not be lower than MINLENGTH")
	}

	m.maxLength = &value
	return *m.maxLength, nil
}

func (m *Metadata) ResetMaxLength() {
	m.maxLength = nil
}

func (m *Metadata) SetMinLength(value int) (int, error) {
	if value < 0 {
		return 0, errors.New("Invalid option assignment: MINLENGTH can not be lower than 0")
	}

	if value > *m.maxLength {
		return 0, errors.New("Invalid option assignment: MINLENGTH can not be greater than MAXLENGTH")
	}

	m.minLength = &value
	return *m.minLength, nil
}

func (m *Metadata) ResetMinLength(value int) {
	m.minLength = nil
}
