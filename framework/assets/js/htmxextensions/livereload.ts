import htmx from "htmx.org";
import {createWebSocketClient} from "../util/ws";

let lastVersion = "";

htmx.defineExtension("livereload", {
    init: function () {

        let enabled = false
        for (const element of Array.from(htmx.findAll("[hx-ext]"))) {
            const value = element.getAttribute("hx-ext");
            if(value?.split(" ").includes("livereload")) {
                enabled = true
                break;
            }
        }

        if(!enabled) {
            return
        }

        console.log('livereload extension initialized.');
        // Create a new EventSource object and point it to your SSE endpoint
        const eventSource = new EventSource('/dev/livereload');
        // Listen for messages from the server
        eventSource.onmessage = function(event) {
            const message = event.data
            // Log the message data received from the server
            if(lastVersion === "") {
                lastVersion = message;
            }
            if(lastVersion !== message) {
                lastVersion = message;
                reload()
            }
        };
        // Handle errors (e.g., when the connection is closed)
        eventSource.onerror = function(error) {
            console.error('EventSource error:', error);
        };

    },
    // @ts-ignore
    onEvent: function (name, evt) {

    },
});

function reload() {
    window.location.reload()
}