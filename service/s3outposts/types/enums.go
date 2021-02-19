// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EndpointStatus string

// Enum values for EndpointStatus
const (
	EndpointStatusPending   EndpointStatus = "PENDING"
	EndpointStatusAvailable EndpointStatus = "AVAILABLE"
)

// Values returns all known values for EndpointStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EndpointStatus) Values() []EndpointStatus {
	return []EndpointStatus{
		"PENDING",
		"AVAILABLE",
	}
}