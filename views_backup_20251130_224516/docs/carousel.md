# Carousel Component

A carousel with motion and swipe built with CSS scroll snapping.

## Basic Usage

```html
<div class="slider">
  <!-- Navigation links -->
  <a href="#slide-1">1</a>
  <a href="#slide-2">2</a>
  <a href="#slide-3">3</a>
  
  <!-- Slides container -->
  <div class="slides">
    <div id="slide-1">Slide 1 Content</div>
    <div id="slide-2">Slide 2 Content</div>
    <div id="slide-3">Slide 3 Content</div>
  </div>
</div>
```

## CSS Classes

### Container Classes
- **`slider`** - Main carousel container with fixed width and overflow hidden
- **`slides`** - Scrollable container with flex layout and scroll snap

### Custom CSS Required
The carousel component requires custom CSS for scroll snapping behavior:

```css
.slider {
  width: 300px;
  text-align: center;
  overflow: hidden;
}

.slides {
  display: flex;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  scroll-behavior: smooth;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none; /* Firefox */
}

.slides::-webkit-scrollbar {
  display: none; /* Webkit browsers */
}

.slides > div {
  scroll-snap-align: start;
  flex-shrink: 0;
  width: 300px;
  height: 300px;
  margin-right: 50px;
  border-radius: 10px;
  background: #eee;
  transform-origin: center center;
  transform: scale(1);
  transition: transform 0.5s;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 100px;
}
```

### Tailwind Utility Alternative
Using Tailwind utilities for responsive carousel:

```css
/* Alternative Tailwind-based carousel */
.carousel-container {
  @apply w-full max-w-lg mx-auto overflow-hidden;
}

.carousel-slides {
  @apply flex overflow-x-auto snap-x snap-mandatory scroll-smooth;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.carousel-slides::-webkit-scrollbar {
  display: none;
}

.carousel-slide {
  @apply snap-start flex-shrink-0 w-full h-64 bg-muted rounded-lg mr-4 flex items-center justify-center text-2xl font-bold;
}
```

## Component Attributes

### Carousel Container
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "slider" for carousel styling | Yes |

### Slides Container  
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "slides" for scroll behavior | Yes |

### Individual Slides
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Unique identifier for navigation links | Yes |

### Navigation Links
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `href` | string | Points to slide ID (e.g., "#slide-1") | Yes |

## No JavaScript Required (Basic)
The carousel uses CSS scroll snapping and hash navigation, requiring no JavaScript for basic functionality.

## HTML Structure

```html
<!-- Basic carousel structure -->
<div class="slider">
  <!-- Navigation (optional) -->
  <nav class="carousel-nav">
    <a href="#slide-1">•</a>
    <a href="#slide-2">•</a>
    <a href="#slide-3">•</a>
  </nav>
  
  <!-- Slides container -->
  <div class="slides">
    <div id="slide-1" class="slide">
      <!-- Slide content -->
    </div>
    <div id="slide-2" class="slide">
      <!-- Slide content -->
    </div>
    <div id="slide-3" class="slide">
      <!-- Slide content -->
    </div>
  </div>
</div>
```

## Examples

### Image Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="image-1" class="slide">
      <img src="/images/photo1.jpg" alt="Beautiful landscape" class="w-full h-full object-cover rounded-lg">
    </div>
    <div id="image-2" class="slide">
      <img src="/images/photo2.jpg" alt="City skyline" class="w-full h-full object-cover rounded-lg">
    </div>
    <div id="image-3" class="slide">
      <img src="/images/photo3.jpg" alt="Ocean view" class="w-full h-full object-cover rounded-lg">
    </div>
  </div>
  
  <!-- Dot navigation -->
  <nav class="flex justify-center gap-2 mt-4">
    <a href="#image-1" class="w-3 h-3 rounded-full bg-muted hover:bg-primary transition-colors"></a>
    <a href="#image-2" class="w-3 h-3 rounded-full bg-muted hover:bg-primary transition-colors"></a>
    <a href="#image-3" class="w-3 h-3 rounded-full bg-muted hover:bg-primary transition-colors"></a>
  </nav>
