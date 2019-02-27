/*
Package navprot provides functions for working with different navigation protocols.
Provided functionality includes:
1)Parsing binary packets
2)Generating binary packets
3)Printing packet information in readable format
Currently EGTS (ERA GLONASS Telematics Standard) and NDTP (Navigation Data Transfer Protocol) are supported
*/
package navprot

// NavProtocol is an interface for arbitrary navigation protocol.
type NavProtocol interface {
	Parse([]byte) ([]byte, error)
	Form() []byte
	Print() string
}