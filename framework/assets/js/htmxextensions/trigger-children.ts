import htmx from "htmx.org";

function kebabEventName(str: string) {
    return str.replace(/([a-z0-9])([A-Z])/g, '$1-$2').toLowerCase()
}

const ignoredEvents = ['htmx:beforeProcessNode', 'htmx:afterProcessNode', 'htmx:configRequest', 'htmx:configResponse', 'htmx:responseError'];

function makeEvent(eventName: string, detail: any) {
    let evt
    if (window.CustomEvent && typeof window.CustomEvent === 'function') {
        // TODO: `composed: true` here is a hack to make global event handlers work with events in shadow DOM
        evt = new CustomEvent(eventName, { bubbles: false, cancelable: true, composed: true, detail })
    } else {
        evt = document.createEvent('CustomEvent')
        evt.initCustomEvent(eventName, true, true, detail)
    }
    return evt
}

function triggerChildren(target: HTMLElement, name: string, event: CustomEvent, triggered: Set<HTMLElement>) {
    if(ignoredEvents.includes(name)) {
        return
    }
    if (target && target.children) {
        Array.from(target.children).forEach((e) => {
            const kehab = kebabEventName(name);
            const eventName = kehab.replace("htmx:", "hx-on::")
            if (!triggered.has(e as HTMLElement)) {
                if(e.hasAttribute(eventName)) {
                    setTimeout(() => {
                        const newEvent = makeEvent(eventName.replace("hx-on::", "htmx:"), {
                            ...event.detail,
                            target: e,
                        })
                        newEvent.detail.meta = 'trigger-children'
                        e.dispatchEvent(newEvent)
                        triggered.add(e as HTMLElement);
                    }, 1)
                }
                if (e.children) {
                    triggerChildren(e as HTMLElement, name, event, triggered);
                }
            }
        });
    }
}


// @ts-ignore
htmx.defineExtension("trigger-children", {
    onEvent: (name, evt: Event | CustomEvent) => {
        if (!(evt instanceof CustomEvent)) {
            return false;
        }
        if(evt.detail.meta === 'trigger-children') {
            return false;
        }
        const triggered = new Set<HTMLElement>();
        const target = evt.target as HTMLElement || evt.detail.target as HTMLElement;
        triggerChildren(target, name, evt, triggered);
        return true;
    },
});
