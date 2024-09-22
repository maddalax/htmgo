import htmx from "htmx.org";
import Idiomorph from './../idiomorph.js'

// @ts-ignore
htmx.defineExtension('diffdom-swap', {
    isInlineSwap: function(swapStyle) {
        return swapStyle === 'diffdom'
    },
    // @ts-ignore
    handleSwap: function(swapStyle: string, target: any, fragment: any) {
        if (swapStyle === 'diffdom') {
            if (fragment.nodeType === Node.DOCUMENT_FRAGMENT_NODE) {
                console.log('diffdom swap', target, fragment)
                console.log(Idiomorph.morph(target, fragment.firstElementChild))
                return [target]
            } else {
                console.log('diffdom swap', target, fragment)
                console.log(Idiomorph.morph(target, fragment.outerHTML))
                return [target]
            }
        }
    }
})