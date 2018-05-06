                                                
                      

                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  
package trezor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

                                                                       
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

                                                                      
                                                                     
                                                                 
                                     
const _ = proto.ProtoPackageIsVersion2                                    

    
                                               
                   
type FailureType int32

const (
	FailureType_Failure_UnexpectedMessage FailureType = 1
	FailureType_Failure_ButtonExpected    FailureType = 2
	FailureType_Failure_DataError         FailureType = 3
	FailureType_Failure_ActionCancelled   FailureType = 4
	FailureType_Failure_PinExpected       FailureType = 5
	FailureType_Failure_PinCancelled      FailureType = 6
	FailureType_Failure_PinInvalid        FailureType = 7
	FailureType_Failure_InvalidSignature  FailureType = 8
	FailureType_Failure_ProcessError      FailureType = 9
	FailureType_Failure_NotEnoughFunds    FailureType = 10
	FailureType_Failure_NotInitialized    FailureType = 11
	FailureType_Failure_FirmwareError     FailureType = 99
)

var FailureType_name = map[int32]string{
	1:  "Failure_UnexpectedMessage",
	2:  "Failure_ButtonExpected",
	3:  "Failure_DataError",
	4:  "Failure_ActionCancelled",
	5:  "Failure_PinExpected",
	6:  "Failure_PinCancelled",
	7:  "Failure_PinInvalid",
	8:  "Failure_InvalidSignature",
	9:  "Failure_ProcessError",
	10: "Failure_NotEnoughFunds",
	11: "Failure_NotInitialized",
	99: "Failure_FirmwareError",
}
var FailureType_value = map[string]int32{
	"Failure_UnexpectedMessage": 1,
	"Failure_ButtonExpected":    2,
	"Failure_DataError":         3,
	"Failure_ActionCancelled":   4,
	"Failure_PinExpected":       5,
	"Failure_PinCancelled":      6,
	"Failure_PinInvalid":        7,
	"Failure_InvalidSignature":  8,
	"Failure_ProcessError":      9,
	"Failure_NotEnoughFunds":    10,
	"Failure_NotInitialized":    11,
	"Failure_FirmwareError":     99,
}

