# Complete Component Refactoring Plan

## Executive Summary

**Objective**: Complete end-to-end refactoring of all component levels to achieve strict atomic design compliance, JSON theme integration, and clean separation of presentation vs. business logic.

---

## 1. Atomic Layer Refactoring

### **Problems Found:**
- Manual class building instead of JSON theme classes
- Custom CSS variables instead of compiled theme classes  
- Missing utils integration (TwMerge, If, IfElse)
- Inconsistent prop patterns
- Missing proper accessibility

### **Required Atoms:**

#### **Core Input Atoms**
```
✅ Button (refactor)
✅ Input (refactor)  
✅ Textarea (refactor)
✅ Select (refactor)
✅ Checkbox (refactor)
✅ Radio (refactor)
✅ Icon (refactor)
➕ Label (create)
➕ Badge (create)
➕ Spinner (create)
➕ Avatar (create)
```

#### **New Atom Principles:**
- **Single Purpose**: Each atom does ONE thing
- **Theme Classes**: Use compiled JSON theme classes (`button`, `button-primary`, etc.)
- **Utils Integration**: TwMerge for class conflicts, If/IfElse for conditionals
- **Clean Props**: Minimal, focused prop interfaces
- **No Business Logic**: Zero conditional rendering based on user/tenant context

---

## 2. Molecule Layer Restructuring

### **Problems Found:**
- FormField doing organism-level work (300+ props)
- Business logic mixed with presentation
- Custom token resolution instead of theme classes

### **Required Molecules:**

#### **Simple Composition Molecules**
```
✅ FormField (simplify to 2-3 atoms max)
➕ SearchBar (Input + Icon + Button)
➕ AlertBox (Icon + Text + Close Button)  
➕ CardHeader (Avatar + Text + Badge)
➕ Breadcrumb (Icon + Link + Separator)
➕ Pagination (Button + Text + Button)
```

#### **New Molecule Principles:**
- **2-3 Atoms Max**: Simple composition only
- **Schema Field Integration**: Accept schema field types for validation display
- **Light Logic Only**: Validation state display, conditional visibility
- **No Business Rules**: No user permissions, tenant logic, data fetching

---

## 3. Organism Layer Implementation

### **Problems Found:**
- Under-utilized layer missing complex business logic
- Forms should be organisms, not molecules

### **Required Organisms:**

#### **Complex Business Components**
```
➕ Form (schema-driven with validation)
➕ DataTable (sorting, filtering, pagination)
➕ Navigation (user context, permissions)
➕ Modal/Dialog (complex interaction patterns)
➕ UserProfile (user data composition)
➕ Dashboard (metrics, widgets composition)
```

#### **New Organism Principles:**
- **Full Business Logic**: User permissions, tenant rules, data validation
- **Schema Integration**: Form schemas, table column definitions
- **Complex State**: HTMX interactions, Alpine.js state management  
- **User Context**: Access to User, Tenant, permissions

---

## 4. Template Layer Structure

### **Required Templates:**

```
➕ DashboardLayout (sidebar, header, main, footer)
➕ AuthLayout (centered forms, branding)
➕ FormLayout (form-specific layouts)
➕ ErrorLayout (error pages)
```

---

## 5. JSON Theme System Integration

### **Implementation Steps:**

#### **A. Theme Compiler Setup**
```bash
# 1. Create theme compiler
cmd/build-themes/main.go

# 2. Default theme
views/style/themes/default.json

# 3. Compile to CSS classes
static/css/themes/default.css
```

#### **B. Component Theme Classes**
```css
/* Compiled from JSON themes */
.button { /* base styles */ }
.button-primary { /* primary variant */ }
.button-md { /* medium size */ }

.input { /* base input */ }
.input-error { /* error state */ }

.alert-success { /* success alert */ }
.card { /* card base */ }
```

#### **C. Replace All Hardcoded Styles**
- Remove all manual class building
- Replace CSS variables with theme classes  
- Use compiled classes throughout

---

## 6. Utils Integration Requirements

### **Update All Components:**

```go
// Old Pattern (Manual)
classes := []string{"inline-flex", "items-center"}
if props.Loading {
    classes = append(classes, "pointer-events-none")  
}

// New Pattern (Utils)
class={ utils.TwMerge(
    "button",
    fmt.Sprintf("button-%s", props.Variant),
    utils.If(props.Loading, "button-loading"),
    props.ClassName,
)}
```

---

## 7. Migration Strategy

### **Phase 1: Foundation (Week 1)**
1. Create JSON theme system + compiler
2. Generate compiled CSS classes
3. Update utils usage patterns

### **Phase 2: Atoms (Week 2)**  
1. Refactor existing atoms (Button, Input, Icon)
2. Create missing atoms (Label, Badge, Spinner)
3. Ensure pure presentation only

### **Phase 3: Molecules (Week 2)**
1. Simplify FormField to basic composition
2. Create new molecules (SearchBar, Alert, Card components)
3. Remove business logic

### **Phase 4: Organisms (Week 3)**
1. Create Form organism with schema integration
2. Build DataTable with business rules  
3. Implement Navigation with user context
4. Create Modal/Dialog patterns

### **Phase 5: Templates & Integration (Week 3)**
1. Create layout templates
2. Update pages to use new architecture
3. Migration scripts for backward compatibility

---

## 8. Validation & Enforcement

### **Automated Checks:**

#### **A. Atom Validation**
```bash
# Check atoms have no business logic
grep -r "user\." views/components/atoms/ && echo "❌ Atoms have business logic"
grep -r "tenant\." views/components/atoms/ && echo "❌ Atoms have business logic"
```

#### **B. Theme Integration Check**
```bash
# Ensure no hardcoded styles
grep -r "bg-blue-500" views/components/ && echo "❌ Hardcoded styles found"
grep -r "text-white" views/components/ && echo "❌ Hardcoded styles found"
```

#### **C. Utils Integration Check**  
```bash
# Ensure utils are used
grep -r "strings.Join.*classes" views/components/ && echo "❌ Manual class building"
```

---

## 9. Success Criteria

### **✅ Atoms:**
- Single purpose, indivisible UI pieces
- JSON theme classes only
- Utils integration throughout  
- No business logic
- Accessibility compliant

### **✅ Molecules:**
- 2-3 atoms maximum
- Schema field integration for validation
- Light logic only
- No business rules

### **✅ Organisms:**
- Complex business logic
- Schema-driven behavior
- User/tenant context
- HTMX/Alpine integration

### **✅ Templates:**
- Layout structures
- Page-level composition
- No business logic

### **✅ System Integration:**
- JSON themes compiled to CSS
- Utils used throughout
- Clean separation of concerns
- Backward compatibility maintained

---

## 10. File Structure (Final)

```
views/components/
├── atoms/           # Pure presentation (Button, Input, Icon, Label, Badge)
├── molecules/       # Simple composition (FormField, SearchBar, Alert) 
├── organisms/       # Business logic (Form, DataTable, Navigation)
├── templates/       # Page layouts (Dashboard, Auth, Form)
├── pages/           # Full page assembly
└── utils/           # Helper functions (TwMerge, If, IfElse)

views/style/themes/   # JSON theme definitions
static/css/themes/    # Compiled CSS classes
cmd/build-themes/     # Theme compiler
```

This refactoring will create a maintainable, scalable, and properly-architected component system that strictly follows atomic design principles while integrating our JSON theme system and utils helpers.