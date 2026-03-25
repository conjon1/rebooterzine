/* ── boot.js ─────────────────────────────────────────────────────────────── */
document.addEventListener('DOMContentLoaded', () => {
  new BlogApp();

  // Nav spinner
  const frames = ['[-]', '[\\]', '[|]', '[/]'];
  let i = 0;
  const el = document.getElementById('nav-spinner');
  if (el) setInterval(() => { el.innerText = frames[i]; i = (i + 1) % frames.length; }, 150);

  // Radar animation
  const radarFrames = ['  . \n -o-\n  . ', ' . .\n -O-\n . .', '  . \n =O=\n  . ', '    \n ~0~\n    '];
  let r = 0;
  const radar = document.getElementById('radar-ascii');
  if (radar) setInterval(() => { radar.innerText = radarFrames[r]; r = (r + 1) % radarFrames.length; }, 500);
});
