import ws from 'k6/ws';
import { check } from 'k6';
import { sleep } from 'k6';

export let options = {
    vus: 1000,           // Number of users
    duration: '10m',     // Total duration of the test
};



function generateRandomString(length) {
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * characters.length));
    }
    return result;
}

export default function () {
    let sessionId = generateRandomString(20);
    const url = 'ws://localhost:3000/ws?sessionId=' + sessionId;

    const response = ws.connect(url, null, function (socket) {
        // Check if the connection was successful
        check(socket, { 'connected successfully': (s) => s.readyState === 1 });

        // Event listener for when the WebSocket is open
        socket.on('open', () => {
            console.log('Connected to WebSocket');
            while (true) {
                socket.send(JSON.stringify({ message: 'ping' }));
                sleep(1); // Wait for 1 second before sending the next message
            }
        });


        // Event listener for incoming messages
        socket.on('message', (msg) => {
            console.log(`Received message: ${msg}`);
        });

        // Event listener for errors
        socket.on('error', (e) => {
            console.error('Error in WebSocket connection:', e);
        });

        // Event listener for close
        socket.on('close', () => {
            console.log('Disconnected from WebSocket');
        });
    });

    // Check if the WebSocket handshake was successful
    check(response, { 'status is 101 (Switching Protocols)': (r) => r && r.status === 101 });

}
