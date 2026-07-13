// Docs search: client-side filter over /docs/index.json.
(function () {
  var input = document.getElementById('docs-search');
  var results = document.getElementById('docs-search-results');
  if (!input || !results) return;

  var index = null;
  var selected = -1;

  function loadIndex() {
    if (index) return Promise.resolve(index);
    return fetch('/docs/index.json')
      .then(function (r) { return r.json(); })
      .then(function (data) { index = data; return index; })
      .catch(function () { return []; });
  }

  function score(entry, q) {
    var title = entry.title.toLowerCase();
    var section = (entry.section || '').toLowerCase();
    if (title.indexOf(q) === 0) return 0;
    if (title.indexOf(q) !== -1) return 1;
    if (section.indexOf(q) !== -1) return 2;
    var heads = entry.headings || [];
    for (var i = 0; i < heads.length; i++) {
      if (heads[i].toLowerCase().indexOf(q) !== -1) return 3;
    }
    return -1;
  }

  function render(matches) {
    selected = -1;
    if (!matches.length) {
      results.innerHTML = '<div class="px-3 py-3 text-sm text-muted">No results</div>';
      results.classList.remove('hidden');
      return;
    }
    results.innerHTML = matches.map(function (m, i) {
      return '<a href="/docs' + (m.path ? '/' + m.path : '') + '" data-i="' + i + '" ' +
        'class="block px-3 py-2 border-b border-line last:border-b-0 hover:bg-signal-soft">' +
        '<span class="block text-sm font-medium text-ink">' + m.title + '</span>' +
        (m.section ? '<span class="block font-mono text-[10px] uppercase tracking-widest text-muted mt-0.5">' + m.section + '</span>' : '') +
        '</a>';
    }).join('');
    results.classList.remove('hidden');
  }

  function highlight() {
    var items = results.querySelectorAll('a');
    items.forEach(function (el, i) {
      el.classList.toggle('bg-signal-soft', i === selected);
    });
    if (selected >= 0 && items[selected]) items[selected].scrollIntoView({ block: 'nearest' });
  }

  function search() {
    var q = input.value.trim().toLowerCase();
    if (q.length < 2) { results.classList.add('hidden'); return; }
    loadIndex().then(function (idx) {
      var matches = idx
        .map(function (e) { return { e: e, s: score(e, q) }; })
        .filter(function (x) { return x.s >= 0; })
        .sort(function (a, b) { return a.s - b.s; })
        .slice(0, 10)
        .map(function (x) { return x.e; });
      render(matches);
    });
  }

  input.addEventListener('input', search);
  input.addEventListener('focus', function () { loadIndex(); if (input.value.trim().length >= 2) search(); });

  input.addEventListener('keydown', function (e) {
    var items = results.querySelectorAll('a');
    if (e.key === 'ArrowDown') { e.preventDefault(); selected = Math.min(selected + 1, items.length - 1); highlight(); }
    else if (e.key === 'ArrowUp') { e.preventDefault(); selected = Math.max(selected - 1, 0); highlight(); }
    else if (e.key === 'Enter' && items.length) { e.preventDefault(); (items[selected >= 0 ? selected : 0]).click(); }
    else if (e.key === 'Escape') { results.classList.add('hidden'); input.blur(); }
  });

  document.addEventListener('keydown', function (e) {
    if ((e.metaKey || e.ctrlKey) && e.key === 'k') { e.preventDefault(); input.focus(); }
    else if (e.key === '/' && document.activeElement !== input &&
             !/^(INPUT|TEXTAREA)$/.test(document.activeElement.tagName)) {
      e.preventDefault(); input.focus();
    }
  });

  document.addEventListener('click', function (e) {
    if (!results.contains(e.target) && e.target !== input) results.classList.add('hidden');
  });
})();
