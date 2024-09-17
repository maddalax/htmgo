import htmx from "htmx.org";
import {createWebSocketClient} from "../util/ws";

let lastVersion = "";

htmx.defineExtension("livereload", {
    init: function () {
        const host = window.location.host;
        console.log('livereload extension initialized.');
        createWebSocketClient({
            url: `ws://${host}/dev/livereload`,
            onOpen: () => {
                console.log('LiveReload connected.');
            },
            onMessage: (message) => {
               if(lastVersion === "") {
                   lastVersion = message;
               }
               if(lastVersion !== message) {
                 window.location.reload();
               }
            },
            onError: (error) => {
            },
            onClose: () => {
            }
        })
    },
    // @ts-ignore
    onEvent: function (name, evt) {

    },
});