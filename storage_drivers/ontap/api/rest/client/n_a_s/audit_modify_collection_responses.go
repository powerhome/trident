// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// AuditModifyCollectionReader is a Reader for the AuditModifyCollection structure.
type AuditModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuditModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAuditModifyCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuditModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuditModifyCollectionOK creates a AuditModifyCollectionOK with default headers values
func NewAuditModifyCollectionOK() *AuditModifyCollectionOK {
	return &AuditModifyCollectionOK{}
}

/*
AuditModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type AuditModifyCollectionOK struct {
	Payload *models.AuditJobLinkResponse
}

// IsSuccess returns true when this audit modify collection o k response has a 2xx status code
func (o *AuditModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this audit modify collection o k response has a 3xx status code
func (o *AuditModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit modify collection o k response has a 4xx status code
func (o *AuditModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this audit modify collection o k response has a 5xx status code
func (o *AuditModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this audit modify collection o k response a status code equal to that given
func (o *AuditModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the audit modify collection o k response
func (o *AuditModifyCollectionOK) Code() int {
	return 200
}

func (o *AuditModifyCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/audit][%d] auditModifyCollectionOK %s", 200, payload)
}

func (o *AuditModifyCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/audit][%d] auditModifyCollectionOK %s", 200, payload)
}

func (o *AuditModifyCollectionOK) GetPayload() *models.AuditJobLinkResponse {
	return o.Payload
}

func (o *AuditModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuditModifyCollectionAccepted creates a AuditModifyCollectionAccepted with default headers values
func NewAuditModifyCollectionAccepted() *AuditModifyCollectionAccepted {
	return &AuditModifyCollectionAccepted{}
}

/*
AuditModifyCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AuditModifyCollectionAccepted struct {
	Payload *models.AuditJobLinkResponse
}

// IsSuccess returns true when this audit modify collection accepted response has a 2xx status code
func (o *AuditModifyCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this audit modify collection accepted response has a 3xx status code
func (o *AuditModifyCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this audit modify collection accepted response has a 4xx status code
func (o *AuditModifyCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this audit modify collection accepted response has a 5xx status code
func (o *AuditModifyCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this audit modify collection accepted response a status code equal to that given
func (o *AuditModifyCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the audit modify collection accepted response
func (o *AuditModifyCollectionAccepted) Code() int {
	return 202
}

func (o *AuditModifyCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/audit][%d] auditModifyCollectionAccepted %s", 202, payload)
}

func (o *AuditModifyCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/audit][%d] auditModifyCollectionAccepted %s", 202, payload)
}

func (o *AuditModifyCollectionAccepted) GetPayload() *models.AuditJobLinkResponse {
	return o.Payload
}

func (o *AuditModifyCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuditModifyCollectionDefault creates a AuditModifyCollectionDefault with default headers values
func NewAuditModifyCollectionDefault(code int) *AuditModifyCollectionDefault {
	return &AuditModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	AuditModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262186     | Field "log.retention.duration" cannot be used with field "log.retention.count"   |
| 9699340    | SVM UUID lookup failed                                                           |
| 9699343    | Audit configuration is absent for modification                                   |
| 9699358    | Audit configuration is absent for enabling                                       |
| 9699359    | Audit configuration is already enabled                                           |
| 9699360    | Final consolidation is in progress, audit enable failed                          |
| 9699365    | Enabling of audit configuration failed                                           |
| 9699373    | Audit configuration is absent for disabling                                      |
| 9699374    | Audit configuration is already disabled                                          |
| 9699375    | Disabling of audit configuration failed                                          |
| 9699384    | The specified log_path does not exist                                            |
| 9699385    | The log_path must be a directory                                                 |
| 9699386    | The log_path must be a canonical path in the SVMs namespace                      |
| 9699387    | The log_path cannot be empty                                                     |
| 9699388    | Rotate size must be greater than or equal to 1024 KB                             |
| 9699389    | The log_path must not contain a symbolic link                                    |
| 9699398    | The log_path exceeds a maximum supported length of characters                    |
| 9699399    | The log_path contains an unsupported read-only (DP/LS) volume                    |
| 9699400    | The specified log_path is not a valid destination for SVM                        |
| 9699402    | The log_path contains an unsupported snaplock volume                             |
| 9699403    | The log_path cannot be accessed for validation                                   |
| 9699406    | The log_path validation failed                                                   |
| 9699407    | Additional fields are provided                                                   |
| 9699409    | Failed to enable multiproto.audit.evtxlog.support support capability             |
| 9699410    | Failed to disable multiproto.audit.evtxlog.support support capability            |
| 9699418    | Audit configuration is absent for rotate                                         |
| 9699419    | Failed to rotate audit log                                                       |
| 9699420    | Cannot rotate audit log, auditing is not enabled for this SVM                    |
| 9699428    | All nodes need to run ONTAP 8.3.0 release to audit CIFS logon-logoff events      |
| 9699429    | Failed to enable multiproto.audit.cifslogonlogoff.support support capability     |
| 9699430    | Failed to disable multiproto.audit.cifslogonlogoff.support support capability    |
| 9699431    | All nodes need to run ONTAP 8.3.0 release to audit CAP staging events            |
| 9699432    | Failed to enable multiproto.audit.capstaging.support support capability          |
| 9699433    | Failed to disable multiproto.audit.capstaging.support support capability         |
*/
type AuditModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this audit modify collection default response has a 2xx status code
func (o *AuditModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this audit modify collection default response has a 3xx status code
func (o *AuditModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this audit modify collection default response has a 4xx status code
func (o *AuditModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this audit modify collection default response has a 5xx status code
func (o *AuditModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this audit modify collection default response a status code equal to that given
func (o *AuditModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the audit modify collection default response
func (o *AuditModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *AuditModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/audit][%d] audit_modify_collection default %s", o._statusCode, payload)
}

func (o *AuditModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /protocols/audit][%d] audit_modify_collection default %s", o._statusCode, payload)
}

func (o *AuditModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AuditModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
AuditModifyCollectionBody audit modify collection body
swagger:model AuditModifyCollectionBody
*/
type AuditModifyCollectionBody struct {

	// audit response inline records
	AuditResponseInlineRecords []*models.Audit `json:"records,omitempty"`

	// Indicates if audit logs generation should incur an extra charge.
	// Example: false
	ChargeQos *bool `json:"charge_qos,omitempty"`

	// Specifies whether or not auditing is enabled on the SVM.
	Enabled *bool `json:"enabled,omitempty"`

	// events
	Events *models.AuditInlineEvents `json:"events,omitempty"`

	// Indicates whether there is a strict Guarantee of Auditing
	// Example: false
	Guarantee *bool `json:"guarantee,omitempty"`

	// log
	Log *models.Log `json:"log,omitempty"`

	// The audit log destination path where consolidated audit logs are stored.
	LogPath *string `json:"log_path,omitempty"`

	// svm
	Svm *models.AuditInlineSvm `json:"svm,omitempty"`
}

// Validate validates this audit modify collection body
func (o *AuditModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuditResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLog(formats); err != nil {
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

func (o *AuditModifyCollectionBody) validateAuditResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.AuditResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.AuditResponseInlineRecords); i++ {
		if swag.IsZero(o.AuditResponseInlineRecords[i]) { // not required
			continue
		}

		if o.AuditResponseInlineRecords[i] != nil {
			if err := o.AuditResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AuditModifyCollectionBody) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(o.Events) { // not required
		return nil
	}

	if o.Events != nil {
		if err := o.Events.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "events")
			}
			return err
		}
	}

	return nil
}

func (o *AuditModifyCollectionBody) validateLog(formats strfmt.Registry) error {
	if swag.IsZero(o.Log) { // not required
		return nil
	}

	if o.Log != nil {
		if err := o.Log.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log")
			}
			return err
		}
	}

	return nil
}

