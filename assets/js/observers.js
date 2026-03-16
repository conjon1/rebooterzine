/* ── observers.js ──────────────────────────────────────────────────────────
   IntersectionObserver that adds .visible to .fade-in elements when they
   enter the viewport, triggering the CSS opacity/translate transition.
   Each element is unobserved after first reveal (fire-once).
   ────────────────────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  initObservers() {
    const observer = new IntersectionObserver(entries => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('visible');
          observer.unobserve(entry.target);
        }
      });
    }, { threshold: 0.1 });

    document.querySelectorAll('.fade-in').forEach(el => observer.observe(el));
  },

});