</div>
```

### Product Showcase Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="product-1" class="slide bg-background border rounded-lg p-6">
      <div class="text-center">
        <img src="/products/laptop.jpg" alt="Premium Laptop" class="w-32 h-32 mx-auto mb-4 object-cover rounded">
        <h3 class="text-lg font-semibold mb-2">Premium Laptop</h3>
        <p class="text-muted-foreground mb-4">High-performance laptop for professionals</p>
        <span class="text-2xl font-bold text-primary">$1,299</span>
      </div>
    </div>
    
    <div id="product-2" class="slide bg-background border rounded-lg p-6">
      <div class="text-center">
        <img src="/products/phone.jpg" alt="Smartphone" class="w-32 h-32 mx-auto mb-4 object-cover rounded">
        <h3 class="text-lg font-semibold mb-2">Smartphone</h3>
        <p class="text-muted-foreground mb-4">Latest flagship with advanced camera</p>
        <span class="text-2xl font-bold text-primary">$899</span>
      </div>
    </div>
    
    <div id="product-3" class="slide bg-background border rounded-lg p-6">
      <div class="text-center">
        <img src="/products/tablet.jpg" alt="Tablet" class="w-32 h-32 mx-auto mb-4 object-cover rounded">
        <h3 class="text-lg font-semibold mb-2">Tablet</h3>
        <p class="text-muted-foreground mb-4">Versatile tablet for work and entertainment</p>
        <span class="text-2xl font-bold text-primary">$599</span>
      </div>
    </div>
  </div>
</div>
```

### Testimonials Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="testimonial-1" class="slide bg-muted rounded-lg p-8">
      <div class="text-center max-w-md mx-auto">
        <svg class="w-8 h-8 mx-auto mb-4 text-primary" fill="currentColor" viewBox="0 0 24 24">
          <path d="M14.17 18.45c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36zm-8 0c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36z"/>
        </svg>
        <blockquote class="text-lg italic mb-4">
          "This product has completely transformed our workflow. Highly recommended!"
        </blockquote>
        <div>
          <img src="/avatars/user1.jpg" alt="Sarah Johnson" class="w-12 h-12 rounded-full mx-auto mb-2">
          <cite class="font-semibold not-italic">Sarah Johnson</cite>
          <div class="text-sm text-muted-foreground">Product Manager</div>
        </div>
      </div>
    </div>
    
    <div id="testimonial-2" class="slide bg-muted rounded-lg p-8">
      <div class="text-center max-w-md mx-auto">
        <svg class="w-8 h-8 mx-auto mb-4 text-primary" fill="currentColor" viewBox="0 0 24 24">
          <path d="M14.17 18.45c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36zm-8 0c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36z"/>
        </svg>
        <blockquote class="text-lg italic mb-4">
          "Amazing user experience and excellent customer support. Five stars!"
        </blockquote>
        <div>
          <img src="/avatars/user2.jpg" alt="Michael Chen" class="w-12 h-12 rounded-full mx-auto mb-2">
          <cite class="font-semibold not-italic">Michael Chen</cite>
          <div class="text-sm text-muted-foreground">Software Developer</div>
        </div>
      </div>
    </div>
  </div>
  
  <!-- Arrow navigation -->
  <div class="flex justify-between items-center mt-4">
    <a href="#testimonial-1" class="btn-icon-outline">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m15 18-6-6 6-6"/>
      </svg>
    </a>
    <div class="flex gap-2">
      <a href="#testimonial-1" class="w-2 h-2 rounded-full bg-primary"></a>
      <a href="#testimonial-2" class="w-2 h-2 rounded-full bg-muted"></a>
    </div>
    <a href="#testimonial-2" class="btn-icon-outline">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m9 18 6-6-6-6"/>
      </svg>
    </a>
  </div>
