package core

const (
	// BlockchainPort constatn with the port for the blockchain process
	BlockchainPort = "9090"

	// NetworkKeySize constant with the generic size for key encoding (points, Public IDs, etc...)
	NetworkKeySize = 80

	// TransactionHeaderSize constant with the size of an encoded transaction header
	// from (80 Bytes) + to (80 Bytes) + Type (4 Bytes) + Timestamp (4 Bytes) + PayloadHash (32 Bytes) + PayloadLength (4 Bytes) + Nonce (4 Bytes)
	TransactionHeaderSize = NetworkKeySize + NetworkKeySize + 4 + 32 + 4 + 4 + 4

	// KeySize constant with the size for the crypto keys
	// In our case Elliptic curve P256 with 32 bytes
	KeySize = 32
)
