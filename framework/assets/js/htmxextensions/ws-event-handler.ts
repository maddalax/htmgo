import {ws} from "./ws";

window.onload = addWsEventHandlers;

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
    console.log('add ws event handlers')
    const observer = new MutationObserver(register)
    observer.observe(document.body, {childList: true, subtree: true})

    let added = new Set<string>();

    function register(mutations: MutationRecord[]) {
        console.log(mutations)

        for (let mutation of mutations) {
            for (let removedNode of Array.from(mutation.removedNodes)) {
                walk(removedNode, (node) => {
                    if (node instanceof HTMLElement) {
                        console.log('removing', node.innerHTML)
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
                console.debug('already added, skipping', id)
                return;
            }
            added.add(id);
            console.debug('adding event listener for ws send', id, event)
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

