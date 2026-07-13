// Mobile menu toggle with accessibility
document.addEventListener('DOMContentLoaded', function () {
  const mobileMenuBtn = document.getElementById('mobile-menu-btn');
  const mobileMenu = document.getElementById('mobile-menu');

  if (mobileMenuBtn && mobileMenu) {
    mobileMenuBtn.addEventListener('click', function () {
      const isHidden = mobileMenu.classList.contains('hidden');
      mobileMenu.classList.toggle('hidden');
      mobileMenuBtn.setAttribute('aria-expanded', isHidden ? 'true' : 'false');
    });

    document.addEventListener('click', function (event) {
      if (!mobileMenu.contains(event.target) && !mobileMenuBtn.contains(event.target)) {
        mobileMenu.classList.add('hidden');
        mobileMenuBtn.setAttribute('aria-expanded', 'false');
      }
    });

    document.addEventListener('keydown', function (event) {
      if (event.key === 'Escape' && !mobileMenu.classList.contains('hidden')) {
        mobileMenu.classList.add('hidden');
        mobileMenuBtn.setAttribute('aria-expanded', 'false');
        mobileMenuBtn.focus();
      }
    });
  }
});

// HTMX loading states
document.body.addEventListener('htmx:beforeRequest', function (evt) {
  const target = evt.detail.target;
  if (target) {
    target.classList.add('opacity-50');
    target.setAttribute('aria-busy', 'true');
  }
});

document.body.addEventListener('htmx:afterRequest', function (evt) {
  const target = evt.detail.target;
  if (target) {
    target.classList.remove('opacity-50');
    target.setAttribute('aria-busy', 'false');
  }
});

document.body.addEventListener('htmx:responseError', function (evt) {
  console.error('HTMX request failed:', evt.detail.error);
  const target = evt.detail.target;
  if (target) {
    target.setAttribute('aria-busy', 'false');
  }
});

// After a form swap, move focus to the response for screen readers
document.body.addEventListener('htmx:afterSwap', function (evt) {
  const target = evt.detail.target;
  if (target && target.id === 'contact-form-container') {
    const response = target.querySelector('[role="alert"], .success-message, .error-message');
    if (response) {
      response.setAttribute('tabindex', '-1');
      response.focus();
    }
  }
});
