
package trezor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageType int32

const (
	MessageType_MessageType_Initialize               MessageType = 0
	MessageType_MessageType_Ping                     MessageType = 1
	MessageType_MessageType_Success                  MessageType = 2
	MessageType_MessageType_Failure                  MessageType = 3
	MessageType_MessageType_ChangePin                MessageType = 4
	MessageType_MessageType_WipeDevice               MessageType = 5
	MessageType_MessageType_FirmwareErase            MessageType = 6
	MessageType_MessageType_FirmwareUpload           MessageType = 7
	MessageType_MessageType_FirmwareRequest          MessageType = 8
	MessageType_MessageType_GetEntropy               MessageType = 9
	MessageType_MessageType_Entropy                  MessageType = 10
	MessageType_MessageType_GetPublicKey             MessageType = 11
	MessageType_MessageType_PublicKey                MessageType = 12
	MessageType_MessageType_LoadDevice               MessageType = 13
	MessageType_MessageType_ResetDevice              MessageType = 14
	MessageType_MessageType_SignTx                   MessageType = 15
	MessageType_MessageType_SimpleSignTx             MessageType = 16
	MessageType_MessageType_Features                 MessageType = 17
	MessageType_MessageType_PinMatrixRequest         MessageType = 18
	MessageType_MessageType_PinMatrixAck             MessageType = 19
	MessageType_MessageType_Cancel                   MessageType = 20
	MessageType_MessageType_TxRequest                MessageType = 21
	MessageType_MessageType_TxAck                    MessageType = 22
	MessageType_MessageType_CipherKeyValue           MessageType = 23
	MessageType_MessageType_ClearSession             MessageType = 24
	MessageType_MessageType_ApplySettings            MessageType = 25
	MessageType_MessageType_ButtonRequest            MessageType = 26
	MessageType_MessageType_ButtonAck                MessageType = 27
	MessageType_MessageType_ApplyFlags               MessageType = 28
	MessageType_MessageType_GetAddress               MessageType = 29
	MessageType_MessageType_Address                  MessageType = 30
	MessageType_MessageType_SelfTest                 MessageType = 32
	MessageType_MessageType_BackupDevice             MessageType = 34
	MessageType_MessageType_EntropyRequest           MessageType = 35
	MessageType_MessageType_EntropyAck               MessageType = 36
	MessageType_MessageType_SignMessage              MessageType = 38
	MessageType_MessageType_VerifyMessage            MessageType = 39
	MessageType_MessageType_MessageSignature         MessageType = 40
	MessageType_MessageType_PassphraseRequest        MessageType = 41
	MessageType_MessageType_PassphraseAck            MessageType = 42
	MessageType_MessageType_EstimateTxSize           MessageType = 43
	MessageType_MessageType_TxSize                   MessageType = 44
	MessageType_MessageType_RecoveryDevice           MessageType = 45
	MessageType_MessageType_WordRequest              MessageType = 46
	MessageType_MessageType_WordAck                  MessageType = 47
	MessageType_MessageType_CipheredKeyValue         MessageType = 48
	MessageType_MessageType_EncryptMessage           MessageType = 49
	MessageType_MessageType_EncryptedMessage         MessageType = 50
	MessageType_MessageType_DecryptMessage           MessageType = 51
	MessageType_MessageType_DecryptedMessage         MessageType = 52
	MessageType_MessageType_SignIdentity             MessageType = 53
	MessageType_MessageType_SignedIdentity           MessageType = 54
	MessageType_MessageType_GetFeatures              MessageType = 55
	MessageType_MessageType_EPVchainGetAddress       MessageType = 56
	MessageType_MessageType_EPVchainAddress					 MessageType = 57
	MessageType_MessageType_EPVchainSignTx           MessageType = 58
	MessageType_MessageType_EPVchainTxRequest        MessageType = 59
	MessageType_MessageType_EPVchainTxAck            MessageType = 60
	MessageType_MessageType_GetECDHSessionKey        MessageType = 61
	MessageType_MessageType_ECDHSessionKey           MessageType = 62
	MessageType_MessageType_SetU2FCounter            MessageType = 63
	MessageType_MessageType_EPVchainSignMessage      MessageType = 64
	MessageType_MessageType_EPVchainVerifyMessage    MessageType = 65
	MessageType_MessageType_EPVchainMessageSignature MessageType = 66
	MessageType_MessageType_DebugLinkDecision        MessageType = 100
	MessageType_MessageType_DebugLinkGetState        MessageType = 101
	MessageType_MessageType_DebugLinkState           MessageType = 102
	MessageType_MessageType_DebugLinkStop            MessageType = 103
	MessageType_MessageType_DebugLinkLog             MessageType = 104
	MessageType_MessageType_DebugLinkMemoryRead      MessageType = 110
	MessageType_MessageType_DebugLinkMemory          MessageType = 111
	MessageType_MessageType_DebugLinkMemoryWrite     MessageType = 112
	MessageType_MessageType_DebugLinkFlashErase      MessageType = 113
)

var MessageType_name = map[int32]string{
	0:   "MessageType_Initialize",
	1:   "MessageType_Ping",
	2:   "MessageType_Success",
	3:   "MessageType_Failure",
	4:   "MessageType_ChangePin",
	5:   "MessageType_WipeDevice",
	6:   "MessageType_FirmwareErase",
	7:   "MessageType_FirmwareUpload",
	8:   "MessageType_FirmwareRequest",
	9:   "MessageType_GetEntropy",
	10:  "MessageType_Entropy",
	11:  "MessageType_GetPublicKey",
	12:  "MessageType_PublicKey",
	13:  "MessageType_LoadDevice",
	14:  "MessageType_ResetDevice",
	15:  "MessageType_SignTx",
	16:  "MessageType_SimpleSignTx",
	17:  "MessageType_Features",
	18:  "MessageType_PinMatrixRequest",
	19:  "MessageType_PinMatrixAck",
	20:  "MessageType_Cancel",
	21:  "MessageType_TxRequest",
	22:  "MessageType_TxAck",
	23:  "MessageType_CipherKeyValue",
	24:  "MessageType_ClearSession",
	25:  "MessageType_ApplySettings",
	26:  "MessageType_ButtonRequest",
	27:  "MessageType_ButtonAck",
	28:  "MessageType_ApplyFlags",
	29:  "MessageType_GetAddress",
	30:  "MessageType_Address",
	32:  "MessageType_SelfTest",
	34:  "MessageType_BackupDevice",
	35:  "MessageType_EntropyRequest",
	36:  "MessageType_EntropyAck",
	38:  "MessageType_SignMessage",
	39:  "MessageType_VerifyMessage",
	40:  "MessageType_MessageSignature",
	41:  "MessageType_PassphraseRequest",
	42:  "MessageType_PassphraseAck",
	43:  "MessageType_EstimateTxSize",
	44:  "MessageType_TxSize",
	45:  "MessageType_RecoveryDevice",
	46:  "MessageType_WordRequest",
	47:  "MessageType_WordAck",
	48:  "MessageType_CipheredKeyValue",
	49:  "MessageType_EncryptMessage",
	50:  "MessageType_EncryptedMessage",
	51:  "MessageType_DecryptMessage",
	52:  "MessageType_DecryptedMessage",
	53:  "MessageType_SignIdentity",
	54:  "MessageType_SignedIdentity",
	55:  "MessageType_GetFeatures",
	56:  "MessageType_EPVchainGetAddress",
	57:  "MessageType_EPVchainAddress",
	58:  "MessageType_EPVchainSignTx",
	59:  "MessageType_EPVchainTxRequest",
	60:  "MessageType_EPVchainTxAck",
	61:  "MessageType_GetECDHSessionKey",
	62:  "MessageType_ECDHSessionKey",
	63:  "MessageType_SetU2FCounter",
	64:  "MessageType_EPVchainSignMessage",
	65:  "MessageType_EPVchainVerifyMessage",
	66:  "MessageType_EPVchainMessageSignature",
	100: "MessageType_DebugLinkDecision",
	101: "MessageType_DebugLinkGetState",
	102: "MessageType_DebugLinkState",
	103: "MessageType_DebugLinkStop",
	104: "MessageType_DebugLinkLog",
	110: "MessageType_DebugLinkMemoryRead",
	111: "MessageType_DebugLinkMemory",
	112: "MessageType_DebugLinkMemoryWrite",
	113: "MessageType_DebugLinkFlashErase",
}
var MessageType_value = map[string]int32{
	"MessageType_Initialize":               0,
	"MessageType_Ping":                     1,
	"MessageType_Success":                  2,
	"MessageType_Failure":                  3,
	"MessageType_ChangePin":                4,
	"MessageType_WipeDevice":               5,
	"MessageType_FirmwareErase":            6,
	"MessageType_FirmwareUpload":           7,
	"MessageType_FirmwareRequest":          8,
	"MessageType_GetEntropy":               9,
	"MessageType_Entropy":                  10,
	"MessageType_GetPublicKey":             11,
	"MessageType_PublicKey":                12,
	"MessageType_LoadDevice":               13,
	"MessageType_ResetDevice":              14,
	"MessageType_SignTx":                   15,
	"MessageType_SimpleSignTx":             16,
	"MessageType_Features":                 17,
	"MessageType_PinMatrixRequest":         18,
	"MessageType_PinMatrixAck":             19,
	"MessageType_Cancel":                   20,
	"MessageType_TxRequest":                21,
	"MessageType_TxAck":                    22,
	"MessageType_CipherKeyValue":           23,
	"MessageType_ClearSession":             24,
	"MessageType_ApplySettings":            25,
	"MessageType_ButtonRequest":            26,
	"MessageType_ButtonAck":                27,
	"MessageType_ApplyFlags":               28,
	"MessageType_GetAddress":               29,
	"MessageType_Address":                  30,
	"MessageType_SelfTest":                 32,
	"MessageType_BackupDevice":             34,
	"MessageType_EntropyRequest":           35,
	"MessageType_EntropyAck":               36,
	"MessageType_SignMessage":              38,
	"MessageType_VerifyMessage":            39,
	"MessageType_MessageSignature":         40,
	"MessageType_PassphraseRequest":        41,
	"MessageType_PassphraseAck":            42,
	"MessageType_EstimateTxSize":           43,
	"MessageType_TxSize":                   44,
	"MessageType_RecoveryDevice":           45,
	"MessageType_WordRequest":              46,
	"MessageType_WordAck":                  47,
	"MessageType_CipheredKeyValue":         48,
	"MessageType_EncryptMessage":           49,
	"MessageType_EncryptedMessage":         50,
	"MessageType_DecryptMessage":           51,
	"MessageType_DecryptedMessage":         52,
	"MessageType_SignIdentity":             53,
	"MessageType_SignedIdentity":           54,
	"MessageType_GetFeatures":              55,
	"MessageType_EPVchainGetAddress":       56,
	"MessageType_EPVchainAddress":          57,
	"MessageType_EPVchainSignTx":           58,
	"MessageType_EPVchainTxRequest":        59,
	"MessageType_EPVchainTxAck":            60,
	"MessageType_GetECDHSessionKey":        61,
	"MessageType_ECDHSessionKey":           62,
	"MessageType_SetU2FCounter":            63,
	"MessageType_EPVchainSignMessage":      64,
	"MessageType_EPVchainVerifyMessage":    65,
	"MessageType_EPVchainMessageSignature": 66,
	"MessageType_DebugLinkDecision":        100,
	"MessageType_DebugLinkGetState":        101,
	"MessageType_DebugLinkState":           102,
	"MessageType_DebugLinkStop":            103,
	"MessageType_DebugLinkLog":             104,
	"MessageType_DebugLinkMemoryRead":      110,
	"MessageType_DebugLinkMemory":          111,
	"MessageType_DebugLinkMemoryWrite":     112,
	"MessageType_DebugLinkFlashErase":      113,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Initialize struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Initialize) Reset()                    { *m = Initialize{} }
