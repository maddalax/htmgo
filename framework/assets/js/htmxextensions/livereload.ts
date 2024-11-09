import htmx from "htmx.org";
import {hasExtension} from "./extension";

let lastVersion = "";

htmx.defineExtension("livereload", {
    init: function () {

        let enabled = hasExtension("livereload")

        if(!enabled) {
            return
        }

        console.info('livereload extension initialized.');
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
