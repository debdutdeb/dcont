# dcont
*D*evelopment *Cont*ainers :/

this is an example of how much I want to **NOT** use vscode

## Quick note 

Start the language server with the correct mounts
```sh
docker container run \
    --platform=linux/amd64 \
    -v $HOME/go:$HOME/go \
    -v $HOME/Documents/Repos/rocketchatctl:$HOME/Documents/Repos/rocketchatctl \
    -p 3001:3001 \
    -d \
    --name gopls \
    --rm \
    lspcontainers/gopls:latest gopls -vv --port 3001
```

Start neovim
```sh
nvim --cmd 'lua goplsHost = "127.0.0.1"; goplsPort = 3001'
```

neovim gopls lspconfig
```lua
local lsp_handlers = require("neoconfig.lsp.handlers")
local goplsOpts = {
    on_attach = lsp_handlers.on_attach,
    capabilities = lsp_handlers.capabilities,
}
if goplsHost ~= nil and goplsPort ~= nil then
    goplsOpts.cmd = vim.lsp.rpc.connect(goplsHost, goplsPort)
end
lspconfig.gopls.setup(goplsOpts)
```
