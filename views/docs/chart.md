# Chart Component

Interactive charts and data visualizations built with Chart.js and integrated with Basecoat theming.

## Basic Usage

```html
<div class="chart-container">
  <canvas id="myChart" width="400" height="200"></canvas>
</div>

<script>
  const ctx = document.getElementById('myChart').getContext('2d');
  new Chart(ctx, {
    type: 'line',
    data: {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May'],
      datasets: [{
        label: 'Sales',
        data: [12, 19, 3, 5, 2],
        borderColor: 'hsl(var(--chart-1))',
        backgroundColor: 'hsla(var(--chart-1), 0.1)',
        tension: 0.3
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false
    }
  });
</script>
```

## CSS Classes

### Container Classes
- **`chart-container`** - Wrapper for chart canvas with proper sizing
- **`card`** - Often wrapped in card component for consistent styling

### Chart Variables
Basecoat provides CSS custom properties for consistent chart colors:
- **`--chart-1`** - Primary chart color (blue)
- **`--chart-2`** - Secondary chart color (emerald)  
- **`--chart-3`** - Tertiary chart color (yellow)
- **`--chart-4`** - Quaternary chart color (red)
- **`--chart-5`** - Quinary chart color (purple)

### Theme Integration
Charts automatically adapt to light/dark mode using CSS variables:
- **`--background`** - Chart background color
- **`--border`** - Grid line and border color
- **`--foreground`** - Text and label color
- **`--muted-foreground`** - Secondary text color

## Component Attributes

### Canvas Element
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Unique identifier for Chart.js initialization | Yes |
| `width` | number | Canvas width in pixels | Optional |
| `height` | number | Canvas height in pixels | Optional |

## JavaScript Required

This component requires Chart.js library:

```html
<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
```

## HTML Structure

```html
<!-- Basic chart structure -->
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Chart Title</h3>
    <p class="text-sm text-muted-foreground">Chart description</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="chart-id"></canvas>
  </div>
</div>
```

## Examples

### Area Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Sales Overview</h3>
    <p class="text-sm text-muted-foreground">Monthly sales data</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="areaChart"></canvas>
  </div>
</div>