</div>
```

### Card Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="feature-1" class="slide">
      <div class="card p-6 h-full">
        <div class="flex items-center gap-3 mb-4">
          <div class="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
              <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
            </svg>
          </div>
          <h3 class="text-lg font-semibold">Easy to Use</h3>
        </div>
        <p class="text-muted-foreground">
          Get started in minutes with our intuitive interface and comprehensive documentation.
        </p>
      </div>
    </div>
    
    <div id="feature-2" class="slide">
      <div class="card p-6 h-full">
        <div class="flex items-center gap-3 mb-4">
          <div class="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
              <path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/>
            </svg>
          </div>
          <h3 class="text-lg font-semibold">Lightning Fast</h3>
        </div>
        <p class="text-muted-foreground">
          Optimized performance ensures your application runs smoothly at any scale.
        </p>
      </div>
    </div>
    
    <div id="feature-3" class="slide">
      <div class="card p-6 h-full">
        <div class="flex items-center gap-3 mb-4">
          <div class="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
              <path d="M9 12l2 2 4-4"/>
              <path d="M21 12c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
              <path d="M3 12c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
              <path d="M12 21c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
              <path d="M12 3c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
            </svg>
          </div>
          <h3 class="text-lg font-semibold">Secure</h3>
        </div>
        <p class="text-muted-foreground">
          Built-in security features protect your data with enterprise-grade encryption.
        </p>
      </div>
    </div>
  </div>
</div>
```

### Responsive Multi-Item Carousel

```html
<style>
.multi-carousel {
  @apply w-full max-w-6xl mx-auto overflow-hidden;
}

.multi-slides {
  @apply flex overflow-x-auto snap-x snap-mandatory scroll-smooth gap-4 pb-4;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.multi-slides::-webkit-scrollbar {
  display: none;
}

.multi-slide {
  @apply snap-start flex-shrink-0 w-64 h-48 bg-muted rounded-lg p-4;
}

@media (min-width: 640px) {
  .multi-slide {
    @apply w-72;
  }
}

@media (min-width: 1024px) {
  .multi-slide {
    @apply w-80;
  }
}
</style>

<div class="multi-carousel">
  <div class="multi-slides">
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 1</h3>
      <p class="text-sm text-muted-foreground mb-4">Lorem ipsum dolor sit amet consectetur adipiscing elit...</p>
      <a href="/blog/post-1" class="text-primary text-sm hover:underline">Read more</a>
    </div>
    
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 2</h3>
      <p class="text-sm text-muted-foreground mb-4">Sed do eiusmod tempor incididunt ut labore et dolore...</p>
      <a href="/blog/post-2" class="text-primary text-sm hover:underline">Read more</a>
    </div>
    
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 3</h3>
      <p class="text-sm text-muted-foreground mb-4">Ut enim ad minim veniam quis nostrud exercitation...</p>
      <a href="/blog/post-3" class="text-primary text-sm hover:underline">Read more</a>
    </div>
    
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 4</h3>
      <p class="text-sm text-muted-foreground mb-4">Duis aute irure dolor in reprehenderit in voluptate...</p>
      <a href="/blog/post-4" class="text-primary text-sm hover:underline">Read more</a>
    </div>
  </div>
</div>
```

## Accessibility Features

- **Keyboard Navigation**: Use arrow keys or tab to navigate
- **Focus Management**: Proper focus states for navigation links
- **Screen Reader Support**: Use descriptive alt text and labels
- **Reduced Motion**: Respect user preferences for motion

### Enhanced Accessibility

```html
<div class="slider" role="region" aria-label="Image carousel" aria-live="polite">
  <div class="slides" role="list">
    <div id="slide-1" class="slide" role="listitem" aria-label="Slide 1 of 3">
      <img src="/images/photo1.jpg" alt="Beautiful mountain landscape with snow-capped peaks">
    </div>
    <div id="slide-2" class="slide" role="listitem" aria-label="Slide 2 of 3">
      <img src="/images/photo2.jpg" alt="Bustling city skyline at sunset">
    </div>
    <div id="slide-3" class="slide" role="listitem" aria-label="Slide 3 of 3">
      <img src="/images/photo3.jpg" alt="Peaceful ocean view with clear blue water">
    </div>
  </div>
  
  <nav aria-label="Carousel navigation" class="flex justify-center gap-2 mt-4">
    <a href="#slide-1" aria-label="Go to slide 1" class="carousel-nav-dot"></a>
    <a href="#slide-2" aria-label="Go to slide 2" class="carousel-nav-dot"></a>
    <a href="#slide-3" aria-label="Go to slide 3" class="carousel-nav-dot"></a>
  </nav>
</div>
```