func (m *Initialize) String() string            { return proto.CompactTextString(m) }
func (*Initialize) ProtoMessage()               {}
func (*Initialize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type GetFeatures struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetFeatures) Reset()                    { *m = GetFeatures{} }
func (m *GetFeatures) String() string            { return proto.CompactTextString(m) }
func (*GetFeatures) ProtoMessage()               {}
func (*GetFeatures) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type Features struct {
	Vendor               *string     `protobuf:"bytes,1,opt,name=vendor" json:"vendor,omitempty"`
	MajorVersion         *uint32     `protobuf:"varint,2,opt,name=major_version,json=majorVersion" json:"major_version,omitempty"`
	MinorVersion         *uint32     `protobuf:"varint,3,opt,name=minor_version,json=minorVersion" json:"minor_version,omitempty"`
	PatchVersion         *uint32     `protobuf:"varint,4,opt,name=patch_version,json=patchVersion" json:"patch_version,omitempty"`
	BootloaderMode       *bool       `protobuf:"varint,5,opt,name=bootloader_mode,json=bootloaderMode" json:"bootloader_mode,omitempty"`
	DeviceId             *string     `protobuf:"bytes,6,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	PinProtection        *bool       `protobuf:"varint,7,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	PassphraseProtection *bool       `protobuf:"varint,8,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	Language             *string     `protobuf:"bytes,9,opt,name=language" json:"language,omitempty"`
	Label                *string     `protobuf:"bytes,10,opt,name=label" json:"label,omitempty"`
	Coins                []*CoinType `protobuf:"bytes,11,rep,name=coins" json:"coins,omitempty"`
	Initialized          *bool       `protobuf:"varint,12,opt,name=initialized" json:"initialized,omitempty"`
	Revision             []byte      `protobuf:"bytes,13,opt,name=revision" json:"revision,omitempty"`
	BootloaderHash       []byte      `protobuf:"bytes,14,opt,name=bootloader_hash,json=bootloaderHash" json:"bootloader_hash,omitempty"`
	Imported             *bool       `protobuf:"varint,15,opt,name=imported" json:"imported,omitempty"`
	PinCached            *bool       `protobuf:"varint,16,opt,name=pin_cached,json=pinCached" json:"pin_cached,omitempty"`
	PassphraseCached     *bool       `protobuf:"varint,17,opt,name=passphrase_cached,json=passphraseCached" json:"passphrase_cached,omitempty"`
	FirmwarePresent      *bool       `protobuf:"varint,18,opt,name=firmware_present,json=firmwarePresent" json:"firmware_present,omitempty"`
	NeedsBackup          *bool       `protobuf:"varint,19,opt,name=needs_backup,json=needsBackup" json:"needs_backup,omitempty"`
	Flags                *uint32     `protobuf:"varint,20,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized     []byte      `json:"-"`
}

func (m *Features) Reset()                    { *m = Features{} }
func (m *Features) String() string            { return proto.CompactTextString(m) }
func (*Features) ProtoMessage()               {}
func (*Features) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Features) GetVendor() string {
	if m != nil && m.Vendor != nil {
		return *m.Vendor
	}
	return ""
}

func (m *Features) GetMajorVersion() uint32 {
	if m != nil && m.MajorVersion != nil {
		return *m.MajorVersion
	}
	return 0
}

func (m *Features) GetMinorVersion() uint32 {
	if m != nil && m.MinorVersion != nil {
		return *m.MinorVersion
	}
	return 0
}

func (m *Features) GetPatchVersion() uint32 {
	if m != nil && m.PatchVersion != nil {
		return *m.PatchVersion
	}
	return 0
}

func (m *Features) GetBootloaderMode() bool {
	if m != nil && m.BootloaderMode != nil {
		return *m.BootloaderMode
	}
	return false
}

func (m *Features) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *Features) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *Features) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *Features) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *Features) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Features) GetCoins() []*CoinType {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Features) GetInitialized() bool {
	if m != nil && m.Initialized != nil {
		return *m.Initialized
	}
	return false
}

func (m *Features) GetRevision() []byte {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *Features) GetBootloaderHash() []byte {
	if m != nil {
		return m.BootloaderHash
	}
	return nil
}

func (m *Features) GetImported() bool {
	if m != nil && m.Imported != nil {
		return *m.Imported
	}
	return false
}

func (m *Features) GetPinCached() bool {
	if m != nil && m.PinCached != nil {
		return *m.PinCached
	}
	return false
}

func (m *Features) GetPassphraseCached() bool {
	if m != nil && m.PassphraseCached != nil {
		return *m.PassphraseCached
	}
	return false
}

func (m *Features) GetFirmwarePresent() bool {
	if m != nil && m.FirmwarePresent != nil {
		return *m.FirmwarePresent
	}
	return false
}

func (m *Features) GetNeedsBackup() bool {
	if m != nil && m.NeedsBackup != nil {
		return *m.NeedsBackup
	}
	return false
}

func (m *Features) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

type ClearSession struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ClearSession) Reset()                    { *m = ClearSession{} }
func (m *ClearSession) String() string            { return proto.CompactTextString(m) }
func (*ClearSession) ProtoMessage()               {}
func (*ClearSession) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type ApplySettings struct {
	Language         *string `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Label            *string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	UsePassphrase    *bool   `protobuf:"varint,3,opt,name=use_passphrase,json=usePassphrase" json:"use_passphrase,omitempty"`
	Homescreen       []byte  `protobuf:"bytes,4,opt,name=homescreen" json:"homescreen,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ApplySettings) Reset()                    { *m = ApplySettings{} }
func (m *ApplySettings) String() string            { return proto.CompactTextString(m) }
func (*ApplySettings) ProtoMessage()               {}
func (*ApplySettings) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ApplySettings) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *ApplySettings) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *ApplySettings) GetUsePassphrase() bool {
	if m != nil && m.UsePassphrase != nil {
		return *m.UsePassphrase
	}
	return false
}

func (m *ApplySettings) GetHomescreen() []byte {
	if m != nil {
		return m.Homescreen
	}
	return nil
}

type ApplyFlags struct {
	Flags            *uint32 `protobuf:"varint,1,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ApplyFlags) Reset()                    { *m = ApplyFlags{} }
func (m *ApplyFlags) String() string            { return proto.CompactTextString(m) }
func (*ApplyFlags) ProtoMessage()               {}
func (*ApplyFlags) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ApplyFlags) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

type ChangePin struct {
	Remove           *bool  `protobuf:"varint,1,opt,name=remove" json:"remove,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChangePin) Reset()                    { *m = ChangePin{} }
func (m *ChangePin) String() string            { return proto.CompactTextString(m) }
func (*ChangePin) ProtoMessage()               {}
func (*ChangePin) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ChangePin) GetRemove() bool {
	if m != nil && m.Remove != nil {
		return *m.Remove
	}
	return false
}

type Ping struct {
	Message              *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	ButtonProtection     *bool   `protobuf:"varint,2,opt,name=button_protection,json=buttonProtection" json:"button_protection,omitempty"`
	PinProtection        *bool   `protobuf:"varint,3,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	PassphraseProtection *bool   `protobuf:"varint,4,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *Ping) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Ping) GetButtonProtection() bool {
	if m != nil && m.ButtonProtection != nil {
		return *m.ButtonProtection
	}
	return false
}

func (m *Ping) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *Ping) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

type Success struct {
	Message          *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Success) Reset()                    { *m = Success{} }
func (m *Success) String() string            { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()               {}
func (*Success) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *Success) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type Failure struct {
	Code             *FailureType `protobuf:"varint,1,opt,name=code,enum=FailureType" json:"code,omitempty"`
	Message          *string      `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Failure) Reset()                    { *m = Failure{} }
func (m *Failure) String() string            { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()               {}
func (*Failure) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *Failure) GetCode() FailureType {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return FailureType_Failure_UnexpectedMessage
}

func (m *Failure) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ButtonRequest struct {
	Code             *ButtonRequestType `protobuf:"varint,1,opt,name=code,enum=ButtonRequestType" json:"code,omitempty"`
	Data             *string            `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ButtonRequest) Reset()                    { *m = ButtonRequest{} }
func (m *ButtonRequest) String() string            { return proto.CompactTextString(m) }
func (*ButtonRequest) ProtoMessage()               {}
func (*ButtonRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *ButtonRequest) GetCode() ButtonRequestType {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ButtonRequestType_ButtonRequest_Other
}

func (m *ButtonRequest) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

type ButtonAck struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ButtonAck) Reset()                    { *m = ButtonAck{} }
func (m *ButtonAck) String() string            { return proto.CompactTextString(m) }
func (*ButtonAck) ProtoMessage()               {}
func (*ButtonAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type PinMatrixRequest struct {
	Type             *PinMatrixRequestType `protobuf:"varint,1,opt,name=type,enum=PinMatrixRequestType" json:"type,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *PinMatrixRequest) Reset()                    { *m = PinMatrixRequest{} }
func (m *PinMatrixRequest) String() string            { return proto.CompactTextString(m) }
func (*PinMatrixRequest) ProtoMessage()               {}
func (*PinMatrixRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *PinMatrixRequest) GetType() PinMatrixRequestType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PinMatrixRequestType_PinMatrixRequestType_Current
}

type PinMatrixAck struct {
	Pin              *string `protobuf:"bytes,1,req,name=pin" json:"pin,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PinMatrixAck) Reset()                    { *m = PinMatrixAck{} }
func (m *PinMatrixAck) String() string            { return proto.CompactTextString(m) }
func (*PinMatrixAck) ProtoMessage()               {}
func (*PinMatrixAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *PinMatrixAck) GetPin() string {
	if m != nil && m.Pin != nil {
		return *m.Pin
	}
	return ""
}

type Cancel struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Cancel) Reset()                    { *m = Cancel{} }
func (m *Cancel) String() string            { return proto.CompactTextString(m) }
func (*Cancel) ProtoMessage()               {}
func (*Cancel) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

type PassphraseRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PassphraseRequest) Reset()                    { *m = PassphraseRequest{} }
func (m *PassphraseRequest) String() string            { return proto.CompactTextString(m) }
func (*PassphraseRequest) ProtoMessage()               {}
func (*PassphraseRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

type PassphraseAck struct {
	Passphrase       *string `protobuf:"bytes,1,req,name=passphrase" json:"passphrase,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassphraseAck) Reset()                    { *m = PassphraseAck{} }
func (m *PassphraseAck) String() string            { return proto.CompactTextString(m) }
func (*PassphraseAck) ProtoMessage()               {}
func (*PassphraseAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *PassphraseAck) GetPassphrase() string {
	if m != nil && m.Passphrase != nil {
		return *m.Passphrase
	}
	return ""
}

type GetEntropy struct {
	Size             *uint32 `protobuf:"varint,1,req,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetEntropy) Reset()                    { *m = GetEntropy{} }
func (m *GetEntropy) String() string            { return proto.CompactTextString(m) }
func (*GetEntropy) ProtoMessage()               {}
func (*GetEntropy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *GetEntropy) GetSize() uint32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

type Entropy struct {
	Entropy          []byte `protobuf:"bytes,1,req,name=entropy" json:"entropy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Entropy) Reset()                    { *m = Entropy{} }
func (m *Entropy) String() string            { return proto.CompactTextString(m) }
func (*Entropy) ProtoMessage()               {}
func (*Entropy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

func (m *Entropy) GetEntropy() []byte {
	if m != nil {
		return m.Entropy
	}
	return nil
}

type GetPublicKey struct {
	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	EcdsaCurveName   *string  `protobuf:"bytes,2,opt,name=ecdsa_curve_name,json=ecdsaCurveName" json:"ecdsa_curve_name,omitempty"`
	ShowDisplay      *bool    `protobuf:"varint,3,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	CoinName         *string  `protobuf:"bytes,4,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetPublicKey) Reset()                    { *m = GetPublicKey{} }
func (m *GetPublicKey) String() string            { return proto.CompactTextString(m) }
func (*GetPublicKey) ProtoMessage()               {}
func (*GetPublicKey) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

const Default_GetPublicKey_CoinName string = "Bitcoin"

func (m *GetPublicKey) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *GetPublicKey) GetEcdsaCurveName() string {
	if m != nil && m.EcdsaCurveName != nil {
		return *m.EcdsaCurveName
	}
	return ""
}

func (m *GetPublicKey) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

func (m *GetPublicKey) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_GetPublicKey_CoinName
}

type PublicKey struct {
	Node             *HDNodeType `protobuf:"bytes,1,req,name=node" json:"node,omitempty"`
	Xpub             *string     `protobuf:"bytes,2,opt,name=xpub" json:"xpub,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{20} }

func (m *PublicKey) GetNode() *HDNodeType {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *PublicKey) GetXpub() string {
	if m != nil && m.Xpub != nil {
		return *m.Xpub
	}
	return ""
}

type GetAddress struct {
	AddressN         []uint32                  `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	CoinName         *string                   `protobuf:"bytes,2,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	ShowDisplay      *bool                     `protobuf:"varint,3,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	Multisig         *MultisigRedeemScriptType `protobuf:"bytes,4,opt,name=multisig" json:"multisig,omitempty"`
	ScriptType       *InputScriptType          `protobuf:"varint,5,opt,name=script_type,json=scriptType,enum=InputScriptType,def=0" json:"script_type,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *GetAddress) Reset()                    { *m = GetAddress{} }
func (m *GetAddress) String() string            { return proto.CompactTextString(m) }
func (*GetAddress) ProtoMessage()               {}
func (*GetAddress) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{21} }

const Default_GetAddress_CoinName string = "Bitcoin"
const Default_GetAddress_ScriptType InputScriptType = InputScriptType_SPENDADDRESS

func (m *GetAddress) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *GetAddress) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_GetAddress_CoinName
}

func (m *GetAddress) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

func (m *GetAddress) GetMultisig() *MultisigRedeemScriptType {
	if m != nil {
		return m.Multisig
	}
	return nil
}

func (m *GetAddress) GetScriptType() InputScriptType {
	if m != nil && m.ScriptType != nil {
		return *m.ScriptType
	}
	return Default_GetAddress_ScriptType
}

type EPVchainGetAddress struct {
	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	ShowDisplay      *bool    `protobuf:"varint,2,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EPVchainGetAddress) Reset()                    { *m = EPVchainGetAddress{} }
func (m *EPVchainGetAddress) String() string            { return proto.CompactTextString(m) }
func (*EPVchainGetAddress) ProtoMessage()               {}
func (*EPVchainGetAddress) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{22} }

func (m *EPVchainGetAddress) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EPVchainGetAddress) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

type Address struct {
	Address          *string `protobuf:"bytes,1,req,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{23} }

func (m *Address) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type EPVchainAddress struct {
	Address          []byte `protobuf:"bytes,1,req,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EPVchainAddress) Reset()                    { *m = EPVchainAddress{} }
func (m *EPVchainAddress) String() string            { return proto.CompactTextString(m) }
func (*EPVchainAddress) ProtoMessage()               {}
func (*EPVchainAddress) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{24} }

func (m *EPVchainAddress) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type WipeDevice struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *WipeDevice) Reset()                    { *m = WipeDevice{} }
func (m *WipeDevice) String() string            { return proto.CompactTextString(m) }
func (*WipeDevice) ProtoMessage()               {}
func (*WipeDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{25} }

type LoadDevice struct {
	Mnemonic             *string     `protobuf:"bytes,1,opt,name=mnemonic" json:"mnemonic,omitempty"`
	Node                 *HDNodeType `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Pin                  *string     `protobuf:"bytes,3,opt,name=pin" json:"pin,omitempty"`
	PassphraseProtection *bool       `protobuf:"varint,4,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	Language             *string     `protobuf:"bytes,5,opt,name=language,def=english" json:"language,omitempty"`
	Label                *string     `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	SkipChecksum         *bool       `protobuf:"varint,7,opt,name=skip_checksum,json=skipChecksum" json:"skip_checksum,omitempty"`
	U2FCounter           *uint32     `protobuf:"varint,8,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	XXX_unrecognized     []byte      `json:"-"`
}

func (m *LoadDevice) Reset()                    { *m = LoadDevice{} }
func (m *LoadDevice) String() string            { return proto.CompactTextString(m) }
func (*LoadDevice) ProtoMessage()               {}
func (*LoadDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{26} }

const Default_LoadDevice_Language string = "english"

func (m *LoadDevice) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

func (m *LoadDevice) GetNode() *HDNodeType {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadDevice) GetPin() string {
	if m != nil && m.Pin != nil {
		return *m.Pin
	}
	return ""
}

func (m *LoadDevice) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *LoadDevice) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return Default_LoadDevice_Language
}

func (m *LoadDevice) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *LoadDevice) GetSkipChecksum() bool {
	if m != nil && m.SkipChecksum != nil {
		return *m.SkipChecksum
	}
	return false
}

func (m *LoadDevice) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}

type ResetDevice struct {
	DisplayRandom        *bool   `protobuf:"varint,1,opt,name=display_random,json=displayRandom" json:"display_random,omitempty"`
	Strength             *uint32 `protobuf:"varint,2,opt,name=strength,def=256" json:"strength,omitempty"`
	PassphraseProtection *bool   `protobuf:"varint,3,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	PinProtection        *bool   `protobuf:"varint,4,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	Language             *string `protobuf:"bytes,5,opt,name=language,def=english" json:"language,omitempty"`
	Label                *string `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	U2FCounter           *uint32 `protobuf:"varint,7,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	SkipBackup           *bool   `protobuf:"varint,8,opt,name=skip_backup,json=skipBackup" json:"skip_backup,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *ResetDevice) Reset()                    { *m = ResetDevice{} }
