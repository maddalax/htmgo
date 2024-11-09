
export function hasExtension(name: string):  boolean {
    for (const element of Array.from(document.querySelectorAll("[hx-ext]"))) {
        const value = element.getAttribute("hx-ext");
        if(value != null) {
            const split = value.split(" ").map(s => s.replace(",", ""))
            if(split.includes(name)) {
                return true;
            }
        }
    }
    return false;
}
