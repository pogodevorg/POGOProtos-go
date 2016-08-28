// Code generated by protoc-gen-go.
// source: map_pokemon.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ActivityType from enums.proto

// Ignoring public import of BadgeType from enums.proto

// Ignoring public import of CameraInterpolation from enums.proto

// Ignoring public import of CameraTarget from enums.proto

// Ignoring public import of Gender from enums.proto

// Ignoring public import of HoloIapItemCategory from enums.proto

// Ignoring public import of ItemCategory from enums.proto

// Ignoring public import of ItemEffect from enums.proto

// Ignoring public import of Platform from enums.proto

// Ignoring public import of PokemonFamilyId from enums.proto

// Ignoring public import of PokemonId from enums.proto

// Ignoring public import of PokemonMove from enums.proto

// Ignoring public import of PokemonMovementType from enums.proto

// Ignoring public import of PokemonRarity from enums.proto

// Ignoring public import of PokemonType from enums.proto

// Ignoring public import of TeamColor from enums.proto

// Ignoring public import of TutorialState from enums.proto

// Ignoring public import of AssetDigestEntry from data.proto

// Ignoring public import of BuddyPokemon from data.proto

// Ignoring public import of DownloadUrlEntry from data.proto

// Ignoring public import of PlayerBadge from data.proto

// Ignoring public import of PlayerData from data.proto

// Ignoring public import of PokedexEntry from data.proto

// Ignoring public import of PokemonData from data.proto

type MapPokemon struct {
	SpawnPointId string    `protobuf:"bytes,1,opt,name=spawn_point_id,json=spawnPointId" json:"spawn_point_id,omitempty"`
	EncounterId  uint64    `protobuf:"fixed64,2,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	PokemonId    PokemonId `protobuf:"varint,3,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	// After this timestamp, the pokemon will be gone.
	ExpirationTimestampMs int64   `protobuf:"varint,4,opt,name=expiration_timestamp_ms,json=expirationTimestampMs" json:"expiration_timestamp_ms,omitempty"`
	Latitude              float64 `protobuf:"fixed64,5,opt,name=latitude" json:"latitude,omitempty"`
	Longitude             float64 `protobuf:"fixed64,6,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *MapPokemon) Reset()                    { *m = MapPokemon{} }
func (m *MapPokemon) String() string            { return proto.CompactTextString(m) }
func (*MapPokemon) ProtoMessage()               {}
func (*MapPokemon) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type NearbyPokemon struct {
	PokemonId        PokemonId `protobuf:"varint,1,opt,name=pokemon_id,json=pokemonId,enum=POGOProtos.Enums.PokemonId" json:"pokemon_id,omitempty"`
	DistanceInMeters float32   `protobuf:"fixed32,2,opt,name=distance_in_meters,json=distanceInMeters" json:"distance_in_meters,omitempty"`
	EncounterId      uint64    `protobuf:"fixed64,3,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	FortId           string    `protobuf:"bytes,4,opt,name=fort_id,json=fortId" json:"fort_id,omitempty"`
	FortImageUrl     string    `protobuf:"bytes,5,opt,name=fort_image_url,json=fortImageUrl" json:"fort_image_url,omitempty"`
}

func (m *NearbyPokemon) Reset()                    { *m = NearbyPokemon{} }
func (m *NearbyPokemon) String() string            { return proto.CompactTextString(m) }
func (*NearbyPokemon) ProtoMessage()               {}
func (*NearbyPokemon) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

type WildPokemon struct {
	EncounterId             uint64       `protobuf:"fixed64,1,opt,name=encounter_id,json=encounterId" json:"encounter_id,omitempty"`
	LastModifiedTimestampMs int64        `protobuf:"varint,2,opt,name=last_modified_timestamp_ms,json=lastModifiedTimestampMs" json:"last_modified_timestamp_ms,omitempty"`
	Latitude                float64      `protobuf:"fixed64,3,opt,name=latitude" json:"latitude,omitempty"`
	Longitude               float64      `protobuf:"fixed64,4,opt,name=longitude" json:"longitude,omitempty"`
	SpawnPointId            string       `protobuf:"bytes,5,opt,name=spawn_point_id,json=spawnPointId" json:"spawn_point_id,omitempty"`
	PokemonData             *PokemonData `protobuf:"bytes,7,opt,name=pokemon_data,json=pokemonData" json:"pokemon_data,omitempty"`
	TimeTillHiddenMs        int32        `protobuf:"varint,11,opt,name=time_till_hidden_ms,json=timeTillHiddenMs" json:"time_till_hidden_ms,omitempty"`
}

func (m *WildPokemon) Reset()                    { *m = WildPokemon{} }
func (m *WildPokemon) String() string            { return proto.CompactTextString(m) }
func (*WildPokemon) ProtoMessage()               {}
func (*WildPokemon) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *WildPokemon) GetPokemonData() *PokemonData {
	if m != nil {
		return m.PokemonData
	}
	return nil
}

func init() {
	proto.RegisterType((*MapPokemon)(nil), "POGOProtos.Map.Pokemon.MapPokemon")
	proto.RegisterType((*NearbyPokemon)(nil), "POGOProtos.Map.Pokemon.NearbyPokemon")
	proto.RegisterType((*WildPokemon)(nil), "POGOProtos.Map.Pokemon.WildPokemon")
}

