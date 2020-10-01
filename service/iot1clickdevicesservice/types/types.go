// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Attributes struct {
}

type Device struct {

	// The device type, such as "button".
	Type *string

	// The unique identifier of the device.
	DeviceId *string

	// The user specified attributes associated with the device for an event.
	Attributes *Attributes
}

type DeviceDescription struct {

	// The unique identifier of the device.
	DeviceId *string

	// The ARN of the device.
	Arn *string

	// An array of zero or more elements of DeviceAttribute objects providing user
	// specified device attributes.
	Attributes map[string]*string

	// A Boolean value indicating whether or not the device is enabled.
	Enabled *bool

	// A value between 0 and 1 inclusive, representing the fraction of life remaining
	// for the device.
	RemainingLife *float64

	// The tags currently associated with the AWS IoT 1-Click device.
	Tags map[string]*string

	// The type of the device, such as "button".
	Type *string
}

type DeviceEvent struct {

	// An object representing the device associated with the event.
	Device *Device

	// A serialized JSON object representing the device-type specific event.
	StdEvent *string
}

type DeviceMethod struct {

	// The name of the method applicable to the deviceType.
	MethodName *string

	// The type of the device, such as "button".
	DeviceType *string
}