func (m *ResetDevice) String() string            { return proto.CompactTextString(m) }
func (*ResetDevice) ProtoMessage()               {}
func (*ResetDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{27} }

const Default_ResetDevice_Strength uint32 = 256
const Default_ResetDevice_Language string = "english"

func (m *ResetDevice) GetDisplayRandom() bool {
	if m != nil && m.DisplayRandom != nil {
		return *m.DisplayRandom
	}
	return false
}

func (m *ResetDevice) GetStrength() uint32 {
	if m != nil && m.Strength != nil {
		return *m.Strength
	}
	return Default_ResetDevice_Strength
}

func (m *ResetDevice) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *ResetDevice) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *ResetDevice) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return Default_ResetDevice_Language
}

func (m *ResetDevice) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *ResetDevice) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}

func (m *ResetDevice) GetSkipBackup() bool {
	if m != nil && m.SkipBackup != nil {
		return *m.SkipBackup
	}
	return false
}

type BackupDevice struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BackupDevice) Reset()                    { *m = BackupDevice{} }
func (m *BackupDevice) String() string            { return proto.CompactTextString(m) }
func (*BackupDevice) ProtoMessage()               {}
func (*BackupDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{28} }

type EntropyRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *EntropyRequest) Reset()                    { *m = EntropyRequest{} }
func (m *EntropyRequest) String() string            { return proto.CompactTextString(m) }
func (*EntropyRequest) ProtoMessage()               {}
func (*EntropyRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{29} }

type EntropyAck struct {
	Entropy          []byte `protobuf:"bytes,1,opt,name=entropy" json:"entropy,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EntropyAck) Reset()                    { *m = EntropyAck{} }
func (m *EntropyAck) String() string            { return proto.CompactTextString(m) }
func (*EntropyAck) ProtoMessage()               {}
func (*EntropyAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{30} }

func (m *EntropyAck) GetEntropy() []byte {
	if m != nil {
		return m.Entropy
	}
	return nil
}

type RecoveryDevice struct {
	WordCount            *uint32 `protobuf:"varint,1,opt,name=word_count,json=wordCount" json:"word_count,omitempty"`
	PassphraseProtection *bool   `protobuf:"varint,2,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	PinProtection        *bool   `protobuf:"varint,3,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	Language             *string `protobuf:"bytes,4,opt,name=language,def=english" json:"language,omitempty"`
	Label                *string `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	EnforceWordlist      *bool   `protobuf:"varint,6,opt,name=enforce_wordlist,json=enforceWordlist" json:"enforce_wordlist,omitempty"`

	Type             *uint32 `protobuf:"varint,8,opt,name=type" json:"type,omitempty"`
	U2FCounter       *uint32 `protobuf:"varint,9,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	DryRun           *bool   `protobuf:"varint,10,opt,name=dry_run,json=dryRun" json:"dry_run,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RecoveryDevice) Reset()                    { *m = RecoveryDevice{} }
func (m *RecoveryDevice) String() string            { return proto.CompactTextString(m) }
func (*RecoveryDevice) ProtoMessage()               {}
func (*RecoveryDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{31} }

const Default_RecoveryDevice_Language string = "english"

func (m *RecoveryDevice) GetWordCount() uint32 {
	if m != nil && m.WordCount != nil {
		return *m.WordCount
	}
	return 0
}

func (m *RecoveryDevice) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *RecoveryDevice) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *RecoveryDevice) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return Default_RecoveryDevice_Language
}

func (m *RecoveryDevice) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *RecoveryDevice) GetEnforceWordlist() bool {
	if m != nil && m.EnforceWordlist != nil {
		return *m.EnforceWordlist
	}
	return false
}

func (m *RecoveryDevice) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *RecoveryDevice) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}

func (m *RecoveryDevice) GetDryRun() bool {
	if m != nil && m.DryRun != nil {
		return *m.DryRun
	}
	return false
}

type WordRequest struct {
	Type             *WordRequestType `protobuf:"varint,1,opt,name=type,enum=WordRequestType" json:"type,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *WordRequest) Reset()                    { *m = WordRequest{} }
func (m *WordRequest) String() string            { return proto.CompactTextString(m) }
func (*WordRequest) ProtoMessage()               {}
func (*WordRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{32} }

func (m *WordRequest) GetType() WordRequestType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return WordRequestType_WordRequestType_Plain
}

type WordAck struct {
	Word             *string `protobuf:"bytes,1,req,name=word" json:"word,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WordAck) Reset()                    { *m = WordAck{} }
func (m *WordAck) String() string            { return proto.CompactTextString(m) }
func (*WordAck) ProtoMessage()               {}
func (*WordAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{33} }

func (m *WordAck) GetWord() string {
	if m != nil && m.Word != nil {
		return *m.Word
	}
	return ""
}

type SignMessage struct {
	AddressN         []uint32         `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Message          []byte           `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	CoinName         *string          `protobuf:"bytes,3,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	ScriptType       *InputScriptType `protobuf:"varint,4,opt,name=script_type,json=scriptType,enum=InputScriptType,def=0" json:"script_type,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SignMessage) Reset()                    { *m = SignMessage{} }
func (m *SignMessage) String() string            { return proto.CompactTextString(m) }
func (*SignMessage) ProtoMessage()               {}
func (*SignMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{34} }

const Default_SignMessage_CoinName string = "Bitcoin"
const Default_SignMessage_ScriptType InputScriptType = InputScriptType_SPENDADDRESS

func (m *SignMessage) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *SignMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignMessage) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_SignMessage_CoinName
}

func (m *SignMessage) GetScriptType() InputScriptType {
	if m != nil && m.ScriptType != nil {
		return *m.ScriptType
	}
	return Default_SignMessage_ScriptType
}

type VerifyMessage struct {
	Address          *string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Signature        []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Message          []byte  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	CoinName         *string `protobuf:"bytes,4,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VerifyMessage) Reset()                    { *m = VerifyMessage{} }
func (m *VerifyMessage) String() string            { return proto.CompactTextString(m) }
func (*VerifyMessage) ProtoMessage()               {}
func (*VerifyMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{35} }

const Default_VerifyMessage_CoinName string = "Bitcoin"

func (m *VerifyMessage) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *VerifyMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *VerifyMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *VerifyMessage) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_VerifyMessage_CoinName
}

type MessageSignature struct {
	Address          *string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Signature        []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MessageSignature) Reset()                    { *m = MessageSignature{} }
func (m *MessageSignature) String() string            { return proto.CompactTextString(m) }
func (*MessageSignature) ProtoMessage()               {}
func (*MessageSignature) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{36} }

func (m *MessageSignature) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *MessageSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type EncryptMessage struct {
	Pubkey           []byte   `protobuf:"bytes,1,opt,name=pubkey" json:"pubkey,omitempty"`
	Message          []byte   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	DisplayOnly      *bool    `protobuf:"varint,3,opt,name=display_only,json=displayOnly" json:"display_only,omitempty"`
	AddressN         []uint32 `protobuf:"varint,4,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	CoinName         *string  `protobuf:"bytes,5,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EncryptMessage) Reset()                    { *m = EncryptMessage{} }
func (m *EncryptMessage) String() string            { return proto.CompactTextString(m) }
func (*EncryptMessage) ProtoMessage()               {}
func (*EncryptMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{37} }

const Default_EncryptMessage_CoinName string = "Bitcoin"

func (m *EncryptMessage) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *EncryptMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *EncryptMessage) GetDisplayOnly() bool {
	if m != nil && m.DisplayOnly != nil {
		return *m.DisplayOnly
	}
	return false
}

func (m *EncryptMessage) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EncryptMessage) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_EncryptMessage_CoinName
}

type EncryptedMessage struct {
	Nonce            []byte `protobuf:"bytes,1,opt,name=nonce" json:"nonce,omitempty"`
	Message          []byte `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Hmac             []byte `protobuf:"bytes,3,opt,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EncryptedMessage) Reset()                    { *m = EncryptedMessage{} }
func (m *EncryptedMessage) String() string            { return proto.CompactTextString(m) }
func (*EncryptedMessage) ProtoMessage()               {}
func (*EncryptedMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{38} }

func (m *EncryptedMessage) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EncryptedMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *EncryptedMessage) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type DecryptMessage struct {
	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Nonce            []byte   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	Message          []byte   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Hmac             []byte   `protobuf:"bytes,4,opt,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DecryptMessage) Reset()                    { *m = DecryptMessage{} }
func (m *DecryptMessage) String() string            { return proto.CompactTextString(m) }
func (*DecryptMessage) ProtoMessage()               {}
func (*DecryptMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{39} }

func (m *DecryptMessage) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *DecryptMessage) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *DecryptMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *DecryptMessage) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type DecryptedMessage struct {
	Message          []byte  `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DecryptedMessage) Reset()                    { *m = DecryptedMessage{} }
func (m *DecryptedMessage) String() string            { return proto.CompactTextString(m) }
func (*DecryptedMessage) ProtoMessage()               {}
func (*DecryptedMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{40} }

func (m *DecryptedMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *DecryptedMessage) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type CipherKeyValue struct {
	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Key              *string  `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Value            []byte   `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Encrypt          *bool    `protobuf:"varint,4,opt,name=encrypt" json:"encrypt,omitempty"`
	AskOnEncrypt     *bool    `protobuf:"varint,5,opt,name=ask_on_encrypt,json=askOnEncrypt" json:"ask_on_encrypt,omitempty"`
	AskOnDecrypt     *bool    `protobuf:"varint,6,opt,name=ask_on_decrypt,json=askOnDecrypt" json:"ask_on_decrypt,omitempty"`
	Iv               []byte   `protobuf:"bytes,7,opt,name=iv" json:"iv,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CipherKeyValue) Reset()                    { *m = CipherKeyValue{} }
func (m *CipherKeyValue) String() string            { return proto.CompactTextString(m) }
func (*CipherKeyValue) ProtoMessage()               {}
func (*CipherKeyValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{41} }

func (m *CipherKeyValue) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *CipherKeyValue) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *CipherKeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *CipherKeyValue) GetEncrypt() bool {
	if m != nil && m.Encrypt != nil {
		return *m.Encrypt
	}
	return false
}

func (m *CipherKeyValue) GetAskOnEncrypt() bool {
	if m != nil && m.AskOnEncrypt != nil {
		return *m.AskOnEncrypt
	}
	return false
}

func (m *CipherKeyValue) GetAskOnDecrypt() bool {
	if m != nil && m.AskOnDecrypt != nil {
		return *m.AskOnDecrypt
	}
	return false
}

func (m *CipherKeyValue) GetIv() []byte {
	if m != nil {
		return m.Iv
	}
	return nil
}

type CipheredKeyValue struct {
	Value            []byte `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CipheredKeyValue) Reset()                    { *m = CipheredKeyValue{} }
func (m *CipheredKeyValue) String() string            { return proto.CompactTextString(m) }
func (*CipheredKeyValue) ProtoMessage()               {}
func (*CipheredKeyValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{42} }

func (m *CipheredKeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type EstimateTxSize struct {
	OutputsCount     *uint32 `protobuf:"varint,1,req,name=outputs_count,json=outputsCount" json:"outputs_count,omitempty"`
	InputsCount      *uint32 `protobuf:"varint,2,req,name=inputs_count,json=inputsCount" json:"inputs_count,omitempty"`
	CoinName         *string `protobuf:"bytes,3,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EstimateTxSize) Reset()                    { *m = EstimateTxSize{} }
func (m *EstimateTxSize) String() string            { return proto.CompactTextString(m) }
func (*EstimateTxSize) ProtoMessage()               {}
func (*EstimateTxSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{43} }

const Default_EstimateTxSize_CoinName string = "Bitcoin"

func (m *EstimateTxSize) GetOutputsCount() uint32 {
	if m != nil && m.OutputsCount != nil {
		return *m.OutputsCount
	}
	return 0
}

func (m *EstimateTxSize) GetInputsCount() uint32 {
	if m != nil && m.InputsCount != nil {
		return *m.InputsCount
	}
	return 0
}

func (m *EstimateTxSize) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_EstimateTxSize_CoinName
}

type TxSize struct {
	TxSize           *uint32 `protobuf:"varint,1,opt,name=tx_size,json=txSize" json:"tx_size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TxSize) Reset()                    { *m = TxSize{} }
func (m *TxSize) String() string            { return proto.CompactTextString(m) }
func (*TxSize) ProtoMessage()               {}
func (*TxSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{44} }

func (m *TxSize) GetTxSize() uint32 {
	if m != nil && m.TxSize != nil {
		return *m.TxSize
	}
	return 0
}

type SignTx struct {
	OutputsCount     *uint32 `protobuf:"varint,1,req,name=outputs_count,json=outputsCount" json:"outputs_count,omitempty"`
	InputsCount      *uint32 `protobuf:"varint,2,req,name=inputs_count,json=inputsCount" json:"inputs_count,omitempty"`
	CoinName         *string `protobuf:"bytes,3,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	Version          *uint32 `protobuf:"varint,4,opt,name=version,def=1" json:"version,omitempty"`
	LockTime         *uint32 `protobuf:"varint,5,opt,name=lock_time,json=lockTime,def=0" json:"lock_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SignTx) Reset()                    { *m = SignTx{} }
func (m *SignTx) String() string            { return proto.CompactTextString(m) }
func (*SignTx) ProtoMessage()               {}
func (*SignTx) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{45} }

const Default_SignTx_CoinName string = "Bitcoin"
const Default_SignTx_Version uint32 = 1
const Default_SignTx_LockTime uint32 = 0

func (m *SignTx) GetOutputsCount() uint32 {
	if m != nil && m.OutputsCount != nil {
		return *m.OutputsCount
	}
	return 0
}

func (m *SignTx) GetInputsCount() uint32 {
	if m != nil && m.InputsCount != nil {
		return *m.InputsCount
	}
	return 0
}

func (m *SignTx) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_SignTx_CoinName
}

func (m *SignTx) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_SignTx_Version
}

func (m *SignTx) GetLockTime() uint32 {
	if m != nil && m.LockTime != nil {
		return *m.LockTime
	}
	return Default_SignTx_LockTime
}

