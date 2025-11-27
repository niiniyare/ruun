/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './js/**/*.{js,ts}',
    './css/**/*.css', 
    '../views/**/*.templ',
    '../views/**/*.go',
  ],
  safelist: [
    // Include all utility classes used in basecoat.css
    'px-4', 'py-3', 'px-2','' 'py-0.5', 'px-3', 'py-2', 'px-1', 'py-1',
    'w-full', 'h-full', 'min-h-4', 'max-w-xs', 'max-w-sm', 'max-w-md', 'max-w-lg',
    'text-xs', 'text-sm', 'text-base', 'text-lg', 'text-xl', 'text-2xl', 'text-3xl',
    'font-medium', 'font-semibold', 'font-bold',
    'rounded-lg', 'rounded-md', 'rounded-sm', 'rounded-full',
    'border', 'border-transparent', 'border-2',
    'bg-card', 'bg-primary', 'bg-secondary', 'bg-destructive', 'bg-muted', 'bg-accent',
    'text-card-foreground', 'text-primary-foreground', 'text-secondary-foreground', 'text-destructive-foreground', 'text-muted-foreground', 'text-accent-foreground', 'text-foreground',
    'hover:bg-primary/90', 'hover:bg-secondary/90', 'hover:bg-destructive/90', 'hover:bg-accent', 'hover:text-accent-foreground',
    'focus-visible:outline-none', 'focus-visible:ring-2', 'focus-visible:ring-ring', 'focus-visible:ring-offset-2',
    'disabled:pointer-events-none', 'disabled:opacity-50',
    'transition-all', 'transition-colors',
    'inline-flex', 'flex', 'grid', 'block', 'inline-block', 'hidden',
    'items-center', 'items-start', 'items-end', 'justify-center', 'justify-start', 'justify-end', 'justify-between',
    'whitespace-nowrap', 'overflow-hidden', 'text-ellipsis',
    'relative', 'absolute', 'fixed', 'sticky',
    'top-0', 'right-0', 'bottom-0', 'left-0',
    'z-10', 'z-20', 'z-30', 'z-40', 'z-50',
    'gap-1', 'gap-2', 'gap-3', 'gap-4', 'gap-x-3', 'gap-y-0.5',
    'space-x-2', 'space-y-2', 'space-y-4',
    'size-3', 'size-4', 'size-5', 'size-6',
    'w-fit', 'shrink-0',
    'line-clamp-1', 'leading-relaxed',
    'list-inside', 'list-disc',
    'col-start-2', 'grid-cols-[calc(var(--spacing)*4)_1fr]', 'grid-cols-[0_1fr]',
    'has-[>svg]:grid-cols-[calc(var(--spacing)*4)_1fr]', 'has-[>svg]:gap-x-3',
    '[&>svg]:size-4', '[&>svg]:translate-y-0.5', '[&>svg]:text-current', '[&>svg]:pointer-events-none', '[&>svg]:shrink-0',
    '[&_svg]:pointer-events-none', '[&_svg:not([class*=\'size-\'])]:size-4',
    '[a&]:hover:bg-primary/90', '[a&]:hover:bg-secondary/90', '[a&]:hover:bg-destructive/90',
    'tracking-tight',
    'justify-items-start',
    '[&_p]:leading-relaxed',
    'text-white',
    'outline-none',
    'focus-visible:border-ring', 'focus-visible:ring-ring/50', 'focus-visible:ring-[3px]',
    'aria-invalid:ring-destructive/20', 'aria-invalid:border-destructive',
    'dark:aria-invalid:ring-destructive/40', 'dark:focus-visible:ring-destructive/40', 'dark:bg-destructive/60',
    'translate-y-0.5',
    'min-h-0',
    'text-current'
  ],
  theme: {
    extend: {
      colors: {
        // These will be populated by CSS custom properties from basecoat
        primary: 'oklch(var(--primary))',
        'primary-foreground': 'oklch(var(--primary-foreground))',
        secondary: 'oklch(var(--secondary))',
        'secondary-foreground': 'oklch(var(--secondary-foreground))',
        destructive: 'oklch(var(--destructive))',
        'destructive-foreground': 'white',
        muted: 'oklch(var(--muted))',
        'muted-foreground': 'oklch(var(--muted-foreground))',
        accent: 'oklch(var(--accent))',
        'accent-foreground': 'oklch(var(--accent-foreground))',
        popover: 'oklch(var(--popover))',
        'popover-foreground': 'oklch(var(--popover-foreground))',
        card: 'oklch(var(--card))',
        'card-foreground': 'oklch(var(--card-foreground))',
        border: 'oklch(var(--border))',
        input: 'oklch(var(--input))',
        ring: 'oklch(var(--ring))',
        background: 'oklch(var(--background))',
        foreground: 'oklch(var(--foreground))',
       destructive: {
          DEFAULT: '#dc2626', // or your destructive color
          foreground: '#ffffff' // or your foreground color
        }
      },
      borderRadius: {
        lg: 'var(--radius)',
        md: 'calc(var(--radius) - 2px)',
        sm: 'calc(var(--radius) - 4px)',
      },
      fontFamily: {
        sans: ['var(--font-sans)', 'system-ui', '-apple-system', 'sans-serif'],
      },
      spacing: {
        'spacing': 'var(--spacing, 1rem)',
      },
    },
  },
  plugins: [],
}
