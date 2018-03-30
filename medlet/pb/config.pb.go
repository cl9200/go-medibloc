// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package medletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	NetworkConfig
	ChainConfig
	RPCConfig
	AppConfig
	PprofConfig
	MiscConfig
	StatsConfig
	InfluxdbConfig
*/
package medletpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reporting modules.
type StatsConfig_ReportingModule int32

const (
	StatsConfig_Influxdb StatsConfig_ReportingModule = 0
)

var StatsConfig_ReportingModule_name = map[int32]string{
	0: "Influxdb",
}
var StatsConfig_ReportingModule_value = map[string]int32{
	"Influxdb": 0,
}

func (x StatsConfig_ReportingModule) String() string {
	return proto.EnumName(StatsConfig_ReportingModule_name, int32(x))
}
func (StatsConfig_ReportingModule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorConfig, []int{7, 0}
}

// Med global configurations.
type Config struct {
	// Network config.
	Network *NetworkConfig `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	// Chain config.
	Chain *ChainConfig `protobuf:"bytes,2,opt,name=chain" json:"chain,omitempty"`
	// RPC config.
	Rpc *RPCConfig `protobuf:"bytes,3,opt,name=rpc" json:"rpc,omitempty"`
	// Stats config.
	Stats *StatsConfig `protobuf:"bytes,100,opt,name=stats" json:"stats,omitempty"`
	// Misc config.
	Misc *MiscConfig `protobuf:"bytes,101,opt,name=misc" json:"misc,omitempty"`
	// App Config.
	App *AppConfig `protobuf:"bytes,102,opt,name=app" json:"app,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetNetwork() *NetworkConfig {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Config) GetChain() *ChainConfig {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetStats() *StatsConfig {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Config) GetMisc() *MiscConfig {
	if m != nil {
		return m.Misc
	}
	return nil
}

func (m *Config) GetApp() *AppConfig {
	if m != nil {
		return m.App
	}
	return nil
}

type NetworkConfig struct {
	// Med seed node address.
	Seed []string `protobuf:"bytes,1,rep,name=seed" json:"seed,omitempty"`
	// Listen addresses.
	Listen []string `protobuf:"bytes,2,rep,name=listen" json:"listen,omitempty"`
	// Network node privateKey address. If nil, generate a new node.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Network ID
	NetworkId uint32 `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (m *NetworkConfig) Reset()                    { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()               {}
func (*NetworkConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *NetworkConfig) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *NetworkConfig) GetListen() []string {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *NetworkConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *NetworkConfig) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

type ChainConfig struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// genesis conf file path
	Genesis string `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis,omitempty"`
	// Data dir.
	Datadir string `protobuf:"bytes,11,opt,name=datadir,proto3" json:"datadir,omitempty"`
	// Key dir.
	Keydir string `protobuf:"bytes,12,opt,name=keydir,proto3" json:"keydir,omitempty"`
	// start mine at launch
	StartMine bool `protobuf:"varint,20,opt,name=start_mine,json=startMine,proto3" json:"start_mine,omitempty"`
	// Coinbase.
	Coinbase string `protobuf:"bytes,21,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	// Miner.
	Miner string `protobuf:"bytes,22,opt,name=miner,proto3" json:"miner,omitempty"`
	// Passphrase.
	Passphrase string `protobuf:"bytes,23,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	// Supported signature cipher list. ["ECC_SECP256K1"]
	SignatureCiphers []string `protobuf:"bytes,24,rep,name=signature_ciphers,json=signatureCiphers" json:"signature_ciphers,omitempty"`
}

func (m *ChainConfig) Reset()                    { *m = ChainConfig{} }
func (m *ChainConfig) String() string            { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()               {}
func (*ChainConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *ChainConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ChainConfig) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *ChainConfig) GetDatadir() string {
	if m != nil {
		return m.Datadir
	}
	return ""
}

func (m *ChainConfig) GetKeydir() string {
	if m != nil {
		return m.Keydir
	}
	return ""
}

func (m *ChainConfig) GetStartMine() bool {
	if m != nil {
		return m.StartMine
	}
	return false
}

func (m *ChainConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *ChainConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *ChainConfig) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ChainConfig) GetSignatureCiphers() []string {
	if m != nil {
		return m.SignatureCiphers
	}
	return nil
}