type SimpleSignTx struct {
	Inputs           []*TxInputType     `protobuf:"bytes,1,rep,name=inputs" json:"inputs,omitempty"`
	Outputs          []*TxOutputType    `protobuf:"bytes,2,rep,name=outputs" json:"outputs,omitempty"`
	Transactions     []*TransactionType `protobuf:"bytes,3,rep,name=transactions" json:"transactions,omitempty"`
	CoinName         *string            `protobuf:"bytes,4,opt,name=coin_name,json=coinName,def=Bitcoin" json:"coin_name,omitempty"`
	Version          *uint32            `protobuf:"varint,5,opt,name=version,def=1" json:"version,omitempty"`
	LockTime         *uint32            `protobuf:"varint,6,opt,name=lock_time,json=lockTime,def=0" json:"lock_time,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SimpleSignTx) Reset()                    { *m = SimpleSignTx{} }
func (m *SimpleSignTx) String() string            { return proto.CompactTextString(m) }
func (*SimpleSignTx) ProtoMessage()               {}
func (*SimpleSignTx) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{46} }

const Default_SimpleSignTx_CoinName string = "Bitcoin"
const Default_SimpleSignTx_Version uint32 = 1
const Default_SimpleSignTx_LockTime uint32 = 0

func (m *SimpleSignTx) GetInputs() []*TxInputType {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *SimpleSignTx) GetOutputs() []*TxOutputType {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *SimpleSignTx) GetTransactions() []*TransactionType {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *SimpleSignTx) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_SimpleSignTx_CoinName
}

func (m *SimpleSignTx) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_SimpleSignTx_Version
}

func (m *SimpleSignTx) GetLockTime() uint32 {
	if m != nil && m.LockTime != nil {
		return *m.LockTime
	}
	return Default_SimpleSignTx_LockTime
}

type TxRequest struct {
	RequestType      *RequestType             `protobuf:"varint,1,opt,name=request_type,json=requestType,enum=RequestType" json:"request_type,omitempty"`
	Details          *TxRequestDetailsType    `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
	Serialized       *TxRequestSerializedType `protobuf:"bytes,3,opt,name=serialized" json:"serialized,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *TxRequest) Reset()                    { *m = TxRequest{} }
func (m *TxRequest) String() string            { return proto.CompactTextString(m) }
func (*TxRequest) ProtoMessage()               {}
func (*TxRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{47} }

func (m *TxRequest) GetRequestType() RequestType {
	if m != nil && m.RequestType != nil {
		return *m.RequestType
	}
	return RequestType_TXINPUT
}

func (m *TxRequest) GetDetails() *TxRequestDetailsType {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *TxRequest) GetSerialized() *TxRequestSerializedType {
	if m != nil {
		return m.Serialized
	}
	return nil
}

type TxAck struct {
	Tx               *TransactionType `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *TxAck) Reset()                    { *m = TxAck{} }
func (m *TxAck) String() string            { return proto.CompactTextString(m) }
func (*TxAck) ProtoMessage()               {}
func (*TxAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{48} }

func (m *TxAck) GetTx() *TransactionType {
	if m != nil {
		return m.Tx
	}
	return nil
}

type EPVchainSignTx struct {
	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Nonce            []byte   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	GasPrice         []byte   `protobuf:"bytes,3,opt,name=gas_price,json=gasPrice" json:"gas_price,omitempty"`
	GasLimit         []byte   `protobuf:"bytes,4,opt,name=gas_limit,json=gasLimit" json:"gas_limit,omitempty"`
	To               []byte   `protobuf:"bytes,5,opt,name=to" json:"to,omitempty"`
	Value            []byte   `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`
	DataInitialChunk []byte   `protobuf:"bytes,7,opt,name=data_initial_chunk,json=dataInitialChunk" json:"data_initial_chunk,omitempty"`
	DataLength       *uint32  `protobuf:"varint,8,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`
	ChainId          *uint32  `protobuf:"varint,9,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EPVchainSignTx) Reset()                    { *m = EPVchainSignTx{} }
func (m *EPVchainSignTx) String() string            { return proto.CompactTextString(m) }
func (*EPVchainSignTx) ProtoMessage()               {}
func (*EPVchainSignTx) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{49} }

func (m *EPVchainSignTx) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EPVchainSignTx) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *EPVchainSignTx) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *EPVchainSignTx) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *EPVchainSignTx) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *EPVchainSignTx) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *EPVchainSignTx) GetDataInitialChunk() []byte {
	if m != nil {
		return m.DataInitialChunk
	}
	return nil
}

func (m *EPVchainSignTx) GetDataLength() uint32 {
	if m != nil && m.DataLength != nil {
		return *m.DataLength
	}
	return 0
}

func (m *EPVchainSignTx) GetChainId() uint32 {
	if m != nil && m.ChainId != nil {
		return *m.ChainId
	}
	return 0
}

type EPVchainTxRequest struct {
	DataLength       *uint32 `protobuf:"varint,1,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`
	SignatureV       *uint32 `protobuf:"varint,2,opt,name=signature_v,json=signatureV" json:"signature_v,omitempty"`
	SignatureR       []byte  `protobuf:"bytes,3,opt,name=signature_r,json=signatureR" json:"signature_r,omitempty"`
	SignatureS       []byte  `protobuf:"bytes,4,opt,name=signature_s,json=signatureS" json:"signature_s,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EPVchainTxRequest) Reset()                    { *m = EPVchainTxRequest{} }
func (m *EPVchainTxRequest) String() string            { return proto.CompactTextString(m) }
func (*EPVchainTxRequest) ProtoMessage()               {}
func (*EPVchainTxRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{50} }

func (m *EPVchainTxRequest) GetDataLength() uint32 {
	if m != nil && m.DataLength != nil {
		return *m.DataLength
	}
	return 0
}

func (m *EPVchainTxRequest) GetSignatureV() uint32 {
	if m != nil && m.SignatureV != nil {
		return *m.SignatureV
	}
	return 0
}

func (m *EPVchainTxRequest) GetSignatureR() []byte {
	if m != nil {
		return m.SignatureR
	}
	return nil
}

func (m *EPVchainTxRequest) GetSignatureS() []byte {
	if m != nil {
		return m.SignatureS
	}
	return nil
}

type EPVchainTxAck struct {
	DataChunk        []byte `protobuf:"bytes,1,opt,name=data_chunk,json=dataChunk" json:"data_chunk,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EPVchainTxAck) Reset()                    { *m = EPVchainTxAck{} }
func (m *EPVchainTxAck) String() string            { return proto.CompactTextString(m) }
func (*EPVchainTxAck) ProtoMessage()               {}
func (*EPVchainTxAck) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{51} }

func (m *EPVchainTxAck) GetDataChunk() []byte {
	if m != nil {
		return m.DataChunk
	}
	return nil
}

type EPVchainSignMessage struct {
	AddressN         []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Message          []byte   `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *EPVchainSignMessage) Reset()                    { *m = EPVchainSignMessage{} }
func (m *EPVchainSignMessage) String() string            { return proto.CompactTextString(m) }
func (*EPVchainSignMessage) ProtoMessage()               {}
func (*EPVchainSignMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{52} }

func (m *EPVchainSignMessage) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *EPVchainSignMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type EPVchainVerifyMessage struct {
	Address          []byte `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Signature        []byte `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Message          []byte `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EPVchainVerifyMessage) Reset()                    { *m = EPVchainVerifyMessage{} }
func (m *EPVchainVerifyMessage) String() string            { return proto.CompactTextString(m) }
func (*EPVchainVerifyMessage) ProtoMessage()               {}
func (*EPVchainVerifyMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{53} }

func (m *EPVchainVerifyMessage) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *EPVchainVerifyMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *EPVchainVerifyMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type EPVchainMessageSignature struct {
	Address          []byte `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Signature        []byte `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EPVchainMessageSignature) Reset()                    { *m = EPVchainMessageSignature{} }
func (m *EPVchainMessageSignature) String() string            { return proto.CompactTextString(m) }
func (*EPVchainMessageSignature) ProtoMessage()               {}
func (*EPVchainMessageSignature) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{54} }

func (m *EPVchainMessageSignature) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *EPVchainMessageSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignIdentity struct {
	Identity         *IdentityType `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	ChallengeHidden  []byte        `protobuf:"bytes,2,opt,name=challenge_hidden,json=challengeHidden" json:"challenge_hidden,omitempty"`
	ChallengeVisual  *string       `protobuf:"bytes,3,opt,name=challenge_visual,json=challengeVisual" json:"challenge_visual,omitempty"`
	EcdsaCurveName   *string       `protobuf:"bytes,4,opt,name=ecdsa_curve_name,json=ecdsaCurveName" json:"ecdsa_curve_name,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SignIdentity) Reset()                    { *m = SignIdentity{} }
func (m *SignIdentity) String() string            { return proto.CompactTextString(m) }
func (*SignIdentity) ProtoMessage()               {}
func (*SignIdentity) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{55} }

func (m *SignIdentity) GetIdentity() *IdentityType {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *SignIdentity) GetChallengeHidden() []byte {
	if m != nil {
		return m.ChallengeHidden
	}
	return nil
}

func (m *SignIdentity) GetChallengeVisual() string {
	if m != nil && m.ChallengeVisual != nil {
		return *m.ChallengeVisual
	}
	return ""
}

func (m *SignIdentity) GetEcdsaCurveName() string {
	if m != nil && m.EcdsaCurveName != nil {
		return *m.EcdsaCurveName
	}
	return ""
}

type SignedIdentity struct {
	Address          *string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	PublicKey        []byte  `protobuf:"bytes,2,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	Signature        []byte  `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SignedIdentity) Reset()                    { *m = SignedIdentity{} }
func (m *SignedIdentity) String() string            { return proto.CompactTextString(m) }
func (*SignedIdentity) ProtoMessage()               {}
func (*SignedIdentity) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{56} }

func (m *SignedIdentity) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *SignedIdentity) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SignedIdentity) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type GetECDHSessionKey struct {
	Identity         *IdentityType `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	PeerPublicKey    []byte        `protobuf:"bytes,2,opt,name=peer_public_key,json=peerPublicKey" json:"peer_public_key,omitempty"`
	EcdsaCurveName   *string       `protobuf:"bytes,3,opt,name=ecdsa_curve_name,json=ecdsaCurveName" json:"ecdsa_curve_name,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GetECDHSessionKey) Reset()                    { *m = GetECDHSessionKey{} }
func (m *GetECDHSessionKey) String() string            { return proto.CompactTextString(m) }
func (*GetECDHSessionKey) ProtoMessage()               {}
func (*GetECDHSessionKey) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{57} }

func (m *GetECDHSessionKey) GetIdentity() *IdentityType {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *GetECDHSessionKey) GetPeerPublicKey() []byte {
	if m != nil {
		return m.PeerPublicKey
	}
	return nil
}

func (m *GetECDHSessionKey) GetEcdsaCurveName() string {
	if m != nil && m.EcdsaCurveName != nil {
		return *m.EcdsaCurveName
	}
	return ""
}

type ECDHSessionKey struct {
	SessionKey       []byte `protobuf:"bytes,1,opt,name=session_key,json=sessionKey" json:"session_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ECDHSessionKey) Reset()                    { *m = ECDHSessionKey{} }
func (m *ECDHSessionKey) String() string            { return proto.CompactTextString(m) }
func (*ECDHSessionKey) ProtoMessage()               {}
func (*ECDHSessionKey) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{58} }

func (m *ECDHSessionKey) GetSessionKey() []byte {
	if m != nil {
		return m.SessionKey
	}
	return nil
}

type SetU2FCounter struct {
	U2FCounter       *uint32 `protobuf:"varint,1,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetU2FCounter) Reset()                    { *m = SetU2FCounter{} }
func (m *SetU2FCounter) String() string            { return proto.CompactTextString(m) }
func (*SetU2FCounter) ProtoMessage()               {}
func (*SetU2FCounter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{59} }

func (m *SetU2FCounter) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}

type FirmwareErase struct {
	Length           *uint32 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FirmwareErase) Reset()                    { *m = FirmwareErase{} }
func (m *FirmwareErase) String() string            { return proto.CompactTextString(m) }
func (*FirmwareErase) ProtoMessage()               {}
func (*FirmwareErase) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{60} }

func (m *FirmwareErase) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type FirmwareRequest struct {
	Offset           *uint32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Length           *uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FirmwareRequest) Reset()                    { *m = FirmwareRequest{} }
func (m *FirmwareRequest) String() string            { return proto.CompactTextString(m) }
func (*FirmwareRequest) ProtoMessage()               {}
func (*FirmwareRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{61} }

func (m *FirmwareRequest) GetOffset() uint32 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *FirmwareRequest) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type FirmwareUpload struct {
	Payload          []byte `protobuf:"bytes,1,req,name=payload" json:"payload,omitempty"`
	Hash             []byte `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FirmwareUpload) Reset()                    { *m = FirmwareUpload{} }
func (m *FirmwareUpload) String() string            { return proto.CompactTextString(m) }
func (*FirmwareUpload) ProtoMessage()               {}
func (*FirmwareUpload) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{62} }

func (m *FirmwareUpload) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *FirmwareUpload) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type SelfTest struct {
	Payload          []byte `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SelfTest) Reset()                    { *m = SelfTest{} }
func (m *SelfTest) String() string            { return proto.CompactTextString(m) }
func (*SelfTest) ProtoMessage()               {}
func (*SelfTest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{63} }

func (m *SelfTest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DebugLinkDecision struct {
	YesNo            *bool  `protobuf:"varint,1,req,name=yes_no,json=yesNo" json:"yes_no,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugLinkDecision) Reset()                    { *m = DebugLinkDecision{} }
func (m *DebugLinkDecision) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkDecision) ProtoMessage()               {}
func (*DebugLinkDecision) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{64} }

func (m *DebugLinkDecision) GetYesNo() bool {
	if m != nil && m.YesNo != nil {
		return *m.YesNo
	}
	return false
}

type DebugLinkGetState struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugLinkGetState) Reset()                    { *m = DebugLinkGetState{} }
func (m *DebugLinkGetState) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkGetState) ProtoMessage()               {}
func (*DebugLinkGetState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{65} }

type DebugLinkState struct {
	Layout               []byte      `protobuf:"bytes,1,opt,name=layout" json:"layout,omitempty"`
	Pin                  *string     `protobuf:"bytes,2,opt,name=pin" json:"pin,omitempty"`
	Matrix               *string     `protobuf:"bytes,3,opt,name=matrix" json:"matrix,omitempty"`
	Mnemonic             *string     `protobuf:"bytes,4,opt,name=mnemonic" json:"mnemonic,omitempty"`
	Node                 *HDNodeType `protobuf:"bytes,5,opt,name=node" json:"node,omitempty"`
	PassphraseProtection *bool       `protobuf:"varint,6,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	ResetWord            *string     `protobuf:"bytes,7,opt,name=reset_word,json=resetWord" json:"reset_word,omitempty"`
	ResetEntropy         []byte      `protobuf:"bytes,8,opt,name=reset_entropy,json=resetEntropy" json:"reset_entropy,omitempty"`
	RecoveryFakeWord     *string     `protobuf:"bytes,9,opt,name=recovery_fake_word,json=recoveryFakeWord" json:"recovery_fake_word,omitempty"`
	RecoveryWordPos      *uint32     `protobuf:"varint,10,opt,name=recovery_word_pos,json=recoveryWordPos" json:"recovery_word_pos,omitempty"`
	XXX_unrecognized     []byte      `json:"-"`
}

func (m *DebugLinkState) Reset()                    { *m = DebugLinkState{} }
func (m *DebugLinkState) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkState) ProtoMessage()               {}
func (*DebugLinkState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{66} }

