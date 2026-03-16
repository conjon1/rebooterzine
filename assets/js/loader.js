/* ── loader.js ─────────────────────────────────────────────────────────────
   Fetches post and author HTML files and injects them into #dynamic-content.
   For posts, also injects a linked author byline below the first <h1>.
   ────────────────────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  async loadPost(id) {
    this.contentDiv.innerHTML = '<p class="text-center animate-pulse">Loading content...</p>';
    try {
      const res = await fetch(`posts/${id}.html`);
      if (!res.ok) throw new Error('Post not found');
      const doc = new DOMParser().parseFromString(await res.text(), 'text/html');

      const authorName = doc.querySelector('meta[name="author"]')?.content;
      if (authorName) {
        const authorObj = this.authorsData.find(a => a.name === authorName);
        const nameHtml  = authorObj
          ? `<a href="#author/${authorObj.id}" class="text-[#c678dd] font-semibold hover:underline hover:text-[#61afef] transition-colors">${authorName}</a>`
          : `<span class="text-[#c678dd] font-semibold">${authorName}</span>`;
        const byline = `<p class="text-sm text-[#5c6370] mb-8 -mt-2">Written by ${nameHtml}</p>`;
        const h1 = doc.querySelector('h1');
        h1
          ? h1.insertAdjacentHTML('afterend', byline)
          : doc.body.insertAdjacentHTML('afterbegin', byline);
      }
      this.contentDiv.innerHTML = doc.body.innerHTML;
    } catch (e) {
      this.contentDiv.innerHTML = `<p class="text-[#e06c75]">Error loading post: ${e.message}</p>`;
    }
  },

  async loadAuthorPage(id) {
    this.contentDiv.innerHTML = '<p class="text-center animate-pulse">Loading author profile...</p>';
    try {
      const res = await fetch(`authors/${id}.html`);
      if (!res.ok) throw new Error('Author profile not found');
      const doc = new DOMParser().parseFromString(await res.text(), 'text/html');
      this.contentDiv.innerHTML = doc.body.innerHTML;
    } catch (e) {
      this.contentDiv.innerHTML = `<p class="text-[#e06c75]">Error loading author: ${e.message}</p>`;
    }
  },

});
