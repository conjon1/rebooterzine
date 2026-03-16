/* ── app.js ────────────────────────────────────────────────────────────────
   BlogApp class skeleton.
   Methods are added to the prototype by the files that follow in concat order:
     router.js → loader.js → search.js → typewriter.js → observers.js → mobile.js
   boot.js fires DOMContentLoaded last.
   ────────────────────────────────────────────────────────────────────────── */
class BlogApp {
  constructor() {
    this.pages       = document.querySelectorAll('.page');
    this.contentDiv  = document.getElementById('dynamic-content');
    this.authorsData = [];
    this.init();
  }

  async init() {
    await this.fetchAuthors();
    this.initRouter();
    this.initMobileMenu();
  }

  async fetchAuthors() {
    try {
      const res = await fetch('authors.json');
      if (res.ok) this.authorsData = await res.json();
    } catch (_) {}
  }
}