<script>
const areaCtx = document.getElementById('areaChart').getContext('2d');
new Chart(areaCtx, {
  type: 'line',
  data: {
    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
    datasets: [{
      label: 'Sales',
      data: [30, 40, 45, 50, 49, 60],
      borderColor: 'hsl(var(--chart-1))',
      backgroundColor: createGradient(areaCtx, 'var(--chart-1)'),
      fill: true,
      tension: 0.4
    }]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        display: false
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
});

function createGradient(ctx, color) {
  const gradient = ctx.createLinearGradient(0, 0, 0, 200);
  gradient.addColorStop(0, `hsla(${color}, 0.3)`);
  gradient.addColorStop(1, `hsla(${color}, 0)`);
  return gradient;
}
</script>
```

### Bar Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Revenue by Product</h3>
    <p class="text-sm text-muted-foreground">Quarterly comparison</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="barChart"></canvas>
  </div>
</div>

<script>
const barCtx = document.getElementById('barChart').getContext('2d');
new Chart(barCtx, {
  type: 'bar',
  data: {
    labels: ['Q1', 'Q2', 'Q3', 'Q4'],
    datasets: [
      {
        label: 'Product A',
        data: [65, 59, 80, 81],
        backgroundColor: 'hsl(var(--chart-1))',
        borderColor: 'hsl(var(--chart-1))',
        borderWidth: 1
      },
      {
        label: 'Product B',
        data: [28, 48, 40, 19],
        backgroundColor: 'hsl(var(--chart-2))',
        borderColor: 'hsl(var(--chart-2))',
        borderWidth: 1
      }
    ]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
});
</script>
```

### Doughnut Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Market Share</h3>
    <p class="text-sm text-muted-foreground">By company segment</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="doughnutChart"></canvas>
  </div>
</div>

<script>
const doughnutCtx = document.getElementById('doughnutChart').getContext('2d');
new Chart(doughnutCtx, {
  type: 'doughnut',
  data: {
    labels: ['Desktop', 'Mobile', 'Tablet'],
    datasets: [{
      data: [300, 50, 100],
      backgroundColor: [
        'hsl(var(--chart-1))',
        'hsl(var(--chart-2))',
        'hsl(var(--chart-3))'
      ],
      borderColor: [
        'hsl(var(--chart-1))',
        'hsl(var(--chart-2))',
        'hsl(var(--chart-3))'
      ],
      borderWidth: 2
    }]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    cutout: '60%'
  }
});
</script>
```

### Line Chart with Multiple Datasets

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Website Analytics</h3>
    <p class="text-sm text-muted-foreground">Visitors and page views</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="lineChart"></canvas>
  </div>
</div>

<script>
const lineCtx = document.getElementById('lineChart').getContext('2d');
new Chart(lineCtx, {
  type: 'line',
  data: {
    labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
    datasets: [
      {
        label: 'Visitors',
        data: [1200, 1900, 3000, 5000, 2000, 3000, 4500],
        borderColor: 'hsl(var(--chart-1))',
        backgroundColor: 'hsla(var(--chart-1), 0.1)',
        tension: 0.3
      },
      {
        label: 'Page Views',
        data: [2400, 3800, 6000, 10000, 4000, 6000, 9000],
        borderColor: 'hsl(var(--chart-2))',
        backgroundColor: 'hsla(var(--chart-2), 0.1)',
        tension: 0.3
      }
    ]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    interaction: {
      intersect: false,
      mode: 'index'
    },
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
});
</script>
```

### Stacked Bar Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Revenue Breakdown</h3>
    <p class="text-sm text-muted-foreground">By region and quarter</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="stackedChart"></canvas>
  </div>
</div>

<script>
const stackedCtx = document.getElementById('stackedChart').getContext('2d');
new Chart(stackedCtx, {
  type: 'bar',
  data: {
    labels: ['Q1', 'Q2', 'Q3', 'Q4'],
    datasets: [
      {
        label: 'North America',
        data: [120, 150, 180, 200],
        backgroundColor: 'hsl(var(--chart-1))',
      },
      {
        label: 'Europe',
        data: [80, 90, 100, 110],
        backgroundColor: 'hsl(var(--chart-2))',
      },
      {
        label: 'Asia Pacific',
        data: [60, 70, 85, 95],
        backgroundColor: 'hsl(var(--chart-3))',
      }
    ]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    scales: {
      x: {
        stacked: true
      },
      y: {
        stacked: true,
        beginAtZero: true
      }
    }
  }
});
</script>
```

## Custom Tooltip Implementation

```javascript
// Custom tooltip for better styling integration
const customTooltip = (context) => {
  // Get or create tooltip element
  let tooltipEl = document.getElementById('chartjs-tooltip');
  
  if (!tooltipEl) {
    tooltipEl = document.createElement('div');
    tooltipEl.id = 'chartjs-tooltip';
    tooltipEl.className = 'absolute bg-background border border-border rounded-lg shadow-lg p-3 text-sm pointer-events-none z-50 opacity-0 transition-opacity';
    document.body.appendChild(tooltipEl);
  }

  // Hide if no tooltip
  const tooltipModel = context.tooltip;
  if (tooltipModel.opacity === 0) {
    tooltipEl.style.opacity = 0;
    return;
  }

  // Set content
  if (tooltipModel.body) {
    const titleLines = tooltipModel.title || [];
    const bodyLines = tooltipModel.body.map(item => item.lines);

    let innerHtml = '<div class="font-medium mb-1">';
    titleLines.forEach(title => {
      innerHtml += title;
    });
    innerHtml += '</div>';

    bodyLines.forEach((body, i) => {
      const colors = tooltipModel.labelColors[i];
      innerHtml += `
        <div class="flex items-center gap-2">
          <div class="w-3 h-3 rounded-full" style="background-color: ${colors.backgroundColor}"></div>
          <span>${body}</span>
        </div>
      `;
    });

    tooltipEl.innerHTML = innerHtml;
  }

  // Position
  const position = context.chart.canvas.getBoundingClientRect();
  tooltipEl.style.opacity = 1;
  tooltipEl.style.left = position.left + window.pageXOffset + tooltipModel.caretX + 'px';
  tooltipEl.style.top = position.top + window.pageYOffset + tooltipModel.caretY + 'px';
};

// Usage in chart configuration
const chartOptions = {
  plugins: {
    tooltip: {
      enabled: false,
      external: customTooltip
    }
  }
};
```

## Helper Functions

```javascript
// Color utilities for charts
const chartHelpers = {
  // Convert chart color variables to usable format
  getChartColor(variable) {
    return `hsl(${getComputedStyle(document.documentElement).getPropertyValue(variable)})`;
  },

  // Add alpha channel to chart colors
  setAlpha(color, alpha) {
    const hslMatch = color.match(/hsl\((.*?)\)/);
    if (hslMatch) {
      return `hsla(${hslMatch[1]}, ${alpha})`;
    }
    return color;
  },

  // Create gradient for area charts
  createGradient(ctx, colorVar, height = 200) {
    const gradient = ctx.createLinearGradient(0, 0, 0, height);
    const color = this.getChartColor(colorVar);
    
    gradient.addColorStop(0, this.setAlpha(color, 0.3));
    gradient.addColorStop(1, this.setAlpha(color, 0));
    
    return gradient;
  },

  // Format dates for chart labels
  formatDateLabel(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', { 
      month: 'short', 
      day: 'numeric' 
    });
  },

  // Generate chart data from API response
  transformData(apiData, labelKey, valueKey) {
    return {
      labels: apiData.map(item => this.formatDateLabel(item[labelKey])),
      datasets: [{
        data: apiData.map(item => item[valueKey]),
        borderColor: this.getChartColor('--chart-1'),
        backgroundColor: this.createGradient(null, '--chart-1')
      }]
    };
  }
};
```

## Responsive Chart Component

```javascript
class ResponsiveChart {
  constructor(canvasId, config) {
    this.canvas = document.getElementById(canvasId);
    this.ctx = this.canvas.getContext('2d');
    this.config = config;
    this.chart = null;
    
    this.init();
    this.setupResize();
  }
  
