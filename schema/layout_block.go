package schema

// IsValid checks if the BlockType is valid (extended method)
func (b BlockType) IsValid() bool {
	switch b {
	case BlockTypeSection, BlockTypeGroup, BlockTypeTab, BlockTypeStep, BlockTypeCard, BlockTypePanel:
		return true
	default:
		return false
	}
}

// Builder
func NewLayoutBlock(id string, blockType BlockType) *LayoutBlockBuilder {
	return &LayoutBlockBuilder{block: LayoutBlock{ID: id, Type: blockType}}
}

type LayoutBlockBuilder struct {
	block LayoutBlock
}

func (b *LayoutBlockBuilder) Title(t string) *LayoutBlockBuilder {
	b.block.Title = t
	return b
}

func (b *LayoutBlockBuilder) Description(d string) *LayoutBlockBuilder {
	b.block.Description = d
	return b
}

func (b *LayoutBlockBuilder) Icon(i string) *LayoutBlockBuilder {
	b.block.Icon = i
	return b
}

func (b *LayoutBlockBuilder) Badge(bg string) *LayoutBlockBuilder {
	b.block.Badge = bg
	return b
}

func (b *LayoutBlockBuilder) Fields(f ...string) *LayoutBlockBuilder {
	b.block.Fields = f
	return b
}

func (b *LayoutBlockBuilder) Blocks(blocks ...LayoutBlock) *LayoutBlockBuilder {
	b.block.Blocks = blocks
	return b
}

func (b *LayoutBlockBuilder) Order(o int) *LayoutBlockBuilder {
	b.block.Order = o
	return b
}

func (b *LayoutBlockBuilder) Disabled(d bool) *LayoutBlockBuilder {
	b.block.Disabled = d
	return b
}

func (b *LayoutBlockBuilder) Hidden(h bool) *LayoutBlockBuilder {
	b.block.Hidden = h
	return b
}

func (b *LayoutBlockBuilder) Collapsible(c bool) *LayoutBlockBuilder {
	b.block.Collapsible = c
	return b
}

func (b *LayoutBlockBuilder) Collapsed(c bool) *LayoutBlockBuilder {
	b.block.Collapsed = c
	return b
}

func (b *LayoutBlockBuilder) Skippable(s bool) *LayoutBlockBuilder {
	b.block.Skippable = s
	return b
}

func (b *LayoutBlockBuilder) Validation(v bool) *LayoutBlockBuilder {
	b.block.Validation = v
	return b
}

func (b *LayoutBlockBuilder) Columns(c int) *LayoutBlockBuilder {
	b.block.Columns = c
	return b
}

func (b *LayoutBlockBuilder) Border(border bool) *LayoutBlockBuilder {
	b.block.Border = border
	return b
}

func (b *LayoutBlockBuilder) Conditional(c *Conditional) *LayoutBlockBuilder {
	b.block.Conditional = c
	return b
}

func (b *LayoutBlockBuilder) Style(s *Style) *LayoutBlockBuilder {
	b.block.Style = s
	return b
}

func (b *LayoutBlockBuilder) Build() LayoutBlock {
	return b.block
}
