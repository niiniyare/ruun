# Pagination Component

Pagination with page navigation, next and previous links.

## Important Note

**There is no dedicated pagination component in Basecoat.** Instead, use the pattern shown below with Basecoat button classes and semantic HTML.

## Basic Usage

```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="#" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li>
      <a href="#" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <a href="#" class="btn-icon-outline">2</a>
    </li>
    <li>
      <a href="#" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="#" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

## CSS Classes

### Primary Classes
Uses existing Basecoat button classes:
- **`btn-ghost`** - For Previous/Next links
- **`btn-icon-ghost`** - For inactive page numbers
- **`btn-icon-outline`** - For current/active page

### Tailwind Utilities Used
- `mx-auto` - Center the pagination horizontally
- `flex w-full justify-center` - Flexible centering layout
- `flex flex-row items-center` - Horizontal list layout
- `gap-1` - Small spacing between pagination items
- `size-9` - Size for ellipsis container
- `size-4 shrink-0` - Icon sizing

## HTML Structure

```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <!-- Previous link -->
    <li>
      <a href="/prev-page" class="btn-ghost">Previous</a>
    </li>
    
    <!-- Page numbers -->
    <li>
      <a href="/page-1" class="btn-icon-ghost">1</a>
    </li>
    
    <!-- Current page (highlighted) -->
    <li>
      <span class="btn-icon-outline" aria-current="page">2</span>
    </li>
    
    <!-- More page numbers -->
    <li>
      <a href="/page-3" class="btn-icon-ghost">3</a>
    </li>
    
    <!-- Ellipsis for truncated pages -->
    <li>
      <div class="size-9 flex items-center justify-center">
        <!-- Ellipsis icon -->
      </div>
    </li>
    
    <!-- Next link -->
    <li>
      <a href="/next-page" class="btn-ghost">Next</a>
    </li>
  </ul>
</nav>
```

## Examples

### Simple Pagination
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/1" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">2</span>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="/page/3" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Pagination with Ellipsis
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/4" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <div class="size-9 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4 shrink-0">
          <circle cx="12" cy="12" r="1" />
          <circle cx="19" cy="12" r="1" />
          <circle cx="5" cy="12" r="1" />
        </svg>
      </div>
    </li>
    <li>
      <a href="/page/4" class="btn-icon-ghost">4</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">5</span>
    </li>
    <li>
      <a href="/page/6" class="btn-icon-ghost">6</a>
    </li>
    <li>
      <div class="size-9 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4 shrink-0">
          <circle cx="12" cy="12" r="1" />
          <circle cx="19" cy="12" r="1" />
          <circle cx="5" cy="12" r="1" />
        </svg>
      </div>
    </li>
    <li>
      <a href="/page/20" class="btn-icon-ghost">20</a>
    </li>
    <li>
      <a href="/page/6" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Disabled States
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <!-- Disabled Previous (first page) -->
    <li>
      <span class="btn-ghost opacity-50 cursor-not-allowed">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </span>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">1</span>
    </li>
    <li>
      <a href="/page/2" class="btn-icon-ghost">2</a>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="/page/2" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Mobile-Friendly Pagination
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <!-- Mobile: Show only Previous/Next on small screens -->
    <li class="sm:hidden">
      <a href="/page/4" class="btn-ghost">Previous</a>
    </li>
    <li class="sm:hidden">
      <span class="btn-outline px-4">Page 5 of 20</span>
    </li>
    <li class="sm:hidden">
      <a href="/page/6" class="btn-ghost">Next</a>
    </li>
    
    <!-- Desktop: Full pagination -->
    <li class="hidden sm:inline">
      <a href="/page/4" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li class="hidden sm:inline">
      <div class="size-9 flex items-center justify-center">…</div>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/4" class="btn-icon-ghost">4</a>
    </li>
    <li class="hidden sm:inline">
      <span class="btn-icon-outline" aria-current="page">5</span>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/6" class="btn-icon-ghost">6</a>
    </li>
    <li class="hidden sm:inline">
      <div class="size-9 flex items-center justify-center">…</div>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/20" class="btn-icon-ghost">20</a>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/6" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Compact Pagination
```html
<nav role="navigation" aria-label="pagination" class="flex justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/1" class="btn-icon-ghost" title="Previous page">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">2</span>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost" title="Next page">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

## Icon SVGs

### Previous Icon
```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <path d="m15 18-6-6 6-6" />
</svg>
```

### Next Icon
```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <path d="m9 18 6-6-6-6" />
</svg>
```

### Ellipsis Icon
```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4 shrink-0">
  <circle cx="12" cy="12" r="1" />
  <circle cx="19" cy="12" r="1" />
  <circle cx="5" cy="12" r="1" />
</svg>
```

## Accessibility Features

- **Semantic HTML**: Uses `<nav>` and `<ul>` for proper navigation structure
- **Navigation Landmark**: `role="navigation"` identifies the pagination
- **ARIA Labels**: `aria-label="pagination"` describes the navigation purpose
- **Current Page**: `aria-current="page"` marks the active page
- **Keyboard Navigation**: All links are keyboard accessible
- **Screen Reader Support**: Proper list and link semantics

### Enhanced Accessibility
```html
<nav role="navigation" aria-label="Pagination Navigation">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/1" class="btn-ghost" aria-label="Go to previous page">
        <svg aria-hidden="true"><!-- Previous icon --></svg>
        Previous
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost" aria-label="Go to page 1">1</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page" aria-label="Current page, page 2">2</span>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost" aria-label="Go to page 3">3</a>
    </li>
    <li>
      <a href="/page/3" class="btn-ghost" aria-label="Go to next page">
        Next
        <svg aria-hidden="true"><!-- Next icon --></svg>
      </a>
    </li>
  </ul>
</nav>
```

