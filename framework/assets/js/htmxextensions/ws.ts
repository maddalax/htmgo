import htmx from 'htmx.org'
import {removeAssociatedScripts} from "./htmgo";


declare module 'htmx.org' {
    interface Htmx {
        defineExtension(name: string, extension: HtmxExtension): void;
        createWebSocket?: (url: string) => WebSocket;
        config: {
            wsReconnectDelay?: 'full-jitter' | ((retryCount: number) => number);
            wsBinaryType?: string;
            [key: string]: any
        };
        [key: string]: any;
    }
}

interface HtmxExtension {
    init: (apiRef: HtmxInternalApi) => void;
    onEvent: (name: string, evt: Event) => void;
    [key: string]: any;
}

interface HtmxInternalApi {
    getInternalData(elt: Element): any;
    bodyContains(elt: Element): boolean;
    getAttributeValue(elt: Element, name: string): string | null;
    triggerEvent(elt: Element, name: string, detail?: any): boolean;
    withExtensions(elt: Element, callback: (extension: any) => void): void;
    makeSettleInfo(elt: Element): any;
    makeFragment(html: string): DocumentFragment;
    oobSwap(swapStyle: string, fragment: Element, settleInfo: any): void;
    settleImmediately(tasks: any): void;
    getClosestMatch(elt: Element, condition: (node: Element) => boolean): Element | null;
    getTriggerSpecs(elt: Element): any[];
    addTriggerHandler(elt: Element, triggerSpec: any, nodeData: any, handler: (elt: Element, evt: Event) => void): void;
    getHeaders(elt: Element, target: Element): any;
    getTarget(elt: Element): Element;
    getInputValues(elt: Element, verb: string): { errors: any[]; values: any };
    getExpressionVars(elt: Element): any;
    mergeObjects(obj1: any, obj2: any): any;
    filterValues(values: any, elt: Element): any;
    triggerErrorEvent(elt?: Element, name?: string, detail?: any): void;
    hasAttribute(elt: Element, name: string): boolean;
    shouldCancel(evt: Event, elt: Element): boolean;
    [key: string]: any;
}

interface WebSocketWrapper {
    socket: WebSocket;
    events : { [key: string]: ((event: Event) => void)[] };
    messageQueue: { message: string; sendElt: Element | null }[];
    retryCount: number;
    sendImmediately(message: string, sendElt: Element | null): void;
    send(message: string, sendElt: Element | null): void;
    addEventListener(event: string, handler: (event: Event) => void): void;
    handleQueuedMessages(): void;
    init(): void;
    close(): void;
    publicInterface: {
        send: (message: string, sendElt: Element | null) => void;
        sendImmediately: (message: string, sendElt: Element | null) => void;
        queue: { message: string; sendElt: Element | null }[];
    };
}

let api: HtmxInternalApi;

function splitOnWhitespace(trigger: string): string[] {
    return trigger.trim().split(/\s+/);
}

function getLegacyWebsocketURL(elt: Element): string | undefined {
    const legacySSEValue = api.getAttributeValue(elt, 'hx-ws');
    if (legacySSEValue) {
        const values = splitOnWhitespace(legacySSEValue);
        for (let i = 0; i < values.length; i++) {
            const value = values[i].split(/:(.+)/);
            if (value[0] === 'connect') {
                return value[1];
            }
        }
    }
    return undefined;
}

function ensureWebSocket(socketElt: HTMLElement): void {
    // If the element containing the WebSocket connection no longer exists, then
    // do not connect/reconnect the WebSocket.
    if (!api.bodyContains(socketElt)) {
        return;
    }

    // Get the source straight from the element's value
    let wssSource = api.getAttributeValue(socketElt, 'ws-connect');

    if (wssSource == null || wssSource === '') {
        const legacySource = getLegacyWebsocketURL(socketElt);
        if (legacySource == null) {
            return;
        } else {
            wssSource = legacySource;
        }
    }

    // Guarantee that the wssSource value is a fully qualified URL
    if (wssSource.indexOf('/') === 0) {
        const base_part = location.hostname + (location.port ? ':' + location.port : '');
        if (location.protocol === 'https:') {
            wssSource = 'wss://' + base_part + wssSource;
        } else if (location.protocol === 'http:') {
            wssSource = 'ws://' + base_part + wssSource;
        }
    }

    const socketWrapper = createWebsocketWrapper(socketElt, () => htmx.createWebSocket!(wssSource));

    socketWrapper.addEventListener('message', (event) => {
        if (maybeCloseWebSocketSource(socketElt)) {
            return;
        }

        let response = (event as MessageEvent).data;
        if (
            !api.triggerEvent(socketElt, 'htmx:wsBeforeMessage', {
                message: response,
                socketWrapper: socketWrapper.publicInterface,
            })
        ) {
            return;
        }

        api.withExtensions(socketElt, (extension) => {
            response = extension.transformResponse(response, null, socketElt);
        });

        const settleInfo = api.makeSettleInfo(socketElt);
        const fragment = api.makeFragment(response);

        if (fragment.children.length) {
            const children = Array.from(fragment.children);
            for (let i = 0; i < children.length; i++) {
                const child = children[i]
                api.oobSwap(api.getAttributeValue(child, 'hx-swap-oob') || 'true', children[i], settleInfo);
                // support htmgo eval__ scripts
                if(child.tagName === 'SCRIPT' && child.id.startsWith("__eval")) {
                    document.body.appendChild(child);
                }
            }
        }

        api.settleImmediately(settleInfo.tasks);
        api.triggerEvent(socketElt, 'htmx:wsAfterMessage', {
            message: response,
            socketWrapper: socketWrapper.publicInterface,
        });
    });

    // Put the WebSocket into the HTML Element's custom data.
    api.getInternalData(socketElt).webSocket = socketWrapper;
}

