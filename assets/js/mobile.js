/* ── mobile.js ─────────────────────────────────────────────────────────────
   Toggles the mobile nav drawer open/closed.
   The hamburger button and drawer are rendered by nav.templ.
   ────────────────────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  initMobileMenu() {
    const btn  = document.getElementById('mobile-menu-btn');
    const menu = document.getElementById('mobile-menu');
    if (!btn || !menu) return;

    btn.addEventListener('click', () => menu.classList.toggle('hidden'));

    menu.querySelectorAll('.mobile-link').forEach(link => {
      link.addEventListener('click', () => menu.classList.add('hidden'));
    });
  },

});