## JavaScript Integration

### Dynamic Pagination
```javascript
function generatePagination(currentPage, totalPages, baseUrl) {
  const paginationContainer = document.getElementById('pagination');
  const maxVisiblePages = 5;
  
  let html = '<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">';
  html += '<ul class="flex flex-row items-center gap-1">';
  
  // Previous button
  if (currentPage > 1) {
    html += `
      <li>
        <a href="${baseUrl}?page=${currentPage - 1}" class="btn-ghost">
          <svg><!-- Previous icon --></svg>
          Previous
        </a>
      </li>
    `;
  } else {
    html += `
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">
          <svg><!-- Previous icon --></svg>
          Previous
        </span>
      </li>
    `;
  }
  
  // Page numbers
  const startPage = Math.max(1, currentPage - Math.floor(maxVisiblePages / 2));
  const endPage = Math.min(totalPages, startPage + maxVisiblePages - 1);
  
  if (startPage > 1) {
    html += `<li><a href="${baseUrl}?page=1" class="btn-icon-ghost">1</a></li>`;
    if (startPage > 2) {
      html += '<li><div class="size-9 flex items-center justify-center">…</div></li>';
    }
  }
  
  for (let page = startPage; page <= endPage; page++) {
    if (page === currentPage) {
      html += `<li><span class="btn-icon-outline" aria-current="page">${page}</span></li>`;
    } else {
      html += `<li><a href="${baseUrl}?page=${page}" class="btn-icon-ghost">${page}</a></li>`;
    }
  }
  
  if (endPage < totalPages) {
    if (endPage < totalPages - 1) {
      html += '<li><div class="size-9 flex items-center justify-center">…</div></li>';
    }
    html += `<li><a href="${baseUrl}?page=${totalPages}" class="btn-icon-ghost">${totalPages}</a></li>`;
  }
  
  // Next button
  if (currentPage < totalPages) {
    html += `
      <li>
        <a href="${baseUrl}?page=${currentPage + 1}" class="btn-ghost">
          Next
          <svg><!-- Next icon --></svg>
        </a>
      </li>
    `;
  } else {
    html += `
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">
          Next
          <svg><!-- Next icon --></svg>
        </span>
      </li>
    `;
  }
  
  html += '</ul></nav>';
  paginationContainer.innerHTML = html;
}

// Usage
generatePagination(5, 20, '/products');
```

### HTMX Integration
```html
<nav id="pagination" hx-get="/api/pagination" hx-trigger="page-change">
  <!-- Pagination content loaded dynamically -->
</nav>

<script>
// Custom event to trigger pagination updates
function changePage(page) {
  document.dispatchEvent(new CustomEvent('page-change', { detail: { page } }));
}
</script>
```

### Keyboard Navigation Enhancement
```javascript
document.addEventListener('keydown', function(e) {
  const currentPage = parseInt(document.querySelector('[aria-current="page"]').textContent);
  const totalPages = document.querySelectorAll('.btn-icon-ghost, .btn-icon-outline').length;
  
  if (e.key === 'ArrowLeft' && currentPage > 1) {
    window.location.href = `?page=${currentPage - 1}`;
  } else if (e.key === 'ArrowRight' && currentPage < totalPages) {
    window.location.href = `?page=${currentPage + 1}`;
  }
});
```

## Best Practices

1. **Provide Context**: Always include total page count and current position
2. **Limit Visible Pages**: Show 5-7 page numbers to avoid overcrowding
3. **Use Ellipsis**: Indicate skipped pages with ellipsis
4. **Disable When Appropriate**: Disable Previous on first page, Next on last page
5. **Mobile Optimization**: Simplify pagination for mobile devices
6. **Clear Labels**: Use descriptive aria-labels for screen readers
7. **Consistent Styling**: Follow button design patterns
8. **Fast Navigation**: Include first/last page links for large datasets

## URL Structure

### Query Parameters
```
/products?page=5
/search?q=term&page=3
/articles?category=tech&page=2&limit=10
```

### RESTful Paths
```
/products/page/5
/search/term/page/3
/articles/tech/page/2
```

## Integration Examples

### PHP/Laravel
```php
// Controller
$products = Product::paginate(10);

// Blade template
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    @if ($products->onFirstPage())
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">Previous</span>
      </li>
    @else
      <li>
        <a href="{{ $products->previousPageUrl() }}" class="btn-ghost">Previous</a>
      </li>
    @endif
    
    @foreach ($elements as $element)
      @if (is_array($element))
        @foreach ($element as $page => $url)
          @if ($page == $products->currentPage())
            <li>
              <span class="btn-icon-outline" aria-current="page">{{ $page }}</span>
            </li>
          @else
            <li>
              <a href="{{ $url }}" class="btn-icon-ghost">{{ $page }}</a>
            </li>
          @endif
        @endforeach
      @endif
    @endforeach
    
    @if ($products->hasMorePages())
      <li>
        <a href="{{ $products->nextPageUrl() }}" class="btn-ghost">Next</a>
      </li>
    @else
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">Next</span>
      </li>
    @endif
  </ul>
</nav>
```

## Related Components

- [Button](./button.md) - For pagination links and controls
- [Breadcrumb](./breadcrumb.md) - For hierarchical navigation
- [Table](./table.md) - Often used with pagination
- [Navigation](./navigation.md) - For primary site navigation