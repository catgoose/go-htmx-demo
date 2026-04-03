/**
 * Alpine.js CSP component registrations.
 *
 * The CSP build of Alpine (@alpinejs/csp) does not use eval(), so every
 * x-data component must be registered via Alpine.data() rather than using
 * inline object expressions.  This file is loaded BEFORE alpine.min.js
 * (both use defer, so execution order follows source order).
 *
 * Registration happens inside the "alpine:init" event, which the CSP build
 * fires before it walks the DOM.
 */
document.addEventListener('alpine:init', function () {

  // -- Alert toast (body in index.templ) --------------------------------
  Alpine.data('alertListener', function () {
    return {
      showAlert: function (event) {
        var t = document.createElement('div');
        t.className = 'toast toast-end toast-top z-50';
        var a = document.createElement('div');
        a.className = 'alert alert-info shadow-lg';
        a.textContent = event.detail;
        t.appendChild(a);
        document.body.appendChild(t);
        setTimeout(function () {
          t.style.transition = 'opacity 0.3s ease';
          t.style.opacity = '0';
          setTimeout(function () { t.remove(); }, 300);
        }, 3000);
      }
    };
  });

  // -- Theme picker (settings_app.templ) --------------------------------
  Alpine.data('themePicker', function () {
    return {
      current: '',
      init: function () {
        this.current = this.$el.dataset.theme || 'dark';
      },
      setTheme: function (theme) {
        this.current = theme;
        document.documentElement.dataset.theme = theme;
        var t = document.querySelector('meta[name="csrf-token"]');
        fetch('/settings/theme', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'X-CSRF-Token': t ? t.content : ''
          },
          body: 'theme=' + theme
        });
        if (window.appChannel) {
          window.appChannel.postMessage({ type: 'theme-change', theme: theme });
        }
      }
    };
  });

  // -- Error trace row (error_traces.templ) -----------------------------
  Alpine.data('traceRow', function () {
    return {
      expanded: false,
      loaded: false,
      toggle: function () {
        this.expanded = !this.expanded;
        if (this.expanded && !this.loaded) {
          this.loaded = true;
          htmx.trigger(this.$el.querySelector('[hx-get]'), 'expand');
        }
      },
      collapse: function () {
        this.expanded = false;
      }
    };
  });

  // -- Expandable (log entry attrs in error_traces.templ) ---------------
  Alpine.data('expandable', function () {
    return {
      open: false,
      toggle: function () { this.open = !this.open; }
    };
  });

  // -- Auto-open modal (report_email.templ) -----------------------------
  Alpine.data('autoModal', function () {
    return {
      init: function () {
        var el = this.$el;
        this.$nextTick(function () { el.showModal(); });
      }
    };
  });

  // -- Bulk select-all checkbox (bulk.templ) ----------------------------
  Alpine.data('bulkSelectAll', function () {
    return {
      toggleAll: function () {
        var checked = this.$el.querySelector('.select-all-check').checked;
        this.$el.closest('table').querySelectorAll('.row-check').forEach(function (cb) {
          cb.checked = checked;
        });
      }
    };
  });

  // -- Bulk row click-to-toggle (bulk.templ) ----------------------------
  Alpine.data('bulkRowToggle', function () {
    return {
      toggleRow: function (event) {
        if (event.target.tagName !== 'INPUT') {
          var cb = this.$el.querySelector('.row-check');
          cb.checked = !cb.checked;
        }
      }
    };
  });

  // -- Locale Intl formatters (hypermedia_components3.templ) -------------
  Alpine.data('intlRelativeTime', function () {
    return {
      formatted: '',
      init: function () {
        this.formatted = new Intl.RelativeTimeFormat(navigator.language, { numeric: 'auto' }).format(-2, 'hour');
      }
    };
  });

  Alpine.data('intlCurrency', function () {
    return {
      formatted: '',
      init: function () {
        this.formatted = new Intl.NumberFormat(navigator.language, { style: 'currency', currency: 'USD' }).format(1234.56);
      }
    };
  });

  Alpine.data('intlList', function () {
    return {
      formatted: '',
      init: function () {
        this.formatted = new Intl.ListFormat(navigator.language, { style: 'long', type: 'conjunction' }).format(['Alice', 'Bob', 'Charlie']);
      }
    };
  });

  Alpine.data('intlDate', function () {
    return {
      formatted: '',
      init: function () {
        this.formatted = new Intl.DateTimeFormat(navigator.language, { dateStyle: 'full' }).format(new Date());
      }
    };
  });

  // -- Range input live output (filter.templ) ----------------------------
  Alpine.data('rangeOutput', function () {
    return {
      updateOutput: function () {
        var input = this.$el.querySelector('input[type="range"]');
        var output = this.$el.querySelector('output');
        if (input && output) {
          output.textContent = input.value;
        }
      }
    };
  });

  // -- NavBar close-on-outside-click (nav.templ) ------------------------
  Alpine.data('navBar', function () {
    return {
      closeOthers: function (event) {
        var el = this.$el;
        el.querySelectorAll('details[open]').forEach(function (d) {
          if (!d.contains(event.target)) {
            d.open = false;
          }
        });
      }
    };
  });

  // -- NavMenu details exclusive toggle (nav.templ) ---------------------
  Alpine.data('navMenuDropdown', function () {
    return {
      closeOtherDropdowns: function () {
        var el = this.$el;
        if (el.open) {
          el.closest('ul.menu-horizontal').querySelectorAll('details').forEach(function (d) {
            if (d !== el) {
              d.open = false;
            }
          });
        }
      }
    };
  });

  // -- Error copy-to-clipboard (error_status.templ) ---------------------
  Alpine.data('errorCopy', function () {
    return {
      copyError: function () {
        navigator.clipboard.writeText(this.$el.dataset.errorJson);
        var tip = this.$refs.copyTip;
        if (tip) {
          tip.classList.remove('hidden');
          setTimeout(function () { tip.classList.add('hidden'); }, 1500);
        }
      }
    };
  });

  // -- Dismiss inline error (error_status.templ) ------------------------
  Alpine.data('dismissError', function () {
    return {
      dismiss: function () {
        var container = this.$el.closest('div[id]');
        if (container) {
          container.innerHTML = '';
        }
      }
    };
  });

  // -- Close parent dialog after HTMX request (modal.templ) -------------
  Alpine.data('modalSubmit', function () {
    return {
      closeDialog: function () {
        var dialog = this.$el.closest('dialog');
        if (dialog) {
          dialog.close();
        }
      }
    };
  });

  // -- Report issue toggle (report_issue.templ) -------------------------
  Alpine.data('reportIssueForm', function () {
    return {
      showMessageField: function () {
        var toggle = this.$refs.addToggle;
        var field = this.$refs.msgField;
        if (toggle) {
          toggle.remove();
        }
        if (field) {
          field.classList.remove('hidden');
          var textarea = field.querySelector('textarea');
          if (textarea) {
            textarea.focus();
          }
        }
      }
    };
  });

  // -- Interval control with unit cycling (interval.templ) ---------------
  Alpine.data('intervalControl', function () {
    var units = ['ms', 's', 'min', 'h'];
    var configs = {
      ms:  { min: 100, max: 2000, step: 100, mult: 1 },
      s:   { min: 1,   max: 60,   step: 1,   mult: 1000 },
      min: { min: 1,   max: 60,   step: 1,   mult: 60000 },
      h:   { min: 1,   max: 24,   step: 1,   mult: 3600000 }
    };
    return {
      unitIdx: 0,
      init: function () {
        var unit = this.$el.dataset.unit || 's';
        var idx = units.indexOf(unit);
        this.unitIdx = idx >= 0 ? idx : 1;
      },
      cycleUnit: function () {
        var input = this.$el.querySelector('input[type=range]');
        var display = this.$el.querySelector('[data-display]');
        var unitEl = this.$el.querySelector('[data-unit-label]');
        if (!input) return;
        var oldCfg = configs[units[this.unitIdx]];
        var ms = parseInt(input.value) * oldCfg.mult;
        this.unitIdx = (this.unitIdx + 1) % units.length;
        var unit = units[this.unitIdx];
        var cfg = configs[unit];
        var val = Math.round(ms / cfg.mult);
        if (val < cfg.min) val = cfg.min;
        if (val > cfg.max) val = cfg.max;
        input.min = cfg.min;
        input.max = cfg.max;
        input.step = cfg.step;
        input.value = val;
        if (display) display.textContent = val;
        if (unitEl) unitEl.textContent = unit;
        this.postInterval(val, cfg.mult);
      },
      onInput: function () {
        var input = this.$el.querySelector('input[type=range]');
        var display = this.$el.querySelector('[data-display]');
        if (input && display) display.textContent = input.value;
      },
      onChange: function () {
        var input = this.$el.querySelector('input[type=range]');
        if (!input) return;
        var cfg = configs[units[this.unitIdx]];
        this.postInterval(parseInt(input.value), cfg.mult);
      },
      postInterval: function (val, mult) {
        var url = this.$el.dataset.postUrl;
        var key = this.$el.dataset.targetKey;
        var value = this.$el.dataset.targetValue;
        if (!url) return;
        var body = {};
        body[key] = value;
        body['interval_ms'] = (val * mult).toString();
        htmx.ajax('POST', url, { values: body, swap: 'none' });
      }
    };
  });

  // -- Existing global functions that need CSP registration -------------
  if (typeof offlineIndicator === 'function') {
    Alpine.data('offlineIndicator', offlineIndicator);
  }

});
