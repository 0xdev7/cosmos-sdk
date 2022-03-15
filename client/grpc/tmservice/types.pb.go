// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/tendermint/v1beta1/types.proto

package tmservice

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/tendermint/tendermint/proto/tendermint/p2p"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
	version "github.com/tendermint/tendermint/proto/tendermint/version"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BlockMeta struct {
	BlockID   types.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id"`
	BlockSize int64         `protobuf:"varint,2,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	Header    Header        `protobuf:"bytes,3,opt,name=header,proto3" json:"header"`
	NumTxs    int64         `protobuf:"varint,4,opt,name=num_txs,json=numTxs,proto3" json:"num_txs,omitempty"`
}

func (m *BlockMeta) Reset()         { *m = BlockMeta{} }
func (m *BlockMeta) String() string { return proto.CompactTextString(m) }
func (*BlockMeta) ProtoMessage()    {}
func (*BlockMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9931519c08e0d6, []int{0}
}
func (m *BlockMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMeta.Merge(m, src)
}
func (m *BlockMeta) XXX_Size() int {
	return m.Size()
}
func (m *BlockMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMeta.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMeta proto.InternalMessageInfo

func (m *BlockMeta) GetBlockID() types.BlockID {
	if m != nil {
		return m.BlockID
	}
	return types.BlockID{}
}

func (m *BlockMeta) GetBlockSize() int64 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *BlockMeta) GetHeader() Header {
	if m != nil {
		return m.Header
	}
	return Header{}
}

func (m *BlockMeta) GetNumTxs() int64 {
	if m != nil {
		return m.NumTxs
	}
	return 0
}

// Header defines the structure of a Tendermint block header.
type Header struct {
	// basic block info
	Version version.Consensus `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	ChainID string            `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Height  int64             `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time    time.Time         `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
	// prev block info
	LastBlockId types.BlockID `protobuf:"bytes,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	// hashes of block data
	LastCommitHash []byte `protobuf:"bytes,6,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	DataHash       []byte `protobuf:"bytes,7,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	// hashes from the app output from the prev block
	ValidatorsHash     []byte `protobuf:"bytes,8,opt,name=validators_hash,json=validatorsHash,proto3" json:"validators_hash,omitempty"`
	NextValidatorsHash []byte `protobuf:"bytes,9,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"next_validators_hash,omitempty"`
	ConsensusHash      []byte `protobuf:"bytes,10,opt,name=consensus_hash,json=consensusHash,proto3" json:"consensus_hash,omitempty"`
	AppHash            []byte `protobuf:"bytes,11,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	LastResultsHash    []byte `protobuf:"bytes,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	// consensus info
	EvidenceHash    []byte `protobuf:"bytes,13,opt,name=evidence_hash,json=evidenceHash,proto3" json:"evidence_hash,omitempty"`
	ProposerAddress string `protobuf:"bytes,14,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9931519c08e0d6, []int{1}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() version.Consensus {
	if m != nil {
		return m.Version
	}
	return version.Consensus{}
}

func (m *Header) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Header) GetLastBlockId() types.BlockID {
	if m != nil {
		return m.LastBlockId
	}
	return types.BlockID{}
}

func (m *Header) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *Header) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *Header) GetValidatorsHash() []byte {
	if m != nil {
		return m.ValidatorsHash
	}
	return nil
}

func (m *Header) GetNextValidatorsHash() []byte {
	if m != nil {
		return m.NextValidatorsHash
	}
	return nil
}

func (m *Header) GetConsensusHash() []byte {
	if m != nil {
		return m.ConsensusHash
	}
	return nil
}

func (m *Header) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *Header) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *Header) GetEvidenceHash() []byte {
	if m != nil {
		return m.EvidenceHash
	}
	return nil
}

func (m *Header) GetProposerAddress() string {
	if m != nil {
		return m.ProposerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*BlockMeta)(nil), "cosmos.base.tendermint.v1beta1.BlockMeta")
	proto.RegisterType((*Header)(nil), "cosmos.base.tendermint.v1beta1.Header")
}

func init() {
	proto.RegisterFile("cosmos/base/tendermint/v1beta1/types.proto", fileDescriptor_bb9931519c08e0d6)
}

var fileDescriptor_bb9931519c08e0d6 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0xe3, 0x4b, 0x6e, 0xfe, 0x4c, 0x08, 0xe1, 0x5a, 0xe8, 0x36, 0xa4, 0xe0, 0x20, 0xaa,
	0xd2, 0x14, 0xa9, 0x76, 0x49, 0x37, 0xdd, 0x74, 0xd1, 0x84, 0x4a, 0x20, 0xb5, 0x1b, 0x17, 0x75,
	0xd1, 0x4d, 0x34, 0xb1, 0xa7, 0xce, 0x88, 0x78, 0x66, 0xea, 0x19, 0x47, 0xc0, 0x53, 0xf0, 0x58,
	0x2c, 0x59, 0x76, 0x45, 0xab, 0x20, 0xf5, 0x39, 0xaa, 0x39, 0x33, 0x86, 0x04, 0xa9, 0xed, 0x0a,
	0xfc, 0x7d, 0xbf, 0x33, 0xe7, 0x7c, 0xe7, 0x00, 0xda, 0x8f, 0xb8, 0x4c, 0xb9, 0x0c, 0xc6, 0x58,
	0x92, 0x40, 0x11, 0x16, 0x93, 0x2c, 0xa5, 0x4c, 0x05, 0xb3, 0x83, 0x31, 0x51, 0xf8, 0x20, 0x50,
	0xe7, 0x82, 0x48, 0x5f, 0x64, 0x5c, 0x71, 0xd7, 0x33, 0xac, 0xaf, 0x59, 0xff, 0x9e, 0xf5, 0x2d,
	0xdb, 0xd9, 0x48, 0x78, 0xc2, 0x01, 0x0d, 0xf4, 0x6f, 0xa6, 0xaa, 0xb3, 0x99, 0x70, 0x9e, 0x4c,
	0x49, 0x00, 0x5f, 0xe3, 0xfc, 0x4b, 0x80, 0xd9, 0xb9, 0xb5, 0xb6, 0xac, 0x85, 0x05, 0x0d, 0x30,
	0x63, 0x5c, 0x61, 0x45, 0x39, 0xb3, 0xed, 0x3a, 0x9d, 0x85, 0x71, 0x44, 0x5f, 0x2c, 0x8e, 0xd2,
	0xd9, 0x5a, 0xf0, 0x40, 0x0f, 0xc6, 0x53, 0x1e, 0x9d, 0xfe, 0xd6, 0x5d, 0xac, 0xf5, 0x16, 0x63,
	0x92, 0x4c, 0x52, 0xce, 0x96, 0xfc, 0xa5, 0x95, 0x7c, 0xcd, 0x49, 0x76, 0x7e, 0xb7, 0x0d, 0x81,
	0x13, 0xca, 0x60, 0xc8, 0x22, 0x9c, 0x61, 0x47, 0x26, 0xb5, 0xdd, 0x8f, 0xb1, 0xba, 0x0f, 0x73,
	0x2b, 0x9a, 0x12, 0xa9, 0x70, 0x2a, 0x0c, 0xb0, 0x7b, 0xed, 0xa0, 0xfa, 0x40, 0x4f, 0xfd, 0x81,
	0x28, 0xec, 0xbe, 0x43, 0x35, 0x88, 0x30, 0xa2, 0x71, 0xdb, 0xd9, 0x71, 0x7a, 0x8d, 0xfe, 0xe6,
	0xe2, 0x8e, 0xcd, 0x80, 0x80, 0x1f, 0x1f, 0x0e, 0x5a, 0x57, 0x37, 0xdd, 0xd2, 0xfc, 0xa6, 0x5b,
	0xb5, 0x42, 0x58, 0x85, 0xda, 0xe3, 0xd8, 0xdd, 0x46, 0xc8, 0x3c, 0x23, 0xe9, 0x05, 0x69, 0xff,
	0xb3, 0xe3, 0xf4, 0x56, 0xc2, 0x3a, 0x28, 0x1f, 0xe9, 0x05, 0x71, 0x0f, 0x51, 0x65, 0x42, 0x70,
	0x4c, 0xb2, 0xf6, 0x0a, 0xf4, 0xd8, 0xf3, 0xff, 0x7c, 0x53, 0xff, 0x08, 0xe8, 0x41, 0x59, 0x37,
	0x0c, 0x6d, 0xad, 0xfb, 0x08, 0x55, 0x59, 0x9e, 0x8e, 0xd4, 0x99, 0x6c, 0x97, 0xa1, 0x43, 0x85,
	0xe5, 0xe9, 0xc9, 0x99, 0xdc, 0xfd, 0x59, 0x46, 0x15, 0x53, 0xe1, 0xbe, 0x41, 0x55, 0xbb, 0x5c,
	0x1b, 0x67, 0x7b, 0xe9, 0x79, 0x63, 0xf9, 0x43, 0xce, 0x24, 0x61, 0x32, 0x97, 0xb6, 0x43, 0x51,
	0xe3, 0xee, 0xa1, 0x5a, 0x34, 0xc1, 0x94, 0xe9, 0x75, 0xe8, 0x14, 0xf5, 0x41, 0x43, 0xe7, 0x1d,
	0x6a, 0x4d, 0xe7, 0x05, 0xf3, 0x38, 0x76, 0xff, 0xd7, 0x81, 0x68, 0x32, 0x51, 0x10, 0x68, 0x25,
	0xb4, 0x5f, 0xee, 0x6b, 0x54, 0xd6, 0xfb, 0x86, 0xf9, 0x1a, 0xfd, 0x8e, 0x6f, 0x8e, 0xe1, 0x17,
	0xc7, 0xf0, 0x4f, 0x8a, 0x63, 0x0c, 0x6a, 0xba, 0xf1, 0xe5, 0xf7, 0xae, 0x13, 0x42, 0x85, 0x3b,
	0x44, 0xcd, 0x29, 0x96, 0x6a, 0x74, 0x77, 0x8d, 0x7f, 0xff, 0x76, 0x0d, 0x33, 0x7a, 0x43, 0x57,
	0x0d, 0xec, 0x19, 0x7a, 0x68, 0x1d, 0x1e, 0x89, 0x78, 0x9a, 0x52, 0x35, 0x9a, 0x60, 0x39, 0x69,
	0x57, 0x76, 0x9c, 0xde, 0x6a, 0xb8, 0xa6, 0xf5, 0x21, 0xc8, 0x47, 0x58, 0x4e, 0xdc, 0xc7, 0xa8,
	0x1e, 0x63, 0x85, 0x0d, 0x52, 0x05, 0xa4, 0xa6, 0x05, 0x30, 0x9f, 0xa1, 0xd6, 0x0c, 0x4f, 0x69,
	0x8c, 0x15, 0xcf, 0xa4, 0x41, 0x6a, 0xe6, 0x95, 0x7b, 0x19, 0xc0, 0x97, 0x68, 0x83, 0x91, 0x33,
	0x35, 0x7a, 0x48, 0xd7, 0x81, 0x76, 0xb5, 0xf7, 0x69, 0xb9, 0xe2, 0x29, 0x5a, 0x8b, 0x8a, 0xe5,
	0x1b, 0x16, 0x01, 0xdb, 0xbc, 0x53, 0x01, 0xdb, 0x44, 0x35, 0x2c, 0x84, 0x01, 0x1a, 0x00, 0x54,
	0xb1, 0x10, 0x60, 0xed, 0xa3, 0xff, 0x20, 0x63, 0x46, 0x64, 0x3e, 0x55, 0xf6, 0x91, 0x55, 0x60,
	0x5a, 0xda, 0x08, 0x8d, 0x0e, 0xec, 0x13, 0xd4, 0x24, 0x33, 0x1a, 0x13, 0x16, 0x11, 0xc3, 0x35,
	0x81, 0x5b, 0x2d, 0x44, 0x80, 0x9e, 0xa3, 0x75, 0x91, 0x71, 0xc1, 0x25, 0xc9, 0x46, 0x38, 0x8e,
	0x33, 0x22, 0x65, 0x7b, 0x4d, 0xdf, 0x3e, 0x6c, 0x15, 0xfa, 0x5b, 0x23, 0x0f, 0xde, 0x5f, 0xcd,
	0x3d, 0xe7, 0x7a, 0xee, 0x39, 0x3f, 0xe6, 0x9e, 0x73, 0x79, 0xeb, 0x95, 0xae, 0x6f, 0xbd, 0xd2,
	0xb7, 0x5b, 0xaf, 0xf4, 0xb9, 0x9f, 0x50, 0x35, 0xc9, 0xc7, 0x7e, 0xc4, 0x53, 0xfb, 0xff, 0x68,
	0x7f, 0xbc, 0x90, 0xf1, 0x69, 0x10, 0x4d, 0x29, 0x61, 0x2a, 0x48, 0x32, 0x11, 0x05, 0x2a, 0x95,
	0x24, 0x9b, 0xd1, 0x88, 0x8c, 0x2b, 0xf0, 0x67, 0xf1, 0xea, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x22, 0x88, 0x44, 0x21, 0x0d, 0x05, 0x00, 0x00,
}

func (m *BlockMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumTxs != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.NumTxs))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.BlockSize != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BlockSize))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.BlockID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProposerAddress) > 0 {
		i -= len(m.ProposerAddress)
		copy(dAtA[i:], m.ProposerAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ProposerAddress)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.EvidenceHash) > 0 {
		i -= len(m.EvidenceHash)
		copy(dAtA[i:], m.EvidenceHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EvidenceHash)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.LastResultsHash) > 0 {
		i -= len(m.LastResultsHash)
		copy(dAtA[i:], m.LastResultsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastResultsHash)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ConsensusHash) > 0 {
		i -= len(m.ConsensusHash)
		copy(dAtA[i:], m.ConsensusHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ConsensusHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.NextValidatorsHash) > 0 {
		i -= len(m.NextValidatorsHash)
		copy(dAtA[i:], m.NextValidatorsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NextValidatorsHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ValidatorsHash) > 0 {
		i -= len(m.ValidatorsHash)
		copy(dAtA[i:], m.ValidatorsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorsHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DataHash) > 0 {
		i -= len(m.DataHash)
		copy(dAtA[i:], m.DataHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DataHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LastCommitHash) > 0 {
		i -= len(m.LastCommitHash)
		copy(dAtA[i:], m.LastCommitHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastCommitHash)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.LastBlockId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintTypes(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BlockID.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.BlockSize != 0 {
		n += 1 + sovTypes(uint64(m.BlockSize))
	}
	l = m.Header.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.NumTxs != 0 {
		n += 1 + sovTypes(uint64(m.NumTxs))
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Version.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	l = m.LastBlockId.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.LastCommitHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DataHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.NextValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ConsensusHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.LastResultsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EvidenceHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ProposerAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockSize", wireType)
			}
			m.BlockSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumTxs", wireType)
			}
			m.NumTxs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumTxs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastBlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastCommitHash = append(m.LastCommitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastCommitHash == nil {
				m.LastCommitHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHash = append(m.DataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.DataHash == nil {
				m.DataHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorsHash = append(m.ValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorsHash == nil {
				m.ValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextValidatorsHash = append(m.NextValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NextValidatorsHash == nil {
				m.NextValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusHash = append(m.ConsensusHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsensusHash == nil {
				m.ConsensusHash = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastResultsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastResultsHash = append(m.LastResultsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastResultsHash == nil {
				m.LastResultsHash = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvidenceHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvidenceHash = append(m.EvidenceHash[:0], dAtA[iNdEx:postIndex]...)
			if m.EvidenceHash == nil {
				m.EvidenceHash = []byte{}
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)