func (m *DebugLinkState) GetLayout() []byte {
	if m != nil {
		return m.Layout
	}
	return nil
}

func (m *DebugLinkState) GetPin() string {
	if m != nil && m.Pin != nil {
		return *m.Pin
	}
	return ""
}

func (m *DebugLinkState) GetMatrix() string {
	if m != nil && m.Matrix != nil {
		return *m.Matrix
	}
	return ""
}

func (m *DebugLinkState) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

func (m *DebugLinkState) GetNode() *HDNodeType {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DebugLinkState) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *DebugLinkState) GetResetWord() string {
	if m != nil && m.ResetWord != nil {
		return *m.ResetWord
	}
	return ""
}

func (m *DebugLinkState) GetResetEntropy() []byte {
	if m != nil {
		return m.ResetEntropy
	}
	return nil
}

func (m *DebugLinkState) GetRecoveryFakeWord() string {
	if m != nil && m.RecoveryFakeWord != nil {
		return *m.RecoveryFakeWord
	}
	return ""
}

func (m *DebugLinkState) GetRecoveryWordPos() uint32 {
	if m != nil && m.RecoveryWordPos != nil {
		return *m.RecoveryWordPos
	}
	return 0
}

type DebugLinkStop struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugLinkStop) Reset()                    { *m = DebugLinkStop{} }
func (m *DebugLinkStop) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkStop) ProtoMessage()               {}
func (*DebugLinkStop) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{67} }

type DebugLinkLog struct {
	Level            *uint32 `protobuf:"varint,1,opt,name=level" json:"level,omitempty"`
	Bucket           *string `protobuf:"bytes,2,opt,name=bucket" json:"bucket,omitempty"`
	Text             *string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugLinkLog) Reset()                    { *m = DebugLinkLog{} }
func (m *DebugLinkLog) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkLog) ProtoMessage()               {}
func (*DebugLinkLog) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{68} }

func (m *DebugLinkLog) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *DebugLinkLog) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *DebugLinkLog) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type DebugLinkMemoryRead struct {
	Address          *uint32 `protobuf:"varint,1,opt,name=address" json:"address,omitempty"`
	Length           *uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugLinkMemoryRead) Reset()                    { *m = DebugLinkMemoryRead{} }
func (m *DebugLinkMemoryRead) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkMemoryRead) ProtoMessage()               {}
func (*DebugLinkMemoryRead) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{69} }

func (m *DebugLinkMemoryRead) GetAddress() uint32 {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return 0
}

func (m *DebugLinkMemoryRead) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type DebugLinkMemory struct {
	Memory           []byte `protobuf:"bytes,1,opt,name=memory" json:"memory,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugLinkMemory) Reset()                    { *m = DebugLinkMemory{} }
func (m *DebugLinkMemory) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkMemory) ProtoMessage()               {}
func (*DebugLinkMemory) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{70} }

func (m *DebugLinkMemory) GetMemory() []byte {
	if m != nil {
		return m.Memory
	}
	return nil
}

type DebugLinkMemoryWrite struct {
	Address          *uint32 `protobuf:"varint,1,opt,name=address" json:"address,omitempty"`
	Memory           []byte  `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	Flash            *bool   `protobuf:"varint,3,opt,name=flash" json:"flash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugLinkMemoryWrite) Reset()                    { *m = DebugLinkMemoryWrite{} }
func (m *DebugLinkMemoryWrite) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkMemoryWrite) ProtoMessage()               {}
func (*DebugLinkMemoryWrite) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{71} }

func (m *DebugLinkMemoryWrite) GetAddress() uint32 {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return 0
}

func (m *DebugLinkMemoryWrite) GetMemory() []byte {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *DebugLinkMemoryWrite) GetFlash() bool {
	if m != nil && m.Flash != nil {
		return *m.Flash
	}
	return false
}

type DebugLinkFlashErase struct {
	Sector           *uint32 `protobuf:"varint,1,opt,name=sector" json:"sector,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugLinkFlashErase) Reset()                    { *m = DebugLinkFlashErase{} }
func (m *DebugLinkFlashErase) String() string            { return proto.CompactTextString(m) }
func (*DebugLinkFlashErase) ProtoMessage()               {}
func (*DebugLinkFlashErase) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{72} }

func (m *DebugLinkFlashErase) GetSector() uint32 {
	if m != nil && m.Sector != nil {
		return *m.Sector
	}
	return 0
}

