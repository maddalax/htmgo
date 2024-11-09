import {ws} from "./ws";

window.onload = function () {
    const elements = document.querySelectorAll("[hx-extension]");
    for (let element of Array.from(elements)) {
        const value = element.getAttribute("hx-extension");
        if(value != null && value.split(" ").includes("ws")) {
            addWsEventHandlers()
            break;
        }
    }
};

function sendWs(message: Record<string, any>) {
    if(ws != null && ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify(message));
    }
}

function walk(node: Node, cb: (node: Node) => void) {
    cb(node);
    for (let child of Array.from(node.childNodes)) {
        walk(child, cb);
    }
}

export function addWsEventHandlers() {
    const observer = new MutationObserver(register)
    observer.observe(document.body, {childList: true, subtree: true})

    let added = new Set<string>();

    function register(mutations: MutationRecord[]) {
        for (let mutation of mutations) {
            for (let removedNode of Array.from(mutation.removedNodes)) {
                walk(removedNode, (node) => {
                    if (node instanceof HTMLElement) {
                        const handlerId = node.getAttribute("data-handler-id")
                        if(handlerId) {
                            added.delete(handlerId)
                            sendWs({id: handlerId, event: 'dom-element-removed'})
                        }
                    }
                })
            }

        }


        let ids = new Set<string>();
        document.querySelectorAll("[data-handler-id]").forEach(element => {
            const id = element.getAttribute("data-handler-id");
            const event = element.getAttribute("data-handler-event");

            if(id == null || event == null) {
                return;
            }

            ids.add(id);
            if (added.has(id)) {
                return;
            }
            added.add(id);
            element.addEventListener(event, (e) => {
                sendWs({id, event})
            });
        })
        for (let id of added) {
            if (!ids.has(id)) {
                added.delete(id);
            }
        }
    }

    register([])
}