func (x FailureType) Enum() *FailureType {
	p := new(FailureType)
	*p = x
	return p
}
func (x FailureType) String() string {
	return proto.EnumName(FailureType_name, int32(x))
}
func (x *FailureType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FailureType_value, data, "FailureType")
	if err != nil {
		return err
	}
	*x = FailureType(value)
	return nil
}
func (FailureType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

    
                                                           
                        
type OutputScriptType int32

const (
	OutputScriptType_PAYTOADDRESS     OutputScriptType = 0
	OutputScriptType_PAYTOSCRIPTHASH  OutputScriptType = 1
	OutputScriptType_PAYTOMULTISIG    OutputScriptType = 2
	OutputScriptType_PAYTOOPRETURN    OutputScriptType = 3
	OutputScriptType_PAYTOWITNESS     OutputScriptType = 4
	OutputScriptType_PAYTOP2SHWITNESS OutputScriptType = 5
)

var OutputScriptType_name = map[int32]string{
	0: "PAYTOADDRESS",
	1: "PAYTOSCRIPTHASH",
	2: "PAYTOMULTISIG",
	3: "PAYTOOPRETURN",
	4: "PAYTOWITNESS",
	5: "PAYTOP2SHWITNESS",
}
var OutputScriptType_value = map[string]int32{
	"PAYTOADDRESS":     0,
	"PAYTOSCRIPTHASH":  1,
	"PAYTOMULTISIG":    2,
	"PAYTOOPRETURN":    3,
	"PAYTOWITNESS":     4,
	"PAYTOP2SHWITNESS": 5,
}

func (x OutputScriptType) Enum() *OutputScriptType {
	p := new(OutputScriptType)
	*p = x
	return p
}
func (x OutputScriptType) String() string {
	return proto.EnumName(OutputScriptType_name, int32(x))
}
func (x *OutputScriptType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OutputScriptType_value, data, "OutputScriptType")
	if err != nil {
		return err
	}
	*x = OutputScriptType(value)
	return nil
}
func (OutputScriptType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

    
                                                           
                       
type InputScriptType int32

const (
	InputScriptType_SPENDADDRESS     InputScriptType = 0
	InputScriptType_SPENDMULTISIG    InputScriptType = 1
	InputScriptType_EXTERNAL         InputScriptType = 2
	InputScriptType_SPENDWITNESS     InputScriptType = 3
	InputScriptType_SPENDP2SHWITNESS InputScriptType = 4
)

var InputScriptType_name = map[int32]string{
	0: "SPENDADDRESS",
	1: "SPENDMULTISIG",
	2: "EXTERNAL",
	3: "SPENDWITNESS",
	4: "SPENDP2SHWITNESS",
}
var InputScriptType_value = map[string]int32{
	"SPENDADDRESS":     0,
	"SPENDMULTISIG":    1,
	"EXTERNAL":         2,
	"SPENDWITNESS":     3,
	"SPENDP2SHWITNESS": 4,
}

func (x InputScriptType) Enum() *InputScriptType {
	p := new(InputScriptType)
	*p = x
	return p
}
func (x InputScriptType) String() string {
	return proto.EnumName(InputScriptType_name, int32(x))
}
func (x *InputScriptType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InputScriptType_value, data, "InputScriptType")
	if err != nil {
		return err
	}
	*x = InputScriptType(value)
	return nil
}
func (InputScriptType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

    
                                                              
                     
type RequestType int32

const (
	RequestType_TXINPUT     RequestType = 0
	RequestType_TXOUTPUT    RequestType = 1
	RequestType_TXMETA      RequestType = 2
	RequestType_TXFINISHED  RequestType = 3
	RequestType_TXEXTRADATA RequestType = 4
)

var RequestType_name = map[int32]string{
	0: "TXINPUT",
	1: "TXOUTPUT",
	2: "TXMETA",
	3: "TXFINISHED",
	4: "TXEXTRADATA",
}
var RequestType_value = map[string]int32{
	"TXINPUT":     0,
	"TXOUTPUT":    1,
	"TXMETA":      2,
	"TXFINISHED":  3,
	"TXEXTRADATA": 4,
}

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}
func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (x *RequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RequestType_value, data, "RequestType")
	if err != nil {
		return err
	}
	*x = RequestType(value)
	return nil
}
func (RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

    
                         
                         
type ButtonRequestType int32

const (
	ButtonRequestType_ButtonRequest_Other            ButtonRequestType = 1
	ButtonRequestType_ButtonRequest_FeeOverThreshold ButtonRequestType = 2
	ButtonRequestType_ButtonRequest_ConfirmOutput    ButtonRequestType = 3
	ButtonRequestType_ButtonRequest_ResetDevice      ButtonRequestType = 4
	ButtonRequestType_ButtonRequest_ConfirmWord      ButtonRequestType = 5
	ButtonRequestType_ButtonRequest_WipeDevice       ButtonRequestType = 6
	ButtonRequestType_ButtonRequest_ProtectCall      ButtonRequestType = 7
	ButtonRequestType_ButtonRequest_SignTx           ButtonRequestType = 8
	ButtonRequestType_ButtonRequest_FirmwareCheck    ButtonRequestType = 9
	ButtonRequestType_ButtonRequest_Address          ButtonRequestType = 10
	ButtonRequestType_ButtonRequest_PublicKey        ButtonRequestType = 11
)

var ButtonRequestType_name = map[int32]string{
	1:  "ButtonRequest_Other",
	2:  "ButtonRequest_FeeOverThreshold",
	3:  "ButtonRequest_ConfirmOutput",
	4:  "ButtonRequest_ResetDevice",
	5:  "ButtonRequest_ConfirmWord",
	6:  "ButtonRequest_WipeDevice",
	7:  "ButtonRequest_ProtectCall",
	8:  "ButtonRequest_SignTx",
	9:  "ButtonRequest_FirmwareCheck",
	10: "ButtonRequest_Address",
	11: "ButtonRequest_PublicKey",
}
var ButtonRequestType_value = map[string]int32{
	"ButtonRequest_Other":            1,
	"ButtonRequest_FeeOverThreshold": 2,
	"ButtonRequest_ConfirmOutput":    3,
	"ButtonRequest_ResetDevice":      4,
	"ButtonRequest_ConfirmWord":      5,
	"ButtonRequest_WipeDevice":       6,
	"ButtonRequest_ProtectCall":      7,
	"ButtonRequest_SignTx":           8,
	"ButtonRequest_FirmwareCheck":    9,
	"ButtonRequest_Address":          10,
	"ButtonRequest_PublicKey":        11,
}

func (x ButtonRequestType) Enum() *ButtonRequestType {
	p := new(ButtonRequestType)
	*p = x
	return p
}
func (x ButtonRequestType) String() string {
	return proto.EnumName(ButtonRequestType_name, int32(x))
}
func (x *ButtonRequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ButtonRequestType_value, data, "ButtonRequestType")
	if err != nil {
		return err
	}
	*x = ButtonRequestType(value)
	return nil
}
func (ButtonRequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

    
                      
                            
type PinMatrixRequestType int32

const (
	PinMatrixRequestType_PinMatrixRequestType_Current   PinMatrixRequestType = 1
	PinMatrixRequestType_PinMatrixRequestType_NewFirst  PinMatrixRequestType = 2
	PinMatrixRequestType_PinMatrixRequestType_NewSecond PinMatrixRequestType = 3
)

var PinMatrixRequestType_name = map[int32]string{
	1: "PinMatrixRequestType_Current",
	2: "PinMatrixRequestType_NewFirst",
	3: "PinMatrixRequestType_NewSecond",
}
var PinMatrixRequestType_value = map[string]int32{
	"PinMatrixRequestType_Current":   1,
	"PinMatrixRequestType_NewFirst":  2,
	"PinMatrixRequestType_NewSecond": 3,
}

func (x PinMatrixRequestType) Enum() *PinMatrixRequestType {
	p := new(PinMatrixRequestType)
	*p = x
	return p
}
func (x PinMatrixRequestType) String() string {
	return proto.EnumName(PinMatrixRequestType_name, int32(x))
}
func (x *PinMatrixRequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PinMatrixRequestType_value, data, "PinMatrixRequestType")
	if err != nil {
		return err
	}
	*x = PinMatrixRequestType(value)
	return nil
}
func (PinMatrixRequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

    
                                                                     
                                                                  
                                                       
  
                                                                     
                                                                 
  
                          
type RecoveryDeviceType int32

const (
	                                              
	RecoveryDeviceType_RecoveryDeviceType_ScrambledWords RecoveryDeviceType = 0
	RecoveryDeviceType_RecoveryDeviceType_Matrix         RecoveryDeviceType = 1
)

var RecoveryDeviceType_name = map[int32]string{
	0: "RecoveryDeviceType_ScrambledWords",
	1: "RecoveryDeviceType_Matrix",
}
var RecoveryDeviceType_value = map[string]int32{
	"RecoveryDeviceType_ScrambledWords": 0,
	"RecoveryDeviceType_Matrix":         1,
}

func (x RecoveryDeviceType) Enum() *RecoveryDeviceType {
	p := new(RecoveryDeviceType)
	*p = x
	return p
}
func (x RecoveryDeviceType) String() string {
	return proto.EnumName(RecoveryDeviceType_name, int32(x))
}
func (x *RecoveryDeviceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RecoveryDeviceType_value, data, "RecoveryDeviceType")
	if err != nil {
		return err
	}
	*x = RecoveryDeviceType(value)
	return nil
}
func (RecoveryDeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

    
                                
                       
type WordRequestType int32

const (
	WordRequestType_WordRequestType_Plain   WordRequestType = 0
	WordRequestType_WordRequestType_Matrix9 WordRequestType = 1
	WordRequestType_WordRequestType_Matrix6 WordRequestType = 2
)

var WordRequestType_name = map[int32]string{
	0: "WordRequestType_Plain",
	1: "WordRequestType_Matrix9",
	2: "WordRequestType_Matrix6",
}
var WordRequestType_value = map[string]int32{
	"WordRequestType_Plain":   0,
	"WordRequestType_Matrix9": 1,
	"WordRequestType_Matrix6": 2,
}

func (x WordRequestType) Enum() *WordRequestType {
	p := new(WordRequestType)
	*p = x
	return p
}
func (x WordRequestType) String() string {
	return proto.EnumName(WordRequestType_name, int32(x))
}
func (x *WordRequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WordRequestType_value, data, "WordRequestType")
	if err != nil {
		return err
	}
	*x = WordRequestType(value)
	return nil
}
func (WordRequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

    
                                                                 
                                                                                         
                     
                      
                          
                   
type HDNodeType struct {
	Depth            *uint32 `protobuf:"varint,1,req,name=depth" json:"depth,omitempty"`
	Fingerprint      *uint32 `protobuf:"varint,2,req,name=fingerprint" json:"fingerprint,omitempty"`
	ChildNum         *uint32 `protobuf:"varint,3,req,name=child_num,json=childNum" json:"child_num,omitempty"`
	ChainCode        []byte  `protobuf:"bytes,4,req,name=chain_code,json=chainCode" json:"chain_code,omitempty"`
	PrivateKey       []byte  `protobuf:"bytes,5,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	PublicKey        []byte  `protobuf:"bytes,6,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HDNodeType) Reset()                    { *m = HDNodeType{} }
func (m *HDNodeType) String() string            { return proto.CompactTextString(m) }
func (*HDNodeType) ProtoMessage()               {}
func (*HDNodeType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HDNodeType) GetDepth() uint32 {
	if m != nil && m.Depth != nil {
		return *m.Depth
	}
	return 0
}

func (m *HDNodeType) GetFingerprint() uint32 {
	if m != nil && m.Fingerprint != nil {
		return *m.Fingerprint
	}
	return 0
}

func (m *HDNodeType) GetChildNum() uint32 {
	if m != nil && m.ChildNum != nil {
		return *m.ChildNum
	}
	return 0
}

func (m *HDNodeType) GetChainCode() []byte {
	if m != nil {
		return m.ChainCode
	}
	return nil
}

func (m *HDNodeType) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *HDNodeType) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type HDNodePathType struct {
	Node             *HDNodeType `protobuf:"bytes,1,req,name=node" json:"node,omitempty"`
	AddressN         []uint32    `protobuf:"varint,2,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *HDNodePathType) Reset()                    { *m = HDNodePathType{} }
func (m *HDNodePathType) String() string            { return proto.CompactTextString(m) }
func (*HDNodePathType) ProtoMessage()               {}
func (*HDNodePathType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HDNodePathType) GetNode() *HDNodeType {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *HDNodePathType) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

    
                              
                    
type CoinType struct {
	CoinName            *string `protobuf:"bytes,1,opt,name=coin_name,json=coinName" json:"coin_name,omitempty"`
	CoinShortcut        *string `protobuf:"bytes,2,opt,name=coin_shortcut,json=coinShortcut" json:"coin_shortcut,omitempty"`
	AddressType         *uint32 `protobuf:"varint,3,opt,name=address_type,json=addressType,def=0" json:"address_type,omitempty"`
	MaxfeeKb            *uint64 `protobuf:"varint,4,opt,name=maxfee_kb,json=maxfeeKb" json:"maxfee_kb,omitempty"`
	AddressTypeP2Sh     *uint32 `protobuf:"varint,5,opt,name=address_type_p2sh,json=addressTypeP2sh,def=5" json:"address_type_p2sh,omitempty"`
	SignedMessageHeader *string `protobuf:"bytes,8,opt,name=signed_message_header,json=signedMessageHeader" json:"signed_message_header,omitempty"`
	XpubMagic           *uint32 `protobuf:"varint,9,opt,name=xpub_magic,json=xpubMagic,def=76067358" json:"xpub_magic,omitempty"`
	XprvMagic           *uint32 `protobuf:"varint,10,opt,name=xprv_magic,json=xprvMagic,def=76066276" json:"xprv_magic,omitempty"`
	Segwit              *bool   `protobuf:"varint,11,opt,name=segwit" json:"segwit,omitempty"`
	Forkid              *uint32 `protobuf:"varint,12,opt,name=forkid" json:"forkid,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *CoinType) Reset()                    { *m = CoinType{} }
func (m *CoinType) String() string            { return proto.CompactTextString(m) }
func (*CoinType) ProtoMessage()               {}
func (*CoinType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_CoinType_AddressType uint32 = 0
const Default_CoinType_AddressTypeP2Sh uint32 = 5
const Default_CoinType_XpubMagic uint32 = 76067358
const Default_CoinType_XprvMagic uint32 = 76066276

func (m *CoinType) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return ""
}

func (m *CoinType) GetCoinShortcut() string {
	if m != nil && m.CoinShortcut != nil {
		return *m.CoinShortcut
	}
	return ""
}

func (m *CoinType) GetAddressType() uint32 {
	if m != nil && m.AddressType != nil {
		return *m.AddressType
	}
	return Default_CoinType_AddressType
}

func (m *CoinType) GetMaxfeeKb() uint64 {
	if m != nil && m.MaxfeeKb != nil {
		return *m.MaxfeeKb
	}
	return 0
}

func (m *CoinType) GetAddressTypeP2Sh() uint32 {
	if m != nil && m.AddressTypeP2Sh != nil {
		return *m.AddressTypeP2Sh
	}
	return Default_CoinType_AddressTypeP2Sh
}

func (m *CoinType) GetSignedMessageHeader() string {
	if m != nil && m.SignedMessageHeader != nil {
		return *m.SignedMessageHeader
	}
	return ""
}

func (m *CoinType) GetXpubMagic() uint32 {
	if m != nil && m.XpubMagic != nil {
		return *m.XpubMagic
	}
	return Default_CoinType_XpubMagic
}

func (m *CoinType) GetXprvMagic() uint32 {
	if m != nil && m.XprvMagic != nil {
		return *m.XprvMagic
	}
	return Default_CoinType_XprvMagic
}

func (m *CoinType) GetSegwit() bool {
	if m != nil && m.Segwit != nil {
		return *m.Segwit
	}
	return false
}

func (m *CoinType) GetForkid() uint32 {
	if m != nil && m.Forkid != nil {
		return *m.Forkid
	}
	return 0
}

    
                                      
                       
type MultisigRedeemScriptType struct {
	Pubkeys          []*HDNodePathType `protobuf:"bytes,1,rep,name=pubkeys" json:"pubkeys,omitempty"`
	Signatures       [][]byte          `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
	M                *uint32           `protobuf:"varint,3,opt,name=m" json:"m,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *MultisigRedeemScriptType) Reset()                    { *m = MultisigRedeemScriptType{} }
func (m *MultisigRedeemScriptType) String() string            { return proto.CompactTextString(m) }
func (*MultisigRedeemScriptType) ProtoMessage()               {}
func (*MultisigRedeemScriptType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MultisigRedeemScriptType) GetPubkeys() []*HDNodePathType {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

func (m *MultisigRedeemScriptType) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *MultisigRedeemScriptType) GetM() uint32 {
	if m != nil && m.M != nil {
		return *m.M
	}
	return 0
}

    
                                           
                        
                           
type TxInputType struct {
	AddressN         []uint32                  `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	PrevHash         []byte                    `protobuf:"bytes,2,req,name=prev_hash,json=prevHash" json:"prev_hash,omitempty"`
	PrevIndex        *uint32                   `protobuf:"varint,3,req,name=prev_index,json=prevIndex" json:"prev_index,omitempty"`
	ScriptSig        []byte                    `protobuf:"bytes,4,opt,name=script_sig,json=scriptSig" json:"script_sig,omitempty"`
	Sequence         *uint32                   `protobuf:"varint,5,opt,name=sequence,def=4294967295" json:"sequence,omitempty"`
	ScriptType       *InputScriptType          `protobuf:"varint,6,opt,name=script_type,json=scriptType,enum=InputScriptType,def=0" json:"script_type,omitempty"`
	Multisig         *MultisigRedeemScriptType `protobuf:"bytes,7,opt,name=multisig" json:"multisig,omitempty"`
	Amount           *uint64                   `protobuf:"varint,8,opt,name=amount" json:"amount,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *TxInputType) Reset()                    { *m = TxInputType{} }
func (m *TxInputType) String() string            { return proto.CompactTextString(m) }
func (*TxInputType) ProtoMessage()               {}
func (*TxInputType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_TxInputType_Sequence uint32 = 4294967295
const Default_TxInputType_ScriptType InputScriptType = InputScriptType_SPENDADDRESS

func (m *TxInputType) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *TxInputType) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *TxInputType) GetPrevIndex() uint32 {
	if m != nil && m.PrevIndex != nil {
		return *m.PrevIndex
	}
	return 0
}

func (m *TxInputType) GetScriptSig() []byte {
	if m != nil {
		return m.ScriptSig
	}
	return nil
}

func (m *TxInputType) GetSequence() uint32 {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return Default_TxInputType_Sequence
}

func (m *TxInputType) GetScriptType() InputScriptType {
	if m != nil && m.ScriptType != nil {
		return *m.ScriptType
	}
	return Default_TxInputType_ScriptType
}

func (m *TxInputType) GetMultisig() *MultisigRedeemScriptType {
	if m != nil {
		return m.Multisig
	}
	return nil
}

func (m *TxInputType) GetAmount() uint64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

    
                                            
                        
                           
type TxOutputType struct {
	Address          *string                   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	AddressN         []uint32                  `protobuf:"varint,2,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Amount           *uint64                   `protobuf:"varint,3,req,name=amount" json:"amount,omitempty"`
	ScriptType       *OutputScriptType         `protobuf:"varint,4,req,name=script_type,json=scriptType,enum=OutputScriptType" json:"script_type,omitempty"`
	Multisig         *MultisigRedeemScriptType `protobuf:"bytes,5,opt,name=multisig" json:"multisig,omitempty"`
	OpReturnData     []byte                    `protobuf:"bytes,6,opt,name=op_return_data,json=opReturnData" json:"op_return_data,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *TxOutputType) Reset()                    { *m = TxOutputType{} }
func (m *TxOutputType) String() string            { return proto.CompactTextString(m) }
func (*TxOutputType) ProtoMessage()               {}
func (*TxOutputType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TxOutputType) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *TxOutputType) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *TxOutputType) GetAmount() uint64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *TxOutputType) GetScriptType() OutputScriptType {
	if m != nil && m.ScriptType != nil {
		return *m.ScriptType
	}
	return OutputScriptType_PAYTOADDRESS
}

func (m *TxOutputType) GetMultisig() *MultisigRedeemScriptType {
	if m != nil {
		return m.Multisig
	}
	return nil
}

func (m *TxOutputType) GetOpReturnData() []byte {
	if m != nil {
		return m.OpReturnData
	}
	return nil
}

    
                                                     
                           
type TxOutputBinType struct {
	Amount           *uint64 `protobuf:"varint,1,req,name=amount" json:"amount,omitempty"`
	ScriptPubkey     []byte  `protobuf:"bytes,2,req,name=script_pubkey,json=scriptPubkey" json:"script_pubkey,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TxOutputBinType) Reset()                    { *m = TxOutputBinType{} }
func (m *TxOutputBinType) String() string            { return proto.CompactTextString(m) }
func (*TxOutputBinType) ProtoMessage()               {}
func (*TxOutputBinType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TxOutputBinType) GetAmount() uint64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *TxOutputBinType) GetScriptPubkey() []byte {
	if m != nil {
		return m.ScriptPubkey
	}
	return nil
}

    
                                     
                        
type TransactionType struct {
	Version          *uint32            `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Inputs           []*TxInputType     `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty"`
	BinOutputs       []*TxOutputBinType `protobuf:"bytes,3,rep,name=bin_outputs,json=binOutputs" json:"bin_outputs,omitempty"`
	Outputs          []*TxOutputType    `protobuf:"bytes,5,rep,name=outputs" json:"outputs,omitempty"`
	LockTime         *uint32            `protobuf:"varint,4,opt,name=lock_time,json=lockTime" json:"lock_time,omitempty"`
	InputsCnt        *uint32            `protobuf:"varint,6,opt,name=inputs_cnt,json=inputsCnt" json:"inputs_cnt,omitempty"`
	OutputsCnt       *uint32            `protobuf:"varint,7,opt,name=outputs_cnt,json=outputsCnt" json:"outputs_cnt,omitempty"`
	ExtraData        []byte             `protobuf:"bytes,8,opt,name=extra_data,json=extraData" json:"extra_data,omitempty"`
	ExtraDataLen     *uint32            `protobuf:"varint,9,opt,name=extra_data_len,json=extraDataLen" json:"extra_data_len,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *TransactionType) Reset()                    { *m = TransactionType{} }
func (m *TransactionType) String() string            { return proto.CompactTextString(m) }
func (*TransactionType) ProtoMessage()               {}
func (*TransactionType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TransactionType) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *TransactionType) GetInputs() []*TxInputType {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TransactionType) GetBinOutputs() []*TxOutputBinType {
	if m != nil {
		return m.BinOutputs
	}
	return nil
}

func (m *TransactionType) GetOutputs() []*TxOutputType {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TransactionType) GetLockTime() uint32 {
	if m != nil && m.LockTime != nil {
		return *m.LockTime
	}
	return 0
}

func (m *TransactionType) GetInputsCnt() uint32 {
	if m != nil && m.InputsCnt != nil {
		return *m.InputsCnt
	}
	return 0
}

func (m *TransactionType) GetOutputsCnt() uint32 {
	if m != nil && m.OutputsCnt != nil {
		return *m.OutputsCnt
	}
	return 0
}

func (m *TransactionType) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func (m *TransactionType) GetExtraDataLen() uint32 {
	if m != nil && m.ExtraDataLen != nil {
		return *m.ExtraDataLen
	}
	return 0
}

    
                                         
                     
type TxRequestDetailsType struct {
	RequestIndex     *uint32 `protobuf:"varint,1,opt,name=request_index,json=requestIndex" json:"request_index,omitempty"`
	TxHash           []byte  `protobuf:"bytes,2,opt,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
	ExtraDataLen     *uint32 `protobuf:"varint,3,opt,name=extra_data_len,json=extraDataLen" json:"extra_data_len,omitempty"`
	ExtraDataOffset  *uint32 `protobuf:"varint,4,opt,name=extra_data_offset,json=extraDataOffset" json:"extra_data_offset,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TxRequestDetailsType) Reset()                    { *m = TxRequestDetailsType{} }
func (m *TxRequestDetailsType) String() string            { return proto.CompactTextString(m) }
func (*TxRequestDetailsType) ProtoMessage()               {}
func (*TxRequestDetailsType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TxRequestDetailsType) GetRequestIndex() uint32 {
	if m != nil && m.RequestIndex != nil {
		return *m.RequestIndex
	}
	return 0
}

func (m *TxRequestDetailsType) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *TxRequestDetailsType) GetExtraDataLen() uint32 {
	if m != nil && m.ExtraDataLen != nil {
		return *m.ExtraDataLen
	}
	return 0
}

func (m *TxRequestDetailsType) GetExtraDataOffset() uint32 {
	if m != nil && m.ExtraDataOffset != nil {
		return *m.ExtraDataOffset
	}
	return 0
}

    
                                         
                     
type TxRequestSerializedType struct {
	SignatureIndex   *uint32 `protobuf:"varint,1,opt,name=signature_index,json=signatureIndex" json:"signature_index,omitempty"`
	Signature        []byte  `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	SerializedTx     []byte  `protobuf:"bytes,3,opt,name=serialized_tx,json=serializedTx" json:"serialized_tx,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TxRequestSerializedType) Reset()                    { *m = TxRequestSerializedType{} }
func (m *TxRequestSerializedType) String() string            { return proto.CompactTextString(m) }
func (*TxRequestSerializedType) ProtoMessage()               {}
func (*TxRequestSerializedType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TxRequestSerializedType) GetSignatureIndex() uint32 {
	if m != nil && m.SignatureIndex != nil {
		return *m.SignatureIndex
	}
	return 0
}

func (m *TxRequestSerializedType) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TxRequestSerializedType) GetSerializedTx() []byte {
	if m != nil {
		return m.SerializedTx
	}
	return nil
}

    
                                       
                        
type IdentityType struct {
	Proto            *string `protobuf:"bytes,1,opt,name=proto" json:"proto,omitempty"`
	User             *string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Host             *string `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	Port             *string `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
	Path             *string `protobuf:"bytes,5,opt,name=path" json:"path,omitempty"`
	Index            *uint32 `protobuf:"varint,6,opt,name=index,def=0" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IdentityType) Reset()                    { *m = IdentityType{} }
func (m *IdentityType) String() string            { return proto.CompactTextString(m) }
func (*IdentityType) ProtoMessage()               {}
func (*IdentityType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

const Default_IdentityType_Index uint32 = 0

func (m *IdentityType) GetProto() string {
	if m != nil && m.Proto != nil {
		return *m.Proto
	}
	return ""
}

func (m *IdentityType) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *IdentityType) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *IdentityType) GetPort() string {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return ""
}

