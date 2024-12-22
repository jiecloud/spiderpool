// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*
GetServiceOK describes a response with status code 200, with default header values.

Success
*/
type GetServiceOK struct {
	Payload []*models.Service
}

// IsSuccess returns true when this get service o k response has a 2xx status code
func (o *GetServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service o k response has a 3xx status code
func (o *GetServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service o k response has a 4xx status code
func (o *GetServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service o k response has a 5xx status code
func (o *GetServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get service o k response a status code equal to that given
func (o *GetServiceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[GET /service][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) String() string {
	return fmt.Sprintf("[GET /service][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) GetPayload() []*models.Service {
	return o.Payload
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
