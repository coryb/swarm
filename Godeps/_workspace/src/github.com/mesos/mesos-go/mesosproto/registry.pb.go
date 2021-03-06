// Code generated by protoc-gen-gogo.
// source: registry.proto
// DO NOT EDIT!

package mesosproto

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Registry struct {
	// Most recent leading master.
	Master *Registry_Master `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	// All admitted slaves.
	Slaves           *Registry_Slaves `protobuf:"bytes,2,opt,name=slaves" json:"slaves,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}

func (m *Registry) GetMaster() *Registry_Master {
	if m != nil {
		return m.Master
	}
	return nil
}

func (m *Registry) GetSlaves() *Registry_Slaves {
	if m != nil {
		return m.Slaves
	}
	return nil
}

type Registry_Master struct {
	Info             *MasterInfo `protobuf:"bytes,1,req,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Registry_Master) Reset()         { *m = Registry_Master{} }
func (m *Registry_Master) String() string { return proto.CompactTextString(m) }
func (*Registry_Master) ProtoMessage()    {}

func (m *Registry_Master) GetInfo() *MasterInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type Registry_Slave struct {
	Info             *SlaveInfo `protobuf:"bytes,1,req,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Registry_Slave) Reset()         { *m = Registry_Slave{} }
func (m *Registry_Slave) String() string { return proto.CompactTextString(m) }
func (*Registry_Slave) ProtoMessage()    {}

func (m *Registry_Slave) GetInfo() *SlaveInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type Registry_Slaves struct {
	Slaves           []*Registry_Slave `protobuf:"bytes,1,rep,name=slaves" json:"slaves,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Registry_Slaves) Reset()         { *m = Registry_Slaves{} }
func (m *Registry_Slaves) String() string { return proto.CompactTextString(m) }
func (*Registry_Slaves) ProtoMessage()    {}

func (m *Registry_Slaves) GetSlaves() []*Registry_Slave {
	if m != nil {
		return m.Slaves
	}
	return nil
}
