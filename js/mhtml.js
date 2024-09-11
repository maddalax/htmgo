// Replace 'ws://example.com/socket' with the URL of your WebSocket server

window.onload = function () {
    document.querySelectorAll('[m\\:onclick]').forEach(element => {
        const value = element.getAttribute('m:onclick')
        element.addEventListener('click', () => {
            fetch('/click/' + value).catch()
        })
    });
}