func init() {
	proto.RegisterType((*Initialize)(nil), "Initialize")
	proto.RegisterType((*GetFeatures)(nil), "GetFeatures")
	proto.RegisterType((*Features)(nil), "Features")
	proto.RegisterType((*ClearSession)(nil), "ClearSession")
	proto.RegisterType((*ApplySettings)(nil), "ApplySettings")
	proto.RegisterType((*ApplyFlags)(nil), "ApplyFlags")
	proto.RegisterType((*ChangePin)(nil), "ChangePin")
	proto.RegisterType((*Ping)(nil), "Ping")
	proto.RegisterType((*Success)(nil), "Success")
	proto.RegisterType((*Failure)(nil), "Failure")
	proto.RegisterType((*ButtonRequest)(nil), "ButtonRequest")
	proto.RegisterType((*ButtonAck)(nil), "ButtonAck")
	proto.RegisterType((*PinMatrixRequest)(nil), "PinMatrixRequest")
	proto.RegisterType((*PinMatrixAck)(nil), "PinMatrixAck")
	proto.RegisterType((*Cancel)(nil), "Cancel")
	proto.RegisterType((*PassphraseRequest)(nil), "PassphraseRequest")
	proto.RegisterType((*PassphraseAck)(nil), "PassphraseAck")
	proto.RegisterType((*GetEntropy)(nil), "GetEntropy")
	proto.RegisterType((*Entropy)(nil), "Entropy")
	proto.RegisterType((*GetPublicKey)(nil), "GetPublicKey")
	proto.RegisterType((*PublicKey)(nil), "PublicKey")
	proto.RegisterType((*GetAddress)(nil), "GetAddress")
	proto.RegisterType((*EPVchainGetAddress)(nil), "EPVchainGetAddress")
	proto.RegisterType((*Address)(nil), "Address")
	proto.RegisterType((*EPVchainAddress)(nil), "EPVchainAddress")
	proto.RegisterType((*WipeDevice)(nil), "WipeDevice")
	proto.RegisterType((*LoadDevice)(nil), "LoadDevice")
	proto.RegisterType((*ResetDevice)(nil), "ResetDevice")
	proto.RegisterType((*BackupDevice)(nil), "BackupDevice")
	proto.RegisterType((*EntropyRequest)(nil), "EntropyRequest")
	proto.RegisterType((*EntropyAck)(nil), "EntropyAck")
	proto.RegisterType((*RecoveryDevice)(nil), "RecoveryDevice")
	proto.RegisterType((*WordRequest)(nil), "WordRequest")
	proto.RegisterType((*WordAck)(nil), "WordAck")
	proto.RegisterType((*SignMessage)(nil), "SignMessage")
	proto.RegisterType((*VerifyMessage)(nil), "VerifyMessage")
	proto.RegisterType((*MessageSignature)(nil), "MessageSignature")
	proto.RegisterType((*EncryptMessage)(nil), "EncryptMessage")
	proto.RegisterType((*EncryptedMessage)(nil), "EncryptedMessage")
	proto.RegisterType((*DecryptMessage)(nil), "DecryptMessage")
	proto.RegisterType((*DecryptedMessage)(nil), "DecryptedMessage")
	proto.RegisterType((*CipherKeyValue)(nil), "CipherKeyValue")
	proto.RegisterType((*CipheredKeyValue)(nil), "CipheredKeyValue")
	proto.RegisterType((*EstimateTxSize)(nil), "EstimateTxSize")
	proto.RegisterType((*TxSize)(nil), "TxSize")
	proto.RegisterType((*SignTx)(nil), "SignTx")
	proto.RegisterType((*SimpleSignTx)(nil), "SimpleSignTx")
	proto.RegisterType((*TxRequest)(nil), "TxRequest")
	proto.RegisterType((*TxAck)(nil), "TxAck")
	proto.RegisterType((*EPVchainSignTx)(nil), "EPVchainSignTx")
	proto.RegisterType((*EPVchainTxRequest)(nil), "EPVchainTxRequest")
	proto.RegisterType((*EPVchainTxAck)(nil), "EPVchainTxAck")
	proto.RegisterType((*EPVchainSignMessage)(nil), "EPVchainSignMessage")
	proto.RegisterType((*EPVchainVerifyMessage)(nil), "EPVchainVerifyMessage")
	proto.RegisterType((*EPVchainMessageSignature)(nil), "EPVchainMessageSignature")
	proto.RegisterType((*SignIdentity)(nil), "SignIdentity")
	proto.RegisterType((*SignedIdentity)(nil), "SignedIdentity")
	proto.RegisterType((*GetECDHSessionKey)(nil), "GetECDHSessionKey")
	proto.RegisterType((*ECDHSessionKey)(nil), "ECDHSessionKey")
	proto.RegisterType((*SetU2FCounter)(nil), "SetU2FCounter")
	proto.RegisterType((*FirmwareErase)(nil), "FirmwareErase")
	proto.RegisterType((*FirmwareRequest)(nil), "FirmwareRequest")
	proto.RegisterType((*FirmwareUpload)(nil), "FirmwareUpload")
	proto.RegisterType((*SelfTest)(nil), "SelfTest")
	proto.RegisterType((*DebugLinkDecision)(nil), "DebugLinkDecision")
	proto.RegisterType((*DebugLinkGetState)(nil), "DebugLinkGetState")
	proto.RegisterType((*DebugLinkState)(nil), "DebugLinkState")
	proto.RegisterType((*DebugLinkStop)(nil), "DebugLinkStop")
	proto.RegisterType((*DebugLinkLog)(nil), "DebugLinkLog")
	proto.RegisterType((*DebugLinkMemoryRead)(nil), "DebugLinkMemoryRead")
	proto.RegisterType((*DebugLinkMemory)(nil), "DebugLinkMemory")
	proto.RegisterType((*DebugLinkMemoryWrite)(nil), "DebugLinkMemoryWrite")
	proto.RegisterType((*DebugLinkFlashErase)(nil), "DebugLinkFlashErase")
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x5a, 0xcb, 0x6f, 0xdc, 0x46,
	0x9a, 0x5f, 0x76, 0xb7, 0xfa, 0xf1, 0x35, 0xbb, 0x55, 0xa2, 0x2d, 0xbb, 0x2d, 0x5b, 0xb6, 0x4c,
	0xc9, 0xb6, 0x64, 0x27, 0xed, 0x44, 0x79, 0x6c, 0xd6, 0xbb, 0x79, 0xc8, 0x7a, 0xd8, 0xde, 0xd8,
	0x8e, 0xc0, 0x56, 0x9c, 0xdb, 0x12, 0x14, 0x59, 0xea, 0xae, 0x55, 0x37, 0xc9, 0xf0, 0xa1, 0xa8,
	0x7d, 0xd8, 0xeb, 0xee, 0x65, 0x81, 0xec, 0x69, 0x73, 0x1a, 0xe4, 0x36, 0x19, 0x04, 0x18, 0x0c,
	0x30, 0x18, 0x60, 0x72, 0x9a, 0x3f, 0x60, 0xfe, 0x8b, 0x39, 0xce, 0x1f, 0x30, 0xe7, 0x41, 0x3d,
	0x48, 0x16, 0x29, 0xb6, 0x6c, 0x27, 0xc0, 0x5c, 0x04, 0xd6, 0x57, 0xbf, 0xfe, 0xea, 0x7b, 0xd5,
	0x57, 0x5f, 0x7d, 0x25, 0xe8, 0x4e, 0x70, 0x18, 0x5a, 0x43, 0x1c, 0xf6, 0xfd, 0xc0, 0x8b, 0xbc,
	0xa5, 0x76, 0x34, 0xf5, 0x93, 0x81, 0xae, 0x02, 0x3c, 0x71, 0x49, 0x44, 0xac, 0x31, 0x79, 0x89,
	0xf5, 0x0e, 0xb4, 0x1f, 0xe1, 0x68, 0x0f, 0x5b, 0x51, 0x1c, 0xe0, 0x50, 0xff, 0x69, 0x0e, 0x9a,
	0xc9, 0x40, 0xbb, 0x04, 0xf5, 0x13, 0xec, 0x3a, 0x5e, 0xd0, 0x53, 0x56, 0x94, 0xf5, 0x96, 0x21,
	0x46, 0xda, 0x2a, 0x74, 0x26, 0xd6, 0x7f, 0x7a, 0x81, 0x79, 0x82, 0x83, 0x90, 0x78, 0x6e, 0xaf,
	0xb2, 0xa2, 0xac, 0x77, 0x0c, 0x95, 0x11, 0x5f, 0x70, 0x1a, 0x03, 0x11, 0x57, 0x02, 0x55, 0x05,
	0x88, 0x12, 0x25, 0x90, 0x6f, 0x45, 0xf6, 0x28, 0x05, 0xd5, 0x38, 0x88, 0x11, 0x13, 0xd0, 0x1d,
	0x98, 0x3f, 0xf4, 0xbc, 0x68, 0xec, 0x59, 0x0e, 0x0e, 0xcc, 0x89, 0xe7, 0xe0, 0xde, 0xdc, 0x8a,
	0xb2, 0xde, 0x34, 0xba, 0x19, 0xf9, 0x99, 0xe7, 0x60, 0xed, 0x2a, 0xb4, 0x1c, 0x7c, 0x42, 0x6c,
	0x6c, 0x12, 0xa7, 0x57, 0x67, 0x22, 0x37, 0x39, 0xe1, 0x89, 0xa3, 0xdd, 0x82, 0xae, 0x4f, 0x5c,
	0x93, 0xda, 0x00, 0xdb, 0x11, 0x5d, 0xab, 0xc1, 0x98, 0x74, 0x7c, 0xe2, 0xee, 0xa7, 0x44, 0xed,
	0x3d, 0x58, 0xf4, 0xad, 0x30, 0xf4, 0x47, 0x81, 0x15, 0x62, 0x19, 0xdd, 0x64, 0xe8, 0x8b, 0xd9,
	0xa4, 0xf4, 0xa3, 0x25, 0x68, 0x8e, 0x2d, 0x77, 0x18, 0x5b, 0x43, 0xdc, 0x6b, 0xf1, 0x75, 0x93,
	0xb1, 0x76, 0x11, 0xe6, 0xc6, 0xd6, 0x21, 0x1e, 0xf7, 0x80, 0x4d, 0xf0, 0x81, 0x76, 0x03, 0xe6,
	0x6c, 0x8f, 0xb8, 0x61, 0xaf, 0xbd, 0x52, 0x5d, 0x6f, 0x6f, 0xb6, 0xfa, 0xdb, 0x1e, 0x71, 0x0f,
	0xa6, 0x3e, 0x36, 0x38, 0x5d, 0x5b, 0x81, 0x36, 0x49, 0xbd, 0xe4, 0xf4, 0x54, 0xb6, 0xba, 0x4c,
	0xa2, 0x8b, 0x06, 0xf8, 0x84, 0x30, 0xb3, 0x75, 0x56, 0x94, 0x75, 0xd5, 0x48, 0xc7, 0x05, 0x93,
	0x8d, 0xac, 0x70, 0xd4, 0xeb, 0x32, 0x88, 0x64, 0xb2, 0xc7, 0x56, 0x38, 0xa2, 0x4c, 0xc8, 0xc4,
	0xf7, 0x82, 0x08, 0x3b, 0xbd, 0x79, 0xb6, 0x46, 0x3a, 0xd6, 0x96, 0x01, 0xa8, 0xc5, 0x6c, 0xcb,
	0x1e, 0x61, 0xa7, 0x87, 0xd8, 0x6c, 0xcb, 0x27, 0xee, 0x36, 0x23, 0x68, 0xf7, 0x60, 0x41, 0xb2,
	0x94, 0x40, 0x2d, 0x30, 0x14, 0xca, 0x26, 0x04, 0x78, 0x03, 0xd0, 0x11, 0x09, 0x26, 0xdf, 0x58,
	0x01, 0x35, 0x2a, 0x0e, 0xb1, 0x1b, 0xf5, 0x34, 0x86, 0x9d, 0x4f, 0xe8, 0xfb, 0x9c, 0xac, 0xdd,
	0x04, 0xd5, 0xc5, 0xd8, 0x09, 0xcd, 0x43, 0xcb, 0x3e, 0x8e, 0xfd, 0xde, 0x05, 0xae, 0x3a, 0xa3,
	0x3d, 0x64, 0x24, 0x6a, 0xd3, 0xa3, 0xb1, 0x35, 0x0c, 0x7b, 0x17, 0x59, 0xb8, 0xf0, 0x81, 0xde,
	0x05, 0x75, 0x7b, 0x8c, 0xad, 0x60, 0x80, 0x43, 0x6a, 0x04, 0xfd, 0x7f, 0x14, 0xe8, 0x6c, 0xf9,
	0xfe, 0x78, 0x3a, 0xc0, 0x51, 0x44, 0xdc, 0x61, 0x98, 0xf3, 0x93, 0x32, 0xcb, 0x4f, 0x15, 0xd9,
	0x4f, 0xb7, 0xa0, 0x1b, 0xd3, 0x38, 0x48, 0xf5, 0x61, 0x61, 0xdc, 0x34, 0x3a, 0x71, 0x88, 0xf7,
	0x53, 0xa2, 0x76, 0x1d, 0x60, 0xe4, 0x4d, 0x70, 0x68, 0x07, 0x18, 0xf3, 0x20, 0x56, 0x0d, 0x89,
	0xa2, 0xeb, 0x00, 0x4c, 0x92, 0x3d, 0x2a, 0x68, 0x26, 0xbe, 0x22, 0x8b, 0xbf, 0x0a, 0xad, 0xed,
	0x91, 0xe5, 0x0e, 0xf1, 0x3e, 0x71, 0xe9, 0xd6, 0x0b, 0xf0, 0xc4, 0x3b, 0xe1, 0x72, 0x36, 0x0d,
	0x31, 0xd2, 0x7f, 0xa3, 0x40, 0x6d, 0x9f, 0xb8, 0x43, 0xad, 0x07, 0x0d, 0xb1, 0xc9, 0x85, 0x26,
	0xc9, 0x90, 0xfa, 0xe5, 0x30, 0x8e, 0x22, 0x2f, 0x17, 0xeb, 0x15, 0xee, 0x17, 0x3e, 0x21, 0x45,
	0xee, 0xd9, 0x5d, 0x51, 0x7d, 0xa3, 0x5d, 0x51, 0x9b, 0xbd, 0x2b, 0xf4, 0x55, 0x68, 0x0c, 0x62,
	0xdb, 0xc6, 0x61, 0x38, 0x5b, 0x5a, 0x7d, 0x17, 0x1a, 0x7b, 0x16, 0x19, 0xc7, 0x01, 0xd6, 0x56,
	0xa0, 0x66, 0xd3, 0xcd, 0x4d, 0x11, 0xdd, 0x4d, 0xb5, 0x2f, 0xe8, 0x6c, 0x57, 0xb0, 0x19, 0x99,
	0x4d, 0x25, 0xcf, 0xe6, 0x73, 0xe8, 0x3c, 0x64, 0xba, 0x19, 0xf8, 0xeb, 0x18, 0x87, 0x91, 0x76,
	0x3b, 0xc7, 0x4c, 0xeb, 0xe7, 0x66, 0x25, 0x96, 0x1a, 0xd4, 0x1c, 0x2b, 0xb2, 0x04, 0x3f, 0xf6,
	0xad, 0xb7, 0xa1, 0xc5, 0xe1, 0x5b, 0xf6, 0xb1, 0xfe, 0x31, 0xa0, 0x7d, 0xe2, 0x3e, 0xb3, 0xa2,
	0x80, 0x9c, 0x26, 0xcc, 0x37, 0xa0, 0x46, 0x33, 0xaa, 0x60, 0xbe, 0xd8, 0x2f, 0x02, 0x38, 0x7f,
	0x0a, 0xd1, 0x57, 0x40, 0x4d, 0x67, 0xb7, 0xec, 0x63, 0x0d, 0x41, 0xd5, 0x27, 0x6e, 0x4f, 0x59,
	0xa9, 0xac, 0xb7, 0x0c, 0xfa, 0xa9, 0x37, 0xa1, 0xbe, 0x6d, 0xb9, 0x36, 0x1e, 0xeb, 0x17, 0x60,
	0x21, 0x8b, 0x29, 0xc1, 0x4a, 0xbf, 0x0f, 0x9d, 0x8c, 0x48, 0x39, 0x5c, 0x07, 0x90, 0xc2, 0x91,
	0x33, 0x92, 0x28, 0xfa, 0x0a, 0xc0, 0x23, 0x1c, 0xed, 0xba, 0x51, 0xe0, 0xf9, 0x53, 0xaa, 0x5f,
	0x48, 0x5e, 0x72, 0x5c, 0xc7, 0x60, 0xdf, 0xd4, 0x31, 0xc9, 0x74, 0x0f, 0x1a, 0x98, 0x7f, 0x32,
	0x84, 0x6a, 0x24, 0x43, 0xfd, 0x57, 0x0a, 0xa8, 0x8f, 0x70, 0xb4, 0x1f, 0x1f, 0x8e, 0x89, 0xfd,
	0x39, 0x9e, 0xd2, 0xec, 0x6a, 0x39, 0x4e, 0x80, 0xc3, 0xd0, 0xa4, 0xf2, 0x57, 0xd7, 0x3b, 0x46,
	0x53, 0x10, 0x9e, 0x6b, 0xeb, 0x80, 0xb0, 0xed, 0x84, 0x96, 0x69, 0xc7, 0xc1, 0x09, 0x36, 0x5d,
	0x6b, 0x92, 0xb8, 0xa8, 0xcb, 0xe8, 0xdb, 0x94, 0xfc, 0xdc, 0x9a, 0x60, 0xba, 0xbd, 0xc3, 0x91,
	0xf7, 0x8d, 0xe9, 0x90, 0xd0, 0x1f, 0x5b, 0x53, 0x11, 0x6f, 0x6d, 0x4a, 0xdb, 0xe1, 0x24, 0x6d,
	0x0d, 0x5a, 0x34, 0x09, 0x72, 0x2e, 0x34, 0xc2, 0x5a, 0x0f, 0x1a, 0x0f, 0x49, 0x44, 0x69, 0x46,
	0x93, 0xfe, 0xa5, 0x8c, 0xf4, 0xcf, 0xa0, 0x95, 0x09, 0x77, 0x03, 0x6a, 0x2e, 0x77, 0x77, 0x65,
	0xbd, 0xbd, 0xd9, 0xee, 0x3f, 0xde, 0x79, 0xee, 0x39, 0x22, 0x74, 0x5c, 0xe1, 0xe7, 0x53, 0x3f,
	0x3e, 0x4c, 0xfc, 0x4c, 0xbf, 0xf5, 0xbf, 0x2a, 0xcc, 0x54, 0x5b, 0x5c, 0x89, 0xf3, 0x15, 0xcc,
	0xc9, 0x54, 0x99, 0x21, 0xd3, 0xeb, 0x28, 0xf7, 0x01, 0x34, 0x27, 0xf1, 0x38, 0x22, 0x21, 0x19,
	0x32, 0xdd, 0xda, 0x9b, 0x57, 0xfa, 0xcf, 0x04, 0xc1, 0xc0, 0x0e, 0xc6, 0x93, 0x81, 0x1d, 0x10,
	0x9f, 0xc7, 0x50, 0x0a, 0xd5, 0x3e, 0x85, 0x76, 0xc8, 0xe8, 0x26, 0x8b, 0xbc, 0x39, 0x16, 0x79,
	0xa8, 0xff, 0xc4, 0xf5, 0xe3, 0x28, 0xfb, 0xc1, 0x03, 0x75, 0xb0, 0xbf, 0xfb, 0x7c, 0x67, 0x6b,
	0x67, 0xc7, 0xd8, 0x1d, 0x0c, 0x0c, 0x08, 0xd3, 0x19, 0xfd, 0x00, 0xb4, 0xdd, 0x68, 0x84, 0x03,
	0x1c, 0x4f, 0x5e, 0x57, 0xe7, 0xa2, 0x36, 0x95, 0x33, 0xda, 0xd0, 0x50, 0x4a, 0x58, 0xf5, 0xa0,
	0x21, 0x7e, 0x29, 0x82, 0x32, 0x19, 0xea, 0xf7, 0x60, 0x3e, 0x59, 0x7a, 0x06, 0x58, 0xcd, 0xc0,
	0x2a, 0xc0, 0x57, 0xc4, 0xc7, 0x3b, 0xec, 0xdc, 0xd6, 0xff, 0xaf, 0x02, 0xf0, 0xd4, 0xb3, 0x1c,
	0x3e, 0xa4, 0x09, 0x7c, 0xe2, 0xe2, 0x89, 0xe7, 0x12, 0x3b, 0x49, 0xe0, 0xc9, 0x38, 0x0d, 0x81,
	0x0a, 0x33, 0x6a, 0x49, 0x08, 0x88, 0xad, 0x57, 0x65, 0xbf, 0xa3, 0x9f, 0x3f, 0x2b, 0xad, 0x69,
	0xab, 0xd2, 0x21, 0x32, 0xc7, 0x03, 0x01, 0xbb, 0xc3, 0x31, 0x09, 0x47, 0x65, 0xa7, 0x49, 0x5d,
	0x3e, 0x4d, 0x56, 0xa1, 0x13, 0x1e, 0x13, 0xdf, 0xb4, 0x47, 0xd8, 0x3e, 0x0e, 0xe3, 0x89, 0x28,
	0x41, 0x54, 0x4a, 0xdc, 0x16, 0x34, 0xed, 0x06, 0xb4, 0xe3, 0xcd, 0x23, 0xd3, 0xf6, 0x62, 0x37,
	0xc2, 0x01, 0xab, 0x3b, 0x3a, 0x06, 0xc4, 0x9b, 0x47, 0xdb, 0x9c, 0xa2, 0xff, 0xb6, 0x02, 0x6d,
	0x03, 0x87, 0x38, 0x12, 0x46, 0xb9, 0x05, 0x5d, 0xe1, 0x21, 0x33, 0xb0, 0x5c, 0xc7, 0x9b, 0x88,
	0x33, 0xa3, 0x23, 0xa8, 0x06, 0x23, 0x6a, 0x37, 0xa0, 0x19, 0x46, 0x01, 0x76, 0x87, 0xd1, 0x88,
	0x17, 0x6c, 0x0f, 0xaa, 0x9b, 0x1f, 0x7c, 0x68, 0xa4, 0xc4, 0xd9, 0xd6, 0xa8, 0x9e, 0x63, 0x8d,
	0xb3, 0x07, 0x48, 0xad, 0xec, 0x00, 0xf9, 0x05, 0x46, 0x2b, 0xd8, 0xa3, 0x51, 0xb4, 0x07, 0x05,
	0x30, 0xab, 0x8a, 0x7a, 0x81, 0x17, 0x6a, 0x40, 0x49, 0xbc, 0x5c, 0xa0, 0x85, 0x01, 0xff, 0x12,
	0x41, 0x85, 0xa0, 0x2b, 0xf2, 0x5f, 0x92, 0x64, 0x6f, 0x03, 0x08, 0x0a, 0xcd, 0xb0, 0xb9, 0xa4,
	0xa8, 0xc8, 0x49, 0xf1, 0x4f, 0x15, 0xe8, 0x1a, 0xd8, 0xf6, 0x4e, 0x70, 0x30, 0x15, 0xd6, 0x5f,
	0x06, 0xf8, 0xc6, 0x0b, 0x1c, 0x2e, 0x9f, 0x38, 0xd1, 0x5b, 0x94, 0xc2, 0xc4, 0x9b, 0x6d, 0xd4,
	0xca, 0x1b, 0x19, 0xb5, 0xfa, 0x2a, 0xa3, 0xd6, 0x5e, 0x69, 0xd4, 0x39, 0xd9, 0xa8, 0x1b, 0x80,
	0xb0, 0x7b, 0xe4, 0x05, 0x36, 0x36, 0xa9, 0xac, 0x63, 0x12, 0x46, 0xcc, 0xea, 0x4d, 0x63, 0x5e,
	0xd0, 0xbf, 0x12, 0x64, 0x9a, 0x39, 0x59, 0xca, 0xe1, 0x81, 0xc8, 0xbe, 0x8b, 0x3e, 0x69, 0x9d,
	0xf1, 0xc9, 0x65, 0x68, 0x38, 0xc1, 0xd4, 0x0c, 0x62, 0x97, 0xd5, 0xbd, 0x4d, 0xa3, 0xee, 0x04,
	0x53, 0x23, 0x76, 0xf5, 0xf7, 0xa0, 0x4d, 0x39, 0x27, 0x27, 0xe9, 0x5a, 0xee, 0x24, 0x45, 0x7d,
	0x69, 0x4e, 0x3a, 0x44, 0x97, 0xa1, 0x41, 0x27, 0xa8, 0x6f, 0x34, 0xa8, 0x51, 0x81, 0x45, 0x8a,
	0x61, 0xdf, 0xfa, 0x8f, 0x0a, 0xb4, 0x07, 0x64, 0xe8, 0x3e, 0x13, 0x15, 0xd0, 0xb9, 0x49, 0x2d,
	0x57, 0x43, 0xb0, 0xcc, 0x93, 0x14, 0x4e, 0xb9, 0x14, 0x5f, 0x9d, 0x95, 0xe2, 0x0b, 0x89, 0xb8,
	0xf6, 0xc6, 0x89, 0xf8, 0xbf, 0x15, 0xe8, 0xbc, 0xc0, 0x01, 0x39, 0x9a, 0x26, 0xf2, 0xe6, 0x92,
	0xa1, 0x22, 0x65, 0x4e, 0xed, 0x1a, 0xb4, 0x42, 0x32, 0x74, 0xd9, 0x7d, 0x8c, 0x45, 0x8c, 0x6a,
	0x64, 0x04, 0x59, 0x95, 0x2a, 0x8f, 0xd3, 0x52, 0x55, 0x66, 0x9e, 0xa0, 0xff, 0x0e, 0x48, 0x88,
	0x30, 0x90, 0x79, 0xfe, 0x1c, 0x59, 0xf4, 0x1f, 0x14, 0xba, 0xa9, 0xec, 0x60, 0xea, 0x47, 0x89,
	0x5a, 0x97, 0xa0, 0xee, 0xc7, 0x87, 0xc7, 0x38, 0xd9, 0x45, 0x62, 0x54, 0xac, 0xe2, 0x24, 0xb1,
	0x6f, 0x82, 0x9a, 0x64, 0x32, 0xcf, 0x1d, 0xa7, 0xc7, 0xa7, 0xa0, 0x7d, 0xe1, 0x8e, 0x0b, 0x55,
	0x48, 0xed, 0xbc, 0x43, 0x7a, 0x6e, 0x96, 0xda, 0x2f, 0x00, 0x09, 0x49, 0xb1, 0x93, 0xc8, 0x7a,
	0x11, 0xe6, 0x5c, 0xcf, 0xb5, 0xb1, 0x10, 0x95, 0x0f, 0xce, 0x91, 0x54, 0x83, 0xda, 0x68, 0x62,
	0xd9, 0xc2, 0xee, 0xec, 0x5b, 0xff, 0x1a, 0xba, 0x3b, 0x38, 0x67, 0x81, 0x73, 0x03, 0x31, 0x5d,
	0xb2, 0x32, 0x63, 0xc9, 0x6a, 0xf9, 0x92, 0x35, 0x69, 0xc9, 0x3d, 0x40, 0x62, 0xc9, 0x4c, 0x95,
	0x42, 0xad, 0x2d, 0x71, 0x90, 0x7c, 0x5b, 0xc9, 0xf9, 0x56, 0xff, 0xb3, 0x02, 0xdd, 0x6d, 0xe2,
	0x8f, 0x70, 0xf0, 0x39, 0x9e, 0xbe, 0xb0, 0xc6, 0xf1, 0x2b, 0x64, 0x47, 0x50, 0xa5, 0x7e, 0xe5,
	0x5c, 0xe8, 0x27, 0xd5, 0xe6, 0x84, 0xfe, 0x4e, 0x48, 0xcd, 0x07, 0x3c, 0x93, 0x32, 0xf9, 0xc4,
	0xb1, 0x90, 0x0c, 0xb5, 0x35, 0xe8, 0x5a, 0xe1, 0xb1, 0xe9, 0xb9, 0x66, 0x02, 0xe0, 0x77, 0x7a,
	0xd5, 0x0a, 0x8f, 0xbf, 0x70, 0x77, 0xcf, 0xa0, 0x1c, 0xae, 0xa6, 0x48, 0x52, 0x1c, 0x25, 0x54,
	0xd7, 0xba, 0x50, 0x21, 0x27, 0xec, 0x60, 0x50, 0x8d, 0x0a, 0x39, 0xd1, 0xd7, 0x01, 0x71, 0x65,
	0xb0, 0x93, 0xaa, 0x93, 0xca, 0xa7, 0x48, 0xf2, 0xe9, 0xff, 0x05, 0xdd, 0xdd, 0x30, 0x22, 0x13,
	0x2b, 0xc2, 0x07, 0xa7, 0x03, 0xf2, 0x12, 0xd3, 0x23, 0xda, 0x8b, 0x23, 0x3f, 0x8e, 0xc2, 0x34,
	0xa3, 0xd3, 0xc2, 0x59, 0x15, 0x44, 0x9e, 0xd4, 0x6f, 0x82, 0x4a, 0x5c, 0x09, 0x53, 0x61, 0x98,
	0x36, 0xa7, 0x71, 0xc8, 0x6b, 0x25, 0x13, 0xfd, 0x26, 0xd4, 0xc5, 0xba, 0x97, 0xa1, 0x11, 0x9d,
	0x9a, 0xa2, 0x54, 0xa7, 0xd9, 0xb4, 0x1e, 0xb1, 0x09, 0xfd, 0xf7, 0x0a, 0xd4, 0xe9, 0xf6, 0x3c,
	0x38, 0xfd, 0xc7, 0xca, 0xa6, 0x5d, 0x85, 0x46, 0xae, 0x2b, 0xf3, 0x40, 0x79, 0xd7, 0x48, 0x28,
	0xda, 0x75, 0x68, 0x8d, 0x3d, 0xfb, 0xd8, 0x8c, 0x88, 0xd8, 0x69, 0x9d, 0x07, 0xca, 0x3b, 0x46,
	0x93, 0xd2, 0x0e, 0xc8, 0x04, 0xeb, 0x7f, 0x53, 0x40, 0x1d, 0x90, 0x89, 0x3f, 0xc6, 0x42, 0xf6,
	0x35, 0xa8, 0x73, 0x11, 0x58, 0x2c, 0xb5, 0x37, 0xd5, 0xfe, 0xc1, 0x29, 0xcb, 0x99, 0x2c, 0xcd,
	0x8b, 0x39, 0xed, 0x0e, 0x34, 0x84, 0x32, 0xbd, 0x0a, 0x83, 0x75, 0xfa, 0x07, 0xa7, 0x5f, 0x30,
	0x0a, 0xc3, 0x25, 0xb3, 0xda, 0xfb, 0xa0, 0x46, 0x81, 0xe5, 0x86, 0x16, 0x3b, 0x09, 0xc3, 0x5e,
	0x95, 0xa1, 0x51, 0xff, 0x20, 0x23, 0xb2, 0x1f, 0xe4, 0x50, 0xaf, 0x97, 0x16, 0x65, 0xc5, 0xe7,
	0xce, 0x57, 0xbc, 0x7e, 0x56, 0xf1, 0x5f, 0x2b, 0xd0, 0x3a, 0x48, 0x2f, 0x8a, 0xf7, 0x41, 0x0d,
	0xf8, 0xa7, 0x29, 0x1d, 0x73, 0x6a, 0x5f, 0x3e, 0xe2, 0xda, 0x41, 0x36, 0xd0, 0xee, 0x43, 0xc3,
	0xc1, 0x91, 0x45, 0xc6, 0xa1, 0xa8, 0x63, 0x17, 0xfb, 0x29, 0xb7, 0x1d, 0x3e, 0xc1, 0x0d, 0x21,
	0x50, 0xda, 0x47, 0x00, 0x21, 0x0e, 0x92, 0x36, 0x51, 0x95, 0xfd, 0xa6, 0x97, 0xfd, 0x66, 0x90,
	0xce, 0xb1, 0x9f, 0x49, 0x58, 0x7d, 0x03, 0xe6, 0x0e, 0xd8, 0x95, 0x74, 0x05, 0x2a, 0xd1, 0x29,
	0x13, 0xad, 0xcc, 0x82, 0x95, 0xe8, 0x54, 0xff, 0xdf, 0x0a, 0x74, 0x93, 0x0a, 0x5e, 0xf8, 0xf3,
	0x67, 0xa4, 0xb6, 0xab, 0xd0, 0x1a, 0x5a, 0xa1, 0xe9, 0x07, 0xc4, 0x4e, 0xd2, 0x44, 0x73, 0x68,
	0x85, 0xfb, 0x74, 0x9c, 0x4c, 0x8e, 0xc9, 0x84, 0x44, 0x22, 0xc5, 0xd1, 0xc9, 0xa7, 0x74, 0x4c,
	0x37, 0x78, 0xe4, 0x31, 0x67, 0xa8, 0x46, 0x25, 0xf2, 0xb2, 0xcd, 0x5c, 0x97, 0x93, 0xcd, 0x5b,
	0xa0, 0xd1, 0xeb, 0xbb, 0x29, 0x9a, 0x64, 0xa6, 0x3d, 0x8a, 0xdd, 0x63, 0x91, 0x16, 0x10, 0x9d,
	0x11, 0x6d, 0xcf, 0x6d, 0x4a, 0xa7, 0x25, 0x0c, 0x43, 0x8f, 0x79, 0x45, 0x2c, 0xca, 0x6c, 0x4a,
	0x7a, 0xca, 0xcb, 0xe1, 0x2b, 0xd0, 0xb4, 0x47, 0x16, 0x71, 0x4d, 0xe2, 0x88, 0x02, 0xa7, 0xc1,
	0xc6, 0x4f, 0x1c, 0xfd, 0xff, 0x15, 0x58, 0x48, 0xec, 0x91, 0x39, 0xbb, 0xc0, 0x51, 0x39, 0xc3,
	0x91, 0x16, 0xaa, 0xc9, 0x81, 0x69, 0x9e, 0x88, 0xae, 0x29, 0xa4, 0xa4, 0x17, 0x79, 0x40, 0x20,
	0x6c, 0x94, 0x01, 0x8c, 0x3c, 0x20, 0x4c, 0x1a, 0x4d, 0x29, 0x69, 0xa0, 0xf7, 0xa1, 0x93, 0x09,
	0x46, 0x9d, 0xbb, 0x0c, 0x4c, 0x02, 0x61, 0x0c, 0x9e, 0xfc, 0x5a, 0x94, 0xc2, 0xac, 0xa0, 0x3f,
	0x85, 0x0b, 0xb2, 0x63, 0x7f, 0x59, 0x05, 0xa5, 0x13, 0x58, 0x4c, 0xb8, 0x9d, 0x5b, 0xe1, 0xa8,
	0xbf, 0xb8, 0xc2, 0xd1, 0x0d, 0xe8, 0x25, 0x4b, 0xbd, 0xaa, 0x86, 0x79, 0xdd, 0xd5, 0xf4, 0x9f,
	0x58, 0xd2, 0x1a, 0xba, 0x4f, 0x1c, 0xec, 0x46, 0x24, 0x9a, 0x6a, 0x1b, 0xd0, 0x24, 0xe2, 0x5b,
	0xec, 0x8f, 0x4e, 0x3f, 0x99, 0xe4, 0xf7, 0x73, 0x92, 0x41, 0x91, 0x3d, 0xb2, 0xc6, 0xd4, 0xf7,
	0xd8, 0x1c, 0x11, 0xc7, 0xc1, 0xae, 0x58, 0x60, 0x3e, 0xa5, 0x3f, 0x66, 0xe4, 0x3c, 0xf4, 0x84,
	0x84, 0xb1, 0x35, 0x16, 0x97, 0xd2, 0x0c, 0xfa, 0x82, 0x91, 0x4b, 0xdb, 0x2a, 0xb5, 0xb2, 0xb6,
	0x8a, 0x3e, 0x84, 0x2e, 0x15, 0x1d, 0x3b, 0xa9, 0xf0, 0xb3, 0x2b, 0xb9, 0x65, 0x00, 0x9f, 0x75,
	0x4e, 0xcc, 0xe4, 0x10, 0x57, 0x8d, 0x96, 0x9f, 0xf6, 0x52, 0x72, 0x46, 0xaa, 0x16, 0x8d, 0xf4,
	0xad, 0x02, 0x0b, 0x8f, 0x70, 0xb4, 0xbb, 0xbd, 0xf3, 0x58, 0x34, 0x5a, 0xe9, 0x6f, 0xde, 0xc0,
	0x52, 0xb7, 0x61, 0xde, 0xc7, 0x38, 0x30, 0xcf, 0x88, 0xd0, 0xa1, 0xe4, 0xac, 0xa5, 0x53, 0xa6,
	0x7b, 0xb5, 0x54, 0xf7, 0x77, 0xa1, 0x5b, 0x10, 0x87, 0xee, 0x13, 0x3e, 0x32, 0xb3, 0xfa, 0x13,
	0xc2, 0x14, 0xa0, 0xbf, 0x03, 0x9d, 0x01, 0x8e, 0xbe, 0xdc, 0xdc, 0x93, 0x2e, 0x91, 0xf2, 0x8d,
	0x46, 0x39, 0x73, 0xeb, 0xbe, 0x03, 0x9d, 0x3d, 0xd1, 0xa9, 0xde, 0x65, 0x3d, 0xdf, 0x4b, 0x50,
	0xcf, 0xed, 0x74, 0x31, 0xd2, 0xb7, 0x60, 0x3e, 0x01, 0x26, 0x99, 0xe1, 0x12, 0xd4, 0xbd, 0xa3,
	0xa3, 0x10, 0x27, 0xf7, 0x43, 0x31, 0x92, 0x58, 0x54, 0x72, 0x2c, 0x3e, 0x81, 0x6e, 0xc2, 0xe2,
	0x4b, 0x7f, 0xec, 0x59, 0x0e, 0x75, 0xa6, 0x6f, 0x4d, 0xe9, 0x67, 0xd2, 0x2f, 0x11, 0x43, 0x56,
	0x16, 0x5a, 0xe1, 0x48, 0xd8, 0x90, 0x7d, 0xeb, 0x6b, 0xd0, 0x1c, 0xe0, 0xf1, 0xd1, 0x01, 0x5d,
	0x3b, 0xf7, 0x4b, 0x45, 0xfa, 0xa5, 0x7e, 0x17, 0x16, 0x76, 0xf0, 0x61, 0x3c, 0x7c, 0x4a, 0xdc,
	0xe3, 0x1d, 0x6c, 0xf3, 0x97, 0x83, 0x45, 0xa8, 0x4f, 0x71, 0x68, 0xba, 0x1e, 0x5b, 0xa7, 0x69,
	0xcc, 0x4d, 0x71, 0xf8, 0xdc, 0xd3, 0x2f, 0x48, 0xd8, 0x47, 0x38, 0x1a, 0x44, 0x56, 0x84, 0xf5,
	0xbf, 0x54, 0x68, 0xc5, 0x2b, 0xa8, 0x8c, 0xc4, 0x34, 0xb2, 0xa6, 0x5e, 0x1c, 0x25, 0x35, 0x3f,
	0x1f, 0x25, 0xbd, 0x97, 0x4a, 0xd6, 0x7b, 0xb9, 0x04, 0xf5, 0x09, 0xeb, 0x8a, 0x0a, 0xa7, 0x8a,
	0x51, 0xae, 0xc5, 0x53, 0x9b, 0xd1, 0xe2, 0x99, 0x9b, 0xd5, 0xe2, 0x99, 0x79, 0xdb, 0xae, 0x9f,
	0x73, 0xdb, 0x5e, 0x06, 0x08, 0x70, 0x88, 0x23, 0x76, 0x13, 0x66, 0xe7, 0x45, 0xcb, 0x68, 0x31,
	0x0a, 0xbd, 0x74, 0xd2, 0xaa, 0x8b, 0x4f, 0x27, 0x3d, 0x81, 0x26, 0xd3, 0x4c, 0x65, 0xc4, 0xa4,
	0x8f, 0xfa, 0x16, 0x68, 0x81, 0xe8, 0x0b, 0x98, 0x47, 0xd6, 0x31, 0xbf, 0x55, 0x8b, 0xb7, 0x20,
	0x94, 0xcc, 0xec, 0x59, 0xc7, 0xec, 0x5a, 0xad, 0xdd, 0x85, 0x85, 0x14, 0xcd, 0x9a, 0x07, 0xbe,
	0x17, 0xb2, 0x7b, 0x72, 0xc7, 0x98, 0x4f, 0x26, 0x28, 0x70, 0xdf, 0x0b, 0xf5, 0x79, 0xe8, 0x48,
	0x36, 0xf6, 0x7c, 0x7d, 0x1f, 0xd4, 0x94, 0xf0, 0xd4, 0x1b, 0xb2, 0x0b, 0x3e, 0x3e, 0xc1, 0xe3,
	0xe4, 0x35, 0x81, 0x0d, 0xa8, 0x79, 0x0f, 0x63, 0xfb, 0x18, 0x47, 0xc2, 0xe6, 0x62, 0xc4, 0x6e,
	0xf3, 0xf8, 0x34, 0x12, 0x46, 0x67, 0xdf, 0xfa, 0x23, 0xb8, 0x90, 0x72, 0x7c, 0x86, 0x27, 0x5e,
	0x30, 0x35, 0x30, 0x8f, 0x39, 0x39, 0x81, 0x74, 0xb2, 0x04, 0x32, 0x2b, 0x6e, 0x37, 0x60, 0xbe,
	0xc0, 0x88, 0xb9, 0x99, 0x7d, 0x25, 0x01, 0xc1, 0x47, 0xfa, 0x7f, 0xc0, 0xc5, 0x02, 0xf4, 0xab,
	0x80, 0x44, 0xf8, 0xfc, 0x45, 0x05, 0xa7, 0x8a, 0xcc, 0x49, 0xbc, 0xa6, 0x84, 0x23, 0x71, 0x5b,
	0xe4, 0x03, 0xfd, 0x6d, 0x49, 0xa7, 0x3d, 0x4a, 0x49, 0x37, 0x6d, 0x88, 0xed, 0xc8, 0x4b, 0x76,
	0xb8, 0x18, 0xdd, 0xfd, 0x71, 0x11, 0xda, 0xe2, 0x1c, 0x61, 0x75, 0xd8, 0x0a, 0x5c, 0x92, 0x86,
	0x66, 0xf6, 0x60, 0x8a, 0xfe, 0x69, 0xa9, 0xf6, 0xed, 0x1f, 0x7a, 0x8a, 0xb6, 0x94, 0x5e, 0x9e,
	0x19, 0x62, 0x9f, 0xb8, 0x43, 0xa4, 0x88, 0xb9, 0x65, 0xb8, 0x20, 0xcf, 0x89, 0x57, 0x10, 0x54,
	0x59, 0xaa, 0x7d, 0x57, 0x32, 0x2d, 0xde, 0x39, 0x50, 0x55, 0x4c, 0xdf, 0x80, 0x45, 0x79, 0x3a,
	0x7d, 0x14, 0x42, 0x35, 0xc1, 0xbe, 0x20, 0x5c, 0xd6, 0x2e, 0x45, 0x73, 0x02, 0x71, 0x07, 0xae,
	0xe4, 0x56, 0x90, 0x13, 0x17, 0xaa, 0x2f, 0x35, 0x29, 0xe8, 0x8f, 0x14, 0xb8, 0x0e, 0x4b, 0x65,
	0x40, 0x9e, 0x75, 0x50, 0x43, 0x42, 0x6e, 0xc0, 0xd5, 0x32, 0xa4, 0x48, 0x71, 0xa8, 0xb9, 0xd4,
	0xfc, 0x2e, 0x81, 0x16, 0xe4, 0xcb, 0x5e, 0x23, 0x50, 0xab, 0xdc, 0x40, 0xc9, 0x34, 0x08, 0x0b,
	0xe8, 0xd0, 0x2b, 0x30, 0x48, 0x8f, 0x05, 0xd4, 0x16, 0x2c, 0x0a, 0x56, 0xca, 0x00, 0xaa, 0x60,
	0x52, 0x90, 0x22, 0xeb, 0x22, 0xa3, 0x8e, 0x60, 0x71, 0x13, 0x2e, 0xcb, 0x08, 0xa9, 0xa7, 0x8a,
	0xba, 0x02, 0x72, 0x0d, 0xb4, 0x9c, 0x27, 0x59, 0xf1, 0x8b, 0xe6, 0xc5, 0xec, 0x5a, 0x5e, 0x4e,
	0xf9, 0xc2, 0x83, 0xd0, 0x52, 0x9d, 0x62, 0x9a, 0x8a, 0x76, 0x1d, 0x2e, 0xe6, 0x2c, 0x27, 0x9e,
	0xd7, 0xd1, 0x82, 0x10, 0xf4, 0x36, 0x5c, 0x2b, 0x44, 0x52, 0xee, 0x31, 0x09, 0x69, 0x29, 0xae,
	0x57, 0x8a, 0xdb, 0xb2, 0x8f, 0xd1, 0x05, 0xee, 0xa9, 0xdf, 0x95, 0xc8, 0xcc, 0x1f, 0x97, 0xd0,
	0xc5, 0x72, 0xbb, 0xa5, 0xe5, 0x2b, 0x5a, 0x14, 0xcb, 0x5c, 0x85, 0x85, 0x3c, 0x80, 0xf2, 0xbf,
	0x94, 0x6a, 0x9c, 0x8b, 0x97, 0x7c, 0xcf, 0x00, 0x5d, 0x16, 0xa8, 0x82, 0xff, 0xe4, 0x57, 0x59,
	0xd4, 0x13, 0x98, 0xd5, 0x7c, 0x88, 0xe6, 0x1e, 0x6a, 0xd1, 0x95, 0x72, 0x50, 0xee, 0x11, 0x0f,
	0x2d, 0x09, 0x81, 0x57, 0xf3, 0x1a, 0xa5, 0x4f, 0x77, 0xe8, 0xaa, 0x64, 0x94, 0x42, 0x34, 0x64,
	0xaf, 0xb1, 0xe8, 0x5a, 0xf9, 0xae, 0xca, 0x1e, 0x49, 0xd0, 0x72, 0x79, 0xd4, 0x26, 0xd3, 0xd7,
	0xd3, 0xa8, 0xcd, 0xf9, 0x39, 0x39, 0x81, 0xd1, 0x8a, 0xb4, 0x8b, 0x0a, 0x96, 0x91, 0xdb, 0xd2,
	0x48, 0x2f, 0xb7, 0x71, 0xbe, 0x55, 0x8d, 0x56, 0xcb, 0xc3, 0x3b, 0x6b, 0x5f, 0xa3, 0xb5, 0xf2,
	0xf0, 0x96, 0xea, 0x7b, 0x74, 0xbb, 0xdc, 0xbe, 0xb9, 0xa2, 0x1d, 0xdd, 0x11, 0xa0, 0x42, 0x7c,
	0x16, 0xcb, 0x6d, 0xb4, 0x2e, 0x24, 0xba, 0x03, 0xcb, 0xb9, 0xf8, 0x2c, 0x3e, 0x65, 0xa2, 0x8d,
	0x14, 0x78, 0xa5, 0x1c, 0x48, 0xa5, 0xbf, 0x2b, 0x39, 0xed, 0x76, 0xc1, 0x12, 0xb9, 0x56, 0x0d,
	0xba, 0x27, 0xed, 0x30, 0x2d, 0x1f, 0xb2, 0x6c, 0xfe, 0xad, 0xa5, 0xfa, 0x77, 0x7c, 0xbe, 0x60,
	0xd1, 0x7c, 0x07, 0x1f, 0xbd, 0x5d, 0x6e, 0x2f, 0xa9, 0x15, 0x8d, 0xfa, 0xe5, 0x99, 0x5b, 0x34,
	0xa5, 0xd1, 0xfd, 0x72, 0x4b, 0x15, 0x9b, 0x50, 0xe8, 0x9d, 0x74, 0x27, 0x17, 0x3c, 0x2c, 0x77,
	0x0d, 0xd1, 0xbb, 0xa9, 0x5e, 0xeb, 0x79, 0x7e, 0xc5, 0xae, 0x25, 0xda, 0x4c, 0x35, 0x2c, 0x70,
	0xcc, 0xf7, 0x21, 0xd1, 0x7b, 0xb3, 0x38, 0x16, 0x9b, 0x87, 0xe8, 0xfd, 0x94, 0xa3, 0x5e, 0xcc,
	0x6d, 0xd9, 0xbd, 0x08, 0x7d, 0x50, 0x1e, 0xa9, 0xf9, 0x0b, 0x08, 0xfa, 0x50, 0x68, 0x5b, 0xb0,
	0xab, 0xf4, 0xef, 0x46, 0xe8, 0x9f, 0x05, 0xa3, 0x75, 0xb8, 0x9e, 0x53, 0xf4, 0xcc, 0x43, 0x25,
	0xfa, 0x48, 0x20, 0x6f, 0xe5, 0x8f, 0xa1, 0xc2, 0xbb, 0x22, 0xfa, 0x17, 0xb1, 0x66, 0x71, 0x0f,
	0xe5, 0x9a, 0x17, 0xe8, 0x41, 0x7a, 0x4c, 0x2e, 0x97, 0xa1, 0xb2, 0x9c, 0xf8, 0xaf, 0x69, 0x8a,
	0xb9, 0x52, 0x0e, 0xa4, 0xde, 0xff, 0xb7, 0x72, 0x6e, 0x67, 0x2e, 0x49, 0xe8, 0xe3, 0x19, 0x1b,
	0x3c, 0x8f, 0xfa, 0xa4, 0x7c, 0xcd, 0xdc, 0x75, 0x05, 0x7d, 0x2a, 0x58, 0x6d, 0xc0, 0x8d, 0x59,
	0x7a, 0x26, 0x2e, 0xfd, 0x4c, 0x40, 0xef, 0xc1, 0xcd, 0x32, 0x68, 0x7e, 0xcf, 0x6f, 0x09, 0x70,
	0x1f, 0xd6, 0xca, 0xc0, 0x67, 0xf6, 0xfe, 0x43, 0x21, 0xec, 0xbd, 0xbc, 0xee, 0x67, 0xee, 0x15,
	0xc8, 0x59, 0x6a, 0x7e, 0x9f, 0x6c, 0xeb, 0x3b, 0x33, 0xc0, 0xc9, 0xc5, 0x02, 0xe1, 0xa5, 0xda,
	0xf7, 0x25, 0x86, 0xca, 0xdf, 0x35, 0xd0, 0xd1, 0x52, 0xed, 0x87, 0x12, 0x43, 0xe5, 0xaa, 0x65,
	0x34, 0x14, 0xac, 0x0a, 0xe1, 0x2c, 0x57, 0xd0, 0x68, 0x24, 0x18, 0x15, 0x8c, 0x59, 0x52, 0x13,
	0x23, 0x57, 0xb0, 0x2b, 0x84, 0x61, 0x01, 0x8a, 0x3c, 0xc1, 0xf1, 0x2e, 0xac, 0x9c, 0x03, 0x63,
	0x15, 0x2f, 0xf2, 0x05, 0xcb, 0x59, 0xab, 0x67, 0xd5, 0x2b, 0xfa, 0x9a, 0x43, 0x1f, 0xbe, 0x0f,
	0xab, 0xb6, 0x37, 0xe9, 0x87, 0x56, 0xe4, 0x85, 0x23, 0x32, 0xb6, 0x0e, 0xc3, 0x7e, 0x14, 0xe0,
	0x97, 0x5e, 0xd0, 0x1f, 0x93, 0x43, 0xfe, 0x6f, 0x7e, 0x87, 0xf1, 0xd1, 0xc3, 0xce, 0x01, 0x23,
	0x0a, 0xae, 0x7f, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xe4, 0xc0, 0x85, 0x16, 0x28, 0x00, 0x00,
}