  init() {
    this.chart = new Chart(this.ctx, {
      ...this.config,
      options: {
        responsive: true,
        maintainAspectRatio: false,
        ...this.config.options
      }
    });
  }
  
  setupResize() {
    window.addEventListener('resize', () => {
      if (this.chart) {
        this.chart.resize();
      }
    });
  }
  
  updateData(newData) {
    this.chart.data = newData;
    this.chart.update();
  }
  
  destroy() {
    if (this.chart) {
      this.chart.destroy();
    }
  }
}

// Usage
const salesChart = new ResponsiveChart('salesChart', {
  type: 'line',
  data: {
    labels: ['Jan', 'Feb', 'Mar'],
    datasets: [{
      label: 'Sales',
      data: [100, 200, 150]
    }]
  }
});
```

## React Chart Component

```jsx
import React, { useEffect, useRef } from 'react';
import Chart from 'chart.js/auto';

function ChartComponent({ 
  type = 'line', 
  data, 
  options = {}, 
  className = '',
  height = 256 
}) {
  const canvasRef = useRef(null);
  const chartRef = useRef(null);
  
  useEffect(() => {
    if (canvasRef.current) {
      // Destroy existing chart
      if (chartRef.current) {
        chartRef.current.destroy();
      }
      
      // Create new chart
      chartRef.current = new Chart(canvasRef.current, {
        type,
        data,
        options: {
          responsive: true,
          maintainAspectRatio: false,
          plugins: {
            tooltip: {
              backgroundColor: 'hsl(var(--background))',
              titleColor: 'hsl(var(--foreground))',
              bodyColor: 'hsl(var(--foreground))',
              borderColor: 'hsl(var(--border))',
              borderWidth: 1
            }
          },
          ...options
        }
      });
    }
    
    return () => {
      if (chartRef.current) {
        chartRef.current.destroy();
      }
    };
  }, [type, data, options]);
  
  return (
    <div className={`chart-container ${className}`} style={{ height }}>
      <canvas ref={canvasRef} />
    </div>
  );
}

