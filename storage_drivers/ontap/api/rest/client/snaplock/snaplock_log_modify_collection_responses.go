// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockLogModifyCollectionReader is a Reader for the SnaplockLogModifyCollection structure.
type SnaplockLogModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLogModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLogModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnaplockLogModifyCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLogModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLogModifyCollectionOK creates a SnaplockLogModifyCollectionOK with default headers values
func NewSnaplockLogModifyCollectionOK() *SnaplockLogModifyCollectionOK {
	return &SnaplockLogModifyCollectionOK{}
}

/*
SnaplockLogModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLogModifyCollectionOK struct {
	Payload *models.SnaplockLogJobLinkResponse
}

// IsSuccess returns true when this snaplock log modify collection o k response has a 2xx status code
func (o *SnaplockLogModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log modify collection o k response has a 3xx status code
func (o *SnaplockLogModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log modify collection o k response has a 4xx status code
func (o *SnaplockLogModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log modify collection o k response has a 5xx status code
func (o *SnaplockLogModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log modify collection o k response a status code equal to that given
func (o *SnaplockLogModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snaplock log modify collection o k response
func (o *SnaplockLogModifyCollectionOK) Code() int {
	return 200
}

func (o *SnaplockLogModifyCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs][%d] snaplockLogModifyCollectionOK %s", 200, payload)
}

func (o *SnaplockLogModifyCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs][%d] snaplockLogModifyCollectionOK %s", 200, payload)
}

func (o *SnaplockLogModifyCollectionOK) GetPayload() *models.SnaplockLogJobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLogJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogModifyCollectionAccepted creates a SnaplockLogModifyCollectionAccepted with default headers values
func NewSnaplockLogModifyCollectionAccepted() *SnaplockLogModifyCollectionAccepted {
	return &SnaplockLogModifyCollectionAccepted{}
}

/*
SnaplockLogModifyCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnaplockLogModifyCollectionAccepted struct {
	Payload *models.SnaplockLogJobLinkResponse
}

// IsSuccess returns true when this snaplock log modify collection accepted response has a 2xx status code
func (o *SnaplockLogModifyCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log modify collection accepted response has a 3xx status code
func (o *SnaplockLogModifyCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log modify collection accepted response has a 4xx status code
func (o *SnaplockLogModifyCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log modify collection accepted response has a 5xx status code
func (o *SnaplockLogModifyCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log modify collection accepted response a status code equal to that given
func (o *SnaplockLogModifyCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snaplock log modify collection accepted response
func (o *SnaplockLogModifyCollectionAccepted) Code() int {
	return 202
}

func (o *SnaplockLogModifyCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs][%d] snaplockLogModifyCollectionAccepted %s", 202, payload)
}

func (o *SnaplockLogModifyCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs][%d] snaplockLogModifyCollectionAccepted %s", 202, payload)
}

func (o *SnaplockLogModifyCollectionAccepted) GetPayload() *models.SnaplockLogJobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogModifyCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLogJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogModifyCollectionDefault creates a SnaplockLogModifyCollectionDefault with default headers values
func NewSnaplockLogModifyCollectionDefault(code int) *SnaplockLogModifyCollectionDefault {
	return &SnaplockLogModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLogModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090344    | If log_volume is specified, then log_archive must not be specified  |
| 14090345    | If log_archive.base_name is specified, then log_archive.archive must also be specified  |
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
*/
type SnaplockLogModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock log modify collection default response has a 2xx status code
func (o *SnaplockLogModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock log modify collection default response has a 3xx status code
func (o *SnaplockLogModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock log modify collection default response has a 4xx status code
func (o *SnaplockLogModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock log modify collection default response has a 5xx status code
func (o *SnaplockLogModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock log modify collection default response a status code equal to that given
func (o *SnaplockLogModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock log modify collection default response
func (o *SnaplockLogModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLogModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs][%d] snaplock_log_modify_collection default %s", o._statusCode, payload)
}

func (o *SnaplockLogModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/snaplock/audit-logs][%d] snaplock_log_modify_collection default %s", o._statusCode, payload)
}

func (o *SnaplockLogModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLogModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SnaplockLogModifyCollectionBody snaplock log modify collection body
swagger:model SnaplockLogModifyCollectionBody
*/
type SnaplockLogModifyCollectionBody struct {

	// links
	Links *models.SnaplockLogInlineLinks `json:"_links,omitempty"`

	// log archive
	LogArchive *models.SnaplockLogInlineLogArchive `json:"log_archive,omitempty"`

	// log volume
	LogVolume *models.SnaplockLogVolume `json:"log_volume,omitempty"`

	// snaplock log inline log files
	// Read Only: true
	SnaplockLogInlineLogFiles []*models.SnaplockLogFile `json:"log_files,omitempty"`

	// snaplock log response inline records
	SnaplockLogResponseInlineRecords []*models.SnaplockLog `json:"records,omitempty"`

	// svm
	Svm *models.SnaplockLogInlineSvm `json:"svm,omitempty"`
}

// Validate validates this snaplock log modify collection body
func (o *SnaplockLogModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogArchive(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLogVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSnaplockLogInlineLogFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSnaplockLogResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogModifyCollectionBody) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) validateLogArchive(formats strfmt.Registry) error {
	if swag.IsZero(o.LogArchive) { // not required
		return nil
	}

	if o.LogArchive != nil {
		if err := o.LogArchive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_archive")
			}
			return err
		}
	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) validateLogVolume(formats strfmt.Registry) error {
	if swag.IsZero(o.LogVolume) { // not required
		return nil
	}

	if o.LogVolume != nil {
		if err := o.LogVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_volume")
			}
			return err
		}
	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) validateSnaplockLogInlineLogFiles(formats strfmt.Registry) error {
	if swag.IsZero(o.SnaplockLogInlineLogFiles) { // not required
		return nil
	}

	for i := 0; i < len(o.SnaplockLogInlineLogFiles); i++ {
		if swag.IsZero(o.SnaplockLogInlineLogFiles[i]) { // not required
			continue
		}

		if o.SnaplockLogInlineLogFiles[i] != nil {
			if err := o.SnaplockLogInlineLogFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "log_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) validateSnaplockLogResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SnaplockLogResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SnaplockLogResponseInlineRecords); i++ {
		if swag.IsZero(o.SnaplockLogResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SnaplockLogResponseInlineRecords[i] != nil {
			if err := o.SnaplockLogResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(o.Svm) { // not required
		return nil
	}

	if o.Svm != nil {
		if err := o.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log modify collection body based on the context it is used
func (o *SnaplockLogModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLogArchive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLogVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSnaplockLogInlineLogFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSnaplockLogResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogModifyCollectionBody) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) contextValidateLogArchive(ctx context.Context, formats strfmt.Registry) error {

	if o.LogArchive != nil {
		if err := o.LogArchive.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_archive")
			}
			return err
		}
	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) contextValidateLogVolume(ctx context.Context, formats strfmt.Registry) error {

	if o.LogVolume != nil {
		if err := o.LogVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_volume")
			}
			return err
		}
	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) contextValidateSnaplockLogInlineLogFiles(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"log_files", "body", []*models.SnaplockLogFile(o.SnaplockLogInlineLogFiles)); err != nil {
		return err
	}

	for i := 0; i < len(o.SnaplockLogInlineLogFiles); i++ {

		if o.SnaplockLogInlineLogFiles[i] != nil {
			if err := o.SnaplockLogInlineLogFiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "log_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) contextValidateSnaplockLogResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SnaplockLogResponseInlineRecords); i++ {

		if o.SnaplockLogResponseInlineRecords[i] != nil {
			if err := o.SnaplockLogResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SnaplockLogModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if o.Svm != nil {
		if err := o.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SnaplockLogModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SnaplockLogModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res SnaplockLogModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SnaplockLogInlineLinks snaplock log inline links
swagger:model snaplock_log_inline__links
*/
type SnaplockLogInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this snaplock log inline links
func (o *SnaplockLogInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log inline links based on the context it is used
func (o *SnaplockLogInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SnaplockLogInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SnaplockLogInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SnaplockLogInlineLogArchive snaplock log inline log archive
swagger:model snaplock_log_inline_log_archive
*/
type SnaplockLogInlineLogArchive struct {

	// links
	Links *models.SnaplockLogInlineLogArchiveInlineLinks `json:"_links,omitempty"`

	// Archive the specified SnapLock log file for the given base_name, and create a new log file. If base_name is not mentioned, archive all log files.
	Archive *bool `json:"archive,omitempty"`

	// Base name of log archive
	// Enum: ["legal_hold","privileged_delete","system"]
	BaseName *string `json:"base_name,omitempty"`
}

// Validate validates this snaplock log inline log archive
func (o *SnaplockLogInlineLogArchive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBaseName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineLogArchive) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_archive" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

var snaplockLogInlineLogArchiveTypeBaseNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["legal_hold","privileged_delete","system"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snaplockLogInlineLogArchiveTypeBaseNamePropEnum = append(snaplockLogInlineLogArchiveTypeBaseNamePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// snaplock_log_inline_log_archive
	// SnaplockLogInlineLogArchive
	// base_name
	// BaseName
	// legal_hold
	// END DEBUGGING
	// SnaplockLogInlineLogArchiveBaseNameLegalHold captures enum value "legal_hold"
	SnaplockLogInlineLogArchiveBaseNameLegalHold string = "legal_hold"

	// BEGIN DEBUGGING
	// snaplock_log_inline_log_archive
	// SnaplockLogInlineLogArchive
	// base_name
	// BaseName
	// privileged_delete
	// END DEBUGGING
	// SnaplockLogInlineLogArchiveBaseNamePrivilegedDelete captures enum value "privileged_delete"
	SnaplockLogInlineLogArchiveBaseNamePrivilegedDelete string = "privileged_delete"

	// BEGIN DEBUGGING
	// snaplock_log_inline_log_archive
	// SnaplockLogInlineLogArchive
	// base_name
	// BaseName
	// system
	// END DEBUGGING
	// SnaplockLogInlineLogArchiveBaseNameSystem captures enum value "system"
	SnaplockLogInlineLogArchiveBaseNameSystem string = "system"
)

// prop value enum
func (o *SnaplockLogInlineLogArchive) validateBaseNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snaplockLogInlineLogArchiveTypeBaseNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SnaplockLogInlineLogArchive) validateBaseName(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseName) { // not required
		return nil
	}

	// value enum
	if err := o.validateBaseNameEnum("info"+"."+"log_archive"+"."+"base_name", "body", *o.BaseName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snaplock log inline log archive based on the context it is used
func (o *SnaplockLogInlineLogArchive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineLogArchive) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_archive" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SnaplockLogInlineLogArchive) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SnaplockLogInlineLogArchive) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineLogArchive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SnaplockLogInlineLogArchiveInlineLinks snaplock log inline log archive inline links
swagger:model snaplock_log_inline_log_archive_inline__links
*/
type SnaplockLogInlineLogArchiveInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this snaplock log inline log archive inline links
func (o *SnaplockLogInlineLogArchiveInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineLogArchiveInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_archive" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log inline log archive inline links based on the context it is used
func (o *SnaplockLogInlineLogArchiveInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineLogArchiveInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log_archive" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SnaplockLogInlineLogArchiveInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SnaplockLogInlineLogArchiveInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineLogArchiveInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SnaplockLogInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model snaplock_log_inline_svm
*/
type SnaplockLogInlineSvm struct {

	// links
	Links *models.SnaplockLogInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snaplock log inline svm
func (o *SnaplockLogInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log inline svm based on the context it is used
func (o *SnaplockLogInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SnaplockLogInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SnaplockLogInlineSvm) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SnaplockLogInlineSvmInlineLinks snaplock log inline svm inline links
swagger:model snaplock_log_inline_svm_inline__links
*/
type SnaplockLogInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this snaplock log inline svm inline links
func (o *SnaplockLogInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snaplock log inline svm inline links based on the context it is used
func (o *SnaplockLogInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SnaplockLogInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SnaplockLogInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SnaplockLogInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnaplockLogInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
