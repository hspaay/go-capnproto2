// Code generated by capnpc-go. DO NOT EDIT.

package stream

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type StreamResult struct{ capnp.Struct }

// StreamResult_TypeID is the unique identifier for the type StreamResult.
const StreamResult_TypeID = 0x995f9a3377c0b16e

func NewStreamResult(s *capnp.Segment) (StreamResult, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StreamResult{st}, err
}

func NewRootStreamResult(s *capnp.Segment) (StreamResult, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return StreamResult{st}, err
}

func ReadRootStreamResult(msg *capnp.Message) (StreamResult, error) {
	root, err := msg.Root()
	return StreamResult{root.Struct()}, err
}

func (s StreamResult) String() string {
	str, _ := text.Marshal(0x995f9a3377c0b16e, s.Struct)
	return str
}

// StreamResult_List is a list of StreamResult.
type StreamResult_List struct{ capnp.List }

// NewStreamResult creates a new list of StreamResult.
func NewStreamResult_List(s *capnp.Segment, sz int32) (StreamResult_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return StreamResult_List{l}, err
}

func (s StreamResult_List) At(i int) StreamResult { return StreamResult{s.List.Struct(i)} }

func (s StreamResult_List) Set(i int, v StreamResult) error { return s.List.SetStruct(i, v.Struct) }

func (s StreamResult_List) String() string {
	str, _ := text.MarshalList(0x995f9a3377c0b16e, s.List)
	return str
}

// StreamResult_Future is a wrapper for a StreamResult promised by a client call.
type StreamResult_Future struct{ *capnp.Future }

func (p StreamResult_Future) Struct() (StreamResult, error) {
	s, err := p.Future.Struct()
	return StreamResult{s}, err
}

const schema_86c366a91393f3f8 = "x\xda\x12\x88p`1\xe4\xcdgb`\x0a\x94ae" +
	"\xfb\x9f\xb7\xf1@\xb9\xf1\xac\xf8\x99\x0c\x82\xbc\x8c\xff\x7f" +
	"|\x9e,\xbc2\xedp\x1b\x03\x0b;\x03\x83\xb0,\xe3" +
	"%aMFv\x06\xd7\xff\xc5%E\xa9\x89\xb9z\xc9" +
	"L\x89\x05y\x05V\xc1`^Pjqi\x0ecI" +
	"\x00## \x00\x00\xff\xff\x13|\x19\xd3"

func init() {
	schemas.Register(schema_86c366a91393f3f8,
		0x995f9a3377c0b16e)
}