// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysqlx.proto

/*
Package mysqlx is a generated protocol buffer package.

It is generated from these files:
	mysqlx.proto

It has these top-level messages:
	ClientMessages
	ServerMessages
	Ok
	Error
*/
package mysqlx

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientMessages_Type int32

const (
	ClientMessages_CON_CAPABILITIES_GET       ClientMessages_Type = 1
	ClientMessages_CON_CAPABILITIES_SET       ClientMessages_Type = 2
	ClientMessages_CON_CLOSE                  ClientMessages_Type = 3
	ClientMessages_SESS_AUTHENTICATE_START    ClientMessages_Type = 4
	ClientMessages_SESS_AUTHENTICATE_CONTINUE ClientMessages_Type = 5
	ClientMessages_SESS_RESET                 ClientMessages_Type = 6
	ClientMessages_SESS_CLOSE                 ClientMessages_Type = 7
	ClientMessages_SQL_STMT_EXECUTE           ClientMessages_Type = 12
	ClientMessages_CRUD_FIND                  ClientMessages_Type = 17
	ClientMessages_CRUD_INSERT                ClientMessages_Type = 18
	ClientMessages_CRUD_UPDATE                ClientMessages_Type = 19
	ClientMessages_CRUD_DELETE                ClientMessages_Type = 20
	ClientMessages_EXPECT_OPEN                ClientMessages_Type = 24
	ClientMessages_EXPECT_CLOSE               ClientMessages_Type = 25
	ClientMessages_CRUD_CREATE_VIEW           ClientMessages_Type = 30
	ClientMessages_CRUD_MODIFY_VIEW           ClientMessages_Type = 31
	ClientMessages_CRUD_DROP_VIEW             ClientMessages_Type = 32
)

var ClientMessages_Type_name = map[int32]string{
	1:  "CON_CAPABILITIES_GET",
	2:  "CON_CAPABILITIES_SET",
	3:  "CON_CLOSE",
	4:  "SESS_AUTHENTICATE_START",
	5:  "SESS_AUTHENTICATE_CONTINUE",
	6:  "SESS_RESET",
	7:  "SESS_CLOSE",
	12: "SQL_STMT_EXECUTE",
	17: "CRUD_FIND",
	18: "CRUD_INSERT",
	19: "CRUD_UPDATE",
	20: "CRUD_DELETE",
	24: "EXPECT_OPEN",
	25: "EXPECT_CLOSE",
	30: "CRUD_CREATE_VIEW",
	31: "CRUD_MODIFY_VIEW",
	32: "CRUD_DROP_VIEW",
}
var ClientMessages_Type_value = map[string]int32{
	"CON_CAPABILITIES_GET":       1,
	"CON_CAPABILITIES_SET":       2,
	"CON_CLOSE":                  3,
	"SESS_AUTHENTICATE_START":    4,
	"SESS_AUTHENTICATE_CONTINUE": 5,
	"SESS_RESET":                 6,
	"SESS_CLOSE":                 7,
	"SQL_STMT_EXECUTE":           12,
	"CRUD_FIND":                  17,
	"CRUD_INSERT":                18,
	"CRUD_UPDATE":                19,
	"CRUD_DELETE":                20,
	"EXPECT_OPEN":                24,
	"EXPECT_CLOSE":               25,
	"CRUD_CREATE_VIEW":           30,
	"CRUD_MODIFY_VIEW":           31,
	"CRUD_DROP_VIEW":             32,
}

