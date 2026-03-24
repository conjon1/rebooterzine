/* ── loader.js ─────────────────────────────────────────────────────────────
   Fetches post and author HTML files and injects them into #dynamic-content.
   For posts, also injects a linked author byline below the first h1.
   ────────────────────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  async loadPost(id) {
    this.contentDiv.innerHTML = '<p class="text-center animate-pulse" style="font-family: var(--font-body); color: var(--ink3);">Fetching transmission...</p>';
    try {
      const res = await fetch(`posts/${id}.html`);
      if (!res.ok) throw new Error('Transmission missing or redacted');
      const doc = new DOMParser().parseFromString(await res.text(), 'text/html');

      const authorName = doc.querySelector('meta[name="author"]')?.content;
      if (authorName) {
        const authorObj = this.authorsData.find(a => a.name === authorName);
        const nameHtml  = authorObj
          ? `<a href="#author/${authorObj.id}" class="text-[#c42a1c] hover:text-[#18130e] hover:underline decoration-2 underline-offset-4 transition-colors">${authorName}</a>`
          : `<span class="text-[#c42a1c]">${authorName}</span>`;
        
        // Zine style byline tag
        const byline = `
          <div class="flex items-center gap-2 mb-8 mt-2 text-[11px] text-[#4a3e30] uppercase tracking-widest" style="font-family: var(--font-body);">
            <span>Author</span>
            <span class="text-[#c42a1c]">✦</span>
            ${nameHtml}
          </div>`;
          
        const h1 = doc.querySelector('h1');
        if (h1) {
            // Apply riso class and font to h1 if it exists
            h1.style.fontFamily = "var(--font-display)";
            h1.style.textTransform = "uppercase";
            h1.insertAdjacentHTML('afterend', byline);
        } else {
            doc.body.insertAdjacentHTML('afterbegin', byline);
        }
      }
      this.contentDiv.innerHTML = doc.body.innerHTML;
      
      // Re-initialize observers for newly injected content
      setTimeout(() => {
        if (typeof this.initObservers === 'function') this.initObservers();
      }, 50);

    } catch (e) {
      this.contentDiv.innerHTML = `
        <div class="border-[2px] border-[#c42a1c] bg-[#e8e2d4] p-4 m-4">
          <h4 class="text-[9px] text-[#c42a1c] uppercase tracking-widest mb-2" style="font-family: var(--font-pixel);">ERROR</h4>
          <p class="text-[12px] text-[#2e251c]" style="font-family: var(--font-body);">${e.message}</p>
        </div>`;
    }
  },

  async loadAuthorPage(id) {
    this.contentDiv.innerHTML = '<p class="text-center animate-pulse" style="font-family: var(--font-body); color: var(--ink3);">Locating personnel file...</p>';
    try {
      const res = await fetch(`authors/${id}.html`);
      if (!res.ok) throw new Error('Personnel file restricted or not found');
      const doc = new DOMParser().parseFromString(await res.text(), 'text/html');
      
      this.contentDiv.innerHTML = doc.body.innerHTML;
      
      // Re-initialize observers for newly injected content
      setTimeout(() => {
        if (typeof this.initObservers === 'function') this.initObservers();
      }, 50);

    } catch (e) {
      this.contentDiv.innerHTML = `
        <div class="border-[2px] border-[#c42a1c] bg-[#e8e2d4] p-4 m-4">
          <h4 class="text-[9px] text-[#c42a1c] uppercase tracking-widest mb-2" style="font-family: var(--font-pixel);">ERROR</h4>
          <p class="text-[12px] text-[#2e251c]" style="font-family: var(--font-body);">${e.message}</p>
        </div>`;
    }
  },

});
