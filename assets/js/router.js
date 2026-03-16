/* ── router.js ─────────────────────────────────────────────────────────────
   Hash-based SPA routing.
   Handles #home, #blog, #authors, #post/<id>, #author/<id>.
   ────────────────────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  initRouter() {
    window.addEventListener('hashchange', () => this.handleRoute());
    this.handleRoute();
    this.initSearch();
    this.initTypewriter();
    this.initObservers();

    document.getElementById('back-btn').addEventListener('click', () => {
      window.history.back();
      setTimeout(() => {
        if (!window.location.hash) window.location.hash = '#blog';
      }, 100);
    });
  },

  async handleRoute() {
    const hash = window.location.hash || '#home';
    if (hash.startsWith('#post/')) {
      await this.loadPost(hash.replace('#post/', ''));
      this.showPage('content-view');
    } else if (hash.startsWith('#author/')) {
      await this.loadAuthorPage(hash.replace('#author/', ''));
      this.showPage('content-view');
    } else {
      this.showPage(hash.replace('#', ''));
    }
    window.scrollTo(0, 0);
  },

  showPage(id) {
    this.pages.forEach(p => {
      if (p.id === id) {
        p.classList.remove('hidden');
        setTimeout(() => {
          p.querySelectorAll('.fade-in').forEach(el => el.classList.add('visible'));
        }, 50);
      } else {
        p.classList.add('hidden');
      }
    });
  },

});