func (x ClientMessages_Type) Enum() *ClientMessages_Type {
	p := new(ClientMessages_Type)
	*p = x
	return p
}
func (x ClientMessages_Type) String() string {
	return proto.EnumName(ClientMessages_Type_name, int32(x))
}
func (x *ClientMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ClientMessages_Type_value, data, "ClientMessages_Type")
	if err != nil {
		return err
	}
	*x = ClientMessages_Type(value)
	return nil
}
func (ClientMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ServerMessages_Type int32

const (
	ServerMessages_OK                         ServerMessages_Type = 0
	ServerMessages_ERROR                      ServerMessages_Type = 1
	ServerMessages_CONN_CAPABILITIES          ServerMessages_Type = 2
	ServerMessages_SESS_AUTHENTICATE_CONTINUE ServerMessages_Type = 3
	ServerMessages_SESS_AUTHENTICATE_OK       ServerMessages_Type = 4
	// NOTICE has to stay at 11 forever
	ServerMessages_NOTICE                               ServerMessages_Type = 11
	ServerMessages_RESULTSET_COLUMN_META_DATA           ServerMessages_Type = 12
	ServerMessages_RESULTSET_ROW                        ServerMessages_Type = 13
	ServerMessages_RESULTSET_FETCH_DONE                 ServerMessages_Type = 14
	ServerMessages_RESULTSET_FETCH_SUSPENDED            ServerMessages_Type = 15
	ServerMessages_RESULTSET_FETCH_DONE_MORE_RESULTSETS ServerMessages_Type = 16
	ServerMessages_SQL_STMT_EXECUTE_OK                  ServerMessages_Type = 17
	ServerMessages_RESULTSET_FETCH_DONE_MORE_OUT_PARAMS ServerMessages_Type = 18
)

var ServerMessages_Type_name = map[int32]string{
	0:  "OK",
	1:  "ERROR",
	2:  "CONN_CAPABILITIES",
	3:  "SESS_AUTHENTICATE_CONTINUE",
	4:  "SESS_AUTHENTICATE_OK",
	11: "NOTICE",
	12: "RESULTSET_COLUMN_META_DATA",
	13: "RESULTSET_ROW",
	14: "RESULTSET_FETCH_DONE",
	15: "RESULTSET_FETCH_SUSPENDED",
	16: "RESULTSET_FETCH_DONE_MORE_RESULTSETS",
	17: "SQL_STMT_EXECUTE_OK",
	18: "RESULTSET_FETCH_DONE_MORE_OUT_PARAMS",
}
var ServerMessages_Type_value = map[string]int32{
	"OK":                                   0,
	"ERROR":                                1,
	"CONN_CAPABILITIES":                    2,
	"SESS_AUTHENTICATE_CONTINUE":           3,
	"SESS_AUTHENTICATE_OK":                 4,
	"NOTICE":                               11,
	"RESULTSET_COLUMN_META_DATA":           12,
	"RESULTSET_ROW":                        13,
	"RESULTSET_FETCH_DONE":                 14,
	"RESULTSET_FETCH_SUSPENDED":            15,
	"RESULTSET_FETCH_DONE_MORE_RESULTSETS": 16,
	"SQL_STMT_EXECUTE_OK":                  17,
	"RESULTSET_FETCH_DONE_MORE_OUT_PARAMS": 18,
}

func (x ServerMessages_Type) Enum() *ServerMessages_Type {
	p := new(ServerMessages_Type)
	*p = x
	return p
}
func (x ServerMessages_Type) String() string {
	return proto.EnumName(ServerMessages_Type_name, int32(x))
}
func (x *ServerMessages_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServerMessages_Type_value, data, "ServerMessages_Type")
	if err != nil {
		return err
	}
	*x = ServerMessages_Type(value)
	return nil
}
func (ServerMessages_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Error_Severity int32

const (
	Error_ERROR Error_Severity = 0
	Error_FATAL Error_Severity = 1
)

var Error_Severity_name = map[int32]string{
	0: "ERROR",
	1: "FATAL",
}
var Error_Severity_value = map[string]int32{
	"ERROR": 0,
	"FATAL": 1,
}

func (x Error_Severity) Enum() *Error_Severity {
	p := new(Error_Severity)
	*p = x
	return p
}
func (x Error_Severity) String() string {
	return proto.EnumName(Error_Severity_name, int32(x))
}
func (x *Error_Severity) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Error_Severity_value, data, "Error_Severity")
	if err != nil {
		return err
	}
	*x = Error_Severity(value)
	return nil
}
func (Error_Severity) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// IDs of messages that can be sent from client to the server
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ClientMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClientMessages) Reset()                    { *m = ClientMessages{} }
func (m *ClientMessages) String() string            { return proto.CompactTextString(m) }
func (*ClientMessages) ProtoMessage()               {}
func (*ClientMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// IDs of messages that can be sent from server to client
//
// .. note::
//   this message is never sent on the wire. It is only used to let ``protoc``
//
//   * generate constants
//   * check for uniqueness
type ServerMessages struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ServerMessages) Reset()                    { *m = ServerMessages{} }
func (m *ServerMessages) String() string            { return proto.CompactTextString(m) }
func (*ServerMessages) ProtoMessage()               {}
func (*ServerMessages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// generic Ok message
type Ok struct {
	Msg              *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Ok) Reset()                    { *m = Ok{} }
func (m *Ok) String() string            { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()               {}
func (*Ok) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ok) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// generic Error message
//
// A ``severity`` of ``ERROR`` indicates the current message sequence is
// aborted for the given error and the session is ready for more.
//
// In case of a ``FATAL`` error message the client should not expect
// the server to continue handling any further messages and should
// close the connection.
//
// :param severity: severity of the error message
// :param code: error-code
// :param sql_state: SQL state
// :param msg: human readable error message
type Error struct {
	Severity         *Error_Severity `protobuf:"varint,1,opt,name=severity,enum=Mysqlx.Error_Severity,def=0" json:"severity,omitempty"`
	Code             *uint32         `protobuf:"varint,2,req,name=code" json:"code,omitempty"`
	SqlState         *string         `protobuf:"bytes,4,req,name=sql_state,json=sqlState" json:"sql_state,omitempty"`
	Msg              *string         `protobuf:"bytes,3,req,name=msg" json:"msg,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_Error_Severity Error_Severity = Error_ERROR

func (m *Error) GetSeverity() Error_Severity {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return Default_Error_Severity
}

func (m *Error) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Error) GetSqlState() string {
	if m != nil && m.SqlState != nil {
		return *m.SqlState
	}
	return ""
}

func (m *Error) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

var E_ClientMessageId = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*ClientMessages_Type)(nil),
	Field:         100001,
	Name:          "Mysqlx.client_message_id",
	Tag:           "varint,100001,opt,name=client_message_id,json=clientMessageId,enum=Mysqlx.ClientMessages_Type",
	Filename:      "mysqlx.proto",
}

var E_ServerMessageId = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*ServerMessages_Type)(nil),
	Field:         100002,
	Name:          "Mysqlx.server_message_id",
	Tag:           "varint,100002,opt,name=server_message_id,json=serverMessageId,enum=Mysqlx.ServerMessages_Type",
	Filename:      "mysqlx.proto",
}

func init() {
	proto.RegisterType((*ClientMessages)(nil), "Mysqlx.ClientMessages")
	proto.RegisterType((*ServerMessages)(nil), "Mysqlx.ServerMessages")
	proto.RegisterType((*Ok)(nil), "Mysqlx.Ok")
	proto.RegisterType((*Error)(nil), "Mysqlx.Error")
	proto.RegisterEnum("Mysqlx.ClientMessages_Type", ClientMessages_Type_name, ClientMessages_Type_value)
	proto.RegisterEnum("Mysqlx.ServerMessages_Type", ServerMessages_Type_name, ServerMessages_Type_value)
	proto.RegisterEnum("Mysqlx.Error_Severity", Error_Severity_name, Error_Severity_value)
	proto.RegisterExtension(E_ClientMessageId)
	proto.RegisterExtension(E_ServerMessageId)
}

func init() { proto.RegisterFile("mysqlx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x52, 0xe3, 0x46,
	0x10, 0xc6, 0xb2, 0x71, 0x70, 0x63, 0x9b, 0xf1, 0x40, 0x82, 0x81, 0x04, 0x5c, 0xae, 0x1c, 0x7c,
	0x12, 0xa9, 0xdc, 0xc2, 0x4d, 0x48, 0xed, 0xa0, 0xc2, 0xd2, 0x38, 0x33, 0xa3, 0x40, 0x4e, 0x53,
	0xc4, 0x56, 0x5c, 0x26, 0x36, 0x32, 0x92, 0x43, 0xc1, 0x43, 0xa4, 0x6a, 0x5f, 0x61, 0xf7, 0xbe,
	0xcf, 0xb2, 0xcf, 0xb0, 0x7b, 0xdc, 0xf3, 0x3e, 0xc0, 0x96, 0x46, 0xb6, 0xf8, 0xad, 0xdd, 0xda,
	0x5b, 0xf7, 0xf7, 0xb5, 0xbe, 0xee, 0xe9, 0x1f, 0x41, 0x75, 0x7a, 0x97, 0x5c, 0x4f, 0x6e, 0xcd,
	0x59, 0x1c, 0xcd, 0x23, 0x5a, 0xf6, 0xb4, 0xb7, 0xdb, 0x1a, 0x45, 0xd1, 0x68, 0x12, 0x1e, 0x6a,
	0xf4, 0xef, 0xff, 0xfe, 0x39, 0x1c, 0x86, 0xc9, 0x20, 0x1e, 0xcf, 0xe6, 0x51, 0x9c, 0x45, 0xb6,
	0x3f, 0x19, 0x50, 0xb7, 0x27, 0xe3, 0xf0, 0x6a, 0xee, 0x85, 0x49, 0x72, 0x31, 0x0a, 0x93, 0xf6,
	0x7b, 0x03, 0x4a, 0xf2, 0x6e, 0x16, 0xd2, 0x26, 0x6c, 0xd9, 0xcc, 0x57, 0xb6, 0xd5, 0xb7, 0x8e,
	0xdd, 0x9e, 0x2b, 0x5d, 0x14, 0xea, 0x77, 0x94, 0xa4, 0xf0, 0x22, 0x23, 0x50, 0x12, 0x83, 0xd6,
	0xa0, 0xa2, 0x99, 0x1e, 0x13, 0x48, 0x8a, 0x74, 0x0f, 0xb6, 0x05, 0x0a, 0xa1, 0xac, 0x40, 0x9e,
	0xa0, 0x2f, 0x5d, 0xdb, 0x92, 0xa8, 0x84, 0xb4, 0xb8, 0x24, 0x25, 0xba, 0x0f, 0xbb, 0xcf, 0x49,
	0x9b, 0xf9, 0xd2, 0xf5, 0x03, 0x24, 0xab, 0xb4, 0x0e, 0xa0, 0x79, 0x8e, 0xa9, 0x76, 0x39, 0xf7,
	0x33, 0xf1, 0xef, 0xe8, 0x16, 0x10, 0xf1, 0x47, 0x4f, 0x09, 0xe9, 0x49, 0x85, 0xe7, 0x68, 0x07,
	0x12, 0x49, 0x55, 0x57, 0xc0, 0x03, 0x47, 0x75, 0x5d, 0xdf, 0x21, 0x0d, 0xba, 0x01, 0xeb, 0xda,
	0x75, 0x7d, 0x81, 0x5c, 0x12, 0x9a, 0x03, 0x41, 0xdf, 0xb1, 0x24, 0x92, 0xcd, 0x1c, 0x70, 0xb0,
	0x87, 0x12, 0xc9, 0x56, 0x0a, 0xe0, 0x79, 0x1f, 0x6d, 0xa9, 0x58, 0x1f, 0x7d, 0xd2, 0xa4, 0x04,
	0xaa, 0x0b, 0x20, 0x4b, 0xbd, 0x93, 0xa6, 0xd6, 0xdf, 0xd8, 0x1c, 0xd3, 0xa2, 0xff, 0x74, 0xf1,
	0x8c, 0xec, 0xe7, 0xa8, 0xc7, 0x1c, 0xb7, 0xfb, 0x57, 0x86, 0x1e, 0x50, 0x0a, 0xf5, 0x4c, 0x9f,
	0xb3, 0x7e, 0x86, 0xb5, 0xda, 0x1f, 0x0c, 0xa8, 0x8b, 0x30, 0xbe, 0x09, 0xe3, 0xbc, 0xed, 0xef,
	0x96, 0x6d, 0x2f, 0x83, 0xc1, 0x4e, 0xc9, 0x0a, 0xad, 0xc0, 0x2a, 0x72, 0xce, 0x38, 0x29, 0xd0,
	0xef, 0xa1, 0x61, 0x33, 0xff, 0x71, 0xc3, 0x89, 0xf1, 0x95, 0x06, 0x16, 0xd3, 0x31, 0x3d, 0xe7,
	0xd9, 0x29, 0x29, 0x51, 0x80, 0xb2, 0xcf, 0xa4, 0x6b, 0x23, 0x59, 0x4f, 0x55, 0x38, 0x8a, 0xa0,
	0x27, 0x05, 0x4a, 0x65, 0xb3, 0x5e, 0xe0, 0xf9, 0xca, 0x43, 0x69, 0x29, 0xc7, 0x92, 0x16, 0xa9,
	0xd2, 0x06, 0xd4, 0xee, 0x79, 0xce, 0xce, 0x48, 0x2d, 0x15, 0xbe, 0x87, 0xba, 0x28, 0xed, 0x13,
	0xe5, 0x30, 0x1f, 0x49, 0x9d, 0xfe, 0x04, 0x3b, 0x4f, 0x19, 0x11, 0x88, 0x3e, 0xfa, 0x0e, 0x3a,
	0x64, 0x83, 0x76, 0xe0, 0xe7, 0x97, 0x3e, 0x54, 0x1e, 0xe3, 0xa8, 0x72, 0x46, 0x10, 0x42, 0xb7,
	0x61, 0xf3, 0xe9, 0x70, 0xd3, 0xd2, 0x1b, 0x5f, 0x96, 0x60, 0x81, 0x54, 0x7d, 0x8b, 0x5b, 0x9e,
	0x20, 0xb4, 0xfd, 0x23, 0x18, 0xec, 0x5f, 0x4a, 0xa0, 0x38, 0x4d, 0x46, 0xcd, 0x42, 0xab, 0xd0,
	0xa9, 0xf0, 0xd4, 0x3c, 0x2a, 0xbd, 0xfa, 0xf8, 0xcb, 0x4a, 0xfb, 0x6d, 0x01, 0x56, 0x31, 0x8e,
	0xa3, 0x98, 0xfe, 0x06, 0x6b, 0x49, 0x78, 0x13, 0xc6, 0xe3, 0xf9, 0x9d, 0x0e, 0xab, 0xff, 0xfa,
	0x83, 0x99, 0x1d, 0x90, 0xa9, 0x03, 0x4c, 0xb1, 0x60, 0x8f, 0xb2, 0x81, 0xf0, 0x3c, 0x9c, 0x52,
	0x28, 0x0d, 0xa2, 0x61, 0xd8, 0x34, 0x5a, 0x46, 0xa7, 0xc6, 0xb5, 0x4d, 0xf7, 0xa0, 0x92, 0x5c,
	0x4f, 0x54, 0x32, 0xbf, 0x98, 0x87, 0xcd, 0x52, 0xcb, 0xe8, 0x54, 0xf8, 0x5a, 0x72, 0x3d, 0x11,
	0xa9, 0xbf, 0xac, 0xa6, 0xa8, 0xe1, 0xd4, 0x6c, 0xb7, 0x60, 0x6d, 0xa9, 0x7f, 0x3f, 0x72, 0x3d,
	0xfd, 0xae, 0x25, 0xad, 0x1e, 0x29, 0xe8, 0x7a, 0x0b, 0x47, 0x97, 0xd0, 0x18, 0xe8, 0x43, 0x55,
	0xd3, 0x6c, 0x65, 0xd4, 0x78, 0x48, 0x0f, 0xcc, 0xec, 0xc2, 0xcd, 0xe5, 0x85, 0x9b, 0x8b, 0x7d,
	0x62, 0xb3, 0xf9, 0x38, 0xba, 0x4a, 0x9a, 0xaf, 0xff, 0x2f, 0xeb, 0x17, 0xed, 0x2d, 0x5f, 0xf4,
	0xf8, 0xd8, 0xcd, 0x74, 0xe3, 0xf8, 0xc6, 0xe0, 0x21, 0xe8, 0x0e, 0xd3, 0x5c, 0x89, 0xde, 0xce,
	0x6f, 0xca, 0xf5, 0xe6, 0x69, 0xae, 0xc7, 0x1b, 0xbe, 0xc8, 0x95, 0x3c, 0x04, 0xdd, 0xe1, 0xf1,
	0x0e, 0x6c, 0x0f, 0xa2, 0xa9, 0xa9, 0xff, 0x5f, 0xe6, 0xe0, 0xd2, 0xbc, 0xcd, 0xf5, 0x3f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x8f, 0xfa, 0x23, 0xd5, 0x04, 0x00, 0x00,
}
