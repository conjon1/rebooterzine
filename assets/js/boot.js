/* ── boot.js ───────────────────────────────────────────────────────────────
   Boots the application. Must be concatenated last so that all prototype
   extensions in the preceding files are in place before instantiation.
   ────────────────────────────────────────────────────────────────────────── */
document.addEventListener('DOMContentLoaded', () => new BlogApp());
