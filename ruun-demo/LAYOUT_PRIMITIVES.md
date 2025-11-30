# Ruun Layout Primitives

The ruun package must export these layout primitives in `atoms/` for pure component-based layouts:

## atoms.Box
Container with padding, margin, background.
```go
type BoxProps struct {
    Padding   Spacing    // P0, P1, P2, P4, P6, P8
    Margin    Spacing
    BgColor   Color      // BgBackground, BgMuted, BgCard
    Border    bool
    Rounded   Rounded    // RoundedNone, RoundedSm, RoundedMd, RoundedLg, RoundedFull
    Shadow    Shadow     // ShadowNone, ShadowSm, ShadowMd, ShadowLg
    Attrs     templ.Attributes
}
```

## atoms.Flex
Flexbox container.
```go
type FlexProps struct {
    Direction   FlexDirection  // Row, Col, RowReverse, ColReverse
    Wrap        FlexWrap       // NoWrap, Wrap, WrapReverse
    Justify     Justify        // JustifyStart, JustifyCenter, JustifyEnd, JustifyBetween, JustifyAround
    Align       Align          // AlignStart, AlignCenter, AlignEnd, AlignStretch, AlignBaseline
    Gap         Spacing        // Gap0, Gap1, Gap2, Gap4, Gap6, Gap8
    Padding     Spacing
    Attrs       templ.Attributes
}
```

## atoms.Grid
CSS Grid container.
```go
type GridProps struct {
    Cols        int            // 1-12
    Rows        int
    Gap         Spacing
    ColGap      Spacing
    RowGap      Spacing
    Padding     Spacing
    Attrs       templ.Attributes
}
```

## atoms.Stack
Vertical or horizontal stack (simplified Flex).
```go
type StackProps struct {
    Direction   StackDirection  // Vertical, Horizontal
    Gap         Spacing
    Align       Align
    Padding     Spacing
    Attrs       templ.Attributes
}
```

## atoms.Center
Centers content both horizontally and vertically.
```go
type CenterProps struct {
    Padding     Spacing
    MaxWidth    MaxWidth  // MaxWidthSm, MaxWidthMd, MaxWidthLg, MaxWidthXl, MaxWidth2xl
    Attrs       templ.Attributes
}
```

## atoms.Spacer
Flexible space for pushing items apart.
```go
type SpacerProps struct {
    Size    Spacing  // Fixed size, or leave empty for flex-grow
}
```

## atoms.Divider
Horizontal or vertical divider line.
```go
type DividerProps struct {
    Orientation  Orientation  // Horizontal, Vertical
    Label        string       // Optional label in middle
}
```

## Spacing Constants
```go
type Spacing int

const (
    Gap0 Spacing = iota  // 0
    Gap1                  // 0.25rem
    Gap2                  // 0.5rem
    Gap3                  // 0.75rem
    Gap4                  // 1rem
    Gap5                  // 1.25rem
    Gap6                  // 1.5rem
    Gap8                  // 2rem
    Gap10                 // 2.5rem
    Gap12                 // 3rem
)

// Aliases
const (
    P0 = Gap0; P1 = Gap1; P2 = Gap2; P4 = Gap4; P6 = Gap6; P8 = Gap8
    M0 = Gap0; M1 = Gap1; M2 = Gap2; M4 = Gap4; M6 = Gap6; M8 = Gap8
)
```

## Usage in Demo
```templ
// Instead of: <div class="flex gap-4 items-center">
@atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
    @atoms.Button(atoms.ButtonProps{Text: "Button 1"})
    @atoms.Button(atoms.ButtonProps{Text: "Button 2"})
}

// Instead of: <div class="grid grid-cols-3 gap-6">
@atoms.Grid(atoms.GridProps{Cols: 3, Gap: atoms.Gap6}) {
    @molecules.Card(...) { ... }
    @molecules.Card(...) { ... }
    @molecules.Card(...) { ... }
}

// Instead of: <div class="space-y-4">
@atoms.Stack(atoms.StackProps{Direction: atoms.Vertical, Gap: atoms.Gap4}) {
    @atoms.Text(...)
    @atoms.Text(...)
}
```
