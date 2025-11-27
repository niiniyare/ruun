# UI Component Refactoring Guide

## Patterns

### Allowed (✅)
- `utils.TwMerge` - resolve Tailwind layout conflicts
- `templ.Classes()` - conditional class composition
- `templ.KV()` - conditional class inclusion
- `templ.Attributes` - spread hx-*, aria-*, data-*
- `templ.Component` - slots, icons, composition
- Basecoat classes: btn, card, field, input, badge
- Tailwind for layout: grid, flex, gap-4

### Not Allowed (❌)
- `ClassName string` prop - leaky abstraction

## Demo Server
```bash
./component-refactor.sh --demo
```
Opens http://localhost:3000 with kitchen-sink component gallery.

## Status
| Phase | Tasks | Status |
|-------|-------|--------|
| 0 Demo Setup | 3 | 🔵 |
| 1 Analysis | 6 | 🔵 |
| 2 Strategy | 5 | 🔵 |
| 3 CSS | 17 | 🔵 |
| 4 Atoms | 25 | 🔵 |
| 5 Molecules | 21 | 🔵 |
| 6 Organisms | 19 | 🔵 |
| 7 Templates | 13 | 🔵 |
| 8 Pages | 15 | 🔵 |
| 9 Validation | 8 | 🔵 |

**Total: 132 tasks**
