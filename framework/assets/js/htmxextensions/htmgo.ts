import htmx from "htmx.org";

const evalFuncRegex =/__eval_[A-Za-z0-9]+\([a-z]+\)/gm

htmx.defineExtension("htmgo", {
    // @ts-ignore
    onEvent: function (name, evt) {
       if(name === "htmx:beforeCleanupElement" && evt.target) {
           removeAssociatedScripts(evt.target as HTMLElement);
       }
       if(name === "htmx:load" && evt.target) {
          invokeOnLoad(evt.target as HTMLElement);
       }
    },
});

/**
 * Browser doesn't support onload for all elements, so we need to manually trigger it
 * this is useful for locality of behavior
 */
function invokeOnLoad(element : Element) {
    if(element == null || !(element instanceof HTMLElement)) {
        return
    }
    const ignored = ['SCRIPT', 'LINK', 'STYLE', 'META', 'BASE', 'TITLE', 'HEAD', 'HTML', 'BODY'];
    if(!ignored.includes(element.tagName)) {
        if(element.hasAttribute("onload")) {
            element.onload!(new Event("load"));
        }
    }
    // check its children
    element.querySelectorAll('[onload]').forEach(invokeOnLoad)
}

export function removeAssociatedScripts(element: HTMLElement) {
   const attributes = Array.from(element.attributes)
    for (let attribute of attributes) {
       const matches = attribute.value.match(evalFuncRegex) || []
        for (let match of matches) {
            const id = match.replace("()", "").replace("(this)", "").replace(";", "")
            const ele = document.getElementById(id)
            if(ele && ele.tagName === "SCRIPT") {
                console.debug("removing associated script with id", id)
                ele.remove()
            }
        }
    }
}
