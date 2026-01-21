// Mobile menu toggle with accessibility
document.addEventListener('DOMContentLoaded', function() {
  const mobileMenuBtn = document.getElementById('mobile-menu-btn');
  const mobileMenu = document.getElementById('mobile-menu');

  if (mobileMenuBtn && mobileMenu) {
    mobileMenuBtn.addEventListener('click', function() {
      const isHidden = mobileMenu.classList.contains('hidden');
      mobileMenu.classList.toggle('hidden');
      // Update aria-expanded for accessibility
      mobileMenuBtn.setAttribute('aria-expanded', isHidden ? 'true' : 'false');
    });
  }

  // Close mobile menu when clicking outside
  document.addEventListener('click', function(event) {
    if (mobileMenu && !mobileMenu.contains(event.target) && !mobileMenuBtn.contains(event.target)) {
      mobileMenu.classList.add('hidden');
      if (mobileMenuBtn) {
        mobileMenuBtn.setAttribute('aria-expanded', 'false');
      }
    }
  });

  // Close mobile menu on Escape key
  document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape' && mobileMenu && !mobileMenu.classList.contains('hidden')) {
      mobileMenu.classList.add('hidden');
      if (mobileMenuBtn) {
        mobileMenuBtn.setAttribute('aria-expanded', 'false');
        mobileMenuBtn.focus();
      }
    }
  });

  // Smooth scroll for anchor links
  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function(e) {
      const href = this.getAttribute('href');
      if (href !== '#') {
        e.preventDefault();
        const target = document.querySelector(href);
        if (target) {
          target.scrollIntoView({
            behavior: 'smooth',
            block: 'start'
          });
          // Set focus to target for accessibility
          target.setAttribute('tabindex', '-1');
          target.focus();
        }
      }
    });
  });

  // Header scroll effect
  const header = document.querySelector('header');
  const headerCta = document.getElementById('header-cta');
  const isHomePage = window.location.pathname === '/';

  if (header) {
    let lastScrollY = window.scrollY;

    // Hide CTA on home page initially
    if (isHomePage && headerCta) {
      headerCta.style.opacity = '0';
      headerCta.style.pointerEvents = 'none';
    }

    window.addEventListener('scroll', function() {
      const currentScrollY = window.scrollY;

      if (currentScrollY > 100) {
        header.classList.add('shadow-md');
      } else {
        header.classList.remove('shadow-md');
      }

      // Show/hide CTA button on home page based on scroll
      if (isHomePage && headerCta) {
        if (currentScrollY > 300) {
          headerCta.style.opacity = '1';
          headerCta.style.pointerEvents = 'auto';
        } else {
          headerCta.style.opacity = '0';
          headerCta.style.pointerEvents = 'none';
        }
      }

      lastScrollY = currentScrollY;
    });
  }

  // Intersection Observer for animations
  const observerOptions = {
    root: null,
    rootMargin: '0px',
    threshold: 0.1
  };

  const observer = new IntersectionObserver(function(entries) {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('animate-fade-in');
        observer.unobserve(entry.target);
      }
    });
  }, observerOptions);

  // Observe elements with data-animate attribute
  document.querySelectorAll('[data-animate]').forEach(el => {
    observer.observe(el);
  });
});

// HTMX event handlers
document.body.addEventListener('htmx:beforeRequest', function(evt) {
  // Add loading state to target element
  const target = evt.detail.target;
  if (target) {
    target.classList.add('opacity-50');
    target.setAttribute('aria-busy', 'true');
  }
});

document.body.addEventListener('htmx:afterRequest', function(evt) {
  // Remove loading state
  const target = evt.detail.target;
  if (target) {
    target.classList.remove('opacity-50');
    target.setAttribute('aria-busy', 'false');
  }
});

document.body.addEventListener('htmx:responseError', function(evt) {
  console.error('HTMX request failed:', evt.detail.error);
  // Announce error to screen readers
  const target = evt.detail.target;
  if (target) {
    target.setAttribute('aria-busy', 'false');
  }
});

// Handle form success - move focus to response for accessibility
document.body.addEventListener('htmx:afterSwap', function(evt) {
  const target = evt.detail.target;
  if (target && target.id === 'contact-form-container') {
    // Focus on the response for screen reader users
    const response = target.querySelector('[role="alert"], .success-message, .error-message');
    if (response) {
      response.setAttribute('tabindex', '-1');
      response.focus();
    }
  }
});
