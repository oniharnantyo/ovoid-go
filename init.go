package ovoid_go

type Config struct {
	AppID string
	AppVersion string
	OSName string
	OsVersion string
	MACAddress string
	BaseEndpoint string
	AWSEndpoint string
	TransferOVO string
	TransferBank string
}

var (
	cfg    Config
)

func Init( config Config) {
	cfg = config
}
