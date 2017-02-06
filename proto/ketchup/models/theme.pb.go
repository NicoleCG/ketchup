// Code generated by protoc-gen-go.
// source: theme.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Theme struct {
	Uuid             *string                   `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name             *string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Templates        map[string]*ThemeTemplate `protobuf:"bytes,10,rep,name=templates" json:"templates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Assets           map[string]*ThemeAsset    `protobuf:"bytes,11,rep,name=assets" json:"assets,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Theme) Reset()                    { *m = Theme{} }
func (m *Theme) String() string            { return proto.CompactTextString(m) }
func (*Theme) ProtoMessage()               {}
func (*Theme) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Theme) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Theme) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Theme) GetTemplates() map[string]*ThemeTemplate {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *Theme) GetAssets() map[string]*ThemeAsset {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ThemeTemplate struct {
	Uuid         *string             `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name         *string             `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Theme        *string             `protobuf:"bytes,3,opt,name=theme" json:"theme,omitempty"`
	Engine       *string             `protobuf:"bytes,4,opt,name=engine,def=html" json:"engine,omitempty"`
	HideContent  *bool               `protobuf:"varint,5,opt,name=hideContent,def=0" json:"hideContent,omitempty"`
	Description  *string             `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Placeholders []*ThemePlaceholder `protobuf:"bytes,10,rep,name=placeholders" json:"placeholders,omitempty"`
	// maybe an option for extensions?
	Data             *string `protobuf:"bytes,11,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ThemeTemplate) Reset()                    { *m = ThemeTemplate{} }
func (m *ThemeTemplate) String() string            { return proto.CompactTextString(m) }
func (*ThemeTemplate) ProtoMessage()               {}
func (*ThemeTemplate) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

const Default_ThemeTemplate_Engine string = "html"
const Default_ThemeTemplate_HideContent bool = false

func (m *ThemeTemplate) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ThemeTemplate) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ThemeTemplate) GetTheme() string {
	if m != nil && m.Theme != nil {
		return *m.Theme
	}
	return ""
}

func (m *ThemeTemplate) GetEngine() string {
	if m != nil && m.Engine != nil {
		return *m.Engine
	}
	return Default_ThemeTemplate_Engine
}

func (m *ThemeTemplate) GetHideContent() bool {
	if m != nil && m.HideContent != nil {
		return *m.HideContent
	}
	return Default_ThemeTemplate_HideContent
}

func (m *ThemeTemplate) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ThemeTemplate) GetPlaceholders() []*ThemePlaceholder {
	if m != nil {
		return m.Placeholders
	}
	return nil
}

func (m *ThemeTemplate) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

type ThemePlaceholder struct {
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*ThemePlaceholder_Short
	//	*ThemePlaceholder_Text
	//	*ThemePlaceholder_Multiple
	Type             IsThemePlaceholder_Type `protobuf_oneof:"type"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ThemePlaceholder) Reset()                    { *m = ThemePlaceholder{} }
func (m *ThemePlaceholder) String() string            { return proto.CompactTextString(m) }
func (*ThemePlaceholder) ProtoMessage()               {}
func (*ThemePlaceholder) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type IsThemePlaceholder_Type interface {
	IsThemePlaceholder_Type()
}

type ThemePlaceholder_Short struct {
	Short *ContentString `protobuf:"bytes,11,opt,name=short,oneof"`
}
type ThemePlaceholder_Text struct {
	Text *ContentText `protobuf:"bytes,12,opt,name=text,oneof"`
}
type ThemePlaceholder_Multiple struct {
	Multiple *ContentMultiple `protobuf:"bytes,13,opt,name=multiple,oneof"`
}

func (*ThemePlaceholder_Short) IsThemePlaceholder_Type()    {}
func (*ThemePlaceholder_Text) IsThemePlaceholder_Type()     {}
func (*ThemePlaceholder_Multiple) IsThemePlaceholder_Type() {}

func (m *ThemePlaceholder) GetType() IsThemePlaceholder_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ThemePlaceholder) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ThemePlaceholder) GetShort() *ContentString {
	if x, ok := m.GetType().(*ThemePlaceholder_Short); ok {
		return x.Short
	}
	return nil
}

func (m *ThemePlaceholder) GetText() *ContentText {
	if x, ok := m.GetType().(*ThemePlaceholder_Text); ok {
		return x.Text
	}
	return nil
}

func (m *ThemePlaceholder) GetMultiple() *ContentMultiple {
	if x, ok := m.GetType().(*ThemePlaceholder_Multiple); ok {
		return x.Multiple
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ThemePlaceholder) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ThemePlaceholder_OneofMarshaler, _ThemePlaceholder_OneofUnmarshaler, _ThemePlaceholder_OneofSizer, []interface{}{
		(*ThemePlaceholder_Short)(nil),
		(*ThemePlaceholder_Text)(nil),
		(*ThemePlaceholder_Multiple)(nil),
	}
}

func _ThemePlaceholder_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ThemePlaceholder)
	// type
	switch x := m.Type.(type) {
	case *ThemePlaceholder_Short:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Short); err != nil {
			return err
		}
	case *ThemePlaceholder_Text:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *ThemePlaceholder_Multiple:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Multiple); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ThemePlaceholder.Type has unexpected type %T", x)
	}
	return nil
}

func _ThemePlaceholder_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ThemePlaceholder)
	switch tag {
	case 11: // type.short
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentString)
		err := b.DecodeMessage(msg)
		m.Type = &ThemePlaceholder_Short{msg}
		return true, err
	case 12: // type.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentText)
		err := b.DecodeMessage(msg)
		m.Type = &ThemePlaceholder_Text{msg}
		return true, err
	case 13: // type.multiple
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentMultiple)
		err := b.DecodeMessage(msg)
		m.Type = &ThemePlaceholder_Multiple{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ThemePlaceholder_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ThemePlaceholder)
	// type
	switch x := m.Type.(type) {
	case *ThemePlaceholder_Short:
		s := proto.Size(x.Short)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ThemePlaceholder_Text:
		s := proto.Size(x.Text)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ThemePlaceholder_Multiple:
		s := proto.Size(x.Multiple)
		n += proto.SizeVarint(13<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ThemeAsset struct {
	Uuid             *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Theme            *string `protobuf:"bytes,3,opt,name=theme" json:"theme,omitempty"`
	Data             *string `protobuf:"bytes,10,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ThemeAsset) Reset()                    { *m = ThemeAsset{} }
func (m *ThemeAsset) String() string            { return proto.CompactTextString(m) }
func (*ThemeAsset) ProtoMessage()               {}
func (*ThemeAsset) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *ThemeAsset) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *ThemeAsset) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ThemeAsset) GetTheme() string {
	if m != nil && m.Theme != nil {
		return *m.Theme
	}
	return ""
}

