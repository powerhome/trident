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
	"github.com/go-openapi/validate"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityKeystoreModifyCollectionReader is a Reader for the SecurityKeystoreModifyCollection structure.
type SecurityKeystoreModifyCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeystoreModifyCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityKeystoreModifyCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSecurityKeystoreModifyCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeystoreModifyCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeystoreModifyCollectionOK creates a SecurityKeystoreModifyCollectionOK with default headers values
func NewSecurityKeystoreModifyCollectionOK() *SecurityKeystoreModifyCollectionOK {
	return &SecurityKeystoreModifyCollectionOK{}
}

/*
SecurityKeystoreModifyCollectionOK describes a response with status code 200, with default header values.

OK
*/
type SecurityKeystoreModifyCollectionOK struct {
	Payload *models.SecurityKeystoreJobLinkResponse
}

// IsSuccess returns true when this security keystore modify collection o k response has a 2xx status code
func (o *SecurityKeystoreModifyCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore modify collection o k response has a 3xx status code
func (o *SecurityKeystoreModifyCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore modify collection o k response has a 4xx status code
func (o *SecurityKeystoreModifyCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore modify collection o k response has a 5xx status code
func (o *SecurityKeystoreModifyCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore modify collection o k response a status code equal to that given
func (o *SecurityKeystoreModifyCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security keystore modify collection o k response
func (o *SecurityKeystoreModifyCollectionOK) Code() int {
	return 200
}

func (o *SecurityKeystoreModifyCollectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores][%d] securityKeystoreModifyCollectionOK %s", 200, payload)
}

func (o *SecurityKeystoreModifyCollectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores][%d] securityKeystoreModifyCollectionOK %s", 200, payload)
}

func (o *SecurityKeystoreModifyCollectionOK) GetPayload() *models.SecurityKeystoreJobLinkResponse {
	return o.Payload
}

func (o *SecurityKeystoreModifyCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityKeystoreJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeystoreModifyCollectionAccepted creates a SecurityKeystoreModifyCollectionAccepted with default headers values
func NewSecurityKeystoreModifyCollectionAccepted() *SecurityKeystoreModifyCollectionAccepted {
	return &SecurityKeystoreModifyCollectionAccepted{}
}

/*
SecurityKeystoreModifyCollectionAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SecurityKeystoreModifyCollectionAccepted struct {
	Payload *models.SecurityKeystoreJobLinkResponse
}

// IsSuccess returns true when this security keystore modify collection accepted response has a 2xx status code
func (o *SecurityKeystoreModifyCollectionAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security keystore modify collection accepted response has a 3xx status code
func (o *SecurityKeystoreModifyCollectionAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security keystore modify collection accepted response has a 4xx status code
func (o *SecurityKeystoreModifyCollectionAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this security keystore modify collection accepted response has a 5xx status code
func (o *SecurityKeystoreModifyCollectionAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this security keystore modify collection accepted response a status code equal to that given
func (o *SecurityKeystoreModifyCollectionAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the security keystore modify collection accepted response
func (o *SecurityKeystoreModifyCollectionAccepted) Code() int {
	return 202
}

func (o *SecurityKeystoreModifyCollectionAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores][%d] securityKeystoreModifyCollectionAccepted %s", 202, payload)
}

func (o *SecurityKeystoreModifyCollectionAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores][%d] securityKeystoreModifyCollectionAccepted %s", 202, payload)
}

func (o *SecurityKeystoreModifyCollectionAccepted) GetPayload() *models.SecurityKeystoreJobLinkResponse {
	return o.Payload
}

func (o *SecurityKeystoreModifyCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityKeystoreJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeystoreModifyCollectionDefault creates a SecurityKeystoreModifyCollectionDefault with default headers values
func NewSecurityKeystoreModifyCollectionDefault(code int) *SecurityKeystoreModifyCollectionDefault {
	return &SecurityKeystoreModifyCollectionDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeystoreModifyCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 262155 | This operation requires an effective cluster version of 9.14.0 or later. |
| 65537605 | Failed to establish connectivity with the cloud key management service. |
| 65538908 | The specified keystore configuration UUID either does not exist or corresponds to a keystore configuration that is not supported by this operation. |
| 65538909 | A value for enabled is required. |
| 65538910 | Disabling an enabled configuration through this method is currently not supported. |
| 65539206 | The SVM associated with the supplied keystore UUID already has a keystore configuration enabled. This command does not support migrating from configurations of that keystore type". |
| 65539212 | Cannot switch the enabled keystore configuration when it is not in the 'active' or 'switching' state. |
| 65539513 | An effective cluster version of ONTAP 9.16.1 or later is required to enable an inactive key manager on the admin SVM. |
| 65539514 | This command does not support enabling key manager configurations with the specified keystore type on the admin SVM. |
| 65539515 | Cannot switch keystore types on the admin SVM. The keystore type for the invalid configuration must be OKM and the enabled configuration must be KMIP, or vice versa. |
| 65539518 | Internal error. Cannot find the enabled configuration. |
| 65539520 | Cannot enable the Onboard Key Manager on the admin SVM because an inactive Onboard Key Manager already exists on the admin SVM. |
| 65539534 | Cannot switch admin SVM Key Manager when system root volumes are present. |
| 65539583 | Cannot switch to the Onboard Key Manager when the external key manager has a policy associated with it |
| 65539585 | Cannot enable an external key manager on the admin SVM because an inactive external key manager already exists on the admin SVM. |
| 65539590 | Cannot switch to the Onboard Key Manager if there are more than two NSE-AKs in the cluster. |
| 65539591 | Cannot switch to the Onboard Key Manager if there are fewer than two NSE-AKs in the cluster. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type SecurityKeystoreModifyCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this security keystore modify collection default response has a 2xx status code
func (o *SecurityKeystoreModifyCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security keystore modify collection default response has a 3xx status code
func (o *SecurityKeystoreModifyCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security keystore modify collection default response has a 4xx status code
func (o *SecurityKeystoreModifyCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security keystore modify collection default response has a 5xx status code
func (o *SecurityKeystoreModifyCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security keystore modify collection default response a status code equal to that given
func (o *SecurityKeystoreModifyCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the security keystore modify collection default response
func (o *SecurityKeystoreModifyCollectionDefault) Code() int {
	return o._statusCode
}

func (o *SecurityKeystoreModifyCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores][%d] security_keystore_modify_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreModifyCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /security/key-stores][%d] security_keystore_modify_collection default %s", o._statusCode, payload)
}

func (o *SecurityKeystoreModifyCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeystoreModifyCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecurityKeystoreModifyCollectionBody security keystore modify collection body
swagger:model SecurityKeystoreModifyCollectionBody
*/
type SecurityKeystoreModifyCollectionBody struct {

	// configuration
	Configuration *models.SecurityKeystoreInlineConfiguration `json:"configuration,omitempty"`

	// Indicates whether the configuration is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Indicates whether the keystore is onboard or external. * 'onboard' - Onboard Key Database * 'external' - External Key Database, including KMIP and Cloud Key Management Systems
	//
	// Read Only: true
	// Enum: ["onboard","external"]
	Location *string `json:"location,omitempty"`

	// scope
	Scope *models.NetworkScopeReadonly `json:"scope,omitempty"`

	// security keystore response inline records
	SecurityKeystoreResponseInlineRecords []*models.SecurityKeystore `json:"records,omitempty"`

	// State of the keystore: * 'active' - The key manager is active and serving new and existing keys. * 'mixed' - The key manager has a mixed configuration. New keys can't be created. * 'svm_kek_rekey' - An SVM key encryption key (KEK) rekey is in progress. New keys can't be created. * 'blocked' - The key manager is blocked and cannot serve new and existing keys. * 'switching' - Switching the enabled key manager keystore configuration. Some operations are blocked. * 'initializing' - The key manager is being initialized. All operations are blocked. * 'disabling' - The key manager is being disabled. All operations are blocked.
	//
	// Read Only: true
	// Enum: ["active","mixed","svm_kek_rekey","blocked","switching","initializing","disabling"]
	State *string `json:"state,omitempty"`

	// svm
	Svm *models.SecurityKeystoreInlineSvm `json:"svm,omitempty"`

	// Type of keystore that is configured: * 'okm' - Onboard Key Manager * 'kmip' - External Key Manager * 'akv' - Azure Key Vault Key Management Service * 'gcp' - Google Cloud Platform Key Management Service * 'aws' - Amazon Web Service Key Management Service * 'ikp' - IBM Key Protect Key Management Service
	//
	// Read Only: true
	// Enum: ["okm","kmip","akv","gcp","aws","ikp"]
	Type *string `json:"type,omitempty"`

	// uuid
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this security keystore modify collection body
func (o *SecurityKeystoreModifyCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecurityKeystoreResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) validateConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(o.Configuration) { // not required
		return nil
	}

	if o.Configuration != nil {
		if err := o.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "configuration")
			}
			return err
		}
	}

	return nil
}

var securityKeystoreModifyCollectionBodyTypeLocationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["onboard","external"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityKeystoreModifyCollectionBodyTypeLocationPropEnum = append(securityKeystoreModifyCollectionBodyTypeLocationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// location
	// Location
	// onboard
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyLocationOnboard captures enum value "onboard"
	SecurityKeystoreModifyCollectionBodyLocationOnboard string = "onboard"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// location
	// Location
	// external
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyLocationExternal captures enum value "external"
	SecurityKeystoreModifyCollectionBodyLocationExternal string = "external"
)

// prop value enum
func (o *SecurityKeystoreModifyCollectionBody) validateLocationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityKeystoreModifyCollectionBodyTypeLocationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	// value enum
	if err := o.validateLocationEnum("info"+"."+"location", "body", *o.Location); err != nil {
		return err
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(o.Scope) { // not required
		return nil
	}

	if o.Scope != nil {
		if err := o.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "scope")
			}
			return err
		}
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) validateSecurityKeystoreResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.SecurityKeystoreResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.SecurityKeystoreResponseInlineRecords); i++ {
		if swag.IsZero(o.SecurityKeystoreResponseInlineRecords[i]) { // not required
			continue
		}

		if o.SecurityKeystoreResponseInlineRecords[i] != nil {
			if err := o.SecurityKeystoreResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var securityKeystoreModifyCollectionBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","mixed","svm_kek_rekey","blocked","switching","initializing","disabling"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityKeystoreModifyCollectionBodyTypeStatePropEnum = append(securityKeystoreModifyCollectionBodyTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// state
	// State
	// active
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyStateActive captures enum value "active"
	SecurityKeystoreModifyCollectionBodyStateActive string = "active"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// state
	// State
	// mixed
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyStateMixed captures enum value "mixed"
	SecurityKeystoreModifyCollectionBodyStateMixed string = "mixed"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// state
	// State
	// svm_kek_rekey
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyStateSvmKekRekey captures enum value "svm_kek_rekey"
	SecurityKeystoreModifyCollectionBodyStateSvmKekRekey string = "svm_kek_rekey"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// state
	// State
	// blocked
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyStateBlocked captures enum value "blocked"
	SecurityKeystoreModifyCollectionBodyStateBlocked string = "blocked"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// state
	// State
	// switching
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyStateSwitching captures enum value "switching"
	SecurityKeystoreModifyCollectionBodyStateSwitching string = "switching"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// state
	// State
	// initializing
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyStateInitializing captures enum value "initializing"
	SecurityKeystoreModifyCollectionBodyStateInitializing string = "initializing"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// state
	// State
	// disabling
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyStateDisabling captures enum value "disabling"
	SecurityKeystoreModifyCollectionBodyStateDisabling string = "disabling"
)

// prop value enum
func (o *SecurityKeystoreModifyCollectionBody) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityKeystoreModifyCollectionBodyTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("info"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) validateSvm(formats strfmt.Registry) error {
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

var securityKeystoreModifyCollectionBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["okm","kmip","akv","gcp","aws","ikp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityKeystoreModifyCollectionBodyTypeTypePropEnum = append(securityKeystoreModifyCollectionBodyTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// type
	// Type
	// okm
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyTypeOkm captures enum value "okm"
	SecurityKeystoreModifyCollectionBodyTypeOkm string = "okm"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// type
	// Type
	// kmip
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyTypeKmip captures enum value "kmip"
	SecurityKeystoreModifyCollectionBodyTypeKmip string = "kmip"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// type
	// Type
	// akv
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyTypeAkv captures enum value "akv"
	SecurityKeystoreModifyCollectionBodyTypeAkv string = "akv"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// type
	// Type
	// gcp
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyTypeGcp captures enum value "gcp"
	SecurityKeystoreModifyCollectionBodyTypeGcp string = "gcp"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// type
	// Type
	// aws
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyTypeAws captures enum value "aws"
	SecurityKeystoreModifyCollectionBodyTypeAws string = "aws"

	// BEGIN DEBUGGING
	// SecurityKeystoreModifyCollectionBody
	// SecurityKeystoreModifyCollectionBody
	// type
	// Type
	// ikp
	// END DEBUGGING
	// SecurityKeystoreModifyCollectionBodyTypeIkp captures enum value "ikp"
	SecurityKeystoreModifyCollectionBodyTypeIkp string = "ikp"
)

// prop value enum
func (o *SecurityKeystoreModifyCollectionBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, securityKeystoreModifyCollectionBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("info"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this security keystore modify collection body based on the context it is used
func (o *SecurityKeystoreModifyCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSecurityKeystoreResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if o.Configuration != nil {
		if err := o.Configuration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "configuration")
			}
			return err
		}
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"location", "body", o.Location); err != nil {
		return err
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if o.Scope != nil {
		if err := o.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "scope")
			}
			return err
		}
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) contextValidateSecurityKeystoreResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SecurityKeystoreResponseInlineRecords); i++ {

		if o.SecurityKeystoreResponseInlineRecords[i] != nil {
			if err := o.SecurityKeystoreResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"state", "body", o.State); err != nil {
		return err
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (o *SecurityKeystoreModifyCollectionBody) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

func (o *SecurityKeystoreModifyCollectionBody) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "info"+"."+"uuid", "body", o.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityKeystoreModifyCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeystoreModifyCollectionBody) UnmarshalBinary(b []byte) error {
	var res SecurityKeystoreModifyCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecurityKeystoreInlineConfiguration Security keystore object reference.
swagger:model security_keystore_inline_configuration
*/
type SecurityKeystoreInlineConfiguration struct {

	// links
	Links *SecurityKeystoreInlineConfigurationInlineLinks `json:"_links,omitempty"`

	// Name of the configuration.
	// Example: default
	Name *string `json:"name,omitempty"`

	// Keystore UUID.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563434
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this security keystore inline configuration
func (o *SecurityKeystoreInlineConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineConfiguration) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "configuration" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security keystore inline configuration based on the context it is used
func (o *SecurityKeystoreInlineConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineConfiguration) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if o.Links != nil {
		if err := o.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "configuration" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityKeystoreInlineConfiguration) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeystoreInlineConfiguration) UnmarshalBinary(b []byte) error {
	var res SecurityKeystoreInlineConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecurityKeystoreInlineConfigurationInlineLinks security keystore inline configuration inline links
swagger:model security_keystore_inline_configuration_inline__links
*/
type SecurityKeystoreInlineConfigurationInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this security keystore inline configuration inline links
func (o *SecurityKeystoreInlineConfigurationInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineConfigurationInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "configuration" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security keystore inline configuration inline links based on the context it is used
func (o *SecurityKeystoreInlineConfigurationInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineConfigurationInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if o.Self != nil {
		if err := o.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "configuration" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SecurityKeystoreInlineConfigurationInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeystoreInlineConfigurationInlineLinks) UnmarshalBinary(b []byte) error {
	var res SecurityKeystoreInlineConfigurationInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecurityKeystoreInlineSvm SVM, applies only to SVM-scoped objects.
swagger:model security_keystore_inline_svm
*/
type SecurityKeystoreInlineSvm struct {

	// links
	Links *models.SecurityKeystoreInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM. This field cannot be specified in a PATCH method.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this security keystore inline svm
func (o *SecurityKeystoreInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this security keystore inline svm based on the context it is used
func (o *SecurityKeystoreInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SecurityKeystoreInlineSvm) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeystoreInlineSvm) UnmarshalBinary(b []byte) error {
	var res SecurityKeystoreInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecurityKeystoreInlineSvmInlineLinks security keystore inline svm inline links
swagger:model security_keystore_inline_svm_inline__links
*/
type SecurityKeystoreInlineSvmInlineLinks struct {

	// self
	Self *models.Href `json:"self,omitempty"`
}

// Validate validates this security keystore inline svm inline links
func (o *SecurityKeystoreInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this security keystore inline svm inline links based on the context it is used
func (o *SecurityKeystoreInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SecurityKeystoreInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (o *SecurityKeystoreInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecurityKeystoreInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SecurityKeystoreInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
