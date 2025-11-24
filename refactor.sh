claude "ultrathink and Execute complete refactoring workflow:
PHASE 1: Analysis

- Load @views/refactoring_guide.md
- Scan all components in @views/components/
- Load @static/css/basecoat.css
- Load @basecoat/docs/src/components as examaple
- Load @static/js/datatable @static/js/basecoat
- Create inventory

PHASE 2: Strategy

- Prioritize components
- Create refactoring order
- Generate templates

PHASE 3: Execute
- Refactor each component \(ask for approval between each\)
- Generate tests
- Update documentation
- Create migration notes
PHASE 4: Validate

- Check compilation
- Verify Basecoat classes exist
- Create visual gallery
Work autonomously but pause for my approval at each phase.
Additional constraints:
- NO utils.TwMerge anywhere
- NO ClassName props of any kind
- NO manual Tailwind utilities that duplicate Basecoat
- Trust @static/css/basecoat.css - read it first
- When in doubt, use data attributes not classes
Check: Does basecoat.css already do this\? \(Answer is usually YES\)"