## JavaScript Enhancement

### Auto-Play Carousel

```javascript
class AutoCarousel {
  constructor(container, options = {}) {
    this.container = container;
    this.slides = container.querySelectorAll('.slides > div');
    this.currentSlide = 0;
    this.interval = options.interval || 5000;
    this.autoPlayId = null;
    
    this.init();
  }
  
  init() {
    this.startAutoPlay();
    this.addEventListeners();
  }
  
  startAutoPlay() {
    this.autoPlayId = setInterval(() => {
      this.nextSlide();
    }, this.interval);
  }
  
  stopAutoPlay() {
    if (this.autoPlayId) {
      clearInterval(this.autoPlayId);
      this.autoPlayId = null;
    }
  }
  
  nextSlide() {
    this.currentSlide = (this.currentSlide + 1) % this.slides.length;
    this.goToSlide(this.currentSlide);
  }
  
  goToSlide(index) {
    const slide = this.slides[index];
    if (slide) {
      slide.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
      this.currentSlide = index;
    }
  }
  
  addEventListeners() {
    // Pause on hover
    this.container.addEventListener('mouseenter', () => this.stopAutoPlay());
    this.container.addEventListener('mouseleave', () => this.startAutoPlay());
    
    // Pause on focus
    this.container.addEventListener('focusin', () => this.stopAutoPlay());
    this.container.addEventListener('focusout', () => this.startAutoPlay());
  }
}

// Usage
const carousel = new AutoCarousel(document.querySelector('.slider'), {
  interval: 3000
});
```

### Touch/Swipe Support

```javascript
class SwipeCarousel {
  constructor(container) {
    this.container = container;
    this.slides = container.querySelector('.slides');
    this.startX = 0;
    this.currentX = 0;
    this.isDragging = false;
    
    this.init();
  }
  
  init() {
    this.slides.addEventListener('touchstart', (e) => this.handleStart(e));
    this.slides.addEventListener('touchmove', (e) => this.handleMove(e));
    this.slides.addEventListener('touchend', () => this.handleEnd());
    
    // Mouse events for desktop
    this.slides.addEventListener('mousedown', (e) => this.handleStart(e));
    this.slides.addEventListener('mousemove', (e) => this.handleMove(e));
    this.slides.addEventListener('mouseup', () => this.handleEnd());
    this.slides.addEventListener('mouseleave', () => this.handleEnd());
  }
  
  handleStart(e) {
    this.isDragging = true;
    this.startX = e.type === 'touchstart' ? e.touches[0].clientX : e.clientX;
    this.slides.style.scrollBehavior = 'auto';
  }
  
  handleMove(e) {
    if (!this.isDragging) return;
    
    e.preventDefault();
    this.currentX = e.type === 'touchmove' ? e.touches[0].clientX : e.clientX;
    const diffX = this.startX - this.currentX;
    this.slides.scrollLeft += diffX;
    this.startX = this.currentX;
  }
  
  handleEnd() {
    this.isDragging = false;
    this.slides.style.scrollBehavior = 'smooth';
  }
}

// Usage
new SwipeCarousel(document.querySelector('.slider'));
```

### React Carousel Component

