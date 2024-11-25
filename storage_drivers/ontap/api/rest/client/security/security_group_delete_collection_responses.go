// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityGroupDeleteCollectionReader is a Reader for the SecurityGroupDeleteCollection structure.
type SecurityGroupDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityGroupDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityGroupDeleteCollectionOK creates a SecurityGroupDeleteCollectionOK with default headers values
func NewSecurityGroupDeleteCollectionOK() *SecurityGroupDeleteCollectionOK {
	return &SecurityGroupDeleteCollectionOK{}
}

/*
SecurityGroupDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SecurityGroupDeleteCollectionOK struct {
}

// IsSuccess returns true when this security group delete collection o k response has a 2xx status code
func (o *SecurityGroupDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security group delete collection o k response has a 3xx status code
func (o *SecurityGroupDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security group delete collection o k response has a 4xx status code
func (o *SecurityGroupDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security group delete collection o k response has a 5xx status code
func (o *SecurityGroupDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security group delete collection o k response a status code equal to that given
func (o *SecurityGroupDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security group delete collection o k response
func (o *SecurityGroupDeleteCollectionOK) Code() int {
	return 200
}

func (o *SecurityGroupDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /security/groups][%d] securityGroupDeleteCollectionOK", 200)
}

func (o *SecurityGroupDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /security/groups][%d] securityGroupDeleteCollectionOK", 200)
}

func (o *SecurityGroupDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecurityGroupDeleteCollectionDefault creates a SecurityGroupDeleteCollectionDefault with default headers values
func NewSecurityGroupDeleteCollectionDefault(code int) *SecurityGroupDeleteCollectionDefault {
	return &SecurityGroupDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	SecurityGroupDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5636244 | The group is part of the role mapping configuration. Delete the mapping first and then delete the group. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityGroupDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security group delete collection default response has a 2xx status code
func (o *SecurityGroupDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security group delete collection default response has a 3xx status code
func (o *SecurityGroupDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security group delete collection default response has a 4xx status code
func (o *SecurityGroupDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security group delete collection default response has a 5xx status code
func (o *SecurityGroupDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security group delete collection default response a status code equal to that given
func (o *SecurityGroupDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security group delete collection default response
func (o *SecurityGroupDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SecurityGroupDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/groups][%d] security_group_delete_collection default %s", o._statusCode, payload)
}

func (o *SecurityGroupDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/groups][%d] security_group_delete_collection default %s", o._statusCode, payload)
}

func (o *SecurityGroupDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityGroupDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecurityGroupDeleteCollectionBody security group delete collection body
swagger:model SecurityGroupDeleteCollectionBody
*/
type SecurityGroupDeleteCollectionBody struct {

	// security group response inline records
	SecurityGroupResponseInlineRecords []*models.SecurityGroup `json:"records,omitempty"`
}

// Validate validates this security group delete collection body
func (o *SecurityGroupDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSecurityGroupResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityGroupDeleteCollectionBody) validateSecurityGroupResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SecurityGroupResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SecurityGroupResponseInlineRecords); i++ {
		if swag.IsZero(o.SecurityGroupResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SecurityGroupResponseInlineRecords[i] != nil {
			if err := o.SecurityGroupResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this security group delete collection body based on the context it is used
func (o *SecurityGroupDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSecurityGroupResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityGroupDeleteCollectionBody) contextValidateSecurityGroupResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SecurityGroupResponseInlineRecords); i++ {

		if o.SecurityGroupResponseInlineRecords[i] != nil {
			if err := o.SecurityGroupResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityGroupDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityGroupDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res SecurityGroupDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
