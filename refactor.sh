claude "ultrathink and Execute complete refactoring workflow:
PHASE 1: Analysis

- Load @views/refactoring_guide.md
- Scan all components in @views/components/
- Load @static/css/basecoat.css
- Load @basecoat/docs/src/components/ examples (button.njk, field.njk, card.njk, etc.)
- Load @basecoat/src/js/ (basecoat.js, dropdown-menu.js, sidebar.js, etc.)
- Load @basecoat/src/jinja/ templates for JS component integration patterns
- Create inventory
- UPDATE @views/refactoring_guide.md after completing analysis

PHASE 2: Strategy

- Prioritize components
- Create refactoring order
- Generate templates
- UPDATE @views/refactoring_guide.md after completing strategy

PHASE 3: Execute
- Refactor each component \(ask for approval between each\)
- UPDATE @views/refactoring_guide.md after each component refactoring
- Generate tests
- Update documentation
- Create migration notes
- UPDATE @views/refactoring_guide.md after completing execution

PHASE 4: Validate

- Check compilation
- Verify Basecoat classes exist
- Create visual gallery
- UPDATE @views/refactoring_guide.md with final validation results

Work autonomously but pause for my approval at each phase.
Additional constraints:
- NO utils.TwMerge anywhere
- NO ClassName props of any kind
- NO manual Tailwind utilities that duplicate Basecoat
- Trust @static/css/basecoat.css - read it first
- When in doubt, use data attributes not classes
- ALWAYS update @views/refactoring_guide.md to track progress after every task
- Study @basecoat/docs/src/components/ examples for proper implementation patterns
- Reference @basecoat/src/jinja/ templates for JavaScript component integration
- For each component, check if Basecoat already provides it before creating custom CSS
Check: Does basecoat.css already do this\? \(Answer is usually YES\)"
