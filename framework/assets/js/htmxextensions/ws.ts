import htmx from 'htmx.org'
import {removeAssociatedScripts} from "./htmgo";

let api : any = null;
let processed = new Set<string>()
export let ws: WebSocket | null = null;

htmx.defineExtension("ws", {
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
            const elements = document.querySelectorAll('[ws-connect]');
            for (let element of Array.from(elements)) {
                const url = element.getAttribute("ws-connect")!;
                if(url && !processed.has(url)) {
                    connectWs(element, url)
                    processed.add(url)
                }
            }
        }
    }
})

function exponentialBackoff(attempt: number, baseDelay = 100, maxDelay = 10000) {
    // Exponential backoff: baseDelay * (2 ^ attempt) with jitter
    const jitter = Math.random(); // Adding randomness to prevent collisions
    return Math.min(baseDelay * Math.pow(2, attempt) * jitter, maxDelay);
}

function connectWs(ele: Element, url: string, attempt: number = 0) {
    if(!url) {
        return
    }
    if(!url.startsWith('ws://') && !url.startsWith('wss://')) {
        const isSecure = window.location.protocol === 'https:'
        url = (isSecure ? 'wss://' : 'ws://') + window.location.host + url
    }
    console.info('connecting to ws', url)

    ws = new WebSocket(url);

    ws.addEventListener("close", function(event) {
        htmx.trigger(ele, "htmx:wsClose", {event: event});
        const delay = exponentialBackoff(attempt);
        setTimeout(() => {
            connectWs(ele, url, attempt + 1)
        }, delay)
    })

    ws.addEventListener("open", function(event) {
        htmx.trigger(ele, "htmx:wsOpen", {event: event});
    })

    ws.addEventListener("error", function(event) {
        htmx.trigger(ele, "htmx:wsError", {event: event});
    })

    ws.addEventListener("message", function(event) {
        const settleInfo = api.makeSettleInfo(ele);
        htmx.trigger(ele, "htmx:wsBeforeMessage", {event: event});
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
        htmx.trigger(ele, "htmx:wsAfterMessage", {event: event});
    })

    return ws
}