func (m *ThemeAsset) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Theme)(nil), "ketchup.models.Theme")
	proto.RegisterType((*ThemeTemplate)(nil), "ketchup.models.ThemeTemplate")
	proto.RegisterType((*ThemePlaceholder)(nil), "ketchup.models.ThemePlaceholder")
	proto.RegisterType((*ThemeAsset)(nil), "ketchup.models.ThemeAsset")
}
func (m *Theme) SetUuid(v *string) {
	m.Uuid = v
}

func (m *Theme) SetName(v *string) {
	m.Name = v
}

func (m *Theme) SetTemplates(v map[string]*ThemeTemplate) {
	m.Templates = v
}

func (m *Theme) SetAssets(v map[string]*ThemeAsset) {
	m.Assets = v
}

func (m *ThemeTemplate) SetUuid(v *string) {
	m.Uuid = v
}

func (m *ThemeTemplate) SetName(v *string) {
	m.Name = v
}

func (m *ThemeTemplate) SetTheme(v *string) {
	m.Theme = v
}

func (m *ThemeTemplate) SetEngine(v *string) {
	m.Engine = v
}

func (m *ThemeTemplate) SetHideContent(v *bool) {
	m.HideContent = v
}

func (m *ThemeTemplate) SetDescription(v *string) {
	m.Description = v
}

func (m *ThemeTemplate) SetPlaceholders(v []*ThemePlaceholder) {
	m.Placeholders = v
}

func (m *ThemeTemplate) SetData(v *string) {
	m.Data = v
}

func (m *ThemePlaceholder) SetKey(v *string) {
	m.Key = v
}

func (m *ThemeAsset) SetUuid(v *string) {
	m.Uuid = v
}

func (m *ThemeAsset) SetName(v *string) {
	m.Name = v
}

func (m *ThemeAsset) SetTheme(v *string) {
	m.Theme = v
}

func (m *ThemeAsset) SetData(v *string) {
	m.Data = v
}

