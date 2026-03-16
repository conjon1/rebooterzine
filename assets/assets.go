package assets

import _ "embed"

// ── CSS ───────────────────────────────────────────────────────────────────
// Files are embedded individually so the concat order is explicit and
// enforced at compile time. variables.css must come first (everything
// else uses its custom properties).

//go:embed css/variables.css
var variablesCSS string

//go:embed css/base.css
var baseCSS string

//go:embed css/animations.css
var animationsCSS string

//go:embed css/prose.css
var proseCSS string

// CSS returns all stylesheets concatenated in dependency order.
// The result is injected directly into the <style> tag in layout.templ.
func CSS() string {
	return variablesCSS + baseCSS + animationsCSS + proseCSS
}

// ── JS ────────────────────────────────────────────────────────────────────
// app.js must come first (defines the class).
// boot.js must come last (instantiates BlogApp after all prototype extensions).

//go:embed js/app.js
var appJS string

//go:embed js/router.js
var routerJS string

//go:embed js/loader.js
var loaderJS string

//go:embed js/search.js
var searchJS string

//go:embed js/typewriter.js
var typewriterJS string

//go:embed js/observers.js
var observersJS string

//go:embed js/mobile.js
var mobileJS string

//go:embed js/boot.js
var bootJS string

// JS returns all scripts concatenated in dependency order.
// The result is injected directly into the <script> tag in layout.templ.
func JS() string {
	return appJS + routerJS + loaderJS + searchJS + typewriterJS + observersJS + mobileJS + bootJS
}