// Usage
function Dashboard() {
  const chartData = {
    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May'],
    datasets: [{
      label: 'Revenue',
      data: [12, 19, 3, 5, 2],
      borderColor: 'hsl(var(--chart-1))',
      backgroundColor: 'hsla(var(--chart-1), 0.1)',
    }]
  };
  
  return (
    <div className="card p-6">
      <header className="mb-6">
        <h3 className="text-lg font-semibold">Sales Dashboard</h3>
        <p className="text-sm text-muted-foreground">Monthly revenue tracking</p>
      </header>
      
      <ChartComponent
        type="line"
        data={chartData}
        height={300}
        options={{
          scales: {
            y: {
              beginAtZero: true
            }
          }
        }}
      />
    </div>
  );
}
```

## Accessibility Features

- **Keyboard Navigation**: Charts support keyboard interaction when focused
- **Screen Reader Support**: Include descriptive labels and summaries
- **High Contrast**: Colors are accessible in both light and dark modes
- **Alternative Text**: Provide text alternatives for chart data

### Enhanced Accessibility

```html
<div class="card p-6">
  <header className="mb-6">
    <h3 className="text-lg font-semibold" id="sales-chart-title">
      Sales Performance
    </h3>
    <p className="text-sm text-muted-foreground">
      Monthly sales data from January to June 2024
    </p>
  </header>
  
  <div className="chart-container h-64">
    <canvas 
      id="salesChart"
      role="img"
      aria-labelledby="sales-chart-title"
      aria-describedby="sales-chart-summary"
    ></canvas>
  </div>
  
  <div id="sales-chart-summary" className="sr-only">
    Chart showing sales performance over 6 months. 
    January: $30k, February: $40k, March: $45k, 
    April: $50k, May: $49k, June: $60k. 
    Overall trend is positive with steady growth.
  </div>
  
  <!-- Data table alternative -->
  <details className="mt-4">
    <summary className="text-sm text-muted-foreground cursor-pointer">
      View data table
    </summary>
    <table className="table mt-2">
      <thead>
        <tr>
          <th>Month</th>
          <th>Sales</th>
        </tr>
      </thead>
      <tbody>
        <tr><td>January</td><td>$30,000</td></tr>
        <tr><td>February</td><td>$40,000</td></tr>
        <tr><td>March</td><td>$45,000</td></tr>
        <tr><td>April</td><td>$50,000</td></tr>
        <tr><td>May</td><td>$49,000</td></tr>
        <tr><td>June</td><td>$60,000</td></tr>
      </tbody>
    </table>
  </details>
</div>
```

## Best Practices

1. **Performance**: Use Chart.js animations sparingly on mobile devices
2. **Accessibility**: Always provide data tables as alternatives
3. **Responsive Design**: Test charts on different screen sizes
4. **Color Choice**: Use the provided chart color variables
5. **Loading States**: Show skeleton loaders while data loads
6. **Error Handling**: Display fallback content when charts fail to load
7. **Data Updates**: Use Chart.js update methods for smooth transitions

## Common Patterns

### Dashboard Grid

```html
<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
  <!-- KPI Cards -->
  <div className="card p-6">
    <div className="flex items-center justify-between">
      <div>
        <p className="text-sm text-muted-foreground">Total Revenue</p>
        <p className="text-2xl font-bold">$45,231</p>
      </div>
      <div className="chart-container h-16 w-16">
        <canvas id="revenueSparkline"></canvas>
      </div>
    </div>
  </div>
  
  <!-- Main Chart -->
  <div className="card p-6 md:col-span-2">
    <div className="chart-container h-64">
      <canvas id="mainChart"></canvas>
    </div>
  </div>
</div>
```

### Chart with Controls

```html
<div className="card p-6">
  <header className="flex items-center justify-between mb-6">
    <div>
      <h3 className="text-lg font-semibold">Analytics</h3>
      <p className="text-sm text-muted-foreground">User engagement metrics</p>
    </div>
    <div className="flex gap-2">
      <select className="select" id="timeRange">
        <option value="7d">Last 7 days</option>
        <option value="30d">Last 30 days</option>
        <option value="90d">Last 90 days</option>
      </select>
    </div>
  </header>
  
  <div className="chart-container h-64">
    <canvas id="analyticsChart"></canvas>
  </div>
</div>
```

## Related Components

- [Card](./card.md) - For chart containers
- [Select](./select.md) - For chart controls
- [Badge](./badge.md) - For chart legends
- [Table](./table.md) - For data table alternatives