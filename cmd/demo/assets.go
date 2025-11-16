package main

// getEmbeddedCSS returns the CSS for the demo
func getEmbeddedCSS() string {
	return `
/* ═══════════════════════════════════════════════════════════════════════════ */
/* DESIGN TOKEN SYSTEM - Three-Tier Architecture */
/* ═══════════════════════════════════════════════════════════════════════════ */

:root {
    /* PRIMITIVE TOKENS - Foundation Layer */
    --color-blue-50: #eff6ff;
    --color-blue-500: #3b82f6;
    --color-blue-600: #2563eb;
    --color-blue-700: #1d4ed8;
    
    --color-gray-50: #f9fafb;
    --color-gray-100: #f3f4f6;
    --color-gray-200: #e5e7eb;
    --color-gray-300: #d1d5db;
    --color-gray-400: #9ca3af;
    --color-gray-500: #6b7280;
    --color-gray-600: #4b5563;
    --color-gray-700: #374151;
    --color-gray-800: #1f2937;
    --color-gray-900: #111827;
    
    --color-green-50: #f0fdf4;
    --color-green-500: #22c55e;
    --color-green-600: #16a34a;
    
    --color-red-50: #fef2f2;
    --color-red-500: #ef4444;
    --color-red-600: #dc2626;
    
    --color-yellow-50: #fefce8;
    --color-yellow-500: #eab308;
    
    --color-white: #ffffff;
    --color-black: #000000;
    
    /* SPACING */
    --spacing-xs: 0.25rem;
    --spacing-sm: 0.5rem;
    --spacing-md: 1rem;
    --spacing-lg: 1.5rem;
    --spacing-xl: 2rem;
    --spacing-2xl: 3rem;
    
    /* TYPOGRAPHY */
    --font-size-xs: 0.75rem;
    --font-size-sm: 0.875rem;
    --font-size-base: 1rem;
    --font-size-lg: 1.125rem;
    --font-size-xl: 1.25rem;
    --font-size-2xl: 1.5rem;
    
    --font-weight-normal: 400;
    --font-weight-medium: 500;
    --font-weight-semibold: 600;
    --font-weight-bold: 700;
    
    --line-height-normal: 1.5;
    
    /* BORDERS */
    --border-radius-sm: 0.125rem;
    --border-radius-md: 0.375rem;
    --border-radius-lg: 0.5rem;
    --border-width-thin: 1px;
    
    /* ANIMATIONS */
    --animation-duration-fast: 150ms;
    --animation-duration-normal: 300ms;
    --animation-easing: cubic-bezier(0.4, 0, 0.2, 1);
}

/* SEMANTIC TOKENS - Functional Meaning Layer */
:root {
    --semantic-bg-default: var(--color-white);
    --semantic-bg-subtle: var(--color-gray-50);
    --semantic-bg-emphasis: var(--color-gray-100);
    
    --semantic-text-default: var(--color-gray-900);
    --semantic-text-subtle: var(--color-gray-600);
    --semantic-text-disabled: var(--color-gray-400);
    --semantic-text-inverted: var(--color-white);
    
    --semantic-border-default: var(--color-gray-200);
    --semantic-border-focus: var(--color-blue-500);
    --semantic-border-strong: var(--color-gray-300);
    --semantic-border-subtle: var(--color-gray-100);
    
    --semantic-interactive-primary-default: var(--color-blue-600);
    --semantic-interactive-primary-hover: var(--color-blue-700);
    
    --semantic-feedback-success-default: var(--color-green-600);
    --semantic-feedback-success-subtle: var(--color-green-50);
    --semantic-feedback-error-default: var(--color-red-600);
    --semantic-feedback-error-subtle: var(--color-red-50);
    --semantic-feedback-warning-default: var(--color-yellow-500);
    --semantic-feedback-warning-subtle: var(--color-yellow-50);
    --semantic-feedback-info-default: var(--color-blue-600);
    --semantic-feedback-info-subtle: var(--color-blue-50);
}

/* DARK MODE SEMANTIC TOKENS */
[data-theme="dark"] {
    --semantic-bg-default: var(--color-gray-900);
    --semantic-bg-subtle: var(--color-gray-800);
    --semantic-bg-emphasis: var(--color-gray-700);
    
    --semantic-text-default: var(--color-gray-50);
    --semantic-text-subtle: var(--color-gray-300);
    --semantic-text-disabled: var(--color-gray-500);
    
    --semantic-border-default: var(--color-gray-700);
    --semantic-border-strong: var(--color-gray-600);
    --semantic-border-subtle: var(--color-gray-800);
}

/* COMPONENT TOKENS */
:root {
    --component-input-bg: var(--semantic-bg-default);
    --component-input-border: var(--semantic-border-default);
    --component-input-border-focus: var(--semantic-border-focus);
    --component-input-border-error: var(--semantic-feedback-error-default);
    --component-input-text: var(--semantic-text-default);
    --component-input-placeholder: var(--semantic-text-subtle);
    
    --component-button-primary-bg: var(--semantic-interactive-primary-default);
    --component-button-primary-bg-hover: var(--semantic-interactive-primary-hover);
    --component-button-primary-text: var(--semantic-text-inverted);
}

/* ═══════════════════════════════════════════════════════════════════════════ */
/* BASE STYLES */
/* ═══════════════════════════════════════════════════════════════════════════ */

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif;
    font-size: var(--font-size-base);
    line-height: var(--line-height-normal);
    color: var(--semantic-text-default);
    background-color: var(--semantic-bg-default);
    transition: background-color var(--animation-duration-normal) var(--animation-easing),
               color var(--animation-duration-normal) var(--animation-easing);
}

/* ═══════════════════════════════════════════════════════════════════════════ */
/* LAYOUT COMPONENTS */
/* ═══════════════════════════════════════════════════════════════════════════ */

.demo-container {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
}

.demo-header {
    background-color: var(--semantic-bg-emphasis);
    border-bottom: var(--border-width-thin) solid var(--semantic-border-default);
    padding: var(--spacing-lg) var(--spacing-xl);
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: var(--spacing-md);
}

.demo-title {
    font-size: var(--font-size-2xl);
    font-weight: var(--font-weight-bold);
    color: var(--semantic-text-default);
}

.demo-subtitle {
    font-size: var(--font-size-sm);
    color: var(--semantic-text-subtle);
    margin-top: var(--spacing-xs);
}

.demo-controls {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
}

.demo-main {
    flex: 1;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 0;
    min-height: 0;
}

@media (max-width: 1024px) {
    .demo-main {
        grid-template-columns: 1fr;
    }
}

.demo-panel {
    padding: var(--spacing-xl);
    border-right: var(--border-width-thin) solid var(--semantic-border-default);
    overflow-y: auto;
    min-height: 0;
}

.demo-panel:last-child {
    border-right: none;
}

.panel-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: var(--spacing-xl);
    padding-bottom: var(--spacing-md);
    border-bottom: var(--border-width-thin) solid var(--semantic-border-subtle);
}

.panel-title {
    font-size: var(--font-size-lg);
    font-weight: var(--font-weight-semibold);
    color: var(--semantic-text-default);
}

.panel-badge {
    background-color: var(--semantic-interactive-primary-default);
    color: var(--semantic-text-inverted);
    font-size: var(--font-size-xs);
    font-weight: var(--font-weight-medium);
    padding: var(--spacing-xs) var(--spacing-sm);
    border-radius: var(--border-radius-sm);
}

/* ═══════════════════════════════════════════════════════════════════════════ */
/* FORM COMPONENTS - Integrates with Templ Components */
/* ═══════════════════════════════════════════════════════════════════════════ */

.form-container {
    background-color: var(--semantic-bg-default);
}

/* FormField component styling */
.form-field {
    margin-bottom: var(--spacing-lg);
    transition: all var(--animation-duration-normal) var(--animation-easing);
}

.form-field.validation-state-validating {
    opacity: 0.8;
}

.form-field.validation-state-valid {
    border-left: 3px solid var(--semantic-feedback-success-default);
    padding-left: var(--spacing-sm);
}

.form-field.validation-state-invalid {
    border-left: 3px solid var(--semantic-feedback-error-default);
    padding-left: var(--spacing-sm);
    animation: field-error var(--animation-duration-fast) ease-out;
}

@keyframes field-error {
    0%, 100% { transform: translateX(0); }
    25% { transform: translateX(-2px); }
    75% { transform: translateX(2px); }
}

/* Schema Form component styling */
.schema-form {
    max-width: none;
}

/* Validation indicators */
.validation-indicator {
    position: absolute;
    right: var(--spacing-sm);
    top: 50%;
    transform: translateY(-50%);
    font-size: var(--font-size-sm);
    pointer-events: none;
}

.validation-state-validating .validation-indicator::before {
    content: '⟲';
    animation: spin 1s linear infinite;
    color: var(--semantic-feedback-info-default);
}

.validation-state-valid .validation-indicator::before {
    content: '✓';
    color: var(--semantic-feedback-success-default);
}

.validation-state-invalid .validation-indicator::before {
    content: '✗';
    color: var(--semantic-feedback-error-default);
}

.validation-state-warning .validation-indicator::before {
    content: '⚠';
    color: var(--semantic-feedback-warning-default);
}

@keyframes spin {
    from { transform: translateY(-50%) rotate(0deg); }
    to { transform: translateY(-50%) rotate(360deg); }
}

/* ═══════════════════════════════════════════════════════════════════════════ */
/* BUTTON COMPONENTS */
/* ═══════════════════════════════════════════════════════════════════════════ */

.btn {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: var(--spacing-xs);
    font-size: var(--font-size-sm);
    font-weight: var(--font-weight-medium);
    padding: var(--spacing-sm) var(--spacing-lg);
    border: none;
    border-radius: var(--border-radius-md);
    cursor: pointer;
    transition: all var(--animation-duration-fast) var(--animation-easing);
    text-decoration: none;
}

.btn-primary {
    background-color: var(--component-button-primary-bg);
    color: var(--component-button-primary-text);
}

.btn-primary:hover {
    background-color: var(--component-button-primary-bg-hover);
}

.btn-secondary {
    background-color: transparent;
    color: var(--semantic-text-default);
    border: var(--border-width-thin) solid var(--semantic-border-default);
}

.btn-secondary:hover {
    background-color: var(--semantic-bg-subtle);
}

.theme-toggle {
    background: none;
    border: var(--border-width-thin) solid var(--semantic-border-default);
    border-radius: var(--border-radius-md);
    padding: var(--spacing-sm);
    cursor: pointer;
    color: var(--semantic-text-default);
    font-size: var(--font-size-lg);
    transition: all var(--animation-duration-fast) var(--animation-easing);
}

.theme-toggle:hover {
    background-color: var(--semantic-bg-subtle);
}

/* ═══════════════════════════════════════════════════════════════════════════ */
/* STATE VIEWER */
/* ═══════════════════════════════════════════════════════════════════════════ */

.state-viewer {
    background-color: var(--semantic-bg-subtle);
    border: var(--border-width-thin) solid var(--semantic-border-default);
    border-radius: var(--border-radius-md);
    padding: var(--spacing-md);
    margin-bottom: var(--spacing-lg);
}

.state-viewer h3 {
    font-size: var(--font-size-sm);
    font-weight: var(--font-weight-semibold);
    color: var(--semantic-text-default);
    margin-bottom: var(--spacing-sm);
}

.state-viewer pre {
    background-color: var(--semantic-bg-default);
    border-radius: var(--border-radius-sm);
    padding: var(--spacing-sm);
    font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
    font-size: var(--font-size-xs);
    line-height: var(--line-height-normal);
    color: var(--semantic-text-default);
    overflow-x: auto;
    white-space: pre-wrap;
    word-break: break-word;
    max-height: 300px;
    overflow-y: auto;
}

/* ═══════════════════════════════════════════════════════════════════════════ */
/* HTMX LOADING STATES */
/* ═══════════════════════════════════════════════════════════════════════════ */

.htmx-request {
    opacity: 0.7;
    transition: opacity var(--animation-duration-fast);
}

.htmx-indicator {
    display: inline-block;
    width: 1rem;
    height: 1rem;
    border: 2px solid var(--semantic-border-default);
    border-radius: 50%;
    border-top-color: var(--semantic-interactive-primary-default);
    animation: spin 1s linear infinite;
}

/* ═══════════════════════════════════════════════════════════════════════════ */
/* RESPONSIVE DESIGN */
/* ═══════════════════════════════════════════════════════════════════════════ */

@media (max-width: 1024px) {
    .demo-panel {
        border-right: none;
        border-bottom: var(--border-width-thin) solid var(--semantic-border-default);
    }
    
    .demo-panel:last-child {
        border-bottom: none;
    }
}

@media (max-width: 768px) {
    .demo-header {
        flex-direction: column;
        align-items: stretch;
        gap: var(--spacing-sm);
    }
    
    .demo-controls {
        justify-content: center;
    }
    
    .demo-panel {
        padding: var(--spacing-md);
    }
    
    .panel-header {
        flex-direction: column;
        align-items: flex-start;
        gap: var(--spacing-sm);
    }
}
	`
}