```jsx
import React, { useState, useEffect, useRef } from 'react';

function Carousel({ children, autoPlay = false, interval = 5000 }) {
  const [currentSlide, setCurrentSlide] = useState(0);
  const slidesRef = useRef(null);
  const autoPlayRef = useRef(null);
  
  const totalSlides = React.Children.count(children);
  
  const goToSlide = (index) => {
    setCurrentSlide(index);
    const slide = slidesRef.current?.children[index];
    if (slide) {
      slide.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
    }
  };
  
  const nextSlide = () => {
    goToSlide((currentSlide + 1) % totalSlides);
  };
  
  const prevSlide = () => {
    goToSlide(currentSlide === 0 ? totalSlides - 1 : currentSlide - 1);
  };
  
  useEffect(() => {
    if (autoPlay) {
      autoPlayRef.current = setInterval(nextSlide, interval);
      return () => clearInterval(autoPlayRef.current);
    }
  }, [autoPlay, interval, currentSlide]);
  
  const pauseAutoPlay = () => {
    if (autoPlayRef.current) {
      clearInterval(autoPlayRef.current);
    }
  };
  
  const resumeAutoPlay = () => {
    if (autoPlay) {
      autoPlayRef.current = setInterval(nextSlide, interval);
    }
  };
  
  return (
    <div 
      className="slider"
      onMouseEnter={pauseAutoPlay}
      onMouseLeave={resumeAutoPlay}
    >
      <div 
        ref={slidesRef}
        className="slides"
      >
        {React.Children.map(children, (child, index) => (
          <div key={index} id={`slide-${index}`} className="slide">
            {child}
          </div>
        ))}
      </div>
      
      {/* Navigation dots */}
      <div className="flex justify-center gap-2 mt-4">
        {Array.from({ length: totalSlides }, (_, index) => (
          <button
            key={index}
            onClick={() => goToSlide(index)}
            className={`w-3 h-3 rounded-full transition-colors ${
              index === currentSlide ? 'bg-primary' : 'bg-muted hover:bg-primary/50'
            }`}
            aria-label={`Go to slide ${index + 1}`}
          />
        ))}
      </div>
      
      {/* Arrow navigation */}
      <div className="flex justify-between items-center mt-4">
        <button
          onClick={prevSlide}
          className="btn-icon-outline"
          aria-label="Previous slide"
        >
          ←
        </button>
        <button
          onClick={nextSlide}
          className="btn-icon-outline"
          aria-label="Next slide"
        >
          →
        </button>
      </div>
    </div>
  );
}

// Usage
function App() {
  return (
    <Carousel autoPlay interval={4000}>
      <div>Slide 1 Content</div>
      <div>Slide 2 Content</div>
      <div>Slide 3 Content</div>
    </Carousel>
  );
}
```

## Best Practices

1. **Performance**: Use CSS scroll-snap for smooth native scrolling
2. **Accessibility**: Provide keyboard navigation and screen reader support
3. **Responsive Design**: Adapt slide sizes for different screen sizes
4. **Auto-play**: Include pause/play controls for auto-playing carousels
5. **Touch Support**: Enable swipe gestures on touch devices
6. **Loading States**: Show placeholders for content that's loading
7. **Navigation**: Provide multiple ways to navigate (dots, arrows, keyboard)

## Common Patterns

### Hero Carousel

```html
<div class="slider w-full h-96">
  <div class="slides">
    <div id="hero-1" class="slide bg-gradient-to-r from-blue-500 to-purple-600 text-white">
      <div class="flex items-center justify-center h-full">
        <div class="text-center">
          <h1 class="text-4xl font-bold mb-4">Welcome to Our Platform</h1>
          <p class="text-xl mb-6">Discover amazing features and capabilities</p>
          <button class="btn bg-white text-blue-600 hover:bg-gray-100">Get Started</button>
        </div>
      </div>
    </div>
  </div>
</div>
```

### Media Carousel with Thumbnails

```html
<div class="slider">
  <!-- Main carousel -->
  <div class="slides mb-4">
    <div id="main-1" class="slide">
      <img src="/media/image1.jpg" alt="Product image 1" class="w-full h-full object-cover">
    </div>
    <div id="main-2" class="slide">
      <img src="/media/image2.jpg" alt="Product image 2" class="w-full h-full object-cover">
    </div>
  </div>
  
  <!-- Thumbnail navigation -->
  <div class="flex gap-2 justify-center">
    <a href="#main-1" class="w-16 h-16 rounded border-2 border-transparent hover:border-primary">
      <img src="/media/thumb1.jpg" alt="Thumbnail 1" class="w-full h-full object-cover rounded">
    </a>
    <a href="#main-2" class="w-16 h-16 rounded border-2 border-transparent hover:border-primary">
      <img src="/media/thumb2.jpg" alt="Thumbnail 2" class="w-full h-full object-cover rounded">
    </a>
  </div>
</div>
```

## Related Components

- [Card](./card.md) - For carousel slide content
- [Button](./button.md) - For navigation controls
- [Image](./avatar.md) - For image carousels
- [Badge](./badge.md) - For slide indicators