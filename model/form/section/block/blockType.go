package block

type BlockType string

const (
	STATIC   BlockType = "STATIC"
	DYNAMIC  BlockType = "DYNAMIC"
	EXISTING BlockType = "EXISTING"
)