type RPCConfig struct {
	// RPC listen addresses.
	RpcListen []string `protobuf:"bytes,1,rep,name=rpc_listen,json=rpcListen" json:"rpc_listen,omitempty"`
	// HTTP listen addresses.
	HttpListen []string `protobuf:"bytes,2,rep,name=http_listen,json=httpListen" json:"http_listen,omitempty"`
	// Enabled HTTP modules.["api", "admin"]
	HttpModule []string `protobuf:"bytes,3,rep,name=http_module,json=httpModule" json:"http_module,omitempty"`
	// Connection limit.
	ConnectionLimits int32 `protobuf:"varint,4,opt,name=connection_limits,json=connectionLimits,proto3" json:"connection_limits,omitempty"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *RPCConfig) GetRpcListen() []string {
	if m != nil {
		return m.RpcListen
	}
	return nil
}

func (m *RPCConfig) GetHttpListen() []string {
	if m != nil {
		return m.HttpListen
	}
	return nil
}

func (m *RPCConfig) GetHttpModule() []string {
	if m != nil {
		return m.HttpModule
	}
	return nil
}

func (m *RPCConfig) GetConnectionLimits() int32 {
	if m != nil {
		return m.ConnectionLimits
	}
	return 0
}

type AppConfig struct {
	// log level
	LogLevel string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// log file path
	LogFile string `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	// log file age, unit is s.
	LogAge uint32 `protobuf:"varint,3,opt,name=log_age,json=logAge,proto3" json:"log_age,omitempty"`
	// pprof config
	Pprof *PprofConfig `protobuf:"bytes,4,opt,name=pprof" json:"pprof,omitempty"`
	// App version
	Version string `protobuf:"bytes,100,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *AppConfig) Reset()                    { *m = AppConfig{} }
func (m *AppConfig) String() string            { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()               {}
func (*AppConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *AppConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *AppConfig) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *AppConfig) GetLogAge() uint32 {
	if m != nil {
		return m.LogAge
	}
	return 0
}

func (m *AppConfig) GetPprof() *PprofConfig {
	if m != nil {
		return m.Pprof
	}
	return nil
}

func (m *AppConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PprofConfig struct {
	// pprof listen address, if not configured, the function closes.
	HttpListen string `protobuf:"bytes,1,opt,name=http_listen,json=httpListen,proto3" json:"http_listen,omitempty"`
	// cpu profiling file, if not configured, the profiling not start
	Cpuprofile string `protobuf:"bytes,2,opt,name=cpuprofile,proto3" json:"cpuprofile,omitempty"`
	// memory profiling file, if not configured, the profiling not start
	Memprofile string `protobuf:"bytes,3,opt,name=memprofile,proto3" json:"memprofile,omitempty"`
}

func (m *PprofConfig) Reset()                    { *m = PprofConfig{} }
func (m *PprofConfig) String() string            { return proto.CompactTextString(m) }
func (*PprofConfig) ProtoMessage()               {}
func (*PprofConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{5} }

func (m *PprofConfig) GetHttpListen() string {
	if m != nil {
		return m.HttpListen
	}
	return ""
}

func (m *PprofConfig) GetCpuprofile() string {
	if m != nil {
		return m.Cpuprofile
	}
	return ""
}

func (m *PprofConfig) GetMemprofile() string {
	if m != nil {
		return m.Memprofile
	}
	return ""
}

type MiscConfig struct {
	// Default encryption ciper when create new keystore file.
	DefaultKeystoreFileCiper string `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper,omitempty"`
}

func (m *MiscConfig) Reset()                    { *m = MiscConfig{} }
func (m *MiscConfig) String() string            { return proto.CompactTextString(m) }
func (*MiscConfig) ProtoMessage()               {}
func (*MiscConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{6} }

func (m *MiscConfig) GetDefaultKeystoreFileCiper() string {
	if m != nil {
		return m.DefaultKeystoreFileCiper
	}
	return ""
}

type StatsConfig struct {
	// Enable metrics of not.
	EnableMetrics   bool                          `protobuf:"varint,1,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"`
	ReportingModule []StatsConfig_ReportingModule `protobuf:"varint,2,rep,packed,name=reporting_module,json=reportingModule,enum=medletpb.StatsConfig_ReportingModule" json:"reporting_module,omitempty"`
	// Influxdb config.
	Influxdb    *InfluxdbConfig `protobuf:"bytes,11,opt,name=influxdb" json:"influxdb,omitempty"`
	MetricsTags []string        `protobuf:"bytes,12,rep,name=metrics_tags,json=metricsTags" json:"metrics_tags,omitempty"`
}

