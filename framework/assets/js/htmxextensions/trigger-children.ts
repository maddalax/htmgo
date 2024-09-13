import htmx, { HtmxSettleInfo, HtmxSwapStyle } from "htmx.org";

htmx.defineExtension("trigger-children", {
  onEvent: (name, evt: Event | CustomEvent) => {
    if (!(evt instanceof CustomEvent)) {
      return false;
    }
    const target = evt.detail.target as HTMLElement;
    if (target && target.children) {
      Array.from(target.children).forEach((e) => {
        htmx.trigger(e, name, null);
      });
    }
    return true;
  },
  init: function (api: any): void {},
  transformResponse: function (
    text: string,
    xhr: XMLHttpRequest,
    elt: Element,
  ): string {
    return text;
  },
  isInlineSwap: function (swapStyle: HtmxSwapStyle): boolean {
    return false;
  },
  handleSwap: function (
    swapStyle: HtmxSwapStyle,
    target: Node,
    fragment: Node,
    settleInfo: HtmxSettleInfo,
  ): boolean | Node[] {
    return false;
  },
  encodeParameters: function (
    xhr: XMLHttpRequest,
    parameters: FormData,
    elt: Node,
  ) {},
  getSelectors: function (): string[] | null {
    return null;
  },
});
