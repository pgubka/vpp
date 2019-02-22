// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/tap.api.json

/*
 Package tap is a generated from VPP binary API module 'tap'.

 It contains following objects:
	  4 services
	  8 messages
*/
package tap

import api "git.fd.io/govpp.git/api"
import struc "github.com/lunixbochs/struc"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
type Services interface {
	DumpSwInterfaceTap(*SwInterfaceTapDump) ([]*SwInterfaceTapDetails, error)
	TapConnect(*TapConnect) (*TapConnectReply, error)
	TapDelete(*TapDelete) (*TapDeleteReply, error)
	TapModify(*TapModify) (*TapModifyReply, error)
}

/* Messages */

// SwInterfaceTapDetails represents VPP binary API message 'sw_interface_tap_details':
type SwInterfaceTapDetails struct {
	SwIfIndex uint32
	DevName   []byte `struc:"[64]byte"`
}

func (*SwInterfaceTapDetails) GetMessageName() string {
	return "sw_interface_tap_details"
}
func (*SwInterfaceTapDetails) GetCrcString() string {
	return "76229a57"
}
func (*SwInterfaceTapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceTapDump represents VPP binary API message 'sw_interface_tap_dump':
type SwInterfaceTapDump struct{}

func (*SwInterfaceTapDump) GetMessageName() string {
	return "sw_interface_tap_dump"
}
func (*SwInterfaceTapDump) GetCrcString() string {
	return "51077d14"
}
func (*SwInterfaceTapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapConnect represents VPP binary API message 'tap_connect':
type TapConnect struct {
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
	IP4AddressSet     uint8
	IP4Address        []byte `struc:"[4]byte"`
	IP4MaskWidth      uint8
	IP6AddressSet     uint8
	IP6Address        []byte `struc:"[16]byte"`
	IP6MaskWidth      uint8
	Tag               []byte `struc:"[64]byte"`
}

func (*TapConnect) GetMessageName() string {
	return "tap_connect"
}
func (*TapConnect) GetCrcString() string {
	return "9b9c396f"
}
func (*TapConnect) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapConnectReply represents VPP binary API message 'tap_connect_reply':
type TapConnectReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapConnectReply) GetMessageName() string {
	return "tap_connect_reply"
}
func (*TapConnectReply) GetCrcString() string {
	return "fda5941f"
}
func (*TapConnectReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TapDelete represents VPP binary API message 'tap_delete':
type TapDelete struct {
	SwIfIndex uint32
}

func (*TapDelete) GetMessageName() string {
	return "tap_delete"
}
func (*TapDelete) GetCrcString() string {
	return "529cb13f"
}
func (*TapDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapDeleteReply represents VPP binary API message 'tap_delete_reply':
type TapDeleteReply struct {
	Retval int32
}

func (*TapDeleteReply) GetMessageName() string {
	return "tap_delete_reply"
}
func (*TapDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*TapDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TapModify represents VPP binary API message 'tap_modify':
type TapModify struct {
	SwIfIndex         uint32
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
}

func (*TapModify) GetMessageName() string {
	return "tap_modify"
}
func (*TapModify) GetCrcString() string {
	return "8047ae5c"
}
func (*TapModify) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapModifyReply represents VPP binary API message 'tap_modify_reply':
type TapModifyReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapModifyReply) GetMessageName() string {
	return "tap_modify_reply"
}
func (*TapModifyReply) GetCrcString() string {
	return "fda5941f"
}
func (*TapModifyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func init() {
	api.RegisterMessage((*SwInterfaceTapDetails)(nil), "tap.SwInterfaceTapDetails")
	api.RegisterMessage((*SwInterfaceTapDump)(nil), "tap.SwInterfaceTapDump")
	api.RegisterMessage((*TapConnect)(nil), "tap.TapConnect")
	api.RegisterMessage((*TapConnectReply)(nil), "tap.TapConnectReply")
	api.RegisterMessage((*TapDelete)(nil), "tap.TapDelete")
	api.RegisterMessage((*TapDeleteReply)(nil), "tap.TapDeleteReply")
	api.RegisterMessage((*TapModify)(nil), "tap.TapModify")
	api.RegisterMessage((*TapModifyReply)(nil), "tap.TapModifyReply")
}

var Messages = []api.Message{
	(*SwInterfaceTapDetails)(nil),
	(*SwInterfaceTapDump)(nil),
	(*TapConnect)(nil),
	(*TapConnectReply)(nil),
	(*TapDelete)(nil),
	(*TapDeleteReply)(nil),
	(*TapModify)(nil),
	(*TapModifyReply)(nil),
}
