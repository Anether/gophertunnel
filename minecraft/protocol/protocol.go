package protocol

type Protocol int

const (
	V560 Protocol = 560
	V557          = 557
	V554          = 554
)

var (
	SupportedProtocols = []Protocol{
		V560,
		V557,
		V554,
	}
)