func (m *IdentityType) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *IdentityType) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return Default_IdentityType_Index
}

var E_WireIn = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50002,
	Name:          "wire_in",
	Tag:           "varint,50002,opt,name=wire_in,json=wireIn",
	Filename:      "types.proto",
}

var E_WireOut = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50003,
	Name:          "wire_out",
	Tag:           "varint,50003,opt,name=wire_out,json=wireOut",
	Filename:      "types.proto",
}

var E_WireDebugIn = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50004,
	Name:          "wire_debug_in",
	Tag:           "varint,50004,opt,name=wire_debug_in,json=wireDebugIn",
	Filename:      "types.proto",
}

var E_WireDebugOut = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50005,
	Name:          "wire_debug_out",
	Tag:           "varint,50005,opt,name=wire_debug_out,json=wireDebugOut",
	Filename:      "types.proto",
}

var E_WireTiny = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50006,
	Name:          "wire_tiny",
	Tag:           "varint,50006,opt,name=wire_tiny,json=wireTiny",
	Filename:      "types.proto",
}

var E_WireBootloader = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50007,
	Name:          "wire_bootloader",
	Tag:           "varint,50007,opt,name=wire_bootloader,json=wireBootloader",
	Filename:      "types.proto",
}

func init() {
	proto.RegisterType((*HDNodeType)(nil), "HDNodeType")
	proto.RegisterType((*HDNodePathType)(nil), "HDNodePathType")
	proto.RegisterType((*CoinType)(nil), "CoinType")
	proto.RegisterType((*MultisigRedeemScriptType)(nil), "MultisigRedeemScriptType")
	proto.RegisterType((*TxInputType)(nil), "TxInputType")
	proto.RegisterType((*TxOutputType)(nil), "TxOutputType")
	proto.RegisterType((*TxOutputBinType)(nil), "TxOutputBinType")
	proto.RegisterType((*TransactionType)(nil), "TransactionType")
	proto.RegisterType((*TxRequestDetailsType)(nil), "TxRequestDetailsType")
	proto.RegisterType((*TxRequestSerializedType)(nil), "TxRequestSerializedType")
	proto.RegisterType((*IdentityType)(nil), "IdentityType")
	proto.RegisterEnum("FailureType", FailureType_name, FailureType_value)
	proto.RegisterEnum("OutputScriptType", OutputScriptType_name, OutputScriptType_value)
	proto.RegisterEnum("InputScriptType", InputScriptType_name, InputScriptType_value)
	proto.RegisterEnum("RequestType", RequestType_name, RequestType_value)
	proto.RegisterEnum("ButtonRequestType", ButtonRequestType_name, ButtonRequestType_value)
	proto.RegisterEnum("PinMatrixRequestType", PinMatrixRequestType_name, PinMatrixRequestType_value)
	proto.RegisterEnum("RecoveryDeviceType", RecoveryDeviceType_name, RecoveryDeviceType_value)
	proto.RegisterEnum("WordRequestType", WordRequestType_name, WordRequestType_value)
	proto.RegisterExtension(E_WireIn)
	proto.RegisterExtension(E_WireOut)
	proto.RegisterExtension(E_WireDebugIn)
	proto.RegisterExtension(E_WireDebugOut)
	proto.RegisterExtension(E_WireTiny)
	proto.RegisterExtension(E_WireBootloader)
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	                                              
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x57, 0xdb, 0x72, 0x1a, 0xc9,
	0x19, 0xf6, 0x00, 0x92, 0xe0, 0x07, 0xc4, 0xa8, 0x7d, 0xd0, 0x78, 0x6d, 0xaf, 0x31, 0x76, 0x62,
	0x45, 0x55, 0x61, 0x77, 0xc9, 0x5a, 0x8e, 0x55, 0xa9, 0x24, 0x3a, 0xa0, 0x15, 0x65, 0x0b, 0x51,
	0xc3, 0x28, 0x56, 0x72, 0x33, 0x35, 0xcc, 0xb4, 0xa0, 0x4b, 0x43, 0x37, 0xe9, 0xe9, 0x91, 0xd1,
	0xde, 0xe4, 0x2a, 0xc9, 0x55, 0x5e, 0x23, 0x6f, 0x91, 0xaa, 0xbc, 0x41, 0xaa, 0x36, 0xa7, 0xcb,
	0xbc, 0x41, 0xae, 0xf2, 0x00, 0xa9, 0x3e, 0x0c, 0x02, 0xc9, 0xde, 0xd2, 0x1d, 0xfd, 0x7d, 0xff,
	0xf9, 0xd0, 0x3d, 0x40, 0x59, 0x5c, 0x4e, 0x70, 0xd2, 0x9c, 0x70, 0x26, 0xd8, 0x67, 0xf5, 0x21,
	0x63, 0xc3, 0x18, 0x7f, 0xa1, 0x4e, 0x83, 0xf4, 0xec, 0x8b, 0x08, 0x27, 0x21, 0x27, 0x13, 0xc1,
	0xb8, 0x96, 0x68, 0xfc, 0xd5, 0x02, 0x38, 0xdc, 0xef, 0xb2, 0x08, 0x7b, 0x97, 0x13, 0x8c, 0xee,
	0xc1, 0x52, 0x84, 0x27, 0x62, 0xe4, 0x58, 0xf5, 0xdc, 0x46, 0xd5, 0xd5, 0x07, 0x54, 0x87, 0xf2,
	0x19, 0xa1, 0x43, 0xcc, 0x27, 0x9c, 0x50, 0xe1, 0xe4, 0x14, 0x37, 0x0f, 0xa1, 0x47, 0x50, 0x0a,
	0x47, 0x24, 0x8e, 0x7c, 0x9a, 0x8e, 0x9d, 0xbc, 0xe2, 0x8b, 0x0a, 0xe8, 0xa6, 0x63, 0xf4, 0x04,
	0x20, 0x1c, 0x05, 0x84, 0xfa, 0x21, 0x8b, 0xb0, 0x53, 0xa8, 0xe7, 0x36, 0x2a, 0x6e, 0x49, 0x21,
	0x7b, 0x2c, 0xc2, 0xe8, 0x29, 0x94, 0x27, 0x9c, 0x5c, 0x04, 0x02, 0xfb, 0xe7, 0xf8, 0xd2, 0x59,
	0xaa, 0x5b, 0x1b, 0x15, 0x17, 0x0c, 0xf4, 0x16, 0x5f, 0x4a, 0xfd, 0x49, 0x3a, 0x88, 0x49, 0xa8,
	0xf8, 0x65, 0xc5, 0x97, 0x34, 0xf2, 0x16, 0x5f, 0x36, 0xba, 0xb0, 0xaa, 0x33, 0xe8, 0x05, 0x62,
	0xa4, 0xb2, 0x78, 0x0a, 0x05, 0x2a, 0x5d, 0xc9, 0x24, 0xca, 0xad, 0x72, 0xf3, 0x2a, 0x41, 0x57,
	0x11, 0x32, 0xdc, 0x20, 0x8a, 0x38, 0x4e, 0x12, 0x9f, 0x3a, 0xb9, 0x7a, 0x5e, 0x86, 0x6b, 0x80,
	0x6e, 0xe3, 0x7f, 0x39, 0x28, 0xee, 0x31, 0x42, 0x95, 0x29, 0x99, 0x18, 0x23, 0xd4, 0xa7, 0xc1,
	0x58, 0xda, 0xb3, 0x36, 0x4a, 0x6e, 0x51, 0x02, 0xdd, 0x60, 0x8c, 0xd1, 0x73, 0xa8, 0x2a, 0x32,
	0x19, 0x31, 0x2e, 0xc2, 0x54, 0x56, 0x46, 0x0a, 0x54, 0x24, 0xd8, 0x37, 0x18, 0x7a, 0x01, 0x95,
	0xcc, 0x97, 0x6c, 0x8d, 0x93, 0xaf, 0x5b, 0x1b, 0xd5, 0x6d, 0xeb, 0x4b, 0xb7, 0x6c, 0xe0, 0xcc,
	0xcf, 0x38, 0x98, 0x9e, 0x61, 0xec, 0x9f, 0x0f, 0x9c, 0x42, 0xdd, 0xda, 0x28, 0xb8, 0x45, 0x0d,
	0xbc, 0x1d, 0xa0, 0x1f, 0xc3, 0xda, 0xbc, 0x09, 0x7f, 0xd2, 0x4a, 0x46, 0xaa, 0x4e, 0xd5, 0x6d,
	0xeb, 0x95, 0x5b, 0x9b, 0xb3, 0xd3, 0x6b, 0x25, 0x23, 0xd4, 0x82, 0xfb, 0x09, 0x19, 0x52, 0x1c,
	0xf9, 0x63, 0x9c, 0x24, 0xc1, 0x10, 0xfb, 0x23, 0x1c, 0x44, 0x98, 0x3b, 0x45, 0x15, 0xde, 0x5d,
	0x4d, 0x1e, 0x69, 0xee, 0x50, 0x51, 0xe8, 0x25, 0xc0, 0x74, 0x92, 0x0e, 0xfc, 0x71, 0x30, 0x24,
	0xa1, 0x53, 0x52, 0xb6, 0x8b, 0xaf, 0xb7, 0xbe, 0xdc, 0x7a, 0xfd, 0x93, 0x57, 0x3f, 0x75, 0x4b,
	0x92, 0x3b, 0x92, 0x94, 0x16, 0xe4, 0x17, 0x46, 0x10, 0xae, 0x04, 0xb7, 0x5a, 0xaf, 0xb7, 0xa4,
	0x20, 0xbf, 0xd0, 0x82, 0x0f, 0x60, 0x39, 0xc1, 0xc3, 0x0f, 0x44, 0x38, 0xe5, 0xba, 0xb5, 0x51,
	0x74, 0xcd, 0x49, 0xe2, 0x67, 0x8c, 0x9f, 0x93, 0xc8, 0xa9, 0x48, 0x65, 0xd7, 0x9c, 0x1a, 0x09,
	0x38, 0x47, 0x69, 0x2c, 0x48, 0x42, 0x86, 0x2e, 0x8e, 0x30, 0x1e, 0xf7, 0xd5, 0xa4, 0xaa, 0xea,
	0xfc, 0x08, 0x56, 0x26, 0xe9, 0xe0, 0x1c, 0x5f, 0x26, 0x8e, 0x55, 0xcf, 0x6f, 0x94, 0x5b, 0xb5,
	0xe6, 0x62, 0xcb, 0xdd, 0x8c, 0x47, 0x9f, 0x03, 0xc8, 0xfc, 0x02, 0x91, 0x72, 0x9c, 0xa8, 0xde,
	0x56, 0xdc, 0x39, 0x04, 0x55, 0xc0, 0x1a, 0xeb, 0x1e, 0xb8, 0xd6, 0xb8, 0xf1, 0x97, 0x1c, 0x94,
	0xbd, 0x69, 0x87, 0x4e, 0x52, 0x91, 0xb5, 0xe1, 0x6a, 0x30, 0xac, 0xc5, 0xc1, 0x90, 0xe4, 0x84,
	0xe3, 0x0b, 0x7f, 0x14, 0x24, 0x23, 0xb5, 0x04, 0x15, 0xb7, 0x28, 0x81, 0xc3, 0x20, 0x19, 0xa9,
	0x21, 0x95, 0x24, 0xa1, 0x11, 0x9e, 0x9a, 0x15, 0x50, 0xe2, 0x1d, 0x09, 0x48, 0x5a, 0x6f, 0x9e,
	0x9f, 0x90, 0xa1, 0x6a, 0x70, 0xc5, 0x2d, 0x69, 0xa4, 0x4f, 0x86, 0xe8, 0x87, 0x50, 0x4c, 0xf0,
	0x6f, 0x53, 0x4c, 0x43, 0x6c, 0x1a, 0x0b, 0x5f, 0xb7, 0xde, 0x7c, 0xfd, 0x66, 0xeb, 0x75, 0xeb,
	0xcd, 0x2b, 0x77, 0xc6, 0xa1, 0x5f, 0x40, 0xd9, 0x98, 0x51, 0xb3, 0x24, 0x77, 0x61, 0xb5, 0x65,
	0x37, 0x55, 0x02, 0x57, 0xf5, 0xda, 0xae, 0xf4, 0x7b, 0xed, 0xee, 0xfe, 0xce, 0xfe, 0xbe, 0xdb,
	0xee, 0xf7, 0x5d, 0xe3, 0x59, 0x25, 0xf8, 0x0a, 0x8a, 0x63, 0x53, 0x65, 0x67, 0xa5, 0x6e, 0x6d,
	0x94, 0x5b, 0x0f, 0x9b, 0x9f, 0x2a, 0xbb, 0x3b, 0x13, 0x95, 0x4d, 0x0b, 0xc6, 0x2c, 0xa5, 0x42,
	0xcd, 0x50, 0xc1, 0x35, 0xa7, 0xc6, 0x7f, 0x2d, 0xa8, 0x78, 0xd3, 0xe3, 0x54, 0x64, 0x05, 0x74,
	0x60, 0xc5, 0xd4, 0xcb, 0x6c, 0x4b, 0x76, 0xfc, 0xde, 0x9d, 0x9b, 0xb3, 0x2f, 0x2b, 0x37, 0xb3,
	0x8f, 0x5a, 0x8b, 0xf9, 0xca, 0xbb, 0x63, 0xb5, 0xb5, 0xd6, 0xd4, 0x0e, 0xe7, 0x22, 0xfd, 0x54,
	0x8a, 0x4b, 0xb7, 0x4f, 0xf1, 0x05, 0xac, 0xb2, 0x89, 0xcf, 0xb1, 0x48, 0x39, 0xf5, 0xa3, 0x40,
	0x04, 0xe6, 0xa6, 0xa9, 0xb0, 0x89, 0xab, 0xc0, 0xfd, 0x40, 0x04, 0x8d, 0x2e, 0xd4, 0xb2, 0x7c,
	0x77, 0xcd, 0x15, 0x71, 0x15, 0xbb, 0xb5, 0x10, 0xfb, 0x73, 0xa8, 0x9a, 0xd8, 0xf5, 0x6c, 0x9a,
	0x91, 0xa9, 0x68, 0xb0, 0xa7, 0xb0, 0xc6, 0xdf, 0x72, 0x50, 0xf3, 0x78, 0x40, 0x93, 0x20, 0x14,
	0x84, 0xd1, 0xac, 0x86, 0x17, 0x98, 0x27, 0x84, 0x51, 0x55, 0xc3, 0xaa, 0x9b, 0x1d, 0xd1, 0x0b,
	0x58, 0x26, 0xb2, 0xd5, 0x7a, 0xb0, 0xcb, 0xad, 0x4a, 0x73, 0x6e, 0x78, 0x5d, 0xc3, 0xa1, 0xaf,
	0xa0, 0x3c, 0x20, 0xd4, 0x67, 0x2a, 0xca, 0xc4, 0xc9, 0x2b, 0x51, 0xbb, 0x79, 0x2d, 0x6e, 0x17,
	0x06, 0x84, 0x6a, 0x24, 0x41, 0x2f, 0x61, 0x25, 0x13, 0x5f, 0x52, 0xe2, 0xd5, 0xe6, 0x7c, 0x5b,
	0xdd, 0x8c, 0x95, 0x5d, 0x8c, 0x59, 0x78, 0xee, 0x0b, 0x32, 0xc6, 0x6a, 0x8c, 0xab, 0x6e, 0x51,
	0x02, 0x1e, 0x19, 0x63, 0x39, 0xe4, 0x3a, 0x04, 0x3f, 0xa4, 0x42, 0x95, 0xaf, 0xea, 0x96, 0x34,
	0xb2, 0x47, 0x85, 0xbc, 0xe8, 0x8d, 0x19, 0xc5, 0xaf, 0x28, 0x1e, 0x0c, 0x24, 0x05, 0x9e, 0x00,
	0xe0, 0xa9, 0xe0, 0x81, 0x2e, 0x7f, 0x51, 0x2f, 0x89, 0x42, 0x64, 0xed, 0x65, 0x87, 0xae, 0x68,
	0x3f, 0xc6, 0x54, 0xdf, 0x53, 0x6e, 0x65, 0x26, 0xf2, 0x0e, 0xd3, 0xc6, 0x9f, 0x2d, 0xb8, 0xe7,
	0x4d, 0x5d, 0xb9, 0x31, 0x89, 0xd8, 0xc7, 0x22, 0x20, 0xb1, 0xbe, 0x62, 0x9f, 0x43, 0x95, 0x6b,
	0xd4, 0x2c, 0xa9, 0x2e, 0x6e, 0xc5, 0x80, 0x7a, 0x4f, 0xd7, 0x61, 0x45, 0x4c, 0xb3, 0x0d, 0x97,
	0xfe, 0x97, 0xc5, 0x54, 0xed, 0xf7, 0x4d, 0xe7, 0xf9, 0x9b, 0xce, 0xd1, 0x26, 0xac, 0xcd, 0x49,
	0xb1, 0xb3, 0xb3, 0x04, 0x0b, 0x53, 0xa6, 0xda, 0x4c, 0xf0, 0x58, 0xc1, 0x8d, 0xdf, 0x5b, 0xb0,
	0x3e, 0x0b, 0xb4, 0x8f, 0x39, 0x09, 0x62, 0xf2, 0x2d, 0x8e, 0x54, 0xac, 0x2f, 0xa1, 0x36, 0xbb,
	0xb3, 0x16, 0xa2, 0x5d, 0x9d, 0xc1, 0x3a, 0xde, 0xc7, 0x50, 0x9a, 0x21, 0x26, 0xe2, 0x2b, 0x40,
	0x8d, 0xe0, 0xcc, 0xb0, 0x2f, 0xa6, 0x2a, 0x66, 0x39, 0x82, 0x57, 0xde, 0xa6, 0x8d, 0x3f, 0x59,
	0x50, 0xe9, 0x44, 0x98, 0x0a, 0x22, 0x2e, 0xb3, 0x8f, 0x00, 0xf5, 0x71, 0x60, 0x36, 0x58, 0x1f,
	0x10, 0x82, 0x42, 0x9a, 0x60, 0x6e, 0xde, 0x38, 0xf5, 0x5b, 0x62, 0x23, 0x96, 0x08, 0x65, 0xb6,
	0xe4, 0xaa, 0xdf, 0x12, 0x9b, 0x30, 0xae, 0xb3, 0x2e, 0xb9, 0xea, 0xb7, 0xc2, 0x02, 0xa1, 0xdf,
	0x2c, 0x89, 0x05, 0x62, 0x84, 0xd6, 0x61, 0x49, 0x27, 0xb6, 0x9c, 0x3d, 0x88, 0xfa, 0xbc, 0xf9,
	0x5d, 0x0e, 0xca, 0x07, 0x01, 0x89, 0x53, 0xae, 0xbf, 0x49, 0x9e, 0xc0, 0x43, 0x73, 0xf4, 0x4f,
	0x28, 0x9e, 0x4e, 0x70, 0x28, 0x66, 0xaf, 0x97, 0x6d, 0xa1, 0xcf, 0xe0, 0x41, 0x46, 0xef, 0xa6,
	0x42, 0x30, 0xda, 0x36, 0x22, 0x76, 0x0e, 0xdd, 0x87, 0xb5, 0x8c, 0x93, 0x85, 0x6f, 0x73, 0xce,
	0xb8, 0x9d, 0x47, 0x8f, 0x60, 0x3d, 0x83, 0x77, 0xd4, 0xda, 0xed, 0x05, 0x34, 0xc4, 0x71, 0x8c,
	0x23, 0xbb, 0x80, 0xd6, 0xe1, 0x6e, 0x46, 0xf6, 0xc8, 0x95, 0xb1, 0x25, 0xe4, 0xc0, 0xbd, 0x39,
	0xe2, 0x4a, 0x65, 0x19, 0x3d, 0x00, 0x34, 0xc7, 0x74, 0xe8, 0x45, 0x10, 0x93, 0xc8, 0x5e, 0x41,
	0x8f, 0xc1, 0xc9, 0x70, 0x03, 0xf6, 0xb3, 0xd6, 0xd8, 0xc5, 0x05, 0x7b, 0x9c, 0x85, 0x38, 0x49,
	0x74, 0x7c, 0xa5, 0xf9, 0x94, 0xba, 0x4c, 0xb4, 0x29, 0x4b, 0x87, 0xa3, 0x83, 0x94, 0x46, 0x89,
	0x0d, 0xd7, 0xb8, 0x0e, 0x25, 0xc2, 0x74, 0xd2, 0x2e, 0xa3, 0x87, 0x70, 0x3f, 0xe3, 0x0e, 0x08,
	0x1f, 0x7f, 0x08, 0x38, 0xd6, 0x26, 0xc3, 0xcd, 0x3f, 0x5a, 0x60, 0x5f, 0xbf, 0x35, 0x91, 0x0d,
	0x95, 0xde, 0xce, 0xaf, 0xbd, 0x63, 0xf3, 0x50, 0xd8, 0x77, 0xd0, 0x5d, 0xa8, 0x29, 0xa4, 0xbf,
	0xe7, 0x76, 0x7a, 0xde, 0xe1, 0x4e, 0xff, 0xd0, 0xb6, 0xd0, 0x1a, 0x54, 0x15, 0x78, 0x74, 0xf2,
	0xce, 0xeb, 0xf4, 0x3b, 0xdf, 0xd8, 0xb9, 0x19, 0x74, 0xdc, 0x73, 0xdb, 0xde, 0x89, 0xdb, 0xb5,
	0xf3, 0x33, 0x63, 0xef, 0x3b, 0x5e, 0x57, 0x1a, 0x2b, 0xa0, 0x7b, 0x60, 0x2b, 0xa4, 0xd7, 0xea,
	0x1f, 0x66, 0xe8, 0xd2, 0x66, 0x0c, 0xb5, 0x6b, 0xcf, 0x95, 0x54, 0x9d, 0x7f, 0xb0, 0xec, 0x3b,
	0xd2, 0xbe, 0x42, 0x66, 0x2e, 0x2d, 0x54, 0x81, 0x62, 0xfb, 0xd4, 0x6b, 0xbb, 0xdd, 0x9d, 0x77,
	0x76, 0x6e, 0xa6, 0x92, 0xd9, 0xcd, 0x4b, 0x6f, 0x0a, 0x99, 0xf7, 0x56, 0xd8, 0x3c, 0x81, 0xb2,
	0xd9, 0x30, 0xe5, 0xa9, 0x0c, 0x2b, 0xde, 0x69, 0xa7, 0xdb, 0x3b, 0xf1, 0xec, 0x3b, 0xd2, 0xa2,
	0x77, 0x7a, 0x7c, 0xe2, 0xc9, 0x93, 0x85, 0x00, 0x96, 0xbd, 0xd3, 0xa3, 0xb6, 0xb7, 0x63, 0xe7,
	0xd0, 0x2a, 0x80, 0x77, 0x7a, 0xd0, 0xe9, 0x76, 0xfa, 0x87, 0xed, 0x7d, 0x3b, 0x8f, 0x6a, 0x50,
	0xf6, 0x4e, 0xdb, 0xa7, 0x9e, 0xbb, 0xb3, 0xbf, 0xe3, 0xed, 0xd8, 0x85, 0xcd, 0xff, 0xe4, 0x60,
	0x4d, 0x4f, 0xdb, 0xbc, 0xf5, 0x75, 0xb8, 0xbb, 0x00, 0xfa, 0xc7, 0x62, 0x84, 0xb9, 0x6d, 0xa1,
	0x06, 0x7c, 0xbe, 0x48, 0x1c, 0x60, 0x7c, 0x7c, 0x81, 0xb9, 0x37, 0xe2, 0x38, 0x19, 0xb1, 0x58,
	0xce, 0xea, 0x53, 0x78, 0xb4, 0x28, 0xb3, 0xc7, 0xe8, 0x19, 0xe1, 0x63, 0xdd, 0x35, 0x3b, 0x2f,
	0xf7, 0x60, 0x51, 0xc0, 0xc5, 0x09, 0x16, 0xfb, 0xf8, 0x82, 0x84, 0xd8, 0x2e, 0xdc, 0xa4, 0x8d,
	0xfe, 0x7b, 0xc6, 0xe5, 0xf4, 0x3e, 0x06, 0x67, 0x91, 0x7e, 0x4f, 0x26, 0xd8, 0x28, 0x2f, 0xdf,
	0x54, 0xee, 0x71, 0x26, 0x70, 0x28, 0xf6, 0x82, 0x38, 0xb6, 0x57, 0xe4, 0xa8, 0x2e, 0xd2, 0x72,
	0x8e, 0xbd, 0xa9, 0x5d, 0xbc, 0x19, 0x75, 0x36, 0x78, 0x7b, 0x23, 0x1c, 0x9e, 0xdb, 0x25, 0x39,
	0x93, 0x8b, 0x02, 0x3b, 0xfa, 0xcd, 0xb7, 0x41, 0xae, 0xe1, 0x35, 0xa7, 0xd9, 0x37, 0xbd, 0x5d,
	0xde, 0xfc, 0x1d, 0xdc, 0xeb, 0x11, 0x7a, 0x14, 0x08, 0x4e, 0xa6, 0xf3, 0x35, 0xae, 0xc3, 0xe3,
	0x8f, 0xe1, 0xfe, 0x5e, 0xca, 0x39, 0xa6, 0xc2, 0xb6, 0xd0, 0x33, 0x78, 0xf2, 0x51, 0x89, 0x2e,
	0xfe, 0x70, 0x40, 0x78, 0x22, 0xec, 0x9c, 0xec, 0xc7, 0xa7, 0x44, 0xfa, 0x38, 0x64, 0x34, 0xb2,
	0xf3, 0x9b, 0xbf, 0x01, 0xe4, 0xe2, 0x90, 0x5d, 0x60, 0x7e, 0xa9, 0xcb, 0xa4, 0xdc, 0xff, 0x00,
	0x9e, 0xdd, 0x44, 0xfd, 0x7e, 0xc8, 0x83, 0xf1, 0x20, 0xc6, 0x91, 0x2c, 0x76, 0x62, 0xdf, 0x91,
	0xf5, 0xfc, 0x88, 0x98, 0x76, 0x68, 0x5b, 0x9b, 0x67, 0x50, 0x93, 0x92, 0xf3, 0x79, 0x3d, 0x84,
	0xfb, 0xd7, 0x20, 0xbf, 0x17, 0x07, 0x84, 0xda, 0x77, 0x64, 0x9d, 0xae, 0x53, 0xda, 0xd2, 0x1b,
	0xdb, 0xfa, 0x34, 0xb9, 0x65, 0xe7, 0xb6, 0x7f, 0x06, 0x2b, 0x1f, 0x88, 0x7a, 0x41, 0xd0, 0xb3,
	0xa6, 0xfe, 0x2f, 0xd8, 0xcc, 0xfe, 0x0b, 0x36, 0xdb, 0x34, 0x1d, 0xff, 0x2a, 0x88, 0x53, 0x7c,
	0x3c, 0x91, 0x77, 0x60, 0xe2, 0x7c, 0xf7, 0x87, 0xbc, 0xfe, 0x52, 0x97, 0x3a, 0x1d, 0xba, 0xfd,
	0x73, 0x28, 0x2a, 0x6d, 0x96, 0x8a, 0xdb, 0xa8, 0xff, 0xdd, 0xa8, 0x2b, 0x97, 0xc7, 0xa9, 0xd8,
	0xfe, 0x06, 0xaa, 0x4a, 0x3f, 0xc2, 0x83, 0x74, 0x78, 0xcb, 0x18, 0xfe, 0x61, 0x8c, 0x94, 0xa5,
	0xe6, 0xbe, 0x54, 0xec, 0xd0, 0xed, 0x0e, 0xac, 0xce, 0x19, 0xba, 0x65, 0x38, 0xff, 0x34, 0x96,
	0x2a, 0x33, 0x4b, 0x32, 0xa6, 0x5f, 0x42, 0x49, 0x99, 0x12, 0x84, 0x5e, 0xde, 0xc6, 0xca, 0xbf,
	0x8c, 0x15, 0x55, 0x09, 0x8f, 0xd0, 0xcb, 0xed, 0x77, 0x50, 0x53, 0x16, 0x06, 0x8c, 0x89, 0x98,
	0xa9, 0x3f, 0x4f, 0xb7, 0xb0, 0xf3, 0x6f, 0x63, 0x47, 0x25, 0xb2, 0x3b, 0x53, 0xdd, 0xfd, 0x0a,
	0x9e, 0x87, 0x6c, 0xdc, 0x4c, 0x02, 0xc1, 0x92, 0x11, 0x89, 0x83, 0x41, 0xd2, 0x14, 0x1c, 0x7f,
	0xcb, 0x78, 0x33, 0x26, 0x83, 0x99, 0xbd, 0x5d, 0xf0, 0x14, 0x28, 0xdb, 0xfb, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x70, 0x88, 0xcd, 0x71, 0xe2, 0x0f, 0x00, 0x00,
}