// getEmbeddedJS returns the JavaScript for the demo
func getEmbeddedJS() string {
	return `
// ═══════════════════════════════════════════════════════════════════════════
// THEME MANAGEMENT
// ═══════════════════════════════════════════════════════════════════════════

function toggleTheme() {
    const html = document.documentElement;
    const currentTheme = html.getAttribute('data-theme') || 'light';
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    
    html.setAttribute('data-theme', newTheme);
    
    // Update icon
    const icon = document.getElementById('theme-icon');
    icon.textContent = newTheme === 'dark' ? '☀️' : '🌙';
    
    // Save preference
    localStorage.setItem('theme-preference', newTheme);
    
    // Update HTMX headers for theme context
    htmx.config.defaultRequestHeaders['X-Theme'] = newTheme;
    
    // Reload form with new theme
    htmx.ajax('GET', '/form', {target:'#form-container'});
}

// Initialize theme
function initializeTheme() {
    const savedTheme = localStorage.getItem('theme-preference') || 'light';
    document.documentElement.setAttribute('data-theme', savedTheme);
    document.getElementById('theme-icon').textContent = savedTheme === 'dark' ? '☀️' : '🌙';
    
    // Set HTMX header
    htmx.config.defaultRequestHeaders = htmx.config.defaultRequestHeaders || {};
    htmx.config.defaultRequestHeaders['X-Theme'] = savedTheme;
}

// ═══════════════════════════════════════════════════════════════════════════
// FORM MANAGEMENT
// ═══════════════════════════════════════════════════════════════════════════

function resetForm() {
    // Clear all form fields
    const form = document.querySelector('#form-container form');
    if (form) {
        form.reset();
        
        // Clear validation states
        const fields = form.querySelectorAll('.form-field');
        fields.forEach(field => {
            field.classList.remove('validation-state-valid', 'validation-state-invalid', 
                'validation-state-warning', 'validation-state-validating');
        });
    }
    
    // Reload fresh form
    htmx.ajax('GET', '/form', {target:'#form-container'});
}

// ═══════════════════════════════════════════════════════════════════════════
// HTMX EVENT HANDLERS
// ═══════════════════════════════════════════════════════════════════════════

// Handle field validation events
document.addEventListener('htmx:trigger', function(e) {
    if (e.detail.fieldValidated) {
        const {field, state, hasErrors} = e.detail.fieldValidated;
        console.log('🔍 Field validated:', {field, state, hasErrors});
    }
});

// Handle form submission results
document.addEventListener('htmx:afterRequest', function(e) {
    if (e.detail.xhr.status === 200 && e.target.closest('form')) {
        try {
            const response = JSON.parse(e.detail.xhr.responseText);
            if (response.success) {
                showNotification('✅ Form validation successful!', 'success');
            } else {
                showNotification('❌ Please fix validation errors', 'error');
            }
        } catch (ex) {
            // Not a JSON response, probably HTML update
        }
    }
});

// ═══════════════════════════════════════════════════════════════════════════
// NOTIFICATIONS
// ═══════════════════════════════════════════════════════════════════════════

function showNotification(message, type = 'info') {
    // Create notification element
    const notification = document.createElement('div');
    let backgroundColor = type === 'success' ? 'var(--semantic-feedback-success-default)' : 
                         type === 'error' ? 'var(--semantic-feedback-error-default)' : 
                         'var(--semantic-interactive-primary-default)';
    notification.style.cssText = 
        'position: fixed;' +
        'top: 20px;' +
        'right: 20px;' +
        'background: ' + backgroundColor + ';' +
        'color: var(--semantic-text-inverted);' +
        'padding: var(--spacing-md);' +
        'border-radius: var(--border-radius-md);' +
        'font-weight: var(--font-weight-medium);' +
        'z-index: 1000;' +
        'opacity: 0;' +
        'transform: translateY(-10px);' +
        'transition: all var(--animation-duration-normal) var(--animation-easing);';
    notification.textContent = message;
    
    document.body.appendChild(notification);
    
    // Animate in
    setTimeout(() => {
        notification.style.opacity = '1';
        notification.style.transform = 'translateY(0)';
    }, 10);
    
    // Remove after 3 seconds
    setTimeout(() => {
        notification.style.opacity = '0';
        notification.style.transform = 'translateY(-10px)';
        setTimeout(() => {
            if (notification.parentNode) {
                notification.parentNode.removeChild(notification);
            }
        }, 300);
    }, 3000);
}

// ═══════════════════════════════════════════════════════════════════════════
// KEYBOARD SHORTCUTS
// ═══════════════════════════════════════════════════════════════════════════

document.addEventListener('keydown', function(e) {
    if (e.ctrlKey || e.metaKey) {
        switch(e.key) {
            case 'r':
                e.preventDefault();
                resetForm();
                showNotification('🔄 Form reset', 'info');
                break;
            case 'Enter':
                e.preventDefault();
                htmx.ajax('POST', '/validate-all', {target:'#form-container'});
                break;
            case 't':
                e.preventDefault();
                toggleTheme();
                showNotification('🎨 Theme toggled', 'info');
                break;
        }
    }
});

// ═══════════════════════════════════════════════════════════════════════════
// INITIALIZATION
// ═══════════════════════════════════════════════════════════════════════════

document.addEventListener('DOMContentLoaded', function() {
    initializeTheme();
    
    // Display schema in the schema viewer
    const schemaDef = {
        "title": "User Registration Demo",
        "type": "object",
        "properties": {
            "username": {"type": "string", "minLength": 3, "maxLength": 20, "pattern": "^[a-zA-Z0-9_-]+$"},
            "email": {"type": "string", "format": "email"},
            "age": {"type": "number", "minimum": 18, "maximum": 120},
            "country": {"type": "string", "enum": ["US", "CA", "UK", "DE", "FR", "AU", "JP", "Other"]},
            "newsletter": {"type": "boolean"},
            "bio": {"type": "string", "maxLength": 500}
        },
        "required": ["username", "email", "age"]
    };
    
    const schemaDisplay = document.getElementById('schema-display');
    if (schemaDisplay) {
        schemaDisplay.textContent = JSON.stringify(schemaDef, null, 2);
    }
    
    console.log('🚀 Ruun Schema-Driven UI Demo initialized');
    console.log('📋 Available shortcuts:');
    console.log('  • Ctrl+R: Reset form');
    console.log('  • Ctrl+Enter: Validate all');
    console.log('  • Ctrl+T: Toggle theme');
});

// ═══════════════════════════════════════════════════════════════════════════
// REAL-TIME UPDATES
// ═══════════════════════════════════════════════════════════════════════════

// Auto-refresh state panels every few seconds
setInterval(function() {
    // Only refresh if elements exist and are visible
    const stateViewer = document.querySelector('[hx-get="/state"]');
    const validationViewer = document.querySelector('[hx-get="/validation-state"]');
    const tokenViewer = document.querySelector('[hx-get="/tokens"]');
    
    if (stateViewer && stateViewer.offsetParent !== null) {
        htmx.trigger(stateViewer, 'refresh');
    }
    if (validationViewer && validationViewer.offsetParent !== null) {
        htmx.trigger(validationViewer, 'refresh');
    }
    if (tokenViewer && tokenViewer.offsetParent !== null) {
        htmx.trigger(tokenViewer, 'refresh');
    }
}, 2000);
	`
}