function createWebsocketWrapper(socketElt: HTMLElement, socketFunc: () => WebSocket): WebSocketWrapper {
    const wrapper: WebSocketWrapper = {
        socket: null as unknown as WebSocket,
        messageQueue: [],
        retryCount: 0,
        events: {} as { [key: string]: ((event: Event) => void)[] },
        addEventListener(event: string, handler: (event: Event) => void) {
            if (this.socket) {
                this.socket.addEventListener(event, handler);
            }

            if (!this.events[event]) {
                this.events[event] = [];
            }

            this.events[event].push(handler);
        },
        sendImmediately(message: string, sendElt: Element | null) {
            if (!this.socket) {
                api.triggerErrorEvent(socketElt, 'htmx:wsError', { error: 'No socket available' });
            }
            if (
                !sendElt ||
                api.triggerEvent(sendElt, 'htmx:wsBeforeSend', {
                    message,
                    socketWrapper: this.publicInterface,
                })
            ) {
                this.socket.send(message);
                if (sendElt) {
                    api.triggerEvent(sendElt, 'htmx:wsAfterSend', {
                        message,
                        socketWrapper: this.publicInterface,
                    });
                }
            }
        },
        send(message: string, sendElt: Element | null) {
            if (this.socket.readyState !== this.socket.OPEN) {
                this.messageQueue.push({ message, sendElt });
            } else {
                this.sendImmediately(message, sendElt);
            }
        },
        handleQueuedMessages() {
            while (this.messageQueue.length > 0) {
                const queuedItem = this.messageQueue[0];
                if (this.socket.readyState === this.socket.OPEN) {
                    this.sendImmediately(queuedItem.message, queuedItem.sendElt);
                    this.messageQueue.shift();
                } else {
                    break;
                }
            }
        },
        init() {
            if (this.socket && this.socket.readyState === this.socket.OPEN) {
                // Close discarded socket
                this.socket.close();
            }

            // Create a new WebSocket and event handlers
            const socket = socketFunc();

            // The event.type detail is added for interface conformance with the
            // other two lifecycle events (open and close) so a single handler method
            // can handle them polymorphically, if required.
            api.triggerEvent(socketElt, 'htmx:wsConnecting', { event: { type: 'connecting' } });

            this.socket = socket;

            socket.onopen = (e) => {
                this.retryCount = 0;
                api.triggerEvent(socketElt, 'htmx:wsOpen', { event: e, socketWrapper: this.publicInterface });
                this.handleQueuedMessages();
            };

            socket.onclose = (e) => {
                // If socket should not be connected, stop further attempts to establish connection
                // If Abnormal Closure/Service Restart/Try Again Later, then set a timer to reconnect after a pause.
                if (!maybeCloseWebSocketSource(socketElt) && [1006, 1012, 1013].indexOf(e.code) >= 0) {
                    const delay = getWebSocketReconnectDelay(this.retryCount);
                    setTimeout(() => {
                        this.retryCount += 1;
                        this.init();
                    }, delay);
                }

                // Notify client code that connection has been closed. Client code can inspect `event` field
                // to determine whether closure has been valid or abnormal
                api.triggerEvent(socketElt, 'htmx:wsClose', { event: e, socketWrapper: this.publicInterface });
            };

            socket.onerror = (e) => {
                api.triggerErrorEvent(socketElt, 'htmx:wsError', { error: e, socketWrapper: this });
                maybeCloseWebSocketSource(socketElt);
            };

            const events = this.events;
            Object.keys(events).forEach((k) => {
                events[k].forEach((e) => {
                    socket.addEventListener(k, e);
                });
            });
        },
        close() {
            this.socket.close();
        },
        publicInterface: {} as any,
    };

    wrapper.init();

    wrapper.publicInterface = {
        send: wrapper.send.bind(wrapper),
        sendImmediately: wrapper.sendImmediately.bind(wrapper),
        queue: wrapper.messageQueue,
    };

    return wrapper;
}

function ensureWebSocketSend(elt: HTMLElement): void {
    const legacyAttribute = api.getAttributeValue(elt, 'hx-ws');
    if (legacyAttribute && legacyAttribute !== 'send') {
        return;
    }

    const webSocketParent = api.getClosestMatch(elt, hasWebSocket);
    if (webSocketParent) {
        processWebSocketSend(webSocketParent as HTMLElement, elt);
    }
}

