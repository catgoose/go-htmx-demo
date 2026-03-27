// setup:feature:demo
/**
 * Alpine.js component that populates the context bar from <link> tags.
 * Reads link relations from the document head and renders them.
 * - rel="bookmark" — frecent (user's frequently visited pages) — left side
 * - rel="related" — declared related pages — right side
 * @returns {AlpineComponent}
 */
function contextBar() {
  return {
    frecent: [],
    related: [],
    init() {
      this.frecent = this.readLinks('bookmark');
      this.related = this.readLinks('related');
    },
    readLinks(rel) {
      return Array.from(document.querySelectorAll('head link[rel="' + rel + '"]'))
        .map(function(el) { return { href: el.getAttribute('href'), title: el.getAttribute('title') }; })
        .filter(function(l) { return l.href && l.title && l.href !== window.location.pathname; });
    }
  };
}
