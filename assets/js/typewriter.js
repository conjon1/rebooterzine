/* ── typewriter.js ─────────────────────────────────────────────────────────
   Cycles through role strings in the #typewriter span with a typing effect.
   The .cursor element (a blinking block) is rendered by home.templ.
   ────────────────────────────────────────────────────────────────────────── */
Object.assign(BlogApp.prototype, {

  initTypewriter() {
    const el = document.getElementById('typewriter');
    if (!el) return;

    const roles = [
      'Software Engineers',
      'Cybersecurity Researchers',
      'Free Software Advocates',
      'Technology Consultants',
      'Data Scientists',
      'Linux Enthusiasts',
      'Software Architects',
      'Legacy Modernizers',
      'Platform Builders',
      'Daemon Slayers',
      'Unix Wizards',
      'Coffee-Driven Developers',
      'Documentation Meditators',
      'GNU/Programmers',
      'idots',
    ];

    let roleIdx = 0, charIdx = 0, isDeleting = false;

    const type = () => {
      const current = roles[roleIdx];
      el.textContent = current.substring(0, charIdx + (isDeleting ? -1 : 1));
      charIdx += isDeleting ? -1 : 1;

      if (!isDeleting && charIdx === current.length) {
        setTimeout(() => { isDeleting = true; }, 2000);
      } else if (isDeleting && charIdx === 0) {
        isDeleting = false;
        roleIdx = (roleIdx + 1) % roles.length;
      }
      setTimeout(type, isDeleting ? 100 : 150);
    };

    type();
  },

});
