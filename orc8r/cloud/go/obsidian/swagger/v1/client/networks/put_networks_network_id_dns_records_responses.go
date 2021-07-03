// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutNetworksNetworkIDDNSRecordsReader is a Reader for the PutNetworksNetworkIDDNSRecords structure.
type PutNetworksNetworkIDDNSRecordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDDNSRecordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDDNSRecordsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDDNSRecordsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDDNSRecordsNoContent creates a PutNetworksNetworkIDDNSRecordsNoContent with default headers values
func NewPutNetworksNetworkIDDNSRecordsNoContent() *PutNetworksNetworkIDDNSRecordsNoContent {
	return &PutNetworksNetworkIDDNSRecordsNoContent{}
}

/*PutNetworksNetworkIDDNSRecordsNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDDNSRecordsNoContent struct {
}

func (o *PutNetworksNetworkIDDNSRecordsNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/dns/records][%d] putNetworksNetworkIdDnsRecordsNoContent ", 204)
}

func (o *PutNetworksNetworkIDDNSRecordsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDDNSRecordsDefault creates a PutNetworksNetworkIDDNSRecordsDefault with default headers values
func NewPutNetworksNetworkIDDNSRecordsDefault(code int) *PutNetworksNetworkIDDNSRecordsDefault {
	return &PutNetworksNetworkIDDNSRecordsDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDDNSRecordsDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDDNSRecordsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID DNS records default response
func (o *PutNetworksNetworkIDDNSRecordsDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDDNSRecordsDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/dns/records][%d] PutNetworksNetworkIDDNSRecords default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDDNSRecordsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDDNSRecordsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}