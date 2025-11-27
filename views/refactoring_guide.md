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

## Phase 1 Analysis Results

✅ **COMPLETED**: Full analysis report available in `./views/analysis-report.md`

### Key Findings
- **36 components analyzed** across atomic design levels
- **75% Basecoat compatibility** - significant acceleration potential
- **4 anti-pattern violations** requiring immediate resolution
- **Progressive enhancement pattern** successfully implemented

### Critical Blockers Identified
1. **ClassName Anti-Patterns** - BaseProps, Icon, ButtonProps contaminated
2. **Compilation Errors** - Navigation/Form templates affected  
3. **Architecture Consistency** - Mixed proper/improper patterns

### Next Phase Requirements
- **Week 1**: Fix anti-patterns and compilation errors
- **Weeks 2-4**: Basecoat CSS integration (75% of components)  
- **Weeks 5-8**: Custom component development (25% remaining)

## Status
| Phase | Tasks | Status | Completion |
|-------|-------|--------|------------|
| 0 Demo Setup | 3 | ✅ | 100% |
| 1 Analysis | 6 | ✅ | 100% |
| 2 Strategy | 5 | 🟡 | Ready to start |
| 3 CSS | 17 | 🔵 | Waiting |
| 4 Atoms | 25 | 🔵 | Waiting |
| 5 Molecules | 21 | 🔵 | Waiting |
| 6 Organisms | 19 | 🔵 | Waiting |
| 7 Templates | 13 | 🔵 | Waiting |
| 8 Pages | 15 | 🔵 | Waiting |
| 9 Validation | 8 | 🔵 | Waiting |

**Phase 1 Complete**: 9/132 tasks (7%)  
**Total Estimated Effort**: 7-9 weeks for complete modernization
