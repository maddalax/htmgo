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
            },
            onMessage: (message) => {
               if(lastVersion === "") {
                   lastVersion = message;
               }
               if(lastVersion !== message) {
                   lastVersion = message;
                   reload()
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

function reload() {
    fetch(window.location.href).then(response => {
        return response.text();
    }).then(html => {
        document.open();
        document.write(html);
        document.close();
    }).catch(err => {
        console.log('failed to fetch live reload', err)
        setTimeout(reload, 100)
    })
}