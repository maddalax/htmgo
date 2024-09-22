// node_modules/htmx.org/dist/htmx.esm.js
var htmx2 = function() {
  const htmx = {
    // Tsc madness here, assigning the functions directly results in an invalid TypeScript output, but reassigning is fine
    /* Event processing */
    /** @type {typeof onLoadHelper} */
    onLoad: null,
    /** @type {typeof processNode} */
    process: null,
    /** @type {typeof addEventListenerImpl} */
    on: null,
    /** @type {typeof removeEventListenerImpl} */
    off: null,
    /** @type {typeof triggerEvent} */
    trigger: null,
    /** @type {typeof ajaxHelper} */
    ajax: null,
    /* DOM querying helpers */
    /** @type {typeof find} */
    find: null,
    /** @type {typeof findAll} */
    findAll: null,
    /** @type {typeof closest} */
    closest: null,
    /**
     * Returns the input values that would resolve for a given element via the htmx value resolution mechanism
     *
     * @see https://htmx.org/api/#values
     *
     * @param {Element} elt the element to resolve values on
     * @param {HttpVerb} type the request type (e.g. **get** or **post**) non-GET's will include the enclosing form of the element. Defaults to **post**
     * @returns {Object}
     */
    values: function(elt, type) {
      const inputValues = getInputValues(elt, type || "post");
      return inputValues.values;
    },
    /* DOM manipulation helpers */
    /** @type {typeof removeElement} */
    remove: null,
    /** @type {typeof addClassToElement} */
    addClass: null,
    /** @type {typeof removeClassFromElement} */
    removeClass: null,
    /** @type {typeof toggleClassOnElement} */
    toggleClass: null,
    /** @type {typeof takeClassForElement} */
    takeClass: null,
    /** @type {typeof swap} */
    swap: null,
    /* Extension entrypoints */
    /** @type {typeof defineExtension} */
    defineExtension: null,
    /** @type {typeof removeExtension} */
    removeExtension: null,
    /* Debugging */
    /** @type {typeof logAll} */
    logAll: null,
    /** @type {typeof logNone} */
    logNone: null,
    /* Debugging */
    /**
     * The logger htmx uses to log with
     *
     * @see https://htmx.org/api/#logger
     */
    logger: null,
    /**
     * A property holding the configuration htmx uses at runtime.
     *
     * Note that using a [meta tag](https://htmx.org/docs/#config) is the preferred mechanism for setting these properties.
     *
     * @see https://htmx.org/api/#config
     */
    config: {
      /**
       * Whether to use history.
       * @type boolean
       * @default true
       */
      historyEnabled: true,
      /**
       * The number of pages to keep in **localStorage** for history support.
       * @type number
       * @default 10
       */
      historyCacheSize: 10,
      /**
       * @type boolean
       * @default false
       */
      refreshOnHistoryMiss: false,
      /**
       * The default swap style to use if **[hx-swap](https://htmx.org/attributes/hx-swap)** is omitted.
       * @type HtmxSwapStyle
       * @default 'innerHTML'
       */
      defaultSwapStyle: "innerHTML",
      /**
       * The default delay between receiving a response from the server and doing the swap.
       * @type number
       * @default 0
       */
      defaultSwapDelay: 0,
      /**
       * The default delay between completing the content swap and settling attributes.
       * @type number
       * @default 20
       */
      defaultSettleDelay: 20,
      /**
       * If true, htmx will inject a small amount of CSS into the page to make indicators invisible unless the **htmx-indicator** class is present.
       * @type boolean
       * @default true
       */
      includeIndicatorStyles: true,
      /**
       * The class to place on indicators when a request is in flight.
       * @type string
       * @default 'htmx-indicator'
       */
      indicatorClass: "htmx-indicator",
      /**
       * The class to place on triggering elements when a request is in flight.
       * @type string
       * @default 'htmx-request'
       */
      requestClass: "htmx-request",
      /**
       * The class to temporarily place on elements that htmx has added to the DOM.
       * @type string
       * @default 'htmx-added'
       */
      addedClass: "htmx-added",
      /**
       * The class to place on target elements when htmx is in the settling phase.
       * @type string
       * @default 'htmx-settling'
       */
      settlingClass: "htmx-settling",
      /**
       * The class to place on target elements when htmx is in the swapping phase.
       * @type string
       * @default 'htmx-swapping'
       */
      swappingClass: "htmx-swapping",
      /**
       * Allows the use of eval-like functionality in htmx, to enable **hx-vars**, trigger conditions & script tag evaluation. Can be set to **false** for CSP compatibility.
       * @type boolean
       * @default true
       */
      allowEval: true,
      /**
       * If set to false, disables the interpretation of script tags.
       * @type boolean
       * @default true
       */
      allowScriptTags: true,
      /**
       * If set, the nonce will be added to inline scripts.
       * @type string
       * @default ''
       */
      inlineScriptNonce: "",
      /**
       * If set, the nonce will be added to inline styles.
       * @type string
       * @default ''
       */
      inlineStyleNonce: "",
      /**
       * The attributes to settle during the settling phase.
       * @type string[]
       * @default ['class', 'style', 'width', 'height']
       */
      attributesToSettle: ["class", "style", "width", "height"],
      /**
       * Allow cross-site Access-Control requests using credentials such as cookies, authorization headers or TLS client certificates.
       * @type boolean
       * @default false
       */
      withCredentials: false,
      /**
       * @type number
       * @default 0
       */
      timeout: 0,
      /**
       * The default implementation of **getWebSocketReconnectDelay** for reconnecting after unexpected connection loss by the event code **Abnormal Closure**, **Service Restart** or **Try Again Later**.
       * @type {'full-jitter' | ((retryCount:number) => number)}
       * @default "full-jitter"
       */
      wsReconnectDelay: "full-jitter",
      /**
       * The type of binary data being received over the WebSocket connection
       * @type BinaryType
       * @default 'blob'
       */
      wsBinaryType: "blob",
      /**
       * @type string
       * @default '[hx-disable], [data-hx-disable]'
       */
      disableSelector: "[hx-disable], [data-hx-disable]",
      /**
       * @type {'auto' | 'instant' | 'smooth'}
       * @default 'instant'
       */
      scrollBehavior: "instant",
      /**
       * If the focused element should be scrolled into view.
       * @type boolean
       * @default false
       */
      defaultFocusScroll: false,
      /**
       * If set to true htmx will include a cache-busting parameter in GET requests to avoid caching partial responses by the browser
       * @type boolean
       * @default false
       */
      getCacheBusterParam: false,
      /**
       * If set to true, htmx will use the View Transition API when swapping in new content.
       * @type boolean
       * @default false
       */
      globalViewTransitions: false,
      /**
       * htmx will format requests with these methods by encoding their parameters in the URL, not the request body
       * @type {(HttpVerb)[]}
       * @default ['get', 'delete']
       */
      methodsThatUseUrlParams: ["get", "delete"],
      /**
       * If set to true, disables htmx-based requests to non-origin hosts.
       * @type boolean
       * @default false
       */
      selfRequestsOnly: true,
      /**
       * If set to true htmx will not update the title of the document when a title tag is found in new content
       * @type boolean
       * @default false
       */
      ignoreTitle: false,
      /**
       * Whether the target of a boosted element is scrolled into the viewport.
       * @type boolean
       * @default true
       */
      scrollIntoViewOnBoost: true,
      /**
       * The cache to store evaluated trigger specifications into.
       * You may define a simple object to use a never-clearing cache, or implement your own system using a [proxy object](https://developer.mozilla.org/docs/Web/JavaScript/Reference/Global_Objects/Proxy)
       * @type {Object|null}
       * @default null
       */
      triggerSpecsCache: null,
      /** @type boolean */
      disableInheritance: false,
      /** @type HtmxResponseHandlingConfig[] */
      responseHandling: [
        { code: "204", swap: false },
        { code: "[23]..", swap: true },
        { code: "[45]..", swap: false, error: true }
      ],
      /**
       * Whether to process OOB swaps on elements that are nested within the main response element.
       * @type boolean
       * @default true
       */
      allowNestedOobSwaps: true
    },
    /** @type {typeof parseInterval} */
    parseInterval: null,
    /** @type {typeof internalEval} */
    _: null,
    version: "2.0.2"
  };
  htmx.onLoad = onLoadHelper;
  htmx.process = processNode;
  htmx.on = addEventListenerImpl;
  htmx.off = removeEventListenerImpl;
  htmx.trigger = triggerEvent;
  htmx.ajax = ajaxHelper;
  htmx.find = find;
  htmx.findAll = findAll;
  htmx.closest = closest;
  htmx.remove = removeElement;
  htmx.addClass = addClassToElement;
  htmx.removeClass = removeClassFromElement;
  htmx.toggleClass = toggleClassOnElement;
  htmx.takeClass = takeClassForElement;
  htmx.swap = swap;
  htmx.defineExtension = defineExtension;
  htmx.removeExtension = removeExtension;
  htmx.logAll = logAll;
  htmx.logNone = logNone;
  htmx.parseInterval = parseInterval;
  htmx._ = internalEval;
  const internalAPI = {
    addTriggerHandler,
    bodyContains,
    canAccessLocalStorage,
    findThisElement,
    filterValues,
    swap,
    hasAttribute,
    getAttributeValue,
    getClosestAttributeValue,
    getClosestMatch,
    getExpressionVars,
    getHeaders,
    getInputValues,
    getInternalData,
    getSwapSpecification,
    getTriggerSpecs,
    getTarget,
    makeFragment,
    mergeObjects,
    makeSettleInfo,
    oobSwap,
    querySelectorExt,
    settleImmediately,
    shouldCancel,
    triggerEvent,
    triggerErrorEvent,
    withExtensions
  };
  const VERBS = ["get", "post", "put", "delete", "patch"];
  const VERB_SELECTOR = VERBS.map(function(verb) {
    return "[hx-" + verb + "], [data-hx-" + verb + "]";
  }).join(", ");
  const HEAD_TAG_REGEX = makeTagRegEx("head");
  function makeTagRegEx(tag, global = false) {
    return new RegExp(
      `<${tag}(\\s[^>]*>|>)([\\s\\S]*?)<\\/${tag}>`,
      global ? "gim" : "im"
    );
  }
  function parseInterval(str2) {
    if (str2 == void 0) {
      return void 0;
    }
    let interval = NaN;
    if (str2.slice(-2) == "ms") {
      interval = parseFloat(str2.slice(0, -2));
    } else if (str2.slice(-1) == "s") {
      interval = parseFloat(str2.slice(0, -1)) * 1e3;
    } else if (str2.slice(-1) == "m") {
      interval = parseFloat(str2.slice(0, -1)) * 1e3 * 60;
    } else {
      interval = parseFloat(str2);
    }
    return isNaN(interval) ? void 0 : interval;
  }
  function getRawAttribute(elt, name) {
    return elt instanceof Element && elt.getAttribute(name);
  }
  function hasAttribute(elt, qualifiedName) {
    return !!elt.hasAttribute && (elt.hasAttribute(qualifiedName) || elt.hasAttribute("data-" + qualifiedName));
  }
  function getAttributeValue(elt, qualifiedName) {
    return getRawAttribute(elt, qualifiedName) || getRawAttribute(elt, "data-" + qualifiedName);
  }
  function parentElt(elt) {
    const parent = elt.parentElement;
    if (!parent && elt.parentNode instanceof ShadowRoot) return elt.parentNode;
    return parent;
  }
  function getDocument() {
    return document;
  }
  function getRootNode(elt, global) {
    return elt.getRootNode ? elt.getRootNode({ composed: global }) : getDocument();
  }
  function getClosestMatch(elt, condition) {
    while (elt && !condition(elt)) {
      elt = parentElt(elt);
    }
    return elt || null;
  }
  function getAttributeValueWithDisinheritance(initialElement, ancestor, attributeName) {
    const attributeValue = getAttributeValue(ancestor, attributeName);
    const disinherit = getAttributeValue(ancestor, "hx-disinherit");
    var inherit = getAttributeValue(ancestor, "hx-inherit");
    if (initialElement !== ancestor) {
      if (htmx.config.disableInheritance) {
        if (inherit && (inherit === "*" || inherit.split(" ").indexOf(attributeName) >= 0)) {
          return attributeValue;
        } else {
          return null;
        }
      }
      if (disinherit && (disinherit === "*" || disinherit.split(" ").indexOf(attributeName) >= 0)) {
        return "unset";
      }
    }
    return attributeValue;
  }
  function getClosestAttributeValue(elt, attributeName) {
    let closestAttr = null;
    getClosestMatch(elt, function(e) {
      return !!(closestAttr = getAttributeValueWithDisinheritance(elt, asElement(e), attributeName));
    });
    if (closestAttr !== "unset") {
      return closestAttr;
    }
  }
  function matches(elt, selector) {
    const matchesFunction = elt instanceof Element && (elt.matches || elt.matchesSelector || elt.msMatchesSelector || elt.mozMatchesSelector || elt.webkitMatchesSelector || elt.oMatchesSelector);
    return !!matchesFunction && matchesFunction.call(elt, selector);
  }
  function getStartTag(str2) {
    const tagMatcher = /<([a-z][^\/\0>\x20\t\r\n\f]*)/i;
    const match = tagMatcher.exec(str2);
    if (match) {
      return match[1].toLowerCase();
    } else {
      return "";
    }
  }
  function parseHTML(resp) {
    const parser = new DOMParser();
    return parser.parseFromString(resp, "text/html");
  }
  function takeChildrenFor(fragment, elt) {
    while (elt.childNodes.length > 0) {
      fragment.append(elt.childNodes[0]);
    }
  }
  function duplicateScript(script) {
    const newScript = getDocument().createElement("script");
    forEach(script.attributes, function(attr) {
      newScript.setAttribute(attr.name, attr.value);
    });
    newScript.textContent = script.textContent;
    newScript.async = false;
    if (htmx.config.inlineScriptNonce) {
      newScript.nonce = htmx.config.inlineScriptNonce;
    }
    return newScript;
  }
  function isJavaScriptScriptNode(script) {
    return script.matches("script") && (script.type === "text/javascript" || script.type === "module" || script.type === "");
  }
  function normalizeScriptTags(fragment) {
    Array.from(fragment.querySelectorAll("script")).forEach(
      /** @param {HTMLScriptElement} script */
      (script) => {
        if (isJavaScriptScriptNode(script)) {
          const newScript = duplicateScript(script);
          const parent = script.parentNode;
          try {
            parent.insertBefore(newScript, script);
          } catch (e) {
            logError(e);
          } finally {
            script.remove();
          }
        }
      }
    );
  }
  function makeFragment(response) {
    const responseWithNoHead = response.replace(HEAD_TAG_REGEX, "");
    const startTag = getStartTag(responseWithNoHead);
    let fragment;
    if (startTag === "html") {
      fragment = /** @type DocumentFragmentWithTitle */
      new DocumentFragment();
      const doc = parseHTML(response);
      takeChildrenFor(fragment, doc.body);
      fragment.title = doc.title;
    } else if (startTag === "body") {
      fragment = /** @type DocumentFragmentWithTitle */
      new DocumentFragment();
      const doc = parseHTML(responseWithNoHead);
      takeChildrenFor(fragment, doc.body);
      fragment.title = doc.title;
    } else {
      const doc = parseHTML('<body><template class="internal-htmx-wrapper">' + responseWithNoHead + "</template></body>");
      fragment = /** @type DocumentFragmentWithTitle */
      doc.querySelector("template").content;
      fragment.title = doc.title;
      var titleElement = fragment.querySelector("title");
      if (titleElement && titleElement.parentNode === fragment) {
        titleElement.remove();
        fragment.title = titleElement.innerText;
      }
    }
    if (fragment) {
      if (htmx.config.allowScriptTags) {
        normalizeScriptTags(fragment);
      } else {
        fragment.querySelectorAll("script").forEach((script) => script.remove());
      }
    }
    return fragment;
  }
  function maybeCall(func) {
    if (func) {
      func();
    }
  }
  function isType(o, type) {
    return Object.prototype.toString.call(o) === "[object " + type + "]";
  }
  function isFunction(o) {
    return typeof o === "function";
  }
  function isRawObject(o) {
    return isType(o, "Object");
  }
  function getInternalData(elt) {
    const dataProp = "htmx-internal-data";
    let data = elt[dataProp];
    if (!data) {
      data = elt[dataProp] = {};
    }
    return data;
  }
  function toArray(arr) {
    const returnArr = [];
    if (arr) {
      for (let i = 0; i < arr.length; i++) {
        returnArr.push(arr[i]);
      }
    }
    return returnArr;
  }
  function forEach(arr, func) {
    if (arr) {
      for (let i = 0; i < arr.length; i++) {
        func(arr[i]);
      }
    }
  }
  function isScrolledIntoView(el) {
    const rect = el.getBoundingClientRect();
    const elemTop = rect.top;
    const elemBottom = rect.bottom;
    return elemTop < window.innerHeight && elemBottom >= 0;
  }
  function bodyContains(elt) {
    const rootNode = elt.getRootNode && elt.getRootNode();
    if (rootNode && rootNode instanceof window.ShadowRoot) {
      return getDocument().body.contains(rootNode.host);
    } else {
      return getDocument().body.contains(elt);
    }
  }
  function splitOnWhitespace(trigger) {
    return trigger.trim().split(/\s+/);
  }
  function mergeObjects(obj1, obj2) {
    for (const key in obj2) {
      if (obj2.hasOwnProperty(key)) {
        obj1[key] = obj2[key];
      }
    }
    return obj1;
  }
  function parseJSON(jString) {
    try {
      return JSON.parse(jString);
    } catch (error) {
      logError(error);
      return null;
    }
  }
  function canAccessLocalStorage() {
    const test = "htmx:localStorageTest";
    try {
      localStorage.setItem(test, test);
      localStorage.removeItem(test);
      return true;
    } catch (e) {
      return false;
    }
  }
  function normalizePath(path) {
    try {
      const url = new URL(path);
      if (url) {
        path = url.pathname + url.search;
      }
      if (!/^\/$/.test(path)) {
        path = path.replace(/\/+$/, "");
      }
      return path;
    } catch (e) {
      return path;
    }
  }
  function internalEval(str) {
    return maybeEval(getDocument().body, function() {
      return eval(str);
    });
  }
  function onLoadHelper(callback) {
    const value = htmx.on(
      "htmx:load",
      /** @param {CustomEvent} evt */
      function(evt) {
        callback(evt.detail.elt);
      }
    );
    return value;
  }
  function logAll() {
    htmx.logger = function(elt, event, data) {
      if (console) {
        console.log(event, elt, data);
      }
    };
  }
  function logNone() {
    htmx.logger = null;
  }
  function find(eltOrSelector, selector) {
    if (typeof eltOrSelector !== "string") {
      return eltOrSelector.querySelector(selector);
    } else {
      return find(getDocument(), eltOrSelector);
    }
  }
  function findAll(eltOrSelector, selector) {
    if (typeof eltOrSelector !== "string") {
      return eltOrSelector.querySelectorAll(selector);
    } else {
      return findAll(getDocument(), eltOrSelector);
    }
  }
  function getWindow() {
    return window;
  }
  function removeElement(elt, delay) {
    elt = resolveTarget(elt);
    if (delay) {
      getWindow().setTimeout(function() {
        removeElement(elt);
        elt = null;
      }, delay);
    } else {
      parentElt(elt).removeChild(elt);
    }
  }
  function asElement(elt) {
    return elt instanceof Element ? elt : null;
  }
  function asHtmlElement(elt) {
    return elt instanceof HTMLElement ? elt : null;
  }
  function asString(value) {
    return typeof value === "string" ? value : null;
  }
  function asParentNode(elt) {
    return elt instanceof Element || elt instanceof Document || elt instanceof DocumentFragment ? elt : null;
  }
  function addClassToElement(elt, clazz, delay) {
    elt = asElement(resolveTarget(elt));
    if (!elt) {
      return;
    }
    if (delay) {
      getWindow().setTimeout(function() {
        addClassToElement(elt, clazz);
        elt = null;
      }, delay);
    } else {
      elt.classList && elt.classList.add(clazz);
    }
  }
  function removeClassFromElement(node, clazz, delay) {
    let elt = asElement(resolveTarget(node));
    if (!elt) {
      return;
    }
    if (delay) {
      getWindow().setTimeout(function() {
        removeClassFromElement(elt, clazz);
        elt = null;
      }, delay);
    } else {
      if (elt.classList) {
        elt.classList.remove(clazz);
        if (elt.classList.length === 0) {
          elt.removeAttribute("class");
        }
      }
    }
  }
  function toggleClassOnElement(elt, clazz) {
    elt = resolveTarget(elt);
    elt.classList.toggle(clazz);
  }
  function takeClassForElement(elt, clazz) {
    elt = resolveTarget(elt);
    forEach(elt.parentElement.children, function(child) {
      removeClassFromElement(child, clazz);
    });
    addClassToElement(asElement(elt), clazz);
  }
  function closest(elt, selector) {
    elt = asElement(resolveTarget(elt));
    if (elt && elt.closest) {
      return elt.closest(selector);
    } else {
      do {
        if (elt == null || matches(elt, selector)) {
          return elt;
        }
      } while (elt = elt && asElement(parentElt(elt)));
      return null;
    }
  }
  function startsWith(str2, prefix) {
    return str2.substring(0, prefix.length) === prefix;
  }
  function endsWith(str2, suffix) {
    return str2.substring(str2.length - suffix.length) === suffix;
  }
  function normalizeSelector(selector) {
    const trimmedSelector = selector.trim();
    if (startsWith(trimmedSelector, "<") && endsWith(trimmedSelector, "/>")) {
      return trimmedSelector.substring(1, trimmedSelector.length - 2);
    } else {
      return trimmedSelector;
    }
  }
  function querySelectorAllExt(elt, selector, global) {
    elt = resolveTarget(elt);
    if (selector.indexOf("closest ") === 0) {
      return [closest(asElement(elt), normalizeSelector(selector.substr(8)))];
    } else if (selector.indexOf("find ") === 0) {
      return [find(asParentNode(elt), normalizeSelector(selector.substr(5)))];
    } else if (selector === "next") {
      return [asElement(elt).nextElementSibling];
    } else if (selector.indexOf("next ") === 0) {
      return [scanForwardQuery(elt, normalizeSelector(selector.substr(5)), !!global)];
    } else if (selector === "previous") {
      return [asElement(elt).previousElementSibling];
    } else if (selector.indexOf("previous ") === 0) {
      return [scanBackwardsQuery(elt, normalizeSelector(selector.substr(9)), !!global)];
    } else if (selector === "document") {
      return [document];
    } else if (selector === "window") {
      return [window];
    } else if (selector === "body") {
      return [document.body];
    } else if (selector === "root") {
      return [getRootNode(elt, !!global)];
    } else if (selector.indexOf("global ") === 0) {
      return querySelectorAllExt(elt, selector.slice(7), true);
    } else {
      return toArray(asParentNode(getRootNode(elt, !!global)).querySelectorAll(normalizeSelector(selector)));
    }
  }
  var scanForwardQuery = function(start, match, global) {
    const results = asParentNode(getRootNode(start, global)).querySelectorAll(match);
    for (let i = 0; i < results.length; i++) {
      const elt = results[i];
      if (elt.compareDocumentPosition(start) === Node.DOCUMENT_POSITION_PRECEDING) {
        return elt;
      }
    }
  };
  var scanBackwardsQuery = function(start, match, global) {
    const results = asParentNode(getRootNode(start, global)).querySelectorAll(match);
    for (let i = results.length - 1; i >= 0; i--) {
      const elt = results[i];
      if (elt.compareDocumentPosition(start) === Node.DOCUMENT_POSITION_FOLLOWING) {
        return elt;
      }
    }
  };
  function querySelectorExt(eltOrSelector, selector) {
    if (typeof eltOrSelector !== "string") {
      return querySelectorAllExt(eltOrSelector, selector)[0];
    } else {
      return querySelectorAllExt(getDocument().body, eltOrSelector)[0];
    }
  }
  function resolveTarget(eltOrSelector, context) {
    if (typeof eltOrSelector === "string") {
      return find(asParentNode(context) || document, eltOrSelector);
    } else {
      return eltOrSelector;
    }
  }
  function processEventArgs(arg1, arg2, arg3) {
    if (isFunction(arg2)) {
      return {
        target: getDocument().body,
        event: asString(arg1),
        listener: arg2
      };
    } else {
      return {
        target: resolveTarget(arg1),
        event: asString(arg2),
        listener: arg3
      };
    }
  }
  function addEventListenerImpl(arg1, arg2, arg3) {
    ready(function() {
      const eventArgs = processEventArgs(arg1, arg2, arg3);
      eventArgs.target.addEventListener(eventArgs.event, eventArgs.listener);
    });
    const b = isFunction(arg2);
    return b ? arg2 : arg3;
  }
  function removeEventListenerImpl(arg1, arg2, arg3) {
    ready(function() {
      const eventArgs = processEventArgs(arg1, arg2, arg3);
      eventArgs.target.removeEventListener(eventArgs.event, eventArgs.listener);
    });
    return isFunction(arg2) ? arg2 : arg3;
  }
  const DUMMY_ELT = getDocument().createElement("output");
  function findAttributeTargets(elt, attrName) {
    const attrTarget = getClosestAttributeValue(elt, attrName);
    if (attrTarget) {
      if (attrTarget === "this") {
        return [findThisElement(elt, attrName)];
      } else {
        const result = querySelectorAllExt(elt, attrTarget);
        if (result.length === 0) {
          logError('The selector "' + attrTarget + '" on ' + attrName + " returned no matches!");
          return [DUMMY_ELT];
        } else {
          return result;
        }
      }
    }
  }
  function findThisElement(elt, attribute) {
    return asElement(getClosestMatch(elt, function(elt2) {
      return getAttributeValue(asElement(elt2), attribute) != null;
    }));
  }
  function getTarget(elt) {
    const targetStr = getClosestAttributeValue(elt, "hx-target");
    if (targetStr) {
      if (targetStr === "this") {
        return findThisElement(elt, "hx-target");
      } else {
        return querySelectorExt(elt, targetStr);
      }
    } else {
      const data = getInternalData(elt);
      if (data.boosted) {
        return getDocument().body;
      } else {
        return elt;
      }
    }
  }
  function shouldSettleAttribute(name) {
    const attributesToSettle = htmx.config.attributesToSettle;
    for (let i = 0; i < attributesToSettle.length; i++) {
      if (name === attributesToSettle[i]) {
        return true;
      }
    }
    return false;
  }
  function cloneAttributes(mergeTo, mergeFrom) {
    forEach(mergeTo.attributes, function(attr) {
      if (!mergeFrom.hasAttribute(attr.name) && shouldSettleAttribute(attr.name)) {
        mergeTo.removeAttribute(attr.name);
      }
    });
    forEach(mergeFrom.attributes, function(attr) {
      if (shouldSettleAttribute(attr.name)) {
        mergeTo.setAttribute(attr.name, attr.value);
      }
    });
  }
  function isInlineSwap(swapStyle, target) {
    const extensions2 = getExtensions(target);
    for (let i = 0; i < extensions2.length; i++) {
      const extension = extensions2[i];
      try {
        if (extension.isInlineSwap(swapStyle)) {
          return true;
        }
      } catch (e) {
        logError(e);
      }
    }
    return swapStyle === "outerHTML";
  }
  function oobSwap(oobValue, oobElement, settleInfo) {
    let selector = "#" + getRawAttribute(oobElement, "id");
    let swapStyle = "outerHTML";
    if (oobValue === "true") ; else if (oobValue.indexOf(":") > 0) {
      swapStyle = oobValue.substr(0, oobValue.indexOf(":"));
      selector = oobValue.substr(oobValue.indexOf(":") + 1, oobValue.length);
    } else {
      swapStyle = oobValue;
    }
    const targets = getDocument().querySelectorAll(selector);
    if (targets) {
      forEach(
        targets,
        function(target) {
          let fragment;
          const oobElementClone = oobElement.cloneNode(true);
          fragment = getDocument().createDocumentFragment();
          fragment.appendChild(oobElementClone);
          if (!isInlineSwap(swapStyle, target)) {
            fragment = asParentNode(oobElementClone);
          }
          const beforeSwapDetails = { shouldSwap: true, target, fragment };
          if (!triggerEvent(target, "htmx:oobBeforeSwap", beforeSwapDetails)) return;
          target = beforeSwapDetails.target;
          if (beforeSwapDetails.shouldSwap) {
            swapWithStyle(swapStyle, target, target, fragment, settleInfo);
          }
          forEach(settleInfo.elts, function(elt) {
            triggerEvent(elt, "htmx:oobAfterSwap", beforeSwapDetails);
          });
        }
      );
      oobElement.parentNode.removeChild(oobElement);
    } else {
      oobElement.parentNode.removeChild(oobElement);
      triggerErrorEvent(getDocument().body, "htmx:oobErrorNoTarget", { content: oobElement });
    }
    return oobValue;
  }
  function handlePreservedElements(fragment) {
    forEach(findAll(fragment, "[hx-preserve], [data-hx-preserve]"), function(preservedElt) {
      const id = getAttributeValue(preservedElt, "id");
      const oldElt = getDocument().getElementById(id);
      if (oldElt != null) {
        preservedElt.parentNode.replaceChild(oldElt, preservedElt);
      }
    });
  }
  function handleAttributes(parentNode, fragment, settleInfo) {
    forEach(fragment.querySelectorAll("[id]"), function(newNode) {
      const id = getRawAttribute(newNode, "id");
      if (id && id.length > 0) {
        const normalizedId = id.replace("'", "\\'");
        const normalizedTag = newNode.tagName.replace(":", "\\:");
        const parentElt2 = asParentNode(parentNode);
        const oldNode = parentElt2 && parentElt2.querySelector(normalizedTag + "[id='" + normalizedId + "']");
        if (oldNode && oldNode !== parentElt2) {
          const newAttributes = newNode.cloneNode();
          cloneAttributes(newNode, oldNode);
          settleInfo.tasks.push(function() {
            cloneAttributes(newNode, newAttributes);
          });
        }
      }
    });
  }
  function makeAjaxLoadTask(child) {
    return function() {
      removeClassFromElement(child, htmx.config.addedClass);
      processNode(asElement(child));
      processFocus(asParentNode(child));
      triggerEvent(child, "htmx:load");
    };
  }
  function processFocus(child) {
    const autofocus = "[autofocus]";
    const autoFocusedElt = asHtmlElement(matches(child, autofocus) ? child : child.querySelector(autofocus));
    if (autoFocusedElt != null) {
      autoFocusedElt.focus();
    }
  }
  function insertNodesBefore(parentNode, insertBefore, fragment, settleInfo) {
    handleAttributes(parentNode, fragment, settleInfo);
    while (fragment.childNodes.length > 0) {
      const child = fragment.firstChild;
      addClassToElement(asElement(child), htmx.config.addedClass);
      parentNode.insertBefore(child, insertBefore);
      if (child.nodeType !== Node.TEXT_NODE && child.nodeType !== Node.COMMENT_NODE) {
        settleInfo.tasks.push(makeAjaxLoadTask(child));
      }
    }
  }
  function stringHash(string, hash) {
    let char = 0;
    while (char < string.length) {
      hash = (hash << 5) - hash + string.charCodeAt(char++) | 0;
    }
    return hash;
  }
  function attributeHash(elt) {
    let hash = 0;
    if (elt.attributes) {
      for (let i = 0; i < elt.attributes.length; i++) {
        const attribute = elt.attributes[i];
        if (attribute.value) {
          hash = stringHash(attribute.name, hash);
          hash = stringHash(attribute.value, hash);
        }
      }
    }
    return hash;
  }
  function deInitOnHandlers(elt) {
    const internalData = getInternalData(elt);
    if (internalData.onHandlers) {
      for (let i = 0; i < internalData.onHandlers.length; i++) {
        const handlerInfo = internalData.onHandlers[i];
        removeEventListenerImpl(elt, handlerInfo.event, handlerInfo.listener);
      }
      delete internalData.onHandlers;
    }
  }
  function deInitNode(element) {
    const internalData = getInternalData(element);
    if (internalData.timeout) {
      clearTimeout(internalData.timeout);
    }
    if (internalData.listenerInfos) {
      forEach(internalData.listenerInfos, function(info) {
        if (info.on) {
          removeEventListenerImpl(info.on, info.trigger, info.listener);
        }
      });
    }
    deInitOnHandlers(element);
    forEach(Object.keys(internalData), function(key) {
      delete internalData[key];
    });
  }
  function cleanUpElement(element) {
    triggerEvent(element, "htmx:beforeCleanupElement");
    deInitNode(element);
    if (element.children) {
      forEach(element.children, function(child) {
        cleanUpElement(child);
      });
    }
  }
  function swapOuterHTML(target, fragment, settleInfo) {
    if (target instanceof Element && target.tagName === "BODY") {
      return swapInnerHTML(target, fragment, settleInfo);
    }
    let newElt;
    const eltBeforeNewContent = target.previousSibling;
    insertNodesBefore(parentElt(target), target, fragment, settleInfo);
    if (eltBeforeNewContent == null) {
      newElt = parentElt(target).firstChild;
    } else {
      newElt = eltBeforeNewContent.nextSibling;
    }
    settleInfo.elts = settleInfo.elts.filter(function(e) {
      return e !== target;
    });
    while (newElt && newElt !== target) {
      if (newElt instanceof Element) {
        settleInfo.elts.push(newElt);
      }
      newElt = newElt.nextSibling;
    }
    cleanUpElement(target);
    if (target instanceof Element) {
      target.remove();
    } else {
      target.parentNode.removeChild(target);
    }
  }
  function swapAfterBegin(target, fragment, settleInfo) {
    return insertNodesBefore(target, target.firstChild, fragment, settleInfo);
  }
  function swapBeforeBegin(target, fragment, settleInfo) {
    return insertNodesBefore(parentElt(target), target, fragment, settleInfo);
  }
  function swapBeforeEnd(target, fragment, settleInfo) {
    return insertNodesBefore(target, null, fragment, settleInfo);
  }
  function swapAfterEnd(target, fragment, settleInfo) {
    return insertNodesBefore(parentElt(target), target.nextSibling, fragment, settleInfo);
  }
  function swapDelete(target) {
    cleanUpElement(target);
    return parentElt(target).removeChild(target);
  }
  function swapInnerHTML(target, fragment, settleInfo) {
    const firstChild = target.firstChild;
    insertNodesBefore(target, firstChild, fragment, settleInfo);
    if (firstChild) {
      while (firstChild.nextSibling) {
        cleanUpElement(firstChild.nextSibling);
        target.removeChild(firstChild.nextSibling);
      }
      cleanUpElement(firstChild);
      target.removeChild(firstChild);
    }
  }
  function swapWithStyle(swapStyle, elt, target, fragment, settleInfo) {
    switch (swapStyle) {
      case "none":
        return;
      case "outerHTML":
        swapOuterHTML(target, fragment, settleInfo);
        return;
      case "afterbegin":
        swapAfterBegin(target, fragment, settleInfo);
        return;
      case "beforebegin":
        swapBeforeBegin(target, fragment, settleInfo);
        return;
      case "beforeend":
        swapBeforeEnd(target, fragment, settleInfo);
        return;
      case "afterend":
        swapAfterEnd(target, fragment, settleInfo);
        return;
      case "delete":
        swapDelete(target);
        return;
      default:
        var extensions2 = getExtensions(elt);
        for (let i = 0; i < extensions2.length; i++) {
          const ext = extensions2[i];
          try {
            const newElements = ext.handleSwap(swapStyle, target, fragment, settleInfo);
            if (newElements) {
              if (Array.isArray(newElements)) {
                for (let j = 0; j < newElements.length; j++) {
                  const child = newElements[j];
                  if (child.nodeType !== Node.TEXT_NODE && child.nodeType !== Node.COMMENT_NODE) {
                    settleInfo.tasks.push(makeAjaxLoadTask(child));
                  }
                }
              }
              return;
            }
          } catch (e) {
            logError(e);
          }
        }
        if (swapStyle === "innerHTML") {
          swapInnerHTML(target, fragment, settleInfo);
        } else {
          swapWithStyle(htmx.config.defaultSwapStyle, elt, target, fragment, settleInfo);
        }
    }
  }
  function findAndSwapOobElements(fragment, settleInfo) {
    var oobElts = findAll(fragment, "[hx-swap-oob], [data-hx-swap-oob]");
    forEach(oobElts, function(oobElement) {
      if (htmx.config.allowNestedOobSwaps || oobElement.parentElement === null) {
        const oobValue = getAttributeValue(oobElement, "hx-swap-oob");
        if (oobValue != null) {
          oobSwap(oobValue, oobElement, settleInfo);
        }
      } else {
        oobElement.removeAttribute("hx-swap-oob");
        oobElement.removeAttribute("data-hx-swap-oob");
      }
    });
    return oobElts.length > 0;
  }
  function swap(target, content, swapSpec, swapOptions) {
    if (!swapOptions) {
      swapOptions = {};
    }
    target = resolveTarget(target);
    const activeElt = document.activeElement;
    let selectionInfo = {};
    try {
      selectionInfo = {
        elt: activeElt,
        // @ts-ignore
        start: activeElt ? activeElt.selectionStart : null,
        // @ts-ignore
        end: activeElt ? activeElt.selectionEnd : null
      };
    } catch (e) {
    }
    const settleInfo = makeSettleInfo(target);
    if (swapSpec.swapStyle === "textContent") {
      target.textContent = content;
    } else {
      let fragment = makeFragment(content);
      settleInfo.title = fragment.title;
      if (swapOptions.selectOOB) {
        const oobSelectValues = swapOptions.selectOOB.split(",");
        for (let i = 0; i < oobSelectValues.length; i++) {
          const oobSelectValue = oobSelectValues[i].split(":", 2);
          let id = oobSelectValue[0].trim();
          if (id.indexOf("#") === 0) {
            id = id.substring(1);
          }
          const oobValue = oobSelectValue[1] || "true";
          const oobElement = fragment.querySelector("#" + id);
          if (oobElement) {
            oobSwap(oobValue, oobElement, settleInfo);
          }
        }
      }
      findAndSwapOobElements(fragment, settleInfo);
      forEach(
        findAll(fragment, "template"),
        /** @param {HTMLTemplateElement} template */
        function(template) {
          if (findAndSwapOobElements(template.content, settleInfo)) {
            template.remove();
          }
        }
      );
      if (swapOptions.select) {
        const newFragment = getDocument().createDocumentFragment();
        forEach(fragment.querySelectorAll(swapOptions.select), function(node) {
          newFragment.appendChild(node);
        });
        fragment = newFragment;
      }
      handlePreservedElements(fragment);
      swapWithStyle(swapSpec.swapStyle, swapOptions.contextElement, target, fragment, settleInfo);
    }
    if (selectionInfo.elt && !bodyContains(selectionInfo.elt) && getRawAttribute(selectionInfo.elt, "id")) {
      const newActiveElt = document.getElementById(getRawAttribute(selectionInfo.elt, "id"));
      const focusOptions = { preventScroll: swapSpec.focusScroll !== void 0 ? !swapSpec.focusScroll : !htmx.config.defaultFocusScroll };
      if (newActiveElt) {
        if (selectionInfo.start && newActiveElt.setSelectionRange) {
          try {
            newActiveElt.setSelectionRange(selectionInfo.start, selectionInfo.end);
          } catch (e) {
          }
        }
        newActiveElt.focus(focusOptions);
      }
    }
    target.classList.remove(htmx.config.swappingClass);
    forEach(settleInfo.elts, function(elt) {
      if (elt.classList) {
        elt.classList.add(htmx.config.settlingClass);
      }
      triggerEvent(elt, "htmx:afterSwap", swapOptions.eventInfo);
    });
    if (swapOptions.afterSwapCallback) {
      swapOptions.afterSwapCallback();
    }
    if (!swapSpec.ignoreTitle) {
      handleTitle(settleInfo.title);
    }
    const doSettle = function() {
      forEach(settleInfo.tasks, function(task) {
        task.call();
      });
      forEach(settleInfo.elts, function(elt) {
        if (elt.classList) {
          elt.classList.remove(htmx.config.settlingClass);
        }
        triggerEvent(elt, "htmx:afterSettle", swapOptions.eventInfo);
      });
      if (swapOptions.anchor) {
        const anchorTarget = asElement(resolveTarget("#" + swapOptions.anchor));
        if (anchorTarget) {
          anchorTarget.scrollIntoView({ block: "start", behavior: "auto" });
        }
      }
      updateScrollState(settleInfo.elts, swapSpec);
      if (swapOptions.afterSettleCallback) {
        swapOptions.afterSettleCallback();
      }
    };
    if (swapSpec.settleDelay > 0) {
      getWindow().setTimeout(doSettle, swapSpec.settleDelay);
    } else {
      doSettle();
    }
  }
  function handleTriggerHeader(xhr, header, elt) {
    const triggerBody = xhr.getResponseHeader(header);
    if (triggerBody.indexOf("{") === 0) {
      const triggers = parseJSON(triggerBody);
      for (const eventName in triggers) {
        if (triggers.hasOwnProperty(eventName)) {
          let detail = triggers[eventName];
          if (isRawObject(detail)) {
            elt = detail.target !== void 0 ? detail.target : elt;
          } else {
            detail = { value: detail };
          }
          triggerEvent(elt, eventName, detail);
        }
      }
    } else {
      const eventNames = triggerBody.split(",");
      for (let i = 0; i < eventNames.length; i++) {
        triggerEvent(elt, eventNames[i].trim(), []);
      }
    }
  }
  const WHITESPACE_OR_COMMA = /[\s,]/;
  const SYMBOL_START = /[_$a-zA-Z]/;
  const SYMBOL_CONT = /[_$a-zA-Z0-9]/;
  const STRINGISH_START = ['"', "'", "/"];
  const NOT_WHITESPACE = /[^\s]/;
  const COMBINED_SELECTOR_START = /[{(]/;
  const COMBINED_SELECTOR_END = /[})]/;
  function tokenizeString(str2) {
    const tokens = [];
    let position = 0;
    while (position < str2.length) {
      if (SYMBOL_START.exec(str2.charAt(position))) {
        var startPosition = position;
        while (SYMBOL_CONT.exec(str2.charAt(position + 1))) {
          position++;
        }
        tokens.push(str2.substr(startPosition, position - startPosition + 1));
      } else if (STRINGISH_START.indexOf(str2.charAt(position)) !== -1) {
        const startChar = str2.charAt(position);
        var startPosition = position;
        position++;
        while (position < str2.length && str2.charAt(position) !== startChar) {
          if (str2.charAt(position) === "\\") {
            position++;
          }
          position++;
        }
        tokens.push(str2.substr(startPosition, position - startPosition + 1));
      } else {
        const symbol = str2.charAt(position);
        tokens.push(symbol);
      }
      position++;
    }
    return tokens;
  }
  function isPossibleRelativeReference(token, last, paramName) {
    return SYMBOL_START.exec(token.charAt(0)) && token !== "true" && token !== "false" && token !== "this" && token !== paramName && last !== ".";
  }
  function maybeGenerateConditional(elt, tokens, paramName) {
    if (tokens[0] === "[") {
      tokens.shift();
      let bracketCount = 1;
      let conditionalSource = " return (function(" + paramName + "){ return (";
      let last = null;
      while (tokens.length > 0) {
        const token = tokens[0];
        if (token === "]") {
          bracketCount--;
          if (bracketCount === 0) {
            if (last === null) {
              conditionalSource = conditionalSource + "true";
            }
            tokens.shift();
            conditionalSource += ")})";
            try {
              const conditionFunction = maybeEval(
                elt,
                function() {
                  return Function(conditionalSource)();
                },
                function() {
                  return true;
                }
              );
              conditionFunction.source = conditionalSource;
              return conditionFunction;
            } catch (e) {
              triggerErrorEvent(getDocument().body, "htmx:syntax:error", { error: e, source: conditionalSource });
              return null;
            }
          }
        } else if (token === "[") {
          bracketCount++;
        }
        if (isPossibleRelativeReference(token, last, paramName)) {
          conditionalSource += "((" + paramName + "." + token + ") ? (" + paramName + "." + token + ") : (window." + token + "))";
        } else {
          conditionalSource = conditionalSource + token;
        }
        last = tokens.shift();
      }
    }
  }
  function consumeUntil(tokens, match) {
    let result = "";
    while (tokens.length > 0 && !match.test(tokens[0])) {
      result += tokens.shift();
    }
    return result;
  }
  function consumeCSSSelector(tokens) {
    let result;
    if (tokens.length > 0 && COMBINED_SELECTOR_START.test(tokens[0])) {
      tokens.shift();
      result = consumeUntil(tokens, COMBINED_SELECTOR_END).trim();
      tokens.shift();
    } else {
      result = consumeUntil(tokens, WHITESPACE_OR_COMMA);
    }
    return result;
  }
  const INPUT_SELECTOR = "input, textarea, select";
  function parseAndCacheTrigger(elt, explicitTrigger, cache) {
    const triggerSpecs = [];
    const tokens = tokenizeString(explicitTrigger);
    do {
      consumeUntil(tokens, NOT_WHITESPACE);
      const initialLength = tokens.length;
      const trigger = consumeUntil(tokens, /[,\[\s]/);
      if (trigger !== "") {
        if (trigger === "every") {
          const every = { trigger: "every" };
          consumeUntil(tokens, NOT_WHITESPACE);
          every.pollInterval = parseInterval(consumeUntil(tokens, /[,\[\s]/));
          consumeUntil(tokens, NOT_WHITESPACE);
          var eventFilter = maybeGenerateConditional(elt, tokens, "event");
          if (eventFilter) {
            every.eventFilter = eventFilter;
          }
          triggerSpecs.push(every);
        } else {
          const triggerSpec = { trigger };
          var eventFilter = maybeGenerateConditional(elt, tokens, "event");
          if (eventFilter) {
            triggerSpec.eventFilter = eventFilter;
          }
          while (tokens.length > 0 && tokens[0] !== ",") {
            consumeUntil(tokens, NOT_WHITESPACE);
            const token = tokens.shift();
            if (token === "changed") {
              triggerSpec.changed = true;
            } else if (token === "once") {
              triggerSpec.once = true;
            } else if (token === "consume") {
              triggerSpec.consume = true;
            } else if (token === "delay" && tokens[0] === ":") {
              tokens.shift();
              triggerSpec.delay = parseInterval(consumeUntil(tokens, WHITESPACE_OR_COMMA));
            } else if (token === "from" && tokens[0] === ":") {
              tokens.shift();
              if (COMBINED_SELECTOR_START.test(tokens[0])) {
                var from_arg = consumeCSSSelector(tokens);
              } else {
                var from_arg = consumeUntil(tokens, WHITESPACE_OR_COMMA);
                if (from_arg === "closest" || from_arg === "find" || from_arg === "next" || from_arg === "previous") {
                  tokens.shift();
                  const selector = consumeCSSSelector(tokens);
                  if (selector.length > 0) {
                    from_arg += " " + selector;
                  }
                }
              }
              triggerSpec.from = from_arg;
            } else if (token === "target" && tokens[0] === ":") {
              tokens.shift();
              triggerSpec.target = consumeCSSSelector(tokens);
            } else if (token === "throttle" && tokens[0] === ":") {
              tokens.shift();
              triggerSpec.throttle = parseInterval(consumeUntil(tokens, WHITESPACE_OR_COMMA));
            } else if (token === "queue" && tokens[0] === ":") {
              tokens.shift();
              triggerSpec.queue = consumeUntil(tokens, WHITESPACE_OR_COMMA);
            } else if (token === "root" && tokens[0] === ":") {
              tokens.shift();
              triggerSpec[token] = consumeCSSSelector(tokens);
            } else if (token === "threshold" && tokens[0] === ":") {
              tokens.shift();
              triggerSpec[token] = consumeUntil(tokens, WHITESPACE_OR_COMMA);
            } else {
              triggerErrorEvent(elt, "htmx:syntax:error", { token: tokens.shift() });
            }
          }
          triggerSpecs.push(triggerSpec);
        }
      }
      if (tokens.length === initialLength) {
        triggerErrorEvent(elt, "htmx:syntax:error", { token: tokens.shift() });
      }
      consumeUntil(tokens, NOT_WHITESPACE);
    } while (tokens[0] === "," && tokens.shift());
    if (cache) {
      cache[explicitTrigger] = triggerSpecs;
    }
    return triggerSpecs;
  }
  function getTriggerSpecs(elt) {
    const explicitTrigger = getAttributeValue(elt, "hx-trigger");
    let triggerSpecs = [];
    if (explicitTrigger) {
      const cache = htmx.config.triggerSpecsCache;
      triggerSpecs = cache && cache[explicitTrigger] || parseAndCacheTrigger(elt, explicitTrigger, cache);
    }
    if (triggerSpecs.length > 0) {
      return triggerSpecs;
    } else if (matches(elt, "form")) {
      return [{ trigger: "submit" }];
    } else if (matches(elt, 'input[type="button"], input[type="submit"]')) {
      return [{ trigger: "click" }];
    } else if (matches(elt, INPUT_SELECTOR)) {
      return [{ trigger: "change" }];
    } else {
      return [{ trigger: "click" }];
    }
  }
  function cancelPolling(elt) {
    getInternalData(elt).cancelled = true;
  }
  function processPolling(elt, handler, spec) {
    const nodeData = getInternalData(elt);
    nodeData.timeout = getWindow().setTimeout(function() {
      if (bodyContains(elt) && nodeData.cancelled !== true) {
        if (!maybeFilterEvent(spec, elt, makeEvent("hx:poll:trigger", {
          triggerSpec: spec,
          target: elt
        }))) {
          handler(elt);
        }
        processPolling(elt, handler, spec);
      }
    }, spec.pollInterval);
  }
  function isLocalLink(elt) {
    return location.hostname === elt.hostname && getRawAttribute(elt, "href") && getRawAttribute(elt, "href").indexOf("#") !== 0;
  }
  function eltIsDisabled(elt) {
    return closest(elt, htmx.config.disableSelector);
  }
  function boostElement(elt, nodeData, triggerSpecs) {
    if (elt instanceof HTMLAnchorElement && isLocalLink(elt) && (elt.target === "" || elt.target === "_self") || elt.tagName === "FORM" && String(getRawAttribute(elt, "method")).toLowerCase() !== "dialog") {
      nodeData.boosted = true;
      let verb, path;
      if (elt.tagName === "A") {
        verb = "get";
        path = getRawAttribute(elt, "href");
      } else {
        const rawAttribute = getRawAttribute(elt, "method");
        verb = rawAttribute ? rawAttribute.toLowerCase() : "get";
        path = getRawAttribute(elt, "action");
      }
      triggerSpecs.forEach(function(triggerSpec) {
        addEventListener(elt, function(node, evt) {
          const elt2 = asElement(node);
          if (eltIsDisabled(elt2)) {
            cleanUpElement(elt2);
            return;
          }
          issueAjaxRequest(verb, path, elt2, evt);
        }, nodeData, triggerSpec, true);
      });
    }
  }
  function shouldCancel(evt, node) {
    const elt = asElement(node);
    if (!elt) {
      return false;
    }
    if (evt.type === "submit" || evt.type === "click") {
      if (elt.tagName === "FORM") {
        return true;
      }
      if (matches(elt, 'input[type="submit"], button') && closest(elt, "form") !== null) {
        return true;
      }
      if (elt instanceof HTMLAnchorElement && elt.href && (elt.getAttribute("href") === "#" || elt.getAttribute("href").indexOf("#") !== 0)) {
        return true;
      }
    }
    return false;
  }
  function ignoreBoostedAnchorCtrlClick(elt, evt) {
    return getInternalData(elt).boosted && elt instanceof HTMLAnchorElement && evt.type === "click" && // @ts-ignore this will resolve to undefined for events that don't define those properties, which is fine
    (evt.ctrlKey || evt.metaKey);
  }
  function maybeFilterEvent(triggerSpec, elt, evt) {
    const eventFilter = triggerSpec.eventFilter;
    if (eventFilter) {
      try {
        return eventFilter.call(elt, evt) !== true;
      } catch (e) {
        const source = eventFilter.source;
        triggerErrorEvent(getDocument().body, "htmx:eventFilter:error", { error: e, source });
        return true;
      }
    }
    return false;
  }
  function addEventListener(elt, handler, nodeData, triggerSpec, explicitCancel) {
    const elementData = getInternalData(elt);
    let eltsToListenOn;
    if (triggerSpec.from) {
      eltsToListenOn = querySelectorAllExt(elt, triggerSpec.from);
    } else {
      eltsToListenOn = [elt];
    }
    if (triggerSpec.changed) {
      eltsToListenOn.forEach(function(eltToListenOn) {
        const eltToListenOnData = getInternalData(eltToListenOn);
        eltToListenOnData.lastValue = eltToListenOn.value;
      });
    }
    forEach(eltsToListenOn, function(eltToListenOn) {
      const eventListener = function(evt) {
        if (!bodyContains(elt)) {
          eltToListenOn.removeEventListener(triggerSpec.trigger, eventListener);
          return;
        }
        if (ignoreBoostedAnchorCtrlClick(elt, evt)) {
          return;
        }
        if (explicitCancel || shouldCancel(evt, elt)) {
          evt.preventDefault();
        }
        if (maybeFilterEvent(triggerSpec, elt, evt)) {
          return;
        }
        const eventData = getInternalData(evt);
        eventData.triggerSpec = triggerSpec;
        if (eventData.handledFor == null) {
          eventData.handledFor = [];
        }
        if (eventData.handledFor.indexOf(elt) < 0) {
          eventData.handledFor.push(elt);
          if (triggerSpec.consume) {
            evt.stopPropagation();
          }
          if (triggerSpec.target && evt.target) {
            if (!matches(asElement(evt.target), triggerSpec.target)) {
              return;
            }
          }
          if (triggerSpec.once) {
            if (elementData.triggeredOnce) {
              return;
            } else {
              elementData.triggeredOnce = true;
            }
          }
          if (triggerSpec.changed) {
            const eltToListenOnData = getInternalData(eltToListenOn);
            const value = eltToListenOn.value;
            if (eltToListenOnData.lastValue === value) {
              return;
            }
            eltToListenOnData.lastValue = value;
          }
          if (elementData.delayed) {
            clearTimeout(elementData.delayed);
          }
          if (elementData.throttle) {
            return;
          }
          if (triggerSpec.throttle > 0) {
            if (!elementData.throttle) {
              triggerEvent(elt, "htmx:trigger");
              handler(elt, evt);
              elementData.throttle = getWindow().setTimeout(function() {
                elementData.throttle = null;
              }, triggerSpec.throttle);
            }
          } else if (triggerSpec.delay > 0) {
            elementData.delayed = getWindow().setTimeout(function() {
              triggerEvent(elt, "htmx:trigger");
              handler(elt, evt);
            }, triggerSpec.delay);
          } else {
            triggerEvent(elt, "htmx:trigger");
            handler(elt, evt);
          }
        }
      };
      if (nodeData.listenerInfos == null) {
        nodeData.listenerInfos = [];
      }
      nodeData.listenerInfos.push({
        trigger: triggerSpec.trigger,
        listener: eventListener,
        on: eltToListenOn
      });
      eltToListenOn.addEventListener(triggerSpec.trigger, eventListener);
    });
  }
  let windowIsScrolling = false;
  let scrollHandler = null;
  function initScrollHandler() {
    if (!scrollHandler) {
      scrollHandler = function() {
        windowIsScrolling = true;
      };
      window.addEventListener("scroll", scrollHandler);
      setInterval(function() {
        if (windowIsScrolling) {
          windowIsScrolling = false;
          forEach(getDocument().querySelectorAll("[hx-trigger*='revealed'],[data-hx-trigger*='revealed']"), function(elt) {
            maybeReveal(elt);
          });
        }
      }, 200);
    }
  }
  function maybeReveal(elt) {
    if (!hasAttribute(elt, "data-hx-revealed") && isScrolledIntoView(elt)) {
      elt.setAttribute("data-hx-revealed", "true");
      const nodeData = getInternalData(elt);
      if (nodeData.initHash) {
        triggerEvent(elt, "revealed");
      } else {
        elt.addEventListener("htmx:afterProcessNode", function() {
          triggerEvent(elt, "revealed");
        }, { once: true });
      }
    }
  }
  function loadImmediately(elt, handler, nodeData, delay) {
    const load = function() {
      if (!nodeData.loaded) {
        nodeData.loaded = true;
        handler(elt);
      }
    };
    if (delay > 0) {
      getWindow().setTimeout(load, delay);
    } else {
      load();
    }
  }
  function processVerbs(elt, nodeData, triggerSpecs) {
    let explicitAction = false;
    forEach(VERBS, function(verb) {
      if (hasAttribute(elt, "hx-" + verb)) {
        const path = getAttributeValue(elt, "hx-" + verb);
        explicitAction = true;
        nodeData.path = path;
        nodeData.verb = verb;
        triggerSpecs.forEach(function(triggerSpec) {
          addTriggerHandler(elt, triggerSpec, nodeData, function(node, evt) {
            const elt2 = asElement(node);
            if (closest(elt2, htmx.config.disableSelector)) {
              cleanUpElement(elt2);
              return;
            }
            issueAjaxRequest(verb, path, elt2, evt);
          });
        });
      }
    });
    return explicitAction;
  }
  function addTriggerHandler(elt, triggerSpec, nodeData, handler) {
    if (triggerSpec.trigger === "revealed") {
      initScrollHandler();
      addEventListener(elt, handler, nodeData, triggerSpec);
      maybeReveal(asElement(elt));
    } else if (triggerSpec.trigger === "intersect") {
      const observerOptions = {};
      if (triggerSpec.root) {
        observerOptions.root = querySelectorExt(elt, triggerSpec.root);
      }
      if (triggerSpec.threshold) {
        observerOptions.threshold = parseFloat(triggerSpec.threshold);
      }
      const observer = new IntersectionObserver(function(entries) {
        for (let i = 0; i < entries.length; i++) {
          const entry = entries[i];
          if (entry.isIntersecting) {
            triggerEvent(elt, "intersect");
            break;
          }
        }
      }, observerOptions);
      observer.observe(asElement(elt));
      addEventListener(asElement(elt), handler, nodeData, triggerSpec);
    } else if (triggerSpec.trigger === "load") {
      if (!maybeFilterEvent(triggerSpec, elt, makeEvent("load", { elt }))) {
        loadImmediately(asElement(elt), handler, nodeData, triggerSpec.delay);
      }
    } else if (triggerSpec.pollInterval > 0) {
      nodeData.polling = true;
      processPolling(asElement(elt), handler, triggerSpec);
    } else {
      addEventListener(elt, handler, nodeData, triggerSpec);
    }
  }
  function shouldProcessHxOn(node) {
    const elt = asElement(node);
    if (!elt) {
      return false;
    }
    const attributes = elt.attributes;
    for (let j = 0; j < attributes.length; j++) {
      const attrName = attributes[j].name;
      if (startsWith(attrName, "hx-on:") || startsWith(attrName, "data-hx-on:") || startsWith(attrName, "hx-on-") || startsWith(attrName, "data-hx-on-")) {
        return true;
      }
    }
    return false;
  }
  const HX_ON_QUERY = new XPathEvaluator().createExpression('.//*[@*[ starts-with(name(), "hx-on:") or starts-with(name(), "data-hx-on:") or starts-with(name(), "hx-on-") or starts-with(name(), "data-hx-on-") ]]');
  function processHXOnRoot(elt, elements) {
    if (shouldProcessHxOn(elt)) {
      elements.push(asElement(elt));
    }
    const iter = HX_ON_QUERY.evaluate(elt);
    let node = null;
    while (node = iter.iterateNext()) elements.push(asElement(node));
  }
  function findHxOnWildcardElements(elt) {
    const elements = [];
    if (elt instanceof DocumentFragment) {
      for (const child of elt.childNodes) {
        processHXOnRoot(child, elements);
      }
    } else {
      processHXOnRoot(elt, elements);
    }
    return elements;
  }
  function findElementsToProcess(elt) {
    if (elt.querySelectorAll) {
      const boostedSelector = ", [hx-boost] a, [data-hx-boost] a, a[hx-boost], a[data-hx-boost]";
      const extensionSelectors = [];
      for (const e in extensions) {
        const extension = extensions[e];
        if (extension.getSelectors) {
          var selectors = extension.getSelectors();
          if (selectors) {
            extensionSelectors.push(selectors);
          }
        }
      }
      const results = elt.querySelectorAll(VERB_SELECTOR + boostedSelector + ", form, [type='submit'], [hx-ext], [data-hx-ext], [hx-trigger], [data-hx-trigger]" + extensionSelectors.flat().map((s) => ", " + s).join(""));
      return results;
    } else {
      return [];
    }
  }
  function maybeSetLastButtonClicked(evt) {
    const elt = (
      /** @type {HTMLButtonElement|HTMLInputElement} */
      closest(asElement(evt.target), "button, input[type='submit']")
    );
    const internalData = getRelatedFormData(evt);
    if (internalData) {
      internalData.lastButtonClicked = elt;
    }
  }
  function maybeUnsetLastButtonClicked(evt) {
    const internalData = getRelatedFormData(evt);
    if (internalData) {
      internalData.lastButtonClicked = null;
    }
  }
  function getRelatedFormData(evt) {
    const elt = closest(asElement(evt.target), "button, input[type='submit']");
    if (!elt) {
      return;
    }
    const form = resolveTarget("#" + getRawAttribute(elt, "form"), elt.getRootNode()) || closest(elt, "form");
    if (!form) {
      return;
    }
    return getInternalData(form);
  }
  function initButtonTracking(elt) {
    elt.addEventListener("click", maybeSetLastButtonClicked);
    elt.addEventListener("focusin", maybeSetLastButtonClicked);
    elt.addEventListener("focusout", maybeUnsetLastButtonClicked);
  }
  function addHxOnEventHandler(elt, eventName, code) {
    const nodeData = getInternalData(elt);
    if (!Array.isArray(nodeData.onHandlers)) {
      nodeData.onHandlers = [];
    }
    let func;
    const listener = function(e) {
      maybeEval(elt, function() {
        if (eltIsDisabled(elt)) {
          return;
        }
        if (!func) {
          func = new Function("event", code);
        }
        func.call(elt, e);
      });
    };
    elt.addEventListener(eventName, listener);
    nodeData.onHandlers.push({ event: eventName, listener });
  }
  function processHxOnWildcard(elt) {
    deInitOnHandlers(elt);
    for (let i = 0; i < elt.attributes.length; i++) {
      const name = elt.attributes[i].name;
      const value = elt.attributes[i].value;
      if (startsWith(name, "hx-on") || startsWith(name, "data-hx-on")) {
        const afterOnPosition = name.indexOf("-on") + 3;
        const nextChar = name.slice(afterOnPosition, afterOnPosition + 1);
        if (nextChar === "-" || nextChar === ":") {
          let eventName = name.slice(afterOnPosition + 1);
          if (startsWith(eventName, ":")) {
            eventName = "htmx" + eventName;
          } else if (startsWith(eventName, "-")) {
            eventName = "htmx:" + eventName.slice(1);
          } else if (startsWith(eventName, "htmx-")) {
            eventName = "htmx:" + eventName.slice(5);
          }
          addHxOnEventHandler(elt, eventName, value);
        }
      }
    }
  }
  function initNode(elt) {
    if (closest(elt, htmx.config.disableSelector)) {
      cleanUpElement(elt);
      return;
    }
    const nodeData = getInternalData(elt);
    if (nodeData.initHash !== attributeHash(elt)) {
      deInitNode(elt);
      nodeData.initHash = attributeHash(elt);
      triggerEvent(elt, "htmx:beforeProcessNode");
      if (elt.value) {
        nodeData.lastValue = elt.value;
      }
      const triggerSpecs = getTriggerSpecs(elt);
      const hasExplicitHttpAction = processVerbs(elt, nodeData, triggerSpecs);
      if (!hasExplicitHttpAction) {
        if (getClosestAttributeValue(elt, "hx-boost") === "true") {
          boostElement(elt, nodeData, triggerSpecs);
        } else if (hasAttribute(elt, "hx-trigger")) {
          triggerSpecs.forEach(function(triggerSpec) {
            addTriggerHandler(elt, triggerSpec, nodeData, function() {
            });
          });
        }
      }
      if (elt.tagName === "FORM" || getRawAttribute(elt, "type") === "submit" && hasAttribute(elt, "form")) {
        initButtonTracking(elt);
      }
      triggerEvent(elt, "htmx:afterProcessNode");
    }
  }
  function processNode(elt) {
    elt = resolveTarget(elt);
    if (closest(elt, htmx.config.disableSelector)) {
      cleanUpElement(elt);
      return;
    }
    initNode(elt);
    forEach(findElementsToProcess(elt), function(child) {
      initNode(child);
    });
    forEach(findHxOnWildcardElements(elt), processHxOnWildcard);
  }
  function kebabEventName(str2) {
    return str2.replace(/([a-z0-9])([A-Z])/g, "$1-$2").toLowerCase();
  }
  function makeEvent(eventName, detail) {
    let evt;
    if (window.CustomEvent && typeof window.CustomEvent === "function") {
      evt = new CustomEvent(eventName, { bubbles: true, cancelable: true, composed: true, detail });
    } else {
      evt = getDocument().createEvent("CustomEvent");
      evt.initCustomEvent(eventName, true, true, detail);
    }
    return evt;
  }
  function triggerErrorEvent(elt, eventName, detail) {
    triggerEvent(elt, eventName, mergeObjects({ error: eventName }, detail));
  }
  function ignoreEventForLogging(eventName) {
    return eventName === "htmx:afterProcessNode";
  }
  function withExtensions(elt, toDo) {
    forEach(getExtensions(elt), function(extension) {
      try {
        toDo(extension);
      } catch (e) {
        logError(e);
      }
    });
  }
  function logError(msg) {
    if (console.error) {
      console.error(msg);
    } else if (console.log) {
      console.log("ERROR: ", msg);
    }
  }
  function triggerEvent(elt, eventName, detail) {
    elt = resolveTarget(elt);
    if (detail == null) {
      detail = {};
    }
    detail.elt = elt;
    const event = makeEvent(eventName, detail);
    if (htmx.logger && !ignoreEventForLogging(eventName)) {
      htmx.logger(elt, eventName, detail);
    }
    if (detail.error) {
      logError(detail.error);
      triggerEvent(elt, "htmx:error", { errorInfo: detail });
    }
    let eventResult = elt.dispatchEvent(event);
    const kebabName = kebabEventName(eventName);
    if (eventResult && kebabName !== eventName) {
      const kebabedEvent = makeEvent(kebabName, event.detail);
      eventResult = eventResult && elt.dispatchEvent(kebabedEvent);
    }
    withExtensions(asElement(elt), function(extension) {
      eventResult = eventResult && (extension.onEvent(eventName, event) !== false && !event.defaultPrevented);
    });
    return eventResult;
  }
  let currentPathForHistory = location.pathname + location.search;
  function getHistoryElement() {
    const historyElt = getDocument().querySelector("[hx-history-elt],[data-hx-history-elt]");
    return historyElt || getDocument().body;
  }
  function saveToHistoryCache(url, rootElt) {
    if (!canAccessLocalStorage()) {
      return;
    }
    const innerHTML = cleanInnerHtmlForHistory(rootElt);
    const title = getDocument().title;
    const scroll = window.scrollY;
    if (htmx.config.historyCacheSize <= 0) {
      localStorage.removeItem("htmx-history-cache");
      return;
    }
    url = normalizePath(url);
    const historyCache = parseJSON(localStorage.getItem("htmx-history-cache")) || [];
    for (let i = 0; i < historyCache.length; i++) {
      if (historyCache[i].url === url) {
        historyCache.splice(i, 1);
        break;
      }
    }
    const newHistoryItem = { url, content: innerHTML, title, scroll };
    triggerEvent(getDocument().body, "htmx:historyItemCreated", { item: newHistoryItem, cache: historyCache });
    historyCache.push(newHistoryItem);
    while (historyCache.length > htmx.config.historyCacheSize) {
      historyCache.shift();
    }
    while (historyCache.length > 0) {
      try {
        localStorage.setItem("htmx-history-cache", JSON.stringify(historyCache));
        break;
      } catch (e) {
        triggerErrorEvent(getDocument().body, "htmx:historyCacheError", { cause: e, cache: historyCache });
        historyCache.shift();
      }
    }
  }
  function getCachedHistory(url) {
    if (!canAccessLocalStorage()) {
      return null;
    }
    url = normalizePath(url);
    const historyCache = parseJSON(localStorage.getItem("htmx-history-cache")) || [];
    for (let i = 0; i < historyCache.length; i++) {
      if (historyCache[i].url === url) {
        return historyCache[i];
      }
    }
    return null;
  }
  function cleanInnerHtmlForHistory(elt) {
    const className = htmx.config.requestClass;
    const clone = (
      /** @type Element */
      elt.cloneNode(true)
    );
    forEach(findAll(clone, "." + className), function(child) {
      removeClassFromElement(child, className);
    });
    forEach(findAll(clone, "[data-disabled-by-htmx]"), function(child) {
      child.removeAttribute("disabled");
    });
    return clone.innerHTML;
  }
  function saveCurrentPageToHistory() {
    const elt = getHistoryElement();
    const path = currentPathForHistory || location.pathname + location.search;
    let disableHistoryCache;
    try {
      disableHistoryCache = getDocument().querySelector('[hx-history="false" i],[data-hx-history="false" i]');
    } catch (e) {
      disableHistoryCache = getDocument().querySelector('[hx-history="false"],[data-hx-history="false"]');
    }
    if (!disableHistoryCache) {
      triggerEvent(getDocument().body, "htmx:beforeHistorySave", { path, historyElt: elt });
      saveToHistoryCache(path, elt);
    }
    if (htmx.config.historyEnabled) history.replaceState({ htmx: true }, getDocument().title, window.location.href);
  }
  function pushUrlIntoHistory(path) {
    if (htmx.config.getCacheBusterParam) {
      path = path.replace(/org\.htmx\.cache-buster=[^&]*&?/, "");
      if (endsWith(path, "&") || endsWith(path, "?")) {
        path = path.slice(0, -1);
      }
    }
    if (htmx.config.historyEnabled) {
      history.pushState({ htmx: true }, "", path);
    }
    currentPathForHistory = path;
  }
  function replaceUrlInHistory(path) {
    if (htmx.config.historyEnabled) history.replaceState({ htmx: true }, "", path);
    currentPathForHistory = path;
  }
  function settleImmediately(tasks) {
    forEach(tasks, function(task) {
      task.call(void 0);
    });
  }
  function loadHistoryFromServer(path) {
    const request = new XMLHttpRequest();
    const details = { path, xhr: request };
    triggerEvent(getDocument().body, "htmx:historyCacheMiss", details);
    request.open("GET", path, true);
    request.setRequestHeader("HX-Request", "true");
    request.setRequestHeader("HX-History-Restore-Request", "true");
    request.setRequestHeader("HX-Current-URL", getDocument().location.href);
    request.onload = function() {
      if (this.status >= 200 && this.status < 400) {
        triggerEvent(getDocument().body, "htmx:historyCacheMissLoad", details);
        const fragment = makeFragment(this.response);
        const content = fragment.querySelector("[hx-history-elt],[data-hx-history-elt]") || fragment;
        const historyElement = getHistoryElement();
        const settleInfo = makeSettleInfo(historyElement);
        handleTitle(fragment.title);
        swapInnerHTML(historyElement, content, settleInfo);
        settleImmediately(settleInfo.tasks);
        currentPathForHistory = path;
        triggerEvent(getDocument().body, "htmx:historyRestore", { path, cacheMiss: true, serverResponse: this.response });
      } else {
        triggerErrorEvent(getDocument().body, "htmx:historyCacheMissLoadError", details);
      }
    };
    request.send();
  }
  function restoreHistory(path) {
    saveCurrentPageToHistory();
    path = path || location.pathname + location.search;
    const cached = getCachedHistory(path);
    if (cached) {
      const fragment = makeFragment(cached.content);
      const historyElement = getHistoryElement();
      const settleInfo = makeSettleInfo(historyElement);
      handleTitle(fragment.title);
      swapInnerHTML(historyElement, fragment, settleInfo);
      settleImmediately(settleInfo.tasks);
      getWindow().setTimeout(function() {
        window.scrollTo(0, cached.scroll);
      }, 0);
      currentPathForHistory = path;
      triggerEvent(getDocument().body, "htmx:historyRestore", { path, item: cached });
    } else {
      if (htmx.config.refreshOnHistoryMiss) {
        window.location.reload(true);
      } else {
        loadHistoryFromServer(path);
      }
    }
  }
  function addRequestIndicatorClasses(elt) {
    let indicators = (
      /** @type Element[] */
      findAttributeTargets(elt, "hx-indicator")
    );
    if (indicators == null) {
      indicators = [elt];
    }
    forEach(indicators, function(ic) {
      const internalData = getInternalData(ic);
      internalData.requestCount = (internalData.requestCount || 0) + 1;
      ic.classList.add.call(ic.classList, htmx.config.requestClass);
    });
    return indicators;
  }
  function disableElements(elt) {
    let disabledElts = (
      /** @type Element[] */
      findAttributeTargets(elt, "hx-disabled-elt")
    );
    if (disabledElts == null) {
      disabledElts = [];
    }
    forEach(disabledElts, function(disabledElement) {
      const internalData = getInternalData(disabledElement);
      internalData.requestCount = (internalData.requestCount || 0) + 1;
      disabledElement.setAttribute("disabled", "");
      disabledElement.setAttribute("data-disabled-by-htmx", "");
    });
    return disabledElts;
  }
  function removeRequestIndicators(indicators, disabled) {
    forEach(indicators, function(ic) {
      const internalData = getInternalData(ic);
      internalData.requestCount = (internalData.requestCount || 0) - 1;
      if (internalData.requestCount === 0) {
        ic.classList.remove.call(ic.classList, htmx.config.requestClass);
      }
    });
    forEach(disabled, function(disabledElement) {
      const internalData = getInternalData(disabledElement);
      internalData.requestCount = (internalData.requestCount || 0) - 1;
      if (internalData.requestCount === 0) {
        disabledElement.removeAttribute("disabled");
        disabledElement.removeAttribute("data-disabled-by-htmx");
      }
    });
  }
  function haveSeenNode(processed, elt) {
    for (let i = 0; i < processed.length; i++) {
      const node = processed[i];
      if (node.isSameNode(elt)) {
        return true;
      }
    }
    return false;
  }
  function shouldInclude(element) {
    const elt = (
      /** @type {HTMLInputElement} */
      element
    );
    if (elt.name === "" || elt.name == null || elt.disabled || closest(elt, "fieldset[disabled]")) {
      return false;
    }
    if (elt.type === "button" || elt.type === "submit" || elt.tagName === "image" || elt.tagName === "reset" || elt.tagName === "file") {
      return false;
    }
    if (elt.type === "checkbox" || elt.type === "radio") {
      return elt.checked;
    }
    return true;
  }
  function addValueToFormData(name, value, formData) {
    if (name != null && value != null) {
      if (Array.isArray(value)) {
        value.forEach(function(v) {
          formData.append(name, v);
        });
      } else {
        formData.append(name, value);
      }
    }
  }
  function removeValueFromFormData(name, value, formData) {
    if (name != null && value != null) {
      let values = formData.getAll(name);
      if (Array.isArray(value)) {
        values = values.filter((v) => value.indexOf(v) < 0);
      } else {
        values = values.filter((v) => v !== value);
      }
      formData.delete(name);
      forEach(values, (v) => formData.append(name, v));
    }
  }
  function processInputValue(processed, formData, errors, elt, validate) {
    if (elt == null || haveSeenNode(processed, elt)) {
      return;
    } else {
      processed.push(elt);
    }
    if (shouldInclude(elt)) {
      const name = getRawAttribute(elt, "name");
      let value = elt.value;
      if (elt instanceof HTMLSelectElement && elt.multiple) {
        value = toArray(elt.querySelectorAll("option:checked")).map(function(e) {
          return (
            /** @type HTMLOptionElement */
            e.value
          );
        });
      }
      if (elt instanceof HTMLInputElement && elt.files) {
        value = toArray(elt.files);
      }
      addValueToFormData(name, value, formData);
      if (validate) {
        validateElement(elt, errors);
      }
    }
    if (elt instanceof HTMLFormElement) {
      forEach(elt.elements, function(input) {
        if (processed.indexOf(input) >= 0) {
          removeValueFromFormData(input.name, input.value, formData);
        } else {
          processed.push(input);
        }
        if (validate) {
          validateElement(input, errors);
        }
      });
      new FormData(elt).forEach(function(value, name) {
        if (value instanceof File && value.name === "") {
          return;
        }
        addValueToFormData(name, value, formData);
      });
    }
  }
  function validateElement(elt, errors) {
    const element = (
      /** @type {HTMLElement & ElementInternals} */
      elt
    );
    if (element.willValidate) {
      triggerEvent(element, "htmx:validation:validate");
      if (!element.checkValidity()) {
        errors.push({ elt: element, message: element.validationMessage, validity: element.validity });
        triggerEvent(element, "htmx:validation:failed", { message: element.validationMessage, validity: element.validity });
      }
    }
  }
  function overrideFormData(receiver, donor) {
    for (const key of donor.keys()) {
      receiver.delete(key);
    }
    donor.forEach(function(value, key) {
      receiver.append(key, value);
    });
    return receiver;
  }
  function getInputValues(elt, verb) {
    const processed = [];
    const formData = new FormData();
    const priorityFormData = new FormData();
    const errors = [];
    const internalData = getInternalData(elt);
    if (internalData.lastButtonClicked && !bodyContains(internalData.lastButtonClicked)) {
      internalData.lastButtonClicked = null;
    }
    let validate = elt instanceof HTMLFormElement && elt.noValidate !== true || getAttributeValue(elt, "hx-validate") === "true";
    if (internalData.lastButtonClicked) {
      validate = validate && internalData.lastButtonClicked.formNoValidate !== true;
    }
    if (verb !== "get") {
      processInputValue(processed, priorityFormData, errors, closest(elt, "form"), validate);
    }
    processInputValue(processed, formData, errors, elt, validate);
    if (internalData.lastButtonClicked || elt.tagName === "BUTTON" || elt.tagName === "INPUT" && getRawAttribute(elt, "type") === "submit") {
      const button = internalData.lastButtonClicked || /** @type HTMLInputElement|HTMLButtonElement */
      elt;
      const name = getRawAttribute(button, "name");
      addValueToFormData(name, button.value, priorityFormData);
    }
    const includes = findAttributeTargets(elt, "hx-include");
    forEach(includes, function(node) {
      processInputValue(processed, formData, errors, asElement(node), validate);
      if (!matches(node, "form")) {
        forEach(asParentNode(node).querySelectorAll(INPUT_SELECTOR), function(descendant) {
          processInputValue(processed, formData, errors, descendant, validate);
        });
      }
    });
    overrideFormData(formData, priorityFormData);
    return { errors, formData, values: formDataProxy(formData) };
  }
  function appendParam(returnStr, name, realValue) {
    if (returnStr !== "") {
      returnStr += "&";
    }
    if (String(realValue) === "[object Object]") {
      realValue = JSON.stringify(realValue);
    }
    const s = encodeURIComponent(realValue);
    returnStr += encodeURIComponent(name) + "=" + s;
    return returnStr;
  }
  function urlEncode(values) {
    values = formDataFromObject(values);
    let returnStr = "";
    values.forEach(function(value, key) {
      returnStr = appendParam(returnStr, key, value);
    });
    return returnStr;
  }
  function getHeaders(elt, target, prompt2) {
    const headers = {
      "HX-Request": "true",
      "HX-Trigger": getRawAttribute(elt, "id"),
      "HX-Trigger-Name": getRawAttribute(elt, "name"),
      "HX-Target": getAttributeValue(target, "id"),
      "HX-Current-URL": getDocument().location.href
    };
    getValuesForElement(elt, "hx-headers", false, headers);
    if (prompt2 !== void 0) {
      headers["HX-Prompt"] = prompt2;
    }
    if (getInternalData(elt).boosted) {
      headers["HX-Boosted"] = "true";
    }
    return headers;
  }
  function filterValues(inputValues, elt) {
    const paramsValue = getClosestAttributeValue(elt, "hx-params");
    if (paramsValue) {
      if (paramsValue === "none") {
        return new FormData();
      } else if (paramsValue === "*") {
        return inputValues;
      } else if (paramsValue.indexOf("not ") === 0) {
        forEach(paramsValue.substr(4).split(","), function(name) {
          name = name.trim();
          inputValues.delete(name);
        });
        return inputValues;
      } else {
        const newValues = new FormData();
        forEach(paramsValue.split(","), function(name) {
          name = name.trim();
          if (inputValues.has(name)) {
            inputValues.getAll(name).forEach(function(value) {
              newValues.append(name, value);
            });
          }
        });
        return newValues;
      }
    } else {
      return inputValues;
    }
  }
  function isAnchorLink(elt) {
    return !!getRawAttribute(elt, "href") && getRawAttribute(elt, "href").indexOf("#") >= 0;
  }
  function getSwapSpecification(elt, swapInfoOverride) {
    const swapInfo = swapInfoOverride || getClosestAttributeValue(elt, "hx-swap");
    const swapSpec = {
      swapStyle: getInternalData(elt).boosted ? "innerHTML" : htmx.config.defaultSwapStyle,
      swapDelay: htmx.config.defaultSwapDelay,
      settleDelay: htmx.config.defaultSettleDelay
    };
    if (htmx.config.scrollIntoViewOnBoost && getInternalData(elt).boosted && !isAnchorLink(elt)) {
      swapSpec.show = "top";
    }
    if (swapInfo) {
      const split = splitOnWhitespace(swapInfo);
      if (split.length > 0) {
        for (let i = 0; i < split.length; i++) {
          const value = split[i];
          if (value.indexOf("swap:") === 0) {
            swapSpec.swapDelay = parseInterval(value.substr(5));
          } else if (value.indexOf("settle:") === 0) {
            swapSpec.settleDelay = parseInterval(value.substr(7));
          } else if (value.indexOf("transition:") === 0) {
            swapSpec.transition = value.substr(11) === "true";
          } else if (value.indexOf("ignoreTitle:") === 0) {
            swapSpec.ignoreTitle = value.substr(12) === "true";
          } else if (value.indexOf("scroll:") === 0) {
            const scrollSpec = value.substr(7);
            var splitSpec = scrollSpec.split(":");
            const scrollVal = splitSpec.pop();
            var selectorVal = splitSpec.length > 0 ? splitSpec.join(":") : null;
            swapSpec.scroll = scrollVal;
            swapSpec.scrollTarget = selectorVal;
          } else if (value.indexOf("show:") === 0) {
            const showSpec = value.substr(5);
            var splitSpec = showSpec.split(":");
            const showVal = splitSpec.pop();
            var selectorVal = splitSpec.length > 0 ? splitSpec.join(":") : null;
            swapSpec.show = showVal;
            swapSpec.showTarget = selectorVal;
          } else if (value.indexOf("focus-scroll:") === 0) {
            const focusScrollVal = value.substr("focus-scroll:".length);
            swapSpec.focusScroll = focusScrollVal == "true";
          } else if (i == 0) {
            swapSpec.swapStyle = value;
          } else {
            logError("Unknown modifier in hx-swap: " + value);
          }
        }
      }
    }
    return swapSpec;
  }
  function usesFormData(elt) {
    return getClosestAttributeValue(elt, "hx-encoding") === "multipart/form-data" || matches(elt, "form") && getRawAttribute(elt, "enctype") === "multipart/form-data";
  }
  function encodeParamsForBody(xhr, elt, filteredParameters) {
    let encodedParameters = null;
    withExtensions(elt, function(extension) {
      if (encodedParameters == null) {
        encodedParameters = extension.encodeParameters(xhr, filteredParameters, elt);
      }
    });
    if (encodedParameters != null) {
      return encodedParameters;
    } else {
      if (usesFormData(elt)) {
        return overrideFormData(new FormData(), formDataFromObject(filteredParameters));
      } else {
        return urlEncode(filteredParameters);
      }
    }
  }
  function makeSettleInfo(target) {
    return { tasks: [], elts: [target] };
  }
  function updateScrollState(content, swapSpec) {
    const first = content[0];
    const last = content[content.length - 1];
    if (swapSpec.scroll) {
      var target = null;
      if (swapSpec.scrollTarget) {
        target = asElement(querySelectorExt(first, swapSpec.scrollTarget));
      }
      if (swapSpec.scroll === "top" && (first || target)) {
        target = target || first;
        target.scrollTop = 0;
      }
      if (swapSpec.scroll === "bottom" && (last || target)) {
        target = target || last;
        target.scrollTop = target.scrollHeight;
      }
    }
    if (swapSpec.show) {
      var target = null;
      if (swapSpec.showTarget) {
        let targetStr = swapSpec.showTarget;
        if (swapSpec.showTarget === "window") {
          targetStr = "body";
        }
        target = asElement(querySelectorExt(first, targetStr));
      }
      if (swapSpec.show === "top" && (first || target)) {
        target = target || first;
        target.scrollIntoView({ block: "start", behavior: htmx.config.scrollBehavior });
      }
      if (swapSpec.show === "bottom" && (last || target)) {
        target = target || last;
        target.scrollIntoView({ block: "end", behavior: htmx.config.scrollBehavior });
      }
    }
  }
  function getValuesForElement(elt, attr, evalAsDefault, values) {
    if (values == null) {
      values = {};
    }
    if (elt == null) {
      return values;
    }
    const attributeValue = getAttributeValue(elt, attr);
    if (attributeValue) {
      let str2 = attributeValue.trim();
      let evaluateValue = evalAsDefault;
      if (str2 === "unset") {
        return null;
      }
      if (str2.indexOf("javascript:") === 0) {
        str2 = str2.substr(11);
        evaluateValue = true;
      } else if (str2.indexOf("js:") === 0) {
        str2 = str2.substr(3);
        evaluateValue = true;
      }
      if (str2.indexOf("{") !== 0) {
        str2 = "{" + str2 + "}";
      }
      let varsValues;
      if (evaluateValue) {
        varsValues = maybeEval(elt, function() {
          return Function("return (" + str2 + ")")();
        }, {});
      } else {
        varsValues = parseJSON(str2);
      }
      for (const key in varsValues) {
        if (varsValues.hasOwnProperty(key)) {
          if (values[key] == null) {
            values[key] = varsValues[key];
          }
        }
      }
    }
    return getValuesForElement(asElement(parentElt(elt)), attr, evalAsDefault, values);
  }
  function maybeEval(elt, toEval, defaultVal) {
    if (htmx.config.allowEval) {
      return toEval();
    } else {
      triggerErrorEvent(elt, "htmx:evalDisallowedError");
      return defaultVal;
    }
  }
  function getHXVarsForElement(elt, expressionVars) {
    return getValuesForElement(elt, "hx-vars", true, expressionVars);
  }
  function getHXValsForElement(elt, expressionVars) {
    return getValuesForElement(elt, "hx-vals", false, expressionVars);
  }
  function getExpressionVars(elt) {
    return mergeObjects(getHXVarsForElement(elt), getHXValsForElement(elt));
  }
  function safelySetHeaderValue(xhr, header, headerValue) {
    if (headerValue !== null) {
      try {
        xhr.setRequestHeader(header, headerValue);
      } catch (e) {
        xhr.setRequestHeader(header, encodeURIComponent(headerValue));
        xhr.setRequestHeader(header + "-URI-AutoEncoded", "true");
      }
    }
  }
  function getPathFromResponse(xhr) {
    if (xhr.responseURL && typeof URL !== "undefined") {
      try {
        const url = new URL(xhr.responseURL);
        return url.pathname + url.search;
      } catch (e) {
        triggerErrorEvent(getDocument().body, "htmx:badResponseUrl", { url: xhr.responseURL });
      }
    }
  }
  function hasHeader(xhr, regexp) {
    return regexp.test(xhr.getAllResponseHeaders());
  }
  function ajaxHelper(verb, path, context) {
    verb = /** @type HttpVerb */
    verb.toLowerCase();
    if (context) {
      if (context instanceof Element || typeof context === "string") {
        return issueAjaxRequest(verb, path, null, null, {
          targetOverride: resolveTarget(context),
          returnPromise: true
        });
      } else {
        return issueAjaxRequest(
          verb,
          path,
          resolveTarget(context.source),
          context.event,
          {
            handler: context.handler,
            headers: context.headers,
            values: context.values,
            targetOverride: resolveTarget(context.target),
            swapOverride: context.swap,
            select: context.select,
            returnPromise: true
          }
        );
      }
    } else {
      return issueAjaxRequest(verb, path, null, null, {
        returnPromise: true
      });
    }
  }
  function hierarchyForElt(elt) {
    const arr = [];
    while (elt) {
      arr.push(elt);
      elt = elt.parentElement;
    }
    return arr;
  }
  function verifyPath(elt, path, requestConfig) {
    let sameHost;
    let url;
    if (typeof URL === "function") {
      url = new URL(path, document.location.href);
      const origin = document.location.origin;
      sameHost = origin === url.origin;
    } else {
      url = path;
      sameHost = startsWith(path, document.location.origin);
    }
    if (htmx.config.selfRequestsOnly) {
      if (!sameHost) {
        return false;
      }
    }
    return triggerEvent(elt, "htmx:validateUrl", mergeObjects({ url, sameHost }, requestConfig));
  }
  function formDataFromObject(obj) {
    if (obj instanceof FormData) return obj;
    const formData = new FormData();
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        if (typeof obj[key].forEach === "function") {
          obj[key].forEach(function(v) {
            formData.append(key, v);
          });
        } else if (typeof obj[key] === "object" && !(obj[key] instanceof Blob)) {
          formData.append(key, JSON.stringify(obj[key]));
        } else {
          formData.append(key, obj[key]);
        }
      }
    }
    return formData;
  }
  function formDataArrayProxy(formData, name, array) {
    return new Proxy(array, {
      get: function(target, key) {
        if (typeof key === "number") return target[key];
        if (key === "length") return target.length;
        if (key === "push") {
          return function(value) {
            target.push(value);
            formData.append(name, value);
          };
        }
        if (typeof target[key] === "function") {
          return function() {
            target[key].apply(target, arguments);
            formData.delete(name);
            target.forEach(function(v) {
              formData.append(name, v);
            });
          };
        }
        if (target[key] && target[key].length === 1) {
          return target[key][0];
        } else {
          return target[key];
        }
      },
      set: function(target, index, value) {
        target[index] = value;
        formData.delete(name);
        target.forEach(function(v) {
          formData.append(name, v);
        });
        return true;
      }
    });
  }
  function formDataProxy(formData) {
    return new Proxy(formData, {
      get: function(target, name) {
        if (typeof name === "symbol") {
          return Reflect.get(target, name);
        }
        if (name === "toJSON") {
          return () => Object.fromEntries(formData);
        }
        if (name in target) {
          if (typeof target[name] === "function") {
            return function() {
              return formData[name].apply(formData, arguments);
            };
          } else {
            return target[name];
          }
        }
        const array = formData.getAll(name);
        if (array.length === 0) {
          return void 0;
        } else if (array.length === 1) {
          return array[0];
        } else {
          return formDataArrayProxy(target, name, array);
        }
      },
      set: function(target, name, value) {
        if (typeof name !== "string") {
          return false;
        }
        target.delete(name);
        if (typeof value.forEach === "function") {
          value.forEach(function(v) {
            target.append(name, v);
          });
        } else if (typeof value === "object" && !(value instanceof Blob)) {
          target.append(name, JSON.stringify(value));
        } else {
          target.append(name, value);
        }
        return true;
      },
      deleteProperty: function(target, name) {
        if (typeof name === "string") {
          target.delete(name);
        }
        return true;
      },
      // Support Object.assign call from proxy
      ownKeys: function(target) {
        return Reflect.ownKeys(Object.fromEntries(target));
      },
      getOwnPropertyDescriptor: function(target, prop) {
        return Reflect.getOwnPropertyDescriptor(Object.fromEntries(target), prop);
      }
    });
  }
  function issueAjaxRequest(verb, path, elt, event, etc, confirmed) {
    let resolve = null;
    let reject = null;
    etc = etc != null ? etc : {};
    if (etc.returnPromise && typeof Promise !== "undefined") {
      var promise = new Promise(function(_resolve, _reject) {
        resolve = _resolve;
        reject = _reject;
      });
    }
    if (elt == null) {
      elt = getDocument().body;
    }
    const responseHandler = etc.handler || handleAjaxResponse;
    const select = etc.select || null;
    if (!bodyContains(elt)) {
      maybeCall(resolve);
      return promise;
    }
    const target = etc.targetOverride || asElement(getTarget(elt));
    if (target == null || target == DUMMY_ELT) {
      triggerErrorEvent(elt, "htmx:targetError", { target: getAttributeValue(elt, "hx-target") });
      maybeCall(reject);
      return promise;
    }
    let eltData = getInternalData(elt);
    const submitter = eltData.lastButtonClicked;
    if (submitter) {
      const buttonPath = getRawAttribute(submitter, "formaction");
      if (buttonPath != null) {
        path = buttonPath;
      }
      const buttonVerb = getRawAttribute(submitter, "formmethod");
      if (buttonVerb != null) {
        if (buttonVerb.toLowerCase() !== "dialog") {
          verb = /** @type HttpVerb */
          buttonVerb;
        }
      }
    }
    const confirmQuestion = getClosestAttributeValue(elt, "hx-confirm");
    if (confirmed === void 0) {
      const issueRequest = function(skipConfirmation) {
        return issueAjaxRequest(verb, path, elt, event, etc, !!skipConfirmation);
      };
      const confirmDetails = { target, elt, path, verb, triggeringEvent: event, etc, issueRequest, question: confirmQuestion };
      if (triggerEvent(elt, "htmx:confirm", confirmDetails) === false) {
        maybeCall(resolve);
        return promise;
      }
    }
    let syncElt = elt;
    let syncStrategy = getClosestAttributeValue(elt, "hx-sync");
    let queueStrategy = null;
    let abortable = false;
    if (syncStrategy) {
      const syncStrings = syncStrategy.split(":");
      const selector = syncStrings[0].trim();
      if (selector === "this") {
        syncElt = findThisElement(elt, "hx-sync");
      } else {
        syncElt = asElement(querySelectorExt(elt, selector));
      }
      syncStrategy = (syncStrings[1] || "drop").trim();
      eltData = getInternalData(syncElt);
      if (syncStrategy === "drop" && eltData.xhr && eltData.abortable !== true) {
        maybeCall(resolve);
        return promise;
      } else if (syncStrategy === "abort") {
        if (eltData.xhr) {
          maybeCall(resolve);
          return promise;
        } else {
          abortable = true;
        }
      } else if (syncStrategy === "replace") {
        triggerEvent(syncElt, "htmx:abort");
      } else if (syncStrategy.indexOf("queue") === 0) {
        const queueStrArray = syncStrategy.split(" ");
        queueStrategy = (queueStrArray[1] || "last").trim();
      }
    }
    if (eltData.xhr) {
      if (eltData.abortable) {
        triggerEvent(syncElt, "htmx:abort");
      } else {
        if (queueStrategy == null) {
          if (event) {
            const eventData = getInternalData(event);
            if (eventData && eventData.triggerSpec && eventData.triggerSpec.queue) {
              queueStrategy = eventData.triggerSpec.queue;
            }
          }
          if (queueStrategy == null) {
            queueStrategy = "last";
          }
        }
        if (eltData.queuedRequests == null) {
          eltData.queuedRequests = [];
        }
        if (queueStrategy === "first" && eltData.queuedRequests.length === 0) {
          eltData.queuedRequests.push(function() {
            issueAjaxRequest(verb, path, elt, event, etc);
          });
        } else if (queueStrategy === "all") {
          eltData.queuedRequests.push(function() {
            issueAjaxRequest(verb, path, elt, event, etc);
          });
        } else if (queueStrategy === "last") {
          eltData.queuedRequests = [];
          eltData.queuedRequests.push(function() {
            issueAjaxRequest(verb, path, elt, event, etc);
          });
        }
        maybeCall(resolve);
        return promise;
      }
    }
    const xhr = new XMLHttpRequest();
    eltData.xhr = xhr;
    eltData.abortable = abortable;
    const endRequestLock = function() {
      eltData.xhr = null;
      eltData.abortable = false;
      if (eltData.queuedRequests != null && eltData.queuedRequests.length > 0) {
        const queuedRequest = eltData.queuedRequests.shift();
        queuedRequest();
      }
    };
    const promptQuestion = getClosestAttributeValue(elt, "hx-prompt");
    if (promptQuestion) {
      var promptResponse = prompt(promptQuestion);
      if (promptResponse === null || !triggerEvent(elt, "htmx:prompt", { prompt: promptResponse, target })) {
        maybeCall(resolve);
        endRequestLock();
        return promise;
      }
    }
    if (confirmQuestion && !confirmed) {
      if (!confirm(confirmQuestion)) {
        maybeCall(resolve);
        endRequestLock();
        return promise;
      }
    }
    let headers = getHeaders(elt, target, promptResponse);
    if (verb !== "get" && !usesFormData(elt)) {
      headers["Content-Type"] = "application/x-www-form-urlencoded";
    }
    if (etc.headers) {
      headers = mergeObjects(headers, etc.headers);
    }
    const results = getInputValues(elt, verb);
    let errors = results.errors;
    const rawFormData = results.formData;
    if (etc.values) {
      overrideFormData(rawFormData, formDataFromObject(etc.values));
    }
    const expressionVars = formDataFromObject(getExpressionVars(elt));
    const allFormData = overrideFormData(rawFormData, expressionVars);
    let filteredFormData = filterValues(allFormData, elt);
    if (htmx.config.getCacheBusterParam && verb === "get") {
      filteredFormData.set("org.htmx.cache-buster", getRawAttribute(target, "id") || "true");
    }
    if (path == null || path === "") {
      path = getDocument().location.href;
    }
    const requestAttrValues = getValuesForElement(elt, "hx-request");
    const eltIsBoosted = getInternalData(elt).boosted;
    let useUrlParams = htmx.config.methodsThatUseUrlParams.indexOf(verb) >= 0;
    const requestConfig = {
      boosted: eltIsBoosted,
      useUrlParams,
      formData: filteredFormData,
      parameters: formDataProxy(filteredFormData),
      unfilteredFormData: allFormData,
      unfilteredParameters: formDataProxy(allFormData),
      headers,
      target,
      verb,
      errors,
      withCredentials: etc.credentials || requestAttrValues.credentials || htmx.config.withCredentials,
      timeout: etc.timeout || requestAttrValues.timeout || htmx.config.timeout,
      path,
      triggeringEvent: event
    };
    if (!triggerEvent(elt, "htmx:configRequest", requestConfig)) {
      maybeCall(resolve);
      endRequestLock();
      return promise;
    }
    path = requestConfig.path;
    verb = requestConfig.verb;
    headers = requestConfig.headers;
    filteredFormData = formDataFromObject(requestConfig.parameters);
    errors = requestConfig.errors;
    useUrlParams = requestConfig.useUrlParams;
    if (errors && errors.length > 0) {
      triggerEvent(elt, "htmx:validation:halted", requestConfig);
      maybeCall(resolve);
      endRequestLock();
      return promise;
    }
    const splitPath = path.split("#");
    const pathNoAnchor = splitPath[0];
    const anchor = splitPath[1];
    let finalPath = path;
    if (useUrlParams) {
      finalPath = pathNoAnchor;
      const hasValues = !filteredFormData.keys().next().done;
      if (hasValues) {
        if (finalPath.indexOf("?") < 0) {
          finalPath += "?";
        } else {
          finalPath += "&";
        }
        finalPath += urlEncode(filteredFormData);
        if (anchor) {
          finalPath += "#" + anchor;
        }
      }
    }
    if (!verifyPath(elt, finalPath, requestConfig)) {
      triggerErrorEvent(elt, "htmx:invalidPath", requestConfig);
      maybeCall(reject);
      return promise;
    }
    xhr.open(verb.toUpperCase(), finalPath, true);
    xhr.overrideMimeType("text/html");
    xhr.withCredentials = requestConfig.withCredentials;
    xhr.timeout = requestConfig.timeout;
    if (requestAttrValues.noHeaders) ; else {
      for (const header in headers) {
        if (headers.hasOwnProperty(header)) {
          const headerValue = headers[header];
          safelySetHeaderValue(xhr, header, headerValue);
        }
      }
    }
    const responseInfo = {
      xhr,
      target,
      requestConfig,
      etc,
      boosted: eltIsBoosted,
      select,
      pathInfo: {
        requestPath: path,
        finalRequestPath: finalPath,
        responsePath: null,
        anchor
      }
    };
    xhr.onload = function() {
      try {
        const hierarchy = hierarchyForElt(elt);
        responseInfo.pathInfo.responsePath = getPathFromResponse(xhr);
        responseHandler(elt, responseInfo);
        if (responseInfo.keepIndicators !== true) {
          removeRequestIndicators(indicators, disableElts);
        }
        triggerEvent(elt, "htmx:afterRequest", responseInfo);
        triggerEvent(elt, "htmx:afterOnLoad", responseInfo);
        if (!bodyContains(elt)) {
          let secondaryTriggerElt = null;
          while (hierarchy.length > 0 && secondaryTriggerElt == null) {
            const parentEltInHierarchy = hierarchy.shift();
            if (bodyContains(parentEltInHierarchy)) {
              secondaryTriggerElt = parentEltInHierarchy;
            }
          }
          if (secondaryTriggerElt) {
            triggerEvent(secondaryTriggerElt, "htmx:afterRequest", responseInfo);
            triggerEvent(secondaryTriggerElt, "htmx:afterOnLoad", responseInfo);
          }
        }
        maybeCall(resolve);
        endRequestLock();
      } catch (e) {
        triggerErrorEvent(elt, "htmx:onLoadError", mergeObjects({ error: e }, responseInfo));
        throw e;
      }
    };
    xhr.onerror = function() {
      removeRequestIndicators(indicators, disableElts);
      triggerErrorEvent(elt, "htmx:afterRequest", responseInfo);
      triggerErrorEvent(elt, "htmx:sendError", responseInfo);
      maybeCall(reject);
      endRequestLock();
    };
    xhr.onabort = function() {
      removeRequestIndicators(indicators, disableElts);
      triggerErrorEvent(elt, "htmx:afterRequest", responseInfo);
      triggerErrorEvent(elt, "htmx:sendAbort", responseInfo);
      maybeCall(reject);
      endRequestLock();
    };
    xhr.ontimeout = function() {
      removeRequestIndicators(indicators, disableElts);
      triggerErrorEvent(elt, "htmx:afterRequest", responseInfo);
      triggerErrorEvent(elt, "htmx:timeout", responseInfo);
      maybeCall(reject);
      endRequestLock();
    };
    if (!triggerEvent(elt, "htmx:beforeRequest", responseInfo)) {
      maybeCall(resolve);
      endRequestLock();
      return promise;
    }
    var indicators = addRequestIndicatorClasses(elt);
    var disableElts = disableElements(elt);
    forEach(["loadstart", "loadend", "progress", "abort"], function(eventName) {
      forEach([xhr, xhr.upload], function(target2) {
        target2.addEventListener(eventName, function(event2) {
          triggerEvent(elt, "htmx:xhr:" + eventName, {
            lengthComputable: event2.lengthComputable,
            loaded: event2.loaded,
            total: event2.total
          });
        });
      });
    });
    triggerEvent(elt, "htmx:beforeSend", responseInfo);
    const params = useUrlParams ? null : encodeParamsForBody(xhr, elt, filteredFormData);
    xhr.send(params);
    return promise;
  }
  function determineHistoryUpdates(elt, responseInfo) {
    const xhr = responseInfo.xhr;
    let pathFromHeaders = null;
    let typeFromHeaders = null;
    if (hasHeader(xhr, /HX-Push:/i)) {
      pathFromHeaders = xhr.getResponseHeader("HX-Push");
      typeFromHeaders = "push";
    } else if (hasHeader(xhr, /HX-Push-Url:/i)) {
      pathFromHeaders = xhr.getResponseHeader("HX-Push-Url");
      typeFromHeaders = "push";
    } else if (hasHeader(xhr, /HX-Replace-Url:/i)) {
      pathFromHeaders = xhr.getResponseHeader("HX-Replace-Url");
      typeFromHeaders = "replace";
    }
    if (pathFromHeaders) {
      if (pathFromHeaders === "false") {
        return {};
      } else {
        return {
          type: typeFromHeaders,
          path: pathFromHeaders
        };
      }
    }
    const requestPath = responseInfo.pathInfo.finalRequestPath;
    const responsePath = responseInfo.pathInfo.responsePath;
    const pushUrl = getClosestAttributeValue(elt, "hx-push-url");
    const replaceUrl = getClosestAttributeValue(elt, "hx-replace-url");
    const elementIsBoosted = getInternalData(elt).boosted;
    let saveType = null;
    let path = null;
    if (pushUrl) {
      saveType = "push";
      path = pushUrl;
    } else if (replaceUrl) {
      saveType = "replace";
      path = replaceUrl;
    } else if (elementIsBoosted) {
      saveType = "push";
      path = responsePath || requestPath;
    }
    if (path) {
      if (path === "false") {
        return {};
      }
      if (path === "true") {
        path = responsePath || requestPath;
      }
      if (responseInfo.pathInfo.anchor && path.indexOf("#") === -1) {
        path = path + "#" + responseInfo.pathInfo.anchor;
      }
      return {
        type: saveType,
        path
      };
    } else {
      return {};
    }
  }
  function codeMatches(responseHandlingConfig, status) {
    var regExp = new RegExp(responseHandlingConfig.code);
    return regExp.test(status.toString(10));
  }
  function resolveResponseHandling(xhr) {
    for (var i = 0; i < htmx.config.responseHandling.length; i++) {
      var responseHandlingElement = htmx.config.responseHandling[i];
      if (codeMatches(responseHandlingElement, xhr.status)) {
        return responseHandlingElement;
      }
    }
    return {
      swap: false
    };
  }
  function handleTitle(title) {
    if (title) {
      const titleElt = find("title");
      if (titleElt) {
        titleElt.innerHTML = title;
      } else {
        window.document.title = title;
      }
    }
  }
  function handleAjaxResponse(elt, responseInfo) {
    const xhr = responseInfo.xhr;
    let target = responseInfo.target;
    const etc = responseInfo.etc;
    const responseInfoSelect = responseInfo.select;
    if (!triggerEvent(elt, "htmx:beforeOnLoad", responseInfo)) return;
    if (hasHeader(xhr, /HX-Trigger:/i)) {
      handleTriggerHeader(xhr, "HX-Trigger", elt);
    }
    if (hasHeader(xhr, /HX-Location:/i)) {
      saveCurrentPageToHistory();
      let redirectPath = xhr.getResponseHeader("HX-Location");
      var redirectSwapSpec;
      if (redirectPath.indexOf("{") === 0) {
        redirectSwapSpec = parseJSON(redirectPath);
        redirectPath = redirectSwapSpec.path;
        delete redirectSwapSpec.path;
      }
      ajaxHelper("get", redirectPath, redirectSwapSpec).then(function() {
        pushUrlIntoHistory(redirectPath);
      });
      return;
    }
    const shouldRefresh = hasHeader(xhr, /HX-Refresh:/i) && xhr.getResponseHeader("HX-Refresh") === "true";
    if (hasHeader(xhr, /HX-Redirect:/i)) {
      responseInfo.keepIndicators = true;
      location.href = xhr.getResponseHeader("HX-Redirect");
      shouldRefresh && location.reload();
      return;
    }
    if (shouldRefresh) {
      responseInfo.keepIndicators = true;
      location.reload();
      return;
    }
    if (hasHeader(xhr, /HX-Retarget:/i)) {
      if (xhr.getResponseHeader("HX-Retarget") === "this") {
        responseInfo.target = elt;
      } else {
        responseInfo.target = asElement(querySelectorExt(elt, xhr.getResponseHeader("HX-Retarget")));
      }
    }
    const historyUpdate = determineHistoryUpdates(elt, responseInfo);
    const responseHandling = resolveResponseHandling(xhr);
    const shouldSwap = responseHandling.swap;
    let isError = !!responseHandling.error;
    let ignoreTitle = htmx.config.ignoreTitle || responseHandling.ignoreTitle;
    let selectOverride = responseHandling.select;
    if (responseHandling.target) {
      responseInfo.target = asElement(querySelectorExt(elt, responseHandling.target));
    }
    var swapOverride = etc.swapOverride;
    if (swapOverride == null && responseHandling.swapOverride) {
      swapOverride = responseHandling.swapOverride;
    }
    if (hasHeader(xhr, /HX-Retarget:/i)) {
      if (xhr.getResponseHeader("HX-Retarget") === "this") {
        responseInfo.target = elt;
      } else {
        responseInfo.target = asElement(querySelectorExt(elt, xhr.getResponseHeader("HX-Retarget")));
      }
    }
    if (hasHeader(xhr, /HX-Reswap:/i)) {
      swapOverride = xhr.getResponseHeader("HX-Reswap");
    }
    var serverResponse = xhr.response;
    var beforeSwapDetails = mergeObjects({
      shouldSwap,
      serverResponse,
      isError,
      ignoreTitle,
      selectOverride
    }, responseInfo);
    if (responseHandling.event && !triggerEvent(target, responseHandling.event, beforeSwapDetails)) return;
    if (!triggerEvent(target, "htmx:beforeSwap", beforeSwapDetails)) return;
    target = beforeSwapDetails.target;
    serverResponse = beforeSwapDetails.serverResponse;
    isError = beforeSwapDetails.isError;
    ignoreTitle = beforeSwapDetails.ignoreTitle;
    selectOverride = beforeSwapDetails.selectOverride;
    responseInfo.target = target;
    responseInfo.failed = isError;
    responseInfo.successful = !isError;
    if (beforeSwapDetails.shouldSwap) {
      if (xhr.status === 286) {
        cancelPolling(elt);
      }
      withExtensions(elt, function(extension) {
        serverResponse = extension.transformResponse(serverResponse, xhr, elt);
      });
      if (historyUpdate.type) {
        saveCurrentPageToHistory();
      }
      if (hasHeader(xhr, /HX-Reswap:/i)) {
        swapOverride = xhr.getResponseHeader("HX-Reswap");
      }
      var swapSpec = getSwapSpecification(elt, swapOverride);
      if (!swapSpec.hasOwnProperty("ignoreTitle")) {
        swapSpec.ignoreTitle = ignoreTitle;
      }
      target.classList.add(htmx.config.swappingClass);
      let settleResolve = null;
      let settleReject = null;
      if (responseInfoSelect) {
        selectOverride = responseInfoSelect;
      }
      if (hasHeader(xhr, /HX-Reselect:/i)) {
        selectOverride = xhr.getResponseHeader("HX-Reselect");
      }
      const selectOOB = getClosestAttributeValue(elt, "hx-select-oob");
      const select = getClosestAttributeValue(elt, "hx-select");
      let doSwap = function() {
        try {
          if (historyUpdate.type) {
            triggerEvent(getDocument().body, "htmx:beforeHistoryUpdate", mergeObjects({ history: historyUpdate }, responseInfo));
            if (historyUpdate.type === "push") {
              pushUrlIntoHistory(historyUpdate.path);
              triggerEvent(getDocument().body, "htmx:pushedIntoHistory", { path: historyUpdate.path });
            } else {
              replaceUrlInHistory(historyUpdate.path);
              triggerEvent(getDocument().body, "htmx:replacedInHistory", { path: historyUpdate.path });
            }
          }
          swap(target, serverResponse, swapSpec, {
            select: selectOverride || select,
            selectOOB,
            eventInfo: responseInfo,
            anchor: responseInfo.pathInfo.anchor,
            contextElement: elt,
            afterSwapCallback: function() {
              if (hasHeader(xhr, /HX-Trigger-After-Swap:/i)) {
                let finalElt = elt;
                if (!bodyContains(elt)) {
                  finalElt = getDocument().body;
                }
                handleTriggerHeader(xhr, "HX-Trigger-After-Swap", finalElt);
              }
            },
            afterSettleCallback: function() {
              if (hasHeader(xhr, /HX-Trigger-After-Settle:/i)) {
                let finalElt = elt;
                if (!bodyContains(elt)) {
                  finalElt = getDocument().body;
                }
                handleTriggerHeader(xhr, "HX-Trigger-After-Settle", finalElt);
              }
              maybeCall(settleResolve);
            }
          });
        } catch (e) {
          triggerErrorEvent(elt, "htmx:swapError", responseInfo);
          maybeCall(settleReject);
          throw e;
        }
      };
      let shouldTransition = htmx.config.globalViewTransitions;
      if (swapSpec.hasOwnProperty("transition")) {
        shouldTransition = swapSpec.transition;
      }
      if (shouldTransition && triggerEvent(elt, "htmx:beforeTransition", responseInfo) && typeof Promise !== "undefined" && // @ts-ignore experimental feature atm
      document.startViewTransition) {
        const settlePromise = new Promise(function(_resolve, _reject) {
          settleResolve = _resolve;
          settleReject = _reject;
        });
        const innerDoSwap = doSwap;
        doSwap = function() {
          document.startViewTransition(function() {
            innerDoSwap();
            return settlePromise;
          });
        };
      }
      if (swapSpec.swapDelay > 0) {
        getWindow().setTimeout(doSwap, swapSpec.swapDelay);
      } else {
        doSwap();
      }
    }
    if (isError) {
      triggerErrorEvent(elt, "htmx:responseError", mergeObjects({ error: "Response Status Error Code " + xhr.status + " from " + responseInfo.pathInfo.requestPath }, responseInfo));
    }
  }
  const extensions = {};
  function extensionBase() {
    return {
      init: function(api2) {
        return null;
      },
      getSelectors: function() {
        return null;
      },
      onEvent: function(name, evt) {
        return true;
      },
      transformResponse: function(text, xhr, elt) {
        return text;
      },
      isInlineSwap: function(swapStyle) {
        return false;
      },
      handleSwap: function(swapStyle, target, fragment, settleInfo) {
        return false;
      },
      encodeParameters: function(xhr, parameters, elt) {
        return null;
      }
    };
  }
  function defineExtension(name, extension) {
    if (extension.init) {
      extension.init(internalAPI);
    }
    extensions[name] = mergeObjects(extensionBase(), extension);
  }
  function removeExtension(name) {
    delete extensions[name];
  }
  function getExtensions(elt, extensionsToReturn, extensionsToIgnore) {
    if (extensionsToReturn == void 0) {
      extensionsToReturn = [];
    }
    if (elt == void 0) {
      return extensionsToReturn;
    }
    if (extensionsToIgnore == void 0) {
      extensionsToIgnore = [];
    }
    const extensionsForElement = getAttributeValue(elt, "hx-ext");
    if (extensionsForElement) {
      forEach(extensionsForElement.split(","), function(extensionName) {
        extensionName = extensionName.replace(/ /g, "");
        if (extensionName.slice(0, 7) == "ignore:") {
          extensionsToIgnore.push(extensionName.slice(7));
          return;
        }
        if (extensionsToIgnore.indexOf(extensionName) < 0) {
          const extension = extensions[extensionName];
          if (extension && extensionsToReturn.indexOf(extension) < 0) {
            extensionsToReturn.push(extension);
          }
        }
      });
    }
    return getExtensions(asElement(parentElt(elt)), extensionsToReturn, extensionsToIgnore);
  }
  var isReady = false;
  getDocument().addEventListener("DOMContentLoaded", function() {
    isReady = true;
  });
  function ready(fn) {
    if (isReady || getDocument().readyState === "complete") {
      fn();
    } else {
      getDocument().addEventListener("DOMContentLoaded", fn);
    }
  }
  function insertIndicatorStyles() {
    if (htmx.config.includeIndicatorStyles !== false) {
      const nonceAttribute = htmx.config.inlineStyleNonce ? ` nonce="${htmx.config.inlineStyleNonce}"` : "";
      getDocument().head.insertAdjacentHTML(
        "beforeend",
        "<style" + nonceAttribute + ">      ." + htmx.config.indicatorClass + "{opacity:0}      ." + htmx.config.requestClass + " ." + htmx.config.indicatorClass + "{opacity:1; transition: opacity 200ms ease-in;}      ." + htmx.config.requestClass + "." + htmx.config.indicatorClass + "{opacity:1; transition: opacity 200ms ease-in;}      </style>"
      );
    }
  }
  function getMetaConfig() {
    const element = getDocument().querySelector('meta[name="htmx-config"]');
    if (element) {
      return parseJSON(element.content);
    } else {
      return null;
    }
  }
  function mergeMetaConfig() {
    const metaConfig = getMetaConfig();
    if (metaConfig) {
      htmx.config = mergeObjects(htmx.config, metaConfig);
    }
  }
  ready(function() {
    mergeMetaConfig();
    insertIndicatorStyles();
    let body = getDocument().body;
    processNode(body);
    const restoredElts = getDocument().querySelectorAll(
      "[hx-trigger='restored'],[data-hx-trigger='restored']"
    );
    body.addEventListener("htmx:abort", function(evt) {
      const target = evt.target;
      const internalData = getInternalData(target);
      if (internalData && internalData.xhr) {
        internalData.xhr.abort();
      }
    });
    const originalPopstate = window.onpopstate ? window.onpopstate.bind(window) : null;
    window.onpopstate = function(event) {
      if (event.state && event.state.htmx) {
        restoreHistory();
        forEach(restoredElts, function(elt) {
          triggerEvent(elt, "htmx:restored", {
            document: getDocument(),
            triggerEvent
          });
        });
      } else {
        if (originalPopstate) {
          originalPopstate(event);
        }
      }
    };
    getWindow().setTimeout(function() {
      triggerEvent(body, "htmx:load", {});
      body = null;
    }, 0);
  });
  return htmx;
}();
var htmx_esm_default = htmx2;

// htmxextensions/pathdeps.ts
function dependsOn(pathSpec, url) {
  if (pathSpec === "ignore") {
    return false;
  }
  const dependencyPath = pathSpec.split("/");
  const urlPath = url.split("/");
  for (let i = 0; i < urlPath.length; i++) {
    const dependencyElement = dependencyPath.shift();
    const pathElement = urlPath[i];
    if (dependencyElement !== pathElement && dependencyElement !== "*") {
      return false;
    }
    if (dependencyPath.length === 0 || dependencyPath.length === 1 && dependencyPath[0] === "") {
      return true;
    }
  }
  return false;
}
function refreshPath(path) {
  const eltsWithDeps = htmx_esm_default.findAll("[path-deps]");
  for (let i = 0; i < eltsWithDeps.length; i++) {
    const elt = eltsWithDeps[i];
    if (dependsOn(elt.getAttribute("path-deps"), path)) {
      htmx_esm_default.trigger(elt, "path-deps", null);
    }
  }
}
htmx_esm_default.defineExtension("path-deps", {
  // @ts-ignore
  onEvent: function(name, evt) {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    if (name === "htmx:beforeOnLoad") {
      const config2 = evt.detail.requestConfig;
      if (config2 && config2.verb !== "get" && evt.target != null && evt.target instanceof Element && evt.target.getAttribute("path-deps") !== "ignore") {
        refreshPath(config2.path);
      }
    }
  }
});

// htmxextensions/trigger-children.ts
htmx_esm_default.defineExtension("trigger-children", {
  onEvent: (name, evt) => {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    const target = evt.detail.target;
    if (target && target.children) {
      Array.from(target.children).forEach((e) => {
        htmx_esm_default.trigger(e, name, null);
      });
    }
    return true;
  },
  init: function(api2) {
  },
  transformResponse: function(text, xhr, elt) {
    return text;
  },
  isInlineSwap: function(swapStyle) {
    return false;
  },
  handleSwap: function(swapStyle, target, fragment, settleInfo) {
    return false;
  },
  encodeParameters: function(xhr, parameters, elt) {
  },
  getSelectors: function() {
    return null;
  }
});

// htmxextensions/debug.ts
htmx_esm_default.defineExtension("debug", {
  // @ts-ignore
  onEvent: function(name, evt) {
    if (console.debug) {
      console.debug(name);
    } else if (console) {
      console.log("DEBUG:", name);
    } else ;
  }
});

// htmxextensions/response-targets.ts
var config = htmx_esm_default.config;
var api;
var attrPrefix = "hx-target-";
function startsWith2(str2, prefix) {
  return str2.substring(0, prefix.length) === prefix;
}
function getRespCodeTarget(elt, respCodeNumber) {
  if (!elt || !respCodeNumber) return null;
  const respCode = respCodeNumber.toString();
  const attrPossibilities = [
    respCode,
    respCode.substr(0, 2) + "*",
    respCode.substr(0, 2) + "x",
    respCode.substr(0, 1) + "*",
    respCode.substr(0, 1) + "x",
    respCode.substr(0, 1) + "**",
    respCode.substr(0, 1) + "xx",
    "*",
    "x",
    "***",
    "xxx"
  ];
  if (startsWith2(respCode, "4") || startsWith2(respCode, "5")) {
    attrPossibilities.push("error");
  }
  for (let i = 0; i < attrPossibilities.length; i++) {
    const attr = attrPrefix + attrPossibilities[i];
    const attrValue = api.getClosestAttributeValue(elt, attr);
    if (attrValue) {
      if (attrValue === "this") {
        return api.findThisElement(elt, attr);
      } else {
        return api.querySelectorExt(elt, attrValue);
      }
    }
  }
  return null;
}
function handleErrorFlag(evt) {
  if (evt.detail.isError) {
    if (config.responseTargetUnsetsError) {
      evt.detail.isError = false;
    }
  } else if (config.responseTargetSetsError) {
    evt.detail.isError = true;
  }
}
htmx_esm_default.defineExtension("response-targets", {
  // @ts-ignore
  init: (apiRef) => {
    api = apiRef;
    if (config.responseTargetUnsetsError === void 0) {
      config.responseTargetUnsetsError = true;
    }
    if (config.responseTargetSetsError === void 0) {
      config.responseTargetSetsError = false;
    }
    if (config.responseTargetPrefersExisting === void 0) {
      config.responseTargetPrefersExisting = false;
    }
    if (config.responseTargetPrefersRetargetHeader === void 0) {
      config.responseTargetPrefersRetargetHeader = true;
    }
  },
  // @ts-ignore
  onEvent: (name, evt) => {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    if (name === "htmx:beforeSwap" && evt.detail.xhr && evt.detail.xhr.status !== 200) {
      if (evt.detail.target) {
        if (config.responseTargetPrefersExisting) {
          evt.detail.shouldSwap = true;
          handleErrorFlag(evt);
          return true;
        }
        if (config.responseTargetPrefersRetargetHeader && evt.detail.xhr.getAllResponseHeaders().match(/HX-Retarget:/i)) {
          evt.detail.shouldSwap = true;
          handleErrorFlag(evt);
          return true;
        }
      }
      if (!evt.detail.requestConfig) {
        return true;
      }
      const target = getRespCodeTarget(
        evt.detail.requestConfig.elt,
        evt.detail.xhr.status
      );
      if (target) {
        handleErrorFlag(evt);
        evt.detail.shouldSwap = true;
        evt.detail.target = target;
      }
      return true;
    }
  }
});

// htmxextensions/mutation-error.ts
htmx_esm_default.defineExtension("mutation-error", {
  // @ts-ignore
  onEvent: (name, evt) => {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    if (name === "htmx:afterRequest") {
      if (!evt.detail || !evt.detail.xhr) {
        return;
      }
      const status = evt.detail.xhr.status;
      if (status >= 400) {
        htmx_esm_default.findAll("[hx-on\\:\\:mutation-error]").forEach((element) => {
          htmx_esm_default.trigger(element, "htmx:mutation-error", { status });
        });
      }
    }
  }
});

// util/ws.ts
function createWebSocketClient(opts) {
  let socket = null;
  const connect = (tries) => {
    console.log("connecting to ws", opts.url, "attempt", tries);
    socket = new WebSocket(opts.url);
    socket.onmessage = (event) => {
      opts.onMessage(event.data);
    };
    socket.onerror = (error) => {
      try {
        socket?.close();
      } catch (ex) {
      }
      socket = null;
      let interval = tries * (opts.reconnectInterval || 1e3);
      setTimeout(() => connect(tries + 1), interval);
    };
    socket.onclose = () => {
      socket = null;
      let interval = tries * (opts.reconnectInterval || 1e3);
      setTimeout(() => connect(tries + 1), interval);
    };
  };
  connect(1);
  const sendMessage = (message) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(message);
    } else {
      setTimeout(() => sendMessage(message), 100);
    }
  };
  return { sendMessage };
}

// htmxextensions/livereload.ts
var lastVersion = "";
htmx_esm_default.defineExtension("livereload", {
  init: function() {
    const host = window.location.host;
    let enabled = false;
    for (const element of Array.from(htmx_esm_default.findAll("[hx-ext]"))) {
      const value = element.getAttribute("hx-ext");
      if (value?.split(" ").includes("livereload")) {
        enabled = true;
        break;
      }
    }
    if (!enabled) {
      return;
    }
    console.log("livereload extension initialized.");
    createWebSocketClient({
      url: `${window.location.protocol === "https:" ? "wss" : "ws"}://${host}/dev/livereload`,
      onOpen: () => {
      },
      onMessage: (message) => {
        if (lastVersion === "") {
          lastVersion = message;
        }
        if (lastVersion !== message) {
          lastVersion = message;
          reload();
        }
      },
      onError: (error) => {
      },
      onClose: () => {
      }
    });
  },
  // @ts-ignore
  onEvent: function(name, evt) {
  }
});
function reload() {
  window.location.reload();
}

// htmxextensions/htmgo.ts
var evalFuncRegex = /__eval_[A-Za-z0-9]+\(\)/gm;
htmx_esm_default.defineExtension("htmgo", {
  // @ts-ignore
  onEvent: function(name, evt) {
    if (name === "htmx:beforeCleanupElement" && evt.target) {
      removeAssociatedScripts(evt.target);
    }
  }
});
function removeAssociatedScripts(element) {
  const attributes = Array.from(element.attributes);
  for (let attribute of attributes) {
    const matches2 = attribute.value.match(evalFuncRegex) || [];
    for (let match of matches2) {
      const id = match.replace("()", "");
      const ele = document.getElementById(id);
      if (ele && ele.tagName === "SCRIPT") {
        ele.remove();
      }
    }
  }
}

// htmgo.ts
function watchUrl(callback) {
  let lastUrl = window.location.href;
  setInterval(() => {
    if (window.location.href !== lastUrl) {
      callback(lastUrl, window.location.href);
      lastUrl = window.location.href;
    }
  }, 100);
}
watchUrl((_, newUrl) => {
  onUrlChange(newUrl);
});
function onUrlChange(newUrl) {
  let url = new URL(newUrl);
  document.querySelectorAll("[hx-trigger]").forEach(function(element) {
    const triggers = element.getAttribute("hx-trigger");
    if (!triggers) {
      return;
    }
    const split = triggers.split(", ");
    if (split.find((s) => s === "url")) {
      htmx_esm_default.swap(element, "url", {
        swapStyle: "outerHTML",
        swapDelay: 0,
        settleDelay: 0
      });
    } else {
      for (let [key, values] of url.searchParams) {
        let eventName = "qs:" + key;
        if (triggers.includes(eventName)) {
          console.log("triggering", eventName);
          htmx_esm_default.trigger(element, eventName, null);
          break;
        }
      }
    }
  });
  document.querySelectorAll("[hx-match-qp]").forEach((el) => {
    let hasMatch = false;
    for (let name of el.getAttributeNames()) {
      if (name.startsWith("hx-match-qp-mapping:")) {
        let match = name.replace("hx-match-qp-mapping:", "");
        let value = url.searchParams.get(match);
        if (value) {
          htmx_esm_default.swap(el, el.getAttribute(name) ?? "", {
            swapStyle: "innerHTML",
            swapDelay: 0,
            settleDelay: 0
          });
          hasMatch = true;
          break;
        }
      }
    }
    if (!hasMatch) {
      let defaultKey = el.getAttribute("hx-match-qp-default");
      if (defaultKey) {
        htmx_esm_default.swap(
          el,
          el.getAttribute("hx-match-qp-mapping:" + defaultKey) ?? "",
          { swapStyle: "innerHTML", swapDelay: 0, settleDelay: 0 }
        );
      }
    }
  });
}
//# sourceMappingURL=htmgo.js.map
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIi4uL2pzL25vZGVfbW9kdWxlcy9odG14Lm9yZy9kaXN0L2h0bXguZXNtLmpzIiwiLi4vanMvaHRteGV4dGVuc2lvbnMvcGF0aGRlcHMudHMiLCIuLi9qcy9odG14ZXh0ZW5zaW9ucy90cmlnZ2VyLWNoaWxkcmVuLnRzIiwiLi4vanMvaHRteGV4dGVuc2lvbnMvZGVidWcudHMiLCIuLi9qcy9odG14ZXh0ZW5zaW9ucy9yZXNwb25zZS10YXJnZXRzLnRzIiwiLi4vanMvaHRteGV4dGVuc2lvbnMvbXV0YXRpb24tZXJyb3IudHMiLCIuLi9qcy91dGlsL3dzLnRzIiwiLi4vanMvaHRteGV4dGVuc2lvbnMvbGl2ZXJlbG9hZC50cyIsIi4uL2pzL2h0bXhleHRlbnNpb25zL2h0bWdvLnRzIiwiLi4vanMvaHRtZ28udHMiXSwibmFtZXMiOlsiaHRteCIsInN0ciIsImVsdCIsImV4dGVuc2lvbnMiLCJwYXJlbnRFbHQiLCJwcm9tcHQiLCJ0YXJnZXQiLCJldmVudCIsImFwaSIsImNvbmZpZyIsInN0YXJ0c1dpdGgiLCJtYXRjaGVzIl0sIm1hcHBpbmdzIjoiO0FBQUEsSUFBSUEsUUFBUSxXQUFXO0FBSXJCLEVBQUEsTUFBTSxJQUFPLEdBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxJQUlYLE1BQVEsRUFBQSxJQUFBO0FBQUE7QUFBQSxJQUVSLE9BQVMsRUFBQSxJQUFBO0FBQUE7QUFBQSxJQUVULEVBQUksRUFBQSxJQUFBO0FBQUE7QUFBQSxJQUVKLEdBQUssRUFBQSxJQUFBO0FBQUE7QUFBQSxJQUVMLE9BQVMsRUFBQSxJQUFBO0FBQUE7QUFBQSxJQUVULElBQU0sRUFBQSxJQUFBO0FBQUE7QUFBQTtBQUFBLElBR04sSUFBTSxFQUFBLElBQUE7QUFBQTtBQUFBLElBRU4sT0FBUyxFQUFBLElBQUE7QUFBQTtBQUFBLElBRVQsT0FBUyxFQUFBLElBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQSxJQVVULE1BQUEsRUFBUSxT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QUEsQ0FBQTtBQUNuRCxRQUFBLElBQUksS0FBUSxHQUFBLEdBQUEsQ0FBSSxZQUFhLENBQUEsR0FBQSxDQUFJLEtBQUssQ0FBQSxDQUFBO0FBQ3RDLFFBQUEsSUFBSSxLQUFPLEVBQUE7QUFDVCxVQUFBLGdCQUFBLENBQUssS0FBSyxFQUFJLEVBQUEsRUFBQSxDQUFHLFlBQWEsQ0FBQSxJQUFJLEtBQUssRUFBSSxFQUFBO0FBQUEsWUFDekMsU0FBVyxFQUFBLFdBQUE7QUFBQSxZQUNYLFNBQVcsRUFBQSxDQUFBO0FBQUEsWUFDWCxXQUFhLEVBQUEsQ0FBQTtBQUFBLFdBQ2QsQ0FBQSxDQUFBO0FBQ0QsVUFBVyxRQUFBLEdBQUEsSUFBQSxDQUFBO0FBQ1gsVUFBQSxNQUFBO0FBQUEsU0FDRjtBQUFBLE9BQ0Y7QUFBQSxLQUNGO0FBQ0EsSUFBQSxJQUFJLENBQUMsUUFBVSxFQUFBO0FBQ2IsTUFBSSxJQUFBLFVBQUEsR0FBYSxFQUFHLENBQUEsWUFBQSxDQUFhLHFCQUFxQixDQUFBLENBQUE7QUFDdEQsTUFBQSxJQUFJLFVBQVksRUFBQTtBQUNkLFFBQUssZ0JBQUEsQ0FBQSxJQUFBO0FBQUEsVUFDSCxFQUFBO0FBQUEsVUFDQSxFQUFHLENBQUEsWUFBQSxDQUFhLHNCQUF5QixHQUFBLFVBQVUsQ0FBSyxJQUFBLEVBQUE7QUFBQSxVQUN4RCxFQUFFLFNBQVcsRUFBQSxXQUFBLEVBQWEsU0FBVyxFQUFBLENBQUEsRUFBRyxhQUFhLENBQUUsRUFBQTtBQUFBLFNBQ3pELENBQUE7QUFBQSxPQUNGO0FBQUEsS0FDRjtBQUFBLEdBQ0QsQ0FBQSxDQUFBO0FBQ0giLCJmaWxlIjoiaHRtZ28uanMiLCJzb3VyY2VzQ29udGVudCI6WyJ2YXIgaHRteCA9IChmdW5jdGlvbigpIHtcbiAgJ3VzZSBzdHJpY3QnXG5cbiAgLy8gUHVibGljIEFQSVxuICBjb25zdCBodG14ID0ge1xuICAgIC8vIFRzYyBtYWRuZXNzIGhlcmUsIGFzc2lnbmluZyB0aGUgZnVuY3Rpb25zIGRpcmVjdGx5IHJlc3VsdHMgaW4gYW4gaW52YWxpZCBUeXBlU2NyaXB0IG91dHB1dCwgYnV0IHJlYXNzaWduaW5nIGlzIGZpbmVcbiAgICAvKiBFdmVudCBwcm9jZXNzaW5nICovXG4gICAgLyoqIEB0eXBlIHt0eXBlb2Ygb25Mb2FkSGVscGVyfSAqL1xuICAgIG9uTG9hZDogbnVsbCxcbiAgICAvKiogQHR5cGUge3R5cGVvZiBwcm9jZXNzTm9kZX0gKi9cbiAgICBwcm9jZXNzOiBudWxsLFxuICAgIC8qKiBAdHlwZSB7dHlwZW9mIGFkZEV2ZW50TGlzdGVuZXJJbXBsfSAqL1xuICAgIG9uOiBudWxsLFxuICAgIC8qKiBAdHlwZSB7dHlwZW9mIHJlbW92ZUV2ZW50TGlzdGVuZXJJbXBsfSAqL1xuICAgIG9mZjogbnVsbCxcbiAgICAvKiogQHR5cGUge3R5cGVvZiB0cmlnZ2VyRXZlbnR9ICovXG4gICAgdHJpZ2dlcjogbnVsbCxcbiAgICAvKiogQHR5cGUge3R5cGVvZiBhamF4SGVscGVyfSAqL1xuICAgIGFqYXg6IG51bGwsXG4gICAgLyogRE9NIHF1ZXJ5aW5nIGhlbHBlcnMgKi9cbiAgICAvKiogQHR5cGUge3R5cGVvZiBmaW5kfSAqL1xuICAgIGZpbmQ6IG51bGwsXG4gICAgLyoqIEB0eXBlIHt0eXBlb2YgZmluZEFsbH0gKi9cbiAgICBmaW5kQWxsOiBudWxsLFxuICAgIC8qKiBAdHlwZSB7dHlwZW9mIGNsb3Nlc3R9ICovXG4gICAgY2xvc2VzdDogbnVsbCxcbiAgICAvKipcbiAgICAgKiBSZXR1cm5zIHRoZSBpbnB1dCB2YWx1ZXMgdGhhdCB3b3VsZCByZXNvbHZlIGZvciBhIGdpdmVuIGVsZW1lbnQgdmlhIHRoZSBodG14IHZhbHVlIHJlc29sdXRpb24gbWVjaGFuaXNtXG4gICAgICpcbiAgICAgKiBAc2VlIGh0dHBzOi8vaHRteC5vcmcvYXBpLyN2YWx1ZXNcbiAgICAgKlxuICAgICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0IHRoZSBlbGVtZW50IHRvIHJlc29sdmUgdmFsdWVzIG9uXG4gICAgICogQHBhcmFtIHtIdHRwVmVyYn0gdHlwZSB0aGUgcmVxdWVzdCB0eXBlIChlLmcuICoqZ2V0Kiogb3IgKipwb3N0KiopIG5vbi1HRVQncyB3aWxsIGluY2x1ZGUgdGhlIGVuY2xvc2luZyBmb3JtIG9mIHRoZSBlbGVtZW50LiBEZWZhdWx0cyB0byAqKnBvc3QqKlxuICAgICAqIEByZXR1cm5zIHtPYmplY3R9XG4gICAgICovXG4gICAgdmFsdWVzOiBmdW5jdGlvbihlbHQsIHR5cGUpIHtcbiAgICAgIGNvbnN0IGlucHV0VmFsdWVzID0gZ2V0SW5wdXRWYWx1ZXMoZWx0LCB0eXBlIHx8ICdwb3N0JylcbiAgICAgIHJldHVybiBpbnB1dFZhbHVlcy52YWx1ZXNcbiAgICB9LFxuICAgIC8qIERPTSBtYW5pcHVsYXRpb24gaGVscGVycyAqL1xuICAgIC8qKiBAdHlwZSB7dHlwZW9mIHJlbW92ZUVsZW1lbnR9ICovXG4gICAgcmVtb3ZlOiBudWxsLFxuICAgIC8qKiBAdHlwZSB7dHlwZW9mIGFkZENsYXNzVG9FbGVtZW50fSAqL1xuICAgIGFkZENsYXNzOiBudWxsLFxuICAgIC8qKiBAdHlwZSB7dHlwZW9mIHJlbW92ZUNsYXNzRnJvbUVsZW1lbnR9ICovXG4gICAgcmVtb3ZlQ2xhc3M6IG51bGwsXG4gICAgLyoqIEB0eXBlIHt0eXBlb2YgdG9nZ2xlQ2xhc3NPbkVsZW1lbnR9ICovXG4gICAgdG9nZ2xlQ2xhc3M6IG51bGwsXG4gICAgLyoqIEB0eXBlIHt0eXBlb2YgdGFrZUNsYXNzRm9yRWxlbWVudH0gKi9cbiAgICB0YWtlQ2xhc3M6IG51bGwsXG4gICAgLyoqIEB0eXBlIHt0eXBlb2Ygc3dhcH0gKi9cbiAgICBzd2FwOiBudWxsLFxuICAgIC8qIEV4dGVuc2lvbiBlbnRyeXBvaW50cyAqL1xuICAgIC8qKiBAdHlwZSB7dHlwZW9mIGRlZmluZUV4dGVuc2lvbn0gKi9cbiAgICBkZWZpbmVFeHRlbnNpb246IG51bGwsXG4gICAgLyoqIEB0eXBlIHt0eXBlb2YgcmVtb3ZlRXh0ZW5zaW9ufSAqL1xuICAgIHJlbW92ZUV4dGVuc2lvbjogbnVsbCxcbiAgICAvKiBEZWJ1Z2dpbmcgKi9cbiAgICAvKiogQHR5cGUge3R5cGVvZiBsb2dBbGx9ICovXG4gICAgbG9nQWxsOiBudWxsLFxuICAgIC8qKiBAdHlwZSB7dHlwZW9mIGxvZ05vbmV9ICovXG4gICAgbG9nTm9uZTogbnVsbCxcbiAgICAvKiBEZWJ1Z2dpbmcgKi9cbiAgICAvKipcbiAgICAgKiBUaGUgbG9nZ2VyIGh0bXggdXNlcyB0byBsb2cgd2l0aFxuICAgICAqXG4gICAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jbG9nZ2VyXG4gICAgICovXG4gICAgbG9nZ2VyOiBudWxsLFxuICAgIC8qKlxuICAgICAqIEEgcHJvcGVydHkgaG9sZGluZyB0aGUgY29uZmlndXJhdGlvbiBodG14IHVzZXMgYXQgcnVudGltZS5cbiAgICAgKlxuICAgICAqIE5vdGUgdGhhdCB1c2luZyBhIFttZXRhIHRhZ10oaHR0cHM6Ly9odG14Lm9yZy9kb2NzLyNjb25maWcpIGlzIHRoZSBwcmVmZXJyZWQgbWVjaGFuaXNtIGZvciBzZXR0aW5nIHRoZXNlIHByb3BlcnRpZXMuXG4gICAgICpcbiAgICAgKiBAc2VlIGh0dHBzOi8vaHRteC5vcmcvYXBpLyNjb25maWdcbiAgICAgKi9cbiAgICBjb25maWc6IHtcbiAgICAgIC8qKlxuICAgICAgICogV2hldGhlciB0byB1c2UgaGlzdG9yeS5cbiAgICAgICAqIEB0eXBlIGJvb2xlYW5cbiAgICAgICAqIEBkZWZhdWx0IHRydWVcbiAgICAgICAqL1xuICAgICAgaGlzdG9yeUVuYWJsZWQ6IHRydWUsXG4gICAgICAvKipcbiAgICAgICAqIFRoZSBudW1iZXIgb2YgcGFnZXMgdG8ga2VlcCBpbiAqKmxvY2FsU3RvcmFnZSoqIGZvciBoaXN0b3J5IHN1cHBvcnQuXG4gICAgICAgKiBAdHlwZSBudW1iZXJcbiAgICAgICAqIEBkZWZhdWx0IDEwXG4gICAgICAgKi9cbiAgICAgIGhpc3RvcnlDYWNoZVNpemU6IDEwLFxuICAgICAgLyoqXG4gICAgICAgKiBAdHlwZSBib29sZWFuXG4gICAgICAgKiBAZGVmYXVsdCBmYWxzZVxuICAgICAgICovXG4gICAgICByZWZyZXNoT25IaXN0b3J5TWlzczogZmFsc2UsXG4gICAgICAvKipcbiAgICAgICAqIFRoZSBkZWZhdWx0IHN3YXAgc3R5bGUgdG8gdXNlIGlmICoqW2h4LXN3YXBdKGh0dHBzOi8vaHRteC5vcmcvYXR0cmlidXRlcy9oeC1zd2FwKSoqIGlzIG9taXR0ZWQuXG4gICAgICAgKiBAdHlwZSBIdG14U3dhcFN0eWxlXG4gICAgICAgKiBAZGVmYXVsdCAnaW5uZXJIVE1MJ1xuICAgICAgICovXG4gICAgICBkZWZhdWx0U3dhcFN0eWxlOiAnaW5uZXJIVE1MJyxcbiAgICAgIC8qKlxuICAgICAgICogVGhlIGRlZmF1bHQgZGVsYXkgYmV0d2VlbiByZWNlaXZpbmcgYSByZXNwb25zZSBmcm9tIHRoZSBzZXJ2ZXIgYW5kIGRvaW5nIHRoZSBzd2FwLlxuICAgICAgICogQHR5cGUgbnVtYmVyXG4gICAgICAgKiBAZGVmYXVsdCAwXG4gICAgICAgKi9cbiAgICAgIGRlZmF1bHRTd2FwRGVsYXk6IDAsXG4gICAgICAvKipcbiAgICAgICAqIFRoZSBkZWZhdWx0IGRlbGF5IGJldHdlZW4gY29tcGxldGluZyB0aGUgY29udGVudCBzd2FwIGFuZCBzZXR0bGluZyBhdHRyaWJ1dGVzLlxuICAgICAgICogQHR5cGUgbnVtYmVyXG4gICAgICAgKiBAZGVmYXVsdCAyMFxuICAgICAgICovXG4gICAgICBkZWZhdWx0U2V0dGxlRGVsYXk6IDIwLFxuICAgICAgLyoqXG4gICAgICAgKiBJZiB0cnVlLCBodG14IHdpbGwgaW5qZWN0IGEgc21hbGwgYW1vdW50IG9mIENTUyBpbnRvIHRoZSBwYWdlIHRvIG1ha2UgaW5kaWNhdG9ycyBpbnZpc2libGUgdW5sZXNzIHRoZSAqKmh0bXgtaW5kaWNhdG9yKiogY2xhc3MgaXMgcHJlc2VudC5cbiAgICAgICAqIEB0eXBlIGJvb2xlYW5cbiAgICAgICAqIEBkZWZhdWx0IHRydWVcbiAgICAgICAqL1xuICAgICAgaW5jbHVkZUluZGljYXRvclN0eWxlczogdHJ1ZSxcbiAgICAgIC8qKlxuICAgICAgICogVGhlIGNsYXNzIHRvIHBsYWNlIG9uIGluZGljYXRvcnMgd2hlbiBhIHJlcXVlc3QgaXMgaW4gZmxpZ2h0LlxuICAgICAgICogQHR5cGUgc3RyaW5nXG4gICAgICAgKiBAZGVmYXVsdCAnaHRteC1pbmRpY2F0b3InXG4gICAgICAgKi9cbiAgICAgIGluZGljYXRvckNsYXNzOiAnaHRteC1pbmRpY2F0b3InLFxuICAgICAgLyoqXG4gICAgICAgKiBUaGUgY2xhc3MgdG8gcGxhY2Ugb24gdHJpZ2dlcmluZyBlbGVtZW50cyB3aGVuIGEgcmVxdWVzdCBpcyBpbiBmbGlnaHQuXG4gICAgICAgKiBAdHlwZSBzdHJpbmdcbiAgICAgICAqIEBkZWZhdWx0ICdodG14LXJlcXVlc3QnXG4gICAgICAgKi9cbiAgICAgIHJlcXVlc3RDbGFzczogJ2h0bXgtcmVxdWVzdCcsXG4gICAgICAvKipcbiAgICAgICAqIFRoZSBjbGFzcyB0byB0ZW1wb3JhcmlseSBwbGFjZSBvbiBlbGVtZW50cyB0aGF0IGh0bXggaGFzIGFkZGVkIHRvIHRoZSBET00uXG4gICAgICAgKiBAdHlwZSBzdHJpbmdcbiAgICAgICAqIEBkZWZhdWx0ICdodG14LWFkZGVkJ1xuICAgICAgICovXG4gICAgICBhZGRlZENsYXNzOiAnaHRteC1hZGRlZCcsXG4gICAgICAvKipcbiAgICAgICAqIFRoZSBjbGFzcyB0byBwbGFjZSBvbiB0YXJnZXQgZWxlbWVudHMgd2hlbiBodG14IGlzIGluIHRoZSBzZXR0bGluZyBwaGFzZS5cbiAgICAgICAqIEB0eXBlIHN0cmluZ1xuICAgICAgICogQGRlZmF1bHQgJ2h0bXgtc2V0dGxpbmcnXG4gICAgICAgKi9cbiAgICAgIHNldHRsaW5nQ2xhc3M6ICdodG14LXNldHRsaW5nJyxcbiAgICAgIC8qKlxuICAgICAgICogVGhlIGNsYXNzIHRvIHBsYWNlIG9uIHRhcmdldCBlbGVtZW50cyB3aGVuIGh0bXggaXMgaW4gdGhlIHN3YXBwaW5nIHBoYXNlLlxuICAgICAgICogQHR5cGUgc3RyaW5nXG4gICAgICAgKiBAZGVmYXVsdCAnaHRteC1zd2FwcGluZydcbiAgICAgICAqL1xuICAgICAgc3dhcHBpbmdDbGFzczogJ2h0bXgtc3dhcHBpbmcnLFxuICAgICAgLyoqXG4gICAgICAgKiBBbGxvd3MgdGhlIHVzZSBvZiBldmFsLWxpa2UgZnVuY3Rpb25hbGl0eSBpbiBodG14LCB0byBlbmFibGUgKipoeC12YXJzKiosIHRyaWdnZXIgY29uZGl0aW9ucyAmIHNjcmlwdCB0YWcgZXZhbHVhdGlvbi4gQ2FuIGJlIHNldCB0byAqKmZhbHNlKiogZm9yIENTUCBjb21wYXRpYmlsaXR5LlxuICAgICAgICogQHR5cGUgYm9vbGVhblxuICAgICAgICogQGRlZmF1bHQgdHJ1ZVxuICAgICAgICovXG4gICAgICBhbGxvd0V2YWw6IHRydWUsXG4gICAgICAvKipcbiAgICAgICAqIElmIHNldCB0byBmYWxzZSwgZGlzYWJsZXMgdGhlIGludGVycHJldGF0aW9uIG9mIHNjcmlwdCB0YWdzLlxuICAgICAgICogQHR5cGUgYm9vbGVhblxuICAgICAgICogQGRlZmF1bHQgdHJ1ZVxuICAgICAgICovXG4gICAgICBhbGxvd1NjcmlwdFRhZ3M6IHRydWUsXG4gICAgICAvKipcbiAgICAgICAqIElmIHNldCwgdGhlIG5vbmNlIHdpbGwgYmUgYWRkZWQgdG8gaW5saW5lIHNjcmlwdHMuXG4gICAgICAgKiBAdHlwZSBzdHJpbmdcbiAgICAgICAqIEBkZWZhdWx0ICcnXG4gICAgICAgKi9cbiAgICAgIGlubGluZVNjcmlwdE5vbmNlOiAnJyxcbiAgICAgIC8qKlxuICAgICAgICogSWYgc2V0LCB0aGUgbm9uY2Ugd2lsbCBiZSBhZGRlZCB0byBpbmxpbmUgc3R5bGVzLlxuICAgICAgICogQHR5cGUgc3RyaW5nXG4gICAgICAgKiBAZGVmYXVsdCAnJ1xuICAgICAgICovXG4gICAgICBpbmxpbmVTdHlsZU5vbmNlOiAnJyxcbiAgICAgIC8qKlxuICAgICAgICogVGhlIGF0dHJpYnV0ZXMgdG8gc2V0dGxlIGR1cmluZyB0aGUgc2V0dGxpbmcgcGhhc2UuXG4gICAgICAgKiBAdHlwZSBzdHJpbmdbXVxuICAgICAgICogQGRlZmF1bHQgWydjbGFzcycsICdzdHlsZScsICd3aWR0aCcsICdoZWlnaHQnXVxuICAgICAgICovXG4gICAgICBhdHRyaWJ1dGVzVG9TZXR0bGU6IFsnY2xhc3MnLCAnc3R5bGUnLCAnd2lkdGgnLCAnaGVpZ2h0J10sXG4gICAgICAvKipcbiAgICAgICAqIEFsbG93IGNyb3NzLXNpdGUgQWNjZXNzLUNvbnRyb2wgcmVxdWVzdHMgdXNpbmcgY3JlZGVudGlhbHMgc3VjaCBhcyBjb29raWVzLCBhdXRob3JpemF0aW9uIGhlYWRlcnMgb3IgVExTIGNsaWVudCBjZXJ0aWZpY2F0ZXMuXG4gICAgICAgKiBAdHlwZSBib29sZWFuXG4gICAgICAgKiBAZGVmYXVsdCBmYWxzZVxuICAgICAgICovXG4gICAgICB3aXRoQ3JlZGVudGlhbHM6IGZhbHNlLFxuICAgICAgLyoqXG4gICAgICAgKiBAdHlwZSBudW1iZXJcbiAgICAgICAqIEBkZWZhdWx0IDBcbiAgICAgICAqL1xuICAgICAgdGltZW91dDogMCxcbiAgICAgIC8qKlxuICAgICAgICogVGhlIGRlZmF1bHQgaW1wbGVtZW50YXRpb24gb2YgKipnZXRXZWJTb2NrZXRSZWNvbm5lY3REZWxheSoqIGZvciByZWNvbm5lY3RpbmcgYWZ0ZXIgdW5leHBlY3RlZCBjb25uZWN0aW9uIGxvc3MgYnkgdGhlIGV2ZW50IGNvZGUgKipBYm5vcm1hbCBDbG9zdXJlKiosICoqU2VydmljZSBSZXN0YXJ0Kiogb3IgKipUcnkgQWdhaW4gTGF0ZXIqKi5cbiAgICAgICAqIEB0eXBlIHsnZnVsbC1qaXR0ZXInIHwgKChyZXRyeUNvdW50Om51bWJlcikgPT4gbnVtYmVyKX1cbiAgICAgICAqIEBkZWZhdWx0IFwiZnVsbC1qaXR0ZXJcIlxuICAgICAgICovXG4gICAgICB3c1JlY29ubmVjdERlbGF5OiAnZnVsbC1qaXR0ZXInLFxuICAgICAgLyoqXG4gICAgICAgKiBUaGUgdHlwZSBvZiBiaW5hcnkgZGF0YSBiZWluZyByZWNlaXZlZCBvdmVyIHRoZSBXZWJTb2NrZXQgY29ubmVjdGlvblxuICAgICAgICogQHR5cGUgQmluYXJ5VHlwZVxuICAgICAgICogQGRlZmF1bHQgJ2Jsb2InXG4gICAgICAgKi9cbiAgICAgIHdzQmluYXJ5VHlwZTogJ2Jsb2InLFxuICAgICAgLyoqXG4gICAgICAgKiBAdHlwZSBzdHJpbmdcbiAgICAgICAqIEBkZWZhdWx0ICdbaHgtZGlzYWJsZV0sIFtkYXRhLWh4LWRpc2FibGVdJ1xuICAgICAgICovXG4gICAgICBkaXNhYmxlU2VsZWN0b3I6ICdbaHgtZGlzYWJsZV0sIFtkYXRhLWh4LWRpc2FibGVdJyxcbiAgICAgIC8qKlxuICAgICAgICogQHR5cGUgeydhdXRvJyB8ICdpbnN0YW50JyB8ICdzbW9vdGgnfVxuICAgICAgICogQGRlZmF1bHQgJ2luc3RhbnQnXG4gICAgICAgKi9cbiAgICAgIHNjcm9sbEJlaGF2aW9yOiAnaW5zdGFudCcsXG4gICAgICAvKipcbiAgICAgICAqIElmIHRoZSBmb2N1c2VkIGVsZW1lbnQgc2hvdWxkIGJlIHNjcm9sbGVkIGludG8gdmlldy5cbiAgICAgICAqIEB0eXBlIGJvb2xlYW5cbiAgICAgICAqIEBkZWZhdWx0IGZhbHNlXG4gICAgICAgKi9cbiAgICAgIGRlZmF1bHRGb2N1c1Njcm9sbDogZmFsc2UsXG4gICAgICAvKipcbiAgICAgICAqIElmIHNldCB0byB0cnVlIGh0bXggd2lsbCBpbmNsdWRlIGEgY2FjaGUtYnVzdGluZyBwYXJhbWV0ZXIgaW4gR0VUIHJlcXVlc3RzIHRvIGF2b2lkIGNhY2hpbmcgcGFydGlhbCByZXNwb25zZXMgYnkgdGhlIGJyb3dzZXJcbiAgICAgICAqIEB0eXBlIGJvb2xlYW5cbiAgICAgICAqIEBkZWZhdWx0IGZhbHNlXG4gICAgICAgKi9cbiAgICAgIGdldENhY2hlQnVzdGVyUGFyYW06IGZhbHNlLFxuICAgICAgLyoqXG4gICAgICAgKiBJZiBzZXQgdG8gdHJ1ZSwgaHRteCB3aWxsIHVzZSB0aGUgVmlldyBUcmFuc2l0aW9uIEFQSSB3aGVuIHN3YXBwaW5nIGluIG5ldyBjb250ZW50LlxuICAgICAgICogQHR5cGUgYm9vbGVhblxuICAgICAgICogQGRlZmF1bHQgZmFsc2VcbiAgICAgICAqL1xuICAgICAgZ2xvYmFsVmlld1RyYW5zaXRpb25zOiBmYWxzZSxcbiAgICAgIC8qKlxuICAgICAgICogaHRteCB3aWxsIGZvcm1hdCByZXF1ZXN0cyB3aXRoIHRoZXNlIG1ldGhvZHMgYnkgZW5jb2RpbmcgdGhlaXIgcGFyYW1ldGVycyBpbiB0aGUgVVJMLCBub3QgdGhlIHJlcXVlc3QgYm9keVxuICAgICAgICogQHR5cGUgeyhIdHRwVmVyYilbXX1cbiAgICAgICAqIEBkZWZhdWx0IFsnZ2V0JywgJ2RlbGV0ZSddXG4gICAgICAgKi9cbiAgICAgIG1ldGhvZHNUaGF0VXNlVXJsUGFyYW1zOiBbJ2dldCcsICdkZWxldGUnXSxcbiAgICAgIC8qKlxuICAgICAgICogSWYgc2V0IHRvIHRydWUsIGRpc2FibGVzIGh0bXgtYmFzZWQgcmVxdWVzdHMgdG8gbm9uLW9yaWdpbiBob3N0cy5cbiAgICAgICAqIEB0eXBlIGJvb2xlYW5cbiAgICAgICAqIEBkZWZhdWx0IGZhbHNlXG4gICAgICAgKi9cbiAgICAgIHNlbGZSZXF1ZXN0c09ubHk6IHRydWUsXG4gICAgICAvKipcbiAgICAgICAqIElmIHNldCB0byB0cnVlIGh0bXggd2lsbCBub3QgdXBkYXRlIHRoZSB0aXRsZSBvZiB0aGUgZG9jdW1lbnQgd2hlbiBhIHRpdGxlIHRhZyBpcyBmb3VuZCBpbiBuZXcgY29udGVudFxuICAgICAgICogQHR5cGUgYm9vbGVhblxuICAgICAgICogQGRlZmF1bHQgZmFsc2VcbiAgICAgICAqL1xuICAgICAgaWdub3JlVGl0bGU6IGZhbHNlLFxuICAgICAgLyoqXG4gICAgICAgKiBXaGV0aGVyIHRoZSB0YXJnZXQgb2YgYSBib29zdGVkIGVsZW1lbnQgaXMgc2Nyb2xsZWQgaW50byB0aGUgdmlld3BvcnQuXG4gICAgICAgKiBAdHlwZSBib29sZWFuXG4gICAgICAgKiBAZGVmYXVsdCB0cnVlXG4gICAgICAgKi9cbiAgICAgIHNjcm9sbEludG9WaWV3T25Cb29zdDogdHJ1ZSxcbiAgICAgIC8qKlxuICAgICAgICogVGhlIGNhY2hlIHRvIHN0b3JlIGV2YWx1YXRlZCB0cmlnZ2VyIHNwZWNpZmljYXRpb25zIGludG8uXG4gICAgICAgKiBZb3UgbWF5IGRlZmluZSBhIHNpbXBsZSBvYmplY3QgdG8gdXNlIGEgbmV2ZXItY2xlYXJpbmcgY2FjaGUsIG9yIGltcGxlbWVudCB5b3VyIG93biBzeXN0ZW0gdXNpbmcgYSBbcHJveHkgb2JqZWN0XShodHRwczovL2RldmVsb3Blci5tb3ppbGxhLm9yZy9kb2NzL1dlYi9KYXZhU2NyaXB0L1JlZmVyZW5jZS9HbG9iYWxfT2JqZWN0cy9Qcm94eSlcbiAgICAgICAqIEB0eXBlIHtPYmplY3R8bnVsbH1cbiAgICAgICAqIEBkZWZhdWx0IG51bGxcbiAgICAgICAqL1xuICAgICAgdHJpZ2dlclNwZWNzQ2FjaGU6IG51bGwsXG4gICAgICAvKiogQHR5cGUgYm9vbGVhbiAqL1xuICAgICAgZGlzYWJsZUluaGVyaXRhbmNlOiBmYWxzZSxcbiAgICAgIC8qKiBAdHlwZSBIdG14UmVzcG9uc2VIYW5kbGluZ0NvbmZpZ1tdICovXG4gICAgICByZXNwb25zZUhhbmRsaW5nOiBbXG4gICAgICAgIHsgY29kZTogJzIwNCcsIHN3YXA6IGZhbHNlIH0sXG4gICAgICAgIHsgY29kZTogJ1syM10uLicsIHN3YXA6IHRydWUgfSxcbiAgICAgICAgeyBjb2RlOiAnWzQ1XS4uJywgc3dhcDogZmFsc2UsIGVycm9yOiB0cnVlIH1cbiAgICAgIF0sXG4gICAgICAvKipcbiAgICAgICAqIFdoZXRoZXIgdG8gcHJvY2VzcyBPT0Igc3dhcHMgb24gZWxlbWVudHMgdGhhdCBhcmUgbmVzdGVkIHdpdGhpbiB0aGUgbWFpbiByZXNwb25zZSBlbGVtZW50LlxuICAgICAgICogQHR5cGUgYm9vbGVhblxuICAgICAgICogQGRlZmF1bHQgdHJ1ZVxuICAgICAgICovXG4gICAgICBhbGxvd05lc3RlZE9vYlN3YXBzOiB0cnVlXG4gICAgfSxcbiAgICAvKiogQHR5cGUge3R5cGVvZiBwYXJzZUludGVydmFsfSAqL1xuICAgIHBhcnNlSW50ZXJ2YWw6IG51bGwsXG4gICAgLyoqIEB0eXBlIHt0eXBlb2YgaW50ZXJuYWxFdmFsfSAqL1xuICAgIF86IG51bGwsXG4gICAgdmVyc2lvbjogJzIuMC4yJ1xuICB9XG4gIC8vIFRzYyBtYWRuZXNzIHBhcnQgMlxuICBodG14Lm9uTG9hZCA9IG9uTG9hZEhlbHBlclxuICBodG14LnByb2Nlc3MgPSBwcm9jZXNzTm9kZVxuICBodG14Lm9uID0gYWRkRXZlbnRMaXN0ZW5lckltcGxcbiAgaHRteC5vZmYgPSByZW1vdmVFdmVudExpc3RlbmVySW1wbFxuICBodG14LnRyaWdnZXIgPSB0cmlnZ2VyRXZlbnRcbiAgaHRteC5hamF4ID0gYWpheEhlbHBlclxuICBodG14LmZpbmQgPSBmaW5kXG4gIGh0bXguZmluZEFsbCA9IGZpbmRBbGxcbiAgaHRteC5jbG9zZXN0ID0gY2xvc2VzdFxuICBodG14LnJlbW92ZSA9IHJlbW92ZUVsZW1lbnRcbiAgaHRteC5hZGRDbGFzcyA9IGFkZENsYXNzVG9FbGVtZW50XG4gIGh0bXgucmVtb3ZlQ2xhc3MgPSByZW1vdmVDbGFzc0Zyb21FbGVtZW50XG4gIGh0bXgudG9nZ2xlQ2xhc3MgPSB0b2dnbGVDbGFzc09uRWxlbWVudFxuICBodG14LnRha2VDbGFzcyA9IHRha2VDbGFzc0ZvckVsZW1lbnRcbiAgaHRteC5zd2FwID0gc3dhcFxuICBodG14LmRlZmluZUV4dGVuc2lvbiA9IGRlZmluZUV4dGVuc2lvblxuICBodG14LnJlbW92ZUV4dGVuc2lvbiA9IHJlbW92ZUV4dGVuc2lvblxuICBodG14LmxvZ0FsbCA9IGxvZ0FsbFxuICBodG14LmxvZ05vbmUgPSBsb2dOb25lXG4gIGh0bXgucGFyc2VJbnRlcnZhbCA9IHBhcnNlSW50ZXJ2YWxcbiAgaHRteC5fID0gaW50ZXJuYWxFdmFsXG5cbiAgY29uc3QgaW50ZXJuYWxBUEkgPSB7XG4gICAgYWRkVHJpZ2dlckhhbmRsZXIsXG4gICAgYm9keUNvbnRhaW5zLFxuICAgIGNhbkFjY2Vzc0xvY2FsU3RvcmFnZSxcbiAgICBmaW5kVGhpc0VsZW1lbnQsXG4gICAgZmlsdGVyVmFsdWVzLFxuICAgIHN3YXAsXG4gICAgaGFzQXR0cmlidXRlLFxuICAgIGdldEF0dHJpYnV0ZVZhbHVlLFxuICAgIGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZSxcbiAgICBnZXRDbG9zZXN0TWF0Y2gsXG4gICAgZ2V0RXhwcmVzc2lvblZhcnMsXG4gICAgZ2V0SGVhZGVycyxcbiAgICBnZXRJbnB1dFZhbHVlcyxcbiAgICBnZXRJbnRlcm5hbERhdGEsXG4gICAgZ2V0U3dhcFNwZWNpZmljYXRpb24sXG4gICAgZ2V0VHJpZ2dlclNwZWNzLFxuICAgIGdldFRhcmdldCxcbiAgICBtYWtlRnJhZ21lbnQsXG4gICAgbWVyZ2VPYmplY3RzLFxuICAgIG1ha2VTZXR0bGVJbmZvLFxuICAgIG9vYlN3YXAsXG4gICAgcXVlcnlTZWxlY3RvckV4dCxcbiAgICBzZXR0bGVJbW1lZGlhdGVseSxcbiAgICBzaG91bGRDYW5jZWwsXG4gICAgdHJpZ2dlckV2ZW50LFxuICAgIHRyaWdnZXJFcnJvckV2ZW50LFxuICAgIHdpdGhFeHRlbnNpb25zXG4gIH1cblxuICBjb25zdCBWRVJCUyA9IFsnZ2V0JywgJ3Bvc3QnLCAncHV0JywgJ2RlbGV0ZScsICdwYXRjaCddXG4gIGNvbnN0IFZFUkJfU0VMRUNUT1IgPSBWRVJCUy5tYXAoZnVuY3Rpb24odmVyYikge1xuICAgIHJldHVybiAnW2h4LScgKyB2ZXJiICsgJ10sIFtkYXRhLWh4LScgKyB2ZXJiICsgJ10nXG4gIH0pLmpvaW4oJywgJylcblxuICBjb25zdCBIRUFEX1RBR19SRUdFWCA9IG1ha2VUYWdSZWdFeCgnaGVhZCcpXG5cbiAgLy89ID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT1cbiAgLy8gVXRpbGl0aWVzXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSB0YWdcbiAgICogQHBhcmFtIHtib29sZWFufSBnbG9iYWxcbiAgICogQHJldHVybnMge1JlZ0V4cH1cbiAgICovXG4gIGZ1bmN0aW9uIG1ha2VUYWdSZWdFeCh0YWcsIGdsb2JhbCA9IGZhbHNlKSB7XG4gICAgcmV0dXJuIG5ldyBSZWdFeHAoYDwke3RhZ30oXFxcXHNbXj5dKj58PikoW1xcXFxzXFxcXFNdKj8pPFxcXFwvJHt0YWd9PmAsXG4gICAgICBnbG9iYWwgPyAnZ2ltJyA6ICdpbScpXG4gIH1cblxuICAvKipcbiAgICogUGFyc2VzIGFuIGludGVydmFsIHN0cmluZyBjb25zaXN0ZW50IHdpdGggdGhlIHdheSBodG14IGRvZXMuIFVzZWZ1bCBmb3IgcGx1Z2lucyB0aGF0IGhhdmUgdGltaW5nLXJlbGF0ZWQgYXR0cmlidXRlcy5cbiAgICpcbiAgICogQ2F1dGlvbjogQWNjZXB0cyBhbiBpbnQgZm9sbG93ZWQgYnkgZWl0aGVyICoqcyoqIG9yICoqbXMqKi4gQWxsIG90aGVyIHZhbHVlcyB1c2UgKipwYXJzZUZsb2F0KipcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jcGFyc2VJbnRlcnZhbFxuICAgKlxuICAgKiBAcGFyYW0ge3N0cmluZ30gc3RyIHRpbWluZyBzdHJpbmdcbiAgICogQHJldHVybnMge251bWJlcnx1bmRlZmluZWR9XG4gICAqL1xuICBmdW5jdGlvbiBwYXJzZUludGVydmFsKHN0cikge1xuICAgIGlmIChzdHIgPT0gdW5kZWZpbmVkKSB7XG4gICAgICByZXR1cm4gdW5kZWZpbmVkXG4gICAgfVxuXG4gICAgbGV0IGludGVydmFsID0gTmFOXG4gICAgaWYgKHN0ci5zbGljZSgtMikgPT0gJ21zJykge1xuICAgICAgaW50ZXJ2YWwgPSBwYXJzZUZsb2F0KHN0ci5zbGljZSgwLCAtMikpXG4gICAgfSBlbHNlIGlmIChzdHIuc2xpY2UoLTEpID09ICdzJykge1xuICAgICAgaW50ZXJ2YWwgPSBwYXJzZUZsb2F0KHN0ci5zbGljZSgwLCAtMSkpICogMTAwMFxuICAgIH0gZWxzZSBpZiAoc3RyLnNsaWNlKC0xKSA9PSAnbScpIHtcbiAgICAgIGludGVydmFsID0gcGFyc2VGbG9hdChzdHIuc2xpY2UoMCwgLTEpKSAqIDEwMDAgKiA2MFxuICAgIH0gZWxzZSB7XG4gICAgICBpbnRlcnZhbCA9IHBhcnNlRmxvYXQoc3RyKVxuICAgIH1cbiAgICByZXR1cm4gaXNOYU4oaW50ZXJ2YWwpID8gdW5kZWZpbmVkIDogaW50ZXJ2YWxcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ30gbmFtZVxuICAgKiBAcmV0dXJucyB7KHN0cmluZyB8IG51bGwpfVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0UmF3QXR0cmlidXRlKGVsdCwgbmFtZSkge1xuICAgIHJldHVybiBlbHQgaW5zdGFuY2VvZiBFbGVtZW50ICYmIGVsdC5nZXRBdHRyaWJ1dGUobmFtZSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ30gcXVhbGlmaWVkTmFtZVxuICAgKiBAcmV0dXJucyB7Ym9vbGVhbn1cbiAgICovXG4gIC8vIHJlc29sdmUgd2l0aCBib3RoIGh4IGFuZCBkYXRhLWh4IHByZWZpeGVzXG4gIGZ1bmN0aW9uIGhhc0F0dHJpYnV0ZShlbHQsIHF1YWxpZmllZE5hbWUpIHtcbiAgICByZXR1cm4gISFlbHQuaGFzQXR0cmlidXRlICYmIChlbHQuaGFzQXR0cmlidXRlKHF1YWxpZmllZE5hbWUpIHx8XG4gICAgICBlbHQuaGFzQXR0cmlidXRlKCdkYXRhLScgKyBxdWFsaWZpZWROYW1lKSlcbiAgfVxuXG4gIC8qKlxuICAgKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ30gcXVhbGlmaWVkTmFtZVxuICAgKiBAcmV0dXJucyB7KHN0cmluZyB8IG51bGwpfVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0QXR0cmlidXRlVmFsdWUoZWx0LCBxdWFsaWZpZWROYW1lKSB7XG4gICAgcmV0dXJuIGdldFJhd0F0dHJpYnV0ZShlbHQsIHF1YWxpZmllZE5hbWUpIHx8IGdldFJhd0F0dHJpYnV0ZShlbHQsICdkYXRhLScgKyBxdWFsaWZpZWROYW1lKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gZWx0XG4gICAqIEByZXR1cm5zIHtOb2RlIHwgbnVsbH1cbiAgICovXG4gIGZ1bmN0aW9uIHBhcmVudEVsdChlbHQpIHtcbiAgICBjb25zdCBwYXJlbnQgPSBlbHQucGFyZW50RWxlbWVudFxuICAgIGlmICghcGFyZW50ICYmIGVsdC5wYXJlbnROb2RlIGluc3RhbmNlb2YgU2hhZG93Um9vdCkgcmV0dXJuIGVsdC5wYXJlbnROb2RlXG4gICAgcmV0dXJuIHBhcmVudFxuICB9XG5cbiAgLyoqXG4gICAqIEByZXR1cm5zIHtEb2N1bWVudH1cbiAgICovXG4gIGZ1bmN0aW9uIGdldERvY3VtZW50KCkge1xuICAgIHJldHVybiBkb2N1bWVudFxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gZWx0XG4gICAqIEBwYXJhbSB7Ym9vbGVhbn0gZ2xvYmFsXG4gICAqIEByZXR1cm5zIHtOb2RlfERvY3VtZW50fVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0Um9vdE5vZGUoZWx0LCBnbG9iYWwpIHtcbiAgICByZXR1cm4gZWx0LmdldFJvb3ROb2RlID8gZWx0LmdldFJvb3ROb2RlKHsgY29tcG9zZWQ6IGdsb2JhbCB9KSA6IGdldERvY3VtZW50KClcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKiBAcGFyYW0geyhlOk5vZGUpID0+IGJvb2xlYW59IGNvbmRpdGlvblxuICAgKiBAcmV0dXJucyB7Tm9kZSB8IG51bGx9XG4gICAqL1xuICBmdW5jdGlvbiBnZXRDbG9zZXN0TWF0Y2goZWx0LCBjb25kaXRpb24pIHtcbiAgICB3aGlsZSAoZWx0ICYmICFjb25kaXRpb24oZWx0KSkge1xuICAgICAgZWx0ID0gcGFyZW50RWx0KGVsdClcbiAgICB9XG5cbiAgICByZXR1cm4gZWx0IHx8IG51bGxcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGluaXRpYWxFbGVtZW50XG4gICAqIEBwYXJhbSB7RWxlbWVudH0gYW5jZXN0b3JcbiAgICogQHBhcmFtIHtzdHJpbmd9IGF0dHJpYnV0ZU5hbWVcbiAgICogQHJldHVybnMge3N0cmluZ3xudWxsfVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0QXR0cmlidXRlVmFsdWVXaXRoRGlzaW5oZXJpdGFuY2UoaW5pdGlhbEVsZW1lbnQsIGFuY2VzdG9yLCBhdHRyaWJ1dGVOYW1lKSB7XG4gICAgY29uc3QgYXR0cmlidXRlVmFsdWUgPSBnZXRBdHRyaWJ1dGVWYWx1ZShhbmNlc3RvciwgYXR0cmlidXRlTmFtZSlcbiAgICBjb25zdCBkaXNpbmhlcml0ID0gZ2V0QXR0cmlidXRlVmFsdWUoYW5jZXN0b3IsICdoeC1kaXNpbmhlcml0JylcbiAgICB2YXIgaW5oZXJpdCA9IGdldEF0dHJpYnV0ZVZhbHVlKGFuY2VzdG9yLCAnaHgtaW5oZXJpdCcpXG4gICAgaWYgKGluaXRpYWxFbGVtZW50ICE9PSBhbmNlc3Rvcikge1xuICAgICAgaWYgKGh0bXguY29uZmlnLmRpc2FibGVJbmhlcml0YW5jZSkge1xuICAgICAgICBpZiAoaW5oZXJpdCAmJiAoaW5oZXJpdCA9PT0gJyonIHx8IGluaGVyaXQuc3BsaXQoJyAnKS5pbmRleE9mKGF0dHJpYnV0ZU5hbWUpID49IDApKSB7XG4gICAgICAgICAgcmV0dXJuIGF0dHJpYnV0ZVZhbHVlXG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgcmV0dXJuIG51bGxcbiAgICAgICAgfVxuICAgICAgfVxuICAgICAgaWYgKGRpc2luaGVyaXQgJiYgKGRpc2luaGVyaXQgPT09ICcqJyB8fCBkaXNpbmhlcml0LnNwbGl0KCcgJykuaW5kZXhPZihhdHRyaWJ1dGVOYW1lKSA+PSAwKSkge1xuICAgICAgICByZXR1cm4gJ3Vuc2V0J1xuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gYXR0cmlidXRlVmFsdWVcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ30gYXR0cmlidXRlTmFtZVxuICAgKiBAcmV0dXJucyB7c3RyaW5nIHwgbnVsbH1cbiAgICovXG4gIGZ1bmN0aW9uIGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZShlbHQsIGF0dHJpYnV0ZU5hbWUpIHtcbiAgICBsZXQgY2xvc2VzdEF0dHIgPSBudWxsXG4gICAgZ2V0Q2xvc2VzdE1hdGNoKGVsdCwgZnVuY3Rpb24oZSkge1xuICAgICAgcmV0dXJuICEhKGNsb3Nlc3RBdHRyID0gZ2V0QXR0cmlidXRlVmFsdWVXaXRoRGlzaW5oZXJpdGFuY2UoZWx0LCBhc0VsZW1lbnQoZSksIGF0dHJpYnV0ZU5hbWUpKVxuICAgIH0pXG4gICAgaWYgKGNsb3Nlc3RBdHRyICE9PSAndW5zZXQnKSB7XG4gICAgICByZXR1cm4gY2xvc2VzdEF0dHJcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfSBlbHRcbiAgICogQHBhcmFtIHtzdHJpbmd9IHNlbGVjdG9yXG4gICAqIEByZXR1cm5zIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gbWF0Y2hlcyhlbHQsIHNlbGVjdG9yKSB7XG4gICAgLy8gQHRzLWlnbm9yZTogbm9uLXN0YW5kYXJkIHByb3BlcnRpZXMgZm9yIGJyb3dzZXIgY29tcGF0aWJpbGl0eVxuICAgIC8vIG5vaW5zcGVjdGlvbiBKU1VucmVzb2x2ZWRWYXJpYWJsZVxuICAgIGNvbnN0IG1hdGNoZXNGdW5jdGlvbiA9IGVsdCBpbnN0YW5jZW9mIEVsZW1lbnQgJiYgKGVsdC5tYXRjaGVzIHx8IGVsdC5tYXRjaGVzU2VsZWN0b3IgfHwgZWx0Lm1zTWF0Y2hlc1NlbGVjdG9yIHx8IGVsdC5tb3pNYXRjaGVzU2VsZWN0b3IgfHwgZWx0LndlYmtpdE1hdGNoZXNTZWxlY3RvciB8fCBlbHQub01hdGNoZXNTZWxlY3RvcilcbiAgICByZXR1cm4gISFtYXRjaGVzRnVuY3Rpb24gJiYgbWF0Y2hlc0Z1bmN0aW9uLmNhbGwoZWx0LCBzZWxlY3RvcilcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge3N0cmluZ30gc3RyXG4gICAqIEByZXR1cm5zIHtzdHJpbmd9XG4gICAqL1xuICBmdW5jdGlvbiBnZXRTdGFydFRhZyhzdHIpIHtcbiAgICBjb25zdCB0YWdNYXRjaGVyID0gLzwoW2Etel1bXlxcL1xcMD5cXHgyMFxcdFxcclxcblxcZl0qKS9pXG4gICAgY29uc3QgbWF0Y2ggPSB0YWdNYXRjaGVyLmV4ZWMoc3RyKVxuICAgIGlmIChtYXRjaCkge1xuICAgICAgcmV0dXJuIG1hdGNoWzFdLnRvTG93ZXJDYXNlKClcbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuICcnXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSByZXNwXG4gICAqIEByZXR1cm5zIHtEb2N1bWVudH1cbiAgICovXG4gIGZ1bmN0aW9uIHBhcnNlSFRNTChyZXNwKSB7XG4gICAgY29uc3QgcGFyc2VyID0gbmV3IERPTVBhcnNlcigpXG4gICAgcmV0dXJuIHBhcnNlci5wYXJzZUZyb21TdHJpbmcocmVzcCwgJ3RleHQvaHRtbCcpXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtEb2N1bWVudEZyYWdtZW50fSBmcmFnbWVudFxuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKi9cbiAgZnVuY3Rpb24gdGFrZUNoaWxkcmVuRm9yKGZyYWdtZW50LCBlbHQpIHtcbiAgICB3aGlsZSAoZWx0LmNoaWxkTm9kZXMubGVuZ3RoID4gMCkge1xuICAgICAgZnJhZ21lbnQuYXBwZW5kKGVsdC5jaGlsZE5vZGVzWzBdKVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0hUTUxTY3JpcHRFbGVtZW50fSBzY3JpcHRcbiAgICogQHJldHVybnMge0hUTUxTY3JpcHRFbGVtZW50fVxuICAgKi9cbiAgZnVuY3Rpb24gZHVwbGljYXRlU2NyaXB0KHNjcmlwdCkge1xuICAgIGNvbnN0IG5ld1NjcmlwdCA9IGdldERvY3VtZW50KCkuY3JlYXRlRWxlbWVudCgnc2NyaXB0JylcbiAgICBmb3JFYWNoKHNjcmlwdC5hdHRyaWJ1dGVzLCBmdW5jdGlvbihhdHRyKSB7XG4gICAgICBuZXdTY3JpcHQuc2V0QXR0cmlidXRlKGF0dHIubmFtZSwgYXR0ci52YWx1ZSlcbiAgICB9KVxuICAgIG5ld1NjcmlwdC50ZXh0Q29udGVudCA9IHNjcmlwdC50ZXh0Q29udGVudFxuICAgIG5ld1NjcmlwdC5hc3luYyA9IGZhbHNlXG4gICAgaWYgKGh0bXguY29uZmlnLmlubGluZVNjcmlwdE5vbmNlKSB7XG4gICAgICBuZXdTY3JpcHQubm9uY2UgPSBodG14LmNvbmZpZy5pbmxpbmVTY3JpcHROb25jZVxuICAgIH1cbiAgICByZXR1cm4gbmV3U2NyaXB0XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtIVE1MU2NyaXB0RWxlbWVudH0gc2NyaXB0XG4gICAqIEByZXR1cm5zIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gaXNKYXZhU2NyaXB0U2NyaXB0Tm9kZShzY3JpcHQpIHtcbiAgICByZXR1cm4gc2NyaXB0Lm1hdGNoZXMoJ3NjcmlwdCcpICYmIChzY3JpcHQudHlwZSA9PT0gJ3RleHQvamF2YXNjcmlwdCcgfHwgc2NyaXB0LnR5cGUgPT09ICdtb2R1bGUnIHx8IHNjcmlwdC50eXBlID09PSAnJylcbiAgfVxuXG4gIC8qKlxuICAgKiB3ZSBoYXZlIHRvIG1ha2UgbmV3IGNvcGllcyBvZiBzY3JpcHQgdGFncyB0aGF0IHdlIGFyZSBnb2luZyB0byBpbnNlcnQgYmVjYXVzZVxuICAgKiBTT01FIGJyb3dzZXJzIChub3Qgc2F5aW5nIHdobywgYnV0IGl0IGludm9sdmVzIGFuIGVsZW1lbnQgYW5kIGFuIGFuaW1hbCkgZG9uJ3RcbiAgICogZXhlY3V0ZSBzY3JpcHRzIGNyZWF0ZWQgaW4gPHRlbXBsYXRlPiB0YWdzIHdoZW4gdGhleSBhcmUgaW5zZXJ0ZWQgaW50byB0aGUgRE9NXG4gICAqIGFuZCBhbGwgdGhlIG90aGVycyBkbyBsbWFvXG4gICAqIEBwYXJhbSB7RG9jdW1lbnRGcmFnbWVudH0gZnJhZ21lbnRcbiAgICovXG4gIGZ1bmN0aW9uIG5vcm1hbGl6ZVNjcmlwdFRhZ3MoZnJhZ21lbnQpIHtcbiAgICBBcnJheS5mcm9tKGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3JBbGwoJ3NjcmlwdCcpKS5mb3JFYWNoKC8qKiBAcGFyYW0ge0hUTUxTY3JpcHRFbGVtZW50fSBzY3JpcHQgKi8gKHNjcmlwdCkgPT4ge1xuICAgICAgaWYgKGlzSmF2YVNjcmlwdFNjcmlwdE5vZGUoc2NyaXB0KSkge1xuICAgICAgICBjb25zdCBuZXdTY3JpcHQgPSBkdXBsaWNhdGVTY3JpcHQoc2NyaXB0KVxuICAgICAgICBjb25zdCBwYXJlbnQgPSBzY3JpcHQucGFyZW50Tm9kZVxuICAgICAgICB0cnkge1xuICAgICAgICAgIHBhcmVudC5pbnNlcnRCZWZvcmUobmV3U2NyaXB0LCBzY3JpcHQpXG4gICAgICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAgICAgICBsb2dFcnJvcihlKVxuICAgICAgICB9IGZpbmFsbHkge1xuICAgICAgICAgIHNjcmlwdC5yZW1vdmUoKVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAdHlwZWRlZiB7RG9jdW1lbnRGcmFnbWVudCAmIHt0aXRsZT86IHN0cmluZ319IERvY3VtZW50RnJhZ21lbnRXaXRoVGl0bGVcbiAgICogQGRlc2NyaXB0aW9uICBhIGRvY3VtZW50IGZyYWdtZW50IHJlcHJlc2VudGluZyB0aGUgcmVzcG9uc2UgSFRNTCwgaW5jbHVkaW5nXG4gICAqIGEgYHRpdGxlYCBwcm9wZXJ0eSBmb3IgYW55IHRpdGxlIGluZm9ybWF0aW9uIGZvdW5kXG4gICAqL1xuXG4gIC8qKlxuICAgKiBAcGFyYW0ge3N0cmluZ30gcmVzcG9uc2UgSFRNTFxuICAgKiBAcmV0dXJucyB7RG9jdW1lbnRGcmFnbWVudFdpdGhUaXRsZX1cbiAgICovXG4gIGZ1bmN0aW9uIG1ha2VGcmFnbWVudChyZXNwb25zZSkge1xuICAgIC8vIHN0cmlwIGhlYWQgdGFnIHRvIGRldGVybWluZSBzaGFwZSBvZiByZXNwb25zZSB3ZSBhcmUgZGVhbGluZyB3aXRoXG4gICAgY29uc3QgcmVzcG9uc2VXaXRoTm9IZWFkID0gcmVzcG9uc2UucmVwbGFjZShIRUFEX1RBR19SRUdFWCwgJycpXG4gICAgY29uc3Qgc3RhcnRUYWcgPSBnZXRTdGFydFRhZyhyZXNwb25zZVdpdGhOb0hlYWQpXG4gICAgLyoqIEB0eXBlIERvY3VtZW50RnJhZ21lbnRXaXRoVGl0bGUgKi9cbiAgICBsZXQgZnJhZ21lbnRcbiAgICBpZiAoc3RhcnRUYWcgPT09ICdodG1sJykge1xuICAgICAgLy8gaWYgaXQgaXMgYSBmdWxsIGRvY3VtZW50LCBwYXJzZSBpdCBhbmQgcmV0dXJuIHRoZSBib2R5XG4gICAgICBmcmFnbWVudCA9IC8qKiBAdHlwZSBEb2N1bWVudEZyYWdtZW50V2l0aFRpdGxlICovIChuZXcgRG9jdW1lbnRGcmFnbWVudCgpKVxuICAgICAgY29uc3QgZG9jID0gcGFyc2VIVE1MKHJlc3BvbnNlKVxuICAgICAgdGFrZUNoaWxkcmVuRm9yKGZyYWdtZW50LCBkb2MuYm9keSlcbiAgICAgIGZyYWdtZW50LnRpdGxlID0gZG9jLnRpdGxlXG4gICAgfSBlbHNlIGlmIChzdGFydFRhZyA9PT0gJ2JvZHknKSB7XG4gICAgICAvLyBwYXJzZSBib2R5IHcvbyB3cmFwcGluZyBpbiB0ZW1wbGF0ZVxuICAgICAgZnJhZ21lbnQgPSAvKiogQHR5cGUgRG9jdW1lbnRGcmFnbWVudFdpdGhUaXRsZSAqLyAobmV3IERvY3VtZW50RnJhZ21lbnQoKSlcbiAgICAgIGNvbnN0IGRvYyA9IHBhcnNlSFRNTChyZXNwb25zZVdpdGhOb0hlYWQpXG4gICAgICB0YWtlQ2hpbGRyZW5Gb3IoZnJhZ21lbnQsIGRvYy5ib2R5KVxuICAgICAgZnJhZ21lbnQudGl0bGUgPSBkb2MudGl0bGVcbiAgICB9IGVsc2Uge1xuICAgICAgLy8gb3RoZXJ3aXNlIHdlIGhhdmUgbm9uLWJvZHkgcGFydGlhbCBIVE1MIGNvbnRlbnQsIHNvIHdyYXAgaXQgaW4gYSB0ZW1wbGF0ZSB0byBtYXhpbWl6ZSBwYXJzaW5nIGZsZXhpYmlsaXR5XG4gICAgICBjb25zdCBkb2MgPSBwYXJzZUhUTUwoJzxib2R5Pjx0ZW1wbGF0ZSBjbGFzcz1cImludGVybmFsLWh0bXgtd3JhcHBlclwiPicgKyByZXNwb25zZVdpdGhOb0hlYWQgKyAnPC90ZW1wbGF0ZT48L2JvZHk+JylcbiAgICAgIGZyYWdtZW50ID0gLyoqIEB0eXBlIERvY3VtZW50RnJhZ21lbnRXaXRoVGl0bGUgKi8gKGRvYy5xdWVyeVNlbGVjdG9yKCd0ZW1wbGF0ZScpLmNvbnRlbnQpXG4gICAgICAvLyBleHRyYWN0IHRpdGxlIGludG8gZnJhZ21lbnQgZm9yIGxhdGVyIHByb2Nlc3NpbmdcbiAgICAgIGZyYWdtZW50LnRpdGxlID0gZG9jLnRpdGxlXG5cbiAgICAgIC8vIGZvciBsZWdhY3kgcmVhc29ucyB3ZSBzdXBwb3J0IGEgdGl0bGUgdGFnIGF0IHRoZSByb290IGxldmVsIG9mIG5vbi1ib2R5IHJlc3BvbnNlcywgc28gd2UgbmVlZCB0byBoYW5kbGUgaXRcbiAgICAgIHZhciB0aXRsZUVsZW1lbnQgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCd0aXRsZScpXG4gICAgICBpZiAodGl0bGVFbGVtZW50ICYmIHRpdGxlRWxlbWVudC5wYXJlbnROb2RlID09PSBmcmFnbWVudCkge1xuICAgICAgICB0aXRsZUVsZW1lbnQucmVtb3ZlKClcbiAgICAgICAgZnJhZ21lbnQudGl0bGUgPSB0aXRsZUVsZW1lbnQuaW5uZXJUZXh0XG4gICAgICB9XG4gICAgfVxuICAgIGlmIChmcmFnbWVudCkge1xuICAgICAgaWYgKGh0bXguY29uZmlnLmFsbG93U2NyaXB0VGFncykge1xuICAgICAgICBub3JtYWxpemVTY3JpcHRUYWdzKGZyYWdtZW50KVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgLy8gcmVtb3ZlIGFsbCBzY3JpcHQgdGFncyBpZiBzY3JpcHRzIGFyZSBkaXNhYmxlZFxuICAgICAgICBmcmFnbWVudC5xdWVyeVNlbGVjdG9yQWxsKCdzY3JpcHQnKS5mb3JFYWNoKChzY3JpcHQpID0+IHNjcmlwdC5yZW1vdmUoKSlcbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuIGZyYWdtZW50XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtGdW5jdGlvbn0gZnVuY1xuICAgKi9cbiAgZnVuY3Rpb24gbWF5YmVDYWxsKGZ1bmMpIHtcbiAgICBpZiAoZnVuYykge1xuICAgICAgZnVuYygpXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7YW55fSBvXG4gICAqIEBwYXJhbSB7c3RyaW5nfSB0eXBlXG4gICAqIEByZXR1cm5zXG4gICAqL1xuICBmdW5jdGlvbiBpc1R5cGUobywgdHlwZSkge1xuICAgIHJldHVybiBPYmplY3QucHJvdG90eXBlLnRvU3RyaW5nLmNhbGwobykgPT09ICdbb2JqZWN0ICcgKyB0eXBlICsgJ10nXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHsqfSBvXG4gICAqIEByZXR1cm5zIHtvIGlzIEZ1bmN0aW9ufVxuICAgKi9cbiAgZnVuY3Rpb24gaXNGdW5jdGlvbihvKSB7XG4gICAgcmV0dXJuIHR5cGVvZiBvID09PSAnZnVuY3Rpb24nXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHsqfSBvXG4gICAqIEByZXR1cm5zIHtvIGlzIE9iamVjdH1cbiAgICovXG4gIGZ1bmN0aW9uIGlzUmF3T2JqZWN0KG8pIHtcbiAgICByZXR1cm4gaXNUeXBlKG8sICdPYmplY3QnKVxuICB9XG5cbiAgLyoqXG4gICAqIEB0eXBlZGVmIHtPYmplY3R9IE9uSGFuZGxlclxuICAgKiBAcHJvcGVydHkgeyhrZXlvZiBIVE1MRWxlbWVudEV2ZW50TWFwKXxzdHJpbmd9IGV2ZW50XG4gICAqIEBwcm9wZXJ0eSB7RXZlbnRMaXN0ZW5lcn0gbGlzdGVuZXJcbiAgICovXG5cbiAgLyoqXG4gICAqIEB0eXBlZGVmIHtPYmplY3R9IExpc3RlbmVySW5mb1xuICAgKiBAcHJvcGVydHkge3N0cmluZ30gdHJpZ2dlclxuICAgKiBAcHJvcGVydHkge0V2ZW50TGlzdGVuZXJ9IGxpc3RlbmVyXG4gICAqIEBwcm9wZXJ0eSB7RXZlbnRUYXJnZXR9IG9uXG4gICAqL1xuXG4gIC8qKlxuICAgKiBAdHlwZWRlZiB7T2JqZWN0fSBIdG14Tm9kZUludGVybmFsRGF0YVxuICAgKiBFbGVtZW50IGRhdGFcbiAgICogQHByb3BlcnR5IHtudW1iZXJ9IFtpbml0SGFzaF1cbiAgICogQHByb3BlcnR5IHtib29sZWFufSBbYm9vc3RlZF1cbiAgICogQHByb3BlcnR5IHtPbkhhbmRsZXJbXX0gW29uSGFuZGxlcnNdXG4gICAqIEBwcm9wZXJ0eSB7bnVtYmVyfSBbdGltZW91dF1cbiAgICogQHByb3BlcnR5IHtMaXN0ZW5lckluZm9bXX0gW2xpc3RlbmVySW5mb3NdXG4gICAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gW2NhbmNlbGxlZF1cbiAgICogQHByb3BlcnR5IHtib29sZWFufSBbdHJpZ2dlcmVkT25jZV1cbiAgICogQHByb3BlcnR5IHtudW1iZXJ9IFtkZWxheWVkXVxuICAgKiBAcHJvcGVydHkge251bWJlcnxudWxsfSBbdGhyb3R0bGVdXG4gICAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBbbGFzdFZhbHVlXVxuICAgKiBAcHJvcGVydHkge2Jvb2xlYW59IFtsb2FkZWRdXG4gICAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBbcGF0aF1cbiAgICogQHByb3BlcnR5IHtzdHJpbmd9IFt2ZXJiXVxuICAgKiBAcHJvcGVydHkge2Jvb2xlYW59IFtwb2xsaW5nXVxuICAgKiBAcHJvcGVydHkge0hUTUxCdXR0b25FbGVtZW50fEhUTUxJbnB1dEVsZW1lbnR8bnVsbH0gW2xhc3RCdXR0b25DbGlja2VkXVxuICAgKiBAcHJvcGVydHkge251bWJlcn0gW3JlcXVlc3RDb3VudF1cbiAgICogQHByb3BlcnR5IHtYTUxIdHRwUmVxdWVzdH0gW3hocl1cbiAgICogQHByb3BlcnR5IHsoKCkgPT4gdm9pZClbXX0gW3F1ZXVlZFJlcXVlc3RzXVxuICAgKiBAcHJvcGVydHkge2Jvb2xlYW59IFthYm9ydGFibGVdXG4gICAqXG4gICAqIEV2ZW50IGRhdGFcbiAgICogQHByb3BlcnR5IHtIdG14VHJpZ2dlclNwZWNpZmljYXRpb259IFt0cmlnZ2VyU3BlY11cbiAgICogQHByb3BlcnR5IHtFdmVudFRhcmdldFtdfSBbaGFuZGxlZEZvcl1cbiAgICovXG5cbiAgLyoqXG4gICAqIGdldEludGVybmFsRGF0YSByZXRyaWV2ZXMgXCJwcml2YXRlXCIgZGF0YSBzdG9yZWQgYnkgaHRteCB3aXRoaW4gYW4gZWxlbWVudFxuICAgKiBAcGFyYW0ge0V2ZW50VGFyZ2V0fEV2ZW50fSBlbHRcbiAgICogQHJldHVybnMge0h0bXhOb2RlSW50ZXJuYWxEYXRhfVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0SW50ZXJuYWxEYXRhKGVsdCkge1xuICAgIGNvbnN0IGRhdGFQcm9wID0gJ2h0bXgtaW50ZXJuYWwtZGF0YSdcbiAgICBsZXQgZGF0YSA9IGVsdFtkYXRhUHJvcF1cbiAgICBpZiAoIWRhdGEpIHtcbiAgICAgIGRhdGEgPSBlbHRbZGF0YVByb3BdID0ge31cbiAgICB9XG4gICAgcmV0dXJuIGRhdGFcbiAgfVxuXG4gIC8qKlxuICAgKiB0b0FycmF5IGNvbnZlcnRzIGFuIEFycmF5TGlrZSBvYmplY3QgaW50byBhIHJlYWwgYXJyYXkuXG4gICAqIEB0ZW1wbGF0ZSBUXG4gICAqIEBwYXJhbSB7QXJyYXlMaWtlPFQ+fSBhcnJcbiAgICogQHJldHVybnMge1RbXX1cbiAgICovXG4gIGZ1bmN0aW9uIHRvQXJyYXkoYXJyKSB7XG4gICAgY29uc3QgcmV0dXJuQXJyID0gW11cbiAgICBpZiAoYXJyKSB7XG4gICAgICBmb3IgKGxldCBpID0gMDsgaSA8IGFyci5sZW5ndGg7IGkrKykge1xuICAgICAgICByZXR1cm5BcnIucHVzaChhcnJbaV0pXG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiByZXR1cm5BcnJcbiAgfVxuXG4gIC8qKlxuICAgKiBAdGVtcGxhdGUgVFxuICAgKiBAcGFyYW0ge1RbXXxOYW1lZE5vZGVNYXB8SFRNTENvbGxlY3Rpb258SFRNTEZvcm1Db250cm9sc0NvbGxlY3Rpb258QXJyYXlMaWtlPFQ+fSBhcnJcbiAgICogQHBhcmFtIHsoVCkgPT4gdm9pZH0gZnVuY1xuICAgKi9cbiAgZnVuY3Rpb24gZm9yRWFjaChhcnIsIGZ1bmMpIHtcbiAgICBpZiAoYXJyKSB7XG4gICAgICBmb3IgKGxldCBpID0gMDsgaSA8IGFyci5sZW5ndGg7IGkrKykge1xuICAgICAgICBmdW5jKGFycltpXSlcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbFxuICAgKiBAcmV0dXJucyB7Ym9vbGVhbn1cbiAgICovXG4gIGZ1bmN0aW9uIGlzU2Nyb2xsZWRJbnRvVmlldyhlbCkge1xuICAgIGNvbnN0IHJlY3QgPSBlbC5nZXRCb3VuZGluZ0NsaWVudFJlY3QoKVxuICAgIGNvbnN0IGVsZW1Ub3AgPSByZWN0LnRvcFxuICAgIGNvbnN0IGVsZW1Cb3R0b20gPSByZWN0LmJvdHRvbVxuICAgIHJldHVybiBlbGVtVG9wIDwgd2luZG93LmlubmVySGVpZ2h0ICYmIGVsZW1Cb3R0b20gPj0gMFxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gZWx0XG4gICAqIEByZXR1cm5zIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gYm9keUNvbnRhaW5zKGVsdCkge1xuICAgIC8vIElFIEZpeFxuICAgIGNvbnN0IHJvb3ROb2RlID0gZWx0LmdldFJvb3ROb2RlICYmIGVsdC5nZXRSb290Tm9kZSgpXG4gICAgaWYgKHJvb3ROb2RlICYmIHJvb3ROb2RlIGluc3RhbmNlb2Ygd2luZG93LlNoYWRvd1Jvb3QpIHtcbiAgICAgIHJldHVybiBnZXREb2N1bWVudCgpLmJvZHkuY29udGFpbnMocm9vdE5vZGUuaG9zdClcbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuIGdldERvY3VtZW50KCkuYm9keS5jb250YWlucyhlbHQpXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSB0cmlnZ2VyXG4gICAqIEByZXR1cm5zIHtzdHJpbmdbXX1cbiAgICovXG4gIGZ1bmN0aW9uIHNwbGl0T25XaGl0ZXNwYWNlKHRyaWdnZXIpIHtcbiAgICByZXR1cm4gdHJpZ2dlci50cmltKCkuc3BsaXQoL1xccysvKVxuICB9XG5cbiAgLyoqXG4gICAqIG1lcmdlT2JqZWN0cyB0YWtlcyBhbGwgdGhlIGtleXMgZnJvbVxuICAgKiBvYmoyIGFuZCBkdXBsaWNhdGVzIHRoZW0gaW50byBvYmoxXG4gICAqIEB0ZW1wbGF0ZSBUMVxuICAgKiBAdGVtcGxhdGUgVDJcbiAgICogQHBhcmFtIHtUMX0gb2JqMVxuICAgKiBAcGFyYW0ge1QyfSBvYmoyXG4gICAqIEByZXR1cm5zIHtUMSAmIFQyfVxuICAgKi9cbiAgZnVuY3Rpb24gbWVyZ2VPYmplY3RzKG9iajEsIG9iajIpIHtcbiAgICBmb3IgKGNvbnN0IGtleSBpbiBvYmoyKSB7XG4gICAgICBpZiAob2JqMi5oYXNPd25Qcm9wZXJ0eShrZXkpKSB7XG4gICAgICAgIC8vIEB0cy1pZ25vcmUgdHNjIGRvZXNuJ3Qgc2VlbSB0byBwcm9wZXJseSBoYW5kbGUgdHlwZXMgbWVyZ2luZ1xuICAgICAgICBvYmoxW2tleV0gPSBvYmoyW2tleV1cbiAgICAgIH1cbiAgICB9XG4gICAgLy8gQHRzLWlnbm9yZSB0c2MgZG9lc24ndCBzZWVtIHRvIHByb3Blcmx5IGhhbmRsZSB0eXBlcyBtZXJnaW5nXG4gICAgcmV0dXJuIG9iajFcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge3N0cmluZ30galN0cmluZ1xuICAgKiBAcmV0dXJucyB7YW55fG51bGx9XG4gICAqL1xuICBmdW5jdGlvbiBwYXJzZUpTT04oalN0cmluZykge1xuICAgIHRyeSB7XG4gICAgICByZXR1cm4gSlNPTi5wYXJzZShqU3RyaW5nKVxuICAgIH0gY2F0Y2ggKGVycm9yKSB7XG4gICAgICBsb2dFcnJvcihlcnJvcilcbiAgICAgIHJldHVybiBudWxsXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEByZXR1cm5zIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gY2FuQWNjZXNzTG9jYWxTdG9yYWdlKCkge1xuICAgIGNvbnN0IHRlc3QgPSAnaHRteDpsb2NhbFN0b3JhZ2VUZXN0J1xuICAgIHRyeSB7XG4gICAgICBsb2NhbFN0b3JhZ2Uuc2V0SXRlbSh0ZXN0LCB0ZXN0KVxuICAgICAgbG9jYWxTdG9yYWdlLnJlbW92ZUl0ZW0odGVzdClcbiAgICAgIHJldHVybiB0cnVlXG4gICAgfSBjYXRjaCAoZSkge1xuICAgICAgcmV0dXJuIGZhbHNlXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBwYXRoXG4gICAqIEByZXR1cm5zIHtzdHJpbmd9XG4gICAqL1xuICBmdW5jdGlvbiBub3JtYWxpemVQYXRoKHBhdGgpIHtcbiAgICB0cnkge1xuICAgICAgY29uc3QgdXJsID0gbmV3IFVSTChwYXRoKVxuICAgICAgaWYgKHVybCkge1xuICAgICAgICBwYXRoID0gdXJsLnBhdGhuYW1lICsgdXJsLnNlYXJjaFxuICAgICAgfVxuICAgICAgLy8gcmVtb3ZlIHRyYWlsaW5nIHNsYXNoLCB1bmxlc3MgaW5kZXggcGFnZVxuICAgICAgaWYgKCEoL15cXC8kLy50ZXN0KHBhdGgpKSkge1xuICAgICAgICBwYXRoID0gcGF0aC5yZXBsYWNlKC9cXC8rJC8sICcnKVxuICAgICAgfVxuICAgICAgcmV0dXJuIHBhdGhcbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBiZSBraW5kIHRvIElFMTEsIHdoaWNoIGRvZXNuJ3Qgc3VwcG9ydCBVUkwoKVxuICAgICAgcmV0dXJuIHBhdGhcbiAgICB9XG4gIH1cblxuICAvLz0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT1cbiAgLy8gcHVibGljIEFQSVxuICAvLz0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IHN0clxuICAgKiBAcmV0dXJucyB7YW55fVxuICAgKi9cbiAgZnVuY3Rpb24gaW50ZXJuYWxFdmFsKHN0cikge1xuICAgIHJldHVybiBtYXliZUV2YWwoZ2V0RG9jdW1lbnQoKS5ib2R5LCBmdW5jdGlvbigpIHtcbiAgICAgIHJldHVybiBldmFsKHN0cilcbiAgICB9KVxuICB9XG5cbiAgLyoqXG4gICAqIEFkZHMgYSBjYWxsYmFjayBmb3IgdGhlICoqaHRteDpsb2FkKiogZXZlbnQuIFRoaXMgY2FuIGJlIHVzZWQgdG8gcHJvY2VzcyBuZXcgY29udGVudCwgZm9yIGV4YW1wbGUgaW5pdGlhbGl6aW5nIHRoZSBjb250ZW50IHdpdGggYSBqYXZhc2NyaXB0IGxpYnJhcnlcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jb25Mb2FkXG4gICAqXG4gICAqIEBwYXJhbSB7KGVsdDogTm9kZSkgPT4gdm9pZH0gY2FsbGJhY2sgdGhlIGNhbGxiYWNrIHRvIGNhbGwgb24gbmV3bHkgbG9hZGVkIGNvbnRlbnRcbiAgICogQHJldHVybnMge0V2ZW50TGlzdGVuZXJ9XG4gICAqL1xuICBmdW5jdGlvbiBvbkxvYWRIZWxwZXIoY2FsbGJhY2spIHtcbiAgICBjb25zdCB2YWx1ZSA9IGh0bXgub24oJ2h0bXg6bG9hZCcsIC8qKiBAcGFyYW0ge0N1c3RvbUV2ZW50fSBldnQgKi8gZnVuY3Rpb24oZXZ0KSB7XG4gICAgICBjYWxsYmFjayhldnQuZGV0YWlsLmVsdClcbiAgICB9KVxuICAgIHJldHVybiB2YWx1ZVxuICB9XG5cbiAgLyoqXG4gICAqIExvZyBhbGwgaHRteCBldmVudHMsIHVzZWZ1bCBmb3IgZGVidWdnaW5nLlxuICAgKlxuICAgKiBAc2VlIGh0dHBzOi8vaHRteC5vcmcvYXBpLyNsb2dBbGxcbiAgICovXG4gIGZ1bmN0aW9uIGxvZ0FsbCgpIHtcbiAgICBodG14LmxvZ2dlciA9IGZ1bmN0aW9uKGVsdCwgZXZlbnQsIGRhdGEpIHtcbiAgICAgIGlmIChjb25zb2xlKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKGV2ZW50LCBlbHQsIGRhdGEpXG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgZnVuY3Rpb24gbG9nTm9uZSgpIHtcbiAgICBodG14LmxvZ2dlciA9IG51bGxcbiAgfVxuXG4gIC8qKlxuICAgKiBGaW5kcyBhbiBlbGVtZW50IG1hdGNoaW5nIHRoZSBzZWxlY3RvclxuICAgKlxuICAgKiBAc2VlIGh0dHBzOi8vaHRteC5vcmcvYXBpLyNmaW5kXG4gICAqXG4gICAqIEBwYXJhbSB7UGFyZW50Tm9kZXxzdHJpbmd9IGVsdE9yU2VsZWN0b3IgIHRoZSByb290IGVsZW1lbnQgdG8gZmluZCB0aGUgbWF0Y2hpbmcgZWxlbWVudCBpbiwgaW5jbHVzaXZlIHwgdGhlIHNlbGVjdG9yIHRvIG1hdGNoXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBbc2VsZWN0b3JdIHRoZSBzZWxlY3RvciB0byBtYXRjaFxuICAgKiBAcmV0dXJucyB7RWxlbWVudHxudWxsfVxuICAgKi9cbiAgZnVuY3Rpb24gZmluZChlbHRPclNlbGVjdG9yLCBzZWxlY3Rvcikge1xuICAgIGlmICh0eXBlb2YgZWx0T3JTZWxlY3RvciAhPT0gJ3N0cmluZycpIHtcbiAgICAgIHJldHVybiBlbHRPclNlbGVjdG9yLnF1ZXJ5U2VsZWN0b3Ioc2VsZWN0b3IpXG4gICAgfSBlbHNlIHtcbiAgICAgIHJldHVybiBmaW5kKGdldERvY3VtZW50KCksIGVsdE9yU2VsZWN0b3IpXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEZpbmRzIGFsbCBlbGVtZW50cyBtYXRjaGluZyB0aGUgc2VsZWN0b3JcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jZmluZEFsbFxuICAgKlxuICAgKiBAcGFyYW0ge1BhcmVudE5vZGV8c3RyaW5nfSBlbHRPclNlbGVjdG9yIHRoZSByb290IGVsZW1lbnQgdG8gZmluZCB0aGUgbWF0Y2hpbmcgZWxlbWVudHMgaW4sIGluY2x1c2l2ZSB8IHRoZSBzZWxlY3RvciB0byBtYXRjaFxuICAgKiBAcGFyYW0ge3N0cmluZ30gW3NlbGVjdG9yXSB0aGUgc2VsZWN0b3IgdG8gbWF0Y2hcbiAgICogQHJldHVybnMge05vZGVMaXN0T2Y8RWxlbWVudD59XG4gICAqL1xuICBmdW5jdGlvbiBmaW5kQWxsKGVsdE9yU2VsZWN0b3IsIHNlbGVjdG9yKSB7XG4gICAgaWYgKHR5cGVvZiBlbHRPclNlbGVjdG9yICE9PSAnc3RyaW5nJykge1xuICAgICAgcmV0dXJuIGVsdE9yU2VsZWN0b3IucXVlcnlTZWxlY3RvckFsbChzZWxlY3RvcilcbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuIGZpbmRBbGwoZ2V0RG9jdW1lbnQoKSwgZWx0T3JTZWxlY3RvcilcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHJldHVybnMgV2luZG93XG4gICAqL1xuICBmdW5jdGlvbiBnZXRXaW5kb3coKSB7XG4gICAgcmV0dXJuIHdpbmRvd1xuICB9XG5cbiAgLyoqXG4gICAqIFJlbW92ZXMgYW4gZWxlbWVudCBmcm9tIHRoZSBET01cbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jcmVtb3ZlXG4gICAqXG4gICAqIEBwYXJhbSB7Tm9kZX0gZWx0XG4gICAqIEBwYXJhbSB7bnVtYmVyfSBbZGVsYXldXG4gICAqL1xuICBmdW5jdGlvbiByZW1vdmVFbGVtZW50KGVsdCwgZGVsYXkpIHtcbiAgICBlbHQgPSByZXNvbHZlVGFyZ2V0KGVsdClcbiAgICBpZiAoZGVsYXkpIHtcbiAgICAgIGdldFdpbmRvdygpLnNldFRpbWVvdXQoZnVuY3Rpb24oKSB7XG4gICAgICAgIHJlbW92ZUVsZW1lbnQoZWx0KVxuICAgICAgICBlbHQgPSBudWxsXG4gICAgICB9LCBkZWxheSlcbiAgICB9IGVsc2Uge1xuICAgICAgcGFyZW50RWx0KGVsdCkucmVtb3ZlQ2hpbGQoZWx0KVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge2FueX0gZWx0XG4gICAqIEByZXR1cm4ge0VsZW1lbnR8bnVsbH1cbiAgICovXG4gIGZ1bmN0aW9uIGFzRWxlbWVudChlbHQpIHtcbiAgICByZXR1cm4gZWx0IGluc3RhbmNlb2YgRWxlbWVudCA/IGVsdCA6IG51bGxcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge2FueX0gZWx0XG4gICAqIEByZXR1cm4ge0hUTUxFbGVtZW50fG51bGx9XG4gICAqL1xuICBmdW5jdGlvbiBhc0h0bWxFbGVtZW50KGVsdCkge1xuICAgIHJldHVybiBlbHQgaW5zdGFuY2VvZiBIVE1MRWxlbWVudCA/IGVsdCA6IG51bGxcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge2FueX0gdmFsdWVcbiAgICogQHJldHVybiB7c3RyaW5nfG51bGx9XG4gICAqL1xuICBmdW5jdGlvbiBhc1N0cmluZyh2YWx1ZSkge1xuICAgIHJldHVybiB0eXBlb2YgdmFsdWUgPT09ICdzdHJpbmcnID8gdmFsdWUgOiBudWxsXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFdmVudFRhcmdldH0gZWx0XG4gICAqIEByZXR1cm4ge1BhcmVudE5vZGV8bnVsbH1cbiAgICovXG4gIGZ1bmN0aW9uIGFzUGFyZW50Tm9kZShlbHQpIHtcbiAgICByZXR1cm4gZWx0IGluc3RhbmNlb2YgRWxlbWVudCB8fCBlbHQgaW5zdGFuY2VvZiBEb2N1bWVudCB8fCBlbHQgaW5zdGFuY2VvZiBEb2N1bWVudEZyYWdtZW50ID8gZWx0IDogbnVsbFxuICB9XG5cbiAgLyoqXG4gICAqIFRoaXMgbWV0aG9kIGFkZHMgYSBjbGFzcyB0byB0aGUgZ2l2ZW4gZWxlbWVudC5cbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jYWRkQ2xhc3NcbiAgICpcbiAgICogQHBhcmFtIHtFbGVtZW50fHN0cmluZ30gZWx0IHRoZSBlbGVtZW50IHRvIGFkZCB0aGUgY2xhc3MgdG9cbiAgICogQHBhcmFtIHtzdHJpbmd9IGNsYXp6IHRoZSBjbGFzcyB0byBhZGRcbiAgICogQHBhcmFtIHtudW1iZXJ9IFtkZWxheV0gdGhlIGRlbGF5IChpbiBtaWxsaXNlY29uZHMpIGJlZm9yZSBjbGFzcyBpcyBhZGRlZFxuICAgKi9cbiAgZnVuY3Rpb24gYWRkQ2xhc3NUb0VsZW1lbnQoZWx0LCBjbGF6eiwgZGVsYXkpIHtcbiAgICBlbHQgPSBhc0VsZW1lbnQocmVzb2x2ZVRhcmdldChlbHQpKVxuICAgIGlmICghZWx0KSB7XG4gICAgICByZXR1cm5cbiAgICB9XG4gICAgaWYgKGRlbGF5KSB7XG4gICAgICBnZXRXaW5kb3coKS5zZXRUaW1lb3V0KGZ1bmN0aW9uKCkge1xuICAgICAgICBhZGRDbGFzc1RvRWxlbWVudChlbHQsIGNsYXp6KVxuICAgICAgICBlbHQgPSBudWxsXG4gICAgICB9LCBkZWxheSlcbiAgICB9IGVsc2Uge1xuICAgICAgZWx0LmNsYXNzTGlzdCAmJiBlbHQuY2xhc3NMaXN0LmFkZChjbGF6eilcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogUmVtb3ZlcyBhIGNsYXNzIGZyb20gdGhlIGdpdmVuIGVsZW1lbnRcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jcmVtb3ZlQ2xhc3NcbiAgICpcbiAgICogQHBhcmFtIHtOb2RlfHN0cmluZ30gbm9kZSBlbGVtZW50IHRvIHJlbW92ZSB0aGUgY2xhc3MgZnJvbVxuICAgKiBAcGFyYW0ge3N0cmluZ30gY2xhenogdGhlIGNsYXNzIHRvIHJlbW92ZVxuICAgKiBAcGFyYW0ge251bWJlcn0gW2RlbGF5XSB0aGUgZGVsYXkgKGluIG1pbGxpc2Vjb25kcyBiZWZvcmUgY2xhc3MgaXMgcmVtb3ZlZClcbiAgICovXG4gIGZ1bmN0aW9uIHJlbW92ZUNsYXNzRnJvbUVsZW1lbnQobm9kZSwgY2xhenosIGRlbGF5KSB7XG4gICAgbGV0IGVsdCA9IGFzRWxlbWVudChyZXNvbHZlVGFyZ2V0KG5vZGUpKVxuICAgIGlmICghZWx0KSB7XG4gICAgICByZXR1cm5cbiAgICB9XG4gICAgaWYgKGRlbGF5KSB7XG4gICAgICBnZXRXaW5kb3coKS5zZXRUaW1lb3V0KGZ1bmN0aW9uKCkge1xuICAgICAgICByZW1vdmVDbGFzc0Zyb21FbGVtZW50KGVsdCwgY2xhenopXG4gICAgICAgIGVsdCA9IG51bGxcbiAgICAgIH0sIGRlbGF5KVxuICAgIH0gZWxzZSB7XG4gICAgICBpZiAoZWx0LmNsYXNzTGlzdCkge1xuICAgICAgICBlbHQuY2xhc3NMaXN0LnJlbW92ZShjbGF6eilcbiAgICAgICAgLy8gaWYgdGhlcmUgYXJlIG5vIGNsYXNzZXMgbGVmdCwgcmVtb3ZlIHRoZSBjbGFzcyBhdHRyaWJ1dGVcbiAgICAgICAgaWYgKGVsdC5jbGFzc0xpc3QubGVuZ3RoID09PSAwKSB7XG4gICAgICAgICAgZWx0LnJlbW92ZUF0dHJpYnV0ZSgnY2xhc3MnKVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIFRvZ2dsZXMgdGhlIGdpdmVuIGNsYXNzIG9uIGFuIGVsZW1lbnRcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jdG9nZ2xlQ2xhc3NcbiAgICpcbiAgICogQHBhcmFtIHtFbGVtZW50fHN0cmluZ30gZWx0IHRoZSBlbGVtZW50IHRvIHRvZ2dsZSB0aGUgY2xhc3Mgb25cbiAgICogQHBhcmFtIHtzdHJpbmd9IGNsYXp6IHRoZSBjbGFzcyB0byB0b2dnbGVcbiAgICovXG4gIGZ1bmN0aW9uIHRvZ2dsZUNsYXNzT25FbGVtZW50KGVsdCwgY2xhenopIHtcbiAgICBlbHQgPSByZXNvbHZlVGFyZ2V0KGVsdClcbiAgICBlbHQuY2xhc3NMaXN0LnRvZ2dsZShjbGF6eilcbiAgfVxuXG4gIC8qKlxuICAgKiBUYWtlcyB0aGUgZ2l2ZW4gY2xhc3MgZnJvbSBpdHMgc2libGluZ3MsIHNvIHRoYXQgYW1vbmcgaXRzIHNpYmxpbmdzLCBvbmx5IHRoZSBnaXZlbiBlbGVtZW50IHdpbGwgaGF2ZSB0aGUgY2xhc3MuXG4gICAqXG4gICAqIEBzZWUgaHR0cHM6Ly9odG14Lm9yZy9hcGkvI3Rha2VDbGFzc1xuICAgKlxuICAgKiBAcGFyYW0ge05vZGV8c3RyaW5nfSBlbHQgdGhlIGVsZW1lbnQgdGhhdCB3aWxsIHRha2UgdGhlIGNsYXNzXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBjbGF6eiB0aGUgY2xhc3MgdG8gdGFrZVxuICAgKi9cbiAgZnVuY3Rpb24gdGFrZUNsYXNzRm9yRWxlbWVudChlbHQsIGNsYXp6KSB7XG4gICAgZWx0ID0gcmVzb2x2ZVRhcmdldChlbHQpXG4gICAgZm9yRWFjaChlbHQucGFyZW50RWxlbWVudC5jaGlsZHJlbiwgZnVuY3Rpb24oY2hpbGQpIHtcbiAgICAgIHJlbW92ZUNsYXNzRnJvbUVsZW1lbnQoY2hpbGQsIGNsYXp6KVxuICAgIH0pXG4gICAgYWRkQ2xhc3NUb0VsZW1lbnQoYXNFbGVtZW50KGVsdCksIGNsYXp6KVxuICB9XG5cbiAgLyoqXG4gICAqIEZpbmRzIHRoZSBjbG9zZXN0IG1hdGNoaW5nIGVsZW1lbnQgaW4gdGhlIGdpdmVuIGVsZW1lbnRzIHBhcmVudGFnZSwgaW5jbHVzaXZlIG9mIHRoZSBlbGVtZW50XG4gICAqXG4gICAqIEBzZWUgaHR0cHM6Ly9odG14Lm9yZy9hcGkvI2Nsb3Nlc3RcbiAgICpcbiAgICogQHBhcmFtIHtFbGVtZW50fHN0cmluZ30gZWx0IHRoZSBlbGVtZW50IHRvIGZpbmQgdGhlIHNlbGVjdG9yIGZyb21cbiAgICogQHBhcmFtIHtzdHJpbmd9IHNlbGVjdG9yIHRoZSBzZWxlY3RvciB0byBmaW5kXG4gICAqIEByZXR1cm5zIHtFbGVtZW50fG51bGx9XG4gICAqL1xuICBmdW5jdGlvbiBjbG9zZXN0KGVsdCwgc2VsZWN0b3IpIHtcbiAgICBlbHQgPSBhc0VsZW1lbnQocmVzb2x2ZVRhcmdldChlbHQpKVxuICAgIGlmIChlbHQgJiYgZWx0LmNsb3Nlc3QpIHtcbiAgICAgIHJldHVybiBlbHQuY2xvc2VzdChzZWxlY3RvcilcbiAgICB9IGVsc2Uge1xuICAgICAgLy8gVE9ETyByZW1vdmUgd2hlbiBJRSBnb2VzIGF3YXlcbiAgICAgIGRvIHtcbiAgICAgICAgaWYgKGVsdCA9PSBudWxsIHx8IG1hdGNoZXMoZWx0LCBzZWxlY3RvcikpIHtcbiAgICAgICAgICByZXR1cm4gZWx0XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICAgIHdoaWxlIChlbHQgPSBlbHQgJiYgYXNFbGVtZW50KHBhcmVudEVsdChlbHQpKSlcbiAgICAgIHJldHVybiBudWxsXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBzdHJcbiAgICogQHBhcmFtIHtzdHJpbmd9IHByZWZpeFxuICAgKiBAcmV0dXJucyB7Ym9vbGVhbn1cbiAgICovXG4gIGZ1bmN0aW9uIHN0YXJ0c1dpdGgoc3RyLCBwcmVmaXgpIHtcbiAgICByZXR1cm4gc3RyLnN1YnN0cmluZygwLCBwcmVmaXgubGVuZ3RoKSA9PT0gcHJlZml4XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IHN0clxuICAgKiBAcGFyYW0ge3N0cmluZ30gc3VmZml4XG4gICAqIEByZXR1cm5zIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gZW5kc1dpdGgoc3RyLCBzdWZmaXgpIHtcbiAgICByZXR1cm4gc3RyLnN1YnN0cmluZyhzdHIubGVuZ3RoIC0gc3VmZml4Lmxlbmd0aCkgPT09IHN1ZmZpeFxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBzZWxlY3RvclxuICAgKiBAcmV0dXJucyB7c3RyaW5nfVxuICAgKi9cbiAgZnVuY3Rpb24gbm9ybWFsaXplU2VsZWN0b3Ioc2VsZWN0b3IpIHtcbiAgICBjb25zdCB0cmltbWVkU2VsZWN0b3IgPSBzZWxlY3Rvci50cmltKClcbiAgICBpZiAoc3RhcnRzV2l0aCh0cmltbWVkU2VsZWN0b3IsICc8JykgJiYgZW5kc1dpdGgodHJpbW1lZFNlbGVjdG9yLCAnLz4nKSkge1xuICAgICAgcmV0dXJuIHRyaW1tZWRTZWxlY3Rvci5zdWJzdHJpbmcoMSwgdHJpbW1lZFNlbGVjdG9yLmxlbmd0aCAtIDIpXG4gICAgfSBlbHNlIHtcbiAgICAgIHJldHVybiB0cmltbWVkU2VsZWN0b3JcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfEVsZW1lbnR8RG9jdW1lbnR8c3RyaW5nfSBlbHRcbiAgICogQHBhcmFtIHtzdHJpbmd9IHNlbGVjdG9yXG4gICAqIEBwYXJhbSB7Ym9vbGVhbj19IGdsb2JhbFxuICAgKiBAcmV0dXJucyB7KE5vZGV8V2luZG93KVtdfVxuICAgKi9cbiAgZnVuY3Rpb24gcXVlcnlTZWxlY3RvckFsbEV4dChlbHQsIHNlbGVjdG9yLCBnbG9iYWwpIHtcbiAgICBlbHQgPSByZXNvbHZlVGFyZ2V0KGVsdClcbiAgICBpZiAoc2VsZWN0b3IuaW5kZXhPZignY2xvc2VzdCAnKSA9PT0gMCkge1xuICAgICAgcmV0dXJuIFtjbG9zZXN0KGFzRWxlbWVudChlbHQpLCBub3JtYWxpemVTZWxlY3RvcihzZWxlY3Rvci5zdWJzdHIoOCkpKV1cbiAgICB9IGVsc2UgaWYgKHNlbGVjdG9yLmluZGV4T2YoJ2ZpbmQgJykgPT09IDApIHtcbiAgICAgIHJldHVybiBbZmluZChhc1BhcmVudE5vZGUoZWx0KSwgbm9ybWFsaXplU2VsZWN0b3Ioc2VsZWN0b3Iuc3Vic3RyKDUpKSldXG4gICAgfSBlbHNlIGlmIChzZWxlY3RvciA9PT0gJ25leHQnKSB7XG4gICAgICByZXR1cm4gW2FzRWxlbWVudChlbHQpLm5leHRFbGVtZW50U2libGluZ11cbiAgICB9IGVsc2UgaWYgKHNlbGVjdG9yLmluZGV4T2YoJ25leHQgJykgPT09IDApIHtcbiAgICAgIHJldHVybiBbc2NhbkZvcndhcmRRdWVyeShlbHQsIG5vcm1hbGl6ZVNlbGVjdG9yKHNlbGVjdG9yLnN1YnN0cig1KSksICEhZ2xvYmFsKV1cbiAgICB9IGVsc2UgaWYgKHNlbGVjdG9yID09PSAncHJldmlvdXMnKSB7XG4gICAgICByZXR1cm4gW2FzRWxlbWVudChlbHQpLnByZXZpb3VzRWxlbWVudFNpYmxpbmddXG4gICAgfSBlbHNlIGlmIChzZWxlY3Rvci5pbmRleE9mKCdwcmV2aW91cyAnKSA9PT0gMCkge1xuICAgICAgcmV0dXJuIFtzY2FuQmFja3dhcmRzUXVlcnkoZWx0LCBub3JtYWxpemVTZWxlY3RvcihzZWxlY3Rvci5zdWJzdHIoOSkpLCAhIWdsb2JhbCldXG4gICAgfSBlbHNlIGlmIChzZWxlY3RvciA9PT0gJ2RvY3VtZW50Jykge1xuICAgICAgcmV0dXJuIFtkb2N1bWVudF1cbiAgICB9IGVsc2UgaWYgKHNlbGVjdG9yID09PSAnd2luZG93Jykge1xuICAgICAgcmV0dXJuIFt3aW5kb3ddXG4gICAgfSBlbHNlIGlmIChzZWxlY3RvciA9PT0gJ2JvZHknKSB7XG4gICAgICByZXR1cm4gW2RvY3VtZW50LmJvZHldXG4gICAgfSBlbHNlIGlmIChzZWxlY3RvciA9PT0gJ3Jvb3QnKSB7XG4gICAgICByZXR1cm4gW2dldFJvb3ROb2RlKGVsdCwgISFnbG9iYWwpXVxuICAgIH0gZWxzZSBpZiAoc2VsZWN0b3IuaW5kZXhPZignZ2xvYmFsICcpID09PSAwKSB7XG4gICAgICByZXR1cm4gcXVlcnlTZWxlY3RvckFsbEV4dChlbHQsIHNlbGVjdG9yLnNsaWNlKDcpLCB0cnVlKVxuICAgIH0gZWxzZSB7XG4gICAgICByZXR1cm4gdG9BcnJheShhc1BhcmVudE5vZGUoZ2V0Um9vdE5vZGUoZWx0LCAhIWdsb2JhbCkpLnF1ZXJ5U2VsZWN0b3JBbGwobm9ybWFsaXplU2VsZWN0b3Ioc2VsZWN0b3IpKSlcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfSBzdGFydFxuICAgKiBAcGFyYW0ge3N0cmluZ30gbWF0Y2hcbiAgICogQHBhcmFtIHtib29sZWFufSBnbG9iYWxcbiAgICogQHJldHVybnMge0VsZW1lbnR9XG4gICAqL1xuICB2YXIgc2NhbkZvcndhcmRRdWVyeSA9IGZ1bmN0aW9uKHN0YXJ0LCBtYXRjaCwgZ2xvYmFsKSB7XG4gICAgY29uc3QgcmVzdWx0cyA9IGFzUGFyZW50Tm9kZShnZXRSb290Tm9kZShzdGFydCwgZ2xvYmFsKSkucXVlcnlTZWxlY3RvckFsbChtYXRjaClcbiAgICBmb3IgKGxldCBpID0gMDsgaSA8IHJlc3VsdHMubGVuZ3RoOyBpKyspIHtcbiAgICAgIGNvbnN0IGVsdCA9IHJlc3VsdHNbaV1cbiAgICAgIGlmIChlbHQuY29tcGFyZURvY3VtZW50UG9zaXRpb24oc3RhcnQpID09PSBOb2RlLkRPQ1VNRU5UX1BPU0lUSU9OX1BSRUNFRElORykge1xuICAgICAgICByZXR1cm4gZWx0XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gc3RhcnRcbiAgICogQHBhcmFtIHtzdHJpbmd9IG1hdGNoXG4gICAqIEBwYXJhbSB7Ym9vbGVhbn0gZ2xvYmFsXG4gICAqIEByZXR1cm5zIHtFbGVtZW50fVxuICAgKi9cbiAgdmFyIHNjYW5CYWNrd2FyZHNRdWVyeSA9IGZ1bmN0aW9uKHN0YXJ0LCBtYXRjaCwgZ2xvYmFsKSB7XG4gICAgY29uc3QgcmVzdWx0cyA9IGFzUGFyZW50Tm9kZShnZXRSb290Tm9kZShzdGFydCwgZ2xvYmFsKSkucXVlcnlTZWxlY3RvckFsbChtYXRjaClcbiAgICBmb3IgKGxldCBpID0gcmVzdWx0cy5sZW5ndGggLSAxOyBpID49IDA7IGktLSkge1xuICAgICAgY29uc3QgZWx0ID0gcmVzdWx0c1tpXVxuICAgICAgaWYgKGVsdC5jb21wYXJlRG9jdW1lbnRQb3NpdGlvbihzdGFydCkgPT09IE5vZGUuRE9DVU1FTlRfUE9TSVRJT05fRk9MTE9XSU5HKSB7XG4gICAgICAgIHJldHVybiBlbHRcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfHN0cmluZ30gZWx0T3JTZWxlY3RvclxuICAgKiBAcGFyYW0ge3N0cmluZz19IHNlbGVjdG9yXG4gICAqIEByZXR1cm5zIHtOb2RlfFdpbmRvd31cbiAgICovXG4gIGZ1bmN0aW9uIHF1ZXJ5U2VsZWN0b3JFeHQoZWx0T3JTZWxlY3Rvciwgc2VsZWN0b3IpIHtcbiAgICBpZiAodHlwZW9mIGVsdE9yU2VsZWN0b3IgIT09ICdzdHJpbmcnKSB7XG4gICAgICByZXR1cm4gcXVlcnlTZWxlY3RvckFsbEV4dChlbHRPclNlbGVjdG9yLCBzZWxlY3RvcilbMF1cbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuIHF1ZXJ5U2VsZWN0b3JBbGxFeHQoZ2V0RG9jdW1lbnQoKS5ib2R5LCBlbHRPclNlbGVjdG9yKVswXVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAdGVtcGxhdGUge0V2ZW50VGFyZ2V0fSBUXG4gICAqIEBwYXJhbSB7VHxzdHJpbmd9IGVsdE9yU2VsZWN0b3JcbiAgICogQHBhcmFtIHtUfSBbY29udGV4dF1cbiAgICogQHJldHVybnMge0VsZW1lbnR8VHxudWxsfVxuICAgKi9cbiAgZnVuY3Rpb24gcmVzb2x2ZVRhcmdldChlbHRPclNlbGVjdG9yLCBjb250ZXh0KSB7XG4gICAgaWYgKHR5cGVvZiBlbHRPclNlbGVjdG9yID09PSAnc3RyaW5nJykge1xuICAgICAgcmV0dXJuIGZpbmQoYXNQYXJlbnROb2RlKGNvbnRleHQpIHx8IGRvY3VtZW50LCBlbHRPclNlbGVjdG9yKVxuICAgIH0gZWxzZSB7XG4gICAgICByZXR1cm4gZWx0T3JTZWxlY3RvclxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAdHlwZWRlZiB7a2V5b2YgSFRNTEVsZW1lbnRFdmVudE1hcHxzdHJpbmd9IEFueUV2ZW50TmFtZVxuICAgKi9cblxuICAvKipcbiAgICogQHR5cGVkZWYge09iamVjdH0gRXZlbnRBcmdzXG4gICAqIEBwcm9wZXJ0eSB7RXZlbnRUYXJnZXR9IHRhcmdldFxuICAgKiBAcHJvcGVydHkge0FueUV2ZW50TmFtZX0gZXZlbnRcbiAgICogQHByb3BlcnR5IHtFdmVudExpc3RlbmVyfSBsaXN0ZW5lclxuICAgKi9cblxuICAvKipcbiAgICogQHBhcmFtIHtFdmVudFRhcmdldHxBbnlFdmVudE5hbWV9IGFyZzFcbiAgICogQHBhcmFtIHtBbnlFdmVudE5hbWV8RXZlbnRMaXN0ZW5lcn0gYXJnMlxuICAgKiBAcGFyYW0ge0V2ZW50TGlzdGVuZXJ9IFthcmczXVxuICAgKiBAcmV0dXJucyB7RXZlbnRBcmdzfVxuICAgKi9cbiAgZnVuY3Rpb24gcHJvY2Vzc0V2ZW50QXJncyhhcmcxLCBhcmcyLCBhcmczKSB7XG4gICAgaWYgKGlzRnVuY3Rpb24oYXJnMikpIHtcbiAgICAgIHJldHVybiB7XG4gICAgICAgIHRhcmdldDogZ2V0RG9jdW1lbnQoKS5ib2R5LFxuICAgICAgICBldmVudDogYXNTdHJpbmcoYXJnMSksXG4gICAgICAgIGxpc3RlbmVyOiBhcmcyXG4gICAgICB9XG4gICAgfSBlbHNlIHtcbiAgICAgIHJldHVybiB7XG4gICAgICAgIHRhcmdldDogcmVzb2x2ZVRhcmdldChhcmcxKSxcbiAgICAgICAgZXZlbnQ6IGFzU3RyaW5nKGFyZzIpLFxuICAgICAgICBsaXN0ZW5lcjogYXJnM1xuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBBZGRzIGFuIGV2ZW50IGxpc3RlbmVyIHRvIGFuIGVsZW1lbnRcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jb25cbiAgICpcbiAgICogQHBhcmFtIHtFdmVudFRhcmdldHxzdHJpbmd9IGFyZzEgdGhlIGVsZW1lbnQgdG8gYWRkIHRoZSBsaXN0ZW5lciB0byB8IHRoZSBldmVudCBuYW1lIHRvIGFkZCB0aGUgbGlzdGVuZXIgZm9yXG4gICAqIEBwYXJhbSB7c3RyaW5nfEV2ZW50TGlzdGVuZXJ9IGFyZzIgdGhlIGV2ZW50IG5hbWUgdG8gYWRkIHRoZSBsaXN0ZW5lciBmb3IgfCB0aGUgbGlzdGVuZXIgdG8gYWRkXG4gICAqIEBwYXJhbSB7RXZlbnRMaXN0ZW5lcn0gW2FyZzNdIHRoZSBsaXN0ZW5lciB0byBhZGRcbiAgICogQHJldHVybnMge0V2ZW50TGlzdGVuZXJ9XG4gICAqL1xuICBmdW5jdGlvbiBhZGRFdmVudExpc3RlbmVySW1wbChhcmcxLCBhcmcyLCBhcmczKSB7XG4gICAgcmVhZHkoZnVuY3Rpb24oKSB7XG4gICAgICBjb25zdCBldmVudEFyZ3MgPSBwcm9jZXNzRXZlbnRBcmdzKGFyZzEsIGFyZzIsIGFyZzMpXG4gICAgICBldmVudEFyZ3MudGFyZ2V0LmFkZEV2ZW50TGlzdGVuZXIoZXZlbnRBcmdzLmV2ZW50LCBldmVudEFyZ3MubGlzdGVuZXIpXG4gICAgfSlcbiAgICBjb25zdCBiID0gaXNGdW5jdGlvbihhcmcyKVxuICAgIHJldHVybiBiID8gYXJnMiA6IGFyZzNcbiAgfVxuXG4gIC8qKlxuICAgKiBSZW1vdmVzIGFuIGV2ZW50IGxpc3RlbmVyIGZyb20gYW4gZWxlbWVudFxuICAgKlxuICAgKiBAc2VlIGh0dHBzOi8vaHRteC5vcmcvYXBpLyNvZmZcbiAgICpcbiAgICogQHBhcmFtIHtFdmVudFRhcmdldHxzdHJpbmd9IGFyZzEgdGhlIGVsZW1lbnQgdG8gcmVtb3ZlIHRoZSBsaXN0ZW5lciBmcm9tIHwgdGhlIGV2ZW50IG5hbWUgdG8gcmVtb3ZlIHRoZSBsaXN0ZW5lciBmcm9tXG4gICAqIEBwYXJhbSB7c3RyaW5nfEV2ZW50TGlzdGVuZXJ9IGFyZzIgdGhlIGV2ZW50IG5hbWUgdG8gcmVtb3ZlIHRoZSBsaXN0ZW5lciBmcm9tIHwgdGhlIGxpc3RlbmVyIHRvIHJlbW92ZVxuICAgKiBAcGFyYW0ge0V2ZW50TGlzdGVuZXJ9IFthcmczXSB0aGUgbGlzdGVuZXIgdG8gcmVtb3ZlXG4gICAqIEByZXR1cm5zIHtFdmVudExpc3RlbmVyfVxuICAgKi9cbiAgZnVuY3Rpb24gcmVtb3ZlRXZlbnRMaXN0ZW5lckltcGwoYXJnMSwgYXJnMiwgYXJnMykge1xuICAgIHJlYWR5KGZ1bmN0aW9uKCkge1xuICAgICAgY29uc3QgZXZlbnRBcmdzID0gcHJvY2Vzc0V2ZW50QXJncyhhcmcxLCBhcmcyLCBhcmczKVxuICAgICAgZXZlbnRBcmdzLnRhcmdldC5yZW1vdmVFdmVudExpc3RlbmVyKGV2ZW50QXJncy5ldmVudCwgZXZlbnRBcmdzLmxpc3RlbmVyKVxuICAgIH0pXG4gICAgcmV0dXJuIGlzRnVuY3Rpb24oYXJnMikgPyBhcmcyIDogYXJnM1xuICB9XG5cbiAgLy89ID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT1cbiAgLy8gTm9kZSBwcm9jZXNzaW5nXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG5cbiAgY29uc3QgRFVNTVlfRUxUID0gZ2V0RG9jdW1lbnQoKS5jcmVhdGVFbGVtZW50KCdvdXRwdXQnKSAvLyBkdW1teSBlbGVtZW50IGZvciBiYWQgc2VsZWN0b3JzXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ30gYXR0ck5hbWVcbiAgICogQHJldHVybnMgeyhOb2RlfFdpbmRvdylbXX1cbiAgICovXG4gIGZ1bmN0aW9uIGZpbmRBdHRyaWJ1dGVUYXJnZXRzKGVsdCwgYXR0ck5hbWUpIHtcbiAgICBjb25zdCBhdHRyVGFyZ2V0ID0gZ2V0Q2xvc2VzdEF0dHJpYnV0ZVZhbHVlKGVsdCwgYXR0ck5hbWUpXG4gICAgaWYgKGF0dHJUYXJnZXQpIHtcbiAgICAgIGlmIChhdHRyVGFyZ2V0ID09PSAndGhpcycpIHtcbiAgICAgICAgcmV0dXJuIFtmaW5kVGhpc0VsZW1lbnQoZWx0LCBhdHRyTmFtZSldXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBjb25zdCByZXN1bHQgPSBxdWVyeVNlbGVjdG9yQWxsRXh0KGVsdCwgYXR0clRhcmdldClcbiAgICAgICAgaWYgKHJlc3VsdC5sZW5ndGggPT09IDApIHtcbiAgICAgICAgICBsb2dFcnJvcignVGhlIHNlbGVjdG9yIFwiJyArIGF0dHJUYXJnZXQgKyAnXCIgb24gJyArIGF0dHJOYW1lICsgJyByZXR1cm5lZCBubyBtYXRjaGVzIScpXG4gICAgICAgICAgcmV0dXJuIFtEVU1NWV9FTFRdXG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgcmV0dXJuIHJlc3VsdFxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEBwYXJhbSB7c3RyaW5nfSBhdHRyaWJ1dGVcbiAgICogQHJldHVybnMge0VsZW1lbnR8bnVsbH1cbiAgICovXG4gIGZ1bmN0aW9uIGZpbmRUaGlzRWxlbWVudChlbHQsIGF0dHJpYnV0ZSkge1xuICAgIHJldHVybiBhc0VsZW1lbnQoZ2V0Q2xvc2VzdE1hdGNoKGVsdCwgZnVuY3Rpb24oZWx0KSB7XG4gICAgICByZXR1cm4gZ2V0QXR0cmlidXRlVmFsdWUoYXNFbGVtZW50KGVsdCksIGF0dHJpYnV0ZSkgIT0gbnVsbFxuICAgIH0pKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm5zIHtOb2RlfFdpbmRvd3xudWxsfVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0VGFyZ2V0KGVsdCkge1xuICAgIGNvbnN0IHRhcmdldFN0ciA9IGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC10YXJnZXQnKVxuICAgIGlmICh0YXJnZXRTdHIpIHtcbiAgICAgIGlmICh0YXJnZXRTdHIgPT09ICd0aGlzJykge1xuICAgICAgICByZXR1cm4gZmluZFRoaXNFbGVtZW50KGVsdCwgJ2h4LXRhcmdldCcpXG4gICAgICB9IGVsc2Uge1xuICAgICAgICByZXR1cm4gcXVlcnlTZWxlY3RvckV4dChlbHQsIHRhcmdldFN0cilcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgY29uc3QgZGF0YSA9IGdldEludGVybmFsRGF0YShlbHQpXG4gICAgICBpZiAoZGF0YS5ib29zdGVkKSB7XG4gICAgICAgIHJldHVybiBnZXREb2N1bWVudCgpLmJvZHlcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHJldHVybiBlbHRcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IG5hbWVcbiAgICogQHJldHVybnMge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiBzaG91bGRTZXR0bGVBdHRyaWJ1dGUobmFtZSkge1xuICAgIGNvbnN0IGF0dHJpYnV0ZXNUb1NldHRsZSA9IGh0bXguY29uZmlnLmF0dHJpYnV0ZXNUb1NldHRsZVxuICAgIGZvciAobGV0IGkgPSAwOyBpIDwgYXR0cmlidXRlc1RvU2V0dGxlLmxlbmd0aDsgaSsrKSB7XG4gICAgICBpZiAobmFtZSA9PT0gYXR0cmlidXRlc1RvU2V0dGxlW2ldKSB7XG4gICAgICAgIHJldHVybiB0cnVlXG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiBmYWxzZVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gbWVyZ2VUb1xuICAgKiBAcGFyYW0ge0VsZW1lbnR9IG1lcmdlRnJvbVxuICAgKi9cbiAgZnVuY3Rpb24gY2xvbmVBdHRyaWJ1dGVzKG1lcmdlVG8sIG1lcmdlRnJvbSkge1xuICAgIGZvckVhY2gobWVyZ2VUby5hdHRyaWJ1dGVzLCBmdW5jdGlvbihhdHRyKSB7XG4gICAgICBpZiAoIW1lcmdlRnJvbS5oYXNBdHRyaWJ1dGUoYXR0ci5uYW1lKSAmJiBzaG91bGRTZXR0bGVBdHRyaWJ1dGUoYXR0ci5uYW1lKSkge1xuICAgICAgICBtZXJnZVRvLnJlbW92ZUF0dHJpYnV0ZShhdHRyLm5hbWUpXG4gICAgICB9XG4gICAgfSlcbiAgICBmb3JFYWNoKG1lcmdlRnJvbS5hdHRyaWJ1dGVzLCBmdW5jdGlvbihhdHRyKSB7XG4gICAgICBpZiAoc2hvdWxkU2V0dGxlQXR0cmlidXRlKGF0dHIubmFtZSkpIHtcbiAgICAgICAgbWVyZ2VUby5zZXRBdHRyaWJ1dGUoYXR0ci5uYW1lLCBhdHRyLnZhbHVlKVxuICAgICAgfVxuICAgIH0pXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtIdG14U3dhcFN0eWxlfSBzd2FwU3R5bGVcbiAgICogQHBhcmFtIHtFbGVtZW50fSB0YXJnZXRcbiAgICogQHJldHVybnMge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiBpc0lubGluZVN3YXAoc3dhcFN0eWxlLCB0YXJnZXQpIHtcbiAgICBjb25zdCBleHRlbnNpb25zID0gZ2V0RXh0ZW5zaW9ucyh0YXJnZXQpXG4gICAgZm9yIChsZXQgaSA9IDA7IGkgPCBleHRlbnNpb25zLmxlbmd0aDsgaSsrKSB7XG4gICAgICBjb25zdCBleHRlbnNpb24gPSBleHRlbnNpb25zW2ldXG4gICAgICB0cnkge1xuICAgICAgICBpZiAoZXh0ZW5zaW9uLmlzSW5saW5lU3dhcChzd2FwU3R5bGUpKSB7XG4gICAgICAgICAgcmV0dXJuIHRydWVcbiAgICAgICAgfVxuICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICBsb2dFcnJvcihlKVxuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gc3dhcFN0eWxlID09PSAnb3V0ZXJIVE1MJ1xuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBvb2JWYWx1ZVxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IG9vYkVsZW1lbnRcbiAgICogQHBhcmFtIHtIdG14U2V0dGxlSW5mb30gc2V0dGxlSW5mb1xuICAgKiBAcmV0dXJuc1xuICAgKi9cbiAgZnVuY3Rpb24gb29iU3dhcChvb2JWYWx1ZSwgb29iRWxlbWVudCwgc2V0dGxlSW5mbykge1xuICAgIGxldCBzZWxlY3RvciA9ICcjJyArIGdldFJhd0F0dHJpYnV0ZShvb2JFbGVtZW50LCAnaWQnKVxuICAgIC8qKiBAdHlwZSBIdG14U3dhcFN0eWxlICovXG4gICAgbGV0IHN3YXBTdHlsZSA9ICdvdXRlckhUTUwnXG4gICAgaWYgKG9vYlZhbHVlID09PSAndHJ1ZScpIHtcbiAgICAgIC8vIGRvIG5vdGhpbmdcbiAgICB9IGVsc2UgaWYgKG9vYlZhbHVlLmluZGV4T2YoJzonKSA+IDApIHtcbiAgICAgIHN3YXBTdHlsZSA9IG9vYlZhbHVlLnN1YnN0cigwLCBvb2JWYWx1ZS5pbmRleE9mKCc6JykpXG4gICAgICBzZWxlY3RvciA9IG9vYlZhbHVlLnN1YnN0cihvb2JWYWx1ZS5pbmRleE9mKCc6JykgKyAxLCBvb2JWYWx1ZS5sZW5ndGgpXG4gICAgfSBlbHNlIHtcbiAgICAgIHN3YXBTdHlsZSA9IG9vYlZhbHVlXG4gICAgfVxuXG4gICAgY29uc3QgdGFyZ2V0cyA9IGdldERvY3VtZW50KCkucXVlcnlTZWxlY3RvckFsbChzZWxlY3RvcilcbiAgICBpZiAodGFyZ2V0cykge1xuICAgICAgZm9yRWFjaChcbiAgICAgICAgdGFyZ2V0cyxcbiAgICAgICAgZnVuY3Rpb24odGFyZ2V0KSB7XG4gICAgICAgICAgbGV0IGZyYWdtZW50XG4gICAgICAgICAgY29uc3Qgb29iRWxlbWVudENsb25lID0gb29iRWxlbWVudC5jbG9uZU5vZGUodHJ1ZSlcbiAgICAgICAgICBmcmFnbWVudCA9IGdldERvY3VtZW50KCkuY3JlYXRlRG9jdW1lbnRGcmFnbWVudCgpXG4gICAgICAgICAgZnJhZ21lbnQuYXBwZW5kQ2hpbGQob29iRWxlbWVudENsb25lKVxuICAgICAgICAgIGlmICghaXNJbmxpbmVTd2FwKHN3YXBTdHlsZSwgdGFyZ2V0KSkge1xuICAgICAgICAgICAgZnJhZ21lbnQgPSBhc1BhcmVudE5vZGUob29iRWxlbWVudENsb25lKSAvLyBpZiB0aGlzIGlzIG5vdCBhbiBpbmxpbmUgc3dhcCwgd2UgdXNlIHRoZSBjb250ZW50IG9mIHRoZSBub2RlLCBub3QgdGhlIG5vZGUgaXRzZWxmXG4gICAgICAgICAgfVxuXG4gICAgICAgICAgY29uc3QgYmVmb3JlU3dhcERldGFpbHMgPSB7IHNob3VsZFN3YXA6IHRydWUsIHRhcmdldCwgZnJhZ21lbnQgfVxuICAgICAgICAgIGlmICghdHJpZ2dlckV2ZW50KHRhcmdldCwgJ2h0bXg6b29iQmVmb3JlU3dhcCcsIGJlZm9yZVN3YXBEZXRhaWxzKSkgcmV0dXJuXG5cbiAgICAgICAgICB0YXJnZXQgPSBiZWZvcmVTd2FwRGV0YWlscy50YXJnZXQgLy8gYWxsb3cgcmUtdGFyZ2V0aW5nXG4gICAgICAgICAgaWYgKGJlZm9yZVN3YXBEZXRhaWxzLnNob3VsZFN3YXApIHtcbiAgICAgICAgICAgIHN3YXBXaXRoU3R5bGUoc3dhcFN0eWxlLCB0YXJnZXQsIHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pXG4gICAgICAgICAgfVxuICAgICAgICAgIGZvckVhY2goc2V0dGxlSW5mby5lbHRzLCBmdW5jdGlvbihlbHQpIHtcbiAgICAgICAgICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14Om9vYkFmdGVyU3dhcCcsIGJlZm9yZVN3YXBEZXRhaWxzKVxuICAgICAgICAgIH0pXG4gICAgICAgIH1cbiAgICAgIClcbiAgICAgIG9vYkVsZW1lbnQucGFyZW50Tm9kZS5yZW1vdmVDaGlsZChvb2JFbGVtZW50KVxuICAgIH0gZWxzZSB7XG4gICAgICBvb2JFbGVtZW50LnBhcmVudE5vZGUucmVtb3ZlQ2hpbGQob29iRWxlbWVudClcbiAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGdldERvY3VtZW50KCkuYm9keSwgJ2h0bXg6b29iRXJyb3JOb1RhcmdldCcsIHsgY29udGVudDogb29iRWxlbWVudCB9KVxuICAgIH1cbiAgICByZXR1cm4gb29iVmFsdWVcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0RvY3VtZW50RnJhZ21lbnR9IGZyYWdtZW50XG4gICAqL1xuICBmdW5jdGlvbiBoYW5kbGVQcmVzZXJ2ZWRFbGVtZW50cyhmcmFnbWVudCkge1xuICAgIGZvckVhY2goZmluZEFsbChmcmFnbWVudCwgJ1toeC1wcmVzZXJ2ZV0sIFtkYXRhLWh4LXByZXNlcnZlXScpLCBmdW5jdGlvbihwcmVzZXJ2ZWRFbHQpIHtcbiAgICAgIGNvbnN0IGlkID0gZ2V0QXR0cmlidXRlVmFsdWUocHJlc2VydmVkRWx0LCAnaWQnKVxuICAgICAgY29uc3Qgb2xkRWx0ID0gZ2V0RG9jdW1lbnQoKS5nZXRFbGVtZW50QnlJZChpZClcbiAgICAgIGlmIChvbGRFbHQgIT0gbnVsbCkge1xuICAgICAgICBwcmVzZXJ2ZWRFbHQucGFyZW50Tm9kZS5yZXBsYWNlQ2hpbGQob2xkRWx0LCBwcmVzZXJ2ZWRFbHQpXG4gICAgICB9XG4gICAgfSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IHBhcmVudE5vZGVcbiAgICogQHBhcmFtIHtQYXJlbnROb2RlfSBmcmFnbWVudFxuICAgKiBAcGFyYW0ge0h0bXhTZXR0bGVJbmZvfSBzZXR0bGVJbmZvXG4gICAqL1xuICBmdW5jdGlvbiBoYW5kbGVBdHRyaWJ1dGVzKHBhcmVudE5vZGUsIGZyYWdtZW50LCBzZXR0bGVJbmZvKSB7XG4gICAgZm9yRWFjaChmcmFnbWVudC5xdWVyeVNlbGVjdG9yQWxsKCdbaWRdJyksIGZ1bmN0aW9uKG5ld05vZGUpIHtcbiAgICAgIGNvbnN0IGlkID0gZ2V0UmF3QXR0cmlidXRlKG5ld05vZGUsICdpZCcpXG4gICAgICBpZiAoaWQgJiYgaWQubGVuZ3RoID4gMCkge1xuICAgICAgICBjb25zdCBub3JtYWxpemVkSWQgPSBpZC5yZXBsYWNlKFwiJ1wiLCBcIlxcXFwnXCIpXG4gICAgICAgIGNvbnN0IG5vcm1hbGl6ZWRUYWcgPSBuZXdOb2RlLnRhZ05hbWUucmVwbGFjZSgnOicsICdcXFxcOicpXG4gICAgICAgIGNvbnN0IHBhcmVudEVsdCA9IGFzUGFyZW50Tm9kZShwYXJlbnROb2RlKVxuICAgICAgICBjb25zdCBvbGROb2RlID0gcGFyZW50RWx0ICYmIHBhcmVudEVsdC5xdWVyeVNlbGVjdG9yKG5vcm1hbGl6ZWRUYWcgKyBcIltpZD0nXCIgKyBub3JtYWxpemVkSWQgKyBcIiddXCIpXG4gICAgICAgIGlmIChvbGROb2RlICYmIG9sZE5vZGUgIT09IHBhcmVudEVsdCkge1xuICAgICAgICAgIGNvbnN0IG5ld0F0dHJpYnV0ZXMgPSBuZXdOb2RlLmNsb25lTm9kZSgpXG4gICAgICAgICAgY2xvbmVBdHRyaWJ1dGVzKG5ld05vZGUsIG9sZE5vZGUpXG4gICAgICAgICAgc2V0dGxlSW5mby50YXNrcy5wdXNoKGZ1bmN0aW9uKCkge1xuICAgICAgICAgICAgY2xvbmVBdHRyaWJ1dGVzKG5ld05vZGUsIG5ld0F0dHJpYnV0ZXMpXG4gICAgICAgICAgfSlcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0pXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfSBjaGlsZFxuICAgKiBAcmV0dXJucyB7SHRteFNldHRsZVRhc2t9XG4gICAqL1xuICBmdW5jdGlvbiBtYWtlQWpheExvYWRUYXNrKGNoaWxkKSB7XG4gICAgcmV0dXJuIGZ1bmN0aW9uKCkge1xuICAgICAgcmVtb3ZlQ2xhc3NGcm9tRWxlbWVudChjaGlsZCwgaHRteC5jb25maWcuYWRkZWRDbGFzcylcbiAgICAgIHByb2Nlc3NOb2RlKGFzRWxlbWVudChjaGlsZCkpXG4gICAgICBwcm9jZXNzRm9jdXMoYXNQYXJlbnROb2RlKGNoaWxkKSlcbiAgICAgIHRyaWdnZXJFdmVudChjaGlsZCwgJ2h0bXg6bG9hZCcpXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7UGFyZW50Tm9kZX0gY2hpbGRcbiAgICovXG4gIGZ1bmN0aW9uIHByb2Nlc3NGb2N1cyhjaGlsZCkge1xuICAgIGNvbnN0IGF1dG9mb2N1cyA9ICdbYXV0b2ZvY3VzXSdcbiAgICBjb25zdCBhdXRvRm9jdXNlZEVsdCA9IGFzSHRtbEVsZW1lbnQobWF0Y2hlcyhjaGlsZCwgYXV0b2ZvY3VzKSA/IGNoaWxkIDogY2hpbGQucXVlcnlTZWxlY3RvcihhdXRvZm9jdXMpKVxuICAgIGlmIChhdXRvRm9jdXNlZEVsdCAhPSBudWxsKSB7XG4gICAgICBhdXRvRm9jdXNlZEVsdC5mb2N1cygpXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gcGFyZW50Tm9kZVxuICAgKiBAcGFyYW0ge05vZGV9IGluc2VydEJlZm9yZVxuICAgKiBAcGFyYW0ge1BhcmVudE5vZGV9IGZyYWdtZW50XG4gICAqIEBwYXJhbSB7SHRteFNldHRsZUluZm99IHNldHRsZUluZm9cbiAgICovXG4gIGZ1bmN0aW9uIGluc2VydE5vZGVzQmVmb3JlKHBhcmVudE5vZGUsIGluc2VydEJlZm9yZSwgZnJhZ21lbnQsIHNldHRsZUluZm8pIHtcbiAgICBoYW5kbGVBdHRyaWJ1dGVzKHBhcmVudE5vZGUsIGZyYWdtZW50LCBzZXR0bGVJbmZvKVxuICAgIHdoaWxlIChmcmFnbWVudC5jaGlsZE5vZGVzLmxlbmd0aCA+IDApIHtcbiAgICAgIGNvbnN0IGNoaWxkID0gZnJhZ21lbnQuZmlyc3RDaGlsZFxuICAgICAgYWRkQ2xhc3NUb0VsZW1lbnQoYXNFbGVtZW50KGNoaWxkKSwgaHRteC5jb25maWcuYWRkZWRDbGFzcylcbiAgICAgIHBhcmVudE5vZGUuaW5zZXJ0QmVmb3JlKGNoaWxkLCBpbnNlcnRCZWZvcmUpXG4gICAgICBpZiAoY2hpbGQubm9kZVR5cGUgIT09IE5vZGUuVEVYVF9OT0RFICYmIGNoaWxkLm5vZGVUeXBlICE9PSBOb2RlLkNPTU1FTlRfTk9ERSkge1xuICAgICAgICBzZXR0bGVJbmZvLnRhc2tzLnB1c2gobWFrZUFqYXhMb2FkVGFzayhjaGlsZCkpXG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIGJhc2VkIG9uIGh0dHBzOi8vZ2lzdC5naXRodWIuY29tL2h5YW1hbW90by9mZDQzNTUwNWQyOWViZmEzZDk3MTZmZDJiZThkNDJmMCxcbiAgICogZGVyaXZlZCBmcm9tIEphdmEncyBzdHJpbmcgaGFzaGNvZGUgaW1wbGVtZW50YXRpb25cbiAgICogQHBhcmFtIHtzdHJpbmd9IHN0cmluZ1xuICAgKiBAcGFyYW0ge251bWJlcn0gaGFzaFxuICAgKiBAcmV0dXJucyB7bnVtYmVyfVxuICAgKi9cbiAgZnVuY3Rpb24gc3RyaW5nSGFzaChzdHJpbmcsIGhhc2gpIHtcbiAgICBsZXQgY2hhciA9IDBcbiAgICB3aGlsZSAoY2hhciA8IHN0cmluZy5sZW5ndGgpIHtcbiAgICAgIGhhc2ggPSAoaGFzaCA8PCA1KSAtIGhhc2ggKyBzdHJpbmcuY2hhckNvZGVBdChjaGFyKyspIHwgMCAvLyBiaXR3aXNlIG9yIGVuc3VyZXMgd2UgaGF2ZSBhIDMyLWJpdCBpbnRcbiAgICB9XG4gICAgcmV0dXJuIGhhc2hcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKiBAcmV0dXJucyB7bnVtYmVyfVxuICAgKi9cbiAgZnVuY3Rpb24gYXR0cmlidXRlSGFzaChlbHQpIHtcbiAgICBsZXQgaGFzaCA9IDBcbiAgICAvLyBJRSBmaXhcbiAgICBpZiAoZWx0LmF0dHJpYnV0ZXMpIHtcbiAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgZWx0LmF0dHJpYnV0ZXMubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgY29uc3QgYXR0cmlidXRlID0gZWx0LmF0dHJpYnV0ZXNbaV1cbiAgICAgICAgaWYgKGF0dHJpYnV0ZS52YWx1ZSkgeyAvLyBvbmx5IGluY2x1ZGUgYXR0cmlidXRlcyB3LyBhY3R1YWwgdmFsdWVzIChlbXB0eSBpcyBzYW1lIGFzIG5vbi1leGlzdGVudClcbiAgICAgICAgICBoYXNoID0gc3RyaW5nSGFzaChhdHRyaWJ1dGUubmFtZSwgaGFzaClcbiAgICAgICAgICBoYXNoID0gc3RyaW5nSGFzaChhdHRyaWJ1dGUudmFsdWUsIGhhc2gpXG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuIGhhc2hcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0V2ZW50VGFyZ2V0fSBlbHRcbiAgICovXG4gIGZ1bmN0aW9uIGRlSW5pdE9uSGFuZGxlcnMoZWx0KSB7XG4gICAgY29uc3QgaW50ZXJuYWxEYXRhID0gZ2V0SW50ZXJuYWxEYXRhKGVsdClcbiAgICBpZiAoaW50ZXJuYWxEYXRhLm9uSGFuZGxlcnMpIHtcbiAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgaW50ZXJuYWxEYXRhLm9uSGFuZGxlcnMubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgY29uc3QgaGFuZGxlckluZm8gPSBpbnRlcm5hbERhdGEub25IYW5kbGVyc1tpXVxuICAgICAgICByZW1vdmVFdmVudExpc3RlbmVySW1wbChlbHQsIGhhbmRsZXJJbmZvLmV2ZW50LCBoYW5kbGVySW5mby5saXN0ZW5lcilcbiAgICAgIH1cbiAgICAgIGRlbGV0ZSBpbnRlcm5hbERhdGEub25IYW5kbGVyc1xuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsZW1lbnRcbiAgICovXG4gIGZ1bmN0aW9uIGRlSW5pdE5vZGUoZWxlbWVudCkge1xuICAgIGNvbnN0IGludGVybmFsRGF0YSA9IGdldEludGVybmFsRGF0YShlbGVtZW50KVxuICAgIGlmIChpbnRlcm5hbERhdGEudGltZW91dCkge1xuICAgICAgY2xlYXJUaW1lb3V0KGludGVybmFsRGF0YS50aW1lb3V0KVxuICAgIH1cbiAgICBpZiAoaW50ZXJuYWxEYXRhLmxpc3RlbmVySW5mb3MpIHtcbiAgICAgIGZvckVhY2goaW50ZXJuYWxEYXRhLmxpc3RlbmVySW5mb3MsIGZ1bmN0aW9uKGluZm8pIHtcbiAgICAgICAgaWYgKGluZm8ub24pIHtcbiAgICAgICAgICByZW1vdmVFdmVudExpc3RlbmVySW1wbChpbmZvLm9uLCBpbmZvLnRyaWdnZXIsIGluZm8ubGlzdGVuZXIpXG4gICAgICAgIH1cbiAgICAgIH0pXG4gICAgfVxuICAgIGRlSW5pdE9uSGFuZGxlcnMoZWxlbWVudClcbiAgICBmb3JFYWNoKE9iamVjdC5rZXlzKGludGVybmFsRGF0YSksIGZ1bmN0aW9uKGtleSkgeyBkZWxldGUgaW50ZXJuYWxEYXRhW2tleV0gfSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsZW1lbnRcbiAgICovXG4gIGZ1bmN0aW9uIGNsZWFuVXBFbGVtZW50KGVsZW1lbnQpIHtcbiAgICB0cmlnZ2VyRXZlbnQoZWxlbWVudCwgJ2h0bXg6YmVmb3JlQ2xlYW51cEVsZW1lbnQnKVxuICAgIGRlSW5pdE5vZGUoZWxlbWVudClcbiAgICAvLyBAdHMtaWdub3JlIElFMTEgY29kZVxuICAgIC8vIG5vaW5zcGVjdGlvbiBKU1VucmVzb2x2ZWRSZWZlcmVuY2VcbiAgICBpZiAoZWxlbWVudC5jaGlsZHJlbikgeyAvLyBJRVxuICAgICAgLy8gQHRzLWlnbm9yZVxuICAgICAgZm9yRWFjaChlbGVtZW50LmNoaWxkcmVuLCBmdW5jdGlvbihjaGlsZCkgeyBjbGVhblVwRWxlbWVudChjaGlsZCkgfSlcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfSB0YXJnZXRcbiAgICogQHBhcmFtIHtQYXJlbnROb2RlfSBmcmFnbWVudFxuICAgKiBAcGFyYW0ge0h0bXhTZXR0bGVJbmZvfSBzZXR0bGVJbmZvXG4gICAqL1xuICBmdW5jdGlvbiBzd2FwT3V0ZXJIVE1MKHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pIHtcbiAgICBpZiAodGFyZ2V0IGluc3RhbmNlb2YgRWxlbWVudCAmJiB0YXJnZXQudGFnTmFtZSA9PT0gJ0JPRFknKSB7IC8vIHNwZWNpYWwgY2FzZSB0aGUgYm9keSB0byBpbm5lckhUTUwgYmVjYXVzZSBEb2N1bWVudEZyYWdtZW50cyBjYW4ndCBjb250YWluIGEgYm9keSBlbHQgdW5mb3J0dW5hdGVseVxuICAgICAgcmV0dXJuIHN3YXBJbm5lckhUTUwodGFyZ2V0LCBmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgICB9XG4gICAgLyoqIEB0eXBlIHtOb2RlfSAqL1xuICAgIGxldCBuZXdFbHRcbiAgICBjb25zdCBlbHRCZWZvcmVOZXdDb250ZW50ID0gdGFyZ2V0LnByZXZpb3VzU2libGluZ1xuICAgIGluc2VydE5vZGVzQmVmb3JlKHBhcmVudEVsdCh0YXJnZXQpLCB0YXJnZXQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKVxuICAgIGlmIChlbHRCZWZvcmVOZXdDb250ZW50ID09IG51bGwpIHtcbiAgICAgIG5ld0VsdCA9IHBhcmVudEVsdCh0YXJnZXQpLmZpcnN0Q2hpbGRcbiAgICB9IGVsc2Uge1xuICAgICAgbmV3RWx0ID0gZWx0QmVmb3JlTmV3Q29udGVudC5uZXh0U2libGluZ1xuICAgIH1cbiAgICBzZXR0bGVJbmZvLmVsdHMgPSBzZXR0bGVJbmZvLmVsdHMuZmlsdGVyKGZ1bmN0aW9uKGUpIHsgcmV0dXJuIGUgIT09IHRhcmdldCB9KVxuICAgIC8vIHNjYW4gdGhyb3VnaCBhbGwgbmV3bHkgYWRkZWQgY29udGVudCBhbmQgYWRkIGFsbCBlbGVtZW50cyB0byB0aGUgc2V0dGxlIGluZm8gc28gd2UgdHJpZ2dlclxuICAgIC8vIGV2ZW50cyBwcm9wZXJseSBvbiB0aGVtXG4gICAgd2hpbGUgKG5ld0VsdCAmJiBuZXdFbHQgIT09IHRhcmdldCkge1xuICAgICAgaWYgKG5ld0VsdCBpbnN0YW5jZW9mIEVsZW1lbnQpIHtcbiAgICAgICAgc2V0dGxlSW5mby5lbHRzLnB1c2gobmV3RWx0KVxuICAgICAgfVxuICAgICAgbmV3RWx0ID0gbmV3RWx0Lm5leHRTaWJsaW5nXG4gICAgfVxuICAgIGNsZWFuVXBFbGVtZW50KHRhcmdldClcbiAgICBpZiAodGFyZ2V0IGluc3RhbmNlb2YgRWxlbWVudCkge1xuICAgICAgdGFyZ2V0LnJlbW92ZSgpXG4gICAgfSBlbHNlIHtcbiAgICAgIHRhcmdldC5wYXJlbnROb2RlLnJlbW92ZUNoaWxkKHRhcmdldClcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfSB0YXJnZXRcbiAgICogQHBhcmFtIHtQYXJlbnROb2RlfSBmcmFnbWVudFxuICAgKiBAcGFyYW0ge0h0bXhTZXR0bGVJbmZvfSBzZXR0bGVJbmZvXG4gICAqL1xuICBmdW5jdGlvbiBzd2FwQWZ0ZXJCZWdpbih0YXJnZXQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKSB7XG4gICAgcmV0dXJuIGluc2VydE5vZGVzQmVmb3JlKHRhcmdldCwgdGFyZ2V0LmZpcnN0Q2hpbGQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gdGFyZ2V0XG4gICAqIEBwYXJhbSB7UGFyZW50Tm9kZX0gZnJhZ21lbnRcbiAgICogQHBhcmFtIHtIdG14U2V0dGxlSW5mb30gc2V0dGxlSW5mb1xuICAgKi9cbiAgZnVuY3Rpb24gc3dhcEJlZm9yZUJlZ2luKHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pIHtcbiAgICByZXR1cm4gaW5zZXJ0Tm9kZXNCZWZvcmUocGFyZW50RWx0KHRhcmdldCksIHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtOb2RlfSB0YXJnZXRcbiAgICogQHBhcmFtIHtQYXJlbnROb2RlfSBmcmFnbWVudFxuICAgKiBAcGFyYW0ge0h0bXhTZXR0bGVJbmZvfSBzZXR0bGVJbmZvXG4gICAqL1xuICBmdW5jdGlvbiBzd2FwQmVmb3JlRW5kKHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pIHtcbiAgICByZXR1cm4gaW5zZXJ0Tm9kZXNCZWZvcmUodGFyZ2V0LCBudWxsLCBmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IHRhcmdldFxuICAgKiBAcGFyYW0ge1BhcmVudE5vZGV9IGZyYWdtZW50XG4gICAqIEBwYXJhbSB7SHRteFNldHRsZUluZm99IHNldHRsZUluZm9cbiAgICovXG4gIGZ1bmN0aW9uIHN3YXBBZnRlckVuZCh0YXJnZXQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKSB7XG4gICAgcmV0dXJuIGluc2VydE5vZGVzQmVmb3JlKHBhcmVudEVsdCh0YXJnZXQpLCB0YXJnZXQubmV4dFNpYmxpbmcsIGZyYWdtZW50LCBzZXR0bGVJbmZvKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gdGFyZ2V0XG4gICAqL1xuICBmdW5jdGlvbiBzd2FwRGVsZXRlKHRhcmdldCkge1xuICAgIGNsZWFuVXBFbGVtZW50KHRhcmdldClcbiAgICByZXR1cm4gcGFyZW50RWx0KHRhcmdldCkucmVtb3ZlQ2hpbGQodGFyZ2V0KVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gdGFyZ2V0XG4gICAqIEBwYXJhbSB7UGFyZW50Tm9kZX0gZnJhZ21lbnRcbiAgICogQHBhcmFtIHtIdG14U2V0dGxlSW5mb30gc2V0dGxlSW5mb1xuICAgKi9cbiAgZnVuY3Rpb24gc3dhcElubmVySFRNTCh0YXJnZXQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKSB7XG4gICAgY29uc3QgZmlyc3RDaGlsZCA9IHRhcmdldC5maXJzdENoaWxkXG4gICAgaW5zZXJ0Tm9kZXNCZWZvcmUodGFyZ2V0LCBmaXJzdENoaWxkLCBmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgICBpZiAoZmlyc3RDaGlsZCkge1xuICAgICAgd2hpbGUgKGZpcnN0Q2hpbGQubmV4dFNpYmxpbmcpIHtcbiAgICAgICAgY2xlYW5VcEVsZW1lbnQoZmlyc3RDaGlsZC5uZXh0U2libGluZylcbiAgICAgICAgdGFyZ2V0LnJlbW92ZUNoaWxkKGZpcnN0Q2hpbGQubmV4dFNpYmxpbmcpXG4gICAgICB9XG4gICAgICBjbGVhblVwRWxlbWVudChmaXJzdENoaWxkKVxuICAgICAgdGFyZ2V0LnJlbW92ZUNoaWxkKGZpcnN0Q2hpbGQpXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7SHRteFN3YXBTdHlsZX0gc3dhcFN0eWxlXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEBwYXJhbSB7Tm9kZX0gdGFyZ2V0XG4gICAqIEBwYXJhbSB7UGFyZW50Tm9kZX0gZnJhZ21lbnRcbiAgICogQHBhcmFtIHtIdG14U2V0dGxlSW5mb30gc2V0dGxlSW5mb1xuICAgKi9cbiAgZnVuY3Rpb24gc3dhcFdpdGhTdHlsZShzd2FwU3R5bGUsIGVsdCwgdGFyZ2V0LCBmcmFnbWVudCwgc2V0dGxlSW5mbykge1xuICAgIHN3aXRjaCAoc3dhcFN0eWxlKSB7XG4gICAgICBjYXNlICdub25lJzpcbiAgICAgICAgcmV0dXJuXG4gICAgICBjYXNlICdvdXRlckhUTUwnOlxuICAgICAgICBzd2FwT3V0ZXJIVE1MKHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pXG4gICAgICAgIHJldHVyblxuICAgICAgY2FzZSAnYWZ0ZXJiZWdpbic6XG4gICAgICAgIHN3YXBBZnRlckJlZ2luKHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pXG4gICAgICAgIHJldHVyblxuICAgICAgY2FzZSAnYmVmb3JlYmVnaW4nOlxuICAgICAgICBzd2FwQmVmb3JlQmVnaW4odGFyZ2V0LCBmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgICAgICAgcmV0dXJuXG4gICAgICBjYXNlICdiZWZvcmVlbmQnOlxuICAgICAgICBzd2FwQmVmb3JlRW5kKHRhcmdldCwgZnJhZ21lbnQsIHNldHRsZUluZm8pXG4gICAgICAgIHJldHVyblxuICAgICAgY2FzZSAnYWZ0ZXJlbmQnOlxuICAgICAgICBzd2FwQWZ0ZXJFbmQodGFyZ2V0LCBmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgICAgICAgcmV0dXJuXG4gICAgICBjYXNlICdkZWxldGUnOlxuICAgICAgICBzd2FwRGVsZXRlKHRhcmdldClcbiAgICAgICAgcmV0dXJuXG4gICAgICBkZWZhdWx0OlxuICAgICAgICB2YXIgZXh0ZW5zaW9ucyA9IGdldEV4dGVuc2lvbnMoZWx0KVxuICAgICAgICBmb3IgKGxldCBpID0gMDsgaSA8IGV4dGVuc2lvbnMubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgICBjb25zdCBleHQgPSBleHRlbnNpb25zW2ldXG4gICAgICAgICAgdHJ5IHtcbiAgICAgICAgICAgIGNvbnN0IG5ld0VsZW1lbnRzID0gZXh0LmhhbmRsZVN3YXAoc3dhcFN0eWxlLCB0YXJnZXQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKVxuICAgICAgICAgICAgaWYgKG5ld0VsZW1lbnRzKSB7XG4gICAgICAgICAgICAgIGlmIChBcnJheS5pc0FycmF5KG5ld0VsZW1lbnRzKSkge1xuICAgICAgICAgICAgICAgIC8vIGlmIGhhbmRsZVN3YXAgcmV0dXJucyBhbiBhcnJheSAobGlrZSkgb2YgZWxlbWVudHMsIHdlIGhhbmRsZSB0aGVtXG4gICAgICAgICAgICAgICAgZm9yIChsZXQgaiA9IDA7IGogPCBuZXdFbGVtZW50cy5sZW5ndGg7IGorKykge1xuICAgICAgICAgICAgICAgICAgY29uc3QgY2hpbGQgPSBuZXdFbGVtZW50c1tqXVxuICAgICAgICAgICAgICAgICAgaWYgKGNoaWxkLm5vZGVUeXBlICE9PSBOb2RlLlRFWFRfTk9ERSAmJiBjaGlsZC5ub2RlVHlwZSAhPT0gTm9kZS5DT01NRU5UX05PREUpIHtcbiAgICAgICAgICAgICAgICAgICAgc2V0dGxlSW5mby50YXNrcy5wdXNoKG1ha2VBamF4TG9hZFRhc2soY2hpbGQpKVxuICAgICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICByZXR1cm5cbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgICAgICBsb2dFcnJvcihlKVxuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBpZiAoc3dhcFN0eWxlID09PSAnaW5uZXJIVE1MJykge1xuICAgICAgICAgIHN3YXBJbm5lckhUTUwodGFyZ2V0LCBmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBzd2FwV2l0aFN0eWxlKGh0bXguY29uZmlnLmRlZmF1bHRTd2FwU3R5bGUsIGVsdCwgdGFyZ2V0LCBmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0RvY3VtZW50RnJhZ21lbnR9IGZyYWdtZW50XG4gICAqIEBwYXJhbSB7SHRteFNldHRsZUluZm99IHNldHRsZUluZm9cbiAgICovXG4gIGZ1bmN0aW9uIGZpbmRBbmRTd2FwT29iRWxlbWVudHMoZnJhZ21lbnQsIHNldHRsZUluZm8pIHtcbiAgICB2YXIgb29iRWx0cyA9IGZpbmRBbGwoZnJhZ21lbnQsICdbaHgtc3dhcC1vb2JdLCBbZGF0YS1oeC1zd2FwLW9vYl0nKVxuICAgIGZvckVhY2gob29iRWx0cywgZnVuY3Rpb24ob29iRWxlbWVudCkge1xuICAgICAgaWYgKGh0bXguY29uZmlnLmFsbG93TmVzdGVkT29iU3dhcHMgfHwgb29iRWxlbWVudC5wYXJlbnRFbGVtZW50ID09PSBudWxsKSB7XG4gICAgICAgIGNvbnN0IG9vYlZhbHVlID0gZ2V0QXR0cmlidXRlVmFsdWUob29iRWxlbWVudCwgJ2h4LXN3YXAtb29iJylcbiAgICAgICAgaWYgKG9vYlZhbHVlICE9IG51bGwpIHtcbiAgICAgICAgICBvb2JTd2FwKG9vYlZhbHVlLCBvb2JFbGVtZW50LCBzZXR0bGVJbmZvKVxuICAgICAgICB9XG4gICAgICB9IGVsc2Uge1xuICAgICAgICBvb2JFbGVtZW50LnJlbW92ZUF0dHJpYnV0ZSgnaHgtc3dhcC1vb2InKVxuICAgICAgICBvb2JFbGVtZW50LnJlbW92ZUF0dHJpYnV0ZSgnZGF0YS1oeC1zd2FwLW9vYicpXG4gICAgICB9XG4gICAgfSlcbiAgICByZXR1cm4gb29iRWx0cy5sZW5ndGggPiAwXG4gIH1cblxuICAvKipcbiAgICogSW1wbGVtZW50cyBjb21wbGV0ZSBzd2FwcGluZyBwaXBlbGluZSwgaW5jbHVkaW5nOiBmb2N1cyBhbmQgc2VsZWN0aW9uIHByZXNlcnZhdGlvbixcbiAgICogdGl0bGUgdXBkYXRlcywgc2Nyb2xsLCBPT0Igc3dhcHBpbmcsIG5vcm1hbCBzd2FwcGluZyBhbmQgc2V0dGxpbmdcbiAgICogQHBhcmFtIHtzdHJpbmd8RWxlbWVudH0gdGFyZ2V0XG4gICAqIEBwYXJhbSB7c3RyaW5nfSBjb250ZW50XG4gICAqIEBwYXJhbSB7SHRteFN3YXBTcGVjaWZpY2F0aW9ufSBzd2FwU3BlY1xuICAgKiBAcGFyYW0ge1N3YXBPcHRpb25zfSBbc3dhcE9wdGlvbnNdXG4gICAqL1xuICBmdW5jdGlvbiBzd2FwKHRhcmdldCwgY29udGVudCwgc3dhcFNwZWMsIHN3YXBPcHRpb25zKSB7XG4gICAgaWYgKCFzd2FwT3B0aW9ucykge1xuICAgICAgc3dhcE9wdGlvbnMgPSB7fVxuICAgIH1cblxuICAgIHRhcmdldCA9IHJlc29sdmVUYXJnZXQodGFyZ2V0KVxuXG4gICAgLy8gcHJlc2VydmUgZm9jdXMgYW5kIHNlbGVjdGlvblxuICAgIGNvbnN0IGFjdGl2ZUVsdCA9IGRvY3VtZW50LmFjdGl2ZUVsZW1lbnRcbiAgICBsZXQgc2VsZWN0aW9uSW5mbyA9IHt9XG4gICAgdHJ5IHtcbiAgICAgIHNlbGVjdGlvbkluZm8gPSB7XG4gICAgICAgIGVsdDogYWN0aXZlRWx0LFxuICAgICAgICAvLyBAdHMtaWdub3JlXG4gICAgICAgIHN0YXJ0OiBhY3RpdmVFbHQgPyBhY3RpdmVFbHQuc2VsZWN0aW9uU3RhcnQgOiBudWxsLFxuICAgICAgICAvLyBAdHMtaWdub3JlXG4gICAgICAgIGVuZDogYWN0aXZlRWx0ID8gYWN0aXZlRWx0LnNlbGVjdGlvbkVuZCA6IG51bGxcbiAgICAgIH1cbiAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBzYWZhcmkgaXNzdWUgLSBzZWUgaHR0cHM6Ly9naXRodWIuY29tL21pY3Jvc29mdC9wbGF5d3JpZ2h0L2lzc3Vlcy81ODk0XG4gICAgfVxuICAgIGNvbnN0IHNldHRsZUluZm8gPSBtYWtlU2V0dGxlSW5mbyh0YXJnZXQpXG5cbiAgICAvLyBGb3IgdGV4dCBjb250ZW50IHN3YXBzLCBkb24ndCBwYXJzZSB0aGUgcmVzcG9uc2UgYXMgSFRNTCwganVzdCBpbnNlcnQgaXRcbiAgICBpZiAoc3dhcFNwZWMuc3dhcFN0eWxlID09PSAndGV4dENvbnRlbnQnKSB7XG4gICAgICB0YXJnZXQudGV4dENvbnRlbnQgPSBjb250ZW50XG4gICAgLy8gT3RoZXJ3aXNlLCBtYWtlIHRoZSBmcmFnbWVudCBhbmQgcHJvY2VzcyBpdFxuICAgIH0gZWxzZSB7XG4gICAgICBsZXQgZnJhZ21lbnQgPSBtYWtlRnJhZ21lbnQoY29udGVudClcblxuICAgICAgc2V0dGxlSW5mby50aXRsZSA9IGZyYWdtZW50LnRpdGxlXG5cbiAgICAgIC8vIHNlbGVjdC1vb2Igc3dhcHNcbiAgICAgIGlmIChzd2FwT3B0aW9ucy5zZWxlY3RPT0IpIHtcbiAgICAgICAgY29uc3Qgb29iU2VsZWN0VmFsdWVzID0gc3dhcE9wdGlvbnMuc2VsZWN0T09CLnNwbGl0KCcsJylcbiAgICAgICAgZm9yIChsZXQgaSA9IDA7IGkgPCBvb2JTZWxlY3RWYWx1ZXMubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgICBjb25zdCBvb2JTZWxlY3RWYWx1ZSA9IG9vYlNlbGVjdFZhbHVlc1tpXS5zcGxpdCgnOicsIDIpXG4gICAgICAgICAgbGV0IGlkID0gb29iU2VsZWN0VmFsdWVbMF0udHJpbSgpXG4gICAgICAgICAgaWYgKGlkLmluZGV4T2YoJyMnKSA9PT0gMCkge1xuICAgICAgICAgICAgaWQgPSBpZC5zdWJzdHJpbmcoMSlcbiAgICAgICAgICB9XG4gICAgICAgICAgY29uc3Qgb29iVmFsdWUgPSBvb2JTZWxlY3RWYWx1ZVsxXSB8fCAndHJ1ZSdcbiAgICAgICAgICBjb25zdCBvb2JFbGVtZW50ID0gZnJhZ21lbnQucXVlcnlTZWxlY3RvcignIycgKyBpZClcbiAgICAgICAgICBpZiAob29iRWxlbWVudCkge1xuICAgICAgICAgICAgb29iU3dhcChvb2JWYWx1ZSwgb29iRWxlbWVudCwgc2V0dGxlSW5mbylcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICAgIC8vIG9vYiBzd2Fwc1xuICAgICAgZmluZEFuZFN3YXBPb2JFbGVtZW50cyhmcmFnbWVudCwgc2V0dGxlSW5mbylcbiAgICAgIGZvckVhY2goZmluZEFsbChmcmFnbWVudCwgJ3RlbXBsYXRlJyksIC8qKiBAcGFyYW0ge0hUTUxUZW1wbGF0ZUVsZW1lbnR9IHRlbXBsYXRlICovZnVuY3Rpb24odGVtcGxhdGUpIHtcbiAgICAgICAgaWYgKGZpbmRBbmRTd2FwT29iRWxlbWVudHModGVtcGxhdGUuY29udGVudCwgc2V0dGxlSW5mbykpIHtcbiAgICAgICAgICAvLyBBdm9pZCBwb2xsdXRpbmcgdGhlIERPTSB3aXRoIGVtcHR5IHRlbXBsYXRlcyB0aGF0IHdlcmUgb25seSB1c2VkIHRvIGVuY2Fwc3VsYXRlIG9vYiBzd2FwXG4gICAgICAgICAgdGVtcGxhdGUucmVtb3ZlKClcbiAgICAgICAgfVxuICAgICAgfSlcblxuICAgICAgLy8gbm9ybWFsIHN3YXBcbiAgICAgIGlmIChzd2FwT3B0aW9ucy5zZWxlY3QpIHtcbiAgICAgICAgY29uc3QgbmV3RnJhZ21lbnQgPSBnZXREb2N1bWVudCgpLmNyZWF0ZURvY3VtZW50RnJhZ21lbnQoKVxuICAgICAgICBmb3JFYWNoKGZyYWdtZW50LnF1ZXJ5U2VsZWN0b3JBbGwoc3dhcE9wdGlvbnMuc2VsZWN0KSwgZnVuY3Rpb24obm9kZSkge1xuICAgICAgICAgIG5ld0ZyYWdtZW50LmFwcGVuZENoaWxkKG5vZGUpXG4gICAgICAgIH0pXG4gICAgICAgIGZyYWdtZW50ID0gbmV3RnJhZ21lbnRcbiAgICAgIH1cbiAgICAgIGhhbmRsZVByZXNlcnZlZEVsZW1lbnRzKGZyYWdtZW50KVxuICAgICAgc3dhcFdpdGhTdHlsZShzd2FwU3BlYy5zd2FwU3R5bGUsIHN3YXBPcHRpb25zLmNvbnRleHRFbGVtZW50LCB0YXJnZXQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKVxuICAgIH1cblxuICAgIC8vIGFwcGx5IHNhdmVkIGZvY3VzIGFuZCBzZWxlY3Rpb24gaW5mb3JtYXRpb24gdG8gc3dhcHBlZCBjb250ZW50XG4gICAgaWYgKHNlbGVjdGlvbkluZm8uZWx0ICYmXG4gICAgICAhYm9keUNvbnRhaW5zKHNlbGVjdGlvbkluZm8uZWx0KSAmJlxuICAgICAgZ2V0UmF3QXR0cmlidXRlKHNlbGVjdGlvbkluZm8uZWx0LCAnaWQnKSkge1xuICAgICAgY29uc3QgbmV3QWN0aXZlRWx0ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoZ2V0UmF3QXR0cmlidXRlKHNlbGVjdGlvbkluZm8uZWx0LCAnaWQnKSlcbiAgICAgIGNvbnN0IGZvY3VzT3B0aW9ucyA9IHsgcHJldmVudFNjcm9sbDogc3dhcFNwZWMuZm9jdXNTY3JvbGwgIT09IHVuZGVmaW5lZCA/ICFzd2FwU3BlYy5mb2N1c1Njcm9sbCA6ICFodG14LmNvbmZpZy5kZWZhdWx0Rm9jdXNTY3JvbGwgfVxuICAgICAgaWYgKG5ld0FjdGl2ZUVsdCkge1xuICAgICAgICAvLyBAdHMtaWdub3JlXG4gICAgICAgIGlmIChzZWxlY3Rpb25JbmZvLnN0YXJ0ICYmIG5ld0FjdGl2ZUVsdC5zZXRTZWxlY3Rpb25SYW5nZSkge1xuICAgICAgICAgIHRyeSB7XG4gICAgICAgICAgICAvLyBAdHMtaWdub3JlXG4gICAgICAgICAgICBuZXdBY3RpdmVFbHQuc2V0U2VsZWN0aW9uUmFuZ2Uoc2VsZWN0aW9uSW5mby5zdGFydCwgc2VsZWN0aW9uSW5mby5lbmQpXG4gICAgICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICAgICAgLy8gdGhlIHNldFNlbGVjdGlvblJhbmdlIG1ldGhvZCBpcyBwcmVzZW50IG9uIGZpZWxkcyB0aGF0IGRvbid0IHN1cHBvcnQgaXQsIHNvIGp1c3QgbGV0IHRoaXMgZmFpbFxuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBuZXdBY3RpdmVFbHQuZm9jdXMoZm9jdXNPcHRpb25zKVxuICAgICAgfVxuICAgIH1cblxuICAgIHRhcmdldC5jbGFzc0xpc3QucmVtb3ZlKGh0bXguY29uZmlnLnN3YXBwaW5nQ2xhc3MpXG4gICAgZm9yRWFjaChzZXR0bGVJbmZvLmVsdHMsIGZ1bmN0aW9uKGVsdCkge1xuICAgICAgaWYgKGVsdC5jbGFzc0xpc3QpIHtcbiAgICAgICAgZWx0LmNsYXNzTGlzdC5hZGQoaHRteC5jb25maWcuc2V0dGxpbmdDbGFzcylcbiAgICAgIH1cbiAgICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14OmFmdGVyU3dhcCcsIHN3YXBPcHRpb25zLmV2ZW50SW5mbylcbiAgICB9KVxuICAgIGlmIChzd2FwT3B0aW9ucy5hZnRlclN3YXBDYWxsYmFjaykge1xuICAgICAgc3dhcE9wdGlvbnMuYWZ0ZXJTd2FwQ2FsbGJhY2soKVxuICAgIH1cblxuICAgIC8vIG1lcmdlIGluIG5ldyB0aXRsZSBhZnRlciBzd2FwIGJ1dCBiZWZvcmUgc2V0dGxlXG4gICAgaWYgKCFzd2FwU3BlYy5pZ25vcmVUaXRsZSkge1xuICAgICAgaGFuZGxlVGl0bGUoc2V0dGxlSW5mby50aXRsZSlcbiAgICB9XG5cbiAgICAvLyBzZXR0bGVcbiAgICBjb25zdCBkb1NldHRsZSA9IGZ1bmN0aW9uKCkge1xuICAgICAgZm9yRWFjaChzZXR0bGVJbmZvLnRhc2tzLCBmdW5jdGlvbih0YXNrKSB7XG4gICAgICAgIHRhc2suY2FsbCgpXG4gICAgICB9KVxuICAgICAgZm9yRWFjaChzZXR0bGVJbmZvLmVsdHMsIGZ1bmN0aW9uKGVsdCkge1xuICAgICAgICBpZiAoZWx0LmNsYXNzTGlzdCkge1xuICAgICAgICAgIGVsdC5jbGFzc0xpc3QucmVtb3ZlKGh0bXguY29uZmlnLnNldHRsaW5nQ2xhc3MpXG4gICAgICAgIH1cbiAgICAgICAgdHJpZ2dlckV2ZW50KGVsdCwgJ2h0bXg6YWZ0ZXJTZXR0bGUnLCBzd2FwT3B0aW9ucy5ldmVudEluZm8pXG4gICAgICB9KVxuXG4gICAgICBpZiAoc3dhcE9wdGlvbnMuYW5jaG9yKSB7XG4gICAgICAgIGNvbnN0IGFuY2hvclRhcmdldCA9IGFzRWxlbWVudChyZXNvbHZlVGFyZ2V0KCcjJyArIHN3YXBPcHRpb25zLmFuY2hvcikpXG4gICAgICAgIGlmIChhbmNob3JUYXJnZXQpIHtcbiAgICAgICAgICBhbmNob3JUYXJnZXQuc2Nyb2xsSW50b1ZpZXcoeyBibG9jazogJ3N0YXJ0JywgYmVoYXZpb3I6ICdhdXRvJyB9KVxuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIHVwZGF0ZVNjcm9sbFN0YXRlKHNldHRsZUluZm8uZWx0cywgc3dhcFNwZWMpXG4gICAgICBpZiAoc3dhcE9wdGlvbnMuYWZ0ZXJTZXR0bGVDYWxsYmFjaykge1xuICAgICAgICBzd2FwT3B0aW9ucy5hZnRlclNldHRsZUNhbGxiYWNrKClcbiAgICAgIH1cbiAgICB9XG5cbiAgICBpZiAoc3dhcFNwZWMuc2V0dGxlRGVsYXkgPiAwKSB7XG4gICAgICBnZXRXaW5kb3coKS5zZXRUaW1lb3V0KGRvU2V0dGxlLCBzd2FwU3BlYy5zZXR0bGVEZWxheSlcbiAgICB9IGVsc2Uge1xuICAgICAgZG9TZXR0bGUoKVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge1hNTEh0dHBSZXF1ZXN0fSB4aHJcbiAgICogQHBhcmFtIHtzdHJpbmd9IGhlYWRlclxuICAgKiBAcGFyYW0ge0V2ZW50VGFyZ2V0fSBlbHRcbiAgICovXG4gIGZ1bmN0aW9uIGhhbmRsZVRyaWdnZXJIZWFkZXIoeGhyLCBoZWFkZXIsIGVsdCkge1xuICAgIGNvbnN0IHRyaWdnZXJCb2R5ID0geGhyLmdldFJlc3BvbnNlSGVhZGVyKGhlYWRlcilcbiAgICBpZiAodHJpZ2dlckJvZHkuaW5kZXhPZigneycpID09PSAwKSB7XG4gICAgICBjb25zdCB0cmlnZ2VycyA9IHBhcnNlSlNPTih0cmlnZ2VyQm9keSlcbiAgICAgIGZvciAoY29uc3QgZXZlbnROYW1lIGluIHRyaWdnZXJzKSB7XG4gICAgICAgIGlmICh0cmlnZ2Vycy5oYXNPd25Qcm9wZXJ0eShldmVudE5hbWUpKSB7XG4gICAgICAgICAgbGV0IGRldGFpbCA9IHRyaWdnZXJzW2V2ZW50TmFtZV1cbiAgICAgICAgICBpZiAoaXNSYXdPYmplY3QoZGV0YWlsKSkge1xuICAgICAgICAgICAgLy8gQHRzLWlnbm9yZVxuICAgICAgICAgICAgZWx0ID0gZGV0YWlsLnRhcmdldCAhPT0gdW5kZWZpbmVkID8gZGV0YWlsLnRhcmdldCA6IGVsdFxuICAgICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgICBkZXRhaWwgPSB7IHZhbHVlOiBkZXRhaWwgfVxuICAgICAgICAgIH1cbiAgICAgICAgICB0cmlnZ2VyRXZlbnQoZWx0LCBldmVudE5hbWUsIGRldGFpbClcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICBjb25zdCBldmVudE5hbWVzID0gdHJpZ2dlckJvZHkuc3BsaXQoJywnKVxuICAgICAgZm9yIChsZXQgaSA9IDA7IGkgPCBldmVudE5hbWVzLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgIHRyaWdnZXJFdmVudChlbHQsIGV2ZW50TmFtZXNbaV0udHJpbSgpLCBbXSlcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICBjb25zdCBXSElURVNQQUNFID0gL1xccy9cbiAgY29uc3QgV0hJVEVTUEFDRV9PUl9DT01NQSA9IC9bXFxzLF0vXG4gIGNvbnN0IFNZTUJPTF9TVEFSVCA9IC9bXyRhLXpBLVpdL1xuICBjb25zdCBTWU1CT0xfQ09OVCA9IC9bXyRhLXpBLVowLTldL1xuICBjb25zdCBTVFJJTkdJU0hfU1RBUlQgPSBbJ1wiJywgXCInXCIsICcvJ11cbiAgY29uc3QgTk9UX1dISVRFU1BBQ0UgPSAvW15cXHNdL1xuICBjb25zdCBDT01CSU5FRF9TRUxFQ1RPUl9TVEFSVCA9IC9beyhdL1xuICBjb25zdCBDT01CSU5FRF9TRUxFQ1RPUl9FTkQgPSAvW30pXS9cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IHN0clxuICAgKiBAcmV0dXJucyB7c3RyaW5nW119XG4gICAqL1xuICBmdW5jdGlvbiB0b2tlbml6ZVN0cmluZyhzdHIpIHtcbiAgICAvKiogQHR5cGUgc3RyaW5nW10gKi9cbiAgICBjb25zdCB0b2tlbnMgPSBbXVxuICAgIGxldCBwb3NpdGlvbiA9IDBcbiAgICB3aGlsZSAocG9zaXRpb24gPCBzdHIubGVuZ3RoKSB7XG4gICAgICBpZiAoU1lNQk9MX1NUQVJULmV4ZWMoc3RyLmNoYXJBdChwb3NpdGlvbikpKSB7XG4gICAgICAgIHZhciBzdGFydFBvc2l0aW9uID0gcG9zaXRpb25cbiAgICAgICAgd2hpbGUgKFNZTUJPTF9DT05ULmV4ZWMoc3RyLmNoYXJBdChwb3NpdGlvbiArIDEpKSkge1xuICAgICAgICAgIHBvc2l0aW9uKytcbiAgICAgICAgfVxuICAgICAgICB0b2tlbnMucHVzaChzdHIuc3Vic3RyKHN0YXJ0UG9zaXRpb24sIHBvc2l0aW9uIC0gc3RhcnRQb3NpdGlvbiArIDEpKVxuICAgICAgfSBlbHNlIGlmIChTVFJJTkdJU0hfU1RBUlQuaW5kZXhPZihzdHIuY2hhckF0KHBvc2l0aW9uKSkgIT09IC0xKSB7XG4gICAgICAgIGNvbnN0IHN0YXJ0Q2hhciA9IHN0ci5jaGFyQXQocG9zaXRpb24pXG4gICAgICAgIHZhciBzdGFydFBvc2l0aW9uID0gcG9zaXRpb25cbiAgICAgICAgcG9zaXRpb24rK1xuICAgICAgICB3aGlsZSAocG9zaXRpb24gPCBzdHIubGVuZ3RoICYmIHN0ci5jaGFyQXQocG9zaXRpb24pICE9PSBzdGFydENoYXIpIHtcbiAgICAgICAgICBpZiAoc3RyLmNoYXJBdChwb3NpdGlvbikgPT09ICdcXFxcJykge1xuICAgICAgICAgICAgcG9zaXRpb24rK1xuICAgICAgICAgIH1cbiAgICAgICAgICBwb3NpdGlvbisrXG4gICAgICAgIH1cbiAgICAgICAgdG9rZW5zLnB1c2goc3RyLnN1YnN0cihzdGFydFBvc2l0aW9uLCBwb3NpdGlvbiAtIHN0YXJ0UG9zaXRpb24gKyAxKSlcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIGNvbnN0IHN5bWJvbCA9IHN0ci5jaGFyQXQocG9zaXRpb24pXG4gICAgICAgIHRva2Vucy5wdXNoKHN5bWJvbClcbiAgICAgIH1cbiAgICAgIHBvc2l0aW9uKytcbiAgICB9XG4gICAgcmV0dXJuIHRva2Vuc1xuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSB0b2tlblxuICAgKiBAcGFyYW0ge3N0cmluZ3xudWxsfSBsYXN0XG4gICAqIEBwYXJhbSB7c3RyaW5nfSBwYXJhbU5hbWVcbiAgICogQHJldHVybnMge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiBpc1Bvc3NpYmxlUmVsYXRpdmVSZWZlcmVuY2UodG9rZW4sIGxhc3QsIHBhcmFtTmFtZSkge1xuICAgIHJldHVybiBTWU1CT0xfU1RBUlQuZXhlYyh0b2tlbi5jaGFyQXQoMCkpICYmXG4gICAgICB0b2tlbiAhPT0gJ3RydWUnICYmXG4gICAgICB0b2tlbiAhPT0gJ2ZhbHNlJyAmJlxuICAgICAgdG9rZW4gIT09ICd0aGlzJyAmJlxuICAgICAgdG9rZW4gIT09IHBhcmFtTmFtZSAmJlxuICAgICAgbGFzdCAhPT0gJy4nXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFdmVudFRhcmdldHxzdHJpbmd9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ1tdfSB0b2tlbnNcbiAgICogQHBhcmFtIHtzdHJpbmd9IHBhcmFtTmFtZVxuICAgKiBAcmV0dXJucyB7Q29uZGl0aW9uYWxGdW5jdGlvbnxudWxsfVxuICAgKi9cbiAgZnVuY3Rpb24gbWF5YmVHZW5lcmF0ZUNvbmRpdGlvbmFsKGVsdCwgdG9rZW5zLCBwYXJhbU5hbWUpIHtcbiAgICBpZiAodG9rZW5zWzBdID09PSAnWycpIHtcbiAgICAgIHRva2Vucy5zaGlmdCgpXG4gICAgICBsZXQgYnJhY2tldENvdW50ID0gMVxuICAgICAgbGV0IGNvbmRpdGlvbmFsU291cmNlID0gJyByZXR1cm4gKGZ1bmN0aW9uKCcgKyBwYXJhbU5hbWUgKyAnKXsgcmV0dXJuICgnXG4gICAgICBsZXQgbGFzdCA9IG51bGxcbiAgICAgIHdoaWxlICh0b2tlbnMubGVuZ3RoID4gMCkge1xuICAgICAgICBjb25zdCB0b2tlbiA9IHRva2Vuc1swXVxuICAgICAgICAvLyBAdHMtaWdub3JlIEZvciBzb21lIHJlYXNvbiB0c2MgZG9lc24ndCB1bmRlcnN0YW5kIHRoZSBzaGlmdCBjYWxsLCBhbmQgdGhpbmtzIHdlJ3JlIGNvbXBhcmluZyB0aGUgc2FtZSB2YWx1ZSBoZXJlLCBpLmUuICdbJyB2cyAnXSdcbiAgICAgICAgaWYgKHRva2VuID09PSAnXScpIHtcbiAgICAgICAgICBicmFja2V0Q291bnQtLVxuICAgICAgICAgIGlmIChicmFja2V0Q291bnQgPT09IDApIHtcbiAgICAgICAgICAgIGlmIChsYXN0ID09PSBudWxsKSB7XG4gICAgICAgICAgICAgIGNvbmRpdGlvbmFsU291cmNlID0gY29uZGl0aW9uYWxTb3VyY2UgKyAndHJ1ZSdcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIHRva2Vucy5zaGlmdCgpXG4gICAgICAgICAgICBjb25kaXRpb25hbFNvdXJjZSArPSAnKX0pJ1xuICAgICAgICAgICAgdHJ5IHtcbiAgICAgICAgICAgICAgY29uc3QgY29uZGl0aW9uRnVuY3Rpb24gPSBtYXliZUV2YWwoZWx0LCBmdW5jdGlvbigpIHtcbiAgICAgICAgICAgICAgICByZXR1cm4gRnVuY3Rpb24oY29uZGl0aW9uYWxTb3VyY2UpKClcbiAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgICAgZnVuY3Rpb24oKSB7IHJldHVybiB0cnVlIH0pXG4gICAgICAgICAgICAgIGNvbmRpdGlvbkZ1bmN0aW9uLnNvdXJjZSA9IGNvbmRpdGlvbmFsU291cmNlXG4gICAgICAgICAgICAgIHJldHVybiBjb25kaXRpb25GdW5jdGlvblxuICAgICAgICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICAgICAgICB0cmlnZ2VyRXJyb3JFdmVudChnZXREb2N1bWVudCgpLmJvZHksICdodG14OnN5bnRheDplcnJvcicsIHsgZXJyb3I6IGUsIHNvdXJjZTogY29uZGl0aW9uYWxTb3VyY2UgfSlcbiAgICAgICAgICAgICAgcmV0dXJuIG51bGxcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgIH0gZWxzZSBpZiAodG9rZW4gPT09ICdbJykge1xuICAgICAgICAgIGJyYWNrZXRDb3VudCsrXG4gICAgICAgIH1cbiAgICAgICAgaWYgKGlzUG9zc2libGVSZWxhdGl2ZVJlZmVyZW5jZSh0b2tlbiwgbGFzdCwgcGFyYW1OYW1lKSkge1xuICAgICAgICAgIGNvbmRpdGlvbmFsU291cmNlICs9ICcoKCcgKyBwYXJhbU5hbWUgKyAnLicgKyB0b2tlbiArICcpID8gKCcgKyBwYXJhbU5hbWUgKyAnLicgKyB0b2tlbiArICcpIDogKHdpbmRvdy4nICsgdG9rZW4gKyAnKSknXG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgY29uZGl0aW9uYWxTb3VyY2UgPSBjb25kaXRpb25hbFNvdXJjZSArIHRva2VuXG4gICAgICAgIH1cbiAgICAgICAgbGFzdCA9IHRva2Vucy5zaGlmdCgpXG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nW119IHRva2Vuc1xuICAgKiBAcGFyYW0ge1JlZ0V4cH0gbWF0Y2hcbiAgICogQHJldHVybnMge3N0cmluZ31cbiAgICovXG4gIGZ1bmN0aW9uIGNvbnN1bWVVbnRpbCh0b2tlbnMsIG1hdGNoKSB7XG4gICAgbGV0IHJlc3VsdCA9ICcnXG4gICAgd2hpbGUgKHRva2Vucy5sZW5ndGggPiAwICYmICFtYXRjaC50ZXN0KHRva2Vuc1swXSkpIHtcbiAgICAgIHJlc3VsdCArPSB0b2tlbnMuc2hpZnQoKVxuICAgIH1cbiAgICByZXR1cm4gcmVzdWx0XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmdbXX0gdG9rZW5zXG4gICAqIEByZXR1cm5zIHtzdHJpbmd9XG4gICAqL1xuICBmdW5jdGlvbiBjb25zdW1lQ1NTU2VsZWN0b3IodG9rZW5zKSB7XG4gICAgbGV0IHJlc3VsdFxuICAgIGlmICh0b2tlbnMubGVuZ3RoID4gMCAmJiBDT01CSU5FRF9TRUxFQ1RPUl9TVEFSVC50ZXN0KHRva2Vuc1swXSkpIHtcbiAgICAgIHRva2Vucy5zaGlmdCgpXG4gICAgICByZXN1bHQgPSBjb25zdW1lVW50aWwodG9rZW5zLCBDT01CSU5FRF9TRUxFQ1RPUl9FTkQpLnRyaW0oKVxuICAgICAgdG9rZW5zLnNoaWZ0KClcbiAgICB9IGVsc2Uge1xuICAgICAgcmVzdWx0ID0gY29uc3VtZVVudGlsKHRva2VucywgV0hJVEVTUEFDRV9PUl9DT01NQSlcbiAgICB9XG4gICAgcmV0dXJuIHJlc3VsdFxuICB9XG5cbiAgY29uc3QgSU5QVVRfU0VMRUNUT1IgPSAnaW5wdXQsIHRleHRhcmVhLCBzZWxlY3QnXG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEBwYXJhbSB7c3RyaW5nfSBleHBsaWNpdFRyaWdnZXJcbiAgICogQHBhcmFtIHtPYmplY3R9IGNhY2hlIGZvciB0cmlnZ2VyIHNwZWNzXG4gICAqIEByZXR1cm5zIHtIdG14VHJpZ2dlclNwZWNpZmljYXRpb25bXX1cbiAgICovXG4gIGZ1bmN0aW9uIHBhcnNlQW5kQ2FjaGVUcmlnZ2VyKGVsdCwgZXhwbGljaXRUcmlnZ2VyLCBjYWNoZSkge1xuICAgIC8qKiBAdHlwZSBIdG14VHJpZ2dlclNwZWNpZmljYXRpb25bXSAqL1xuICAgIGNvbnN0IHRyaWdnZXJTcGVjcyA9IFtdXG4gICAgY29uc3QgdG9rZW5zID0gdG9rZW5pemVTdHJpbmcoZXhwbGljaXRUcmlnZ2VyKVxuICAgIGRvIHtcbiAgICAgIGNvbnN1bWVVbnRpbCh0b2tlbnMsIE5PVF9XSElURVNQQUNFKVxuICAgICAgY29uc3QgaW5pdGlhbExlbmd0aCA9IHRva2Vucy5sZW5ndGhcbiAgICAgIGNvbnN0IHRyaWdnZXIgPSBjb25zdW1lVW50aWwodG9rZW5zLCAvWyxcXFtcXHNdLylcbiAgICAgIGlmICh0cmlnZ2VyICE9PSAnJykge1xuICAgICAgICBpZiAodHJpZ2dlciA9PT0gJ2V2ZXJ5Jykge1xuICAgICAgICAgIC8qKiBAdHlwZSBIdG14VHJpZ2dlclNwZWNpZmljYXRpb24gKi9cbiAgICAgICAgICBjb25zdCBldmVyeSA9IHsgdHJpZ2dlcjogJ2V2ZXJ5JyB9XG4gICAgICAgICAgY29uc3VtZVVudGlsKHRva2VucywgTk9UX1dISVRFU1BBQ0UpXG4gICAgICAgICAgZXZlcnkucG9sbEludGVydmFsID0gcGFyc2VJbnRlcnZhbChjb25zdW1lVW50aWwodG9rZW5zLCAvWyxcXFtcXHNdLykpXG4gICAgICAgICAgY29uc3VtZVVudGlsKHRva2VucywgTk9UX1dISVRFU1BBQ0UpXG4gICAgICAgICAgdmFyIGV2ZW50RmlsdGVyID0gbWF5YmVHZW5lcmF0ZUNvbmRpdGlvbmFsKGVsdCwgdG9rZW5zLCAnZXZlbnQnKVxuICAgICAgICAgIGlmIChldmVudEZpbHRlcikge1xuICAgICAgICAgICAgZXZlcnkuZXZlbnRGaWx0ZXIgPSBldmVudEZpbHRlclxuICAgICAgICAgIH1cbiAgICAgICAgICB0cmlnZ2VyU3BlY3MucHVzaChldmVyeSlcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAvKiogQHR5cGUgSHRteFRyaWdnZXJTcGVjaWZpY2F0aW9uICovXG4gICAgICAgICAgY29uc3QgdHJpZ2dlclNwZWMgPSB7IHRyaWdnZXIgfVxuICAgICAgICAgIHZhciBldmVudEZpbHRlciA9IG1heWJlR2VuZXJhdGVDb25kaXRpb25hbChlbHQsIHRva2VucywgJ2V2ZW50JylcbiAgICAgICAgICBpZiAoZXZlbnRGaWx0ZXIpIHtcbiAgICAgICAgICAgIHRyaWdnZXJTcGVjLmV2ZW50RmlsdGVyID0gZXZlbnRGaWx0ZXJcbiAgICAgICAgICB9XG4gICAgICAgICAgd2hpbGUgKHRva2Vucy5sZW5ndGggPiAwICYmIHRva2Vuc1swXSAhPT0gJywnKSB7XG4gICAgICAgICAgICBjb25zdW1lVW50aWwodG9rZW5zLCBOT1RfV0hJVEVTUEFDRSlcbiAgICAgICAgICAgIGNvbnN0IHRva2VuID0gdG9rZW5zLnNoaWZ0KClcbiAgICAgICAgICAgIGlmICh0b2tlbiA9PT0gJ2NoYW5nZWQnKSB7XG4gICAgICAgICAgICAgIHRyaWdnZXJTcGVjLmNoYW5nZWQgPSB0cnVlXG4gICAgICAgICAgICB9IGVsc2UgaWYgKHRva2VuID09PSAnb25jZScpIHtcbiAgICAgICAgICAgICAgdHJpZ2dlclNwZWMub25jZSA9IHRydWVcbiAgICAgICAgICAgIH0gZWxzZSBpZiAodG9rZW4gPT09ICdjb25zdW1lJykge1xuICAgICAgICAgICAgICB0cmlnZ2VyU3BlYy5jb25zdW1lID0gdHJ1ZVxuICAgICAgICAgICAgfSBlbHNlIGlmICh0b2tlbiA9PT0gJ2RlbGF5JyAmJiB0b2tlbnNbMF0gPT09ICc6Jykge1xuICAgICAgICAgICAgICB0b2tlbnMuc2hpZnQoKVxuICAgICAgICAgICAgICB0cmlnZ2VyU3BlYy5kZWxheSA9IHBhcnNlSW50ZXJ2YWwoY29uc3VtZVVudGlsKHRva2VucywgV0hJVEVTUEFDRV9PUl9DT01NQSkpXG4gICAgICAgICAgICB9IGVsc2UgaWYgKHRva2VuID09PSAnZnJvbScgJiYgdG9rZW5zWzBdID09PSAnOicpIHtcbiAgICAgICAgICAgICAgdG9rZW5zLnNoaWZ0KClcbiAgICAgICAgICAgICAgaWYgKENPTUJJTkVEX1NFTEVDVE9SX1NUQVJULnRlc3QodG9rZW5zWzBdKSkge1xuICAgICAgICAgICAgICAgIHZhciBmcm9tX2FyZyA9IGNvbnN1bWVDU1NTZWxlY3Rvcih0b2tlbnMpXG4gICAgICAgICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgICAgICAgdmFyIGZyb21fYXJnID0gY29uc3VtZVVudGlsKHRva2VucywgV0hJVEVTUEFDRV9PUl9DT01NQSlcbiAgICAgICAgICAgICAgICBpZiAoZnJvbV9hcmcgPT09ICdjbG9zZXN0JyB8fCBmcm9tX2FyZyA9PT0gJ2ZpbmQnIHx8IGZyb21fYXJnID09PSAnbmV4dCcgfHwgZnJvbV9hcmcgPT09ICdwcmV2aW91cycpIHtcbiAgICAgICAgICAgICAgICAgIHRva2Vucy5zaGlmdCgpXG4gICAgICAgICAgICAgICAgICBjb25zdCBzZWxlY3RvciA9IGNvbnN1bWVDU1NTZWxlY3Rvcih0b2tlbnMpXG4gICAgICAgICAgICAgICAgICAvLyBgbmV4dGAgYW5kIGBwcmV2aW91c2AgYWxsb3cgYSBzZWxlY3Rvci1sZXNzIHN5bnRheFxuICAgICAgICAgICAgICAgICAgaWYgKHNlbGVjdG9yLmxlbmd0aCA+IDApIHtcbiAgICAgICAgICAgICAgICAgICAgZnJvbV9hcmcgKz0gJyAnICsgc2VsZWN0b3JcbiAgICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICAgIH1cbiAgICAgICAgICAgICAgdHJpZ2dlclNwZWMuZnJvbSA9IGZyb21fYXJnXG4gICAgICAgICAgICB9IGVsc2UgaWYgKHRva2VuID09PSAndGFyZ2V0JyAmJiB0b2tlbnNbMF0gPT09ICc6Jykge1xuICAgICAgICAgICAgICB0b2tlbnMuc2hpZnQoKVxuICAgICAgICAgICAgICB0cmlnZ2VyU3BlYy50YXJnZXQgPSBjb25zdW1lQ1NTU2VsZWN0b3IodG9rZW5zKVxuICAgICAgICAgICAgfSBlbHNlIGlmICh0b2tlbiA9PT0gJ3Rocm90dGxlJyAmJiB0b2tlbnNbMF0gPT09ICc6Jykge1xuICAgICAgICAgICAgICB0b2tlbnMuc2hpZnQoKVxuICAgICAgICAgICAgICB0cmlnZ2VyU3BlYy50aHJvdHRsZSA9IHBhcnNlSW50ZXJ2YWwoY29uc3VtZVVudGlsKHRva2VucywgV0hJVEVTUEFDRV9PUl9DT01NQSkpXG4gICAgICAgICAgICB9IGVsc2UgaWYgKHRva2VuID09PSAncXVldWUnICYmIHRva2Vuc1swXSA9PT0gJzonKSB7XG4gICAgICAgICAgICAgIHRva2Vucy5zaGlmdCgpXG4gICAgICAgICAgICAgIHRyaWdnZXJTcGVjLnF1ZXVlID0gY29uc3VtZVVudGlsKHRva2VucywgV0hJVEVTUEFDRV9PUl9DT01NQSlcbiAgICAgICAgICAgIH0gZWxzZSBpZiAodG9rZW4gPT09ICdyb290JyAmJiB0b2tlbnNbMF0gPT09ICc6Jykge1xuICAgICAgICAgICAgICB0b2tlbnMuc2hpZnQoKVxuICAgICAgICAgICAgICB0cmlnZ2VyU3BlY1t0b2tlbl0gPSBjb25zdW1lQ1NTU2VsZWN0b3IodG9rZW5zKVxuICAgICAgICAgICAgfSBlbHNlIGlmICh0b2tlbiA9PT0gJ3RocmVzaG9sZCcgJiYgdG9rZW5zWzBdID09PSAnOicpIHtcbiAgICAgICAgICAgICAgdG9rZW5zLnNoaWZ0KClcbiAgICAgICAgICAgICAgdHJpZ2dlclNwZWNbdG9rZW5dID0gY29uc3VtZVVudGlsKHRva2VucywgV0hJVEVTUEFDRV9PUl9DT01NQSlcbiAgICAgICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGVsdCwgJ2h0bXg6c3ludGF4OmVycm9yJywgeyB0b2tlbjogdG9rZW5zLnNoaWZ0KCkgfSlcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgICAgdHJpZ2dlclNwZWNzLnB1c2godHJpZ2dlclNwZWMpXG4gICAgICAgIH1cbiAgICAgIH1cbiAgICAgIGlmICh0b2tlbnMubGVuZ3RoID09PSBpbml0aWFsTGVuZ3RoKSB7XG4gICAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGVsdCwgJ2h0bXg6c3ludGF4OmVycm9yJywgeyB0b2tlbjogdG9rZW5zLnNoaWZ0KCkgfSlcbiAgICAgIH1cbiAgICAgIGNvbnN1bWVVbnRpbCh0b2tlbnMsIE5PVF9XSElURVNQQUNFKVxuICAgIH0gd2hpbGUgKHRva2Vuc1swXSA9PT0gJywnICYmIHRva2Vucy5zaGlmdCgpKVxuICAgIGlmIChjYWNoZSkge1xuICAgICAgY2FjaGVbZXhwbGljaXRUcmlnZ2VyXSA9IHRyaWdnZXJTcGVjc1xuICAgIH1cbiAgICByZXR1cm4gdHJpZ2dlclNwZWNzXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHJldHVybnMge0h0bXhUcmlnZ2VyU3BlY2lmaWNhdGlvbltdfVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0VHJpZ2dlclNwZWNzKGVsdCkge1xuICAgIGNvbnN0IGV4cGxpY2l0VHJpZ2dlciA9IGdldEF0dHJpYnV0ZVZhbHVlKGVsdCwgJ2h4LXRyaWdnZXInKVxuICAgIGxldCB0cmlnZ2VyU3BlY3MgPSBbXVxuICAgIGlmIChleHBsaWNpdFRyaWdnZXIpIHtcbiAgICAgIGNvbnN0IGNhY2hlID0gaHRteC5jb25maWcudHJpZ2dlclNwZWNzQ2FjaGVcbiAgICAgIHRyaWdnZXJTcGVjcyA9IChjYWNoZSAmJiBjYWNoZVtleHBsaWNpdFRyaWdnZXJdKSB8fCBwYXJzZUFuZENhY2hlVHJpZ2dlcihlbHQsIGV4cGxpY2l0VHJpZ2dlciwgY2FjaGUpXG4gICAgfVxuXG4gICAgaWYgKHRyaWdnZXJTcGVjcy5sZW5ndGggPiAwKSB7XG4gICAgICByZXR1cm4gdHJpZ2dlclNwZWNzXG4gICAgfSBlbHNlIGlmIChtYXRjaGVzKGVsdCwgJ2Zvcm0nKSkge1xuICAgICAgcmV0dXJuIFt7IHRyaWdnZXI6ICdzdWJtaXQnIH1dXG4gICAgfSBlbHNlIGlmIChtYXRjaGVzKGVsdCwgJ2lucHV0W3R5cGU9XCJidXR0b25cIl0sIGlucHV0W3R5cGU9XCJzdWJtaXRcIl0nKSkge1xuICAgICAgcmV0dXJuIFt7IHRyaWdnZXI6ICdjbGljaycgfV1cbiAgICB9IGVsc2UgaWYgKG1hdGNoZXMoZWx0LCBJTlBVVF9TRUxFQ1RPUikpIHtcbiAgICAgIHJldHVybiBbeyB0cmlnZ2VyOiAnY2hhbmdlJyB9XVxuICAgIH0gZWxzZSB7XG4gICAgICByZXR1cm4gW3sgdHJpZ2dlcjogJ2NsaWNrJyB9XVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKi9cbiAgZnVuY3Rpb24gY2FuY2VsUG9sbGluZyhlbHQpIHtcbiAgICBnZXRJbnRlcm5hbERhdGEoZWx0KS5jYW5jZWxsZWQgPSB0cnVlXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHBhcmFtIHtUcmlnZ2VySGFuZGxlcn0gaGFuZGxlclxuICAgKiBAcGFyYW0ge0h0bXhUcmlnZ2VyU3BlY2lmaWNhdGlvbn0gc3BlY1xuICAgKi9cbiAgZnVuY3Rpb24gcHJvY2Vzc1BvbGxpbmcoZWx0LCBoYW5kbGVyLCBzcGVjKSB7XG4gICAgY29uc3Qgbm9kZURhdGEgPSBnZXRJbnRlcm5hbERhdGEoZWx0KVxuICAgIG5vZGVEYXRhLnRpbWVvdXQgPSBnZXRXaW5kb3coKS5zZXRUaW1lb3V0KGZ1bmN0aW9uKCkge1xuICAgICAgaWYgKGJvZHlDb250YWlucyhlbHQpICYmIG5vZGVEYXRhLmNhbmNlbGxlZCAhPT0gdHJ1ZSkge1xuICAgICAgICBpZiAoIW1heWJlRmlsdGVyRXZlbnQoc3BlYywgZWx0LCBtYWtlRXZlbnQoJ2h4OnBvbGw6dHJpZ2dlcicsIHtcbiAgICAgICAgICB0cmlnZ2VyU3BlYzogc3BlYyxcbiAgICAgICAgICB0YXJnZXQ6IGVsdFxuICAgICAgICB9KSkpIHtcbiAgICAgICAgICBoYW5kbGVyKGVsdClcbiAgICAgICAgfVxuICAgICAgICBwcm9jZXNzUG9sbGluZyhlbHQsIGhhbmRsZXIsIHNwZWMpXG4gICAgICB9XG4gICAgfSwgc3BlYy5wb2xsSW50ZXJ2YWwpXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtIVE1MQW5jaG9yRWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm5zIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gaXNMb2NhbExpbmsoZWx0KSB7XG4gICAgcmV0dXJuIGxvY2F0aW9uLmhvc3RuYW1lID09PSBlbHQuaG9zdG5hbWUgJiZcbiAgICAgIGdldFJhd0F0dHJpYnV0ZShlbHQsICdocmVmJykgJiZcbiAgICAgIGdldFJhd0F0dHJpYnV0ZShlbHQsICdocmVmJykuaW5kZXhPZignIycpICE9PSAwXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICovXG4gIGZ1bmN0aW9uIGVsdElzRGlzYWJsZWQoZWx0KSB7XG4gICAgcmV0dXJuIGNsb3Nlc3QoZWx0LCBodG14LmNvbmZpZy5kaXNhYmxlU2VsZWN0b3IpXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHBhcmFtIHtIdG14Tm9kZUludGVybmFsRGF0YX0gbm9kZURhdGFcbiAgICogQHBhcmFtIHtIdG14VHJpZ2dlclNwZWNpZmljYXRpb25bXX0gdHJpZ2dlclNwZWNzXG4gICAqL1xuICBmdW5jdGlvbiBib29zdEVsZW1lbnQoZWx0LCBub2RlRGF0YSwgdHJpZ2dlclNwZWNzKSB7XG4gICAgaWYgKChlbHQgaW5zdGFuY2VvZiBIVE1MQW5jaG9yRWxlbWVudCAmJiBpc0xvY2FsTGluayhlbHQpICYmIChlbHQudGFyZ2V0ID09PSAnJyB8fCBlbHQudGFyZ2V0ID09PSAnX3NlbGYnKSkgfHwgKGVsdC50YWdOYW1lID09PSAnRk9STScgJiYgU3RyaW5nKGdldFJhd0F0dHJpYnV0ZShlbHQsICdtZXRob2QnKSkudG9Mb3dlckNhc2UoKSAhPT0gJ2RpYWxvZycpKSB7XG4gICAgICBub2RlRGF0YS5ib29zdGVkID0gdHJ1ZVxuICAgICAgbGV0IHZlcmIsIHBhdGhcbiAgICAgIGlmIChlbHQudGFnTmFtZSA9PT0gJ0EnKSB7XG4gICAgICAgIHZlcmIgPSAnZ2V0J1xuICAgICAgICBwYXRoID0gZ2V0UmF3QXR0cmlidXRlKGVsdCwgJ2hyZWYnKVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgY29uc3QgcmF3QXR0cmlidXRlID0gZ2V0UmF3QXR0cmlidXRlKGVsdCwgJ21ldGhvZCcpXG4gICAgICAgIHZlcmIgPSByYXdBdHRyaWJ1dGUgPyByYXdBdHRyaWJ1dGUudG9Mb3dlckNhc2UoKSA6ICdnZXQnXG4gICAgICAgIGlmICh2ZXJiID09PSAnZ2V0Jykge1xuICAgICAgICB9XG4gICAgICAgIHBhdGggPSBnZXRSYXdBdHRyaWJ1dGUoZWx0LCAnYWN0aW9uJylcbiAgICAgIH1cbiAgICAgIHRyaWdnZXJTcGVjcy5mb3JFYWNoKGZ1bmN0aW9uKHRyaWdnZXJTcGVjKSB7XG4gICAgICAgIGFkZEV2ZW50TGlzdGVuZXIoZWx0LCBmdW5jdGlvbihub2RlLCBldnQpIHtcbiAgICAgICAgICBjb25zdCBlbHQgPSBhc0VsZW1lbnQobm9kZSlcbiAgICAgICAgICBpZiAoZWx0SXNEaXNhYmxlZChlbHQpKSB7XG4gICAgICAgICAgICBjbGVhblVwRWxlbWVudChlbHQpXG4gICAgICAgICAgICByZXR1cm5cbiAgICAgICAgICB9XG4gICAgICAgICAgaXNzdWVBamF4UmVxdWVzdCh2ZXJiLCBwYXRoLCBlbHQsIGV2dClcbiAgICAgICAgfSwgbm9kZURhdGEsIHRyaWdnZXJTcGVjLCB0cnVlKVxuICAgICAgfSlcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFdmVudH0gZXZ0XG4gICAqIEBwYXJhbSB7Tm9kZX0gbm9kZVxuICAgKiBAcmV0dXJucyB7Ym9vbGVhbn1cbiAgICovXG4gIGZ1bmN0aW9uIHNob3VsZENhbmNlbChldnQsIG5vZGUpIHtcbiAgICBjb25zdCBlbHQgPSBhc0VsZW1lbnQobm9kZSlcbiAgICBpZiAoIWVsdCkge1xuICAgICAgcmV0dXJuIGZhbHNlXG4gICAgfVxuICAgIGlmIChldnQudHlwZSA9PT0gJ3N1Ym1pdCcgfHwgZXZ0LnR5cGUgPT09ICdjbGljaycpIHtcbiAgICAgIGlmIChlbHQudGFnTmFtZSA9PT0gJ0ZPUk0nKSB7XG4gICAgICAgIHJldHVybiB0cnVlXG4gICAgICB9XG4gICAgICBpZiAobWF0Y2hlcyhlbHQsICdpbnB1dFt0eXBlPVwic3VibWl0XCJdLCBidXR0b24nKSAmJiBjbG9zZXN0KGVsdCwgJ2Zvcm0nKSAhPT0gbnVsbCkge1xuICAgICAgICByZXR1cm4gdHJ1ZVxuICAgICAgfVxuICAgICAgaWYgKGVsdCBpbnN0YW5jZW9mIEhUTUxBbmNob3JFbGVtZW50ICYmIGVsdC5ocmVmICYmXG4gICAgICAgIChlbHQuZ2V0QXR0cmlidXRlKCdocmVmJykgPT09ICcjJyB8fCBlbHQuZ2V0QXR0cmlidXRlKCdocmVmJykuaW5kZXhPZignIycpICE9PSAwKSkge1xuICAgICAgICByZXR1cm4gdHJ1ZVxuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gZmFsc2VcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKiBAcGFyYW0ge0V2ZW50fE1vdXNlRXZlbnR8S2V5Ym9hcmRFdmVudHxUb3VjaEV2ZW50fSBldnRcbiAgICogQHJldHVybnMge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiBpZ25vcmVCb29zdGVkQW5jaG9yQ3RybENsaWNrKGVsdCwgZXZ0KSB7XG4gICAgcmV0dXJuIGdldEludGVybmFsRGF0YShlbHQpLmJvb3N0ZWQgJiYgZWx0IGluc3RhbmNlb2YgSFRNTEFuY2hvckVsZW1lbnQgJiYgZXZ0LnR5cGUgPT09ICdjbGljaycgJiZcbiAgICAgIC8vIEB0cy1pZ25vcmUgdGhpcyB3aWxsIHJlc29sdmUgdG8gdW5kZWZpbmVkIGZvciBldmVudHMgdGhhdCBkb24ndCBkZWZpbmUgdGhvc2UgcHJvcGVydGllcywgd2hpY2ggaXMgZmluZVxuICAgICAgKGV2dC5jdHJsS2V5IHx8IGV2dC5tZXRhS2V5KVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7SHRteFRyaWdnZXJTcGVjaWZpY2F0aW9ufSB0cmlnZ2VyU3BlY1xuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKiBAcGFyYW0ge0V2ZW50fSBldnRcbiAgICogQHJldHVybnMge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiBtYXliZUZpbHRlckV2ZW50KHRyaWdnZXJTcGVjLCBlbHQsIGV2dCkge1xuICAgIGNvbnN0IGV2ZW50RmlsdGVyID0gdHJpZ2dlclNwZWMuZXZlbnRGaWx0ZXJcbiAgICBpZiAoZXZlbnRGaWx0ZXIpIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIHJldHVybiBldmVudEZpbHRlci5jYWxsKGVsdCwgZXZ0KSAhPT0gdHJ1ZVxuICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICBjb25zdCBzb3VyY2UgPSBldmVudEZpbHRlci5zb3VyY2VcbiAgICAgICAgdHJpZ2dlckVycm9yRXZlbnQoZ2V0RG9jdW1lbnQoKS5ib2R5LCAnaHRteDpldmVudEZpbHRlcjplcnJvcicsIHsgZXJyb3I6IGUsIHNvdXJjZSB9KVxuICAgICAgICByZXR1cm4gdHJ1ZVxuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gZmFsc2VcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKiBAcGFyYW0ge1RyaWdnZXJIYW5kbGVyfSBoYW5kbGVyXG4gICAqIEBwYXJhbSB7SHRteE5vZGVJbnRlcm5hbERhdGF9IG5vZGVEYXRhXG4gICAqIEBwYXJhbSB7SHRteFRyaWdnZXJTcGVjaWZpY2F0aW9ufSB0cmlnZ2VyU3BlY1xuICAgKiBAcGFyYW0ge2Jvb2xlYW59IFtleHBsaWNpdENhbmNlbF1cbiAgICovXG4gIGZ1bmN0aW9uIGFkZEV2ZW50TGlzdGVuZXIoZWx0LCBoYW5kbGVyLCBub2RlRGF0YSwgdHJpZ2dlclNwZWMsIGV4cGxpY2l0Q2FuY2VsKSB7XG4gICAgY29uc3QgZWxlbWVudERhdGEgPSBnZXRJbnRlcm5hbERhdGEoZWx0KVxuICAgIC8qKiBAdHlwZSB7KE5vZGV8V2luZG93KVtdfSAqL1xuICAgIGxldCBlbHRzVG9MaXN0ZW5PblxuICAgIGlmICh0cmlnZ2VyU3BlYy5mcm9tKSB7XG4gICAgICBlbHRzVG9MaXN0ZW5PbiA9IHF1ZXJ5U2VsZWN0b3JBbGxFeHQoZWx0LCB0cmlnZ2VyU3BlYy5mcm9tKVxuICAgIH0gZWxzZSB7XG4gICAgICBlbHRzVG9MaXN0ZW5PbiA9IFtlbHRdXG4gICAgfVxuICAgIC8vIHN0b3JlIHRoZSBpbml0aWFsIHZhbHVlcyBvZiB0aGUgZWxlbWVudHMsIHNvIHdlIGNhbiB0ZWxsIGlmIHRoZXkgY2hhbmdlXG4gICAgaWYgKHRyaWdnZXJTcGVjLmNoYW5nZWQpIHtcbiAgICAgIGVsdHNUb0xpc3Rlbk9uLmZvckVhY2goZnVuY3Rpb24oZWx0VG9MaXN0ZW5Pbikge1xuICAgICAgICBjb25zdCBlbHRUb0xpc3Rlbk9uRGF0YSA9IGdldEludGVybmFsRGF0YShlbHRUb0xpc3Rlbk9uKVxuICAgICAgICAvLyBAdHMtaWdub3JlIHZhbHVlIHdpbGwgYmUgdW5kZWZpbmVkIGZvciBub24taW5wdXQgZWxlbWVudHMsIHdoaWNoIGlzIGZpbmVcbiAgICAgICAgZWx0VG9MaXN0ZW5PbkRhdGEubGFzdFZhbHVlID0gZWx0VG9MaXN0ZW5Pbi52YWx1ZVxuICAgICAgfSlcbiAgICB9XG4gICAgZm9yRWFjaChlbHRzVG9MaXN0ZW5PbiwgZnVuY3Rpb24oZWx0VG9MaXN0ZW5Pbikge1xuICAgICAgLyoqIEB0eXBlIEV2ZW50TGlzdGVuZXIgKi9cbiAgICAgIGNvbnN0IGV2ZW50TGlzdGVuZXIgPSBmdW5jdGlvbihldnQpIHtcbiAgICAgICAgaWYgKCFib2R5Q29udGFpbnMoZWx0KSkge1xuICAgICAgICAgIGVsdFRvTGlzdGVuT24ucmVtb3ZlRXZlbnRMaXN0ZW5lcih0cmlnZ2VyU3BlYy50cmlnZ2VyLCBldmVudExpc3RlbmVyKVxuICAgICAgICAgIHJldHVyblxuICAgICAgICB9XG4gICAgICAgIGlmIChpZ25vcmVCb29zdGVkQW5jaG9yQ3RybENsaWNrKGVsdCwgZXZ0KSkge1xuICAgICAgICAgIHJldHVyblxuICAgICAgICB9XG4gICAgICAgIGlmIChleHBsaWNpdENhbmNlbCB8fCBzaG91bGRDYW5jZWwoZXZ0LCBlbHQpKSB7XG4gICAgICAgICAgZXZ0LnByZXZlbnREZWZhdWx0KClcbiAgICAgICAgfVxuICAgICAgICBpZiAobWF5YmVGaWx0ZXJFdmVudCh0cmlnZ2VyU3BlYywgZWx0LCBldnQpKSB7XG4gICAgICAgICAgcmV0dXJuXG4gICAgICAgIH1cbiAgICAgICAgY29uc3QgZXZlbnREYXRhID0gZ2V0SW50ZXJuYWxEYXRhKGV2dClcbiAgICAgICAgZXZlbnREYXRhLnRyaWdnZXJTcGVjID0gdHJpZ2dlclNwZWNcbiAgICAgICAgaWYgKGV2ZW50RGF0YS5oYW5kbGVkRm9yID09IG51bGwpIHtcbiAgICAgICAgICBldmVudERhdGEuaGFuZGxlZEZvciA9IFtdXG4gICAgICAgIH1cbiAgICAgICAgaWYgKGV2ZW50RGF0YS5oYW5kbGVkRm9yLmluZGV4T2YoZWx0KSA8IDApIHtcbiAgICAgICAgICBldmVudERhdGEuaGFuZGxlZEZvci5wdXNoKGVsdClcbiAgICAgICAgICBpZiAodHJpZ2dlclNwZWMuY29uc3VtZSkge1xuICAgICAgICAgICAgZXZ0LnN0b3BQcm9wYWdhdGlvbigpXG4gICAgICAgICAgfVxuICAgICAgICAgIGlmICh0cmlnZ2VyU3BlYy50YXJnZXQgJiYgZXZ0LnRhcmdldCkge1xuICAgICAgICAgICAgaWYgKCFtYXRjaGVzKGFzRWxlbWVudChldnQudGFyZ2V0KSwgdHJpZ2dlclNwZWMudGFyZ2V0KSkge1xuICAgICAgICAgICAgICByZXR1cm5cbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgICAgaWYgKHRyaWdnZXJTcGVjLm9uY2UpIHtcbiAgICAgICAgICAgIGlmIChlbGVtZW50RGF0YS50cmlnZ2VyZWRPbmNlKSB7XG4gICAgICAgICAgICAgIHJldHVyblxuICAgICAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgICAgZWxlbWVudERhdGEudHJpZ2dlcmVkT25jZSA9IHRydWVcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9XG4gICAgICAgICAgaWYgKHRyaWdnZXJTcGVjLmNoYW5nZWQpIHtcbiAgICAgICAgICAgIGNvbnN0IGVsdFRvTGlzdGVuT25EYXRhID0gZ2V0SW50ZXJuYWxEYXRhKGVsdFRvTGlzdGVuT24pXG4gICAgICAgICAgICAvLyBAdHMtaWdub3JlIHZhbHVlIHdpbGwgYmUgdW5kZWZpbmVkIGZvciBub24taW5wdXQgZWxlbWVudHMsIHdoaWNoIGlzIGZpbmVcbiAgICAgICAgICAgIGNvbnN0IHZhbHVlID0gZWx0VG9MaXN0ZW5Pbi52YWx1ZVxuICAgICAgICAgICAgaWYgKGVsdFRvTGlzdGVuT25EYXRhLmxhc3RWYWx1ZSA9PT0gdmFsdWUpIHtcbiAgICAgICAgICAgICAgcmV0dXJuXG4gICAgICAgICAgICB9XG4gICAgICAgICAgICBlbHRUb0xpc3Rlbk9uRGF0YS5sYXN0VmFsdWUgPSB2YWx1ZVxuICAgICAgICAgIH1cbiAgICAgICAgICBpZiAoZWxlbWVudERhdGEuZGVsYXllZCkge1xuICAgICAgICAgICAgY2xlYXJUaW1lb3V0KGVsZW1lbnREYXRhLmRlbGF5ZWQpXG4gICAgICAgICAgfVxuICAgICAgICAgIGlmIChlbGVtZW50RGF0YS50aHJvdHRsZSkge1xuICAgICAgICAgICAgcmV0dXJuXG4gICAgICAgICAgfVxuXG4gICAgICAgICAgaWYgKHRyaWdnZXJTcGVjLnRocm90dGxlID4gMCkge1xuICAgICAgICAgICAgaWYgKCFlbGVtZW50RGF0YS50aHJvdHRsZSkge1xuICAgICAgICAgICAgICB0cmlnZ2VyRXZlbnQoZWx0LCAnaHRteDp0cmlnZ2VyJylcbiAgICAgICAgICAgICAgaGFuZGxlcihlbHQsIGV2dClcbiAgICAgICAgICAgICAgZWxlbWVudERhdGEudGhyb3R0bGUgPSBnZXRXaW5kb3coKS5zZXRUaW1lb3V0KGZ1bmN0aW9uKCkge1xuICAgICAgICAgICAgICAgIGVsZW1lbnREYXRhLnRocm90dGxlID0gbnVsbFxuICAgICAgICAgICAgICB9LCB0cmlnZ2VyU3BlYy50aHJvdHRsZSlcbiAgICAgICAgICAgIH1cbiAgICAgICAgICB9IGVsc2UgaWYgKHRyaWdnZXJTcGVjLmRlbGF5ID4gMCkge1xuICAgICAgICAgICAgZWxlbWVudERhdGEuZGVsYXllZCA9IGdldFdpbmRvdygpLnNldFRpbWVvdXQoZnVuY3Rpb24oKSB7XG4gICAgICAgICAgICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14OnRyaWdnZXInKVxuICAgICAgICAgICAgICBoYW5kbGVyKGVsdCwgZXZ0KVxuICAgICAgICAgICAgfSwgdHJpZ2dlclNwZWMuZGVsYXkpXG4gICAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14OnRyaWdnZXInKVxuICAgICAgICAgICAgaGFuZGxlcihlbHQsIGV2dClcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICAgIGlmIChub2RlRGF0YS5saXN0ZW5lckluZm9zID09IG51bGwpIHtcbiAgICAgICAgbm9kZURhdGEubGlzdGVuZXJJbmZvcyA9IFtdXG4gICAgICB9XG4gICAgICBub2RlRGF0YS5saXN0ZW5lckluZm9zLnB1c2goe1xuICAgICAgICB0cmlnZ2VyOiB0cmlnZ2VyU3BlYy50cmlnZ2VyLFxuICAgICAgICBsaXN0ZW5lcjogZXZlbnRMaXN0ZW5lcixcbiAgICAgICAgb246IGVsdFRvTGlzdGVuT25cbiAgICAgIH0pXG4gICAgICBlbHRUb0xpc3Rlbk9uLmFkZEV2ZW50TGlzdGVuZXIodHJpZ2dlclNwZWMudHJpZ2dlciwgZXZlbnRMaXN0ZW5lcilcbiAgICB9KVxuICB9XG5cbiAgbGV0IHdpbmRvd0lzU2Nyb2xsaW5nID0gZmFsc2UgLy8gdXNlZCBieSBpbml0U2Nyb2xsSGFuZGxlclxuICBsZXQgc2Nyb2xsSGFuZGxlciA9IG51bGxcbiAgZnVuY3Rpb24gaW5pdFNjcm9sbEhhbmRsZXIoKSB7XG4gICAgaWYgKCFzY3JvbGxIYW5kbGVyKSB7XG4gICAgICBzY3JvbGxIYW5kbGVyID0gZnVuY3Rpb24oKSB7XG4gICAgICAgIHdpbmRvd0lzU2Nyb2xsaW5nID0gdHJ1ZVxuICAgICAgfVxuICAgICAgd2luZG93LmFkZEV2ZW50TGlzdGVuZXIoJ3Njcm9sbCcsIHNjcm9sbEhhbmRsZXIpXG4gICAgICBzZXRJbnRlcnZhbChmdW5jdGlvbigpIHtcbiAgICAgICAgaWYgKHdpbmRvd0lzU2Nyb2xsaW5nKSB7XG4gICAgICAgICAgd2luZG93SXNTY3JvbGxpbmcgPSBmYWxzZVxuICAgICAgICAgIGZvckVhY2goZ2V0RG9jdW1lbnQoKS5xdWVyeVNlbGVjdG9yQWxsKFwiW2h4LXRyaWdnZXIqPSdyZXZlYWxlZCddLFtkYXRhLWh4LXRyaWdnZXIqPSdyZXZlYWxlZCddXCIpLCBmdW5jdGlvbihlbHQpIHtcbiAgICAgICAgICAgIG1heWJlUmV2ZWFsKGVsdClcbiAgICAgICAgICB9KVxuICAgICAgICB9XG4gICAgICB9LCAyMDApXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqL1xuICBmdW5jdGlvbiBtYXliZVJldmVhbChlbHQpIHtcbiAgICBpZiAoIWhhc0F0dHJpYnV0ZShlbHQsICdkYXRhLWh4LXJldmVhbGVkJykgJiYgaXNTY3JvbGxlZEludG9WaWV3KGVsdCkpIHtcbiAgICAgIGVsdC5zZXRBdHRyaWJ1dGUoJ2RhdGEtaHgtcmV2ZWFsZWQnLCAndHJ1ZScpXG4gICAgICBjb25zdCBub2RlRGF0YSA9IGdldEludGVybmFsRGF0YShlbHQpXG4gICAgICBpZiAobm9kZURhdGEuaW5pdEhhc2gpIHtcbiAgICAgICAgdHJpZ2dlckV2ZW50KGVsdCwgJ3JldmVhbGVkJylcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIC8vIGlmIHRoZSBub2RlIGlzbid0IGluaXRpYWxpemVkLCB3YWl0IGZvciBpdCBiZWZvcmUgdHJpZ2dlcmluZyB0aGUgcmVxdWVzdFxuICAgICAgICBlbHQuYWRkRXZlbnRMaXN0ZW5lcignaHRteDphZnRlclByb2Nlc3NOb2RlJywgZnVuY3Rpb24oKSB7IHRyaWdnZXJFdmVudChlbHQsICdyZXZlYWxlZCcpIH0sIHsgb25jZTogdHJ1ZSB9KVxuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEBwYXJhbSB7VHJpZ2dlckhhbmRsZXJ9IGhhbmRsZXJcbiAgICogQHBhcmFtIHtIdG14Tm9kZUludGVybmFsRGF0YX0gbm9kZURhdGFcbiAgICogQHBhcmFtIHtudW1iZXJ9IGRlbGF5XG4gICAqL1xuICBmdW5jdGlvbiBsb2FkSW1tZWRpYXRlbHkoZWx0LCBoYW5kbGVyLCBub2RlRGF0YSwgZGVsYXkpIHtcbiAgICBjb25zdCBsb2FkID0gZnVuY3Rpb24oKSB7XG4gICAgICBpZiAoIW5vZGVEYXRhLmxvYWRlZCkge1xuICAgICAgICBub2RlRGF0YS5sb2FkZWQgPSB0cnVlXG4gICAgICAgIGhhbmRsZXIoZWx0KVxuICAgICAgfVxuICAgIH1cbiAgICBpZiAoZGVsYXkgPiAwKSB7XG4gICAgICBnZXRXaW5kb3coKS5zZXRUaW1lb3V0KGxvYWQsIGRlbGF5KVxuICAgIH0gZWxzZSB7XG4gICAgICBsb2FkKClcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHBhcmFtIHtIdG14Tm9kZUludGVybmFsRGF0YX0gbm9kZURhdGFcbiAgICogQHBhcmFtIHtIdG14VHJpZ2dlclNwZWNpZmljYXRpb25bXX0gdHJpZ2dlclNwZWNzXG4gICAqIEByZXR1cm5zIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gcHJvY2Vzc1ZlcmJzKGVsdCwgbm9kZURhdGEsIHRyaWdnZXJTcGVjcykge1xuICAgIGxldCBleHBsaWNpdEFjdGlvbiA9IGZhbHNlXG4gICAgZm9yRWFjaChWRVJCUywgZnVuY3Rpb24odmVyYikge1xuICAgICAgaWYgKGhhc0F0dHJpYnV0ZShlbHQsICdoeC0nICsgdmVyYikpIHtcbiAgICAgICAgY29uc3QgcGF0aCA9IGdldEF0dHJpYnV0ZVZhbHVlKGVsdCwgJ2h4LScgKyB2ZXJiKVxuICAgICAgICBleHBsaWNpdEFjdGlvbiA9IHRydWVcbiAgICAgICAgbm9kZURhdGEucGF0aCA9IHBhdGhcbiAgICAgICAgbm9kZURhdGEudmVyYiA9IHZlcmJcbiAgICAgICAgdHJpZ2dlclNwZWNzLmZvckVhY2goZnVuY3Rpb24odHJpZ2dlclNwZWMpIHtcbiAgICAgICAgICBhZGRUcmlnZ2VySGFuZGxlcihlbHQsIHRyaWdnZXJTcGVjLCBub2RlRGF0YSwgZnVuY3Rpb24obm9kZSwgZXZ0KSB7XG4gICAgICAgICAgICBjb25zdCBlbHQgPSBhc0VsZW1lbnQobm9kZSlcbiAgICAgICAgICAgIGlmIChjbG9zZXN0KGVsdCwgaHRteC5jb25maWcuZGlzYWJsZVNlbGVjdG9yKSkge1xuICAgICAgICAgICAgICBjbGVhblVwRWxlbWVudChlbHQpXG4gICAgICAgICAgICAgIHJldHVyblxuICAgICAgICAgICAgfVxuICAgICAgICAgICAgaXNzdWVBamF4UmVxdWVzdCh2ZXJiLCBwYXRoLCBlbHQsIGV2dClcbiAgICAgICAgICB9KVxuICAgICAgICB9KVxuICAgICAgfVxuICAgIH0pXG4gICAgcmV0dXJuIGV4cGxpY2l0QWN0aW9uXG4gIH1cblxuICAvKipcbiAgICogQGNhbGxiYWNrIFRyaWdnZXJIYW5kbGVyXG4gICAqIEBwYXJhbSB7Tm9kZX0gZWx0XG4gICAqIEBwYXJhbSB7RXZlbnR9IFtldnRdXG4gICAqL1xuXG4gIC8qKlxuICAgKiBAcGFyYW0ge05vZGV9IGVsdFxuICAgKiBAcGFyYW0ge0h0bXhUcmlnZ2VyU3BlY2lmaWNhdGlvbn0gdHJpZ2dlclNwZWNcbiAgICogQHBhcmFtIHtIdG14Tm9kZUludGVybmFsRGF0YX0gbm9kZURhdGFcbiAgICogQHBhcmFtIHtUcmlnZ2VySGFuZGxlcn0gaGFuZGxlclxuICAgKi9cbiAgZnVuY3Rpb24gYWRkVHJpZ2dlckhhbmRsZXIoZWx0LCB0cmlnZ2VyU3BlYywgbm9kZURhdGEsIGhhbmRsZXIpIHtcbiAgICBpZiAodHJpZ2dlclNwZWMudHJpZ2dlciA9PT0gJ3JldmVhbGVkJykge1xuICAgICAgaW5pdFNjcm9sbEhhbmRsZXIoKVxuICAgICAgYWRkRXZlbnRMaXN0ZW5lcihlbHQsIGhhbmRsZXIsIG5vZGVEYXRhLCB0cmlnZ2VyU3BlYylcbiAgICAgIG1heWJlUmV2ZWFsKGFzRWxlbWVudChlbHQpKVxuICAgIH0gZWxzZSBpZiAodHJpZ2dlclNwZWMudHJpZ2dlciA9PT0gJ2ludGVyc2VjdCcpIHtcbiAgICAgIGNvbnN0IG9ic2VydmVyT3B0aW9ucyA9IHt9XG4gICAgICBpZiAodHJpZ2dlclNwZWMucm9vdCkge1xuICAgICAgICBvYnNlcnZlck9wdGlvbnMucm9vdCA9IHF1ZXJ5U2VsZWN0b3JFeHQoZWx0LCB0cmlnZ2VyU3BlYy5yb290KVxuICAgICAgfVxuICAgICAgaWYgKHRyaWdnZXJTcGVjLnRocmVzaG9sZCkge1xuICAgICAgICBvYnNlcnZlck9wdGlvbnMudGhyZXNob2xkID0gcGFyc2VGbG9hdCh0cmlnZ2VyU3BlYy50aHJlc2hvbGQpXG4gICAgICB9XG4gICAgICBjb25zdCBvYnNlcnZlciA9IG5ldyBJbnRlcnNlY3Rpb25PYnNlcnZlcihmdW5jdGlvbihlbnRyaWVzKSB7XG4gICAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgZW50cmllcy5sZW5ndGg7IGkrKykge1xuICAgICAgICAgIGNvbnN0IGVudHJ5ID0gZW50cmllc1tpXVxuICAgICAgICAgIGlmIChlbnRyeS5pc0ludGVyc2VjdGluZykge1xuICAgICAgICAgICAgdHJpZ2dlckV2ZW50KGVsdCwgJ2ludGVyc2VjdCcpXG4gICAgICAgICAgICBicmVha1xuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgfSwgb2JzZXJ2ZXJPcHRpb25zKVxuICAgICAgb2JzZXJ2ZXIub2JzZXJ2ZShhc0VsZW1lbnQoZWx0KSlcbiAgICAgIGFkZEV2ZW50TGlzdGVuZXIoYXNFbGVtZW50KGVsdCksIGhhbmRsZXIsIG5vZGVEYXRhLCB0cmlnZ2VyU3BlYylcbiAgICB9IGVsc2UgaWYgKHRyaWdnZXJTcGVjLnRyaWdnZXIgPT09ICdsb2FkJykge1xuICAgICAgaWYgKCFtYXliZUZpbHRlckV2ZW50KHRyaWdnZXJTcGVjLCBlbHQsIG1ha2VFdmVudCgnbG9hZCcsIHsgZWx0IH0pKSkge1xuICAgICAgICBsb2FkSW1tZWRpYXRlbHkoYXNFbGVtZW50KGVsdCksIGhhbmRsZXIsIG5vZGVEYXRhLCB0cmlnZ2VyU3BlYy5kZWxheSlcbiAgICAgIH1cbiAgICB9IGVsc2UgaWYgKHRyaWdnZXJTcGVjLnBvbGxJbnRlcnZhbCA+IDApIHtcbiAgICAgIG5vZGVEYXRhLnBvbGxpbmcgPSB0cnVlXG4gICAgICBwcm9jZXNzUG9sbGluZyhhc0VsZW1lbnQoZWx0KSwgaGFuZGxlciwgdHJpZ2dlclNwZWMpXG4gICAgfSBlbHNlIHtcbiAgICAgIGFkZEV2ZW50TGlzdGVuZXIoZWx0LCBoYW5kbGVyLCBub2RlRGF0YSwgdHJpZ2dlclNwZWMpXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gbm9kZVxuICAgKiBAcmV0dXJucyB7Ym9vbGVhbn1cbiAgICovXG4gIGZ1bmN0aW9uIHNob3VsZFByb2Nlc3NIeE9uKG5vZGUpIHtcbiAgICBjb25zdCBlbHQgPSBhc0VsZW1lbnQobm9kZSlcbiAgICBpZiAoIWVsdCkge1xuICAgICAgcmV0dXJuIGZhbHNlXG4gICAgfVxuICAgIGNvbnN0IGF0dHJpYnV0ZXMgPSBlbHQuYXR0cmlidXRlc1xuICAgIGZvciAobGV0IGogPSAwOyBqIDwgYXR0cmlidXRlcy5sZW5ndGg7IGorKykge1xuICAgICAgY29uc3QgYXR0ck5hbWUgPSBhdHRyaWJ1dGVzW2pdLm5hbWVcbiAgICAgIGlmIChzdGFydHNXaXRoKGF0dHJOYW1lLCAnaHgtb246JykgfHwgc3RhcnRzV2l0aChhdHRyTmFtZSwgJ2RhdGEtaHgtb246JykgfHxcbiAgICAgICAgc3RhcnRzV2l0aChhdHRyTmFtZSwgJ2h4LW9uLScpIHx8IHN0YXJ0c1dpdGgoYXR0ck5hbWUsICdkYXRhLWh4LW9uLScpKSB7XG4gICAgICAgIHJldHVybiB0cnVlXG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiBmYWxzZVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Tm9kZX0gZWx0XG4gICAqIEByZXR1cm5zIHtFbGVtZW50W119XG4gICAqL1xuICBjb25zdCBIWF9PTl9RVUVSWSA9IG5ldyBYUGF0aEV2YWx1YXRvcigpXG4gICAgLmNyZWF0ZUV4cHJlc3Npb24oJy4vLypbQCpbIHN0YXJ0cy13aXRoKG5hbWUoKSwgXCJoeC1vbjpcIikgb3Igc3RhcnRzLXdpdGgobmFtZSgpLCBcImRhdGEtaHgtb246XCIpIG9yJyArXG4gICAgICAnIHN0YXJ0cy13aXRoKG5hbWUoKSwgXCJoeC1vbi1cIikgb3Igc3RhcnRzLXdpdGgobmFtZSgpLCBcImRhdGEtaHgtb24tXCIpIF1dJylcblxuICBmdW5jdGlvbiBwcm9jZXNzSFhPblJvb3QoZWx0LCBlbGVtZW50cykge1xuICAgIGlmIChzaG91bGRQcm9jZXNzSHhPbihlbHQpKSB7XG4gICAgICBlbGVtZW50cy5wdXNoKGFzRWxlbWVudChlbHQpKVxuICAgIH1cbiAgICBjb25zdCBpdGVyID0gSFhfT05fUVVFUlkuZXZhbHVhdGUoZWx0KVxuICAgIGxldCBub2RlID0gbnVsbFxuICAgIHdoaWxlIChub2RlID0gaXRlci5pdGVyYXRlTmV4dCgpKSBlbGVtZW50cy5wdXNoKGFzRWxlbWVudChub2RlKSlcbiAgfVxuXG4gIGZ1bmN0aW9uIGZpbmRIeE9uV2lsZGNhcmRFbGVtZW50cyhlbHQpIHtcbiAgICAvKiogQHR5cGUge0VsZW1lbnRbXX0gKi9cbiAgICBjb25zdCBlbGVtZW50cyA9IFtdXG4gICAgaWYgKGVsdCBpbnN0YW5jZW9mIERvY3VtZW50RnJhZ21lbnQpIHtcbiAgICAgIGZvciAoY29uc3QgY2hpbGQgb2YgZWx0LmNoaWxkTm9kZXMpIHtcbiAgICAgICAgcHJvY2Vzc0hYT25Sb290KGNoaWxkLCBlbGVtZW50cylcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgcHJvY2Vzc0hYT25Sb290KGVsdCwgZWxlbWVudHMpXG4gICAgfVxuICAgIHJldHVybiBlbGVtZW50c1xuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm5zIHtOb2RlTGlzdE9mPEVsZW1lbnQ+fFtdfVxuICAgKi9cbiAgZnVuY3Rpb24gZmluZEVsZW1lbnRzVG9Qcm9jZXNzKGVsdCkge1xuICAgIGlmIChlbHQucXVlcnlTZWxlY3RvckFsbCkge1xuICAgICAgY29uc3QgYm9vc3RlZFNlbGVjdG9yID0gJywgW2h4LWJvb3N0XSBhLCBbZGF0YS1oeC1ib29zdF0gYSwgYVtoeC1ib29zdF0sIGFbZGF0YS1oeC1ib29zdF0nXG5cbiAgICAgIGNvbnN0IGV4dGVuc2lvblNlbGVjdG9ycyA9IFtdXG4gICAgICBmb3IgKGNvbnN0IGUgaW4gZXh0ZW5zaW9ucykge1xuICAgICAgICBjb25zdCBleHRlbnNpb24gPSBleHRlbnNpb25zW2VdXG4gICAgICAgIGlmIChleHRlbnNpb24uZ2V0U2VsZWN0b3JzKSB7XG4gICAgICAgICAgdmFyIHNlbGVjdG9ycyA9IGV4dGVuc2lvbi5nZXRTZWxlY3RvcnMoKVxuICAgICAgICAgIGlmIChzZWxlY3RvcnMpIHtcbiAgICAgICAgICAgIGV4dGVuc2lvblNlbGVjdG9ycy5wdXNoKHNlbGVjdG9ycylcbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgY29uc3QgcmVzdWx0cyA9IGVsdC5xdWVyeVNlbGVjdG9yQWxsKFZFUkJfU0VMRUNUT1IgKyBib29zdGVkU2VsZWN0b3IgKyBcIiwgZm9ybSwgW3R5cGU9J3N1Ym1pdCddLFwiICtcbiAgICAgICAgJyBbaHgtZXh0XSwgW2RhdGEtaHgtZXh0XSwgW2h4LXRyaWdnZXJdLCBbZGF0YS1oeC10cmlnZ2VyXScgKyBleHRlbnNpb25TZWxlY3RvcnMuZmxhdCgpLm1hcChzID0+ICcsICcgKyBzKS5qb2luKCcnKSlcblxuICAgICAgcmV0dXJuIHJlc3VsdHNcbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuIFtdXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEhhbmRsZSBzdWJtaXQgYnV0dG9ucy9pbnB1dHMgdGhhdCBoYXZlIHRoZSBmb3JtIGF0dHJpYnV0ZSBzZXRcbiAgICogc2VlIGh0dHBzOi8vZGV2ZWxvcGVyLm1vemlsbGEub3JnL2RvY3MvV2ViL0hUTUwvRWxlbWVudC9idXR0b25cbiAgICogQHBhcmFtIHtFdmVudH0gZXZ0XG4gICAqL1xuICBmdW5jdGlvbiBtYXliZVNldExhc3RCdXR0b25DbGlja2VkKGV2dCkge1xuICAgIGNvbnN0IGVsdCA9IC8qKiBAdHlwZSB7SFRNTEJ1dHRvbkVsZW1lbnR8SFRNTElucHV0RWxlbWVudH0gKi8gKGNsb3Nlc3QoYXNFbGVtZW50KGV2dC50YXJnZXQpLCBcImJ1dHRvbiwgaW5wdXRbdHlwZT0nc3VibWl0J11cIikpXG4gICAgY29uc3QgaW50ZXJuYWxEYXRhID0gZ2V0UmVsYXRlZEZvcm1EYXRhKGV2dClcbiAgICBpZiAoaW50ZXJuYWxEYXRhKSB7XG4gICAgICBpbnRlcm5hbERhdGEubGFzdEJ1dHRvbkNsaWNrZWQgPSBlbHRcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFdmVudH0gZXZ0XG4gICAqL1xuICBmdW5jdGlvbiBtYXliZVVuc2V0TGFzdEJ1dHRvbkNsaWNrZWQoZXZ0KSB7XG4gICAgY29uc3QgaW50ZXJuYWxEYXRhID0gZ2V0UmVsYXRlZEZvcm1EYXRhKGV2dClcbiAgICBpZiAoaW50ZXJuYWxEYXRhKSB7XG4gICAgICBpbnRlcm5hbERhdGEubGFzdEJ1dHRvbkNsaWNrZWQgPSBudWxsXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RXZlbnR9IGV2dFxuICAgKiBAcmV0dXJucyB7SHRteE5vZGVJbnRlcm5hbERhdGF8dW5kZWZpbmVkfVxuICAgKi9cbiAgZnVuY3Rpb24gZ2V0UmVsYXRlZEZvcm1EYXRhKGV2dCkge1xuICAgIGNvbnN0IGVsdCA9IGNsb3Nlc3QoYXNFbGVtZW50KGV2dC50YXJnZXQpLCBcImJ1dHRvbiwgaW5wdXRbdHlwZT0nc3VibWl0J11cIilcbiAgICBpZiAoIWVsdCkge1xuICAgICAgcmV0dXJuXG4gICAgfVxuICAgIGNvbnN0IGZvcm0gPSByZXNvbHZlVGFyZ2V0KCcjJyArIGdldFJhd0F0dHJpYnV0ZShlbHQsICdmb3JtJyksIGVsdC5nZXRSb290Tm9kZSgpKSB8fCBjbG9zZXN0KGVsdCwgJ2Zvcm0nKVxuICAgIGlmICghZm9ybSkge1xuICAgICAgcmV0dXJuXG4gICAgfVxuICAgIHJldHVybiBnZXRJbnRlcm5hbERhdGEoZm9ybSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0V2ZW50VGFyZ2V0fSBlbHRcbiAgICovXG4gIGZ1bmN0aW9uIGluaXRCdXR0b25UcmFja2luZyhlbHQpIHtcbiAgICAvLyBuZWVkIHRvIGhhbmRsZSBib3RoIGNsaWNrIGFuZCBmb2N1cyBpbjpcbiAgICAvLyAgIGZvY3VzaW4gLSBpbiBjYXNlIHNvbWVvbmUgdGFicyBpbiB0byBhIGJ1dHRvbiBhbmQgaGl0cyB0aGUgc3BhY2UgYmFyXG4gICAgLy8gICBjbGljayAtIG9uIE9TWCBidXR0b25zIGRvIG5vdCBmb2N1cyBvbiBjbGljayBzZWUgaHR0cHM6Ly9idWdzLndlYmtpdC5vcmcvc2hvd19idWcuY2dpP2lkPTEzNzI0XG4gICAgZWx0LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgbWF5YmVTZXRMYXN0QnV0dG9uQ2xpY2tlZClcbiAgICBlbHQuYWRkRXZlbnRMaXN0ZW5lcignZm9jdXNpbicsIG1heWJlU2V0TGFzdEJ1dHRvbkNsaWNrZWQpXG4gICAgZWx0LmFkZEV2ZW50TGlzdGVuZXIoJ2ZvY3Vzb3V0JywgbWF5YmVVbnNldExhc3RCdXR0b25DbGlja2VkKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEBwYXJhbSB7c3RyaW5nfSBldmVudE5hbWVcbiAgICogQHBhcmFtIHtzdHJpbmd9IGNvZGVcbiAgICovXG4gIGZ1bmN0aW9uIGFkZEh4T25FdmVudEhhbmRsZXIoZWx0LCBldmVudE5hbWUsIGNvZGUpIHtcbiAgICBjb25zdCBub2RlRGF0YSA9IGdldEludGVybmFsRGF0YShlbHQpXG4gICAgaWYgKCFBcnJheS5pc0FycmF5KG5vZGVEYXRhLm9uSGFuZGxlcnMpKSB7XG4gICAgICBub2RlRGF0YS5vbkhhbmRsZXJzID0gW11cbiAgICB9XG4gICAgbGV0IGZ1bmNcbiAgICAvKiogQHR5cGUgRXZlbnRMaXN0ZW5lciAqL1xuICAgIGNvbnN0IGxpc3RlbmVyID0gZnVuY3Rpb24oZSkge1xuICAgICAgbWF5YmVFdmFsKGVsdCwgZnVuY3Rpb24oKSB7XG4gICAgICAgIGlmIChlbHRJc0Rpc2FibGVkKGVsdCkpIHtcbiAgICAgICAgICByZXR1cm5cbiAgICAgICAgfVxuICAgICAgICBpZiAoIWZ1bmMpIHtcbiAgICAgICAgICBmdW5jID0gbmV3IEZ1bmN0aW9uKCdldmVudCcsIGNvZGUpXG4gICAgICAgIH1cbiAgICAgICAgZnVuYy5jYWxsKGVsdCwgZSlcbiAgICAgIH0pXG4gICAgfVxuICAgIGVsdC5hZGRFdmVudExpc3RlbmVyKGV2ZW50TmFtZSwgbGlzdGVuZXIpXG4gICAgbm9kZURhdGEub25IYW5kbGVycy5wdXNoKHsgZXZlbnQ6IGV2ZW50TmFtZSwgbGlzdGVuZXIgfSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKi9cbiAgZnVuY3Rpb24gcHJvY2Vzc0h4T25XaWxkY2FyZChlbHQpIHtcbiAgICAvLyB3aXBlIGFueSBwcmV2aW91cyBvbiBoYW5kbGVycyBzbyB0aGF0IHRoaXMgZnVuY3Rpb24gdGFrZXMgcHJlY2VkZW5jZVxuICAgIGRlSW5pdE9uSGFuZGxlcnMoZWx0KVxuXG4gICAgZm9yIChsZXQgaSA9IDA7IGkgPCBlbHQuYXR0cmlidXRlcy5sZW5ndGg7IGkrKykge1xuICAgICAgY29uc3QgbmFtZSA9IGVsdC5hdHRyaWJ1dGVzW2ldLm5hbWVcbiAgICAgIGNvbnN0IHZhbHVlID0gZWx0LmF0dHJpYnV0ZXNbaV0udmFsdWVcbiAgICAgIGlmIChzdGFydHNXaXRoKG5hbWUsICdoeC1vbicpIHx8IHN0YXJ0c1dpdGgobmFtZSwgJ2RhdGEtaHgtb24nKSkge1xuICAgICAgICBjb25zdCBhZnRlck9uUG9zaXRpb24gPSBuYW1lLmluZGV4T2YoJy1vbicpICsgM1xuICAgICAgICBjb25zdCBuZXh0Q2hhciA9IG5hbWUuc2xpY2UoYWZ0ZXJPblBvc2l0aW9uLCBhZnRlck9uUG9zaXRpb24gKyAxKVxuICAgICAgICBpZiAobmV4dENoYXIgPT09ICctJyB8fCBuZXh0Q2hhciA9PT0gJzonKSB7XG4gICAgICAgICAgbGV0IGV2ZW50TmFtZSA9IG5hbWUuc2xpY2UoYWZ0ZXJPblBvc2l0aW9uICsgMSlcbiAgICAgICAgICAvLyBpZiB0aGUgZXZlbnROYW1lIHN0YXJ0cyB3aXRoIGEgY29sb24gb3IgZGFzaCwgcHJlcGVuZCBcImh0bXhcIiBmb3Igc2hvcnRoYW5kIHN1cHBvcnRcbiAgICAgICAgICBpZiAoc3RhcnRzV2l0aChldmVudE5hbWUsICc6JykpIHtcbiAgICAgICAgICAgIGV2ZW50TmFtZSA9ICdodG14JyArIGV2ZW50TmFtZVxuICAgICAgICAgIH0gZWxzZSBpZiAoc3RhcnRzV2l0aChldmVudE5hbWUsICctJykpIHtcbiAgICAgICAgICAgIGV2ZW50TmFtZSA9ICdodG14OicgKyBldmVudE5hbWUuc2xpY2UoMSlcbiAgICAgICAgICB9IGVsc2UgaWYgKHN0YXJ0c1dpdGgoZXZlbnROYW1lLCAnaHRteC0nKSkge1xuICAgICAgICAgICAgZXZlbnROYW1lID0gJ2h0bXg6JyArIGV2ZW50TmFtZS5zbGljZSg1KVxuICAgICAgICAgIH1cblxuICAgICAgICAgIGFkZEh4T25FdmVudEhhbmRsZXIoZWx0LCBldmVudE5hbWUsIHZhbHVlKVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudHxIVE1MSW5wdXRFbGVtZW50fSBlbHRcbiAgICovXG4gIGZ1bmN0aW9uIGluaXROb2RlKGVsdCkge1xuICAgIGlmIChjbG9zZXN0KGVsdCwgaHRteC5jb25maWcuZGlzYWJsZVNlbGVjdG9yKSkge1xuICAgICAgY2xlYW5VcEVsZW1lbnQoZWx0KVxuICAgICAgcmV0dXJuXG4gICAgfVxuICAgIGNvbnN0IG5vZGVEYXRhID0gZ2V0SW50ZXJuYWxEYXRhKGVsdClcbiAgICBpZiAobm9kZURhdGEuaW5pdEhhc2ggIT09IGF0dHJpYnV0ZUhhc2goZWx0KSkge1xuICAgICAgLy8gY2xlYW4gdXAgYW55IHByZXZpb3VzbHkgcHJvY2Vzc2VkIGluZm9cbiAgICAgIGRlSW5pdE5vZGUoZWx0KVxuXG4gICAgICBub2RlRGF0YS5pbml0SGFzaCA9IGF0dHJpYnV0ZUhhc2goZWx0KVxuXG4gICAgICB0cmlnZ2VyRXZlbnQoZWx0LCAnaHRteDpiZWZvcmVQcm9jZXNzTm9kZScpXG5cbiAgICAgIC8vIEB0cy1pZ25vcmUgdmFsdWUgd2lsbCBiZSB1bmRlZmluZWQgZm9yIG5vbi1pbnB1dCBlbGVtZW50cywgd2hpY2ggaXMgZmluZVxuICAgICAgaWYgKGVsdC52YWx1ZSkge1xuICAgICAgICAvLyBAdHMtaWdub3JlXG4gICAgICAgIG5vZGVEYXRhLmxhc3RWYWx1ZSA9IGVsdC52YWx1ZVxuICAgICAgfVxuXG4gICAgICBjb25zdCB0cmlnZ2VyU3BlY3MgPSBnZXRUcmlnZ2VyU3BlY3MoZWx0KVxuICAgICAgY29uc3QgaGFzRXhwbGljaXRIdHRwQWN0aW9uID0gcHJvY2Vzc1ZlcmJzKGVsdCwgbm9kZURhdGEsIHRyaWdnZXJTcGVjcylcblxuICAgICAgaWYgKCFoYXNFeHBsaWNpdEh0dHBBY3Rpb24pIHtcbiAgICAgICAgaWYgKGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC1ib29zdCcpID09PSAndHJ1ZScpIHtcbiAgICAgICAgICBib29zdEVsZW1lbnQoZWx0LCBub2RlRGF0YSwgdHJpZ2dlclNwZWNzKVxuICAgICAgICB9IGVsc2UgaWYgKGhhc0F0dHJpYnV0ZShlbHQsICdoeC10cmlnZ2VyJykpIHtcbiAgICAgICAgICB0cmlnZ2VyU3BlY3MuZm9yRWFjaChmdW5jdGlvbih0cmlnZ2VyU3BlYykge1xuICAgICAgICAgICAgLy8gRm9yIFwibmFrZWRcIiB0cmlnZ2VycywgZG9uJ3QgZG8gYW55dGhpbmcgYXQgYWxsXG4gICAgICAgICAgICBhZGRUcmlnZ2VySGFuZGxlcihlbHQsIHRyaWdnZXJTcGVjLCBub2RlRGF0YSwgZnVuY3Rpb24oKSB7XG4gICAgICAgICAgICB9KVxuICAgICAgICAgIH0pXG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgLy8gSGFuZGxlIHN1Ym1pdCBidXR0b25zL2lucHV0cyB0aGF0IGhhdmUgdGhlIGZvcm0gYXR0cmlidXRlIHNldFxuICAgICAgLy8gc2VlIGh0dHBzOi8vZGV2ZWxvcGVyLm1vemlsbGEub3JnL2RvY3MvV2ViL0hUTUwvRWxlbWVudC9idXR0b25cbiAgICAgIGlmIChlbHQudGFnTmFtZSA9PT0gJ0ZPUk0nIHx8IChnZXRSYXdBdHRyaWJ1dGUoZWx0LCAndHlwZScpID09PSAnc3VibWl0JyAmJiBoYXNBdHRyaWJ1dGUoZWx0LCAnZm9ybScpKSkge1xuICAgICAgICBpbml0QnV0dG9uVHJhY2tpbmcoZWx0KVxuICAgICAgfVxuXG4gICAgICB0cmlnZ2VyRXZlbnQoZWx0LCAnaHRteDphZnRlclByb2Nlc3NOb2RlJylcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogUHJvY2Vzc2VzIG5ldyBjb250ZW50LCBlbmFibGluZyBodG14IGJlaGF2aW9yLiBUaGlzIGNhbiBiZSB1c2VmdWwgaWYgeW91IGhhdmUgY29udGVudCB0aGF0IGlzIGFkZGVkIHRvIHRoZSBET00gb3V0c2lkZSBvZiB0aGUgbm9ybWFsIGh0bXggcmVxdWVzdCBjeWNsZSBidXQgc3RpbGwgd2FudCBodG14IGF0dHJpYnV0ZXMgdG8gd29yay5cbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jcHJvY2Vzc1xuICAgKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR8c3RyaW5nfSBlbHQgZWxlbWVudCB0byBwcm9jZXNzXG4gICAqL1xuICBmdW5jdGlvbiBwcm9jZXNzTm9kZShlbHQpIHtcbiAgICBlbHQgPSByZXNvbHZlVGFyZ2V0KGVsdClcbiAgICBpZiAoY2xvc2VzdChlbHQsIGh0bXguY29uZmlnLmRpc2FibGVTZWxlY3RvcikpIHtcbiAgICAgIGNsZWFuVXBFbGVtZW50KGVsdClcbiAgICAgIHJldHVyblxuICAgIH1cbiAgICBpbml0Tm9kZShlbHQpXG4gICAgZm9yRWFjaChmaW5kRWxlbWVudHNUb1Byb2Nlc3MoZWx0KSwgZnVuY3Rpb24oY2hpbGQpIHsgaW5pdE5vZGUoY2hpbGQpIH0pXG4gICAgZm9yRWFjaChmaW5kSHhPbldpbGRjYXJkRWxlbWVudHMoZWx0KSwgcHJvY2Vzc0h4T25XaWxkY2FyZClcbiAgfVxuXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG4gIC8vIEV2ZW50L0xvZyBTdXBwb3J0XG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBzdHJcbiAgICogQHJldHVybnMge3N0cmluZ31cbiAgICovXG4gIGZ1bmN0aW9uIGtlYmFiRXZlbnROYW1lKHN0cikge1xuICAgIHJldHVybiBzdHIucmVwbGFjZSgvKFthLXowLTldKShbQS1aXSkvZywgJyQxLSQyJykudG9Mb3dlckNhc2UoKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBldmVudE5hbWVcbiAgICogQHBhcmFtIHthbnl9IGRldGFpbFxuICAgKiBAcmV0dXJucyB7Q3VzdG9tRXZlbnR9XG4gICAqL1xuICBmdW5jdGlvbiBtYWtlRXZlbnQoZXZlbnROYW1lLCBkZXRhaWwpIHtcbiAgICBsZXQgZXZ0XG4gICAgaWYgKHdpbmRvdy5DdXN0b21FdmVudCAmJiB0eXBlb2Ygd2luZG93LkN1c3RvbUV2ZW50ID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAvLyBUT0RPOiBgY29tcG9zZWQ6IHRydWVgIGhlcmUgaXMgYSBoYWNrIHRvIG1ha2UgZ2xvYmFsIGV2ZW50IGhhbmRsZXJzIHdvcmsgd2l0aCBldmVudHMgaW4gc2hhZG93IERPTVxuICAgICAgLy8gVGhpcyBicmVha3MgZXhwZWN0ZWQgZW5jYXBzdWxhdGlvbiBidXQgbmVlZHMgdG8gYmUgaGVyZSB1bnRpbCBkZWNpZGVkIG90aGVyd2lzZSBieSBjb3JlIGRldnNcbiAgICAgIGV2dCA9IG5ldyBDdXN0b21FdmVudChldmVudE5hbWUsIHsgYnViYmxlczogdHJ1ZSwgY2FuY2VsYWJsZTogdHJ1ZSwgY29tcG9zZWQ6IHRydWUsIGRldGFpbCB9KVxuICAgIH0gZWxzZSB7XG4gICAgICBldnQgPSBnZXREb2N1bWVudCgpLmNyZWF0ZUV2ZW50KCdDdXN0b21FdmVudCcpXG4gICAgICBldnQuaW5pdEN1c3RvbUV2ZW50KGV2ZW50TmFtZSwgdHJ1ZSwgdHJ1ZSwgZGV0YWlsKVxuICAgIH1cbiAgICByZXR1cm4gZXZ0XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFdmVudFRhcmdldHxzdHJpbmd9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ30gZXZlbnROYW1lXG4gICAqIEBwYXJhbSB7YW55PX0gZGV0YWlsXG4gICAqL1xuICBmdW5jdGlvbiB0cmlnZ2VyRXJyb3JFdmVudChlbHQsIGV2ZW50TmFtZSwgZGV0YWlsKSB7XG4gICAgdHJpZ2dlckV2ZW50KGVsdCwgZXZlbnROYW1lLCBtZXJnZU9iamVjdHMoeyBlcnJvcjogZXZlbnROYW1lIH0sIGRldGFpbCkpXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IGV2ZW50TmFtZVxuICAgKiBAcmV0dXJucyB7Ym9vbGVhbn1cbiAgICovXG4gIGZ1bmN0aW9uIGlnbm9yZUV2ZW50Rm9yTG9nZ2luZyhldmVudE5hbWUpIHtcbiAgICByZXR1cm4gZXZlbnROYW1lID09PSAnaHRteDphZnRlclByb2Nlc3NOb2RlJ1xuICB9XG5cbiAgLyoqXG4gICAqIGB3aXRoRXh0ZW5zaW9uc2AgbG9jYXRlcyBhbGwgYWN0aXZlIGV4dGVuc2lvbnMgZm9yIGEgcHJvdmlkZWQgZWxlbWVudCwgdGhlblxuICAgKiBleGVjdXRlcyB0aGUgcHJvdmlkZWQgZnVuY3Rpb24gdXNpbmcgZWFjaCBvZiB0aGUgYWN0aXZlIGV4dGVuc2lvbnMuICBJdCBzaG91bGRcbiAgICogYmUgY2FsbGVkIGludGVybmFsbHkgYXQgZXZlcnkgZXh0ZW5kYWJsZSBleGVjdXRpb24gcG9pbnQgaW4gaHRteC5cbiAgICpcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHBhcmFtIHsoZXh0ZW5zaW9uOkh0bXhFeHRlbnNpb24pID0+IHZvaWR9IHRvRG9cbiAgICogQHJldHVybnMgdm9pZFxuICAgKi9cbiAgZnVuY3Rpb24gd2l0aEV4dGVuc2lvbnMoZWx0LCB0b0RvKSB7XG4gICAgZm9yRWFjaChnZXRFeHRlbnNpb25zKGVsdCksIGZ1bmN0aW9uKGV4dGVuc2lvbikge1xuICAgICAgdHJ5IHtcbiAgICAgICAgdG9EbyhleHRlbnNpb24pXG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAgIGxvZ0Vycm9yKGUpXG4gICAgICB9XG4gICAgfSlcbiAgfVxuXG4gIGZ1bmN0aW9uIGxvZ0Vycm9yKG1zZykge1xuICAgIGlmIChjb25zb2xlLmVycm9yKSB7XG4gICAgICBjb25zb2xlLmVycm9yKG1zZylcbiAgICB9IGVsc2UgaWYgKGNvbnNvbGUubG9nKSB7XG4gICAgICBjb25zb2xlLmxvZygnRVJST1I6ICcsIG1zZylcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogVHJpZ2dlcnMgYSBnaXZlbiBldmVudCBvbiBhbiBlbGVtZW50XG4gICAqXG4gICAqIEBzZWUgaHR0cHM6Ly9odG14Lm9yZy9hcGkvI3RyaWdnZXJcbiAgICpcbiAgICogQHBhcmFtIHtFdmVudFRhcmdldHxzdHJpbmd9IGVsdCB0aGUgZWxlbWVudCB0byB0cmlnZ2VyIHRoZSBldmVudCBvblxuICAgKiBAcGFyYW0ge3N0cmluZ30gZXZlbnROYW1lIHRoZSBuYW1lIG9mIHRoZSBldmVudCB0byB0cmlnZ2VyXG4gICAqIEBwYXJhbSB7YW55PX0gZGV0YWlsIGRldGFpbHMgZm9yIHRoZSBldmVudFxuICAgKiBAcmV0dXJucyB7Ym9vbGVhbn1cbiAgICovXG4gIGZ1bmN0aW9uIHRyaWdnZXJFdmVudChlbHQsIGV2ZW50TmFtZSwgZGV0YWlsKSB7XG4gICAgZWx0ID0gcmVzb2x2ZVRhcmdldChlbHQpXG4gICAgaWYgKGRldGFpbCA9PSBudWxsKSB7XG4gICAgICBkZXRhaWwgPSB7fVxuICAgIH1cbiAgICBkZXRhaWwuZWx0ID0gZWx0XG4gICAgY29uc3QgZXZlbnQgPSBtYWtlRXZlbnQoZXZlbnROYW1lLCBkZXRhaWwpXG4gICAgaWYgKGh0bXgubG9nZ2VyICYmICFpZ25vcmVFdmVudEZvckxvZ2dpbmcoZXZlbnROYW1lKSkge1xuICAgICAgaHRteC5sb2dnZXIoZWx0LCBldmVudE5hbWUsIGRldGFpbClcbiAgICB9XG4gICAgaWYgKGRldGFpbC5lcnJvcikge1xuICAgICAgbG9nRXJyb3IoZGV0YWlsLmVycm9yKVxuICAgICAgdHJpZ2dlckV2ZW50KGVsdCwgJ2h0bXg6ZXJyb3InLCB7IGVycm9ySW5mbzogZGV0YWlsIH0pXG4gICAgfVxuICAgIGxldCBldmVudFJlc3VsdCA9IGVsdC5kaXNwYXRjaEV2ZW50KGV2ZW50KVxuICAgIGNvbnN0IGtlYmFiTmFtZSA9IGtlYmFiRXZlbnROYW1lKGV2ZW50TmFtZSlcbiAgICBpZiAoZXZlbnRSZXN1bHQgJiYga2ViYWJOYW1lICE9PSBldmVudE5hbWUpIHtcbiAgICAgIGNvbnN0IGtlYmFiZWRFdmVudCA9IG1ha2VFdmVudChrZWJhYk5hbWUsIGV2ZW50LmRldGFpbClcbiAgICAgIGV2ZW50UmVzdWx0ID0gZXZlbnRSZXN1bHQgJiYgZWx0LmRpc3BhdGNoRXZlbnQoa2ViYWJlZEV2ZW50KVxuICAgIH1cbiAgICB3aXRoRXh0ZW5zaW9ucyhhc0VsZW1lbnQoZWx0KSwgZnVuY3Rpb24oZXh0ZW5zaW9uKSB7XG4gICAgICBldmVudFJlc3VsdCA9IGV2ZW50UmVzdWx0ICYmIChleHRlbnNpb24ub25FdmVudChldmVudE5hbWUsIGV2ZW50KSAhPT0gZmFsc2UgJiYgIWV2ZW50LmRlZmF1bHRQcmV2ZW50ZWQpXG4gICAgfSlcbiAgICByZXR1cm4gZXZlbnRSZXN1bHRcbiAgfVxuXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG4gIC8vIEhpc3RvcnkgU3VwcG9ydFxuICAvLz0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PVxuICBsZXQgY3VycmVudFBhdGhGb3JIaXN0b3J5ID0gbG9jYXRpb24ucGF0aG5hbWUgKyBsb2NhdGlvbi5zZWFyY2hcblxuICAvKipcbiAgICogQHJldHVybnMge0VsZW1lbnR9XG4gICAqL1xuICBmdW5jdGlvbiBnZXRIaXN0b3J5RWxlbWVudCgpIHtcbiAgICBjb25zdCBoaXN0b3J5RWx0ID0gZ2V0RG9jdW1lbnQoKS5xdWVyeVNlbGVjdG9yKCdbaHgtaGlzdG9yeS1lbHRdLFtkYXRhLWh4LWhpc3RvcnktZWx0XScpXG4gICAgcmV0dXJuIGhpc3RvcnlFbHQgfHwgZ2V0RG9jdW1lbnQoKS5ib2R5XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IHVybFxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IHJvb3RFbHRcbiAgICovXG4gIGZ1bmN0aW9uIHNhdmVUb0hpc3RvcnlDYWNoZSh1cmwsIHJvb3RFbHQpIHtcbiAgICBpZiAoIWNhbkFjY2Vzc0xvY2FsU3RvcmFnZSgpKSB7XG4gICAgICByZXR1cm5cbiAgICB9XG5cbiAgICAvLyBnZXQgc3RhdGUgdG8gc2F2ZVxuICAgIGNvbnN0IGlubmVySFRNTCA9IGNsZWFuSW5uZXJIdG1sRm9ySGlzdG9yeShyb290RWx0KVxuICAgIGNvbnN0IHRpdGxlID0gZ2V0RG9jdW1lbnQoKS50aXRsZVxuICAgIGNvbnN0IHNjcm9sbCA9IHdpbmRvdy5zY3JvbGxZXG5cbiAgICBpZiAoaHRteC5jb25maWcuaGlzdG9yeUNhY2hlU2l6ZSA8PSAwKSB7XG4gICAgICAvLyBtYWtlIHN1cmUgdGhhdCBhbiBldmVudHVhbGx5IGFscmVhZHkgZXhpc3RpbmcgY2FjaGUgaXMgcHVyZ2VkXG4gICAgICBsb2NhbFN0b3JhZ2UucmVtb3ZlSXRlbSgnaHRteC1oaXN0b3J5LWNhY2hlJylcbiAgICAgIHJldHVyblxuICAgIH1cblxuICAgIHVybCA9IG5vcm1hbGl6ZVBhdGgodXJsKVxuXG4gICAgY29uc3QgaGlzdG9yeUNhY2hlID0gcGFyc2VKU09OKGxvY2FsU3RvcmFnZS5nZXRJdGVtKCdodG14LWhpc3RvcnktY2FjaGUnKSkgfHwgW11cbiAgICBmb3IgKGxldCBpID0gMDsgaSA8IGhpc3RvcnlDYWNoZS5sZW5ndGg7IGkrKykge1xuICAgICAgaWYgKGhpc3RvcnlDYWNoZVtpXS51cmwgPT09IHVybCkge1xuICAgICAgICBoaXN0b3J5Q2FjaGUuc3BsaWNlKGksIDEpXG4gICAgICAgIGJyZWFrXG4gICAgICB9XG4gICAgfVxuXG4gICAgLyoqIEB0eXBlIEh0bXhIaXN0b3J5SXRlbSAqL1xuICAgIGNvbnN0IG5ld0hpc3RvcnlJdGVtID0geyB1cmwsIGNvbnRlbnQ6IGlubmVySFRNTCwgdGl0bGUsIHNjcm9sbCB9XG5cbiAgICB0cmlnZ2VyRXZlbnQoZ2V0RG9jdW1lbnQoKS5ib2R5LCAnaHRteDpoaXN0b3J5SXRlbUNyZWF0ZWQnLCB7IGl0ZW06IG5ld0hpc3RvcnlJdGVtLCBjYWNoZTogaGlzdG9yeUNhY2hlIH0pXG5cbiAgICBoaXN0b3J5Q2FjaGUucHVzaChuZXdIaXN0b3J5SXRlbSlcbiAgICB3aGlsZSAoaGlzdG9yeUNhY2hlLmxlbmd0aCA+IGh0bXguY29uZmlnLmhpc3RvcnlDYWNoZVNpemUpIHtcbiAgICAgIGhpc3RvcnlDYWNoZS5zaGlmdCgpXG4gICAgfVxuXG4gICAgLy8ga2VlcCB0cnlpbmcgdG8gc2F2ZSB0aGUgY2FjaGUgdW50aWwgaXQgc3VjY2VlZHMgb3IgaXMgZW1wdHlcbiAgICB3aGlsZSAoaGlzdG9yeUNhY2hlLmxlbmd0aCA+IDApIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIGxvY2FsU3RvcmFnZS5zZXRJdGVtKCdodG14LWhpc3RvcnktY2FjaGUnLCBKU09OLnN0cmluZ2lmeShoaXN0b3J5Q2FjaGUpKVxuICAgICAgICBicmVha1xuICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICB0cmlnZ2VyRXJyb3JFdmVudChnZXREb2N1bWVudCgpLmJvZHksICdodG14Omhpc3RvcnlDYWNoZUVycm9yJywgeyBjYXVzZTogZSwgY2FjaGU6IGhpc3RvcnlDYWNoZSB9KVxuICAgICAgICBoaXN0b3J5Q2FjaGUuc2hpZnQoKSAvLyBzaHJpbmsgdGhlIGNhY2hlIGFuZCByZXRyeVxuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAdHlwZWRlZiB7T2JqZWN0fSBIdG14SGlzdG9yeUl0ZW1cbiAgICogQHByb3BlcnR5IHtzdHJpbmd9IHVybFxuICAgKiBAcHJvcGVydHkge3N0cmluZ30gY29udGVudFxuICAgKiBAcHJvcGVydHkge3N0cmluZ30gdGl0bGVcbiAgICogQHByb3BlcnR5IHtudW1iZXJ9IHNjcm9sbFxuICAgKi9cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IHVybFxuICAgKiBAcmV0dXJucyB7SHRteEhpc3RvcnlJdGVtfG51bGx9XG4gICAqL1xuICBmdW5jdGlvbiBnZXRDYWNoZWRIaXN0b3J5KHVybCkge1xuICAgIGlmICghY2FuQWNjZXNzTG9jYWxTdG9yYWdlKCkpIHtcbiAgICAgIHJldHVybiBudWxsXG4gICAgfVxuXG4gICAgdXJsID0gbm9ybWFsaXplUGF0aCh1cmwpXG5cbiAgICBjb25zdCBoaXN0b3J5Q2FjaGUgPSBwYXJzZUpTT04obG9jYWxTdG9yYWdlLmdldEl0ZW0oJ2h0bXgtaGlzdG9yeS1jYWNoZScpKSB8fCBbXVxuICAgIGZvciAobGV0IGkgPSAwOyBpIDwgaGlzdG9yeUNhY2hlLmxlbmd0aDsgaSsrKSB7XG4gICAgICBpZiAoaGlzdG9yeUNhY2hlW2ldLnVybCA9PT0gdXJsKSB7XG4gICAgICAgIHJldHVybiBoaXN0b3J5Q2FjaGVbaV1cbiAgICAgIH1cbiAgICB9XG4gICAgcmV0dXJuIG51bGxcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKiBAcmV0dXJucyB7c3RyaW5nfVxuICAgKi9cbiAgZnVuY3Rpb24gY2xlYW5Jbm5lckh0bWxGb3JIaXN0b3J5KGVsdCkge1xuICAgIGNvbnN0IGNsYXNzTmFtZSA9IGh0bXguY29uZmlnLnJlcXVlc3RDbGFzc1xuICAgIGNvbnN0IGNsb25lID0gLyoqIEB0eXBlIEVsZW1lbnQgKi8gKGVsdC5jbG9uZU5vZGUodHJ1ZSkpXG4gICAgZm9yRWFjaChmaW5kQWxsKGNsb25lLCAnLicgKyBjbGFzc05hbWUpLCBmdW5jdGlvbihjaGlsZCkge1xuICAgICAgcmVtb3ZlQ2xhc3NGcm9tRWxlbWVudChjaGlsZCwgY2xhc3NOYW1lKVxuICAgIH0pXG4gICAgLy8gcmVtb3ZlIHRoZSBkaXNhYmxlZCBhdHRyaWJ1dGUgZm9yIGFueSBlbGVtZW50IGRpc2FibGVkIGR1ZSB0byBhbiBodG14IHJlcXVlc3RcbiAgICBmb3JFYWNoKGZpbmRBbGwoY2xvbmUsICdbZGF0YS1kaXNhYmxlZC1ieS1odG14XScpLCBmdW5jdGlvbihjaGlsZCkge1xuICAgICAgY2hpbGQucmVtb3ZlQXR0cmlidXRlKCdkaXNhYmxlZCcpXG4gICAgfSlcbiAgICByZXR1cm4gY2xvbmUuaW5uZXJIVE1MXG4gIH1cblxuICBmdW5jdGlvbiBzYXZlQ3VycmVudFBhZ2VUb0hpc3RvcnkoKSB7XG4gICAgY29uc3QgZWx0ID0gZ2V0SGlzdG9yeUVsZW1lbnQoKVxuICAgIGNvbnN0IHBhdGggPSBjdXJyZW50UGF0aEZvckhpc3RvcnkgfHwgbG9jYXRpb24ucGF0aG5hbWUgKyBsb2NhdGlvbi5zZWFyY2hcblxuICAgIC8vIEFsbG93IGhpc3Rvcnkgc25hcHNob3QgZmVhdHVyZSB0byBiZSBkaXNhYmxlZCB3aGVyZSBoeC1oaXN0b3J5PVwiZmFsc2VcIlxuICAgIC8vIGlzIHByZXNlbnQgKmFueXdoZXJlKiBpbiB0aGUgY3VycmVudCBkb2N1bWVudCB3ZSdyZSBhYm91dCB0byBzYXZlLFxuICAgIC8vIHNvIHdlIGNhbiBwcmV2ZW50IHByaXZpbGVnZWQgZGF0YSBlbnRlcmluZyB0aGUgY2FjaGUuXG4gICAgLy8gVGhlIHBhZ2Ugd2lsbCBzdGlsbCBiZSByZWFjaGFibGUgYXMgYSBoaXN0b3J5IGVudHJ5LCBidXQgaHRteCB3aWxsIGZldGNoIGl0XG4gICAgLy8gbGl2ZSBmcm9tIHRoZSBzZXJ2ZXIgb25wb3BzdGF0ZSByYXRoZXIgdGhhbiBsb29rIGluIHRoZSBsb2NhbFN0b3JhZ2UgY2FjaGVcbiAgICBsZXQgZGlzYWJsZUhpc3RvcnlDYWNoZVxuICAgIHRyeSB7XG4gICAgICBkaXNhYmxlSGlzdG9yeUNhY2hlID0gZ2V0RG9jdW1lbnQoKS5xdWVyeVNlbGVjdG9yKCdbaHgtaGlzdG9yeT1cImZhbHNlXCIgaV0sW2RhdGEtaHgtaGlzdG9yeT1cImZhbHNlXCIgaV0nKVxuICAgIH0gY2F0Y2ggKGUpIHtcbiAgICAvLyBJRTExOiBpbnNlbnNpdGl2ZSBtb2RpZmllciBub3Qgc3VwcG9ydGVkIHNvIGZhbGxiYWNrIHRvIGNhc2Ugc2Vuc2l0aXZlIHNlbGVjdG9yXG4gICAgICBkaXNhYmxlSGlzdG9yeUNhY2hlID0gZ2V0RG9jdW1lbnQoKS5xdWVyeVNlbGVjdG9yKCdbaHgtaGlzdG9yeT1cImZhbHNlXCJdLFtkYXRhLWh4LWhpc3Rvcnk9XCJmYWxzZVwiXScpXG4gICAgfVxuICAgIGlmICghZGlzYWJsZUhpc3RvcnlDYWNoZSkge1xuICAgICAgdHJpZ2dlckV2ZW50KGdldERvY3VtZW50KCkuYm9keSwgJ2h0bXg6YmVmb3JlSGlzdG9yeVNhdmUnLCB7IHBhdGgsIGhpc3RvcnlFbHQ6IGVsdCB9KVxuICAgICAgc2F2ZVRvSGlzdG9yeUNhY2hlKHBhdGgsIGVsdClcbiAgICB9XG5cbiAgICBpZiAoaHRteC5jb25maWcuaGlzdG9yeUVuYWJsZWQpIGhpc3RvcnkucmVwbGFjZVN0YXRlKHsgaHRteDogdHJ1ZSB9LCBnZXREb2N1bWVudCgpLnRpdGxlLCB3aW5kb3cubG9jYXRpb24uaHJlZilcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge3N0cmluZ30gcGF0aFxuICAgKi9cbiAgZnVuY3Rpb24gcHVzaFVybEludG9IaXN0b3J5KHBhdGgpIHtcbiAgLy8gcmVtb3ZlIHRoZSBjYWNoZSBidXN0ZXIgcGFyYW1ldGVyLCBpZiBhbnlcbiAgICBpZiAoaHRteC5jb25maWcuZ2V0Q2FjaGVCdXN0ZXJQYXJhbSkge1xuICAgICAgcGF0aCA9IHBhdGgucmVwbGFjZSgvb3JnXFwuaHRteFxcLmNhY2hlLWJ1c3Rlcj1bXiZdKiY/LywgJycpXG4gICAgICBpZiAoZW5kc1dpdGgocGF0aCwgJyYnKSB8fCBlbmRzV2l0aChwYXRoLCAnPycpKSB7XG4gICAgICAgIHBhdGggPSBwYXRoLnNsaWNlKDAsIC0xKVxuICAgICAgfVxuICAgIH1cbiAgICBpZiAoaHRteC5jb25maWcuaGlzdG9yeUVuYWJsZWQpIHtcbiAgICAgIGhpc3RvcnkucHVzaFN0YXRlKHsgaHRteDogdHJ1ZSB9LCAnJywgcGF0aClcbiAgICB9XG4gICAgY3VycmVudFBhdGhGb3JIaXN0b3J5ID0gcGF0aFxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBwYXRoXG4gICAqL1xuICBmdW5jdGlvbiByZXBsYWNlVXJsSW5IaXN0b3J5KHBhdGgpIHtcbiAgICBpZiAoaHRteC5jb25maWcuaGlzdG9yeUVuYWJsZWQpIGhpc3RvcnkucmVwbGFjZVN0YXRlKHsgaHRteDogdHJ1ZSB9LCAnJywgcGF0aClcbiAgICBjdXJyZW50UGF0aEZvckhpc3RvcnkgPSBwYXRoXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtIdG14U2V0dGxlVGFza1tdfSB0YXNrc1xuICAgKi9cbiAgZnVuY3Rpb24gc2V0dGxlSW1tZWRpYXRlbHkodGFza3MpIHtcbiAgICBmb3JFYWNoKHRhc2tzLCBmdW5jdGlvbih0YXNrKSB7XG4gICAgICB0YXNrLmNhbGwodW5kZWZpbmVkKVxuICAgIH0pXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IHBhdGhcbiAgICovXG4gIGZ1bmN0aW9uIGxvYWRIaXN0b3J5RnJvbVNlcnZlcihwYXRoKSB7XG4gICAgY29uc3QgcmVxdWVzdCA9IG5ldyBYTUxIdHRwUmVxdWVzdCgpXG4gICAgY29uc3QgZGV0YWlscyA9IHsgcGF0aCwgeGhyOiByZXF1ZXN0IH1cbiAgICB0cmlnZ2VyRXZlbnQoZ2V0RG9jdW1lbnQoKS5ib2R5LCAnaHRteDpoaXN0b3J5Q2FjaGVNaXNzJywgZGV0YWlscylcbiAgICByZXF1ZXN0Lm9wZW4oJ0dFVCcsIHBhdGgsIHRydWUpXG4gICAgcmVxdWVzdC5zZXRSZXF1ZXN0SGVhZGVyKCdIWC1SZXF1ZXN0JywgJ3RydWUnKVxuICAgIHJlcXVlc3Quc2V0UmVxdWVzdEhlYWRlcignSFgtSGlzdG9yeS1SZXN0b3JlLVJlcXVlc3QnLCAndHJ1ZScpXG4gICAgcmVxdWVzdC5zZXRSZXF1ZXN0SGVhZGVyKCdIWC1DdXJyZW50LVVSTCcsIGdldERvY3VtZW50KCkubG9jYXRpb24uaHJlZilcbiAgICByZXF1ZXN0Lm9ubG9hZCA9IGZ1bmN0aW9uKCkge1xuICAgICAgaWYgKHRoaXMuc3RhdHVzID49IDIwMCAmJiB0aGlzLnN0YXR1cyA8IDQwMCkge1xuICAgICAgICB0cmlnZ2VyRXZlbnQoZ2V0RG9jdW1lbnQoKS5ib2R5LCAnaHRteDpoaXN0b3J5Q2FjaGVNaXNzTG9hZCcsIGRldGFpbHMpXG4gICAgICAgIGNvbnN0IGZyYWdtZW50ID0gbWFrZUZyYWdtZW50KHRoaXMucmVzcG9uc2UpXG4gICAgICAgIC8qKiBAdHlwZSBQYXJlbnROb2RlICovXG4gICAgICAgIGNvbnN0IGNvbnRlbnQgPSBmcmFnbWVudC5xdWVyeVNlbGVjdG9yKCdbaHgtaGlzdG9yeS1lbHRdLFtkYXRhLWh4LWhpc3RvcnktZWx0XScpIHx8IGZyYWdtZW50XG4gICAgICAgIGNvbnN0IGhpc3RvcnlFbGVtZW50ID0gZ2V0SGlzdG9yeUVsZW1lbnQoKVxuICAgICAgICBjb25zdCBzZXR0bGVJbmZvID0gbWFrZVNldHRsZUluZm8oaGlzdG9yeUVsZW1lbnQpXG4gICAgICAgIGhhbmRsZVRpdGxlKGZyYWdtZW50LnRpdGxlKVxuXG4gICAgICAgIHN3YXBJbm5lckhUTUwoaGlzdG9yeUVsZW1lbnQsIGNvbnRlbnQsIHNldHRsZUluZm8pXG4gICAgICAgIHNldHRsZUltbWVkaWF0ZWx5KHNldHRsZUluZm8udGFza3MpXG4gICAgICAgIGN1cnJlbnRQYXRoRm9ySGlzdG9yeSA9IHBhdGhcbiAgICAgICAgdHJpZ2dlckV2ZW50KGdldERvY3VtZW50KCkuYm9keSwgJ2h0bXg6aGlzdG9yeVJlc3RvcmUnLCB7IHBhdGgsIGNhY2hlTWlzczogdHJ1ZSwgc2VydmVyUmVzcG9uc2U6IHRoaXMucmVzcG9uc2UgfSlcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGdldERvY3VtZW50KCkuYm9keSwgJ2h0bXg6aGlzdG9yeUNhY2hlTWlzc0xvYWRFcnJvcicsIGRldGFpbHMpXG4gICAgICB9XG4gICAgfVxuICAgIHJlcXVlc3Quc2VuZCgpXG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtzdHJpbmd9IFtwYXRoXVxuICAgKi9cbiAgZnVuY3Rpb24gcmVzdG9yZUhpc3RvcnkocGF0aCkge1xuICAgIHNhdmVDdXJyZW50UGFnZVRvSGlzdG9yeSgpXG4gICAgcGF0aCA9IHBhdGggfHwgbG9jYXRpb24ucGF0aG5hbWUgKyBsb2NhdGlvbi5zZWFyY2hcbiAgICBjb25zdCBjYWNoZWQgPSBnZXRDYWNoZWRIaXN0b3J5KHBhdGgpXG4gICAgaWYgKGNhY2hlZCkge1xuICAgICAgY29uc3QgZnJhZ21lbnQgPSBtYWtlRnJhZ21lbnQoY2FjaGVkLmNvbnRlbnQpXG4gICAgICBjb25zdCBoaXN0b3J5RWxlbWVudCA9IGdldEhpc3RvcnlFbGVtZW50KClcbiAgICAgIGNvbnN0IHNldHRsZUluZm8gPSBtYWtlU2V0dGxlSW5mbyhoaXN0b3J5RWxlbWVudClcbiAgICAgIGhhbmRsZVRpdGxlKGZyYWdtZW50LnRpdGxlKVxuICAgICAgc3dhcElubmVySFRNTChoaXN0b3J5RWxlbWVudCwgZnJhZ21lbnQsIHNldHRsZUluZm8pXG4gICAgICBzZXR0bGVJbW1lZGlhdGVseShzZXR0bGVJbmZvLnRhc2tzKVxuICAgICAgZ2V0V2luZG93KCkuc2V0VGltZW91dChmdW5jdGlvbigpIHtcbiAgICAgICAgd2luZG93LnNjcm9sbFRvKDAsIGNhY2hlZC5zY3JvbGwpXG4gICAgICB9LCAwKSAvLyBuZXh0ICd0aWNrJywgc28gYnJvd3NlciBoYXMgdGltZSB0byByZW5kZXIgbGF5b3V0XG4gICAgICBjdXJyZW50UGF0aEZvckhpc3RvcnkgPSBwYXRoXG4gICAgICB0cmlnZ2VyRXZlbnQoZ2V0RG9jdW1lbnQoKS5ib2R5LCAnaHRteDpoaXN0b3J5UmVzdG9yZScsIHsgcGF0aCwgaXRlbTogY2FjaGVkIH0pXG4gICAgfSBlbHNlIHtcbiAgICAgIGlmIChodG14LmNvbmZpZy5yZWZyZXNoT25IaXN0b3J5TWlzcykge1xuICAgICAgICAvLyBAdHMtaWdub3JlOiBvcHRpb25hbCBwYXJhbWV0ZXIgaW4gcmVsb2FkKCkgZnVuY3Rpb24gdGhyb3dzIGVycm9yXG4gICAgICAgIC8vIG5vaW5zcGVjdGlvbiBKU1VucmVzb2x2ZWRSZWZlcmVuY2VcbiAgICAgICAgd2luZG93LmxvY2F0aW9uLnJlbG9hZCh0cnVlKVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgbG9hZEhpc3RvcnlGcm9tU2VydmVyKHBhdGgpXG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm5zIHtFbGVtZW50W119XG4gICAqL1xuICBmdW5jdGlvbiBhZGRSZXF1ZXN0SW5kaWNhdG9yQ2xhc3NlcyhlbHQpIHtcbiAgICBsZXQgaW5kaWNhdG9ycyA9IC8qKiBAdHlwZSBFbGVtZW50W10gKi8gKGZpbmRBdHRyaWJ1dGVUYXJnZXRzKGVsdCwgJ2h4LWluZGljYXRvcicpKVxuICAgIGlmIChpbmRpY2F0b3JzID09IG51bGwpIHtcbiAgICAgIGluZGljYXRvcnMgPSBbZWx0XVxuICAgIH1cbiAgICBmb3JFYWNoKGluZGljYXRvcnMsIGZ1bmN0aW9uKGljKSB7XG4gICAgICBjb25zdCBpbnRlcm5hbERhdGEgPSBnZXRJbnRlcm5hbERhdGEoaWMpXG4gICAgICBpbnRlcm5hbERhdGEucmVxdWVzdENvdW50ID0gKGludGVybmFsRGF0YS5yZXF1ZXN0Q291bnQgfHwgMCkgKyAxXG4gICAgICBpYy5jbGFzc0xpc3QuYWRkLmNhbGwoaWMuY2xhc3NMaXN0LCBodG14LmNvbmZpZy5yZXF1ZXN0Q2xhc3MpXG4gICAgfSlcbiAgICByZXR1cm4gaW5kaWNhdG9yc1xuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm5zIHtFbGVtZW50W119XG4gICAqL1xuICBmdW5jdGlvbiBkaXNhYmxlRWxlbWVudHMoZWx0KSB7XG4gICAgbGV0IGRpc2FibGVkRWx0cyA9IC8qKiBAdHlwZSBFbGVtZW50W10gKi8gKGZpbmRBdHRyaWJ1dGVUYXJnZXRzKGVsdCwgJ2h4LWRpc2FibGVkLWVsdCcpKVxuICAgIGlmIChkaXNhYmxlZEVsdHMgPT0gbnVsbCkge1xuICAgICAgZGlzYWJsZWRFbHRzID0gW11cbiAgICB9XG4gICAgZm9yRWFjaChkaXNhYmxlZEVsdHMsIGZ1bmN0aW9uKGRpc2FibGVkRWxlbWVudCkge1xuICAgICAgY29uc3QgaW50ZXJuYWxEYXRhID0gZ2V0SW50ZXJuYWxEYXRhKGRpc2FibGVkRWxlbWVudClcbiAgICAgIGludGVybmFsRGF0YS5yZXF1ZXN0Q291bnQgPSAoaW50ZXJuYWxEYXRhLnJlcXVlc3RDb3VudCB8fCAwKSArIDFcbiAgICAgIGRpc2FibGVkRWxlbWVudC5zZXRBdHRyaWJ1dGUoJ2Rpc2FibGVkJywgJycpXG4gICAgICBkaXNhYmxlZEVsZW1lbnQuc2V0QXR0cmlidXRlKCdkYXRhLWRpc2FibGVkLWJ5LWh0bXgnLCAnJylcbiAgICB9KVxuICAgIHJldHVybiBkaXNhYmxlZEVsdHNcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnRbXX0gaW5kaWNhdG9yc1xuICAgKiBAcGFyYW0ge0VsZW1lbnRbXX0gZGlzYWJsZWRcbiAgICovXG4gIGZ1bmN0aW9uIHJlbW92ZVJlcXVlc3RJbmRpY2F0b3JzKGluZGljYXRvcnMsIGRpc2FibGVkKSB7XG4gICAgZm9yRWFjaChpbmRpY2F0b3JzLCBmdW5jdGlvbihpYykge1xuICAgICAgY29uc3QgaW50ZXJuYWxEYXRhID0gZ2V0SW50ZXJuYWxEYXRhKGljKVxuICAgICAgaW50ZXJuYWxEYXRhLnJlcXVlc3RDb3VudCA9IChpbnRlcm5hbERhdGEucmVxdWVzdENvdW50IHx8IDApIC0gMVxuICAgICAgaWYgKGludGVybmFsRGF0YS5yZXF1ZXN0Q291bnQgPT09IDApIHtcbiAgICAgICAgaWMuY2xhc3NMaXN0LnJlbW92ZS5jYWxsKGljLmNsYXNzTGlzdCwgaHRteC5jb25maWcucmVxdWVzdENsYXNzKVxuICAgICAgfVxuICAgIH0pXG4gICAgZm9yRWFjaChkaXNhYmxlZCwgZnVuY3Rpb24oZGlzYWJsZWRFbGVtZW50KSB7XG4gICAgICBjb25zdCBpbnRlcm5hbERhdGEgPSBnZXRJbnRlcm5hbERhdGEoZGlzYWJsZWRFbGVtZW50KVxuICAgICAgaW50ZXJuYWxEYXRhLnJlcXVlc3RDb3VudCA9IChpbnRlcm5hbERhdGEucmVxdWVzdENvdW50IHx8IDApIC0gMVxuICAgICAgaWYgKGludGVybmFsRGF0YS5yZXF1ZXN0Q291bnQgPT09IDApIHtcbiAgICAgICAgZGlzYWJsZWRFbGVtZW50LnJlbW92ZUF0dHJpYnV0ZSgnZGlzYWJsZWQnKVxuICAgICAgICBkaXNhYmxlZEVsZW1lbnQucmVtb3ZlQXR0cmlidXRlKCdkYXRhLWRpc2FibGVkLWJ5LWh0bXgnKVxuICAgICAgfVxuICAgIH0pXG4gIH1cblxuICAvLz0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PVxuICAvLyBJbnB1dCBWYWx1ZSBQcm9jZXNzaW5nXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudFtdfSBwcm9jZXNzZWRcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHJldHVybnMge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiBoYXZlU2Vlbk5vZGUocHJvY2Vzc2VkLCBlbHQpIHtcbiAgICBmb3IgKGxldCBpID0gMDsgaSA8IHByb2Nlc3NlZC5sZW5ndGg7IGkrKykge1xuICAgICAgY29uc3Qgbm9kZSA9IHByb2Nlc3NlZFtpXVxuICAgICAgaWYgKG5vZGUuaXNTYW1lTm9kZShlbHQpKSB7XG4gICAgICAgIHJldHVybiB0cnVlXG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiBmYWxzZVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWxlbWVudFxuICAgKiBAcmV0dXJuIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gc2hvdWxkSW5jbHVkZShlbGVtZW50KSB7XG4gICAgLy8gQ2FzdCB0byB0cmljayB0c2MsIHVuZGVmaW5lZCB2YWx1ZXMgd2lsbCB3b3JrIGZpbmUgaGVyZVxuICAgIGNvbnN0IGVsdCA9IC8qKiBAdHlwZSB7SFRNTElucHV0RWxlbWVudH0gKi8gKGVsZW1lbnQpXG4gICAgaWYgKGVsdC5uYW1lID09PSAnJyB8fCBlbHQubmFtZSA9PSBudWxsIHx8IGVsdC5kaXNhYmxlZCB8fCBjbG9zZXN0KGVsdCwgJ2ZpZWxkc2V0W2Rpc2FibGVkXScpKSB7XG4gICAgICByZXR1cm4gZmFsc2VcbiAgICB9XG4gICAgLy8gaWdub3JlIFwic3VibWl0dGVyXCIgdHlwZXMgKHNlZSBqUXVlcnkgc3JjL3NlcmlhbGl6ZS5qcylcbiAgICBpZiAoZWx0LnR5cGUgPT09ICdidXR0b24nIHx8IGVsdC50eXBlID09PSAnc3VibWl0JyB8fCBlbHQudGFnTmFtZSA9PT0gJ2ltYWdlJyB8fCBlbHQudGFnTmFtZSA9PT0gJ3Jlc2V0JyB8fCBlbHQudGFnTmFtZSA9PT0gJ2ZpbGUnKSB7XG4gICAgICByZXR1cm4gZmFsc2VcbiAgICB9XG4gICAgaWYgKGVsdC50eXBlID09PSAnY2hlY2tib3gnIHx8IGVsdC50eXBlID09PSAncmFkaW8nKSB7XG4gICAgICByZXR1cm4gZWx0LmNoZWNrZWRcbiAgICB9XG4gICAgcmV0dXJuIHRydWVcbiAgfVxuXG4gIC8qKiBAcGFyYW0ge3N0cmluZ30gbmFtZVxuICAgKiBAcGFyYW0ge3N0cmluZ3xBcnJheXxGb3JtRGF0YUVudHJ5VmFsdWV9IHZhbHVlXG4gICAqIEBwYXJhbSB7Rm9ybURhdGF9IGZvcm1EYXRhICovXG4gIGZ1bmN0aW9uIGFkZFZhbHVlVG9Gb3JtRGF0YShuYW1lLCB2YWx1ZSwgZm9ybURhdGEpIHtcbiAgICBpZiAobmFtZSAhPSBudWxsICYmIHZhbHVlICE9IG51bGwpIHtcbiAgICAgIGlmIChBcnJheS5pc0FycmF5KHZhbHVlKSkge1xuICAgICAgICB2YWx1ZS5mb3JFYWNoKGZ1bmN0aW9uKHYpIHsgZm9ybURhdGEuYXBwZW5kKG5hbWUsIHYpIH0pXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBmb3JtRGF0YS5hcHBlbmQobmFtZSwgdmFsdWUpXG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqIEBwYXJhbSB7c3RyaW5nfSBuYW1lXG4gICAqIEBwYXJhbSB7c3RyaW5nfEFycmF5fSB2YWx1ZVxuICAgKiBAcGFyYW0ge0Zvcm1EYXRhfSBmb3JtRGF0YSAqL1xuICBmdW5jdGlvbiByZW1vdmVWYWx1ZUZyb21Gb3JtRGF0YShuYW1lLCB2YWx1ZSwgZm9ybURhdGEpIHtcbiAgICBpZiAobmFtZSAhPSBudWxsICYmIHZhbHVlICE9IG51bGwpIHtcbiAgICAgIGxldCB2YWx1ZXMgPSBmb3JtRGF0YS5nZXRBbGwobmFtZSlcbiAgICAgIGlmIChBcnJheS5pc0FycmF5KHZhbHVlKSkge1xuICAgICAgICB2YWx1ZXMgPSB2YWx1ZXMuZmlsdGVyKHYgPT4gdmFsdWUuaW5kZXhPZih2KSA8IDApXG4gICAgICB9IGVsc2Uge1xuICAgICAgICB2YWx1ZXMgPSB2YWx1ZXMuZmlsdGVyKHYgPT4gdiAhPT0gdmFsdWUpXG4gICAgICB9XG4gICAgICBmb3JtRGF0YS5kZWxldGUobmFtZSlcbiAgICAgIGZvckVhY2godmFsdWVzLCB2ID0+IGZvcm1EYXRhLmFwcGVuZChuYW1lLCB2KSlcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50W119IHByb2Nlc3NlZFxuICAgKiBAcGFyYW0ge0Zvcm1EYXRhfSBmb3JtRGF0YVxuICAgKiBAcGFyYW0ge0h0bXhFbGVtZW50VmFsaWRhdGlvbkVycm9yW119IGVycm9yc1xuICAgKiBAcGFyYW0ge0VsZW1lbnR8SFRNTElucHV0RWxlbWVudHxIVE1MU2VsZWN0RWxlbWVudHxIVE1MRm9ybUVsZW1lbnR9IGVsdFxuICAgKiBAcGFyYW0ge2Jvb2xlYW59IHZhbGlkYXRlXG4gICAqL1xuICBmdW5jdGlvbiBwcm9jZXNzSW5wdXRWYWx1ZShwcm9jZXNzZWQsIGZvcm1EYXRhLCBlcnJvcnMsIGVsdCwgdmFsaWRhdGUpIHtcbiAgICBpZiAoZWx0ID09IG51bGwgfHwgaGF2ZVNlZW5Ob2RlKHByb2Nlc3NlZCwgZWx0KSkge1xuICAgICAgcmV0dXJuXG4gICAgfSBlbHNlIHtcbiAgICAgIHByb2Nlc3NlZC5wdXNoKGVsdClcbiAgICB9XG4gICAgaWYgKHNob3VsZEluY2x1ZGUoZWx0KSkge1xuICAgICAgY29uc3QgbmFtZSA9IGdldFJhd0F0dHJpYnV0ZShlbHQsICduYW1lJylcbiAgICAgIC8vIEB0cy1pZ25vcmUgdmFsdWUgd2lsbCBiZSB1bmRlZmluZWQgZm9yIG5vbi1pbnB1dCBlbGVtZW50cywgd2hpY2ggaXMgZmluZVxuICAgICAgbGV0IHZhbHVlID0gZWx0LnZhbHVlXG4gICAgICBpZiAoZWx0IGluc3RhbmNlb2YgSFRNTFNlbGVjdEVsZW1lbnQgJiYgZWx0Lm11bHRpcGxlKSB7XG4gICAgICAgIHZhbHVlID0gdG9BcnJheShlbHQucXVlcnlTZWxlY3RvckFsbCgnb3B0aW9uOmNoZWNrZWQnKSkubWFwKGZ1bmN0aW9uKGUpIHsgcmV0dXJuICgvKiogQHR5cGUgSFRNTE9wdGlvbkVsZW1lbnQgKi8oZSkpLnZhbHVlIH0pXG4gICAgICB9XG4gICAgICAvLyBpbmNsdWRlIGZpbGUgaW5wdXRzXG4gICAgICBpZiAoZWx0IGluc3RhbmNlb2YgSFRNTElucHV0RWxlbWVudCAmJiBlbHQuZmlsZXMpIHtcbiAgICAgICAgdmFsdWUgPSB0b0FycmF5KGVsdC5maWxlcylcbiAgICAgIH1cbiAgICAgIGFkZFZhbHVlVG9Gb3JtRGF0YShuYW1lLCB2YWx1ZSwgZm9ybURhdGEpXG4gICAgICBpZiAodmFsaWRhdGUpIHtcbiAgICAgICAgdmFsaWRhdGVFbGVtZW50KGVsdCwgZXJyb3JzKVxuICAgICAgfVxuICAgIH1cbiAgICBpZiAoZWx0IGluc3RhbmNlb2YgSFRNTEZvcm1FbGVtZW50KSB7XG4gICAgICBmb3JFYWNoKGVsdC5lbGVtZW50cywgZnVuY3Rpb24oaW5wdXQpIHtcbiAgICAgICAgaWYgKHByb2Nlc3NlZC5pbmRleE9mKGlucHV0KSA+PSAwKSB7XG4gICAgICAgICAgLy8gVGhlIGlucHV0IGhhcyBhbHJlYWR5IGJlZW4gcHJvY2Vzc2VkIGFuZCBhZGRlZCB0byB0aGUgdmFsdWVzLCBidXQgdGhlIEZvcm1EYXRhIHRoYXQgd2lsbCBiZVxuICAgICAgICAgIC8vICBjb25zdHJ1Y3RlZCByaWdodCBhZnRlciBvbiB0aGUgZm9ybSwgd2lsbCBpbmNsdWRlIGl0IG9uY2UgYWdhaW4uIFNvIHJlbW92ZSB0aGF0IGlucHV0J3MgdmFsdWVcbiAgICAgICAgICAvLyAgbm93IHRvIGF2b2lkIGR1cGxpY2F0ZXNcbiAgICAgICAgICByZW1vdmVWYWx1ZUZyb21Gb3JtRGF0YShpbnB1dC5uYW1lLCBpbnB1dC52YWx1ZSwgZm9ybURhdGEpXG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgcHJvY2Vzc2VkLnB1c2goaW5wdXQpXG4gICAgICAgIH1cbiAgICAgICAgaWYgKHZhbGlkYXRlKSB7XG4gICAgICAgICAgdmFsaWRhdGVFbGVtZW50KGlucHV0LCBlcnJvcnMpXG4gICAgICAgIH1cbiAgICAgIH0pXG4gICAgICBuZXcgRm9ybURhdGEoZWx0KS5mb3JFYWNoKGZ1bmN0aW9uKHZhbHVlLCBuYW1lKSB7XG4gICAgICAgIGlmICh2YWx1ZSBpbnN0YW5jZW9mIEZpbGUgJiYgdmFsdWUubmFtZSA9PT0gJycpIHtcbiAgICAgICAgICByZXR1cm4gLy8gaWdub3JlIG5vLW5hbWUgZmlsZXNcbiAgICAgICAgfVxuICAgICAgICBhZGRWYWx1ZVRvRm9ybURhdGEobmFtZSwgdmFsdWUsIGZvcm1EYXRhKVxuICAgICAgfSlcbiAgICB9XG4gIH1cblxuICAvKipcbiAgICpcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHBhcmFtIHtIdG14RWxlbWVudFZhbGlkYXRpb25FcnJvcltdfSBlcnJvcnNcbiAgICovXG4gIGZ1bmN0aW9uIHZhbGlkYXRlRWxlbWVudChlbHQsIGVycm9ycykge1xuICAgIGNvbnN0IGVsZW1lbnQgPSAvKiogQHR5cGUge0hUTUxFbGVtZW50ICYgRWxlbWVudEludGVybmFsc30gKi8gKGVsdClcbiAgICBpZiAoZWxlbWVudC53aWxsVmFsaWRhdGUpIHtcbiAgICAgIHRyaWdnZXJFdmVudChlbGVtZW50LCAnaHRteDp2YWxpZGF0aW9uOnZhbGlkYXRlJylcbiAgICAgIGlmICghZWxlbWVudC5jaGVja1ZhbGlkaXR5KCkpIHtcbiAgICAgICAgZXJyb3JzLnB1c2goeyBlbHQ6IGVsZW1lbnQsIG1lc3NhZ2U6IGVsZW1lbnQudmFsaWRhdGlvbk1lc3NhZ2UsIHZhbGlkaXR5OiBlbGVtZW50LnZhbGlkaXR5IH0pXG4gICAgICAgIHRyaWdnZXJFdmVudChlbGVtZW50LCAnaHRteDp2YWxpZGF0aW9uOmZhaWxlZCcsIHsgbWVzc2FnZTogZWxlbWVudC52YWxpZGF0aW9uTWVzc2FnZSwgdmFsaWRpdHk6IGVsZW1lbnQudmFsaWRpdHkgfSlcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogT3ZlcnJpZGUgdmFsdWVzIGluIHRoZSBvbmUgRm9ybURhdGEgd2l0aCB0aG9zZSBmcm9tIGFub3RoZXIuXG4gICAqIEBwYXJhbSB7Rm9ybURhdGF9IHJlY2VpdmVyIHRoZSBmb3JtZGF0YSB0aGF0IHdpbGwgYmUgbXV0YXRlZFxuICAgKiBAcGFyYW0ge0Zvcm1EYXRhfSBkb25vciB0aGUgZm9ybWRhdGEgdGhhdCB3aWxsIHByb3ZpZGUgdGhlIG92ZXJyaWRpbmcgdmFsdWVzXG4gICAqIEByZXR1cm5zIHtGb3JtRGF0YX0gdGhlIHtAbGlua2NvZGUgcmVjZWl2ZXJ9XG4gICAqL1xuICBmdW5jdGlvbiBvdmVycmlkZUZvcm1EYXRhKHJlY2VpdmVyLCBkb25vcikge1xuICAgIGZvciAoY29uc3Qga2V5IG9mIGRvbm9yLmtleXMoKSkge1xuICAgICAgcmVjZWl2ZXIuZGVsZXRlKGtleSlcbiAgICB9XG4gICAgZG9ub3IuZm9yRWFjaChmdW5jdGlvbih2YWx1ZSwga2V5KSB7XG4gICAgICByZWNlaXZlci5hcHBlbmQoa2V5LCB2YWx1ZSlcbiAgICB9KVxuICAgIHJldHVybiByZWNlaXZlclxuICB9XG5cbiAgLyoqXG4gKiBAcGFyYW0ge0VsZW1lbnR8SFRNTEZvcm1FbGVtZW50fSBlbHRcbiAqIEBwYXJhbSB7SHR0cFZlcmJ9IHZlcmJcbiAqIEByZXR1cm5zIHt7ZXJyb3JzOiBIdG14RWxlbWVudFZhbGlkYXRpb25FcnJvcltdLCBmb3JtRGF0YTogRm9ybURhdGEsIHZhbHVlczogT2JqZWN0fX1cbiAqL1xuICBmdW5jdGlvbiBnZXRJbnB1dFZhbHVlcyhlbHQsIHZlcmIpIHtcbiAgICAvKiogQHR5cGUgRWxlbWVudFtdICovXG4gICAgY29uc3QgcHJvY2Vzc2VkID0gW11cbiAgICBjb25zdCBmb3JtRGF0YSA9IG5ldyBGb3JtRGF0YSgpXG4gICAgY29uc3QgcHJpb3JpdHlGb3JtRGF0YSA9IG5ldyBGb3JtRGF0YSgpXG4gICAgLyoqIEB0eXBlIEh0bXhFbGVtZW50VmFsaWRhdGlvbkVycm9yW10gKi9cbiAgICBjb25zdCBlcnJvcnMgPSBbXVxuICAgIGNvbnN0IGludGVybmFsRGF0YSA9IGdldEludGVybmFsRGF0YShlbHQpXG4gICAgaWYgKGludGVybmFsRGF0YS5sYXN0QnV0dG9uQ2xpY2tlZCAmJiAhYm9keUNvbnRhaW5zKGludGVybmFsRGF0YS5sYXN0QnV0dG9uQ2xpY2tlZCkpIHtcbiAgICAgIGludGVybmFsRGF0YS5sYXN0QnV0dG9uQ2xpY2tlZCA9IG51bGxcbiAgICB9XG5cbiAgICAvLyBvbmx5IHZhbGlkYXRlIHdoZW4gZm9ybSBpcyBkaXJlY3RseSBzdWJtaXR0ZWQgYW5kIG5vdmFsaWRhdGUgb3IgZm9ybW5vdmFsaWRhdGUgYXJlIG5vdCBzZXRcbiAgICAvLyBvciBpZiB0aGUgZWxlbWVudCBoYXMgYW4gZXhwbGljaXQgaHgtdmFsaWRhdGU9XCJ0cnVlXCIgb24gaXRcbiAgICBsZXQgdmFsaWRhdGUgPSAoZWx0IGluc3RhbmNlb2YgSFRNTEZvcm1FbGVtZW50ICYmIGVsdC5ub1ZhbGlkYXRlICE9PSB0cnVlKSB8fCBnZXRBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC12YWxpZGF0ZScpID09PSAndHJ1ZSdcbiAgICBpZiAoaW50ZXJuYWxEYXRhLmxhc3RCdXR0b25DbGlja2VkKSB7XG4gICAgICB2YWxpZGF0ZSA9IHZhbGlkYXRlICYmIGludGVybmFsRGF0YS5sYXN0QnV0dG9uQ2xpY2tlZC5mb3JtTm9WYWxpZGF0ZSAhPT0gdHJ1ZVxuICAgIH1cblxuICAgIC8vIGZvciBhIG5vbi1HRVQgaW5jbHVkZSB0aGUgY2xvc2VzdCBmb3JtXG4gICAgaWYgKHZlcmIgIT09ICdnZXQnKSB7XG4gICAgICBwcm9jZXNzSW5wdXRWYWx1ZShwcm9jZXNzZWQsIHByaW9yaXR5Rm9ybURhdGEsIGVycm9ycywgY2xvc2VzdChlbHQsICdmb3JtJyksIHZhbGlkYXRlKVxuICAgIH1cblxuICAgIC8vIGluY2x1ZGUgdGhlIGVsZW1lbnQgaXRzZWxmXG4gICAgcHJvY2Vzc0lucHV0VmFsdWUocHJvY2Vzc2VkLCBmb3JtRGF0YSwgZXJyb3JzLCBlbHQsIHZhbGlkYXRlKVxuXG4gICAgLy8gaWYgYSBidXR0b24gb3Igc3VibWl0IHdhcyBjbGlja2VkIGxhc3QsIGluY2x1ZGUgaXRzIHZhbHVlXG4gICAgaWYgKGludGVybmFsRGF0YS5sYXN0QnV0dG9uQ2xpY2tlZCB8fCBlbHQudGFnTmFtZSA9PT0gJ0JVVFRPTicgfHxcbiAgICAoZWx0LnRhZ05hbWUgPT09ICdJTlBVVCcgJiYgZ2V0UmF3QXR0cmlidXRlKGVsdCwgJ3R5cGUnKSA9PT0gJ3N1Ym1pdCcpKSB7XG4gICAgICBjb25zdCBidXR0b24gPSBpbnRlcm5hbERhdGEubGFzdEJ1dHRvbkNsaWNrZWQgfHwgKC8qKiBAdHlwZSBIVE1MSW5wdXRFbGVtZW50fEhUTUxCdXR0b25FbGVtZW50ICovKGVsdCkpXG4gICAgICBjb25zdCBuYW1lID0gZ2V0UmF3QXR0cmlidXRlKGJ1dHRvbiwgJ25hbWUnKVxuICAgICAgYWRkVmFsdWVUb0Zvcm1EYXRhKG5hbWUsIGJ1dHRvbi52YWx1ZSwgcHJpb3JpdHlGb3JtRGF0YSlcbiAgICB9XG5cbiAgICAvLyBpbmNsdWRlIGFueSBleHBsaWNpdCBpbmNsdWRlc1xuICAgIGNvbnN0IGluY2x1ZGVzID0gZmluZEF0dHJpYnV0ZVRhcmdldHMoZWx0LCAnaHgtaW5jbHVkZScpXG4gICAgZm9yRWFjaChpbmNsdWRlcywgZnVuY3Rpb24obm9kZSkge1xuICAgICAgcHJvY2Vzc0lucHV0VmFsdWUocHJvY2Vzc2VkLCBmb3JtRGF0YSwgZXJyb3JzLCBhc0VsZW1lbnQobm9kZSksIHZhbGlkYXRlKVxuICAgICAgLy8gaWYgYSBub24tZm9ybSBpcyBpbmNsdWRlZCwgaW5jbHVkZSBhbnkgaW5wdXQgdmFsdWVzIHdpdGhpbiBpdFxuICAgICAgaWYgKCFtYXRjaGVzKG5vZGUsICdmb3JtJykpIHtcbiAgICAgICAgZm9yRWFjaChhc1BhcmVudE5vZGUobm9kZSkucXVlcnlTZWxlY3RvckFsbChJTlBVVF9TRUxFQ1RPUiksIGZ1bmN0aW9uKGRlc2NlbmRhbnQpIHtcbiAgICAgICAgICBwcm9jZXNzSW5wdXRWYWx1ZShwcm9jZXNzZWQsIGZvcm1EYXRhLCBlcnJvcnMsIGRlc2NlbmRhbnQsIHZhbGlkYXRlKVxuICAgICAgICB9KVxuICAgICAgfVxuICAgIH0pXG5cbiAgICAvLyB2YWx1ZXMgZnJvbSBhIDxmb3JtPiB0YWtlIHByZWNlZGVuY2UsIG92ZXJyaWRpbmcgdGhlIHJlZ3VsYXIgdmFsdWVzXG4gICAgb3ZlcnJpZGVGb3JtRGF0YShmb3JtRGF0YSwgcHJpb3JpdHlGb3JtRGF0YSlcblxuICAgIHJldHVybiB7IGVycm9ycywgZm9ybURhdGEsIHZhbHVlczogZm9ybURhdGFQcm94eShmb3JtRGF0YSkgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSByZXR1cm5TdHJcbiAgICogQHBhcmFtIHtzdHJpbmd9IG5hbWVcbiAgICogQHBhcmFtIHthbnl9IHJlYWxWYWx1ZVxuICAgKiBAcmV0dXJucyB7c3RyaW5nfVxuICAgKi9cbiAgZnVuY3Rpb24gYXBwZW5kUGFyYW0ocmV0dXJuU3RyLCBuYW1lLCByZWFsVmFsdWUpIHtcbiAgICBpZiAocmV0dXJuU3RyICE9PSAnJykge1xuICAgICAgcmV0dXJuU3RyICs9ICcmJ1xuICAgIH1cbiAgICBpZiAoU3RyaW5nKHJlYWxWYWx1ZSkgPT09ICdbb2JqZWN0IE9iamVjdF0nKSB7XG4gICAgICByZWFsVmFsdWUgPSBKU09OLnN0cmluZ2lmeShyZWFsVmFsdWUpXG4gICAgfVxuICAgIGNvbnN0IHMgPSBlbmNvZGVVUklDb21wb25lbnQocmVhbFZhbHVlKVxuICAgIHJldHVyblN0ciArPSBlbmNvZGVVUklDb21wb25lbnQobmFtZSkgKyAnPScgKyBzXG4gICAgcmV0dXJuIHJldHVyblN0clxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Rm9ybURhdGF8T2JqZWN0fSB2YWx1ZXNcbiAgICogQHJldHVybnMgc3RyaW5nXG4gICAqL1xuICBmdW5jdGlvbiB1cmxFbmNvZGUodmFsdWVzKSB7XG4gICAgdmFsdWVzID0gZm9ybURhdGFGcm9tT2JqZWN0KHZhbHVlcylcbiAgICBsZXQgcmV0dXJuU3RyID0gJydcbiAgICB2YWx1ZXMuZm9yRWFjaChmdW5jdGlvbih2YWx1ZSwga2V5KSB7XG4gICAgICByZXR1cm5TdHIgPSBhcHBlbmRQYXJhbShyZXR1cm5TdHIsIGtleSwgdmFsdWUpXG4gICAgfSlcbiAgICByZXR1cm4gcmV0dXJuU3RyXG4gIH1cblxuICAvLz0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PVxuICAvLyBBamF4XG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG5cbiAgLyoqXG4gKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICogQHBhcmFtIHtFbGVtZW50fSB0YXJnZXRcbiAqIEBwYXJhbSB7c3RyaW5nfSBwcm9tcHRcbiAqIEByZXR1cm5zIHtIdG14SGVhZGVyU3BlY2lmaWNhdGlvbn1cbiAqL1xuICBmdW5jdGlvbiBnZXRIZWFkZXJzKGVsdCwgdGFyZ2V0LCBwcm9tcHQpIHtcbiAgICAvKiogQHR5cGUgSHRteEhlYWRlclNwZWNpZmljYXRpb24gKi9cbiAgICBjb25zdCBoZWFkZXJzID0ge1xuICAgICAgJ0hYLVJlcXVlc3QnOiAndHJ1ZScsXG4gICAgICAnSFgtVHJpZ2dlcic6IGdldFJhd0F0dHJpYnV0ZShlbHQsICdpZCcpLFxuICAgICAgJ0hYLVRyaWdnZXItTmFtZSc6IGdldFJhd0F0dHJpYnV0ZShlbHQsICduYW1lJyksXG4gICAgICAnSFgtVGFyZ2V0JzogZ2V0QXR0cmlidXRlVmFsdWUodGFyZ2V0LCAnaWQnKSxcbiAgICAgICdIWC1DdXJyZW50LVVSTCc6IGdldERvY3VtZW50KCkubG9jYXRpb24uaHJlZlxuICAgIH1cbiAgICBnZXRWYWx1ZXNGb3JFbGVtZW50KGVsdCwgJ2h4LWhlYWRlcnMnLCBmYWxzZSwgaGVhZGVycylcbiAgICBpZiAocHJvbXB0ICE9PSB1bmRlZmluZWQpIHtcbiAgICAgIGhlYWRlcnNbJ0hYLVByb21wdCddID0gcHJvbXB0XG4gICAgfVxuICAgIGlmIChnZXRJbnRlcm5hbERhdGEoZWx0KS5ib29zdGVkKSB7XG4gICAgICBoZWFkZXJzWydIWC1Cb29zdGVkJ10gPSAndHJ1ZSdcbiAgICB9XG4gICAgcmV0dXJuIGhlYWRlcnNcbiAgfVxuXG4gIC8qKlxuICogZmlsdGVyVmFsdWVzIHRha2VzIGFuIG9iamVjdCBjb250YWluaW5nIGZvcm0gaW5wdXQgdmFsdWVzXG4gKiBhbmQgcmV0dXJucyBhIG5ldyBvYmplY3QgdGhhdCBvbmx5IGNvbnRhaW5zIGtleXMgdGhhdCBhcmVcbiAqIHNwZWNpZmllZCBieSB0aGUgY2xvc2VzdCBcImh4LXBhcmFtc1wiIGF0dHJpYnV0ZVxuICogQHBhcmFtIHtGb3JtRGF0YX0gaW5wdXRWYWx1ZXNcbiAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gKiBAcmV0dXJucyB7Rm9ybURhdGF9XG4gKi9cbiAgZnVuY3Rpb24gZmlsdGVyVmFsdWVzKGlucHV0VmFsdWVzLCBlbHQpIHtcbiAgICBjb25zdCBwYXJhbXNWYWx1ZSA9IGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC1wYXJhbXMnKVxuICAgIGlmIChwYXJhbXNWYWx1ZSkge1xuICAgICAgaWYgKHBhcmFtc1ZhbHVlID09PSAnbm9uZScpIHtcbiAgICAgICAgcmV0dXJuIG5ldyBGb3JtRGF0YSgpXG4gICAgICB9IGVsc2UgaWYgKHBhcmFtc1ZhbHVlID09PSAnKicpIHtcbiAgICAgICAgcmV0dXJuIGlucHV0VmFsdWVzXG4gICAgICB9IGVsc2UgaWYgKHBhcmFtc1ZhbHVlLmluZGV4T2YoJ25vdCAnKSA9PT0gMCkge1xuICAgICAgICBmb3JFYWNoKHBhcmFtc1ZhbHVlLnN1YnN0cig0KS5zcGxpdCgnLCcpLCBmdW5jdGlvbihuYW1lKSB7XG4gICAgICAgICAgbmFtZSA9IG5hbWUudHJpbSgpXG4gICAgICAgICAgaW5wdXRWYWx1ZXMuZGVsZXRlKG5hbWUpXG4gICAgICAgIH0pXG4gICAgICAgIHJldHVybiBpbnB1dFZhbHVlc1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgY29uc3QgbmV3VmFsdWVzID0gbmV3IEZvcm1EYXRhKClcbiAgICAgICAgZm9yRWFjaChwYXJhbXNWYWx1ZS5zcGxpdCgnLCcpLCBmdW5jdGlvbihuYW1lKSB7XG4gICAgICAgICAgbmFtZSA9IG5hbWUudHJpbSgpXG4gICAgICAgICAgaWYgKGlucHV0VmFsdWVzLmhhcyhuYW1lKSkge1xuICAgICAgICAgICAgaW5wdXRWYWx1ZXMuZ2V0QWxsKG5hbWUpLmZvckVhY2goZnVuY3Rpb24odmFsdWUpIHsgbmV3VmFsdWVzLmFwcGVuZChuYW1lLCB2YWx1ZSkgfSlcbiAgICAgICAgICB9XG4gICAgICAgIH0pXG4gICAgICAgIHJldHVybiBuZXdWYWx1ZXNcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuIGlucHV0VmFsdWVzXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm4ge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiBpc0FuY2hvckxpbmsoZWx0KSB7XG4gICAgcmV0dXJuICEhZ2V0UmF3QXR0cmlidXRlKGVsdCwgJ2hyZWYnKSAmJiBnZXRSYXdBdHRyaWJ1dGUoZWx0LCAnaHJlZicpLmluZGV4T2YoJyMnKSA+PSAwXG4gIH1cblxuICAvKipcbiAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gKiBAcGFyYW0ge0h0bXhTd2FwU3R5bGV9IFtzd2FwSW5mb092ZXJyaWRlXVxuICogQHJldHVybnMge0h0bXhTd2FwU3BlY2lmaWNhdGlvbn1cbiAqL1xuICBmdW5jdGlvbiBnZXRTd2FwU3BlY2lmaWNhdGlvbihlbHQsIHN3YXBJbmZvT3ZlcnJpZGUpIHtcbiAgICBjb25zdCBzd2FwSW5mbyA9IHN3YXBJbmZvT3ZlcnJpZGUgfHwgZ2V0Q2xvc2VzdEF0dHJpYnV0ZVZhbHVlKGVsdCwgJ2h4LXN3YXAnKVxuICAgIC8qKiBAdHlwZSBIdG14U3dhcFNwZWNpZmljYXRpb24gKi9cbiAgICBjb25zdCBzd2FwU3BlYyA9IHtcbiAgICAgIHN3YXBTdHlsZTogZ2V0SW50ZXJuYWxEYXRhKGVsdCkuYm9vc3RlZCA/ICdpbm5lckhUTUwnIDogaHRteC5jb25maWcuZGVmYXVsdFN3YXBTdHlsZSxcbiAgICAgIHN3YXBEZWxheTogaHRteC5jb25maWcuZGVmYXVsdFN3YXBEZWxheSxcbiAgICAgIHNldHRsZURlbGF5OiBodG14LmNvbmZpZy5kZWZhdWx0U2V0dGxlRGVsYXlcbiAgICB9XG4gICAgaWYgKGh0bXguY29uZmlnLnNjcm9sbEludG9WaWV3T25Cb29zdCAmJiBnZXRJbnRlcm5hbERhdGEoZWx0KS5ib29zdGVkICYmICFpc0FuY2hvckxpbmsoZWx0KSkge1xuICAgICAgc3dhcFNwZWMuc2hvdyA9ICd0b3AnXG4gICAgfVxuICAgIGlmIChzd2FwSW5mbykge1xuICAgICAgY29uc3Qgc3BsaXQgPSBzcGxpdE9uV2hpdGVzcGFjZShzd2FwSW5mbylcbiAgICAgIGlmIChzcGxpdC5sZW5ndGggPiAwKSB7XG4gICAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgc3BsaXQubGVuZ3RoOyBpKyspIHtcbiAgICAgICAgICBjb25zdCB2YWx1ZSA9IHNwbGl0W2ldXG4gICAgICAgICAgaWYgKHZhbHVlLmluZGV4T2YoJ3N3YXA6JykgPT09IDApIHtcbiAgICAgICAgICAgIHN3YXBTcGVjLnN3YXBEZWxheSA9IHBhcnNlSW50ZXJ2YWwodmFsdWUuc3Vic3RyKDUpKVxuICAgICAgICAgIH0gZWxzZSBpZiAodmFsdWUuaW5kZXhPZignc2V0dGxlOicpID09PSAwKSB7XG4gICAgICAgICAgICBzd2FwU3BlYy5zZXR0bGVEZWxheSA9IHBhcnNlSW50ZXJ2YWwodmFsdWUuc3Vic3RyKDcpKVxuICAgICAgICAgIH0gZWxzZSBpZiAodmFsdWUuaW5kZXhPZigndHJhbnNpdGlvbjonKSA9PT0gMCkge1xuICAgICAgICAgICAgc3dhcFNwZWMudHJhbnNpdGlvbiA9IHZhbHVlLnN1YnN0cigxMSkgPT09ICd0cnVlJ1xuICAgICAgICAgIH0gZWxzZSBpZiAodmFsdWUuaW5kZXhPZignaWdub3JlVGl0bGU6JykgPT09IDApIHtcbiAgICAgICAgICAgIHN3YXBTcGVjLmlnbm9yZVRpdGxlID0gdmFsdWUuc3Vic3RyKDEyKSA9PT0gJ3RydWUnXG4gICAgICAgICAgfSBlbHNlIGlmICh2YWx1ZS5pbmRleE9mKCdzY3JvbGw6JykgPT09IDApIHtcbiAgICAgICAgICAgIGNvbnN0IHNjcm9sbFNwZWMgPSB2YWx1ZS5zdWJzdHIoNylcbiAgICAgICAgICAgIHZhciBzcGxpdFNwZWMgPSBzY3JvbGxTcGVjLnNwbGl0KCc6JylcbiAgICAgICAgICAgIGNvbnN0IHNjcm9sbFZhbCA9IHNwbGl0U3BlYy5wb3AoKVxuICAgICAgICAgICAgdmFyIHNlbGVjdG9yVmFsID0gc3BsaXRTcGVjLmxlbmd0aCA+IDAgPyBzcGxpdFNwZWMuam9pbignOicpIDogbnVsbFxuICAgICAgICAgICAgLy8gQHRzLWlnbm9yZVxuICAgICAgICAgICAgc3dhcFNwZWMuc2Nyb2xsID0gc2Nyb2xsVmFsXG4gICAgICAgICAgICBzd2FwU3BlYy5zY3JvbGxUYXJnZXQgPSBzZWxlY3RvclZhbFxuICAgICAgICAgIH0gZWxzZSBpZiAodmFsdWUuaW5kZXhPZignc2hvdzonKSA9PT0gMCkge1xuICAgICAgICAgICAgY29uc3Qgc2hvd1NwZWMgPSB2YWx1ZS5zdWJzdHIoNSlcbiAgICAgICAgICAgIHZhciBzcGxpdFNwZWMgPSBzaG93U3BlYy5zcGxpdCgnOicpXG4gICAgICAgICAgICBjb25zdCBzaG93VmFsID0gc3BsaXRTcGVjLnBvcCgpXG4gICAgICAgICAgICB2YXIgc2VsZWN0b3JWYWwgPSBzcGxpdFNwZWMubGVuZ3RoID4gMCA/IHNwbGl0U3BlYy5qb2luKCc6JykgOiBudWxsXG4gICAgICAgICAgICBzd2FwU3BlYy5zaG93ID0gc2hvd1ZhbFxuICAgICAgICAgICAgc3dhcFNwZWMuc2hvd1RhcmdldCA9IHNlbGVjdG9yVmFsXG4gICAgICAgICAgfSBlbHNlIGlmICh2YWx1ZS5pbmRleE9mKCdmb2N1cy1zY3JvbGw6JykgPT09IDApIHtcbiAgICAgICAgICAgIGNvbnN0IGZvY3VzU2Nyb2xsVmFsID0gdmFsdWUuc3Vic3RyKCdmb2N1cy1zY3JvbGw6Jy5sZW5ndGgpXG4gICAgICAgICAgICBzd2FwU3BlYy5mb2N1c1Njcm9sbCA9IGZvY3VzU2Nyb2xsVmFsID09ICd0cnVlJ1xuICAgICAgICAgIH0gZWxzZSBpZiAoaSA9PSAwKSB7XG4gICAgICAgICAgICBzd2FwU3BlYy5zd2FwU3R5bGUgPSB2YWx1ZVxuICAgICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgICBsb2dFcnJvcignVW5rbm93biBtb2RpZmllciBpbiBoeC1zd2FwOiAnICsgdmFsdWUpXG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiBzd2FwU3BlY1xuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm4ge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiB1c2VzRm9ybURhdGEoZWx0KSB7XG4gICAgcmV0dXJuIGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC1lbmNvZGluZycpID09PSAnbXVsdGlwYXJ0L2Zvcm0tZGF0YScgfHxcbiAgICAobWF0Y2hlcyhlbHQsICdmb3JtJykgJiYgZ2V0UmF3QXR0cmlidXRlKGVsdCwgJ2VuY3R5cGUnKSA9PT0gJ211bHRpcGFydC9mb3JtLWRhdGEnKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7WE1MSHR0cFJlcXVlc3R9IHhoclxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKiBAcGFyYW0ge0Zvcm1EYXRhfSBmaWx0ZXJlZFBhcmFtZXRlcnNcbiAgICogQHJldHVybnMgeyp8c3RyaW5nfG51bGx9XG4gICAqL1xuICBmdW5jdGlvbiBlbmNvZGVQYXJhbXNGb3JCb2R5KHhociwgZWx0LCBmaWx0ZXJlZFBhcmFtZXRlcnMpIHtcbiAgICBsZXQgZW5jb2RlZFBhcmFtZXRlcnMgPSBudWxsXG4gICAgd2l0aEV4dGVuc2lvbnMoZWx0LCBmdW5jdGlvbihleHRlbnNpb24pIHtcbiAgICAgIGlmIChlbmNvZGVkUGFyYW1ldGVycyA9PSBudWxsKSB7XG4gICAgICAgIGVuY29kZWRQYXJhbWV0ZXJzID0gZXh0ZW5zaW9uLmVuY29kZVBhcmFtZXRlcnMoeGhyLCBmaWx0ZXJlZFBhcmFtZXRlcnMsIGVsdClcbiAgICAgIH1cbiAgICB9KVxuICAgIGlmIChlbmNvZGVkUGFyYW1ldGVycyAhPSBudWxsKSB7XG4gICAgICByZXR1cm4gZW5jb2RlZFBhcmFtZXRlcnNcbiAgICB9IGVsc2Uge1xuICAgICAgaWYgKHVzZXNGb3JtRGF0YShlbHQpKSB7XG4gICAgICAgIC8vIEZvcmNlIGNvbnZlcnNpb24gdG8gYW4gYWN0dWFsIEZvcm1EYXRhIG9iamVjdCBpbiBjYXNlIGZpbHRlcmVkUGFyYW1ldGVycyBpcyBhIGZvcm1EYXRhUHJveHlcbiAgICAgICAgLy8gU2VlIGh0dHBzOi8vZ2l0aHViLmNvbS9iaWdza3lzb2Z0d2FyZS9odG14L2lzc3Vlcy8yMzE3XG4gICAgICAgIHJldHVybiBvdmVycmlkZUZvcm1EYXRhKG5ldyBGb3JtRGF0YSgpLCBmb3JtRGF0YUZyb21PYmplY3QoZmlsdGVyZWRQYXJhbWV0ZXJzKSlcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHJldHVybiB1cmxFbmNvZGUoZmlsdGVyZWRQYXJhbWV0ZXJzKVxuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICpcbiAqIEBwYXJhbSB7RWxlbWVudH0gdGFyZ2V0XG4gKiBAcmV0dXJucyB7SHRteFNldHRsZUluZm99XG4gKi9cbiAgZnVuY3Rpb24gbWFrZVNldHRsZUluZm8odGFyZ2V0KSB7XG4gICAgcmV0dXJuIHsgdGFza3M6IFtdLCBlbHRzOiBbdGFyZ2V0XSB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50W119IGNvbnRlbnRcbiAgICogQHBhcmFtIHtIdG14U3dhcFNwZWNpZmljYXRpb259IHN3YXBTcGVjXG4gICAqL1xuICBmdW5jdGlvbiB1cGRhdGVTY3JvbGxTdGF0ZShjb250ZW50LCBzd2FwU3BlYykge1xuICAgIGNvbnN0IGZpcnN0ID0gY29udGVudFswXVxuICAgIGNvbnN0IGxhc3QgPSBjb250ZW50W2NvbnRlbnQubGVuZ3RoIC0gMV1cbiAgICBpZiAoc3dhcFNwZWMuc2Nyb2xsKSB7XG4gICAgICB2YXIgdGFyZ2V0ID0gbnVsbFxuICAgICAgaWYgKHN3YXBTcGVjLnNjcm9sbFRhcmdldCkge1xuICAgICAgICB0YXJnZXQgPSBhc0VsZW1lbnQocXVlcnlTZWxlY3RvckV4dChmaXJzdCwgc3dhcFNwZWMuc2Nyb2xsVGFyZ2V0KSlcbiAgICAgIH1cbiAgICAgIGlmIChzd2FwU3BlYy5zY3JvbGwgPT09ICd0b3AnICYmIChmaXJzdCB8fCB0YXJnZXQpKSB7XG4gICAgICAgIHRhcmdldCA9IHRhcmdldCB8fCBmaXJzdFxuICAgICAgICB0YXJnZXQuc2Nyb2xsVG9wID0gMFxuICAgICAgfVxuICAgICAgaWYgKHN3YXBTcGVjLnNjcm9sbCA9PT0gJ2JvdHRvbScgJiYgKGxhc3QgfHwgdGFyZ2V0KSkge1xuICAgICAgICB0YXJnZXQgPSB0YXJnZXQgfHwgbGFzdFxuICAgICAgICB0YXJnZXQuc2Nyb2xsVG9wID0gdGFyZ2V0LnNjcm9sbEhlaWdodFxuICAgICAgfVxuICAgIH1cbiAgICBpZiAoc3dhcFNwZWMuc2hvdykge1xuICAgICAgdmFyIHRhcmdldCA9IG51bGxcbiAgICAgIGlmIChzd2FwU3BlYy5zaG93VGFyZ2V0KSB7XG4gICAgICAgIGxldCB0YXJnZXRTdHIgPSBzd2FwU3BlYy5zaG93VGFyZ2V0XG4gICAgICAgIGlmIChzd2FwU3BlYy5zaG93VGFyZ2V0ID09PSAnd2luZG93Jykge1xuICAgICAgICAgIHRhcmdldFN0ciA9ICdib2R5J1xuICAgICAgICB9XG4gICAgICAgIHRhcmdldCA9IGFzRWxlbWVudChxdWVyeVNlbGVjdG9yRXh0KGZpcnN0LCB0YXJnZXRTdHIpKVxuICAgICAgfVxuICAgICAgaWYgKHN3YXBTcGVjLnNob3cgPT09ICd0b3AnICYmIChmaXJzdCB8fCB0YXJnZXQpKSB7XG4gICAgICAgIHRhcmdldCA9IHRhcmdldCB8fCBmaXJzdFxuICAgICAgICAvLyBAdHMtaWdub3JlIEZvciBzb21lIHJlYXNvbiB0c2MgZG9lc24ndCByZWNvZ25pemUgXCJpbnN0YW50XCIgYXMgYSB2YWxpZCBvcHRpb24gZm9yIG5vd1xuICAgICAgICB0YXJnZXQuc2Nyb2xsSW50b1ZpZXcoeyBibG9jazogJ3N0YXJ0JywgYmVoYXZpb3I6IGh0bXguY29uZmlnLnNjcm9sbEJlaGF2aW9yIH0pXG4gICAgICB9XG4gICAgICBpZiAoc3dhcFNwZWMuc2hvdyA9PT0gJ2JvdHRvbScgJiYgKGxhc3QgfHwgdGFyZ2V0KSkge1xuICAgICAgICB0YXJnZXQgPSB0YXJnZXQgfHwgbGFzdFxuICAgICAgICAvLyBAdHMtaWdub3JlIEZvciBzb21lIHJlYXNvbiB0c2MgZG9lc24ndCByZWNvZ25pemUgXCJpbnN0YW50XCIgYXMgYSB2YWxpZCBvcHRpb24gZm9yIG5vd1xuICAgICAgICB0YXJnZXQuc2Nyb2xsSW50b1ZpZXcoeyBibG9jazogJ2VuZCcsIGJlaGF2aW9yOiBodG14LmNvbmZpZy5zY3JvbGxCZWhhdmlvciB9KVxuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAqIEBwYXJhbSB7c3RyaW5nfSBhdHRyXG4gKiBAcGFyYW0ge2Jvb2xlYW49fSBldmFsQXNEZWZhdWx0XG4gKiBAcGFyYW0ge09iamVjdD19IHZhbHVlc1xuICogQHJldHVybnMge09iamVjdH1cbiAqL1xuICBmdW5jdGlvbiBnZXRWYWx1ZXNGb3JFbGVtZW50KGVsdCwgYXR0ciwgZXZhbEFzRGVmYXVsdCwgdmFsdWVzKSB7XG4gICAgaWYgKHZhbHVlcyA9PSBudWxsKSB7XG4gICAgICB2YWx1ZXMgPSB7fVxuICAgIH1cbiAgICBpZiAoZWx0ID09IG51bGwpIHtcbiAgICAgIHJldHVybiB2YWx1ZXNcbiAgICB9XG4gICAgY29uc3QgYXR0cmlidXRlVmFsdWUgPSBnZXRBdHRyaWJ1dGVWYWx1ZShlbHQsIGF0dHIpXG4gICAgaWYgKGF0dHJpYnV0ZVZhbHVlKSB7XG4gICAgICBsZXQgc3RyID0gYXR0cmlidXRlVmFsdWUudHJpbSgpXG4gICAgICBsZXQgZXZhbHVhdGVWYWx1ZSA9IGV2YWxBc0RlZmF1bHRcbiAgICAgIGlmIChzdHIgPT09ICd1bnNldCcpIHtcbiAgICAgICAgcmV0dXJuIG51bGxcbiAgICAgIH1cbiAgICAgIGlmIChzdHIuaW5kZXhPZignamF2YXNjcmlwdDonKSA9PT0gMCkge1xuICAgICAgICBzdHIgPSBzdHIuc3Vic3RyKDExKVxuICAgICAgICBldmFsdWF0ZVZhbHVlID0gdHJ1ZVxuICAgICAgfSBlbHNlIGlmIChzdHIuaW5kZXhPZignanM6JykgPT09IDApIHtcbiAgICAgICAgc3RyID0gc3RyLnN1YnN0cigzKVxuICAgICAgICBldmFsdWF0ZVZhbHVlID0gdHJ1ZVxuICAgICAgfVxuICAgICAgaWYgKHN0ci5pbmRleE9mKCd7JykgIT09IDApIHtcbiAgICAgICAgc3RyID0gJ3snICsgc3RyICsgJ30nXG4gICAgICB9XG4gICAgICBsZXQgdmFyc1ZhbHVlc1xuICAgICAgaWYgKGV2YWx1YXRlVmFsdWUpIHtcbiAgICAgICAgdmFyc1ZhbHVlcyA9IG1heWJlRXZhbChlbHQsIGZ1bmN0aW9uKCkgeyByZXR1cm4gRnVuY3Rpb24oJ3JldHVybiAoJyArIHN0ciArICcpJykoKSB9LCB7fSlcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHZhcnNWYWx1ZXMgPSBwYXJzZUpTT04oc3RyKVxuICAgICAgfVxuICAgICAgZm9yIChjb25zdCBrZXkgaW4gdmFyc1ZhbHVlcykge1xuICAgICAgICBpZiAodmFyc1ZhbHVlcy5oYXNPd25Qcm9wZXJ0eShrZXkpKSB7XG4gICAgICAgICAgaWYgKHZhbHVlc1trZXldID09IG51bGwpIHtcbiAgICAgICAgICAgIHZhbHVlc1trZXldID0gdmFyc1ZhbHVlc1trZXldXG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiBnZXRWYWx1ZXNGb3JFbGVtZW50KGFzRWxlbWVudChwYXJlbnRFbHQoZWx0KSksIGF0dHIsIGV2YWxBc0RlZmF1bHQsIHZhbHVlcylcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0V2ZW50VGFyZ2V0fHN0cmluZ30gZWx0XG4gICAqIEBwYXJhbSB7KCkgPT4gYW55fSB0b0V2YWxcbiAgICogQHBhcmFtIHthbnk9fSBkZWZhdWx0VmFsXG4gICAqIEByZXR1cm5zIHthbnl9XG4gICAqL1xuICBmdW5jdGlvbiBtYXliZUV2YWwoZWx0LCB0b0V2YWwsIGRlZmF1bHRWYWwpIHtcbiAgICBpZiAoaHRteC5jb25maWcuYWxsb3dFdmFsKSB7XG4gICAgICByZXR1cm4gdG9FdmFsKClcbiAgICB9IGVsc2Uge1xuICAgICAgdHJpZ2dlckVycm9yRXZlbnQoZWx0LCAnaHRteDpldmFsRGlzYWxsb3dlZEVycm9yJylcbiAgICAgIHJldHVybiBkZWZhdWx0VmFsXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICogQHBhcmFtIHsqP30gZXhwcmVzc2lvblZhcnNcbiAqIEByZXR1cm5zXG4gKi9cbiAgZnVuY3Rpb24gZ2V0SFhWYXJzRm9yRWxlbWVudChlbHQsIGV4cHJlc3Npb25WYXJzKSB7XG4gICAgcmV0dXJuIGdldFZhbHVlc0ZvckVsZW1lbnQoZWx0LCAnaHgtdmFycycsIHRydWUsIGV4cHJlc3Npb25WYXJzKVxuICB9XG5cbiAgLyoqXG4gKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICogQHBhcmFtIHsqP30gZXhwcmVzc2lvblZhcnNcbiAqIEByZXR1cm5zXG4gKi9cbiAgZnVuY3Rpb24gZ2V0SFhWYWxzRm9yRWxlbWVudChlbHQsIGV4cHJlc3Npb25WYXJzKSB7XG4gICAgcmV0dXJuIGdldFZhbHVlc0ZvckVsZW1lbnQoZWx0LCAnaHgtdmFscycsIGZhbHNlLCBleHByZXNzaW9uVmFycylcbiAgfVxuXG4gIC8qKlxuICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAqIEByZXR1cm5zIHtGb3JtRGF0YX1cbiAqL1xuICBmdW5jdGlvbiBnZXRFeHByZXNzaW9uVmFycyhlbHQpIHtcbiAgICByZXR1cm4gbWVyZ2VPYmplY3RzKGdldEhYVmFyc0ZvckVsZW1lbnQoZWx0KSwgZ2V0SFhWYWxzRm9yRWxlbWVudChlbHQpKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7WE1MSHR0cFJlcXVlc3R9IHhoclxuICAgKiBAcGFyYW0ge3N0cmluZ30gaGVhZGVyXG4gICAqIEBwYXJhbSB7c3RyaW5nfG51bGx9IGhlYWRlclZhbHVlXG4gICAqL1xuICBmdW5jdGlvbiBzYWZlbHlTZXRIZWFkZXJWYWx1ZSh4aHIsIGhlYWRlciwgaGVhZGVyVmFsdWUpIHtcbiAgICBpZiAoaGVhZGVyVmFsdWUgIT09IG51bGwpIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIHhoci5zZXRSZXF1ZXN0SGVhZGVyKGhlYWRlciwgaGVhZGVyVmFsdWUpXG4gICAgICB9IGNhdGNoIChlKSB7XG4gICAgICAvLyBPbiBhbiBleGNlcHRpb24sIHRyeSB0byBzZXQgdGhlIGhlYWRlciBVUkkgZW5jb2RlZCBpbnN0ZWFkXG4gICAgICAgIHhoci5zZXRSZXF1ZXN0SGVhZGVyKGhlYWRlciwgZW5jb2RlVVJJQ29tcG9uZW50KGhlYWRlclZhbHVlKSlcbiAgICAgICAgeGhyLnNldFJlcXVlc3RIZWFkZXIoaGVhZGVyICsgJy1VUkktQXV0b0VuY29kZWQnLCAndHJ1ZScpXG4gICAgICB9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7WE1MSHR0cFJlcXVlc3R9IHhoclxuICAgKiBAcmV0dXJuIHtzdHJpbmd9XG4gICAqL1xuICBmdW5jdGlvbiBnZXRQYXRoRnJvbVJlc3BvbnNlKHhocikge1xuICAvLyBOQjogSUUxMSBkb2VzIG5vdCBzdXBwb3J0IHRoaXMgc3R1ZmZcbiAgICBpZiAoeGhyLnJlc3BvbnNlVVJMICYmIHR5cGVvZiAoVVJMKSAhPT0gJ3VuZGVmaW5lZCcpIHtcbiAgICAgIHRyeSB7XG4gICAgICAgIGNvbnN0IHVybCA9IG5ldyBVUkwoeGhyLnJlc3BvbnNlVVJMKVxuICAgICAgICByZXR1cm4gdXJsLnBhdGhuYW1lICsgdXJsLnNlYXJjaFxuICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICB0cmlnZ2VyRXJyb3JFdmVudChnZXREb2N1bWVudCgpLmJvZHksICdodG14OmJhZFJlc3BvbnNlVXJsJywgeyB1cmw6IHhoci5yZXNwb25zZVVSTCB9KVxuICAgICAgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge1hNTEh0dHBSZXF1ZXN0fSB4aHJcbiAgICogQHBhcmFtIHtSZWdFeHB9IHJlZ2V4cFxuICAgKiBAcmV0dXJuIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gaGFzSGVhZGVyKHhociwgcmVnZXhwKSB7XG4gICAgcmV0dXJuIHJlZ2V4cC50ZXN0KHhoci5nZXRBbGxSZXNwb25zZUhlYWRlcnMoKSlcbiAgfVxuXG4gIC8qKlxuICAgKiBJc3N1ZXMgYW4gaHRteC1zdHlsZSBBSkFYIHJlcXVlc3RcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jYWpheFxuICAgKlxuICAgKiBAcGFyYW0ge0h0dHBWZXJifSB2ZXJiXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBwYXRoIHRoZSBVUkwgcGF0aCB0byBtYWtlIHRoZSBBSkFYXG4gICAqIEBwYXJhbSB7RWxlbWVudHxzdHJpbmd8SHRteEFqYXhIZWxwZXJDb250ZXh0fSBjb250ZXh0IHRoZSBlbGVtZW50IHRvIHRhcmdldCAoZGVmYXVsdHMgdG8gdGhlICoqYm9keSoqKSB8IGEgc2VsZWN0b3IgZm9yIHRoZSB0YXJnZXQgfCBhIGNvbnRleHQgb2JqZWN0IHRoYXQgY29udGFpbnMgYW55IG9mIHRoZSBmb2xsb3dpbmdcbiAgICogQHJldHVybiB7UHJvbWlzZTx2b2lkPn0gUHJvbWlzZSB0aGF0IHJlc29sdmVzIGltbWVkaWF0ZWx5IGlmIG5vIHJlcXVlc3QgaXMgc2VudCwgb3Igd2hlbiB0aGUgcmVxdWVzdCBpcyBjb21wbGV0ZVxuICAgKi9cbiAgZnVuY3Rpb24gYWpheEhlbHBlcih2ZXJiLCBwYXRoLCBjb250ZXh0KSB7XG4gICAgdmVyYiA9ICgvKiogQHR5cGUgSHR0cFZlcmIgKi8odmVyYi50b0xvd2VyQ2FzZSgpKSlcbiAgICBpZiAoY29udGV4dCkge1xuICAgICAgaWYgKGNvbnRleHQgaW5zdGFuY2VvZiBFbGVtZW50IHx8IHR5cGVvZiBjb250ZXh0ID09PSAnc3RyaW5nJykge1xuICAgICAgICByZXR1cm4gaXNzdWVBamF4UmVxdWVzdCh2ZXJiLCBwYXRoLCBudWxsLCBudWxsLCB7XG4gICAgICAgICAgdGFyZ2V0T3ZlcnJpZGU6IHJlc29sdmVUYXJnZXQoY29udGV4dCksXG4gICAgICAgICAgcmV0dXJuUHJvbWlzZTogdHJ1ZVxuICAgICAgICB9KVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgcmV0dXJuIGlzc3VlQWpheFJlcXVlc3QodmVyYiwgcGF0aCwgcmVzb2x2ZVRhcmdldChjb250ZXh0LnNvdXJjZSksIGNvbnRleHQuZXZlbnQsXG4gICAgICAgICAge1xuICAgICAgICAgICAgaGFuZGxlcjogY29udGV4dC5oYW5kbGVyLFxuICAgICAgICAgICAgaGVhZGVyczogY29udGV4dC5oZWFkZXJzLFxuICAgICAgICAgICAgdmFsdWVzOiBjb250ZXh0LnZhbHVlcyxcbiAgICAgICAgICAgIHRhcmdldE92ZXJyaWRlOiByZXNvbHZlVGFyZ2V0KGNvbnRleHQudGFyZ2V0KSxcbiAgICAgICAgICAgIHN3YXBPdmVycmlkZTogY29udGV4dC5zd2FwLFxuICAgICAgICAgICAgc2VsZWN0OiBjb250ZXh0LnNlbGVjdCxcbiAgICAgICAgICAgIHJldHVyblByb21pc2U6IHRydWVcbiAgICAgICAgICB9KVxuICAgICAgfVxuICAgIH0gZWxzZSB7XG4gICAgICByZXR1cm4gaXNzdWVBamF4UmVxdWVzdCh2ZXJiLCBwYXRoLCBudWxsLCBudWxsLCB7XG4gICAgICAgIHJldHVyblByb21pc2U6IHRydWVcbiAgICAgIH0pXG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEByZXR1cm4ge0VsZW1lbnRbXX1cbiAgICovXG4gIGZ1bmN0aW9uIGhpZXJhcmNoeUZvckVsdChlbHQpIHtcbiAgICBjb25zdCBhcnIgPSBbXVxuICAgIHdoaWxlIChlbHQpIHtcbiAgICAgIGFyci5wdXNoKGVsdClcbiAgICAgIGVsdCA9IGVsdC5wYXJlbnRFbGVtZW50XG4gICAgfVxuICAgIHJldHVybiBhcnJcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0VsZW1lbnR9IGVsdFxuICAgKiBAcGFyYW0ge3N0cmluZ30gcGF0aFxuICAgKiBAcGFyYW0ge0h0bXhSZXF1ZXN0Q29uZmlnfSByZXF1ZXN0Q29uZmlnXG4gICAqIEByZXR1cm4ge2Jvb2xlYW59XG4gICAqL1xuICBmdW5jdGlvbiB2ZXJpZnlQYXRoKGVsdCwgcGF0aCwgcmVxdWVzdENvbmZpZykge1xuICAgIGxldCBzYW1lSG9zdFxuICAgIGxldCB1cmxcbiAgICBpZiAodHlwZW9mIFVSTCA9PT0gJ2Z1bmN0aW9uJykge1xuICAgICAgdXJsID0gbmV3IFVSTChwYXRoLCBkb2N1bWVudC5sb2NhdGlvbi5ocmVmKVxuICAgICAgY29uc3Qgb3JpZ2luID0gZG9jdW1lbnQubG9jYXRpb24ub3JpZ2luXG4gICAgICBzYW1lSG9zdCA9IG9yaWdpbiA9PT0gdXJsLm9yaWdpblxuICAgIH0gZWxzZSB7XG4gICAgLy8gSUUxMSBkb2Vzbid0IHN1cHBvcnQgVVJMXG4gICAgICB1cmwgPSBwYXRoXG4gICAgICBzYW1lSG9zdCA9IHN0YXJ0c1dpdGgocGF0aCwgZG9jdW1lbnQubG9jYXRpb24ub3JpZ2luKVxuICAgIH1cblxuICAgIGlmIChodG14LmNvbmZpZy5zZWxmUmVxdWVzdHNPbmx5KSB7XG4gICAgICBpZiAoIXNhbWVIb3N0KSB7XG4gICAgICAgIHJldHVybiBmYWxzZVxuICAgICAgfVxuICAgIH1cbiAgICByZXR1cm4gdHJpZ2dlckV2ZW50KGVsdCwgJ2h0bXg6dmFsaWRhdGVVcmwnLCBtZXJnZU9iamVjdHMoeyB1cmwsIHNhbWVIb3N0IH0sIHJlcXVlc3RDb25maWcpKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7T2JqZWN0fEZvcm1EYXRhfSBvYmpcbiAgICogQHJldHVybiB7Rm9ybURhdGF9XG4gICAqL1xuICBmdW5jdGlvbiBmb3JtRGF0YUZyb21PYmplY3Qob2JqKSB7XG4gICAgaWYgKG9iaiBpbnN0YW5jZW9mIEZvcm1EYXRhKSByZXR1cm4gb2JqXG4gICAgY29uc3QgZm9ybURhdGEgPSBuZXcgRm9ybURhdGEoKVxuICAgIGZvciAoY29uc3Qga2V5IGluIG9iaikge1xuICAgICAgaWYgKG9iai5oYXNPd25Qcm9wZXJ0eShrZXkpKSB7XG4gICAgICAgIGlmICh0eXBlb2Ygb2JqW2tleV0uZm9yRWFjaCA9PT0gJ2Z1bmN0aW9uJykge1xuICAgICAgICAgIG9ialtrZXldLmZvckVhY2goZnVuY3Rpb24odikgeyBmb3JtRGF0YS5hcHBlbmQoa2V5LCB2KSB9KVxuICAgICAgICB9IGVsc2UgaWYgKHR5cGVvZiBvYmpba2V5XSA9PT0gJ29iamVjdCcgJiYgIShvYmpba2V5XSBpbnN0YW5jZW9mIEJsb2IpKSB7XG4gICAgICAgICAgZm9ybURhdGEuYXBwZW5kKGtleSwgSlNPTi5zdHJpbmdpZnkob2JqW2tleV0pKVxuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIGZvcm1EYXRhLmFwcGVuZChrZXksIG9ialtrZXldKVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICAgIHJldHVybiBmb3JtRGF0YVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7Rm9ybURhdGF9IGZvcm1EYXRhXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBuYW1lXG4gICAqIEBwYXJhbSB7QXJyYXl9IGFycmF5XG4gICAqIEByZXR1cm5zIHtBcnJheX1cbiAgICovXG4gIGZ1bmN0aW9uIGZvcm1EYXRhQXJyYXlQcm94eShmb3JtRGF0YSwgbmFtZSwgYXJyYXkpIHtcbiAgICAvLyBtdXRhdGluZyB0aGUgYXJyYXkgc2hvdWxkIG11dGF0ZSB0aGUgdW5kZXJseWluZyBmb3JtIGRhdGFcbiAgICByZXR1cm4gbmV3IFByb3h5KGFycmF5LCB7XG4gICAgICBnZXQ6IGZ1bmN0aW9uKHRhcmdldCwga2V5KSB7XG4gICAgICAgIGlmICh0eXBlb2Yga2V5ID09PSAnbnVtYmVyJykgcmV0dXJuIHRhcmdldFtrZXldXG4gICAgICAgIGlmIChrZXkgPT09ICdsZW5ndGgnKSByZXR1cm4gdGFyZ2V0Lmxlbmd0aFxuICAgICAgICBpZiAoa2V5ID09PSAncHVzaCcpIHtcbiAgICAgICAgICByZXR1cm4gZnVuY3Rpb24odmFsdWUpIHtcbiAgICAgICAgICAgIHRhcmdldC5wdXNoKHZhbHVlKVxuICAgICAgICAgICAgZm9ybURhdGEuYXBwZW5kKG5hbWUsIHZhbHVlKVxuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBpZiAodHlwZW9mIHRhcmdldFtrZXldID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAgICAgcmV0dXJuIGZ1bmN0aW9uKCkge1xuICAgICAgICAgICAgdGFyZ2V0W2tleV0uYXBwbHkodGFyZ2V0LCBhcmd1bWVudHMpXG4gICAgICAgICAgICBmb3JtRGF0YS5kZWxldGUobmFtZSlcbiAgICAgICAgICAgIHRhcmdldC5mb3JFYWNoKGZ1bmN0aW9uKHYpIHsgZm9ybURhdGEuYXBwZW5kKG5hbWUsIHYpIH0pXG4gICAgICAgICAgfVxuICAgICAgICB9XG5cbiAgICAgICAgaWYgKHRhcmdldFtrZXldICYmIHRhcmdldFtrZXldLmxlbmd0aCA9PT0gMSkge1xuICAgICAgICAgIHJldHVybiB0YXJnZXRba2V5XVswXVxuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIHJldHVybiB0YXJnZXRba2V5XVxuICAgICAgICB9XG4gICAgICB9LFxuICAgICAgc2V0OiBmdW5jdGlvbih0YXJnZXQsIGluZGV4LCB2YWx1ZSkge1xuICAgICAgICB0YXJnZXRbaW5kZXhdID0gdmFsdWVcbiAgICAgICAgZm9ybURhdGEuZGVsZXRlKG5hbWUpXG4gICAgICAgIHRhcmdldC5mb3JFYWNoKGZ1bmN0aW9uKHYpIHsgZm9ybURhdGEuYXBwZW5kKG5hbWUsIHYpIH0pXG4gICAgICAgIHJldHVybiB0cnVlXG4gICAgICB9XG4gICAgfSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0Zvcm1EYXRhfSBmb3JtRGF0YVxuICAgKiBAcmV0dXJucyB7T2JqZWN0fVxuICAgKi9cbiAgZnVuY3Rpb24gZm9ybURhdGFQcm94eShmb3JtRGF0YSkge1xuICAgIHJldHVybiBuZXcgUHJveHkoZm9ybURhdGEsIHtcbiAgICAgIGdldDogZnVuY3Rpb24odGFyZ2V0LCBuYW1lKSB7XG4gICAgICAgIGlmICh0eXBlb2YgbmFtZSA9PT0gJ3N5bWJvbCcpIHtcbiAgICAgICAgICAvLyBGb3J3YXJkIHN5bWJvbCBjYWxscyB0byB0aGUgRm9ybURhdGEgaXRzZWxmIGRpcmVjdGx5XG4gICAgICAgICAgcmV0dXJuIFJlZmxlY3QuZ2V0KHRhcmdldCwgbmFtZSlcbiAgICAgICAgfVxuICAgICAgICBpZiAobmFtZSA9PT0gJ3RvSlNPTicpIHtcbiAgICAgICAgICAvLyBTdXBwb3J0IEpTT04uc3RyaW5naWZ5IGNhbGwgb24gcHJveHlcbiAgICAgICAgICByZXR1cm4gKCkgPT4gT2JqZWN0LmZyb21FbnRyaWVzKGZvcm1EYXRhKVxuICAgICAgICB9XG4gICAgICAgIGlmIChuYW1lIGluIHRhcmdldCkge1xuICAgICAgICAgIC8vIFdyYXAgaW4gZnVuY3Rpb24gd2l0aCBhcHBseSB0byBjb3JyZWN0bHkgYmluZCB0aGUgRm9ybURhdGEgY29udGV4dCwgYXMgYSBkaXJlY3QgY2FsbCB3b3VsZCByZXN1bHQgaW4gYW4gaWxsZWdhbCBpbnZvY2F0aW9uIGVycm9yXG4gICAgICAgICAgaWYgKHR5cGVvZiB0YXJnZXRbbmFtZV0gPT09ICdmdW5jdGlvbicpIHtcbiAgICAgICAgICAgIHJldHVybiBmdW5jdGlvbigpIHtcbiAgICAgICAgICAgICAgcmV0dXJuIGZvcm1EYXRhW25hbWVdLmFwcGx5KGZvcm1EYXRhLCBhcmd1bWVudHMpXG4gICAgICAgICAgICB9XG4gICAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIHJldHVybiB0YXJnZXRbbmFtZV1cbiAgICAgICAgICB9XG4gICAgICAgIH1cbiAgICAgICAgY29uc3QgYXJyYXkgPSBmb3JtRGF0YS5nZXRBbGwobmFtZSlcbiAgICAgICAgLy8gVGhvc2UgMiB1bmRlZmluZWQgJiBzaW5nbGUgdmFsdWUgcmV0dXJucyBhcmUgZm9yIHJldHJvLWNvbXBhdGliaWxpdHkgYXMgd2Ugd2VyZW4ndCB1c2luZyBGb3JtRGF0YSBiZWZvcmVcbiAgICAgICAgaWYgKGFycmF5Lmxlbmd0aCA9PT0gMCkge1xuICAgICAgICAgIHJldHVybiB1bmRlZmluZWRcbiAgICAgICAgfSBlbHNlIGlmIChhcnJheS5sZW5ndGggPT09IDEpIHtcbiAgICAgICAgICByZXR1cm4gYXJyYXlbMF1cbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICByZXR1cm4gZm9ybURhdGFBcnJheVByb3h5KHRhcmdldCwgbmFtZSwgYXJyYXkpXG4gICAgICAgIH1cbiAgICAgIH0sXG4gICAgICBzZXQ6IGZ1bmN0aW9uKHRhcmdldCwgbmFtZSwgdmFsdWUpIHtcbiAgICAgICAgaWYgKHR5cGVvZiBuYW1lICE9PSAnc3RyaW5nJykge1xuICAgICAgICAgIHJldHVybiBmYWxzZVxuICAgICAgICB9XG4gICAgICAgIHRhcmdldC5kZWxldGUobmFtZSlcbiAgICAgICAgaWYgKHR5cGVvZiB2YWx1ZS5mb3JFYWNoID09PSAnZnVuY3Rpb24nKSB7XG4gICAgICAgICAgdmFsdWUuZm9yRWFjaChmdW5jdGlvbih2KSB7IHRhcmdldC5hcHBlbmQobmFtZSwgdikgfSlcbiAgICAgICAgfSBlbHNlIGlmICh0eXBlb2YgdmFsdWUgPT09ICdvYmplY3QnICYmICEodmFsdWUgaW5zdGFuY2VvZiBCbG9iKSkge1xuICAgICAgICAgIHRhcmdldC5hcHBlbmQobmFtZSwgSlNPTi5zdHJpbmdpZnkodmFsdWUpKVxuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIHRhcmdldC5hcHBlbmQobmFtZSwgdmFsdWUpXG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuIHRydWVcbiAgICAgIH0sXG4gICAgICBkZWxldGVQcm9wZXJ0eTogZnVuY3Rpb24odGFyZ2V0LCBuYW1lKSB7XG4gICAgICAgIGlmICh0eXBlb2YgbmFtZSA9PT0gJ3N0cmluZycpIHtcbiAgICAgICAgICB0YXJnZXQuZGVsZXRlKG5hbWUpXG4gICAgICAgIH1cbiAgICAgICAgcmV0dXJuIHRydWVcbiAgICAgIH0sXG4gICAgICAvLyBTdXBwb3J0IE9iamVjdC5hc3NpZ24gY2FsbCBmcm9tIHByb3h5XG4gICAgICBvd25LZXlzOiBmdW5jdGlvbih0YXJnZXQpIHtcbiAgICAgICAgcmV0dXJuIFJlZmxlY3Qub3duS2V5cyhPYmplY3QuZnJvbUVudHJpZXModGFyZ2V0KSlcbiAgICAgIH0sXG4gICAgICBnZXRPd25Qcm9wZXJ0eURlc2NyaXB0b3I6IGZ1bmN0aW9uKHRhcmdldCwgcHJvcCkge1xuICAgICAgICByZXR1cm4gUmVmbGVjdC5nZXRPd25Qcm9wZXJ0eURlc2NyaXB0b3IoT2JqZWN0LmZyb21FbnRyaWVzKHRhcmdldCksIHByb3ApXG4gICAgICB9XG4gICAgfSlcbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge0h0dHBWZXJifSB2ZXJiXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBwYXRoXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEBwYXJhbSB7RXZlbnR9IGV2ZW50XG4gICAqIEBwYXJhbSB7SHRteEFqYXhFdGN9IFtldGNdXG4gICAqIEBwYXJhbSB7Ym9vbGVhbn0gW2NvbmZpcm1lZF1cbiAgICogQHJldHVybiB7UHJvbWlzZTx2b2lkPn1cbiAgICovXG4gIGZ1bmN0aW9uIGlzc3VlQWpheFJlcXVlc3QodmVyYiwgcGF0aCwgZWx0LCBldmVudCwgZXRjLCBjb25maXJtZWQpIHtcbiAgICBsZXQgcmVzb2x2ZSA9IG51bGxcbiAgICBsZXQgcmVqZWN0ID0gbnVsbFxuICAgIGV0YyA9IGV0YyAhPSBudWxsID8gZXRjIDoge31cbiAgICBpZiAoZXRjLnJldHVyblByb21pc2UgJiYgdHlwZW9mIFByb21pc2UgIT09ICd1bmRlZmluZWQnKSB7XG4gICAgICB2YXIgcHJvbWlzZSA9IG5ldyBQcm9taXNlKGZ1bmN0aW9uKF9yZXNvbHZlLCBfcmVqZWN0KSB7XG4gICAgICAgIHJlc29sdmUgPSBfcmVzb2x2ZVxuICAgICAgICByZWplY3QgPSBfcmVqZWN0XG4gICAgICB9KVxuICAgIH1cbiAgICBpZiAoZWx0ID09IG51bGwpIHtcbiAgICAgIGVsdCA9IGdldERvY3VtZW50KCkuYm9keVxuICAgIH1cbiAgICBjb25zdCByZXNwb25zZUhhbmRsZXIgPSBldGMuaGFuZGxlciB8fCBoYW5kbGVBamF4UmVzcG9uc2VcbiAgICBjb25zdCBzZWxlY3QgPSBldGMuc2VsZWN0IHx8IG51bGxcblxuICAgIGlmICghYm9keUNvbnRhaW5zKGVsdCkpIHtcbiAgICAvLyBkbyBub3QgaXNzdWUgcmVxdWVzdHMgZm9yIGVsZW1lbnRzIHJlbW92ZWQgZnJvbSB0aGUgRE9NXG4gICAgICBtYXliZUNhbGwocmVzb2x2ZSlcbiAgICAgIHJldHVybiBwcm9taXNlXG4gICAgfVxuICAgIGNvbnN0IHRhcmdldCA9IGV0Yy50YXJnZXRPdmVycmlkZSB8fCBhc0VsZW1lbnQoZ2V0VGFyZ2V0KGVsdCkpXG4gICAgaWYgKHRhcmdldCA9PSBudWxsIHx8IHRhcmdldCA9PSBEVU1NWV9FTFQpIHtcbiAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGVsdCwgJ2h0bXg6dGFyZ2V0RXJyb3InLCB7IHRhcmdldDogZ2V0QXR0cmlidXRlVmFsdWUoZWx0LCAnaHgtdGFyZ2V0JykgfSlcbiAgICAgIG1heWJlQ2FsbChyZWplY3QpXG4gICAgICByZXR1cm4gcHJvbWlzZVxuICAgIH1cblxuICAgIGxldCBlbHREYXRhID0gZ2V0SW50ZXJuYWxEYXRhKGVsdClcbiAgICBjb25zdCBzdWJtaXR0ZXIgPSBlbHREYXRhLmxhc3RCdXR0b25DbGlja2VkXG5cbiAgICBpZiAoc3VibWl0dGVyKSB7XG4gICAgICBjb25zdCBidXR0b25QYXRoID0gZ2V0UmF3QXR0cmlidXRlKHN1Ym1pdHRlciwgJ2Zvcm1hY3Rpb24nKVxuICAgICAgaWYgKGJ1dHRvblBhdGggIT0gbnVsbCkge1xuICAgICAgICBwYXRoID0gYnV0dG9uUGF0aFxuICAgICAgfVxuXG4gICAgICBjb25zdCBidXR0b25WZXJiID0gZ2V0UmF3QXR0cmlidXRlKHN1Ym1pdHRlciwgJ2Zvcm1tZXRob2QnKVxuICAgICAgaWYgKGJ1dHRvblZlcmIgIT0gbnVsbCkge1xuICAgICAgLy8gaWdub3JlIGJ1dHRvbnMgd2l0aCBmb3JtbWV0aG9kPVwiZGlhbG9nXCJcbiAgICAgICAgaWYgKGJ1dHRvblZlcmIudG9Mb3dlckNhc2UoKSAhPT0gJ2RpYWxvZycpIHtcbiAgICAgICAgICB2ZXJiID0gKC8qKiBAdHlwZSBIdHRwVmVyYiAqLyhidXR0b25WZXJiKSlcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIGNvbnN0IGNvbmZpcm1RdWVzdGlvbiA9IGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC1jb25maXJtJylcbiAgICAvLyBhbGxvdyBldmVudC1iYXNlZCBjb25maXJtYXRpb24gdy8gYSBjYWxsYmFja1xuICAgIGlmIChjb25maXJtZWQgPT09IHVuZGVmaW5lZCkge1xuICAgICAgY29uc3QgaXNzdWVSZXF1ZXN0ID0gZnVuY3Rpb24oc2tpcENvbmZpcm1hdGlvbikge1xuICAgICAgICByZXR1cm4gaXNzdWVBamF4UmVxdWVzdCh2ZXJiLCBwYXRoLCBlbHQsIGV2ZW50LCBldGMsICEhc2tpcENvbmZpcm1hdGlvbilcbiAgICAgIH1cbiAgICAgIGNvbnN0IGNvbmZpcm1EZXRhaWxzID0geyB0YXJnZXQsIGVsdCwgcGF0aCwgdmVyYiwgdHJpZ2dlcmluZ0V2ZW50OiBldmVudCwgZXRjLCBpc3N1ZVJlcXVlc3QsIHF1ZXN0aW9uOiBjb25maXJtUXVlc3Rpb24gfVxuICAgICAgaWYgKHRyaWdnZXJFdmVudChlbHQsICdodG14OmNvbmZpcm0nLCBjb25maXJtRGV0YWlscykgPT09IGZhbHNlKSB7XG4gICAgICAgIG1heWJlQ2FsbChyZXNvbHZlKVxuICAgICAgICByZXR1cm4gcHJvbWlzZVxuICAgICAgfVxuICAgIH1cblxuICAgIGxldCBzeW5jRWx0ID0gZWx0XG4gICAgbGV0IHN5bmNTdHJhdGVneSA9IGdldENsb3Nlc3RBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC1zeW5jJylcbiAgICBsZXQgcXVldWVTdHJhdGVneSA9IG51bGxcbiAgICBsZXQgYWJvcnRhYmxlID0gZmFsc2VcbiAgICBpZiAoc3luY1N0cmF0ZWd5KSB7XG4gICAgICBjb25zdCBzeW5jU3RyaW5ncyA9IHN5bmNTdHJhdGVneS5zcGxpdCgnOicpXG4gICAgICBjb25zdCBzZWxlY3RvciA9IHN5bmNTdHJpbmdzWzBdLnRyaW0oKVxuICAgICAgaWYgKHNlbGVjdG9yID09PSAndGhpcycpIHtcbiAgICAgICAgc3luY0VsdCA9IGZpbmRUaGlzRWxlbWVudChlbHQsICdoeC1zeW5jJylcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHN5bmNFbHQgPSBhc0VsZW1lbnQocXVlcnlTZWxlY3RvckV4dChlbHQsIHNlbGVjdG9yKSlcbiAgICAgIH1cbiAgICAgIC8vIGRlZmF1bHQgdG8gdGhlIGRyb3Agc3RyYXRlZ3lcbiAgICAgIHN5bmNTdHJhdGVneSA9IChzeW5jU3RyaW5nc1sxXSB8fCAnZHJvcCcpLnRyaW0oKVxuICAgICAgZWx0RGF0YSA9IGdldEludGVybmFsRGF0YShzeW5jRWx0KVxuICAgICAgaWYgKHN5bmNTdHJhdGVneSA9PT0gJ2Ryb3AnICYmIGVsdERhdGEueGhyICYmIGVsdERhdGEuYWJvcnRhYmxlICE9PSB0cnVlKSB7XG4gICAgICAgIG1heWJlQ2FsbChyZXNvbHZlKVxuICAgICAgICByZXR1cm4gcHJvbWlzZVxuICAgICAgfSBlbHNlIGlmIChzeW5jU3RyYXRlZ3kgPT09ICdhYm9ydCcpIHtcbiAgICAgICAgaWYgKGVsdERhdGEueGhyKSB7XG4gICAgICAgICAgbWF5YmVDYWxsKHJlc29sdmUpXG4gICAgICAgICAgcmV0dXJuIHByb21pc2VcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBhYm9ydGFibGUgPSB0cnVlXG4gICAgICAgIH1cbiAgICAgIH0gZWxzZSBpZiAoc3luY1N0cmF0ZWd5ID09PSAncmVwbGFjZScpIHtcbiAgICAgICAgdHJpZ2dlckV2ZW50KHN5bmNFbHQsICdodG14OmFib3J0JykgLy8gYWJvcnQgdGhlIGN1cnJlbnQgcmVxdWVzdCBhbmQgY29udGludWVcbiAgICAgIH0gZWxzZSBpZiAoc3luY1N0cmF0ZWd5LmluZGV4T2YoJ3F1ZXVlJykgPT09IDApIHtcbiAgICAgICAgY29uc3QgcXVldWVTdHJBcnJheSA9IHN5bmNTdHJhdGVneS5zcGxpdCgnICcpXG4gICAgICAgIHF1ZXVlU3RyYXRlZ3kgPSAocXVldWVTdHJBcnJheVsxXSB8fCAnbGFzdCcpLnRyaW0oKVxuICAgICAgfVxuICAgIH1cblxuICAgIGlmIChlbHREYXRhLnhocikge1xuICAgICAgaWYgKGVsdERhdGEuYWJvcnRhYmxlKSB7XG4gICAgICAgIHRyaWdnZXJFdmVudChzeW5jRWx0LCAnaHRteDphYm9ydCcpIC8vIGFib3J0IHRoZSBjdXJyZW50IHJlcXVlc3QgYW5kIGNvbnRpbnVlXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBpZiAocXVldWVTdHJhdGVneSA9PSBudWxsKSB7XG4gICAgICAgICAgaWYgKGV2ZW50KSB7XG4gICAgICAgICAgICBjb25zdCBldmVudERhdGEgPSBnZXRJbnRlcm5hbERhdGEoZXZlbnQpXG4gICAgICAgICAgICBpZiAoZXZlbnREYXRhICYmIGV2ZW50RGF0YS50cmlnZ2VyU3BlYyAmJiBldmVudERhdGEudHJpZ2dlclNwZWMucXVldWUpIHtcbiAgICAgICAgICAgICAgcXVldWVTdHJhdGVneSA9IGV2ZW50RGF0YS50cmlnZ2VyU3BlYy5xdWV1ZVxuICAgICAgICAgICAgfVxuICAgICAgICAgIH1cbiAgICAgICAgICBpZiAocXVldWVTdHJhdGVneSA9PSBudWxsKSB7XG4gICAgICAgICAgICBxdWV1ZVN0cmF0ZWd5ID0gJ2xhc3QnXG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICAgIGlmIChlbHREYXRhLnF1ZXVlZFJlcXVlc3RzID09IG51bGwpIHtcbiAgICAgICAgICBlbHREYXRhLnF1ZXVlZFJlcXVlc3RzID0gW11cbiAgICAgICAgfVxuICAgICAgICBpZiAocXVldWVTdHJhdGVneSA9PT0gJ2ZpcnN0JyAmJiBlbHREYXRhLnF1ZXVlZFJlcXVlc3RzLmxlbmd0aCA9PT0gMCkge1xuICAgICAgICAgIGVsdERhdGEucXVldWVkUmVxdWVzdHMucHVzaChmdW5jdGlvbigpIHtcbiAgICAgICAgICAgIGlzc3VlQWpheFJlcXVlc3QodmVyYiwgcGF0aCwgZWx0LCBldmVudCwgZXRjKVxuICAgICAgICAgIH0pXG4gICAgICAgIH0gZWxzZSBpZiAocXVldWVTdHJhdGVneSA9PT0gJ2FsbCcpIHtcbiAgICAgICAgICBlbHREYXRhLnF1ZXVlZFJlcXVlc3RzLnB1c2goZnVuY3Rpb24oKSB7XG4gICAgICAgICAgICBpc3N1ZUFqYXhSZXF1ZXN0KHZlcmIsIHBhdGgsIGVsdCwgZXZlbnQsIGV0YylcbiAgICAgICAgICB9KVxuICAgICAgICB9IGVsc2UgaWYgKHF1ZXVlU3RyYXRlZ3kgPT09ICdsYXN0Jykge1xuICAgICAgICAgIGVsdERhdGEucXVldWVkUmVxdWVzdHMgPSBbXSAvLyBkdW1wIGV4aXN0aW5nIHF1ZXVlXG4gICAgICAgICAgZWx0RGF0YS5xdWV1ZWRSZXF1ZXN0cy5wdXNoKGZ1bmN0aW9uKCkge1xuICAgICAgICAgICAgaXNzdWVBamF4UmVxdWVzdCh2ZXJiLCBwYXRoLCBlbHQsIGV2ZW50LCBldGMpXG4gICAgICAgICAgfSlcbiAgICAgICAgfVxuICAgICAgICBtYXliZUNhbGwocmVzb2x2ZSlcbiAgICAgICAgcmV0dXJuIHByb21pc2VcbiAgICAgIH1cbiAgICB9XG5cbiAgICBjb25zdCB4aHIgPSBuZXcgWE1MSHR0cFJlcXVlc3QoKVxuICAgIGVsdERhdGEueGhyID0geGhyXG4gICAgZWx0RGF0YS5hYm9ydGFibGUgPSBhYm9ydGFibGVcbiAgICBjb25zdCBlbmRSZXF1ZXN0TG9jayA9IGZ1bmN0aW9uKCkge1xuICAgICAgZWx0RGF0YS54aHIgPSBudWxsXG4gICAgICBlbHREYXRhLmFib3J0YWJsZSA9IGZhbHNlXG4gICAgICBpZiAoZWx0RGF0YS5xdWV1ZWRSZXF1ZXN0cyAhPSBudWxsICYmXG4gICAgICBlbHREYXRhLnF1ZXVlZFJlcXVlc3RzLmxlbmd0aCA+IDApIHtcbiAgICAgICAgY29uc3QgcXVldWVkUmVxdWVzdCA9IGVsdERhdGEucXVldWVkUmVxdWVzdHMuc2hpZnQoKVxuICAgICAgICBxdWV1ZWRSZXF1ZXN0KClcbiAgICAgIH1cbiAgICB9XG4gICAgY29uc3QgcHJvbXB0UXVlc3Rpb24gPSBnZXRDbG9zZXN0QXR0cmlidXRlVmFsdWUoZWx0LCAnaHgtcHJvbXB0JylcbiAgICBpZiAocHJvbXB0UXVlc3Rpb24pIHtcbiAgICAgIHZhciBwcm9tcHRSZXNwb25zZSA9IHByb21wdChwcm9tcHRRdWVzdGlvbilcbiAgICAgIC8vIHByb21wdCByZXR1cm5zIG51bGwgaWYgY2FuY2VsbGVkIGFuZCBlbXB0eSBzdHJpbmcgaWYgYWNjZXB0ZWQgd2l0aCBubyBlbnRyeVxuICAgICAgaWYgKHByb21wdFJlc3BvbnNlID09PSBudWxsIHx8XG4gICAgICAhdHJpZ2dlckV2ZW50KGVsdCwgJ2h0bXg6cHJvbXB0JywgeyBwcm9tcHQ6IHByb21wdFJlc3BvbnNlLCB0YXJnZXQgfSkpIHtcbiAgICAgICAgbWF5YmVDYWxsKHJlc29sdmUpXG4gICAgICAgIGVuZFJlcXVlc3RMb2NrKClcbiAgICAgICAgcmV0dXJuIHByb21pc2VcbiAgICAgIH1cbiAgICB9XG5cbiAgICBpZiAoY29uZmlybVF1ZXN0aW9uICYmICFjb25maXJtZWQpIHtcbiAgICAgIGlmICghY29uZmlybShjb25maXJtUXVlc3Rpb24pKSB7XG4gICAgICAgIG1heWJlQ2FsbChyZXNvbHZlKVxuICAgICAgICBlbmRSZXF1ZXN0TG9jaygpXG4gICAgICAgIHJldHVybiBwcm9taXNlXG4gICAgICB9XG4gICAgfVxuXG4gICAgbGV0IGhlYWRlcnMgPSBnZXRIZWFkZXJzKGVsdCwgdGFyZ2V0LCBwcm9tcHRSZXNwb25zZSlcblxuICAgIGlmICh2ZXJiICE9PSAnZ2V0JyAmJiAhdXNlc0Zvcm1EYXRhKGVsdCkpIHtcbiAgICAgIGhlYWRlcnNbJ0NvbnRlbnQtVHlwZSddID0gJ2FwcGxpY2F0aW9uL3gtd3d3LWZvcm0tdXJsZW5jb2RlZCdcbiAgICB9XG5cbiAgICBpZiAoZXRjLmhlYWRlcnMpIHtcbiAgICAgIGhlYWRlcnMgPSBtZXJnZU9iamVjdHMoaGVhZGVycywgZXRjLmhlYWRlcnMpXG4gICAgfVxuICAgIGNvbnN0IHJlc3VsdHMgPSBnZXRJbnB1dFZhbHVlcyhlbHQsIHZlcmIpXG4gICAgbGV0IGVycm9ycyA9IHJlc3VsdHMuZXJyb3JzXG4gICAgY29uc3QgcmF3Rm9ybURhdGEgPSByZXN1bHRzLmZvcm1EYXRhXG4gICAgaWYgKGV0Yy52YWx1ZXMpIHtcbiAgICAgIG92ZXJyaWRlRm9ybURhdGEocmF3Rm9ybURhdGEsIGZvcm1EYXRhRnJvbU9iamVjdChldGMudmFsdWVzKSlcbiAgICB9XG4gICAgY29uc3QgZXhwcmVzc2lvblZhcnMgPSBmb3JtRGF0YUZyb21PYmplY3QoZ2V0RXhwcmVzc2lvblZhcnMoZWx0KSlcbiAgICBjb25zdCBhbGxGb3JtRGF0YSA9IG92ZXJyaWRlRm9ybURhdGEocmF3Rm9ybURhdGEsIGV4cHJlc3Npb25WYXJzKVxuICAgIGxldCBmaWx0ZXJlZEZvcm1EYXRhID0gZmlsdGVyVmFsdWVzKGFsbEZvcm1EYXRhLCBlbHQpXG5cbiAgICBpZiAoaHRteC5jb25maWcuZ2V0Q2FjaGVCdXN0ZXJQYXJhbSAmJiB2ZXJiID09PSAnZ2V0Jykge1xuICAgICAgZmlsdGVyZWRGb3JtRGF0YS5zZXQoJ29yZy5odG14LmNhY2hlLWJ1c3RlcicsIGdldFJhd0F0dHJpYnV0ZSh0YXJnZXQsICdpZCcpIHx8ICd0cnVlJylcbiAgICB9XG5cbiAgICAvLyBiZWhhdmlvciBvZiBhbmNob3JzIHcvIGVtcHR5IGhyZWYgaXMgdG8gdXNlIHRoZSBjdXJyZW50IFVSTFxuICAgIGlmIChwYXRoID09IG51bGwgfHwgcGF0aCA9PT0gJycpIHtcbiAgICAgIHBhdGggPSBnZXREb2N1bWVudCgpLmxvY2F0aW9uLmhyZWZcbiAgICB9XG5cbiAgICAvKipcbiAgICAgKiBAdHlwZSB7T2JqZWN0fVxuICAgICAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gW2NyZWRlbnRpYWxzXVxuICAgICAqIEBwcm9wZXJ0eSB7bnVtYmVyfSBbdGltZW91dF1cbiAgICAgKiBAcHJvcGVydHkge2Jvb2xlYW59IFtub0hlYWRlcnNdXG4gICAgICovXG4gICAgY29uc3QgcmVxdWVzdEF0dHJWYWx1ZXMgPSBnZXRWYWx1ZXNGb3JFbGVtZW50KGVsdCwgJ2h4LXJlcXVlc3QnKVxuXG4gICAgY29uc3QgZWx0SXNCb29zdGVkID0gZ2V0SW50ZXJuYWxEYXRhKGVsdCkuYm9vc3RlZFxuXG4gICAgbGV0IHVzZVVybFBhcmFtcyA9IGh0bXguY29uZmlnLm1ldGhvZHNUaGF0VXNlVXJsUGFyYW1zLmluZGV4T2YodmVyYikgPj0gMFxuXG4gICAgLyoqIEB0eXBlIEh0bXhSZXF1ZXN0Q29uZmlnICovXG4gICAgY29uc3QgcmVxdWVzdENvbmZpZyA9IHtcbiAgICAgIGJvb3N0ZWQ6IGVsdElzQm9vc3RlZCxcbiAgICAgIHVzZVVybFBhcmFtcyxcbiAgICAgIGZvcm1EYXRhOiBmaWx0ZXJlZEZvcm1EYXRhLFxuICAgICAgcGFyYW1ldGVyczogZm9ybURhdGFQcm94eShmaWx0ZXJlZEZvcm1EYXRhKSxcbiAgICAgIHVuZmlsdGVyZWRGb3JtRGF0YTogYWxsRm9ybURhdGEsXG4gICAgICB1bmZpbHRlcmVkUGFyYW1ldGVyczogZm9ybURhdGFQcm94eShhbGxGb3JtRGF0YSksXG4gICAgICBoZWFkZXJzLFxuICAgICAgdGFyZ2V0LFxuICAgICAgdmVyYixcbiAgICAgIGVycm9ycyxcbiAgICAgIHdpdGhDcmVkZW50aWFsczogZXRjLmNyZWRlbnRpYWxzIHx8IHJlcXVlc3RBdHRyVmFsdWVzLmNyZWRlbnRpYWxzIHx8IGh0bXguY29uZmlnLndpdGhDcmVkZW50aWFscyxcbiAgICAgIHRpbWVvdXQ6IGV0Yy50aW1lb3V0IHx8IHJlcXVlc3RBdHRyVmFsdWVzLnRpbWVvdXQgfHwgaHRteC5jb25maWcudGltZW91dCxcbiAgICAgIHBhdGgsXG4gICAgICB0cmlnZ2VyaW5nRXZlbnQ6IGV2ZW50XG4gICAgfVxuXG4gICAgaWYgKCF0cmlnZ2VyRXZlbnQoZWx0LCAnaHRteDpjb25maWdSZXF1ZXN0JywgcmVxdWVzdENvbmZpZykpIHtcbiAgICAgIG1heWJlQ2FsbChyZXNvbHZlKVxuICAgICAgZW5kUmVxdWVzdExvY2soKVxuICAgICAgcmV0dXJuIHByb21pc2VcbiAgICB9XG5cbiAgICAvLyBjb3B5IG91dCBpbiBjYXNlIHRoZSBvYmplY3Qgd2FzIG92ZXJ3cml0dGVuXG4gICAgcGF0aCA9IHJlcXVlc3RDb25maWcucGF0aFxuICAgIHZlcmIgPSByZXF1ZXN0Q29uZmlnLnZlcmJcbiAgICBoZWFkZXJzID0gcmVxdWVzdENvbmZpZy5oZWFkZXJzXG4gICAgZmlsdGVyZWRGb3JtRGF0YSA9IGZvcm1EYXRhRnJvbU9iamVjdChyZXF1ZXN0Q29uZmlnLnBhcmFtZXRlcnMpXG4gICAgZXJyb3JzID0gcmVxdWVzdENvbmZpZy5lcnJvcnNcbiAgICB1c2VVcmxQYXJhbXMgPSByZXF1ZXN0Q29uZmlnLnVzZVVybFBhcmFtc1xuXG4gICAgaWYgKGVycm9ycyAmJiBlcnJvcnMubGVuZ3RoID4gMCkge1xuICAgICAgdHJpZ2dlckV2ZW50KGVsdCwgJ2h0bXg6dmFsaWRhdGlvbjpoYWx0ZWQnLCByZXF1ZXN0Q29uZmlnKVxuICAgICAgbWF5YmVDYWxsKHJlc29sdmUpXG4gICAgICBlbmRSZXF1ZXN0TG9jaygpXG4gICAgICByZXR1cm4gcHJvbWlzZVxuICAgIH1cblxuICAgIGNvbnN0IHNwbGl0UGF0aCA9IHBhdGguc3BsaXQoJyMnKVxuICAgIGNvbnN0IHBhdGhOb0FuY2hvciA9IHNwbGl0UGF0aFswXVxuICAgIGNvbnN0IGFuY2hvciA9IHNwbGl0UGF0aFsxXVxuXG4gICAgbGV0IGZpbmFsUGF0aCA9IHBhdGhcbiAgICBpZiAodXNlVXJsUGFyYW1zKSB7XG4gICAgICBmaW5hbFBhdGggPSBwYXRoTm9BbmNob3JcbiAgICAgIGNvbnN0IGhhc1ZhbHVlcyA9ICFmaWx0ZXJlZEZvcm1EYXRhLmtleXMoKS5uZXh0KCkuZG9uZVxuICAgICAgaWYgKGhhc1ZhbHVlcykge1xuICAgICAgICBpZiAoZmluYWxQYXRoLmluZGV4T2YoJz8nKSA8IDApIHtcbiAgICAgICAgICBmaW5hbFBhdGggKz0gJz8nXG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgZmluYWxQYXRoICs9ICcmJ1xuICAgICAgICB9XG4gICAgICAgIGZpbmFsUGF0aCArPSB1cmxFbmNvZGUoZmlsdGVyZWRGb3JtRGF0YSlcbiAgICAgICAgaWYgKGFuY2hvcikge1xuICAgICAgICAgIGZpbmFsUGF0aCArPSAnIycgKyBhbmNob3JcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIGlmICghdmVyaWZ5UGF0aChlbHQsIGZpbmFsUGF0aCwgcmVxdWVzdENvbmZpZykpIHtcbiAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGVsdCwgJ2h0bXg6aW52YWxpZFBhdGgnLCByZXF1ZXN0Q29uZmlnKVxuICAgICAgbWF5YmVDYWxsKHJlamVjdClcbiAgICAgIHJldHVybiBwcm9taXNlXG4gICAgfVxuXG4gICAgeGhyLm9wZW4odmVyYi50b1VwcGVyQ2FzZSgpLCBmaW5hbFBhdGgsIHRydWUpXG4gICAgeGhyLm92ZXJyaWRlTWltZVR5cGUoJ3RleHQvaHRtbCcpXG4gICAgeGhyLndpdGhDcmVkZW50aWFscyA9IHJlcXVlc3RDb25maWcud2l0aENyZWRlbnRpYWxzXG4gICAgeGhyLnRpbWVvdXQgPSByZXF1ZXN0Q29uZmlnLnRpbWVvdXRcblxuICAgIC8vIHJlcXVlc3QgaGVhZGVyc1xuICAgIGlmIChyZXF1ZXN0QXR0clZhbHVlcy5ub0hlYWRlcnMpIHtcbiAgICAvLyBpZ25vcmUgYWxsIGhlYWRlcnNcbiAgICB9IGVsc2Uge1xuICAgICAgZm9yIChjb25zdCBoZWFkZXIgaW4gaGVhZGVycykge1xuICAgICAgICBpZiAoaGVhZGVycy5oYXNPd25Qcm9wZXJ0eShoZWFkZXIpKSB7XG4gICAgICAgICAgY29uc3QgaGVhZGVyVmFsdWUgPSBoZWFkZXJzW2hlYWRlcl1cbiAgICAgICAgICBzYWZlbHlTZXRIZWFkZXJWYWx1ZSh4aHIsIGhlYWRlciwgaGVhZGVyVmFsdWUpXG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICAvKiogQHR5cGUge0h0bXhSZXNwb25zZUluZm99ICovXG4gICAgY29uc3QgcmVzcG9uc2VJbmZvID0ge1xuICAgICAgeGhyLFxuICAgICAgdGFyZ2V0LFxuICAgICAgcmVxdWVzdENvbmZpZyxcbiAgICAgIGV0YyxcbiAgICAgIGJvb3N0ZWQ6IGVsdElzQm9vc3RlZCxcbiAgICAgIHNlbGVjdCxcbiAgICAgIHBhdGhJbmZvOiB7XG4gICAgICAgIHJlcXVlc3RQYXRoOiBwYXRoLFxuICAgICAgICBmaW5hbFJlcXVlc3RQYXRoOiBmaW5hbFBhdGgsXG4gICAgICAgIHJlc3BvbnNlUGF0aDogbnVsbCxcbiAgICAgICAgYW5jaG9yXG4gICAgICB9XG4gICAgfVxuXG4gICAgeGhyLm9ubG9hZCA9IGZ1bmN0aW9uKCkge1xuICAgICAgdHJ5IHtcbiAgICAgICAgY29uc3QgaGllcmFyY2h5ID0gaGllcmFyY2h5Rm9yRWx0KGVsdClcbiAgICAgICAgcmVzcG9uc2VJbmZvLnBhdGhJbmZvLnJlc3BvbnNlUGF0aCA9IGdldFBhdGhGcm9tUmVzcG9uc2UoeGhyKVxuICAgICAgICByZXNwb25zZUhhbmRsZXIoZWx0LCByZXNwb25zZUluZm8pXG4gICAgICAgIGlmIChyZXNwb25zZUluZm8ua2VlcEluZGljYXRvcnMgIT09IHRydWUpIHtcbiAgICAgICAgICByZW1vdmVSZXF1ZXN0SW5kaWNhdG9ycyhpbmRpY2F0b3JzLCBkaXNhYmxlRWx0cylcbiAgICAgICAgfVxuICAgICAgICB0cmlnZ2VyRXZlbnQoZWx0LCAnaHRteDphZnRlclJlcXVlc3QnLCByZXNwb25zZUluZm8pXG4gICAgICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14OmFmdGVyT25Mb2FkJywgcmVzcG9uc2VJbmZvKVxuICAgICAgICAvLyBpZiB0aGUgYm9keSBubyBsb25nZXIgY29udGFpbnMgdGhlIGVsZW1lbnQsIHRyaWdnZXIgdGhlIGV2ZW50IG9uIHRoZSBjbG9zZXN0IHBhcmVudFxuICAgICAgICAvLyByZW1haW5pbmcgaW4gdGhlIERPTVxuICAgICAgICBpZiAoIWJvZHlDb250YWlucyhlbHQpKSB7XG4gICAgICAgICAgbGV0IHNlY29uZGFyeVRyaWdnZXJFbHQgPSBudWxsXG4gICAgICAgICAgd2hpbGUgKGhpZXJhcmNoeS5sZW5ndGggPiAwICYmIHNlY29uZGFyeVRyaWdnZXJFbHQgPT0gbnVsbCkge1xuICAgICAgICAgICAgY29uc3QgcGFyZW50RWx0SW5IaWVyYXJjaHkgPSBoaWVyYXJjaHkuc2hpZnQoKVxuICAgICAgICAgICAgaWYgKGJvZHlDb250YWlucyhwYXJlbnRFbHRJbkhpZXJhcmNoeSkpIHtcbiAgICAgICAgICAgICAgc2Vjb25kYXJ5VHJpZ2dlckVsdCA9IHBhcmVudEVsdEluSGllcmFyY2h5XG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuICAgICAgICAgIGlmIChzZWNvbmRhcnlUcmlnZ2VyRWx0KSB7XG4gICAgICAgICAgICB0cmlnZ2VyRXZlbnQoc2Vjb25kYXJ5VHJpZ2dlckVsdCwgJ2h0bXg6YWZ0ZXJSZXF1ZXN0JywgcmVzcG9uc2VJbmZvKVxuICAgICAgICAgICAgdHJpZ2dlckV2ZW50KHNlY29uZGFyeVRyaWdnZXJFbHQsICdodG14OmFmdGVyT25Mb2FkJywgcmVzcG9uc2VJbmZvKVxuICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBtYXliZUNhbGwocmVzb2x2ZSlcbiAgICAgICAgZW5kUmVxdWVzdExvY2soKVxuICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICB0cmlnZ2VyRXJyb3JFdmVudChlbHQsICdodG14Om9uTG9hZEVycm9yJywgbWVyZ2VPYmplY3RzKHsgZXJyb3I6IGUgfSwgcmVzcG9uc2VJbmZvKSlcbiAgICAgICAgdGhyb3cgZVxuICAgICAgfVxuICAgIH1cbiAgICB4aHIub25lcnJvciA9IGZ1bmN0aW9uKCkge1xuICAgICAgcmVtb3ZlUmVxdWVzdEluZGljYXRvcnMoaW5kaWNhdG9ycywgZGlzYWJsZUVsdHMpXG4gICAgICB0cmlnZ2VyRXJyb3JFdmVudChlbHQsICdodG14OmFmdGVyUmVxdWVzdCcsIHJlc3BvbnNlSW5mbylcbiAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGVsdCwgJ2h0bXg6c2VuZEVycm9yJywgcmVzcG9uc2VJbmZvKVxuICAgICAgbWF5YmVDYWxsKHJlamVjdClcbiAgICAgIGVuZFJlcXVlc3RMb2NrKClcbiAgICB9XG4gICAgeGhyLm9uYWJvcnQgPSBmdW5jdGlvbigpIHtcbiAgICAgIHJlbW92ZVJlcXVlc3RJbmRpY2F0b3JzKGluZGljYXRvcnMsIGRpc2FibGVFbHRzKVxuICAgICAgdHJpZ2dlckVycm9yRXZlbnQoZWx0LCAnaHRteDphZnRlclJlcXVlc3QnLCByZXNwb25zZUluZm8pXG4gICAgICB0cmlnZ2VyRXJyb3JFdmVudChlbHQsICdodG14OnNlbmRBYm9ydCcsIHJlc3BvbnNlSW5mbylcbiAgICAgIG1heWJlQ2FsbChyZWplY3QpXG4gICAgICBlbmRSZXF1ZXN0TG9jaygpXG4gICAgfVxuICAgIHhoci5vbnRpbWVvdXQgPSBmdW5jdGlvbigpIHtcbiAgICAgIHJlbW92ZVJlcXVlc3RJbmRpY2F0b3JzKGluZGljYXRvcnMsIGRpc2FibGVFbHRzKVxuICAgICAgdHJpZ2dlckVycm9yRXZlbnQoZWx0LCAnaHRteDphZnRlclJlcXVlc3QnLCByZXNwb25zZUluZm8pXG4gICAgICB0cmlnZ2VyRXJyb3JFdmVudChlbHQsICdodG14OnRpbWVvdXQnLCByZXNwb25zZUluZm8pXG4gICAgICBtYXliZUNhbGwocmVqZWN0KVxuICAgICAgZW5kUmVxdWVzdExvY2soKVxuICAgIH1cbiAgICBpZiAoIXRyaWdnZXJFdmVudChlbHQsICdodG14OmJlZm9yZVJlcXVlc3QnLCByZXNwb25zZUluZm8pKSB7XG4gICAgICBtYXliZUNhbGwocmVzb2x2ZSlcbiAgICAgIGVuZFJlcXVlc3RMb2NrKClcbiAgICAgIHJldHVybiBwcm9taXNlXG4gICAgfVxuICAgIHZhciBpbmRpY2F0b3JzID0gYWRkUmVxdWVzdEluZGljYXRvckNsYXNzZXMoZWx0KVxuICAgIHZhciBkaXNhYmxlRWx0cyA9IGRpc2FibGVFbGVtZW50cyhlbHQpXG5cbiAgICBmb3JFYWNoKFsnbG9hZHN0YXJ0JywgJ2xvYWRlbmQnLCAncHJvZ3Jlc3MnLCAnYWJvcnQnXSwgZnVuY3Rpb24oZXZlbnROYW1lKSB7XG4gICAgICBmb3JFYWNoKFt4aHIsIHhoci51cGxvYWRdLCBmdW5jdGlvbih0YXJnZXQpIHtcbiAgICAgICAgdGFyZ2V0LmFkZEV2ZW50TGlzdGVuZXIoZXZlbnROYW1lLCBmdW5jdGlvbihldmVudCkge1xuICAgICAgICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14OnhocjonICsgZXZlbnROYW1lLCB7XG4gICAgICAgICAgICBsZW5ndGhDb21wdXRhYmxlOiBldmVudC5sZW5ndGhDb21wdXRhYmxlLFxuICAgICAgICAgICAgbG9hZGVkOiBldmVudC5sb2FkZWQsXG4gICAgICAgICAgICB0b3RhbDogZXZlbnQudG90YWxcbiAgICAgICAgICB9KVxuICAgICAgICB9KVxuICAgICAgfSlcbiAgICB9KVxuICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14OmJlZm9yZVNlbmQnLCByZXNwb25zZUluZm8pXG4gICAgY29uc3QgcGFyYW1zID0gdXNlVXJsUGFyYW1zID8gbnVsbCA6IGVuY29kZVBhcmFtc0ZvckJvZHkoeGhyLCBlbHQsIGZpbHRlcmVkRm9ybURhdGEpXG4gICAgeGhyLnNlbmQocGFyYW1zKVxuICAgIHJldHVybiBwcm9taXNlXG4gIH1cblxuICAvKipcbiAgICogQHR5cGVkZWYge09iamVjdH0gSHRteEhpc3RvcnlVcGRhdGVcbiAgICogQHByb3BlcnR5IHtzdHJpbmd8bnVsbH0gW3R5cGVdXG4gICAqIEBwcm9wZXJ0eSB7c3RyaW5nfG51bGx9IFtwYXRoXVxuICAgKi9cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHBhcmFtIHtIdG14UmVzcG9uc2VJbmZvfSByZXNwb25zZUluZm9cbiAgICogQHJldHVybiB7SHRteEhpc3RvcnlVcGRhdGV9XG4gICAqL1xuICBmdW5jdGlvbiBkZXRlcm1pbmVIaXN0b3J5VXBkYXRlcyhlbHQsIHJlc3BvbnNlSW5mbykge1xuICAgIGNvbnN0IHhociA9IHJlc3BvbnNlSW5mby54aHJcblxuICAgIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT1cbiAgICAvLyBGaXJzdCBjb25zdWx0IHJlc3BvbnNlIGhlYWRlcnNcbiAgICAvLz0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG4gICAgbGV0IHBhdGhGcm9tSGVhZGVycyA9IG51bGxcbiAgICBsZXQgdHlwZUZyb21IZWFkZXJzID0gbnVsbFxuICAgIGlmIChoYXNIZWFkZXIoeGhyLCAvSFgtUHVzaDovaSkpIHtcbiAgICAgIHBhdGhGcm9tSGVhZGVycyA9IHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUHVzaCcpXG4gICAgICB0eXBlRnJvbUhlYWRlcnMgPSAncHVzaCdcbiAgICB9IGVsc2UgaWYgKGhhc0hlYWRlcih4aHIsIC9IWC1QdXNoLVVybDovaSkpIHtcbiAgICAgIHBhdGhGcm9tSGVhZGVycyA9IHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUHVzaC1VcmwnKVxuICAgICAgdHlwZUZyb21IZWFkZXJzID0gJ3B1c2gnXG4gICAgfSBlbHNlIGlmIChoYXNIZWFkZXIoeGhyLCAvSFgtUmVwbGFjZS1Vcmw6L2kpKSB7XG4gICAgICBwYXRoRnJvbUhlYWRlcnMgPSB4aHIuZ2V0UmVzcG9uc2VIZWFkZXIoJ0hYLVJlcGxhY2UtVXJsJylcbiAgICAgIHR5cGVGcm9tSGVhZGVycyA9ICdyZXBsYWNlJ1xuICAgIH1cblxuICAgIC8vIGlmIHRoZXJlIHdhcyBhIHJlc3BvbnNlIGhlYWRlciwgdGhhdCBoYXMgcHJpb3JpdHlcbiAgICBpZiAocGF0aEZyb21IZWFkZXJzKSB7XG4gICAgICBpZiAocGF0aEZyb21IZWFkZXJzID09PSAnZmFsc2UnKSB7XG4gICAgICAgIHJldHVybiB7fVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgcmV0dXJuIHtcbiAgICAgICAgICB0eXBlOiB0eXBlRnJvbUhlYWRlcnMsXG4gICAgICAgICAgcGF0aDogcGF0aEZyb21IZWFkZXJzXG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG5cbiAgICAvLz0gPT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG4gICAgLy8gTmV4dCByZXNvbHZlIHZpYSBET00gdmFsdWVzXG4gICAgLy89ID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PVxuICAgIGNvbnN0IHJlcXVlc3RQYXRoID0gcmVzcG9uc2VJbmZvLnBhdGhJbmZvLmZpbmFsUmVxdWVzdFBhdGhcbiAgICBjb25zdCByZXNwb25zZVBhdGggPSByZXNwb25zZUluZm8ucGF0aEluZm8ucmVzcG9uc2VQYXRoXG5cbiAgICBjb25zdCBwdXNoVXJsID0gZ2V0Q2xvc2VzdEF0dHJpYnV0ZVZhbHVlKGVsdCwgJ2h4LXB1c2gtdXJsJylcbiAgICBjb25zdCByZXBsYWNlVXJsID0gZ2V0Q2xvc2VzdEF0dHJpYnV0ZVZhbHVlKGVsdCwgJ2h4LXJlcGxhY2UtdXJsJylcbiAgICBjb25zdCBlbGVtZW50SXNCb29zdGVkID0gZ2V0SW50ZXJuYWxEYXRhKGVsdCkuYm9vc3RlZFxuXG4gICAgbGV0IHNhdmVUeXBlID0gbnVsbFxuICAgIGxldCBwYXRoID0gbnVsbFxuXG4gICAgaWYgKHB1c2hVcmwpIHtcbiAgICAgIHNhdmVUeXBlID0gJ3B1c2gnXG4gICAgICBwYXRoID0gcHVzaFVybFxuICAgIH0gZWxzZSBpZiAocmVwbGFjZVVybCkge1xuICAgICAgc2F2ZVR5cGUgPSAncmVwbGFjZSdcbiAgICAgIHBhdGggPSByZXBsYWNlVXJsXG4gICAgfSBlbHNlIGlmIChlbGVtZW50SXNCb29zdGVkKSB7XG4gICAgICBzYXZlVHlwZSA9ICdwdXNoJ1xuICAgICAgcGF0aCA9IHJlc3BvbnNlUGF0aCB8fCByZXF1ZXN0UGF0aCAvLyBpZiB0aGVyZSBpcyBubyByZXNwb25zZSBwYXRoLCBnbyB3aXRoIHRoZSBvcmlnaW5hbCByZXF1ZXN0IHBhdGhcbiAgICB9XG5cbiAgICBpZiAocGF0aCkge1xuICAgIC8vIGZhbHNlIGluZGljYXRlcyBubyBwdXNoLCByZXR1cm4gZW1wdHkgb2JqZWN0XG4gICAgICBpZiAocGF0aCA9PT0gJ2ZhbHNlJykge1xuICAgICAgICByZXR1cm4ge31cbiAgICAgIH1cblxuICAgICAgLy8gdHJ1ZSBpbmRpY2F0ZXMgd2Ugd2FudCB0byBmb2xsb3cgd2hlcmV2ZXIgdGhlIHNlcnZlciBlbmRlZCB1cCBzZW5kaW5nIHVzXG4gICAgICBpZiAocGF0aCA9PT0gJ3RydWUnKSB7XG4gICAgICAgIHBhdGggPSByZXNwb25zZVBhdGggfHwgcmVxdWVzdFBhdGggLy8gaWYgdGhlcmUgaXMgbm8gcmVzcG9uc2UgcGF0aCwgZ28gd2l0aCB0aGUgb3JpZ2luYWwgcmVxdWVzdCBwYXRoXG4gICAgICB9XG5cbiAgICAgIC8vIHJlc3RvcmUgYW55IGFuY2hvciBhc3NvY2lhdGVkIHdpdGggdGhlIHJlcXVlc3RcbiAgICAgIGlmIChyZXNwb25zZUluZm8ucGF0aEluZm8uYW5jaG9yICYmIHBhdGguaW5kZXhPZignIycpID09PSAtMSkge1xuICAgICAgICBwYXRoID0gcGF0aCArICcjJyArIHJlc3BvbnNlSW5mby5wYXRoSW5mby5hbmNob3JcbiAgICAgIH1cblxuICAgICAgcmV0dXJuIHtcbiAgICAgICAgdHlwZTogc2F2ZVR5cGUsXG4gICAgICAgIHBhdGhcbiAgICAgIH1cbiAgICB9IGVsc2Uge1xuICAgICAgcmV0dXJuIHt9XG4gICAgfVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7SHRteFJlc3BvbnNlSGFuZGxpbmdDb25maWd9IHJlc3BvbnNlSGFuZGxpbmdDb25maWdcbiAgICogQHBhcmFtIHtudW1iZXJ9IHN0YXR1c1xuICAgKiBAcmV0dXJuIHtib29sZWFufVxuICAgKi9cbiAgZnVuY3Rpb24gY29kZU1hdGNoZXMocmVzcG9uc2VIYW5kbGluZ0NvbmZpZywgc3RhdHVzKSB7XG4gICAgdmFyIHJlZ0V4cCA9IG5ldyBSZWdFeHAocmVzcG9uc2VIYW5kbGluZ0NvbmZpZy5jb2RlKVxuICAgIHJldHVybiByZWdFeHAudGVzdChzdGF0dXMudG9TdHJpbmcoMTApKVxuICB9XG5cbiAgLyoqXG4gICAqIEBwYXJhbSB7WE1MSHR0cFJlcXVlc3R9IHhoclxuICAgKiBAcmV0dXJuIHtIdG14UmVzcG9uc2VIYW5kbGluZ0NvbmZpZ31cbiAgICovXG4gIGZ1bmN0aW9uIHJlc29sdmVSZXNwb25zZUhhbmRsaW5nKHhocikge1xuICAgIGZvciAodmFyIGkgPSAwOyBpIDwgaHRteC5jb25maWcucmVzcG9uc2VIYW5kbGluZy5sZW5ndGg7IGkrKykge1xuICAgICAgLyoqIEB0eXBlIEh0bXhSZXNwb25zZUhhbmRsaW5nQ29uZmlnICovXG4gICAgICB2YXIgcmVzcG9uc2VIYW5kbGluZ0VsZW1lbnQgPSBodG14LmNvbmZpZy5yZXNwb25zZUhhbmRsaW5nW2ldXG4gICAgICBpZiAoY29kZU1hdGNoZXMocmVzcG9uc2VIYW5kbGluZ0VsZW1lbnQsIHhoci5zdGF0dXMpKSB7XG4gICAgICAgIHJldHVybiByZXNwb25zZUhhbmRsaW5nRWxlbWVudFxuICAgICAgfVxuICAgIH1cbiAgICAvLyBubyBtYXRjaGVzLCByZXR1cm4gbm8gc3dhcFxuICAgIHJldHVybiB7XG4gICAgICBzd2FwOiBmYWxzZVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBAcGFyYW0ge3N0cmluZ30gdGl0bGVcbiAgICovXG4gIGZ1bmN0aW9uIGhhbmRsZVRpdGxlKHRpdGxlKSB7XG4gICAgaWYgKHRpdGxlKSB7XG4gICAgICBjb25zdCB0aXRsZUVsdCA9IGZpbmQoJ3RpdGxlJylcbiAgICAgIGlmICh0aXRsZUVsdCkge1xuICAgICAgICB0aXRsZUVsdC5pbm5lckhUTUwgPSB0aXRsZVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgd2luZG93LmRvY3VtZW50LnRpdGxlID0gdGl0bGVcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICAvKipcbiAgICogQHBhcmFtIHtFbGVtZW50fSBlbHRcbiAgICogQHBhcmFtIHtIdG14UmVzcG9uc2VJbmZvfSByZXNwb25zZUluZm9cbiAgICovXG4gIGZ1bmN0aW9uIGhhbmRsZUFqYXhSZXNwb25zZShlbHQsIHJlc3BvbnNlSW5mbykge1xuICAgIGNvbnN0IHhociA9IHJlc3BvbnNlSW5mby54aHJcbiAgICBsZXQgdGFyZ2V0ID0gcmVzcG9uc2VJbmZvLnRhcmdldFxuICAgIGNvbnN0IGV0YyA9IHJlc3BvbnNlSW5mby5ldGNcbiAgICBjb25zdCByZXNwb25zZUluZm9TZWxlY3QgPSByZXNwb25zZUluZm8uc2VsZWN0XG5cbiAgICBpZiAoIXRyaWdnZXJFdmVudChlbHQsICdodG14OmJlZm9yZU9uTG9hZCcsIHJlc3BvbnNlSW5mbykpIHJldHVyblxuXG4gICAgaWYgKGhhc0hlYWRlcih4aHIsIC9IWC1UcmlnZ2VyOi9pKSkge1xuICAgICAgaGFuZGxlVHJpZ2dlckhlYWRlcih4aHIsICdIWC1UcmlnZ2VyJywgZWx0KVxuICAgIH1cblxuICAgIGlmIChoYXNIZWFkZXIoeGhyLCAvSFgtTG9jYXRpb246L2kpKSB7XG4gICAgICBzYXZlQ3VycmVudFBhZ2VUb0hpc3RvcnkoKVxuICAgICAgbGV0IHJlZGlyZWN0UGF0aCA9IHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtTG9jYXRpb24nKVxuICAgICAgLyoqIEB0eXBlIHtIdG14QWpheEhlbHBlckNvbnRleHQme3BhdGg6c3RyaW5nfX0gKi9cbiAgICAgIHZhciByZWRpcmVjdFN3YXBTcGVjXG4gICAgICBpZiAocmVkaXJlY3RQYXRoLmluZGV4T2YoJ3snKSA9PT0gMCkge1xuICAgICAgICByZWRpcmVjdFN3YXBTcGVjID0gcGFyc2VKU09OKHJlZGlyZWN0UGF0aClcbiAgICAgICAgLy8gd2hhdCdzIHRoZSBiZXN0IHdheSB0byB0aHJvdyBhbiBlcnJvciBpZiB0aGUgdXNlciBkaWRuJ3QgaW5jbHVkZSB0aGlzXG4gICAgICAgIHJlZGlyZWN0UGF0aCA9IHJlZGlyZWN0U3dhcFNwZWMucGF0aFxuICAgICAgICBkZWxldGUgcmVkaXJlY3RTd2FwU3BlYy5wYXRoXG4gICAgICB9XG4gICAgICBhamF4SGVscGVyKCdnZXQnLCByZWRpcmVjdFBhdGgsIHJlZGlyZWN0U3dhcFNwZWMpLnRoZW4oZnVuY3Rpb24oKSB7XG4gICAgICAgIHB1c2hVcmxJbnRvSGlzdG9yeShyZWRpcmVjdFBhdGgpXG4gICAgICB9KVxuICAgICAgcmV0dXJuXG4gICAgfVxuXG4gICAgY29uc3Qgc2hvdWxkUmVmcmVzaCA9IGhhc0hlYWRlcih4aHIsIC9IWC1SZWZyZXNoOi9pKSAmJiB4aHIuZ2V0UmVzcG9uc2VIZWFkZXIoJ0hYLVJlZnJlc2gnKSA9PT0gJ3RydWUnXG5cbiAgICBpZiAoaGFzSGVhZGVyKHhociwgL0hYLVJlZGlyZWN0Oi9pKSkge1xuICAgICAgcmVzcG9uc2VJbmZvLmtlZXBJbmRpY2F0b3JzID0gdHJ1ZVxuICAgICAgbG9jYXRpb24uaHJlZiA9IHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUmVkaXJlY3QnKVxuICAgICAgc2hvdWxkUmVmcmVzaCAmJiBsb2NhdGlvbi5yZWxvYWQoKVxuICAgICAgcmV0dXJuXG4gICAgfVxuXG4gICAgaWYgKHNob3VsZFJlZnJlc2gpIHtcbiAgICAgIHJlc3BvbnNlSW5mby5rZWVwSW5kaWNhdG9ycyA9IHRydWVcbiAgICAgIGxvY2F0aW9uLnJlbG9hZCgpXG4gICAgICByZXR1cm5cbiAgICB9XG5cbiAgICBpZiAoaGFzSGVhZGVyKHhociwgL0hYLVJldGFyZ2V0Oi9pKSkge1xuICAgICAgaWYgKHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUmV0YXJnZXQnKSA9PT0gJ3RoaXMnKSB7XG4gICAgICAgIHJlc3BvbnNlSW5mby50YXJnZXQgPSBlbHRcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHJlc3BvbnNlSW5mby50YXJnZXQgPSBhc0VsZW1lbnQocXVlcnlTZWxlY3RvckV4dChlbHQsIHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUmV0YXJnZXQnKSkpXG4gICAgICB9XG4gICAgfVxuXG4gICAgY29uc3QgaGlzdG9yeVVwZGF0ZSA9IGRldGVybWluZUhpc3RvcnlVcGRhdGVzKGVsdCwgcmVzcG9uc2VJbmZvKVxuXG4gICAgY29uc3QgcmVzcG9uc2VIYW5kbGluZyA9IHJlc29sdmVSZXNwb25zZUhhbmRsaW5nKHhocilcbiAgICBjb25zdCBzaG91bGRTd2FwID0gcmVzcG9uc2VIYW5kbGluZy5zd2FwXG4gICAgbGV0IGlzRXJyb3IgPSAhIXJlc3BvbnNlSGFuZGxpbmcuZXJyb3JcbiAgICBsZXQgaWdub3JlVGl0bGUgPSBodG14LmNvbmZpZy5pZ25vcmVUaXRsZSB8fCByZXNwb25zZUhhbmRsaW5nLmlnbm9yZVRpdGxlXG4gICAgbGV0IHNlbGVjdE92ZXJyaWRlID0gcmVzcG9uc2VIYW5kbGluZy5zZWxlY3RcbiAgICBpZiAocmVzcG9uc2VIYW5kbGluZy50YXJnZXQpIHtcbiAgICAgIHJlc3BvbnNlSW5mby50YXJnZXQgPSBhc0VsZW1lbnQocXVlcnlTZWxlY3RvckV4dChlbHQsIHJlc3BvbnNlSGFuZGxpbmcudGFyZ2V0KSlcbiAgICB9XG4gICAgdmFyIHN3YXBPdmVycmlkZSA9IGV0Yy5zd2FwT3ZlcnJpZGVcbiAgICBpZiAoc3dhcE92ZXJyaWRlID09IG51bGwgJiYgcmVzcG9uc2VIYW5kbGluZy5zd2FwT3ZlcnJpZGUpIHtcbiAgICAgIHN3YXBPdmVycmlkZSA9IHJlc3BvbnNlSGFuZGxpbmcuc3dhcE92ZXJyaWRlXG4gICAgfVxuXG4gICAgLy8gcmVzcG9uc2UgaGVhZGVycyBvdmVycmlkZSByZXNwb25zZSBoYW5kbGluZyBjb25maWdcbiAgICBpZiAoaGFzSGVhZGVyKHhociwgL0hYLVJldGFyZ2V0Oi9pKSkge1xuICAgICAgaWYgKHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUmV0YXJnZXQnKSA9PT0gJ3RoaXMnKSB7XG4gICAgICAgIHJlc3BvbnNlSW5mby50YXJnZXQgPSBlbHRcbiAgICAgIH0gZWxzZSB7XG4gICAgICAgIHJlc3BvbnNlSW5mby50YXJnZXQgPSBhc0VsZW1lbnQocXVlcnlTZWxlY3RvckV4dChlbHQsIHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUmV0YXJnZXQnKSkpXG4gICAgICB9XG4gICAgfVxuICAgIGlmIChoYXNIZWFkZXIoeGhyLCAvSFgtUmVzd2FwOi9pKSkge1xuICAgICAgc3dhcE92ZXJyaWRlID0geGhyLmdldFJlc3BvbnNlSGVhZGVyKCdIWC1SZXN3YXAnKVxuICAgIH1cblxuICAgIHZhciBzZXJ2ZXJSZXNwb25zZSA9IHhoci5yZXNwb25zZVxuICAgIC8qKiBAdHlwZSBIdG14QmVmb3JlU3dhcERldGFpbHMgKi9cbiAgICB2YXIgYmVmb3JlU3dhcERldGFpbHMgPSBtZXJnZU9iamVjdHMoe1xuICAgICAgc2hvdWxkU3dhcCxcbiAgICAgIHNlcnZlclJlc3BvbnNlLFxuICAgICAgaXNFcnJvcixcbiAgICAgIGlnbm9yZVRpdGxlLFxuICAgICAgc2VsZWN0T3ZlcnJpZGVcbiAgICB9LCByZXNwb25zZUluZm8pXG5cbiAgICBpZiAocmVzcG9uc2VIYW5kbGluZy5ldmVudCAmJiAhdHJpZ2dlckV2ZW50KHRhcmdldCwgcmVzcG9uc2VIYW5kbGluZy5ldmVudCwgYmVmb3JlU3dhcERldGFpbHMpKSByZXR1cm5cblxuICAgIGlmICghdHJpZ2dlckV2ZW50KHRhcmdldCwgJ2h0bXg6YmVmb3JlU3dhcCcsIGJlZm9yZVN3YXBEZXRhaWxzKSkgcmV0dXJuXG5cbiAgICB0YXJnZXQgPSBiZWZvcmVTd2FwRGV0YWlscy50YXJnZXQgLy8gYWxsb3cgcmUtdGFyZ2V0aW5nXG4gICAgc2VydmVyUmVzcG9uc2UgPSBiZWZvcmVTd2FwRGV0YWlscy5zZXJ2ZXJSZXNwb25zZSAvLyBhbGxvdyB1cGRhdGluZyBjb250ZW50XG4gICAgaXNFcnJvciA9IGJlZm9yZVN3YXBEZXRhaWxzLmlzRXJyb3IgLy8gYWxsb3cgdXBkYXRpbmcgZXJyb3JcbiAgICBpZ25vcmVUaXRsZSA9IGJlZm9yZVN3YXBEZXRhaWxzLmlnbm9yZVRpdGxlIC8vIGFsbG93IHVwZGF0aW5nIGlnbm9yaW5nIHRpdGxlXG4gICAgc2VsZWN0T3ZlcnJpZGUgPSBiZWZvcmVTd2FwRGV0YWlscy5zZWxlY3RPdmVycmlkZSAvLyBhbGxvdyB1cGRhdGluZyBzZWxlY3Qgb3ZlcnJpZGVcblxuICAgIHJlc3BvbnNlSW5mby50YXJnZXQgPSB0YXJnZXQgLy8gTWFrZSB1cGRhdGVkIHRhcmdldCBhdmFpbGFibGUgdG8gcmVzcG9uc2UgZXZlbnRzXG4gICAgcmVzcG9uc2VJbmZvLmZhaWxlZCA9IGlzRXJyb3IgLy8gTWFrZSBmYWlsZWQgcHJvcGVydHkgYXZhaWxhYmxlIHRvIHJlc3BvbnNlIGV2ZW50c1xuICAgIHJlc3BvbnNlSW5mby5zdWNjZXNzZnVsID0gIWlzRXJyb3IgLy8gTWFrZSBzdWNjZXNzZnVsIHByb3BlcnR5IGF2YWlsYWJsZSB0byByZXNwb25zZSBldmVudHNcblxuICAgIGlmIChiZWZvcmVTd2FwRGV0YWlscy5zaG91bGRTd2FwKSB7XG4gICAgICBpZiAoeGhyLnN0YXR1cyA9PT0gMjg2KSB7XG4gICAgICAgIGNhbmNlbFBvbGxpbmcoZWx0KVxuICAgICAgfVxuXG4gICAgICB3aXRoRXh0ZW5zaW9ucyhlbHQsIGZ1bmN0aW9uKGV4dGVuc2lvbikge1xuICAgICAgICBzZXJ2ZXJSZXNwb25zZSA9IGV4dGVuc2lvbi50cmFuc2Zvcm1SZXNwb25zZShzZXJ2ZXJSZXNwb25zZSwgeGhyLCBlbHQpXG4gICAgICB9KVxuXG4gICAgICAvLyBTYXZlIGN1cnJlbnQgcGFnZSBpZiB0aGVyZSB3aWxsIGJlIGEgaGlzdG9yeSB1cGRhdGVcbiAgICAgIGlmIChoaXN0b3J5VXBkYXRlLnR5cGUpIHtcbiAgICAgICAgc2F2ZUN1cnJlbnRQYWdlVG9IaXN0b3J5KClcbiAgICAgIH1cblxuICAgICAgaWYgKGhhc0hlYWRlcih4aHIsIC9IWC1SZXN3YXA6L2kpKSB7XG4gICAgICAgIHN3YXBPdmVycmlkZSA9IHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUmVzd2FwJylcbiAgICAgIH1cbiAgICAgIHZhciBzd2FwU3BlYyA9IGdldFN3YXBTcGVjaWZpY2F0aW9uKGVsdCwgc3dhcE92ZXJyaWRlKVxuXG4gICAgICBpZiAoIXN3YXBTcGVjLmhhc093blByb3BlcnR5KCdpZ25vcmVUaXRsZScpKSB7XG4gICAgICAgIHN3YXBTcGVjLmlnbm9yZVRpdGxlID0gaWdub3JlVGl0bGVcbiAgICAgIH1cblxuICAgICAgdGFyZ2V0LmNsYXNzTGlzdC5hZGQoaHRteC5jb25maWcuc3dhcHBpbmdDbGFzcylcblxuICAgICAgLy8gb3B0aW9uYWwgdHJhbnNpdGlvbiBBUEkgcHJvbWlzZSBjYWxsYmFja3NcbiAgICAgIGxldCBzZXR0bGVSZXNvbHZlID0gbnVsbFxuICAgICAgbGV0IHNldHRsZVJlamVjdCA9IG51bGxcblxuICAgICAgaWYgKHJlc3BvbnNlSW5mb1NlbGVjdCkge1xuICAgICAgICBzZWxlY3RPdmVycmlkZSA9IHJlc3BvbnNlSW5mb1NlbGVjdFxuICAgICAgfVxuXG4gICAgICBpZiAoaGFzSGVhZGVyKHhociwgL0hYLVJlc2VsZWN0Oi9pKSkge1xuICAgICAgICBzZWxlY3RPdmVycmlkZSA9IHhoci5nZXRSZXNwb25zZUhlYWRlcignSFgtUmVzZWxlY3QnKVxuICAgICAgfVxuXG4gICAgICBjb25zdCBzZWxlY3RPT0IgPSBnZXRDbG9zZXN0QXR0cmlidXRlVmFsdWUoZWx0LCAnaHgtc2VsZWN0LW9vYicpXG4gICAgICBjb25zdCBzZWxlY3QgPSBnZXRDbG9zZXN0QXR0cmlidXRlVmFsdWUoZWx0LCAnaHgtc2VsZWN0JylcblxuICAgICAgbGV0IGRvU3dhcCA9IGZ1bmN0aW9uKCkge1xuICAgICAgICB0cnkge1xuICAgICAgICAgIC8vIGlmIHdlIG5lZWQgdG8gc2F2ZSBoaXN0b3J5LCBkbyBzbywgYmVmb3JlIHN3YXBwaW5nIHNvIHRoYXQgcmVsYXRpdmUgcmVzb3VyY2VzIGhhdmUgdGhlIGNvcnJlY3QgYmFzZSBVUkxcbiAgICAgICAgICBpZiAoaGlzdG9yeVVwZGF0ZS50eXBlKSB7XG4gICAgICAgICAgICB0cmlnZ2VyRXZlbnQoZ2V0RG9jdW1lbnQoKS5ib2R5LCAnaHRteDpiZWZvcmVIaXN0b3J5VXBkYXRlJywgbWVyZ2VPYmplY3RzKHsgaGlzdG9yeTogaGlzdG9yeVVwZGF0ZSB9LCByZXNwb25zZUluZm8pKVxuICAgICAgICAgICAgaWYgKGhpc3RvcnlVcGRhdGUudHlwZSA9PT0gJ3B1c2gnKSB7XG4gICAgICAgICAgICAgIHB1c2hVcmxJbnRvSGlzdG9yeShoaXN0b3J5VXBkYXRlLnBhdGgpXG4gICAgICAgICAgICAgIHRyaWdnZXJFdmVudChnZXREb2N1bWVudCgpLmJvZHksICdodG14OnB1c2hlZEludG9IaXN0b3J5JywgeyBwYXRoOiBoaXN0b3J5VXBkYXRlLnBhdGggfSlcbiAgICAgICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgICAgIHJlcGxhY2VVcmxJbkhpc3RvcnkoaGlzdG9yeVVwZGF0ZS5wYXRoKVxuICAgICAgICAgICAgICB0cmlnZ2VyRXZlbnQoZ2V0RG9jdW1lbnQoKS5ib2R5LCAnaHRteDpyZXBsYWNlZEluSGlzdG9yeScsIHsgcGF0aDogaGlzdG9yeVVwZGF0ZS5wYXRoIH0pXG4gICAgICAgICAgICB9XG4gICAgICAgICAgfVxuXG4gICAgICAgICAgc3dhcCh0YXJnZXQsIHNlcnZlclJlc3BvbnNlLCBzd2FwU3BlYywge1xuICAgICAgICAgICAgc2VsZWN0OiBzZWxlY3RPdmVycmlkZSB8fCBzZWxlY3QsXG4gICAgICAgICAgICBzZWxlY3RPT0IsXG4gICAgICAgICAgICBldmVudEluZm86IHJlc3BvbnNlSW5mbyxcbiAgICAgICAgICAgIGFuY2hvcjogcmVzcG9uc2VJbmZvLnBhdGhJbmZvLmFuY2hvcixcbiAgICAgICAgICAgIGNvbnRleHRFbGVtZW50OiBlbHQsXG4gICAgICAgICAgICBhZnRlclN3YXBDYWxsYmFjazogZnVuY3Rpb24oKSB7XG4gICAgICAgICAgICAgIGlmIChoYXNIZWFkZXIoeGhyLCAvSFgtVHJpZ2dlci1BZnRlci1Td2FwOi9pKSkge1xuICAgICAgICAgICAgICAgIGxldCBmaW5hbEVsdCA9IGVsdFxuICAgICAgICAgICAgICAgIGlmICghYm9keUNvbnRhaW5zKGVsdCkpIHtcbiAgICAgICAgICAgICAgICAgIGZpbmFsRWx0ID0gZ2V0RG9jdW1lbnQoKS5ib2R5XG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgICAgIGhhbmRsZVRyaWdnZXJIZWFkZXIoeGhyLCAnSFgtVHJpZ2dlci1BZnRlci1Td2FwJywgZmluYWxFbHQpXG4gICAgICAgICAgICAgIH1cbiAgICAgICAgICAgIH0sXG4gICAgICAgICAgICBhZnRlclNldHRsZUNhbGxiYWNrOiBmdW5jdGlvbigpIHtcbiAgICAgICAgICAgICAgaWYgKGhhc0hlYWRlcih4aHIsIC9IWC1UcmlnZ2VyLUFmdGVyLVNldHRsZTovaSkpIHtcbiAgICAgICAgICAgICAgICBsZXQgZmluYWxFbHQgPSBlbHRcbiAgICAgICAgICAgICAgICBpZiAoIWJvZHlDb250YWlucyhlbHQpKSB7XG4gICAgICAgICAgICAgICAgICBmaW5hbEVsdCA9IGdldERvY3VtZW50KCkuYm9keVxuICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgICAgICBoYW5kbGVUcmlnZ2VySGVhZGVyKHhociwgJ0hYLVRyaWdnZXItQWZ0ZXItU2V0dGxlJywgZmluYWxFbHQpXG4gICAgICAgICAgICAgIH1cbiAgICAgICAgICAgICAgbWF5YmVDYWxsKHNldHRsZVJlc29sdmUpXG4gICAgICAgICAgICB9XG4gICAgICAgICAgfSlcbiAgICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICAgIHRyaWdnZXJFcnJvckV2ZW50KGVsdCwgJ2h0bXg6c3dhcEVycm9yJywgcmVzcG9uc2VJbmZvKVxuICAgICAgICAgIG1heWJlQ2FsbChzZXR0bGVSZWplY3QpXG4gICAgICAgICAgdGhyb3cgZVxuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIGxldCBzaG91bGRUcmFuc2l0aW9uID0gaHRteC5jb25maWcuZ2xvYmFsVmlld1RyYW5zaXRpb25zXG4gICAgICBpZiAoc3dhcFNwZWMuaGFzT3duUHJvcGVydHkoJ3RyYW5zaXRpb24nKSkge1xuICAgICAgICBzaG91bGRUcmFuc2l0aW9uID0gc3dhcFNwZWMudHJhbnNpdGlvblxuICAgICAgfVxuXG4gICAgICBpZiAoc2hvdWxkVHJhbnNpdGlvbiAmJlxuICAgICAgICAgICAgICB0cmlnZ2VyRXZlbnQoZWx0LCAnaHRteDpiZWZvcmVUcmFuc2l0aW9uJywgcmVzcG9uc2VJbmZvKSAmJlxuICAgICAgICAgICAgICB0eXBlb2YgUHJvbWlzZSAhPT0gJ3VuZGVmaW5lZCcgJiZcbiAgICAgICAgICAgICAgLy8gQHRzLWlnbm9yZSBleHBlcmltZW50YWwgZmVhdHVyZSBhdG1cbiAgICAgICAgICAgICAgZG9jdW1lbnQuc3RhcnRWaWV3VHJhbnNpdGlvbikge1xuICAgICAgICBjb25zdCBzZXR0bGVQcm9taXNlID0gbmV3IFByb21pc2UoZnVuY3Rpb24oX3Jlc29sdmUsIF9yZWplY3QpIHtcbiAgICAgICAgICBzZXR0bGVSZXNvbHZlID0gX3Jlc29sdmVcbiAgICAgICAgICBzZXR0bGVSZWplY3QgPSBfcmVqZWN0XG4gICAgICAgIH0pXG4gICAgICAgIC8vIHdyYXAgdGhlIG9yaWdpbmFsIGRvU3dhcCgpIGluIGEgY2FsbCB0byBzdGFydFZpZXdUcmFuc2l0aW9uKClcbiAgICAgICAgY29uc3QgaW5uZXJEb1N3YXAgPSBkb1N3YXBcbiAgICAgICAgZG9Td2FwID0gZnVuY3Rpb24oKSB7XG4gICAgICAgICAgLy8gQHRzLWlnbm9yZSBleHBlcmltZW50YWwgZmVhdHVyZSBhdG1cbiAgICAgICAgICBkb2N1bWVudC5zdGFydFZpZXdUcmFuc2l0aW9uKGZ1bmN0aW9uKCkge1xuICAgICAgICAgICAgaW5uZXJEb1N3YXAoKVxuICAgICAgICAgICAgcmV0dXJuIHNldHRsZVByb21pc2VcbiAgICAgICAgICB9KVxuICAgICAgICB9XG4gICAgICB9XG5cbiAgICAgIGlmIChzd2FwU3BlYy5zd2FwRGVsYXkgPiAwKSB7XG4gICAgICAgIGdldFdpbmRvdygpLnNldFRpbWVvdXQoZG9Td2FwLCBzd2FwU3BlYy5zd2FwRGVsYXkpXG4gICAgICB9IGVsc2Uge1xuICAgICAgICBkb1N3YXAoKVxuICAgICAgfVxuICAgIH1cbiAgICBpZiAoaXNFcnJvcikge1xuICAgICAgdHJpZ2dlckVycm9yRXZlbnQoZWx0LCAnaHRteDpyZXNwb25zZUVycm9yJywgbWVyZ2VPYmplY3RzKHsgZXJyb3I6ICdSZXNwb25zZSBTdGF0dXMgRXJyb3IgQ29kZSAnICsgeGhyLnN0YXR1cyArICcgZnJvbSAnICsgcmVzcG9uc2VJbmZvLnBhdGhJbmZvLnJlcXVlc3RQYXRoIH0sIHJlc3BvbnNlSW5mbykpXG4gICAgfVxuICB9XG5cbiAgLy89ID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT1cbiAgLy8gRXh0ZW5zaW9ucyBBUElcbiAgLy89ID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT1cblxuICAvKiogQHR5cGUge09iamVjdDxzdHJpbmcsIEh0bXhFeHRlbnNpb24+fSAqL1xuICBjb25zdCBleHRlbnNpb25zID0ge31cblxuICAvKipcbiAgICogZXh0ZW5zaW9uQmFzZSBkZWZpbmVzIHRoZSBkZWZhdWx0IGZ1bmN0aW9ucyBmb3IgYWxsIGV4dGVuc2lvbnMuXG4gICAqIEByZXR1cm5zIHtIdG14RXh0ZW5zaW9ufVxuICAgKi9cbiAgZnVuY3Rpb24gZXh0ZW5zaW9uQmFzZSgpIHtcbiAgICByZXR1cm4ge1xuICAgICAgaW5pdDogZnVuY3Rpb24oYXBpKSB7IHJldHVybiBudWxsIH0sXG4gICAgICBnZXRTZWxlY3RvcnM6IGZ1bmN0aW9uKCkgeyByZXR1cm4gbnVsbCB9LFxuICAgICAgb25FdmVudDogZnVuY3Rpb24obmFtZSwgZXZ0KSB7IHJldHVybiB0cnVlIH0sXG4gICAgICB0cmFuc2Zvcm1SZXNwb25zZTogZnVuY3Rpb24odGV4dCwgeGhyLCBlbHQpIHsgcmV0dXJuIHRleHQgfSxcbiAgICAgIGlzSW5saW5lU3dhcDogZnVuY3Rpb24oc3dhcFN0eWxlKSB7IHJldHVybiBmYWxzZSB9LFxuICAgICAgaGFuZGxlU3dhcDogZnVuY3Rpb24oc3dhcFN0eWxlLCB0YXJnZXQsIGZyYWdtZW50LCBzZXR0bGVJbmZvKSB7IHJldHVybiBmYWxzZSB9LFxuICAgICAgZW5jb2RlUGFyYW1ldGVyczogZnVuY3Rpb24oeGhyLCBwYXJhbWV0ZXJzLCBlbHQpIHsgcmV0dXJuIG51bGwgfVxuICAgIH1cbiAgfVxuXG4gIC8qKlxuICAgKiBkZWZpbmVFeHRlbnNpb24gaW5pdGlhbGl6ZXMgdGhlIGV4dGVuc2lvbiBhbmQgYWRkcyBpdCB0byB0aGUgaHRteCByZWdpc3RyeVxuICAgKlxuICAgKiBAc2VlIGh0dHBzOi8vaHRteC5vcmcvYXBpLyNkZWZpbmVFeHRlbnNpb25cbiAgICpcbiAgICogQHBhcmFtIHtzdHJpbmd9IG5hbWUgdGhlIGV4dGVuc2lvbiBuYW1lXG4gICAqIEBwYXJhbSB7SHRteEV4dGVuc2lvbn0gZXh0ZW5zaW9uIHRoZSBleHRlbnNpb24gZGVmaW5pdGlvblxuICAgKi9cbiAgZnVuY3Rpb24gZGVmaW5lRXh0ZW5zaW9uKG5hbWUsIGV4dGVuc2lvbikge1xuICAgIGlmIChleHRlbnNpb24uaW5pdCkge1xuICAgICAgZXh0ZW5zaW9uLmluaXQoaW50ZXJuYWxBUEkpXG4gICAgfVxuICAgIGV4dGVuc2lvbnNbbmFtZV0gPSBtZXJnZU9iamVjdHMoZXh0ZW5zaW9uQmFzZSgpLCBleHRlbnNpb24pXG4gIH1cblxuICAvKipcbiAgICogcmVtb3ZlRXh0ZW5zaW9uIHJlbW92ZXMgYW4gZXh0ZW5zaW9uIGZyb20gdGhlIGh0bXggcmVnaXN0cnlcbiAgICpcbiAgICogQHNlZSBodHRwczovL2h0bXgub3JnL2FwaS8jcmVtb3ZlRXh0ZW5zaW9uXG4gICAqXG4gICAqIEBwYXJhbSB7c3RyaW5nfSBuYW1lXG4gICAqL1xuICBmdW5jdGlvbiByZW1vdmVFeHRlbnNpb24obmFtZSkge1xuICAgIGRlbGV0ZSBleHRlbnNpb25zW25hbWVdXG4gIH1cblxuICAvKipcbiAgICogZ2V0RXh0ZW5zaW9ucyBzZWFyY2hlcyB1cCB0aGUgRE9NIHRyZWUgdG8gcmV0dXJuIGFsbCBleHRlbnNpb25zIHRoYXQgY2FuIGJlIGFwcGxpZWQgdG8gYSBnaXZlbiBlbGVtZW50XG4gICAqXG4gICAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gICAqIEBwYXJhbSB7SHRteEV4dGVuc2lvbltdPX0gZXh0ZW5zaW9uc1RvUmV0dXJuXG4gICAqIEBwYXJhbSB7c3RyaW5nW109fSBleHRlbnNpb25zVG9JZ25vcmVcbiAgICogQHJldHVybnMge0h0bXhFeHRlbnNpb25bXX1cbiAgICovXG4gIGZ1bmN0aW9uIGdldEV4dGVuc2lvbnMoZWx0LCBleHRlbnNpb25zVG9SZXR1cm4sIGV4dGVuc2lvbnNUb0lnbm9yZSkge1xuICAgIGlmIChleHRlbnNpb25zVG9SZXR1cm4gPT0gdW5kZWZpbmVkKSB7XG4gICAgICBleHRlbnNpb25zVG9SZXR1cm4gPSBbXVxuICAgIH1cbiAgICBpZiAoZWx0ID09IHVuZGVmaW5lZCkge1xuICAgICAgcmV0dXJuIGV4dGVuc2lvbnNUb1JldHVyblxuICAgIH1cbiAgICBpZiAoZXh0ZW5zaW9uc1RvSWdub3JlID09IHVuZGVmaW5lZCkge1xuICAgICAgZXh0ZW5zaW9uc1RvSWdub3JlID0gW11cbiAgICB9XG4gICAgY29uc3QgZXh0ZW5zaW9uc0ZvckVsZW1lbnQgPSBnZXRBdHRyaWJ1dGVWYWx1ZShlbHQsICdoeC1leHQnKVxuICAgIGlmIChleHRlbnNpb25zRm9yRWxlbWVudCkge1xuICAgICAgZm9yRWFjaChleHRlbnNpb25zRm9yRWxlbWVudC5zcGxpdCgnLCcpLCBmdW5jdGlvbihleHRlbnNpb25OYW1lKSB7XG4gICAgICAgIGV4dGVuc2lvbk5hbWUgPSBleHRlbnNpb25OYW1lLnJlcGxhY2UoLyAvZywgJycpXG4gICAgICAgIGlmIChleHRlbnNpb25OYW1lLnNsaWNlKDAsIDcpID09ICdpZ25vcmU6Jykge1xuICAgICAgICAgIGV4dGVuc2lvbnNUb0lnbm9yZS5wdXNoKGV4dGVuc2lvbk5hbWUuc2xpY2UoNykpXG4gICAgICAgICAgcmV0dXJuXG4gICAgICAgIH1cbiAgICAgICAgaWYgKGV4dGVuc2lvbnNUb0lnbm9yZS5pbmRleE9mKGV4dGVuc2lvbk5hbWUpIDwgMCkge1xuICAgICAgICAgIGNvbnN0IGV4dGVuc2lvbiA9IGV4dGVuc2lvbnNbZXh0ZW5zaW9uTmFtZV1cbiAgICAgICAgICBpZiAoZXh0ZW5zaW9uICYmIGV4dGVuc2lvbnNUb1JldHVybi5pbmRleE9mKGV4dGVuc2lvbikgPCAwKSB7XG4gICAgICAgICAgICBleHRlbnNpb25zVG9SZXR1cm4ucHVzaChleHRlbnNpb24pXG4gICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICB9KVxuICAgIH1cbiAgICByZXR1cm4gZ2V0RXh0ZW5zaW9ucyhhc0VsZW1lbnQocGFyZW50RWx0KGVsdCkpLCBleHRlbnNpb25zVG9SZXR1cm4sIGV4dGVuc2lvbnNUb0lnbm9yZSlcbiAgfVxuXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG4gIC8vIEluaXRpYWxpemF0aW9uXG4gIC8vPSA9PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09XG4gIHZhciBpc1JlYWR5ID0gZmFsc2VcbiAgZ2V0RG9jdW1lbnQoKS5hZGRFdmVudExpc3RlbmVyKCdET01Db250ZW50TG9hZGVkJywgZnVuY3Rpb24oKSB7XG4gICAgaXNSZWFkeSA9IHRydWVcbiAgfSlcblxuICAvKipcbiAgICogRXhlY3V0ZSBhIGZ1bmN0aW9uIG5vdyBpZiBET01Db250ZW50TG9hZGVkIGhhcyBmaXJlZCwgb3RoZXJ3aXNlIGxpc3RlbiBmb3IgaXQuXG4gICAqXG4gICAqIFRoaXMgZnVuY3Rpb24gdXNlcyBpc1JlYWR5IGJlY2F1c2UgdGhlcmUgaXMgbm8gcmVsaWFibGUgd2F5IHRvIGFzayB0aGUgYnJvd3NlciB3aGV0aGVyXG4gICAqIHRoZSBET01Db250ZW50TG9hZGVkIGV2ZW50IGhhcyBhbHJlYWR5IGJlZW4gZmlyZWQ7IHRoZXJlJ3MgYSBnYXAgYmV0d2VlbiBET01Db250ZW50TG9hZGVkXG4gICAqIGZpcmluZyBhbmQgcmVhZHlzdGF0ZT1jb21wbGV0ZS5cbiAgICovXG4gIGZ1bmN0aW9uIHJlYWR5KGZuKSB7XG4gICAgLy8gQ2hlY2tpbmcgcmVhZHlTdGF0ZSBoZXJlIGlzIGEgZmFpbHNhZmUgaW4gY2FzZSB0aGUgaHRteCBzY3JpcHQgdGFnIGVudGVyZWQgdGhlIERPTSBieVxuICAgIC8vIHNvbWUgbWVhbnMgb3RoZXIgdGhhbiB0aGUgaW5pdGlhbCBwYWdlIGxvYWQuXG4gICAgaWYgKGlzUmVhZHkgfHwgZ2V0RG9jdW1lbnQoKS5yZWFkeVN0YXRlID09PSAnY29tcGxldGUnKSB7XG4gICAgICBmbigpXG4gICAgfSBlbHNlIHtcbiAgICAgIGdldERvY3VtZW50KCkuYWRkRXZlbnRMaXN0ZW5lcignRE9NQ29udGVudExvYWRlZCcsIGZuKVxuICAgIH1cbiAgfVxuXG4gIGZ1bmN0aW9uIGluc2VydEluZGljYXRvclN0eWxlcygpIHtcbiAgICBpZiAoaHRteC5jb25maWcuaW5jbHVkZUluZGljYXRvclN0eWxlcyAhPT0gZmFsc2UpIHtcbiAgICAgIGNvbnN0IG5vbmNlQXR0cmlidXRlID0gaHRteC5jb25maWcuaW5saW5lU3R5bGVOb25jZSA/IGAgbm9uY2U9XCIke2h0bXguY29uZmlnLmlubGluZVN0eWxlTm9uY2V9XCJgIDogJydcbiAgICAgIGdldERvY3VtZW50KCkuaGVhZC5pbnNlcnRBZGphY2VudEhUTUwoJ2JlZm9yZWVuZCcsXG4gICAgICAgICc8c3R5bGUnICsgbm9uY2VBdHRyaWJ1dGUgKyAnPlxcXG4gICAgICAuJyArIGh0bXguY29uZmlnLmluZGljYXRvckNsYXNzICsgJ3tvcGFjaXR5OjB9XFxcbiAgICAgIC4nICsgaHRteC5jb25maWcucmVxdWVzdENsYXNzICsgJyAuJyArIGh0bXguY29uZmlnLmluZGljYXRvckNsYXNzICsgJ3tvcGFjaXR5OjE7IHRyYW5zaXRpb246IG9wYWNpdHkgMjAwbXMgZWFzZS1pbjt9XFxcbiAgICAgIC4nICsgaHRteC5jb25maWcucmVxdWVzdENsYXNzICsgJy4nICsgaHRteC5jb25maWcuaW5kaWNhdG9yQ2xhc3MgKyAne29wYWNpdHk6MTsgdHJhbnNpdGlvbjogb3BhY2l0eSAyMDBtcyBlYXNlLWluO31cXFxuICAgICAgPC9zdHlsZT4nKVxuICAgIH1cbiAgfVxuXG4gIGZ1bmN0aW9uIGdldE1ldGFDb25maWcoKSB7XG4gICAgLyoqIEB0eXBlIEhUTUxNZXRhRWxlbWVudCAqL1xuICAgIGNvbnN0IGVsZW1lbnQgPSBnZXREb2N1bWVudCgpLnF1ZXJ5U2VsZWN0b3IoJ21ldGFbbmFtZT1cImh0bXgtY29uZmlnXCJdJylcbiAgICBpZiAoZWxlbWVudCkge1xuICAgICAgcmV0dXJuIHBhcnNlSlNPTihlbGVtZW50LmNvbnRlbnQpXG4gICAgfSBlbHNlIHtcbiAgICAgIHJldHVybiBudWxsXG4gICAgfVxuICB9XG5cbiAgZnVuY3Rpb24gbWVyZ2VNZXRhQ29uZmlnKCkge1xuICAgIGNvbnN0IG1ldGFDb25maWcgPSBnZXRNZXRhQ29uZmlnKClcbiAgICBpZiAobWV0YUNvbmZpZykge1xuICAgICAgaHRteC5jb25maWcgPSBtZXJnZU9iamVjdHMoaHRteC5jb25maWcsIG1ldGFDb25maWcpXG4gICAgfVxuICB9XG5cbiAgLy8gaW5pdGlhbGl6ZSB0aGUgZG9jdW1lbnRcbiAgcmVhZHkoZnVuY3Rpb24oKSB7XG4gICAgbWVyZ2VNZXRhQ29uZmlnKClcbiAgICBpbnNlcnRJbmRpY2F0b3JTdHlsZXMoKVxuICAgIGxldCBib2R5ID0gZ2V0RG9jdW1lbnQoKS5ib2R5XG4gICAgcHJvY2Vzc05vZGUoYm9keSlcbiAgICBjb25zdCByZXN0b3JlZEVsdHMgPSBnZXREb2N1bWVudCgpLnF1ZXJ5U2VsZWN0b3JBbGwoXG4gICAgICBcIltoeC10cmlnZ2VyPSdyZXN0b3JlZCddLFtkYXRhLWh4LXRyaWdnZXI9J3Jlc3RvcmVkJ11cIlxuICAgIClcbiAgICBib2R5LmFkZEV2ZW50TGlzdGVuZXIoJ2h0bXg6YWJvcnQnLCBmdW5jdGlvbihldnQpIHtcbiAgICAgIGNvbnN0IHRhcmdldCA9IGV2dC50YXJnZXRcbiAgICAgIGNvbnN0IGludGVybmFsRGF0YSA9IGdldEludGVybmFsRGF0YSh0YXJnZXQpXG4gICAgICBpZiAoaW50ZXJuYWxEYXRhICYmIGludGVybmFsRGF0YS54aHIpIHtcbiAgICAgICAgaW50ZXJuYWxEYXRhLnhoci5hYm9ydCgpXG4gICAgICB9XG4gICAgfSlcbiAgICAvKiogQHR5cGUgeyhldjogUG9wU3RhdGVFdmVudCkgPT4gYW55fSAqL1xuICAgIGNvbnN0IG9yaWdpbmFsUG9wc3RhdGUgPSB3aW5kb3cub25wb3BzdGF0ZSA/IHdpbmRvdy5vbnBvcHN0YXRlLmJpbmQod2luZG93KSA6IG51bGxcbiAgICAvKiogQHR5cGUgeyhldjogUG9wU3RhdGVFdmVudCkgPT4gYW55fSAqL1xuICAgIHdpbmRvdy5vbnBvcHN0YXRlID0gZnVuY3Rpb24oZXZlbnQpIHtcbiAgICAgIGlmIChldmVudC5zdGF0ZSAmJiBldmVudC5zdGF0ZS5odG14KSB7XG4gICAgICAgIHJlc3RvcmVIaXN0b3J5KClcbiAgICAgICAgZm9yRWFjaChyZXN0b3JlZEVsdHMsIGZ1bmN0aW9uKGVsdCkge1xuICAgICAgICAgIHRyaWdnZXJFdmVudChlbHQsICdodG14OnJlc3RvcmVkJywge1xuICAgICAgICAgICAgZG9jdW1lbnQ6IGdldERvY3VtZW50KCksXG4gICAgICAgICAgICB0cmlnZ2VyRXZlbnRcbiAgICAgICAgICB9KVxuICAgICAgICB9KVxuICAgICAgfSBlbHNlIHtcbiAgICAgICAgaWYgKG9yaWdpbmFsUG9wc3RhdGUpIHtcbiAgICAgICAgICBvcmlnaW5hbFBvcHN0YXRlKGV2ZW50KVxuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICAgIGdldFdpbmRvdygpLnNldFRpbWVvdXQoZnVuY3Rpb24oKSB7XG4gICAgICB0cmlnZ2VyRXZlbnQoYm9keSwgJ2h0bXg6bG9hZCcsIHt9KSAvLyBnaXZlIHJlYWR5IGhhbmRsZXJzIGEgY2hhbmNlIHRvIGxvYWQgdXAgYmVmb3JlIGZpcmluZyB0aGlzIGV2ZW50XG4gICAgICBib2R5ID0gbnVsbCAvLyBraWxsIHJlZmVyZW5jZSBmb3IgZ2NcbiAgICB9LCAwKVxuICB9KVxuXG4gIHJldHVybiBodG14XG59KSgpXG5cbi8qKiBAdHlwZWRlZiB7J2dldCd8J2hlYWQnfCdwb3N0J3wncHV0J3wnZGVsZXRlJ3wnY29ubmVjdCd8J29wdGlvbnMnfCd0cmFjZSd8J3BhdGNoJ30gSHR0cFZlcmIgKi9cblxuLyoqXG4gKiBAdHlwZWRlZiB7T2JqZWN0fSBTd2FwT3B0aW9uc1xuICogQHByb3BlcnR5IHtzdHJpbmd9IFtzZWxlY3RdXG4gKiBAcHJvcGVydHkge3N0cmluZ30gW3NlbGVjdE9PQl1cbiAqIEBwcm9wZXJ0eSB7Kn0gW2V2ZW50SW5mb11cbiAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBbYW5jaG9yXVxuICogQHByb3BlcnR5IHtFbGVtZW50fSBbY29udGV4dEVsZW1lbnRdXG4gKiBAcHJvcGVydHkge3N3YXBDYWxsYmFja30gW2FmdGVyU3dhcENhbGxiYWNrXVxuICogQHByb3BlcnR5IHtzd2FwQ2FsbGJhY2t9IFthZnRlclNldHRsZUNhbGxiYWNrXVxuICovXG5cbi8qKlxuICogQGNhbGxiYWNrIHN3YXBDYWxsYmFja1xuICovXG5cbi8qKlxuICogQHR5cGVkZWYgeydpbm5lckhUTUwnIHwgJ291dGVySFRNTCcgfCAnYmVmb3JlYmVnaW4nIHwgJ2FmdGVyYmVnaW4nIHwgJ2JlZm9yZWVuZCcgfCAnYWZ0ZXJlbmQnIHwgJ2RlbGV0ZScgfCAnbm9uZScgfCBzdHJpbmd9IEh0bXhTd2FwU3R5bGVcbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIEh0bXhTd2FwU3BlY2lmaWNhdGlvblxuICogQHByb3BlcnR5IHtIdG14U3dhcFN0eWxlfSBzd2FwU3R5bGVcbiAqIEBwcm9wZXJ0eSB7bnVtYmVyfSBzd2FwRGVsYXlcbiAqIEBwcm9wZXJ0eSB7bnVtYmVyfSBzZXR0bGVEZWxheVxuICogQHByb3BlcnR5IHtib29sZWFufSBbdHJhbnNpdGlvbl1cbiAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gW2lnbm9yZVRpdGxlXVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtoZWFkXVxuICogQHByb3BlcnR5IHsndG9wJyB8ICdib3R0b20nfSBbc2Nyb2xsXVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtzY3JvbGxUYXJnZXRdXG4gKiBAcHJvcGVydHkge3N0cmluZ30gW3Nob3ddXG4gKiBAcHJvcGVydHkge3N0cmluZ30gW3Nob3dUYXJnZXRdXG4gKiBAcHJvcGVydHkge2Jvb2xlYW59IFtmb2N1c1Njcm9sbF1cbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIHsoKHRoaXM6Tm9kZSwgZXZ0OkV2ZW50KSA9PiBib29sZWFuKSAmIHtzb3VyY2U6IHN0cmluZ319IENvbmRpdGlvbmFsRnVuY3Rpb25cbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIHtPYmplY3R9IEh0bXhUcmlnZ2VyU3BlY2lmaWNhdGlvblxuICogQHByb3BlcnR5IHtzdHJpbmd9IHRyaWdnZXJcbiAqIEBwcm9wZXJ0eSB7bnVtYmVyfSBbcG9sbEludGVydmFsXVxuICogQHByb3BlcnR5IHtDb25kaXRpb25hbEZ1bmN0aW9ufSBbZXZlbnRGaWx0ZXJdXG4gKiBAcHJvcGVydHkge2Jvb2xlYW59IFtjaGFuZ2VkXVxuICogQHByb3BlcnR5IHtib29sZWFufSBbb25jZV1cbiAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gW2NvbnN1bWVdXG4gKiBAcHJvcGVydHkge251bWJlcn0gW2RlbGF5XVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtmcm9tXVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFt0YXJnZXRdXG4gKiBAcHJvcGVydHkge251bWJlcn0gW3Rocm90dGxlXVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtxdWV1ZV1cbiAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBbcm9vdF1cbiAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBbdGhyZXNob2xkXVxuICovXG5cbi8qKlxuICogQHR5cGVkZWYge3tlbHQ6IEVsZW1lbnQsIG1lc3NhZ2U6IHN0cmluZywgdmFsaWRpdHk6IFZhbGlkaXR5U3RhdGV9fSBIdG14RWxlbWVudFZhbGlkYXRpb25FcnJvclxuICovXG5cbi8qKlxuICogQHR5cGVkZWYge1JlY29yZDxzdHJpbmcsIHN0cmluZz59IEh0bXhIZWFkZXJTcGVjaWZpY2F0aW9uXG4gKiBAcHJvcGVydHkgeyd0cnVlJ30gSFgtUmVxdWVzdFxuICogQHByb3BlcnR5IHtzdHJpbmd8bnVsbH0gSFgtVHJpZ2dlclxuICogQHByb3BlcnR5IHtzdHJpbmd8bnVsbH0gSFgtVHJpZ2dlci1OYW1lXG4gKiBAcHJvcGVydHkge3N0cmluZ3xudWxsfSBIWC1UYXJnZXRcbiAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBIWC1DdXJyZW50LVVSTFxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtIWC1Qcm9tcHRdXG4gKiBAcHJvcGVydHkgeyd0cnVlJ30gW0hYLUJvb3N0ZWRdXG4gKiBAcHJvcGVydHkge3N0cmluZ30gW0NvbnRlbnQtVHlwZV1cbiAqIEBwcm9wZXJ0eSB7J3RydWUnfSBbSFgtSGlzdG9yeS1SZXN0b3JlLVJlcXVlc3RdXG4gKi9cblxuLyoqIEB0eXBlZGVmIEh0bXhBamF4SGVscGVyQ29udGV4dFxuICogQHByb3BlcnR5IHtFbGVtZW50fHN0cmluZ30gW3NvdXJjZV1cbiAqIEBwcm9wZXJ0eSB7RXZlbnR9IFtldmVudF1cbiAqIEBwcm9wZXJ0eSB7SHRteEFqYXhIYW5kbGVyfSBbaGFuZGxlcl1cbiAqIEBwcm9wZXJ0eSB7RWxlbWVudHxzdHJpbmd9IFt0YXJnZXRdXG4gKiBAcHJvcGVydHkge0h0bXhTd2FwU3R5bGV9IFtzd2FwXVxuICogQHByb3BlcnR5IHtPYmplY3R8Rm9ybURhdGF9IFt2YWx1ZXNdXG4gKiBAcHJvcGVydHkge1JlY29yZDxzdHJpbmcsc3RyaW5nPn0gW2hlYWRlcnNdXG4gKiBAcHJvcGVydHkge3N0cmluZ30gW3NlbGVjdF1cbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIHtPYmplY3R9IEh0bXhSZXF1ZXN0Q29uZmlnXG4gKiBAcHJvcGVydHkge2Jvb2xlYW59IGJvb3N0ZWRcbiAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gdXNlVXJsUGFyYW1zXG4gKiBAcHJvcGVydHkge0Zvcm1EYXRhfSBmb3JtRGF0YVxuICogQHByb3BlcnR5IHtPYmplY3R9IHBhcmFtZXRlcnMgZm9ybURhdGEgcHJveHlcbiAqIEBwcm9wZXJ0eSB7Rm9ybURhdGF9IHVuZmlsdGVyZWRGb3JtRGF0YVxuICogQHByb3BlcnR5IHtPYmplY3R9IHVuZmlsdGVyZWRQYXJhbWV0ZXJzIHVuZmlsdGVyZWRGb3JtRGF0YSBwcm94eVxuICogQHByb3BlcnR5IHtIdG14SGVhZGVyU3BlY2lmaWNhdGlvbn0gaGVhZGVyc1xuICogQHByb3BlcnR5IHtFbGVtZW50fSB0YXJnZXRcbiAqIEBwcm9wZXJ0eSB7SHR0cFZlcmJ9IHZlcmJcbiAqIEBwcm9wZXJ0eSB7SHRteEVsZW1lbnRWYWxpZGF0aW9uRXJyb3JbXX0gZXJyb3JzXG4gKiBAcHJvcGVydHkge2Jvb2xlYW59IHdpdGhDcmVkZW50aWFsc1xuICogQHByb3BlcnR5IHtudW1iZXJ9IHRpbWVvdXRcbiAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBwYXRoXG4gKiBAcHJvcGVydHkge0V2ZW50fSB0cmlnZ2VyaW5nRXZlbnRcbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIHtPYmplY3R9IEh0bXhSZXNwb25zZUluZm9cbiAqIEBwcm9wZXJ0eSB7WE1MSHR0cFJlcXVlc3R9IHhoclxuICogQHByb3BlcnR5IHtFbGVtZW50fSB0YXJnZXRcbiAqIEBwcm9wZXJ0eSB7SHRteFJlcXVlc3RDb25maWd9IHJlcXVlc3RDb25maWdcbiAqIEBwcm9wZXJ0eSB7SHRteEFqYXhFdGN9IGV0Y1xuICogQHByb3BlcnR5IHtib29sZWFufSBib29zdGVkXG4gKiBAcHJvcGVydHkge3N0cmluZ30gc2VsZWN0XG4gKiBAcHJvcGVydHkge3tyZXF1ZXN0UGF0aDogc3RyaW5nLCBmaW5hbFJlcXVlc3RQYXRoOiBzdHJpbmcsIHJlc3BvbnNlUGF0aDogc3RyaW5nfG51bGwsIGFuY2hvcjogc3RyaW5nfX0gcGF0aEluZm9cbiAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gW2ZhaWxlZF1cbiAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gW3N1Y2Nlc3NmdWxdXG4gKiBAcHJvcGVydHkge2Jvb2xlYW59IFtrZWVwSW5kaWNhdG9yc11cbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIHtPYmplY3R9IEh0bXhBamF4RXRjXG4gKiBAcHJvcGVydHkge2Jvb2xlYW59IFtyZXR1cm5Qcm9taXNlXVxuICogQHByb3BlcnR5IHtIdG14QWpheEhhbmRsZXJ9IFtoYW5kbGVyXVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtzZWxlY3RdXG4gKiBAcHJvcGVydHkge0VsZW1lbnR9IFt0YXJnZXRPdmVycmlkZV1cbiAqIEBwcm9wZXJ0eSB7SHRteFN3YXBTdHlsZX0gW3N3YXBPdmVycmlkZV1cbiAqIEBwcm9wZXJ0eSB7UmVjb3JkPHN0cmluZyxzdHJpbmc+fSBbaGVhZGVyc11cbiAqIEBwcm9wZXJ0eSB7T2JqZWN0fEZvcm1EYXRhfSBbdmFsdWVzXVxuICogQHByb3BlcnR5IHtib29sZWFufSBbY3JlZGVudGlhbHNdXG4gKiBAcHJvcGVydHkge251bWJlcn0gW3RpbWVvdXRdXG4gKi9cblxuLyoqXG4gKiBAdHlwZWRlZiB7T2JqZWN0fSBIdG14UmVzcG9uc2VIYW5kbGluZ0NvbmZpZ1xuICogQHByb3BlcnR5IHtzdHJpbmd9IFtjb2RlXVxuICogQHByb3BlcnR5IHtib29sZWFufSBzd2FwXG4gKiBAcHJvcGVydHkge2Jvb2xlYW59IFtlcnJvcl1cbiAqIEBwcm9wZXJ0eSB7Ym9vbGVhbn0gW2lnbm9yZVRpdGxlXVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtzZWxlY3RdXG4gKiBAcHJvcGVydHkge3N0cmluZ30gW3RhcmdldF1cbiAqIEBwcm9wZXJ0eSB7c3RyaW5nfSBbc3dhcE92ZXJyaWRlXVxuICogQHByb3BlcnR5IHtzdHJpbmd9IFtldmVudF1cbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIHtIdG14UmVzcG9uc2VJbmZvICYge3Nob3VsZFN3YXA6IGJvb2xlYW4sIHNlcnZlclJlc3BvbnNlOiBhbnksIGlzRXJyb3I6IGJvb2xlYW4sIGlnbm9yZVRpdGxlOiBib29sZWFuLCBzZWxlY3RPdmVycmlkZTpzdHJpbmd9fSBIdG14QmVmb3JlU3dhcERldGFpbHNcbiAqL1xuXG4vKipcbiAqIEBjYWxsYmFjayBIdG14QWpheEhhbmRsZXJcbiAqIEBwYXJhbSB7RWxlbWVudH0gZWx0XG4gKiBAcGFyYW0ge0h0bXhSZXNwb25zZUluZm99IHJlc3BvbnNlSW5mb1xuICovXG5cbi8qKlxuICogQHR5cGVkZWYgeygoKSA9PiB2b2lkKX0gSHRteFNldHRsZVRhc2tcbiAqL1xuXG4vKipcbiAqIEB0eXBlZGVmIHtPYmplY3R9IEh0bXhTZXR0bGVJbmZvXG4gKiBAcHJvcGVydHkge0h0bXhTZXR0bGVUYXNrW119IHRhc2tzXG4gKiBAcHJvcGVydHkge0VsZW1lbnRbXX0gZWx0c1xuICogQHByb3BlcnR5IHtzdHJpbmd9IFt0aXRsZV1cbiAqL1xuXG4vKipcbiAqIEBzZWUgaHR0cHM6Ly9naXRodWIuY29tL2JpZ3NreXNvZnR3YXJlL2h0bXgtZXh0ZW5zaW9ucy9ibG9iL21haW4vUkVBRE1FLm1kXG4gKiBAdHlwZWRlZiB7T2JqZWN0fSBIdG14RXh0ZW5zaW9uXG4gKiBAcHJvcGVydHkgeyhhcGk6IGFueSkgPT4gdm9pZH0gaW5pdFxuICogQHByb3BlcnR5IHsobmFtZTogc3RyaW5nLCBldmVudDogRXZlbnR8Q3VzdG9tRXZlbnQpID0+IGJvb2xlYW59IG9uRXZlbnRcbiAqIEBwcm9wZXJ0eSB7KHRleHQ6IHN0cmluZywgeGhyOiBYTUxIdHRwUmVxdWVzdCwgZWx0OiBFbGVtZW50KSA9PiBzdHJpbmd9IHRyYW5zZm9ybVJlc3BvbnNlXG4gKiBAcHJvcGVydHkgeyhzd2FwU3R5bGU6IEh0bXhTd2FwU3R5bGUpID0+IGJvb2xlYW59IGlzSW5saW5lU3dhcFxuICogQHByb3BlcnR5IHsoc3dhcFN0eWxlOiBIdG14U3dhcFN0eWxlLCB0YXJnZXQ6IE5vZGUsIGZyYWdtZW50OiBOb2RlLCBzZXR0bGVJbmZvOiBIdG14U2V0dGxlSW5mbykgPT4gYm9vbGVhbnxOb2RlW119IGhhbmRsZVN3YXBcbiAqIEBwcm9wZXJ0eSB7KHhocjogWE1MSHR0cFJlcXVlc3QsIHBhcmFtZXRlcnM6IEZvcm1EYXRhLCBlbHQ6IE5vZGUpID0+ICp8c3RyaW5nfG51bGx9IGVuY29kZVBhcmFtZXRlcnNcbiAqIEBwcm9wZXJ0eSB7KCkgPT4gc3RyaW5nW118bnVsbH0gZ2V0U2VsZWN0b3JzXG4gKi9cbmV4cG9ydCBkZWZhdWx0IGh0bXhcbiIsImltcG9ydCBodG14IGZyb20gXCJodG14Lm9yZ1wiO1xuXG5mdW5jdGlvbiBkZXBlbmRzT24ocGF0aFNwZWM6IGFueSwgdXJsOiBzdHJpbmcpIHtcbiAgaWYgKHBhdGhTcGVjID09PSBcImlnbm9yZVwiKSB7XG4gICAgcmV0dXJuIGZhbHNlO1xuICB9XG4gIGNvbnN0IGRlcGVuZGVuY3lQYXRoID0gcGF0aFNwZWMuc3BsaXQoXCIvXCIpO1xuICBjb25zdCB1cmxQYXRoID0gdXJsLnNwbGl0KFwiL1wiKTtcbiAgZm9yIChsZXQgaSA9IDA7IGkgPCB1cmxQYXRoLmxlbmd0aDsgaSsrKSB7XG4gICAgY29uc3QgZGVwZW5kZW5jeUVsZW1lbnQgPSBkZXBlbmRlbmN5UGF0aC5zaGlmdCgpO1xuICAgIGNvbnN0IHBhdGhFbGVtZW50ID0gdXJsUGF0aFtpXTtcbiAgICBpZiAoZGVwZW5kZW5jeUVsZW1lbnQgIT09IHBhdGhFbGVtZW50ICYmIGRlcGVuZGVuY3lFbGVtZW50ICE9PSBcIipcIikge1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgICBpZiAoXG4gICAgICBkZXBlbmRlbmN5UGF0aC5sZW5ndGggPT09IDAgfHxcbiAgICAgIChkZXBlbmRlbmN5UGF0aC5sZW5ndGggPT09IDEgJiYgZGVwZW5kZW5jeVBhdGhbMF0gPT09IFwiXCIpXG4gICAgKSB7XG4gICAgICByZXR1cm4gdHJ1ZTtcbiAgICB9XG4gIH1cbiAgcmV0dXJuIGZhbHNlO1xufVxuXG5mdW5jdGlvbiByZWZyZXNoUGF0aChwYXRoOiBzdHJpbmcpIHtcbiAgY29uc3QgZWx0c1dpdGhEZXBzID0gaHRteC5maW5kQWxsKFwiW3BhdGgtZGVwc11cIik7XG4gIGZvciAobGV0IGkgPSAwOyBpIDwgZWx0c1dpdGhEZXBzLmxlbmd0aDsgaSsrKSB7XG4gICAgY29uc3QgZWx0ID0gZWx0c1dpdGhEZXBzW2ldO1xuICAgIGlmIChkZXBlbmRzT24oZWx0LmdldEF0dHJpYnV0ZShcInBhdGgtZGVwc1wiKSwgcGF0aCkpIHtcbiAgICAgIGh0bXgudHJpZ2dlcihlbHQsIFwicGF0aC1kZXBzXCIsIG51bGwpO1xuICAgIH1cbiAgfVxufVxuXG5odG14LmRlZmluZUV4dGVuc2lvbihcInBhdGgtZGVwc1wiLCB7XG4gIC8vIEB0cy1pZ25vcmVcbiAgb25FdmVudDogZnVuY3Rpb24gKG5hbWUsIGV2dCkge1xuICAgIGlmICghKGV2dCBpbnN0YW5jZW9mIEN1c3RvbUV2ZW50KSkge1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgICBpZiAobmFtZSA9PT0gXCJodG14OmJlZm9yZU9uTG9hZFwiKSB7XG4gICAgICBjb25zdCBjb25maWcgPSBldnQuZGV0YWlsLnJlcXVlc3RDb25maWc7XG4gICAgICAvLyBtdXRhdGluZyBjYWxsXG4gICAgICBpZiAoXG4gICAgICAgIGNvbmZpZyAmJlxuICAgICAgICBjb25maWcudmVyYiAhPT0gXCJnZXRcIiAmJlxuICAgICAgICBldnQudGFyZ2V0ICE9IG51bGwgJiZcbiAgICAgICAgZXZ0LnRhcmdldCBpbnN0YW5jZW9mIEVsZW1lbnQgJiZcbiAgICAgICAgZXZ0LnRhcmdldC5nZXRBdHRyaWJ1dGUoXCJwYXRoLWRlcHNcIikgIT09IFwiaWdub3JlXCJcbiAgICAgICkge1xuICAgICAgICByZWZyZXNoUGF0aChjb25maWcucGF0aCk7XG4gICAgICB9XG4gICAgfVxuICB9LFxufSk7XG4iLCJpbXBvcnQgaHRteCwgeyBIdG14U2V0dGxlSW5mbywgSHRteFN3YXBTdHlsZSB9IGZyb20gXCJodG14Lm9yZ1wiO1xuXG5odG14LmRlZmluZUV4dGVuc2lvbihcInRyaWdnZXItY2hpbGRyZW5cIiwge1xuICBvbkV2ZW50OiAobmFtZSwgZXZ0OiBFdmVudCB8IEN1c3RvbUV2ZW50KSA9PiB7XG4gICAgaWYgKCEoZXZ0IGluc3RhbmNlb2YgQ3VzdG9tRXZlbnQpKSB7XG4gICAgICByZXR1cm4gZmFsc2U7XG4gICAgfVxuICAgIGNvbnN0IHRhcmdldCA9IGV2dC5kZXRhaWwudGFyZ2V0IGFzIEhUTUxFbGVtZW50O1xuICAgIGlmICh0YXJnZXQgJiYgdGFyZ2V0LmNoaWxkcmVuKSB7XG4gICAgICBBcnJheS5mcm9tKHRhcmdldC5jaGlsZHJlbikuZm9yRWFjaCgoZSkgPT4ge1xuICAgICAgICBodG14LnRyaWdnZXIoZSwgbmFtZSwgbnVsbCk7XG4gICAgICB9KTtcbiAgICB9XG4gICAgcmV0dXJuIHRydWU7XG4gIH0sXG4gIGluaXQ6IGZ1bmN0aW9uIChhcGk6IGFueSk6IHZvaWQge30sXG4gIHRyYW5zZm9ybVJlc3BvbnNlOiBmdW5jdGlvbiAoXG4gICAgdGV4dDogc3RyaW5nLFxuICAgIHhocjogWE1MSHR0cFJlcXVlc3QsXG4gICAgZWx0OiBFbGVtZW50LFxuICApOiBzdHJpbmcge1xuICAgIHJldHVybiB0ZXh0O1xuICB9LFxuICBpc0lubGluZVN3YXA6IGZ1bmN0aW9uIChzd2FwU3R5bGU6IEh0bXhTd2FwU3R5bGUpOiBib29sZWFuIHtcbiAgICByZXR1cm4gZmFsc2U7XG4gIH0sXG4gIGhhbmRsZVN3YXA6IGZ1bmN0aW9uIChcbiAgICBzd2FwU3R5bGU6IEh0bXhTd2FwU3R5bGUsXG4gICAgdGFyZ2V0OiBOb2RlLFxuICAgIGZyYWdtZW50OiBOb2RlLFxuICAgIHNldHRsZUluZm86IEh0bXhTZXR0bGVJbmZvLFxuICApOiBib29sZWFuIHwgTm9kZVtdIHtcbiAgICByZXR1cm4gZmFsc2U7XG4gIH0sXG4gIGVuY29kZVBhcmFtZXRlcnM6IGZ1bmN0aW9uIChcbiAgICB4aHI6IFhNTEh0dHBSZXF1ZXN0LFxuICAgIHBhcmFtZXRlcnM6IEZvcm1EYXRhLFxuICAgIGVsdDogTm9kZSxcbiAgKSB7fSxcbiAgZ2V0U2VsZWN0b3JzOiBmdW5jdGlvbiAoKTogc3RyaW5nW10gfCBudWxsIHtcbiAgICByZXR1cm4gbnVsbDtcbiAgfSxcbn0pO1xuIiwiaW1wb3J0IGh0bXggZnJvbSBcImh0bXgub3JnXCI7XG5cbmh0bXguZGVmaW5lRXh0ZW5zaW9uKFwiZGVidWdcIiwge1xuICAvLyBAdHMtaWdub3JlXG4gIG9uRXZlbnQ6IGZ1bmN0aW9uIChuYW1lLCBldnQpIHtcbiAgICBpZiAoY29uc29sZS5kZWJ1Zykge1xuICAgICAgY29uc29sZS5kZWJ1ZyhuYW1lKTtcbiAgICB9IGVsc2UgaWYgKGNvbnNvbGUpIHtcbiAgICAgIGNvbnNvbGUubG9nKFwiREVCVUc6XCIsIG5hbWUpO1xuICAgIH0gZWxzZSB7XG4gICAgICAvLyBub29wXG4gICAgfVxuICB9LFxufSk7XG4iLCJpbXBvcnQgaHRteCBmcm9tIFwiaHRteC5vcmdcIjtcbmNvbnN0IGNvbmZpZzogYW55ID0gaHRteC5jb25maWc7XG5cbi8qKiBAdHlwZSB7aW1wb3J0KFwiLi4vaHRteFwiKS5IdG14SW50ZXJuYWxBcGl9ICovXG5sZXQgYXBpOiBhbnk7XG5cbmNvbnN0IGF0dHJQcmVmaXggPSBcImh4LXRhcmdldC1cIjtcblxuLy8gSUUxMSBkb2Vzbid0IHN1cHBvcnQgc3RyaW5nLnN0YXJ0c1dpdGhcbmZ1bmN0aW9uIHN0YXJ0c1dpdGgoc3RyOiBzdHJpbmcsIHByZWZpeDogc3RyaW5nKSB7XG4gIHJldHVybiBzdHIuc3Vic3RyaW5nKDAsIHByZWZpeC5sZW5ndGgpID09PSBwcmVmaXg7XG59XG5cbi8qKlxuICogQHBhcmFtIHtIVE1MRWxlbWVudH0gZWx0XG4gKiBAcGFyYW0gcmVzcENvZGVOdW1iZXJcbiAqIEByZXR1cm5zIHtIVE1MRWxlbWVudCB8IG51bGx9XG4gKi9cbmZ1bmN0aW9uIGdldFJlc3BDb2RlVGFyZ2V0KGVsdDogRWxlbWVudCwgcmVzcENvZGVOdW1iZXI6IG51bWJlcikge1xuICBpZiAoIWVsdCB8fCAhcmVzcENvZGVOdW1iZXIpIHJldHVybiBudWxsO1xuXG4gIGNvbnN0IHJlc3BDb2RlID0gcmVzcENvZGVOdW1iZXIudG9TdHJpbmcoKTtcblxuICAvLyAnKicgaXMgdGhlIG9yaWdpbmFsIHN5bnRheCwgYXMgdGhlIG9idmlvdXMgY2hhcmFjdGVyIGZvciBhIHdpbGRjYXJkLlxuICAvLyBUaGUgJ3gnIGFsdGVybmF0aXZlIHdhcyBhZGRlZCBmb3IgbWF4aW11bSBjb21wYXRpYmlsaXR5IHdpdGggSFRNTFxuICAvLyB0ZW1wbGF0aW5nIGVuZ2luZXMsIGR1ZSB0byBhbWJpZ3VpdHkgYXJvdW5kIHdoaWNoIGNoYXJhY3RlcnMgYXJlXG4gIC8vIHN1cHBvcnRlZCBpbiBIVE1MIGF0dHJpYnV0ZXMuXG4gIC8vXG4gIC8vIFN0YXJ0IHdpdGggdGhlIG1vc3Qgc3BlY2lmaWMgcG9zc2libGUgYXR0cmlidXRlIGFuZCBnZW5lcmFsaXplIGZyb21cbiAgLy8gdGhlcmUuXG4gIGNvbnN0IGF0dHJQb3NzaWJpbGl0aWVzID0gW1xuICAgIHJlc3BDb2RlLFxuXG4gICAgcmVzcENvZGUuc3Vic3RyKDAsIDIpICsgXCIqXCIsXG4gICAgcmVzcENvZGUuc3Vic3RyKDAsIDIpICsgXCJ4XCIsXG5cbiAgICByZXNwQ29kZS5zdWJzdHIoMCwgMSkgKyBcIipcIixcbiAgICByZXNwQ29kZS5zdWJzdHIoMCwgMSkgKyBcInhcIixcbiAgICByZXNwQ29kZS5zdWJzdHIoMCwgMSkgKyBcIioqXCIsXG4gICAgcmVzcENvZGUuc3Vic3RyKDAsIDEpICsgXCJ4eFwiLFxuXG4gICAgXCIqXCIsXG4gICAgXCJ4XCIsXG4gICAgXCIqKipcIixcbiAgICBcInh4eFwiLFxuICBdO1xuICBpZiAoc3RhcnRzV2l0aChyZXNwQ29kZSwgXCI0XCIpIHx8IHN0YXJ0c1dpdGgocmVzcENvZGUsIFwiNVwiKSkge1xuICAgIGF0dHJQb3NzaWJpbGl0aWVzLnB1c2goXCJlcnJvclwiKTtcbiAgfVxuXG4gIGZvciAobGV0IGkgPSAwOyBpIDwgYXR0clBvc3NpYmlsaXRpZXMubGVuZ3RoOyBpKyspIHtcbiAgICBjb25zdCBhdHRyID0gYXR0clByZWZpeCArIGF0dHJQb3NzaWJpbGl0aWVzW2ldO1xuICAgIGNvbnN0IGF0dHJWYWx1ZSA9IGFwaS5nZXRDbG9zZXN0QXR0cmlidXRlVmFsdWUoZWx0LCBhdHRyKTtcbiAgICBpZiAoYXR0clZhbHVlKSB7XG4gICAgICBpZiAoYXR0clZhbHVlID09PSBcInRoaXNcIikge1xuICAgICAgICByZXR1cm4gYXBpLmZpbmRUaGlzRWxlbWVudChlbHQsIGF0dHIpO1xuICAgICAgfSBlbHNlIHtcbiAgICAgICAgcmV0dXJuIGFwaS5xdWVyeVNlbGVjdG9yRXh0KGVsdCwgYXR0clZhbHVlKTtcbiAgICAgIH1cbiAgICB9XG4gIH1cblxuICByZXR1cm4gbnVsbDtcbn1cblxuLyoqIEBwYXJhbSB7RXZlbnR9IGV2dCAqL1xuZnVuY3Rpb24gaGFuZGxlRXJyb3JGbGFnKGV2dDogQ3VzdG9tRXZlbnQpIHtcbiAgaWYgKGV2dC5kZXRhaWwuaXNFcnJvcikge1xuICAgIGlmIChjb25maWcucmVzcG9uc2VUYXJnZXRVbnNldHNFcnJvcikge1xuICAgICAgZXZ0LmRldGFpbC5pc0Vycm9yID0gZmFsc2U7XG4gICAgfVxuICB9IGVsc2UgaWYgKGNvbmZpZy5yZXNwb25zZVRhcmdldFNldHNFcnJvcikge1xuICAgIGV2dC5kZXRhaWwuaXNFcnJvciA9IHRydWU7XG4gIH1cbn1cblxuaHRteC5kZWZpbmVFeHRlbnNpb24oXCJyZXNwb25zZS10YXJnZXRzXCIsIHtcbiAgLy8gQHRzLWlnbm9yZVxuICBpbml0OiAoYXBpUmVmKSA9PiB7XG4gICAgYXBpID0gYXBpUmVmO1xuXG4gICAgaWYgKGNvbmZpZy5yZXNwb25zZVRhcmdldFVuc2V0c0Vycm9yID09PSB1bmRlZmluZWQpIHtcbiAgICAgIGNvbmZpZy5yZXNwb25zZVRhcmdldFVuc2V0c0Vycm9yID0gdHJ1ZTtcbiAgICB9XG4gICAgaWYgKGNvbmZpZy5yZXNwb25zZVRhcmdldFNldHNFcnJvciA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBjb25maWcucmVzcG9uc2VUYXJnZXRTZXRzRXJyb3IgPSBmYWxzZTtcbiAgICB9XG4gICAgaWYgKGNvbmZpZy5yZXNwb25zZVRhcmdldFByZWZlcnNFeGlzdGluZyA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBjb25maWcucmVzcG9uc2VUYXJnZXRQcmVmZXJzRXhpc3RpbmcgPSBmYWxzZTtcbiAgICB9XG4gICAgaWYgKGNvbmZpZy5yZXNwb25zZVRhcmdldFByZWZlcnNSZXRhcmdldEhlYWRlciA9PT0gdW5kZWZpbmVkKSB7XG4gICAgICBjb25maWcucmVzcG9uc2VUYXJnZXRQcmVmZXJzUmV0YXJnZXRIZWFkZXIgPSB0cnVlO1xuICAgIH1cbiAgfSxcblxuICAvLyBAdHMtaWdub3JlXG4gIG9uRXZlbnQ6IChuYW1lLCBldnQpID0+IHtcbiAgICBpZiAoIShldnQgaW5zdGFuY2VvZiBDdXN0b21FdmVudCkpIHtcbiAgICAgIHJldHVybiBmYWxzZTtcbiAgICB9XG4gICAgaWYgKFxuICAgICAgbmFtZSA9PT0gXCJodG14OmJlZm9yZVN3YXBcIiAmJlxuICAgICAgZXZ0LmRldGFpbC54aHIgJiZcbiAgICAgIGV2dC5kZXRhaWwueGhyLnN0YXR1cyAhPT0gMjAwXG4gICAgKSB7XG4gICAgICBpZiAoZXZ0LmRldGFpbC50YXJnZXQpIHtcbiAgICAgICAgaWYgKGNvbmZpZy5yZXNwb25zZVRhcmdldFByZWZlcnNFeGlzdGluZykge1xuICAgICAgICAgIGV2dC5kZXRhaWwuc2hvdWxkU3dhcCA9IHRydWU7XG4gICAgICAgICAgaGFuZGxlRXJyb3JGbGFnKGV2dCk7XG4gICAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICAgIH1cbiAgICAgICAgaWYgKFxuICAgICAgICAgIGNvbmZpZy5yZXNwb25zZVRhcmdldFByZWZlcnNSZXRhcmdldEhlYWRlciAmJlxuICAgICAgICAgIGV2dC5kZXRhaWwueGhyLmdldEFsbFJlc3BvbnNlSGVhZGVycygpLm1hdGNoKC9IWC1SZXRhcmdldDovaSlcbiAgICAgICAgKSB7XG4gICAgICAgICAgZXZ0LmRldGFpbC5zaG91bGRTd2FwID0gdHJ1ZTtcbiAgICAgICAgICBoYW5kbGVFcnJvckZsYWcoZXZ0KTtcbiAgICAgICAgICByZXR1cm4gdHJ1ZTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgICAgaWYgKCFldnQuZGV0YWlsLnJlcXVlc3RDb25maWcpIHtcbiAgICAgICAgcmV0dXJuIHRydWU7XG4gICAgICB9XG4gICAgICBjb25zdCB0YXJnZXQgPSBnZXRSZXNwQ29kZVRhcmdldChcbiAgICAgICAgZXZ0LmRldGFpbC5yZXF1ZXN0Q29uZmlnLmVsdCxcbiAgICAgICAgZXZ0LmRldGFpbC54aHIuc3RhdHVzLFxuICAgICAgKTtcbiAgICAgIGlmICh0YXJnZXQpIHtcbiAgICAgICAgaGFuZGxlRXJyb3JGbGFnKGV2dCk7XG4gICAgICAgIGV2dC5kZXRhaWwuc2hvdWxkU3dhcCA9IHRydWU7XG4gICAgICAgIGV2dC5kZXRhaWwudGFyZ2V0ID0gdGFyZ2V0O1xuICAgICAgfVxuICAgICAgcmV0dXJuIHRydWU7XG4gICAgfVxuICB9LFxufSk7XG4iLCJpbXBvcnQgaHRteCBmcm9tIFwiaHRteC5vcmdcIjtcblxuaHRteC5kZWZpbmVFeHRlbnNpb24oXCJtdXRhdGlvbi1lcnJvclwiLCB7XG4gIC8vIEB0cy1pZ25vcmVcbiAgb25FdmVudDogKG5hbWUsIGV2dCkgPT4ge1xuICAgIGlmICghKGV2dCBpbnN0YW5jZW9mIEN1c3RvbUV2ZW50KSkge1xuICAgICAgcmV0dXJuIGZhbHNlO1xuICAgIH1cbiAgICBpZiAobmFtZSA9PT0gXCJodG14OmFmdGVyUmVxdWVzdFwiKSB7XG4gICAgICBpZiAoIWV2dC5kZXRhaWwgfHwgIWV2dC5kZXRhaWwueGhyKSB7XG4gICAgICAgIHJldHVybjtcbiAgICAgIH1cbiAgICAgIGNvbnN0IHN0YXR1cyA9IGV2dC5kZXRhaWwueGhyLnN0YXR1cztcbiAgICAgIGlmIChzdGF0dXMgPj0gNDAwKSB7XG4gICAgICAgIGh0bXguZmluZEFsbChcIltoeC1vblxcXFw6XFxcXDptdXRhdGlvbi1lcnJvcl1cIikuZm9yRWFjaCgoZWxlbWVudCkgPT4ge1xuICAgICAgICAgIGh0bXgudHJpZ2dlcihlbGVtZW50LCBcImh0bXg6bXV0YXRpb24tZXJyb3JcIiwgeyBzdGF0dXMgfSk7XG4gICAgICAgIH0pO1xuICAgICAgfVxuICAgIH1cbiAgfSxcbn0pO1xuIiwidHlwZSBXc09wdHMgPSB7XG4gICAgdXJsOiBzdHJpbmc7XG4gICAgcmVjb25uZWN0SW50ZXJ2YWw/OiBudW1iZXI7XG4gICAgb25PcGVuPzogKCkgPT4gdm9pZDtcbiAgICBvbk1lc3NhZ2U6IChtZXNzYWdlOiBzdHJpbmcpID0+IHZvaWQ7XG4gICAgb25FcnJvcj86IChlcnJvcjogRXZlbnQpID0+IHZvaWQ7XG4gICAgb25DbG9zZT86ICgpID0+IHZvaWQ7XG59XG5cbmV4cG9ydCBmdW5jdGlvbiBjcmVhdGVXZWJTb2NrZXRDbGllbnQob3B0czogV3NPcHRzKSB7XG4gICAgbGV0IHNvY2tldDogV2ViU29ja2V0IHwgbnVsbCA9IG51bGw7XG4gICAgY29uc3QgY29ubmVjdCA9ICh0cmllczogbnVtYmVyKSA9PiB7XG4gICAgICAgIGNvbnNvbGUubG9nKCdjb25uZWN0aW5nIHRvIHdzJywgb3B0cy51cmwsICdhdHRlbXB0JywgdHJpZXMpXG4gICAgICAgIHNvY2tldCA9IG5ldyBXZWJTb2NrZXQob3B0cy51cmwpO1xuICAgICAgICAvLyBIYW5kbGUgaW5jb21pbmcgbWVzc2FnZXNcbiAgICAgICAgc29ja2V0Lm9ubWVzc2FnZSA9IChldmVudCkgPT4ge1xuICAgICAgICAgICAgb3B0cy5vbk1lc3NhZ2UoZXZlbnQuZGF0YSlcbiAgICAgICAgfTtcbiAgICAgICAgLy8gSGFuZGxlIGNvbm5lY3Rpb24gZXJyb3JzXG4gICAgICAgIHNvY2tldC5vbmVycm9yID0gKGVycm9yKSA9PiB7XG4gICAgICAgICAgIHRyeSB7XG4gICAgICAgICAgICAgICBzb2NrZXQ/LmNsb3NlKClcbiAgICAgICAgICAgfSBjYXRjaChleCkge1xuICAgICAgICAgICAgICAgLy8gbm9vcFxuICAgICAgICAgICB9XG4gICAgICAgICAgIHNvY2tldCA9IG51bGxcbiAgICAgICAgICAgIGxldCBpbnRlcnZhbCA9IHRyaWVzICogKG9wdHMucmVjb25uZWN0SW50ZXJ2YWwgfHwgMTAwMCk7XG4gICAgICAgICAgICBzZXRUaW1lb3V0KCgpID0+IGNvbm5lY3QodHJpZXMgKyAxKSwgaW50ZXJ2YWwpO1xuICAgICAgICB9O1xuICAgICAgICAvLyBIYW5kbGUgY29ubmVjdGlvbiBjbG9zZSBhbmQgYXR0ZW1wdCByZWNvbm5lY3Rpb25cbiAgICAgICAgc29ja2V0Lm9uY2xvc2UgPSAoKSA9PiB7XG4gICAgICAgICAgICBzb2NrZXQgPSBudWxsO1xuICAgICAgICAgICAgbGV0IGludGVydmFsID0gdHJpZXMgKiAob3B0cy5yZWNvbm5lY3RJbnRlcnZhbCB8fCAxMDAwKTtcbiAgICAgICAgICAgIHNldFRpbWVvdXQoKCkgPT4gY29ubmVjdCh0cmllcyArIDEpLCBpbnRlcnZhbCk7XG4gICAgICAgIH07XG4gICAgfTtcbiAgICBjb25uZWN0KDEpO1xuICAgIGNvbnN0IHNlbmRNZXNzYWdlID0gKG1lc3NhZ2U6IHN0cmluZykgPT4ge1xuICAgICAgICBpZiAoc29ja2V0ICYmIHNvY2tldC5yZWFkeVN0YXRlID09PSBXZWJTb2NrZXQuT1BFTikge1xuICAgICAgICAgICAgc29ja2V0LnNlbmQobWVzc2FnZSk7XG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgIHNldFRpbWVvdXQoKCkgPT4gc2VuZE1lc3NhZ2UobWVzc2FnZSksIDEwMCk7XG4gICAgICAgIH1cbiAgICB9O1xuICAgIHJldHVybiB7IHNlbmRNZXNzYWdlIH07XG59IiwiaW1wb3J0IGh0bXggZnJvbSBcImh0bXgub3JnXCI7XG5pbXBvcnQge2NyZWF0ZVdlYlNvY2tldENsaWVudH0gZnJvbSBcIi4uL3V0aWwvd3NcIjtcblxubGV0IGxhc3RWZXJzaW9uID0gXCJcIjtcblxuaHRteC5kZWZpbmVFeHRlbnNpb24oXCJsaXZlcmVsb2FkXCIsIHtcbiAgICBpbml0OiBmdW5jdGlvbiAoKSB7XG4gICAgICAgIGNvbnN0IGhvc3QgPSB3aW5kb3cubG9jYXRpb24uaG9zdDtcblxuICAgICAgICBsZXQgZW5hYmxlZCA9IGZhbHNlXG4gICAgICAgIGZvciAoY29uc3QgZWxlbWVudCBvZiBBcnJheS5mcm9tKGh0bXguZmluZEFsbChcIltoeC1leHRdXCIpKSkge1xuICAgICAgICAgICAgY29uc3QgdmFsdWUgPSBlbGVtZW50LmdldEF0dHJpYnV0ZShcImh4LWV4dFwiKTtcbiAgICAgICAgICAgIGlmKHZhbHVlPy5zcGxpdChcIiBcIikuaW5jbHVkZXMoXCJsaXZlcmVsb2FkXCIpKSB7XG4gICAgICAgICAgICAgICAgZW5hYmxlZCA9IHRydWVcbiAgICAgICAgICAgICAgICBicmVhaztcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuXG4gICAgICAgIGlmKCFlbmFibGVkKSB7XG4gICAgICAgICAgICByZXR1cm5cbiAgICAgICAgfVxuXG4gICAgICAgIGNvbnNvbGUubG9nKCdsaXZlcmVsb2FkIGV4dGVuc2lvbiBpbml0aWFsaXplZC4nKTtcblxuICAgICAgICBjcmVhdGVXZWJTb2NrZXRDbGllbnQoe1xuICAgICAgICAgICAgdXJsOiBgJHt3aW5kb3cubG9jYXRpb24ucHJvdG9jb2wgPT09ICdodHRwczonID8gJ3dzcycgOiAnd3MnfTovLyR7aG9zdH0vZGV2L2xpdmVyZWxvYWRgLFxuICAgICAgICAgICAgb25PcGVuOiAoKSA9PiB7XG4gICAgICAgICAgICB9LFxuICAgICAgICAgICAgb25NZXNzYWdlOiAobWVzc2FnZSkgPT4ge1xuICAgICAgICAgICAgICAgaWYobGFzdFZlcnNpb24gPT09IFwiXCIpIHtcbiAgICAgICAgICAgICAgICAgICBsYXN0VmVyc2lvbiA9IG1lc3NhZ2U7XG4gICAgICAgICAgICAgICB9XG4gICAgICAgICAgICAgICBpZihsYXN0VmVyc2lvbiAhPT0gbWVzc2FnZSkge1xuICAgICAgICAgICAgICAgICAgIGxhc3RWZXJzaW9uID0gbWVzc2FnZTtcbiAgICAgICAgICAgICAgICAgICByZWxvYWQoKVxuICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIG9uRXJyb3I6IChlcnJvcikgPT4ge1xuICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIG9uQ2xvc2U6ICgpID0+IHtcbiAgICAgICAgICAgIH1cbiAgICAgICAgfSlcbiAgICB9LFxuICAgIC8vIEB0cy1pZ25vcmVcbiAgICBvbkV2ZW50OiBmdW5jdGlvbiAobmFtZSwgZXZ0KSB7XG5cbiAgICB9LFxufSk7XG5cbmZ1bmN0aW9uIHJlbG9hZCgpIHtcbiAgICB3aW5kb3cubG9jYXRpb24ucmVsb2FkKClcbiAgICAvLyBmZXRjaCh3aW5kb3cubG9jYXRpb24uaHJlZikudGhlbihyZXNwb25zZSA9PiB7XG4gICAgLy8gICAgIHJldHVybiByZXNwb25zZS50ZXh0KCk7XG4gICAgLy8gfSkudGhlbihodG1sID0+IHtcbiAgICAvLyAgICBkb2N1bWVudC5vcGVuKCk7XG4gICAgLy8gICAgIGRvY3VtZW50LndyaXRlKGh0bWwpO1xuICAgIC8vICAgICBkb2N1bWVudC5jbG9zZSgpO1xuICAgIC8vIH0pLmNhdGNoKGVyciA9PiB7XG4gICAgLy8gICAgIGNvbnNvbGUubG9nKCdmYWlsZWQgdG8gZmV0Y2ggbGl2ZSByZWxvYWQnLCBlcnIpXG4gICAgLy8gICAgIHNldFRpbWVvdXQocmVsb2FkLCAxMDApXG4gICAgLy8gfSlcbn0iLCJpbXBvcnQgaHRteCBmcm9tIFwiaHRteC5vcmdcIjtcblxuY29uc3QgZXZhbEZ1bmNSZWdleCA9IC9fX2V2YWxfW0EtWmEtejAtOV0rXFwoXFwpL2dtXG5cbmh0bXguZGVmaW5lRXh0ZW5zaW9uKFwiaHRtZ29cIiwge1xuICAgIC8vIEB0cy1pZ25vcmVcbiAgICBvbkV2ZW50OiBmdW5jdGlvbiAobmFtZSwgZXZ0KSB7XG4gICAgICAgaWYobmFtZSA9PT0gXCJodG14OmJlZm9yZUNsZWFudXBFbGVtZW50XCIgJiYgZXZ0LnRhcmdldCkge1xuICAgICAgICAgICByZW1vdmVBc3NvY2lhdGVkU2NyaXB0cyhldnQudGFyZ2V0IGFzIEhUTUxFbGVtZW50KTtcbiAgICAgICB9XG4gICAgfSxcbn0pO1xuXG5mdW5jdGlvbiByZW1vdmVBc3NvY2lhdGVkU2NyaXB0cyhlbGVtZW50OiBIVE1MRWxlbWVudCkge1xuICAgY29uc3QgYXR0cmlidXRlcyA9IEFycmF5LmZyb20oZWxlbWVudC5hdHRyaWJ1dGVzKVxuICAgIGZvciAobGV0IGF0dHJpYnV0ZSBvZiBhdHRyaWJ1dGVzKSB7XG4gICAgICAgY29uc3QgbWF0Y2hlcyA9IGF0dHJpYnV0ZS52YWx1ZS5tYXRjaChldmFsRnVuY1JlZ2V4KSB8fCBbXVxuICAgICAgICBmb3IgKGxldCBtYXRjaCBvZiBtYXRjaGVzKSB7XG4gICAgICAgICAgICBjb25zdCBpZCA9IG1hdGNoLnJlcGxhY2UoXCIoKVwiLCBcIlwiKVxuICAgICAgICAgICAgY29uc3QgZWxlID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoaWQpXG4gICAgICAgICAgICBpZihlbGUgJiYgZWxlLnRhZ05hbWUgPT09IFwiU0NSSVBUXCIpIHtcbiAgICAgICAgICAgICAgICBlbGUucmVtb3ZlKClcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgIH1cbn0iLCJpbXBvcnQgaHRteCBmcm9tIFwiaHRteC5vcmdcIjtcbmltcG9ydCBcIi4vaHRteGV4dGVuc2lvbnMvcGF0aGRlcHNcIjtcbmltcG9ydCBcIi4vaHRteGV4dGVuc2lvbnMvdHJpZ2dlci1jaGlsZHJlblwiO1xuaW1wb3J0IFwiLi9odG14ZXh0ZW5zaW9ucy9kZWJ1Z1wiO1xuaW1wb3J0IFwiLi9odG14ZXh0ZW5zaW9ucy9yZXNwb25zZS10YXJnZXRzXCI7XG5pbXBvcnQgXCIuL2h0bXhleHRlbnNpb25zL211dGF0aW9uLWVycm9yXCI7XG5pbXBvcnQgXCIuL2h0bXhleHRlbnNpb25zL2xpdmVyZWxvYWRcIlxuaW1wb3J0IFwiLi9odG14ZXh0ZW5zaW9ucy9odG1nb1wiO1xuXG5mdW5jdGlvbiB3YXRjaFVybChjYWxsYmFjazogKG9sZFVybDogc3RyaW5nLCBuZXdVcmw6IHN0cmluZykgPT4gdm9pZCkge1xuICBsZXQgbGFzdFVybCA9IHdpbmRvdy5sb2NhdGlvbi5ocmVmO1xuICBzZXRJbnRlcnZhbCgoKSA9PiB7XG4gICAgaWYgKHdpbmRvdy5sb2NhdGlvbi5ocmVmICE9PSBsYXN0VXJsKSB7XG4gICAgICBjYWxsYmFjayhsYXN0VXJsLCB3aW5kb3cubG9jYXRpb24uaHJlZik7XG4gICAgICBsYXN0VXJsID0gd2luZG93LmxvY2F0aW9uLmhyZWY7XG4gICAgfVxuICB9LCAxMDApO1xufVxuXG53YXRjaFVybCgoXywgbmV3VXJsKSA9PiB7XG4gIG9uVXJsQ2hhbmdlKG5ld1VybCk7XG59KTtcblxuZnVuY3Rpb24gb25VcmxDaGFuZ2UobmV3VXJsOiBzdHJpbmcpIHtcbiAgbGV0IHVybCA9IG5ldyBVUkwobmV3VXJsKTtcblxuICBkb2N1bWVudC5xdWVyeVNlbGVjdG9yQWxsKFwiW2h4LXRyaWdnZXJdXCIpLmZvckVhY2goZnVuY3Rpb24gKGVsZW1lbnQpIHtcbiAgICBjb25zdCB0cmlnZ2VycyA9IGVsZW1lbnQuZ2V0QXR0cmlidXRlKFwiaHgtdHJpZ2dlclwiKTtcbiAgICBpZiAoIXRyaWdnZXJzKSB7XG4gICAgICByZXR1cm47XG4gICAgfVxuICAgIGNvbnN0IHNwbGl0ID0gdHJpZ2dlcnMuc3BsaXQoXCIsIFwiKTtcbiAgICBpZiAoc3BsaXQuZmluZCgocykgPT4gcyA9PT0gXCJ1cmxcIikpIHtcbiAgICAgIGh0bXguc3dhcChlbGVtZW50LCBcInVybFwiLCB7XG4gICAgICAgIHN3YXBTdHlsZTogXCJvdXRlckhUTUxcIixcbiAgICAgICAgc3dhcERlbGF5OiAwLFxuICAgICAgICBzZXR0bGVEZWxheTogMCxcbiAgICAgIH0pO1xuICAgIH0gZWxzZSB7XG4gICAgICBmb3IgKGxldCBba2V5LCB2YWx1ZXNdIG9mIHVybC5zZWFyY2hQYXJhbXMpIHtcbiAgICAgICAgbGV0IGV2ZW50TmFtZSA9IFwicXM6XCIgKyBrZXk7XG4gICAgICAgIGlmICh0cmlnZ2Vycy5pbmNsdWRlcyhldmVudE5hbWUpKSB7XG4gICAgICAgICAgY29uc29sZS5sb2coXCJ0cmlnZ2VyaW5nXCIsIGV2ZW50TmFtZSk7XG4gICAgICAgICAgaHRteC50cmlnZ2VyKGVsZW1lbnQsIGV2ZW50TmFtZSwgbnVsbCk7XG4gICAgICAgICAgYnJlYWs7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG4gIH0pO1xuXG4gIGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3JBbGwoXCJbaHgtbWF0Y2gtcXBdXCIpLmZvckVhY2goKGVsKSA9PiB7XG4gICAgbGV0IGhhc01hdGNoID0gZmFsc2U7XG4gICAgZm9yIChsZXQgbmFtZSBvZiBlbC5nZXRBdHRyaWJ1dGVOYW1lcygpKSB7XG4gICAgICBpZiAobmFtZS5zdGFydHNXaXRoKFwiaHgtbWF0Y2gtcXAtbWFwcGluZzpcIikpIHtcbiAgICAgICAgbGV0IG1hdGNoID0gbmFtZS5yZXBsYWNlKFwiaHgtbWF0Y2gtcXAtbWFwcGluZzpcIiwgXCJcIik7XG4gICAgICAgIGxldCB2YWx1ZSA9IHVybC5zZWFyY2hQYXJhbXMuZ2V0KG1hdGNoKTtcbiAgICAgICAgaWYgKHZhbHVlKSB7XG4gICAgICAgICAgaHRteC5zd2FwKGVsLCBlbC5nZXRBdHRyaWJ1dGUobmFtZSkgPz8gXCJcIiwge1xuICAgICAgICAgICAgc3dhcFN0eWxlOiBcImlubmVySFRNTFwiLFxuICAgICAgICAgICAgc3dhcERlbGF5OiAwLFxuICAgICAgICAgICAgc2V0dGxlRGVsYXk6IDAsXG4gICAgICAgICAgfSk7XG4gICAgICAgICAgaGFzTWF0Y2ggPSB0cnVlO1xuICAgICAgICAgIGJyZWFrO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuICAgIGlmICghaGFzTWF0Y2gpIHtcbiAgICAgIGxldCBkZWZhdWx0S2V5ID0gZWwuZ2V0QXR0cmlidXRlKFwiaHgtbWF0Y2gtcXAtZGVmYXVsdFwiKTtcbiAgICAgIGlmIChkZWZhdWx0S2V5KSB7XG4gICAgICAgIGh0bXguc3dhcChcbiAgICAgICAgICBlbCxcbiAgICAgICAgICBlbC5nZXRBdHRyaWJ1dGUoXCJoeC1tYXRjaC1xcC1tYXBwaW5nOlwiICsgZGVmYXVsdEtleSkgPz8gXCJcIixcbiAgICAgICAgICB7IHN3YXBTdHlsZTogXCJpbm5lckhUTUxcIiwgc3dhcERlbGF5OiAwLCBzZXR0bGVEZWxheTogMCB9LFxuICAgICAgICApO1xuICAgICAgfVxuICAgIH1cbiAgfSk7XG59XG4iXX0=