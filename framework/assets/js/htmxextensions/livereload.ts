import htmx from "htmx.org";
import {createWebSocketClient} from "../util/ws";

let lastVersion = "";

htmx.defineExtension("livereload", {
    init: function () {
        const host = window.location.host;

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

        createWebSocketClient({
            url: `${window.location.protocol === 'https:' ? 'wss' : 'ws'}://${host}/dev/livereload`,
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
    window.location.reload()
    // fetch(window.location.href).then(response => {
    //     return response.text();
    // }).then(html => {
    //    document.open();
    //     document.write(html);
    //     document.close();
    // }).catch(err => {
    //     console.log('failed to fetch live reload', err)
    //     setTimeout(reload, 100)
    // })
}