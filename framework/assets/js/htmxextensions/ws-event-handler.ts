import {ws} from "./ws";

window.onload = addWsEventHandlers;

export function addWsEventHandlers() {
    console.log('add ws event handlers')
    const observer = new MutationObserver(register)
    observer.observe(document.body, {childList: true, subtree: true})

    let added = new Set<string>();

    function register() {
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
                console.log('sending event', id, event, ws)
                if(ws != null && ws.readyState === WebSocket.OPEN) {
                    ws.send(JSON.stringify({id, event}));
                }
            });
        })
        for (let id of added) {
            if (!ids.has(id)) {
                added.delete(id);
            }
        }
        console.log('size', added.size)
    }

    register()
}

