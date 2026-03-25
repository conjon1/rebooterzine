/* ── loader.js ───────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  async loadPost(id) {
    this.contentDiv.innerHTML = '<p class="p-8 text-[#555555]" style="font-family:var(--font-term);">> fetching transmission...</p>';
    try {
      const res = await fetch(`posts/${id}.html`);
      if (!res.ok) throw new Error('transmission missing or redacted');
      const doc = new DOMParser().parseFromString(await res.text(), 'text/html');

      const authorName = doc.querySelector('meta[name="author"]')?.content;
      if (authorName) {
        const authorObj = this.authorsData.find(a => a.name === authorName);
        const nameHtml  = authorObj
          ? `<a href="#author/${authorObj.id}" class="text-[#5ec4d4] hover:text-white transition-colors">${authorName}</a>`
          : `<span class="text-[#5ec4d4]">${authorName}</span>`;
        const byline = `
          <div class="flex items-center gap-2 mb-6 mt-1 text-[13px] text-[#555555]" style="font-family:var(--font-term);">
            <span>>></span><span>author ::</span>${nameHtml}
          </div>`;
        const h1 = doc.querySelector('h1');
        if (h1) { h1.style.fontFamily = 'var(--font-term)'; h1.insertAdjacentHTML('afterend', byline); }
        else doc.body.insertAdjacentHTML('afterbegin', byline);
      }
      this.contentDiv.innerHTML = doc.body.innerHTML;
      setTimeout(() => { if (typeof this.initObservers === 'function') this.initObservers(); }, 50);
    } catch (e) {
      this.contentDiv.innerHTML = `
        <div class="border border-[#e05555] p-6 m-6">
          <p class="text-[10px] text-[#e05555] mb-2" style="font-family:var(--font-pixel);">[ERR] FAILED</p>
          <p class="text-[13px] text-[#888888]" style="font-family:var(--font-term);">${e.message}</p>
        </div>`;
    }
  },

  async loadAuthorPage(id) {
    this.contentDiv.innerHTML = '<p class="p-8 text-[#555555]" style="font-family:var(--font-term);">> locating personnel file...</p>';
    try {
      const res = await fetch(`authors/${id}.html`);
      if (!res.ok) throw new Error('personnel file not found');
      const doc = new DOMParser().parseFromString(await res.text(), 'text/html');
      this.contentDiv.innerHTML = doc.body.innerHTML;
      setTimeout(() => { if (typeof this.initObservers === 'function') this.initObservers(); }, 50);
    } catch (e) {
      this.contentDiv.innerHTML = `
        <div class="border border-[#e05555] p-6 m-6">
          <p class="text-[10px] text-[#e05555] mb-2" style="font-family:var(--font-pixel);">[ERR] ACCESS DENIED</p>
          <p class="text-[13px] text-[#888888]" style="font-family:var(--font-term);">${e.message}</p>
        </div>`;
    }
  },

});
