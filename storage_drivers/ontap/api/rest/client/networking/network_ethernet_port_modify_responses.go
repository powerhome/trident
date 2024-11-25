// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NetworkEthernetPortModifyReader is a Reader for the NetworkEthernetPortModify structure.
type NetworkEthernetPortModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetPortModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetPortModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetPortModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetPortModifyOK creates a NetworkEthernetPortModifyOK with default headers values
func NewNetworkEthernetPortModifyOK() *NetworkEthernetPortModifyOK {
	return &NetworkEthernetPortModifyOK{}
}

/*
NetworkEthernetPortModifyOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetPortModifyOK struct {
}

// IsSuccess returns true when this network ethernet port modify o k response has a 2xx status code
func (o *NetworkEthernetPortModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet port modify o k response has a 3xx status code
func (o *NetworkEthernetPortModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet port modify o k response has a 4xx status code
func (o *NetworkEthernetPortModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet port modify o k response has a 5xx status code
func (o *NetworkEthernetPortModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet port modify o k response a status code equal to that given
func (o *NetworkEthernetPortModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the network ethernet port modify o k response
func (o *NetworkEthernetPortModifyOK) Code() int {
	return 200
}

func (o *NetworkEthernetPortModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ethernet/ports/{uuid}][%d] networkEthernetPortModifyOK", 200)
}

func (o *NetworkEthernetPortModifyOK) String() string {
	return fmt.Sprintf("[PATCH /network/ethernet/ports/{uuid}][%d] networkEthernetPortModifyOK", 200)
}

func (o *NetworkEthernetPortModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetPortModifyDefault creates a NetworkEthernetPortModifyDefault with default headers values
func NewNetworkEthernetPortModifyDefault(code int) *NetworkEthernetPortModifyDefault {
	return &NetworkEthernetPortModifyDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetPortModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1376361 | Port is already a member of a LAG. |
| 1376369 | Cannot add port to the interface group because the port hosts one or more VLANs. |
| 1376488 | Disabling the last operational cluster port on a node is not allowed. |
| 1376492 | Cannot change the MTU of a VLAN to be greater than the MTU of the port hosting it. |
| 1377563 | Port is already a member of a LAG. |
| 1377608 | Port cannot be used because it is currently the home port or current port of an interface. |
| 1966288 | Disabling the cluster ports can only be done on the local node. |
| 1967087 | The specified broadcast domain UUID is not valid. |
| 1967088 | The specified broadcast domain name does not exist in the specified IPspace. |
| 1967089 | The specified broadcast domain UUID, name and IPspace name do not match. |
| 1967094 | The specified LAG member port UUID is not valid. |
| 1967095 | The specified LAG member port name and node name combination is not valid. |
| 1967096 | The specified node does not match the specified LAG member port node. |
| 1967097 | The specified LAG member ports UUID, name, and node name do not match. |
| 1967126 | A LAG requires at least one member port. |
| 1967148 | Failure to remove port from broadcast domain. |
| 1967149 | Failure to add port to broadcast domain. |
| 1967184 | The reachability parameter cannot be patched in the same request as other parameters that might affect the target port's reachability status. |
| 1967185 | The port cannot be repaired because the port is deemed as non-repairable. |
| 1967186 | Invalid value for the reachability parameter. |
| 1967195 | Missing or incomplete name retrieval for specified port UUID. |
| 1967580 | This command is not supported as the effective cluster version is earlier than 9.8. |
| 1967582 | The reachability parameter is not supported on this cluster. |
| 53216932 | Failed to determine whether newly introduced PFC flow control is supported. |
| 53280899 | Not all nodes support enabling the PFC feature. |
| 53280900 | No PFC queues specified. |
| 53280901 | Failed to modify PFC. |
| 53280902 | Cluster ports does not support PFC flow control. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type NetworkEthernetPortModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this network ethernet port modify default response has a 2xx status code
func (o *NetworkEthernetPortModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet port modify default response has a 3xx status code
func (o *NetworkEthernetPortModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet port modify default response has a 4xx status code
func (o *NetworkEthernetPortModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet port modify default response has a 5xx status code
func (o *NetworkEthernetPortModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet port modify default response a status code equal to that given
func (o *NetworkEthernetPortModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the network ethernet port modify default response
func (o *NetworkEthernetPortModifyDefault) Code() int {
	return o._statusCode
}

func (o *NetworkEthernetPortModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/ports/{uuid}][%d] network_ethernet_port_modify default %s", o._statusCode, payload)
}

func (o *NetworkEthernetPortModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /network/ethernet/ports/{uuid}][%d] network_ethernet_port_modify default %s", o._statusCode, payload)
}

func (o *NetworkEthernetPortModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetPortModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