function hasWebSocket(node: HTMLElement): boolean {
    return api.getInternalData(node).webSocket != null;
}

function processWebSocketSend(socketElt: HTMLElement, sendElt: HTMLElement): void {
    const nodeData = api.getInternalData(sendElt);
    const triggerSpecs = api.getTriggerSpecs(sendElt);
    triggerSpecs.forEach((ts) => {
        api.addTriggerHandler(sendElt, ts, nodeData, (elt: Element, evt: Event) => {
            if (maybeCloseWebSocketSource(socketElt)) {
                return;
            }

            const socketWrapper: WebSocketWrapper = api.getInternalData(socketElt).webSocket;
            const headers = api.getHeaders(sendElt, api.getTarget(sendElt));
            const results = api.getInputValues(sendElt, 'post');
            const errors = results.errors;
            const rawParameters = Object.assign({}, results.values);
            const expressionVars = api.getExpressionVars(sendElt);
            const allParameters = api.mergeObjects(rawParameters, expressionVars);
            const filteredParameters = api.filterValues(allParameters, sendElt);

            const sendConfig = {
                parameters: filteredParameters,
                unfilteredParameters: allParameters,
                headers,
                errors,

                triggeringEvent: evt,
                messageBody: undefined as string | undefined,
                socketWrapper: socketWrapper.publicInterface,
            };

            if (!api.triggerEvent(elt, 'htmx:wsConfigSend', sendConfig)) {
                return;
            }

            if (errors && errors.length > 0) {
                api.triggerEvent(elt, 'htmx:validation:halted', errors);
                return;
            }

            let body = sendConfig.messageBody;
            if (body === undefined) {
                const toSend = Object.assign({}, sendConfig.parameters);
                if (sendConfig.headers) {
                    toSend.HEADERS = headers;
                }
                body = JSON.stringify(toSend);
            }

            socketWrapper.send(body, elt as Element);

            if (evt && api.shouldCancel(evt, elt as Element)) {
                evt.preventDefault();
            }
        });
    });
}

function getWebSocketReconnectDelay(retryCount: number): number {
    const delay = htmx.config.wsReconnectDelay;
    if (typeof delay === 'function') {
        return delay(retryCount);
    }
    if (delay === 'full-jitter') {
        const exp = Math.min(retryCount, 6);
        const maxDelay = 1000 * Math.pow(2, exp);
        return maxDelay * Math.random();
    }
    return 0;
}

function maybeCloseWebSocketSource(elt: HTMLElement): boolean {
    if (!api.bodyContains(elt)) {
        api.getInternalData(elt).webSocket.close();
        return true;
    }
    return false;
}

function createWebSocket(url: string): WebSocket {
    const sock = new WebSocket(url, []);
    sock.binaryType = (htmx.config.wsBinaryType || 'blob') as unknown as BinaryType;
    return sock;
}

function queryAttributeOnThisOrChildren(elt: HTMLElement, attributeName: string): HTMLElement[] {
    const result: HTMLElement[] = [];

    // If the parent element also contains the requested attribute, then add it to the results too.
    if (api.hasAttribute(elt, attributeName) || api.hasAttribute(elt, 'hx-ws')) {
        result.push(elt);
    }

    // Search all child nodes that match the requested attribute
    elt.querySelectorAll('[' + attributeName + '], [data-' + attributeName + '], [data-hx-ws], [hx-ws]').forEach((node) => {
        result.push(node as HTMLElement);
    });

    return result;
}

function forEach<T>(arr: T[], func: (item: T) => void): void {
    if (arr) {
        arr.forEach(func);
    }
}

htmx.defineExtension('ws', {
    init: (apiRef: HtmxInternalApi) => {
        // Store reference to internal API
        api = apiRef;

        // Default function for creating new WebSocket objects
        if (!htmx.createWebSocket) {
            htmx.createWebSocket = createWebSocket;
        }

        // Default setting for reconnect delay
        if (!htmx.config.wsReconnectDelay) {
            htmx.config.wsReconnectDelay = 'full-jitter';
        }
    },

    onEvent: (name: string, evt: Event) => {
        const parent: Element = evt.target as Element || (evt as CustomEvent).detail.elt;

        if(!(parent instanceof HTMLElement)) {
            return
        }

        switch (name) {
            // Try to close the socket when elements are removed
            case 'htmx:beforeCleanupElement':
                removeAssociatedScripts(parent);
                const internalData = api.getInternalData(parent);
                if (internalData.webSocket) {
                    internalData.webSocket.close();
                }
                return;

            // Try to create websockets when elements are processed
            case 'htmx:beforeProcessNode':
                forEach(queryAttributeOnThisOrChildren(parent, 'ws-connect'), (child) => {
                    ensureWebSocket(child);
                });
                forEach(queryAttributeOnThisOrChildren(parent, 'ws-send'), (child) => {
                    ensureWebSocketSend(child);
                });
        }
    },
});
