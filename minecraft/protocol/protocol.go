package protocol

type Protocol int

const (
	V560 Protocol = 560 // 1.19.50
	V557          = 557 // 1.19.40
	V554          = 554 // 1.19.30
	V545          = 545 // 1.19.21
	V544          = 544 // 1.19.20
	V534          = 534 // 1.19.10
)

var (
	SupportedProtocols = []Protocol{
		V560,
		V557,
		V554,
		V545,
		V544,
		V534,
	}
)
