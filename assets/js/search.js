/* ── search.js ─────────────────────────────────────────────────────────────
   Live search: filters #all-posts articles by title and excerpt text.
   Hides non-matching articles immediately on input.
   ────────────────────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  initSearch() {
    const input = document.getElementById('post-search');
    if (!input) return;
    input.addEventListener('input', e => {
      const term = e.target.value.toLowerCase();
      document.querySelectorAll('#all-posts article').forEach(post => {
        const title   = post.querySelector('h3').textContent.toLowerCase();
        const excerpt = post.querySelector('.excerpt').textContent.toLowerCase();
        
        // Use 'flex' instead of 'block' to maintain the column layout of the zine post items
        post.style.display = (title.includes(term) || excerpt.includes(term)) ? 'flex' : 'none';
      });
    });
  },

});
