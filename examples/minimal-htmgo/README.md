Minimal example that just uses htmgo for html rendering / js support and nothing else.

Removes automatic support for:

1. live reloading
2. tailwind recompilation
3. page/partial route registration
4. Single binary (since /public/ assets is required to be there), normally htmgo uses the embedded file system in other examples such as https://github.com/maddalax/htmgo/blob/master/templates/starter/assets_prod.go