func init() { proto.RegisterFile("theme.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x9b, 0x36, 0x29, 0x77, 0x93, 0xeb, 0x71, 0x2c, 0x3e, 0x2c, 0x51, 0x31, 0x16, 0xc1,
	0x3e, 0x05, 0x3d, 0x11, 0xb4, 0xe0, 0x83, 0xa7, 0x42, 0x5f, 0x04, 0x59, 0xeb, 0x93, 0x20, 0x2c,
	0xcd, 0xd8, 0x84, 0xdb, 0x24, 0x4b, 0x76, 0x22, 0xd7, 0x6f, 0xe1, 0x97, 0xf3, 0x0b, 0xf8, 0x49,
	0x24, 0xbb, 0xa9, 0x6d, 0x24, 0x07, 0xc2, 0xbd, 0xcd, 0xce, 0xfe, 0x7f, 0xff, 0xd9, 0x99, 0x59,
	0x08, 0x29, 0xc3, 0x02, 0x13, 0x5d, 0x57, 0x54, 0xb1, 0xf3, 0x6b, 0xa4, 0x4d, 0xd6, 0xe8, 0xa4,
	0xa8, 0x52, 0x54, 0x26, 0x02, 0x2d, 0xb7, 0xdd, 0x5d, 0x34, 0xdb, 0x54, 0x25, 0x61, 0x49, 0xee,
	0x38, 0xff, 0x3d, 0x86, 0x60, 0xdd, 0xa2, 0x8c, 0x81, 0xdf, 0x34, 0x79, 0xca, 0xbd, 0xd8, 0x5b,
	0x9c, 0x0a, 0x1b, 0xb7, 0xb9, 0x52, 0x16, 0xc8, 0xc7, 0x2e, 0xd7, 0xc6, 0xec, 0x0a, 0x4e, 0x09,
	0x0b, 0xad, 0x24, 0xa1, 0xe1, 0x10, 0x4f, 0x16, 0xe1, 0xe5, 0x93, 0xa4, 0x5f, 0x30, 0xb1, 0x8e,
	0xc9, 0x7a, 0x2f, 0xfb, 0x50, 0x52, 0xbd, 0x13, 0x07, 0x8c, 0xbd, 0x86, 0xa9, 0x34, 0x06, 0xc9,
	0xf0, 0xd0, 0x1a, 0x3c, 0x1e, 0x36, 0x78, 0x6b, 0x35, 0x8e, 0xee, 0x80, 0xe8, 0x2b, 0x9c, 0xf7,
	0x7d, 0xd9, 0x05, 0x4c, 0xae, 0x71, 0xd7, 0xbd, 0xbb, 0x0d, 0xd9, 0x0b, 0x08, 0x7e, 0x48, 0xd5,
	0xb8, 0x77, 0x87, 0x97, 0x0f, 0x07, 0xdd, 0xf7, 0x2e, 0xc2, 0x69, 0x97, 0xe3, 0x57, 0x5e, 0xf4,
	0x05, 0xc2, 0xa3, 0x9a, 0x03, 0xce, 0xcf, 0xfa, 0xce, 0xd1, 0xa0, 0xb3, 0xb5, 0x38, 0xb2, 0x9d,
	0xff, 0x1c, 0xc3, 0xac, 0x57, 0xf3, 0xbf, 0x87, 0x7d, 0x0f, 0x02, 0xbb, 0x58, 0x3e, 0xb1, 0x49,
	0x77, 0x60, 0x0f, 0x60, 0x8a, 0xe5, 0x36, 0x2f, 0x91, 0xfb, 0x6d, 0x7a, 0xe9, 0x67, 0x54, 0x28,
	0xd1, 0xe5, 0xd8, 0x53, 0x08, 0xb3, 0x3c, 0xc5, 0x77, 0x6e, 0xcf, 0x3c, 0x88, 0xbd, 0xc5, 0xc9,
	0x32, 0xf8, 0x2e, 0x95, 0x41, 0x71, 0x7c, 0xc3, 0x62, 0x08, 0x53, 0x34, 0x9b, 0x3a, 0xd7, 0x94,
	0x57, 0x25, 0x9f, 0xda, 0x12, 0xc7, 0x29, 0xf6, 0x1e, 0xce, 0xb4, 0x92, 0x1b, 0xcc, 0x2a, 0x95,
	0x62, 0xbd, 0x5f, 0x77, 0x3c, 0xd8, 0xf5, 0xa7, 0x83, 0x50, 0xf4, 0xa8, 0xb6, 0xb1, 0x54, 0x92,
	0xe4, 0xa1, 0x6b, 0xac, 0x8d, 0xe7, 0xbf, 0x3c, 0xb8, 0xf8, 0x17, 0x1b, 0x98, 0xf7, 0x4b, 0x08,
	0x4c, 0x56, 0xd5, 0x64, 0xd9, 0x81, 0x4d, 0x76, 0xad, 0x7c, 0xa6, 0x3a, 0x2f, 0xb7, 0xab, 0x91,
	0x70, 0x6a, 0xf6, 0x1c, 0x7c, 0xc2, 0x1b, 0xe2, 0x67, 0x96, 0xba, 0x7f, 0x0b, 0xb5, 0xc6, 0x1b,
	0x5a, 0x8d, 0x84, 0x95, 0xb2, 0x37, 0x70, 0x52, 0x34, 0x8a, 0x72, 0xad, 0x90, 0xcf, 0x2c, 0xf6,
	0xe8, 0x16, 0xec, 0x63, 0x27, 0x5b, 0x8d, 0xc4, 0x5f, 0xe4, 0x6a, 0x0a, 0x3e, 0xed, 0x34, 0xce,
	0xbf, 0x01, 0x1c, 0xfe, 0xc0, 0x1d, 0xd7, 0xbc, 0x9f, 0x1b, 0x1c, 0xe6, 0xf6, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x04, 0x8c, 0x8c, 0xa9, 0xe8, 0x03, 0x00, 0x00,
}