func (o *AuditModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this audit modify collection body based on the context it is used
func (o *AuditModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAuditResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLog(ctx, formats); err != nil {
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

func (o *AuditModifyCollectionBody) contextValidateAuditResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AuditResponseInlineRecords); i++ {

		if o.AuditResponseInlineRecords[i] != nil {
			if err := o.AuditResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AuditModifyCollectionBody) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	if o.Events != nil {
		if err := o.Events.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "events")
			}
			return err
		}
	}

	return nil
}

func (o *AuditModifyCollectionBody) contextValidateLog(ctx context.Context, formats strfmt.Registry) error {

	if o.Log != nil {
		if err := o.Log.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "log")
			}
			return err
		}
	}

	return nil
}

func (o *AuditModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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
func (o *AuditModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AuditModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res AuditModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AuditInlineEvents audit inline events
swagger:model audit_inline_events
*/
type AuditInlineEvents struct {

	// Volume file async delete events
	AsyncDelete *bool `json:"async_delete,omitempty"`

	// Audit policy change events
	AuditPolicyChange *bool `json:"audit_policy_change,omitempty"`

	// Authorization policy change events
	AuthorizationPolicy *bool `json:"authorization_policy,omitempty"`

	// Central access policy staging events
	CapStaging *bool `json:"cap_staging,omitempty"`

	// CIFS logon and logoff events
	CifsLogonLogoff *bool `json:"cifs_logon_logoff,omitempty"`

	// File operation events
	FileOperations *bool `json:"file_operations,omitempty"`

	// File share category events
	FileShare *bool `json:"file_share,omitempty"`

	// Local security group management events
	SecurityGroup *bool `json:"security_group,omitempty"`

	// Local user account management events
	UserAccount *bool `json:"user_account,omitempty"`
}

// Validate validates this audit inline events
func (o *AuditInlineEvents) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this audit inline events based on context it is used
func (o *AuditInlineEvents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AuditInlineEvents) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AuditInlineEvents) UnmarshalBinary(b []byte) error {
	var res AuditInlineEvents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AuditInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model audit_inline_svm
*/
type AuditInlineSvm struct {

	// links
	Links *models.AuditInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this audit inline svm
func (o *AuditInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AuditInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this audit inline svm based on the context it is used
func (o *AuditInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AuditInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *AuditInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AuditInlineSvm) UnmarshalBinary(b []byte) error {
	var res AuditInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AuditInlineSvmInlineLinks audit inline svm inline links
swagger:model audit_inline_svm_inline__links
*/
type AuditInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this audit inline svm inline links
func (o *AuditInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AuditInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this audit inline svm inline links based on the context it is used
func (o *AuditInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AuditInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *AuditInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AuditInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res AuditInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
