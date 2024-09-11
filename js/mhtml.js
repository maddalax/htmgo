window.onload = function () {
    // htmx.logger = function(elt, event, data) {
    //     if(console) {
    //         console.log(event);
    //     }
    // }
    // onUrlChange(window.location.href);


    window.history.pushState = new Proxy(window.history.pushState, {
        apply: (target, thisArg, argArray) => {
            if(argArray.length > 2) {
                onUrlChange(window.location.origin + argArray[2]);
            }
            return target.apply(thisArg, argArray);
        },
    });
}

function onUrlChange(newUrl) {
    let url = new URL(newUrl);

    setTimeout(() => {
        document.querySelectorAll('[hx-trigger]').forEach(function(element) {
            const triggers = element.getAttribute('hx-trigger');
            const split = triggers.split(", ");
            console.log(split)
            if(split.find(s => s === 'url')) {
                htmx.trigger(element, "url");
            } else {
                for (let [key, values] of url.searchParams) {
                    let eventName = "qs:" + key
                    if (triggers.includes(eventName)) {
                        htmx.trigger(element, eventName);
                        break
                    }
                }
            }
        });
    }, 50)

    document.querySelectorAll('[hx-match-qp]').forEach((el) => {
        let hasMatch = false;
        for (let name of el.getAttributeNames()) {
            if(name.startsWith("hx-match-qp-mapping:")) {
               let match = name.replace("hx-match-qp-mapping:", "");
               let value = url.searchParams.get(match);
               if(value) {
                   htmx.swap(el, el.getAttribute(name), {swapStyle: 'innerHTML'})
                   hasMatch = true;
                   break
               }
            }
        }
        if(!hasMatch) {
            let defaultKey = el.getAttribute("hx-match-qp-default")
            if(defaultKey) {
                htmx.swap(el, el.getAttribute("hx-match-qp-mapping:" + defaultKey), {swapStyle: 'innerHTML'})
            }
        }
    })
}