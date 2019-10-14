// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	zacksclientmodels "github/IAD/zacks/internal/app/server/gen/client/zacksclientmodels"
)

// GetTickerHistoryReader is a Reader for the GetTickerHistory structure.
type GetTickerHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTickerHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTickerHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTickerHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTickerHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /{ticker}/history returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetTickerHistoryOK creates a GetTickerHistoryOK with default headers values
func NewGetTickerHistoryOK() *GetTickerHistoryOK {
	return &GetTickerHistoryOK{}
}

/*GetTickerHistoryOK handles this case with default header values.

OK
*/
type GetTickerHistoryOK struct {
	Payload zacksclientmodels.RankCollection
}

func (o *GetTickerHistoryOK) Error() string {
	return fmt.Sprintf("[GET /{ticker}/history][%d] getTickerHistoryOK  %+v", 200, o.Payload)
}

func (o *GetTickerHistoryOK) GetPayload() zacksclientmodels.RankCollection {
	return o.Payload
}

func (o *GetTickerHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTickerHistoryNotFound creates a GetTickerHistoryNotFound with default headers values
func NewGetTickerHistoryNotFound() *GetTickerHistoryNotFound {
	return &GetTickerHistoryNotFound{}
}

/*GetTickerHistoryNotFound handles this case with default header values.

error
*/
type GetTickerHistoryNotFound struct {
}

func (o *GetTickerHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /{ticker}/history][%d] getTickerHistoryNotFound ", 404)
}

func (o *GetTickerHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTickerHistoryInternalServerError creates a GetTickerHistoryInternalServerError with default headers values
func NewGetTickerHistoryInternalServerError() *GetTickerHistoryInternalServerError {
	return &GetTickerHistoryInternalServerError{}
}

/*GetTickerHistoryInternalServerError handles this case with default header values.

error
*/
type GetTickerHistoryInternalServerError struct {
	Payload *zacksclientmodels.Message
}

func (o *GetTickerHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /{ticker}/history][%d] getTickerHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTickerHistoryInternalServerError) GetPayload() *zacksclientmodels.Message {
	return o.Payload
}

func (o *GetTickerHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(zacksclientmodels.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
