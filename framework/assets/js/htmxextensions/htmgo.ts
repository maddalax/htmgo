import htmx from "htmx.org";

const evalFuncRegex =/__eval_[A-Za-z0-9]+\([a-z]+\)/gm

htmx.defineExtension("htmgo", {
    // @ts-ignore
    onEvent: function (name, evt) {
       if(name === "htmx:beforeCleanupElement" && evt.target) {
           removeAssociatedScripts(evt.target as HTMLElement);
       }
    },
});

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
