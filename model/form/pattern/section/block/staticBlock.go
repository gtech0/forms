package block

type StaticBlock struct {
	Block
	Variants []Variant
}

type StaticBlockSlice []*StaticBlock

func (d *StaticBlockSlice) ToInterface() []IBlock {
	blocks := make([]IBlock, 0)
	for _, staticBlock := range *d {
		blocks = append(blocks, staticBlock)
	}
	return blocks
}
