# Ruun UI Library - Kitchen Sink Demo

This demo showcases how to use the `github.com/niiniyare/ruun` UI component library.

## Running the Demo
```bash
go run main.go
```

Then open http://localhost:3000

## Package Structure
```go
import (
    "github.com/niiniyare/ruun/atoms"      // Basic components
    "github.com/niiniyare/ruun/molecules"  // Combined components
    "github.com/niiniyare/ruun/organisms"  // Complex components
    "github.com/niiniyare/ruun/templates"  // Page layouts
)
```

## Usage Pattern

Everything is a component function call. No raw HTML.
```go
// Layout with Flex
@atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {
    @atoms.Button(atoms.ButtonProps{Text: "Save"})
    @atoms.Button(atoms.ButtonProps{Text: "Cancel", Variant: atoms.ButtonOutline})
}

// Grid layout
@atoms.Grid(atoms.GridProps{Cols: 3, Gap: atoms.Gap6}) {
    @molecules.Card(molecules.CardProps{Title: "Card 1"}) { ... }
    @molecules.Card(molecules.CardProps{Title: "Card 2"}) { ... }
    @molecules.Card(molecules.CardProps{Title: "Card 3"}) { ... }
}
```