func (m *StatsConfig) Reset()                    { *m = StatsConfig{} }
func (m *StatsConfig) String() string            { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()               {}
func (*StatsConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{7} }

func (m *StatsConfig) GetEnableMetrics() bool {
	if m != nil {
		return m.EnableMetrics
	}
	return false
}

func (m *StatsConfig) GetReportingModule() []StatsConfig_ReportingModule {
	if m != nil {
		return m.ReportingModule
	}
	return nil
}

func (m *StatsConfig) GetInfluxdb() *InfluxdbConfig {
	if m != nil {
		return m.Influxdb
	}
	return nil
}

func (m *StatsConfig) GetMetricsTags() []string {
	if m != nil {
		return m.MetricsTags
	}
	return nil
}

type InfluxdbConfig struct {
	// Host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Port.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Database name.
	Db string `protobuf:"bytes,3,opt,name=db,proto3" json:"db,omitempty"`
	// Auth user.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// Auth password.
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *InfluxdbConfig) Reset()                    { *m = InfluxdbConfig{} }
func (m *InfluxdbConfig) String() string            { return proto.CompactTextString(m) }
func (*InfluxdbConfig) ProtoMessage()               {}
func (*InfluxdbConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{8} }

func (m *InfluxdbConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *InfluxdbConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *InfluxdbConfig) GetDb() string {
	if m != nil {
		return m.Db
	}
	return ""
}

func (m *InfluxdbConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *InfluxdbConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "medletpb.Config")
	proto.RegisterType((*NetworkConfig)(nil), "medletpb.NetworkConfig")
	proto.RegisterType((*ChainConfig)(nil), "medletpb.ChainConfig")
	proto.RegisterType((*RPCConfig)(nil), "medletpb.RPCConfig")
	proto.RegisterType((*AppConfig)(nil), "medletpb.AppConfig")
	proto.RegisterType((*PprofConfig)(nil), "medletpb.PprofConfig")
	proto.RegisterType((*MiscConfig)(nil), "medletpb.MiscConfig")
	proto.RegisterType((*StatsConfig)(nil), "medletpb.StatsConfig")
	proto.RegisterType((*InfluxdbConfig)(nil), "medletpb.InfluxdbConfig")
	proto.RegisterEnum("medletpb.StatsConfig_ReportingModule", StatsConfig_ReportingModule_name, StatsConfig_ReportingModule_value)
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0xc1, 0x6e, 0xe3, 0x36,
	0x10, 0xad, 0x9d, 0x38, 0xb1, 0xc6, 0x76, 0x36, 0x65, 0xb3, 0x1b, 0xb6, 0x8b, 0xee, 0xa6, 0x02,
	0x02, 0x18, 0x58, 0x20, 0x40, 0xd3, 0x5e, 0x7b, 0x58, 0x18, 0x28, 0x10, 0x24, 0x2e, 0x02, 0xb5,
	0x77, 0x41, 0x96, 0xc6, 0x32, 0x11, 0x59, 0x24, 0x48, 0x3a, 0xbb, 0x41, 0x2f, 0xfd, 0x81, 0x5e,
	0x7a, 0xeb, 0xb9, 0x3f, 0x5a, 0xcc, 0x88, 0xb2, 0x6c, 0x63, 0x6f, 0x9c, 0xf7, 0x1e, 0x87, 0xa3,
	0x79, 0x33, 0x36, 0x8c, 0x73, 0x5d, 0x2f, 0x55, 0x79, 0x63, 0xac, 0xf6, 0x5a, 0x0c, 0xd7, 0x58,
	0x54, 0xe8, 0xcd, 0x22, 0xfe, 0xbb, 0x0f, 0x27, 0x33, 0xa6, 0xc4, 0x8f, 0x70, 0x5a, 0xa3, 0xff,
	0xa4, 0xed, 0x93, 0xec, 0x5d, 0xf5, 0xa6, 0xa3, 0xdb, 0xcb, 0x9b, 0x56, 0x76, 0xf3, 0x5b, 0x43,
	0x34, 0xca, 0xa4, 0xd5, 0x89, 0x0f, 0x30, 0xc8, 0x57, 0x99, 0xaa, 0x65, 0x9f, 0x2f, 0xbc, 0xee,
	0x2e, 0xcc, 0x08, 0x0e, 0xf2, 0x46, 0x23, 0xae, 0xe1, 0xc8, 0x9a, 0x5c, 0x1e, 0xb1, 0xf4, 0x9b,
	0x4e, 0x9a, 0x3c, 0xce, 0x82, 0x90, 0x78, 0xca, 0xe9, 0x7c, 0xe6, 0x9d, 0x2c, 0x0e, 0x73, 0xfe,
	0x4e, 0x70, 0x9b, 0x93, 0x35, 0x62, 0x0a, 0xc7, 0x6b, 0xe5, 0x72, 0x89, 0xac, 0xbd, 0xe8, 0xb4,
	0x73, 0xe5, 0xf2, 0x20, 0x65, 0x05, 0xbd, 0x9e, 0x19, 0x23, 0x97, 0x87, 0xaf, 0x7f, 0x34, 0xa6,
	0x7d, 0x3d, 0x33, 0x26, 0xfe, 0x13, 0x26, 0x7b, 0xdf, 0x2a, 0x04, 0x1c, 0x3b, 0xc4, 0x42, 0xf6,
	0xae, 0x8e, 0xa6, 0x51, 0xc2, 0x67, 0xf1, 0x06, 0x4e, 0x2a, 0xe5, 0x3c, 0xd2, 0x77, 0x13, 0x1a,
	0x22, 0xf1, 0x1e, 0x46, 0xc6, 0xaa, 0xe7, 0xcc, 0x63, 0xfa, 0x84, 0x2f, 0xfc, 0xa5, 0x51, 0x02,
	0x01, 0xba, 0xc7, 0x17, 0xf1, 0x3d, 0x40, 0x68, 0x5d, 0xaa, 0x0a, 0x79, 0x7c, 0xd5, 0x9b, 0x4e,
	0x92, 0x28, 0x20, 0x77, 0x45, 0xfc, 0x4f, 0x1f, 0x46, 0x3b, 0x8d, 0x13, 0xdf, 0xc2, 0x90, 0x5b,
	0x47, 0xe2, 0x1e, 0x8b, 0x4f, 0x39, 0xbe, 0x2b, 0x84, 0x84, 0xd3, 0x12, 0x6b, 0x74, 0xca, 0x71,
	0xef, 0xa3, 0xa4, 0x0d, 0x89, 0x29, 0x32, 0x9f, 0x15, 0xca, 0xca, 0x51, 0xc3, 0x84, 0x90, 0xca,
	0x7e, 0xc2, 0x17, 0x22, 0xc6, 0x4c, 0x84, 0x88, 0xaa, 0x72, 0x3e, 0xb3, 0x3e, 0x5d, 0xab, 0x1a,
	0xe5, 0xc5, 0x55, 0x6f, 0x3a, 0x4c, 0x22, 0x46, 0xe6, 0xaa, 0x46, 0xf1, 0x1d, 0x0c, 0x73, 0xad,
	0xea, 0x45, 0xe6, 0x50, 0xbe, 0xe6, 0x8b, 0xdb, 0x58, 0x5c, 0xc0, 0x80, 0x2e, 0x59, 0xf9, 0x86,
	0x89, 0x26, 0x10, 0xef, 0x00, 0x4c, 0xe6, 0x9c, 0x59, 0x59, 0xba, 0x73, 0x19, 0xda, 0xb0, 0x45,
	0xc4, 0x07, 0xf8, 0xda, 0xa9, 0xb2, 0xce, 0xfc, 0xc6, 0x62, 0x9a, 0x2b, 0xb3, 0x42, 0xeb, 0xa4,
	0xe4, 0x56, 0x9e, 0x6f, 0x89, 0x59, 0x83, 0xc7, 0xff, 0xf6, 0x20, 0xda, 0x8e, 0x08, 0xd5, 0x6a,
	0x4d, 0x9e, 0x86, 0xf6, 0x37, 0xa6, 0x44, 0xd6, 0xe4, 0x0f, 0x5b, 0x07, 0x56, 0xde, 0x9b, 0x74,
	0xcf, 0x1e, 0x20, 0xe8, 0x40, 0xb0, 0xd6, 0xc5, 0xa6, 0x42, 0x79, 0xd4, 0x09, 0xe6, 0x8c, 0x50,
	0x6d, 0xb9, 0xae, 0x6b, 0xcc, 0xbd, 0xd2, 0x75, 0x5a, 0xa9, 0xb5, 0xf2, 0x8e, 0x9d, 0x1a, 0x24,
	0xe7, 0x1d, 0xf1, 0xc0, 0x78, 0xfc, 0x5f, 0x0f, 0xa2, 0xed, 0x00, 0x89, 0xb7, 0x10, 0x55, 0xba,
	0x4c, 0x2b, 0x7c, 0xc6, 0x8a, 0xfd, 0x8a, 0x92, 0x61, 0xa5, 0xcb, 0x07, 0x8a, 0xc9, 0x4b, 0x22,
	0x97, 0xaa, 0xc2, 0xd6, 0xb1, 0x4a, 0x97, 0xbf, 0xaa, 0x0a, 0xc5, 0x25, 0xd0, 0x31, 0xcd, 0x4a,
	0xe4, 0x91, 0x99, 0x24, 0x27, 0x95, 0x2e, 0x3f, 0x96, 0x54, 0xcb, 0xc0, 0x18, 0xab, 0x97, 0xfc,
	0xfe, 0xde, 0x2a, 0x3c, 0x12, 0xdc, 0xae, 0x02, 0x6b, 0xc8, 0xf7, 0x67, 0xb4, 0x4e, 0xe9, 0x9a,
	0x37, 0x27, 0x4a, 0xda, 0x30, 0xae, 0x61, 0xb4, 0xa3, 0x3f, 0xec, 0x51, 0x53, 0xe8, 0x6e, 0x8f,
	0xde, 0x01, 0xe4, 0x66, 0x43, 0x37, 0xba, 0x62, 0x77, 0x10, 0xe2, 0xd7, 0xb8, 0x6e, 0xf9, 0x30,
	0xe5, 0x1d, 0x12, 0xdf, 0x03, 0x74, 0xeb, 0x27, 0x7e, 0x81, 0xb7, 0x05, 0x2e, 0xb3, 0x4d, 0xe5,
	0x69, 0x29, 0x9c, 0xd7, 0x16, 0xb9, 0x0b, 0x64, 0x3c, 0xda, 0xf0, 0xbc, 0x0c, 0x92, 0xfb, 0xa0,
	0xa0, 0xbe, 0xcc, 0x88, 0x8f, 0xff, 0xea, 0xc3, 0x68, 0x67, 0xf1, 0xc5, 0x35, 0x9c, 0x61, 0x9d,
	0x2d, 0x2a, 0x4c, 0xd7, 0xe8, 0xad, 0xca, 0x1d, 0x67, 0x18, 0x26, 0x93, 0x06, 0x9d, 0x37, 0xa0,
	0x78, 0x84, 0x73, 0x8b, 0x46, 0x5b, 0xaf, 0xea, 0xb2, 0x35, 0x9b, 0xa6, 0xe1, 0xec, 0xf6, 0xfa,
	0x8b, 0x3f, 0x28, 0x37, 0x49, 0xab, 0x6e, 0xe6, 0x20, 0x79, 0x65, 0xf7, 0x01, 0xf1, 0x33, 0x0c,
	0x55, 0xbd, 0xac, 0x36, 0x9f, 0x8b, 0x05, 0x2f, 0xd6, 0xe8, 0x56, 0x76, 0x99, 0xee, 0x02, 0x13,
	0x2c, 0xd9, 0x2a, 0xc5, 0x0f, 0x30, 0x0e, 0x75, 0xa6, 0x3e, 0x2b, 0x9d, 0x1c, 0xf3, 0xc0, 0x8d,
	0x02, 0xf6, 0x47, 0x56, 0xba, 0xf8, 0x3d, 0xbc, 0x3a, 0x78, 0x5c, 0x8c, 0x61, 0xd8, 0x66, 0x3c,
	0xff, 0x2a, 0xfe, 0x0c, 0x67, 0xfb, 0xf9, 0xe9, 0x47, 0x69, 0xa5, 0x9d, 0x0f, 0xcd, 0xe3, 0x33,
	0x61, 0x94, 0x84, 0xfd, 0x9a, 0x24, 0x7c, 0x16, 0x67, 0xd0, 0x2f, 0x16, 0xc1, 0xa1, 0x7e, 0xb1,
	0x20, 0xcd, 0xc6, 0xa1, 0xe5, 0x79, 0x8a, 0x12, 0x3e, 0xd3, 0x7a, 0xd3, 0x6a, 0x7e, 0xd2, 0xb6,
	0x90, 0x83, 0x66, 0x68, 0xdb, 0x78, 0x71, 0xc2, 0x7f, 0x17, 0x3f, 0xfd, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x87, 0xe7, 0xf4, 0x99, 0x3e, 0x06, 0x00, 0x00,
}
