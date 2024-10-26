import htmx from 'htmx.org'
import {removeAssociatedScripts} from "./htmgo";

let api : any = null;
let processed = new Set<string>()

htmx.defineExtension("sse", {
    init: function (apiRef) {
        api = apiRef;
    },
    // @ts-ignore
    onEvent: function (name, evt) {
        const target = evt.target;
        if(!(target instanceof HTMLElement)) {
            return
        }

        if(name === 'htmx:beforeCleanupElement') {
            removeAssociatedScripts(target);
        }

        if(name === 'htmx:beforeProcessNode') {
            const elements = document.querySelectorAll('[sse-connect]');
            for (let element of Array.from(elements)) {
                const url = element.getAttribute("sse-connect")!;
                if(url && !processed.has(url)) {
                    connectEventSource(element, url)
                    processed.add(url)
                }
            }
        }
    }
})

function connectEventSource(ele: Element, url: string) {
    if(!url) {
        return
    }
    console.info('Connecting to EventSource', url)
    const eventSource = new EventSource(url);

    eventSource.addEventListener("close", function(event) {
        htmx.trigger(ele, "htmx:sseClose", {event: event});
    })

    eventSource.onopen = function(event) {
        htmx.trigger(ele, "htmx:sseOpen", {event: event});
    }

    eventSource.onerror = function(event) {
        htmx.trigger(ele, "htmx:sseError", {event: event});
        if (eventSource.readyState == EventSource.CLOSED) {
            htmx.trigger(ele, "htmx:sseClose", {event: event});
        }
    }

    eventSource.onmessage = function(event) {
        const settleInfo = api.makeSettleInfo(ele);
        htmx.trigger(ele, "htmx:sseBeforeMessage", {event: event});
        const response = event.data
        const fragment = api.makeFragment(response) as DocumentFragment;
        const children = Array.from(fragment.children);
        for (let child of children) {
            api.oobSwap(api.getAttributeValue(child, 'hx-swap-oob') || 'true', child, settleInfo);
            // support htmgo eval__ scripts
            if(child.tagName === 'SCRIPT' && child.id.startsWith("__eval")) {
                document.body.appendChild(child);
            }
        }
        htmx.trigger(ele, "htmx:sseAfterMessage", {event: event});
    }
}