func init() { proto.RegisterFile("map_pokemon.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xce, 0x47, 0x9b, 0x71, 0xa8, 0x8a, 0x11, 0x24, 0x0a, 0x3d, 0x40, 0xc4, 0x81, 0x03,
	0xe4, 0x50, 0x24, 0x0e, 0x70, 0x40, 0x42, 0x20, 0xc8, 0xc1, 0xd4, 0xb2, 0x8a, 0x90, 0xb8, 0xac,
	0xb6, 0xdd, 0x6d, 0x59, 0xb1, 0x5f, 0xf2, 0x6e, 0x04, 0x9c, 0xf9, 0x5d, 0xfc, 0x12, 0xfe, 0x0c,
	0xb3, 0x6b, 0xbb, 0x69, 0x9b, 0x80, 0xd4, 0xdb, 0xec, 0x7b, 0xcf, 0xbb, 0xf3, 0xe6, 0x8d, 0xe1,
	0x8e, 0xa2, 0x96, 0x58, 0xf3, 0x8d, 0x2b, 0xa3, 0x17, 0xb6, 0x36, 0xde, 0xe4, 0xf7, 0xcb, 0xa3,
	0xf7, 0x47, 0x65, 0x28, 0xdd, 0xa2, 0xa0, 0x76, 0x51, 0x36, 0xec, 0x2c, 0xe3, 0x7a, 0xa5, 0x5c,
	0x23, 0x9a, 0x01, 0xa3, 0x9e, 0x36, 0xf5, 0xfc, 0x57, 0x0a, 0x80, 0xc2, 0x56, 0x97, 0x3f, 0x86,
	0x3d, 0x67, 0xe9, 0x77, 0x8d, 0xd7, 0x0a, 0xed, 0x89, 0x60, 0xd3, 0xe4, 0x61, 0xf2, 0x64, 0x54,
	0x8d, 0x23, 0x5a, 0x06, 0x70, 0xc9, 0xf2, 0x47, 0x30, 0xe6, 0xfa, 0xd4, 0xac, 0xb4, 0xe7, 0x75,
	0xd0, 0xa4, 0xa8, 0x19, 0x56, 0xd9, 0x05, 0x86, 0x92, 0x97, 0x00, 0x6d, 0x67, 0x41, 0xd0, 0x43,
	0xc1, 0xde, 0xe1, 0x83, 0xc5, 0xa5, 0xee, 0xde, 0xc5, 0x86, 0xda, 0x77, 0x97, 0xac, 0x1a, 0xd9,
	0xae, 0xcc, 0x5f, 0xc0, 0x84, 0xff, 0xb0, 0xa2, 0xa6, 0x5e, 0xe0, 0xe7, 0x5e, 0x28, 0xee, 0x3c,
	0x55, 0x96, 0x28, 0x37, 0xed, 0xe3, 0x45, 0xbd, 0xea, 0xde, 0x9a, 0x3e, 0xee, 0xd8, 0xc2, 0xe5,
	0x33, 0xd8, 0x95, 0x08, 0xfa, 0x15, 0xe3, 0xd3, 0x01, 0x0a, 0x93, 0xea, 0xe2, 0x9c, 0x1f, 0xc0,
	0x48, 0x1a, 0x7d, 0xde, 0x90, 0xc3, 0x48, 0xae, 0x81, 0xf9, 0x9f, 0x04, 0x6e, 0x7f, 0xe4, 0xb4,
	0x3e, 0xf9, 0xd9, 0x0d, 0xe2, 0x6a, 0xff, 0xc9, 0x8d, 0xfa, 0x7f, 0x0a, 0x39, 0x13, 0xd8, 0x94,
	0x3e, 0xe5, 0x44, 0x68, 0xa2, 0x38, 0x8e, 0xc4, 0xc5, 0x21, 0xa5, 0xd5, 0x7e, 0xc7, 0x2c, 0x75,
	0x11, 0xf1, 0x8d, 0x61, 0xf6, 0x36, 0x87, 0x39, 0x81, 0x9d, 0x33, 0x53, 0xc7, 0x38, 0xfa, 0x31,
	0x8e, 0x61, 0x38, 0x22, 0x81, 0x71, 0x35, 0x84, 0xa2, 0xe7, 0x9c, 0xac, 0x6a, 0x19, 0x7d, 0x63,
	0x5c, 0x91, 0x0f, 0xe0, 0xa7, 0x5a, 0xce, 0x7f, 0xa7, 0x90, 0x7d, 0x16, 0x92, 0x75, 0xde, 0xae,
	0xbf, 0x98, 0x6c, 0xbe, 0xf8, 0x0a, 0x66, 0x92, 0x3a, 0x4f, 0x94, 0x61, 0xe2, 0x4c, 0x70, 0x76,
	0x35, 0x85, 0x34, 0xa6, 0x30, 0x09, 0x8a, 0xa2, 0x15, 0xfc, 0x2b, 0x87, 0xde, 0xff, 0x72, 0xe8,
	0x5f, 0xcb, 0x61, 0xcb, 0xfa, 0x0d, 0xb6, 0xac, 0xdf, 0x6b, 0x18, 0x77, 0xd9, 0x84, 0x4d, 0x9e,
	0xee, 0xa0, 0x26, 0x3b, 0x3c, 0xb8, 0x9c, 0xce, 0xdb, 0xb0, 0xe1, 0xad, 0xdf, 0x50, 0x57, 0x99,
	0x5d, 0x1f, 0xf2, 0x67, 0x70, 0x37, 0xf8, 0x41, 0x53, 0x52, 0x92, 0xaf, 0x82, 0x31, 0xae, 0x83,
	0xad, 0x0c, 0xef, 0x19, 0x54, 0xfb, 0x81, 0x3a, 0x46, 0xe6, 0x43, 0x24, 0x0a, 0xf7, 0x66, 0xf7,
	0xcb, 0x30, 0xfe, 0x2c, 0xae, 0xbc, 0x55, 0x26, 0x27, 0x4d, 0xfd, 0xfc, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xed, 0xe3, 0x16, 0x7e, 0x7e, 0x03, 0x00, 0x00,
}
