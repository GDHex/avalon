package params

// MaxDataSize represents the max data to be allowed to verified
const MaxDataSize int = 1000000

const (
	// Sol = Solidity
	Sol = ".sol"
	// Pdf = Reports
	Pdf = ".pdf"

	// KeyDir is the directory that keypair is stored
	KeyDir = "./keys/"
	// DataDir is the directory that the data is stored
	DataDir = "./data/"
	// SignatureDir is the directory that signatures is stored
	SignatureDir = "./signatures/"
	// DirPrefix is a helper for the moment
	DirPrefix = "./data/"

	// PublicKeySuffix is the filename extension for the public keys
	PublicKeySuffix = "_pblk.sec"
	// PrivateKeySuffix is the filename extension for the private keys
	PrivateKeySuffix = "_prvk.sec"
	// SignatureSuffix is the filename extension for the signatures
	SignatureSuffix = "_sig.sec"
)

// FileTypes represents all the types that avalon can parse
var FileTypes = []string{Sol, Pdf}
