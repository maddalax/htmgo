type WsOpts = {
    url: string;
    reconnectInterval?: number;
    onOpen?: () => void;
    onMessage: (message: string) => void;
    onError?: (error: Event) => void;
    onClose?: () => void;
}

export function createWebSocketClient(opts: WsOpts) {
    let socket: WebSocket | null = null;
    const connect = (tries: number) => {
        console.log('connecting to ws', opts.url, 'attempt', tries)
        socket = new WebSocket(opts.url);
        // Handle incoming messages
        socket.onmessage = (event) => {
            opts.onMessage(event.data)
        };
        // Handle connection errors
        socket.onerror = (error) => {
           try {
               socket?.close()
           } catch(ex) {
               // noop
           }
           socket = null
            let interval = tries * (opts.reconnectInterval || 1000);
            setTimeout(() => connect(tries + 1), interval);
        };
        // Handle connection close and attempt reconnection
        socket.onclose = () => {
            socket = null;
            let interval = tries * (opts.reconnectInterval || 1000);
            setTimeout(() => connect(tries + 1), interval);
        };
    };
    connect(1);
    const sendMessage = (message: string) => {
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(message);
        } else {
           setTimeout(() => sendMessage(message), 100);
        }
    };
    return { sendMessage };
}