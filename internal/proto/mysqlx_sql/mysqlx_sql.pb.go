// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mysqlx_sql.proto

/*
Package mysqlx_sql is a generated protocol buffer package.

Messages of the MySQL Package

It is generated from these files:
	mysqlx_sql.proto

It has these top-level messages:
	StmtExecute
	StmtExecuteOk
*/
package mysqlx_sql

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/AlekSi/mysqlx/internal/proto/mysqlx"
import Mysqlx_Datatypes "github.com/AlekSi/mysqlx/internal/proto/mysqlx_datatypes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// execute a statement in the given namespace
//
// .. uml::
//
//   client -> server: StmtExecute
//   ... zero or more Resultsets ...
//   server --> client: StmtExecuteOk
//
// Notices:
//   This message may generate a notice containing WARNINGs generated by its execution.
//   This message may generate a notice containing INFO messages generated by its execution.
//
// :param namespace: namespace of the statement to be executed
// :param stmt: statement that shall be executed.
// :param args: values for wildcard replacements
// :param compact_metadata: send only type information for :protobuf:msg:`Mysqlx.Resultset::ColumnMetadata`, skipping names and others
// :returns:
//    * zero or one :protobuf:msg:`Mysqlx.Resultset::` followed by :protobuf:msg:`Mysqlx.Sql::StmtExecuteOk`
type StmtExecute struct {
	Namespace        *string                 `protobuf:"bytes,3,opt,name=namespace,def=sql" json:"namespace,omitempty"`
	Stmt             []byte                  `protobuf:"bytes,1,req,name=stmt" json:"stmt,omitempty"`
	Args             []*Mysqlx_Datatypes.Any `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
	CompactMetadata  *bool                   `protobuf:"varint,4,opt,name=compact_metadata,json=compactMetadata,def=0" json:"compact_metadata,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *StmtExecute) Reset()                    { *m = StmtExecute{} }
func (m *StmtExecute) String() string            { return proto.CompactTextString(m) }
func (*StmtExecute) ProtoMessage()               {}
func (*StmtExecute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_StmtExecute_Namespace string = "sql"
const Default_StmtExecute_CompactMetadata bool = false

func (m *StmtExecute) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return Default_StmtExecute_Namespace
}

func (m *StmtExecute) GetStmt() []byte {
	if m != nil {
		return m.Stmt
	}
	return nil
}

func (m *StmtExecute) GetArgs() []*Mysqlx_Datatypes.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *StmtExecute) GetCompactMetadata() bool {
	if m != nil && m.CompactMetadata != nil {
		return *m.CompactMetadata
	}
	return Default_StmtExecute_CompactMetadata
}

// statement executed successful
type StmtExecuteOk struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StmtExecuteOk) Reset()                    { *m = StmtExecuteOk{} }
func (m *StmtExecuteOk) String() string            { return proto.CompactTextString(m) }
func (*StmtExecuteOk) ProtoMessage()               {}
func (*StmtExecuteOk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*StmtExecute)(nil), "Mysqlx.Sql.StmtExecute")
	proto.RegisterType((*StmtExecuteOk)(nil), "Mysqlx.Sql.StmtExecuteOk")
}

func init() { proto.RegisterFile("mysqlx_sql.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xad, 0x2c, 0x2e,
	0xcc, 0xa9, 0x88, 0x2f, 0x2e, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xf2, 0x05,
	0x8b, 0xe8, 0x05, 0x17, 0xe6, 0x48, 0xf1, 0x40, 0x64, 0x21, 0x32, 0x52, 0x62, 0x50, 0xb5, 0x29,
	0x89, 0x25, 0x89, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x10, 0x71, 0xa5, 0xe5, 0x8c, 0x5c, 0xdc, 0xc1,
	0x25, 0xb9, 0x25, 0xae, 0x15, 0xa9, 0xc9, 0xa5, 0x25, 0xa9, 0x42, 0x8a, 0x5c, 0x9c, 0x79, 0x89,
	0xb9, 0xa9, 0xc5, 0x05, 0x89, 0xc9, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x56, 0xcc, 0xc5,
	0x85, 0x39, 0x41, 0x08, 0x51, 0x21, 0x21, 0x2e, 0x96, 0xe2, 0x92, 0xdc, 0x12, 0x09, 0x46, 0x05,
	0x26, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0x48, 0x93, 0x8b, 0x25, 0xb1, 0x28, 0xbd, 0x58, 0x82, 0x49,
	0x81, 0x59, 0x83, 0xdb, 0x48, 0x54, 0x0f, 0xea, 0x0e, 0x17, 0xb8, 0x6d, 0x8e, 0x79, 0x95, 0x41,
	0x60, 0x25, 0x42, 0x06, 0x5c, 0x02, 0xc9, 0xf9, 0xb9, 0x05, 0x89, 0xc9, 0x25, 0xf1, 0xb9, 0xa9,
	0x25, 0x89, 0x20, 0x07, 0x49, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x58, 0xb1, 0xa6, 0x25, 0xe6, 0x14,
	0xa7, 0x06, 0xf1, 0x43, 0xa5, 0x7d, 0xa1, 0xb2, 0x56, 0x2c, 0x1d, 0xaf, 0x0c, 0x78, 0x94, 0x44,
	0xb9, 0x78, 0x91, 0x1c, 0xea, 0x9f, 0x6d, 0xc5, 0x32, 0xe1, 0x95, 0x81, 0xa0, 0x93, 0x24, 0x97,
	0x78, 0x72, 0x7e, 0xae, 0x1e, 0xd8, 0x7b, 0x7a, 0xc9, 0x59, 0x7a, 0x50, 0x0f, 0x27, 0x95, 0xa6,
	0x01, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0x49, 0x1d, 0xa3, 0x20, 0x01, 0x00, 0x00,
}
