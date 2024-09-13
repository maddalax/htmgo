import htmx from "htmx.org";
const config: any = htmx.config;

/** @type {import("../htmx").HtmxInternalApi} */
let api: any;

const attrPrefix = "hx-target-";

// IE11 doesn't support string.startsWith
function startsWith(str: string, prefix: string) {
  return str.substring(0, prefix.length) === prefix;
}

/**
 * @param {HTMLElement} elt
 * @param respCodeNumber
 * @returns {HTMLElement | null}
 */
function getRespCodeTarget(elt: Element, respCodeNumber: number) {
  if (!elt || !respCodeNumber) return null;

  const respCode = respCodeNumber.toString();

  // '*' is the original syntax, as the obvious character for a wildcard.
  // The 'x' alternative was added for maximum compatibility with HTML
  // templating engines, due to ambiguity around which characters are
  // supported in HTML attributes.
  //
  // Start with the most specific possible attribute and generalize from
  // there.
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
    "xxx",
  ];
  if (startsWith(respCode, "4") || startsWith(respCode, "5")) {
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

/** @param {Event} evt */
function handleErrorFlag(evt: CustomEvent) {
  if (evt.detail.isError) {
    if (config.responseTargetUnsetsError) {
      evt.detail.isError = false;
    }
  } else if (config.responseTargetSetsError) {
    evt.detail.isError = true;
  }
}

htmx.defineExtension("response-targets", {
  // @ts-ignore
  init: (apiRef) => {
    api = apiRef;

    if (config.responseTargetUnsetsError === undefined) {
      config.responseTargetUnsetsError = true;
    }
    if (config.responseTargetSetsError === undefined) {
      config.responseTargetSetsError = false;
    }
    if (config.responseTargetPrefersExisting === undefined) {
      config.responseTargetPrefersExisting = false;
    }
    if (config.responseTargetPrefersRetargetHeader === undefined) {
      config.responseTargetPrefersRetargetHeader = true;
    }
  },

  // @ts-ignore
  onEvent: (name, evt) => {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    if (
      name === "htmx:beforeSwap" &&
      evt.detail.xhr &&
      evt.detail.xhr.status !== 200
    ) {
      if (evt.detail.target) {
        if (config.responseTargetPrefersExisting) {
          evt.detail.shouldSwap = true;
          handleErrorFlag(evt);
          return true;
        }
        if (
          config.responseTargetPrefersRetargetHeader &&
          evt.detail.xhr.getAllResponseHeaders().match(/HX-Retarget:/i)
        ) {
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
        evt.detail.xhr.status,
      );
      if (target) {
        handleErrorFlag(evt);
        evt.detail.shouldSwap = true;
        evt.detail.target = target;
      }
      return true;
    }
  },
});
