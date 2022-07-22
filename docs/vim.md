## 1/4. Install `dagger` >= 0.2.26

Dagger LSP is a subcommand on the Dagger CLI: `dagger lsp`

Ensure that you have `dagger` **v0.2.26** or newer installed.
[This is how you can install `dagger`](https://docs.dagger.io/install).

To check what version you have installed, run: `dagger version`

As a last step, check the version of `dagger` that vim has access to by running the following command in vim: `!dagger version`

## 2/4. Install `vim-lsp`

Refer to [vim-lsp documentation](https://github.com/prabirshrestha/vim-lsp#installing).

## 3/4. Configure `.vimrc`

Add this to your `.vimrc`:

```vim
if executable('dagger')
  au User lsp_setup call lsp#register_server({
      \ 'name': 'dagger_cue_lsp',
      \ 'cmd': {server_info->['dagger', 'lsp']},
      \ 'allowlist': ['cue'],
      \ })
endif

function! s:on_lsp_buffer_enabled() abort
    setlocal omnifunc=lsp#complete
    setlocal signcolumn=yes
    if exists('+tagfunc') | setlocal tagfunc=lsp#tagfunc | endif
    nmap <buffer> gd <plug>(lsp-definition)
    " nmap <buffer> gs <plug>(lsp-document-symbol-search)
    " nmap <buffer> gS <plug>(lsp-workspace-symbol-search)
    " nmap <buffer> gr <plug>(lsp-references)
    " nmap <buffer> gi <plug>(lsp-implementation)
    " nmap <buffer> gt <plug>(lsp-type-definition)
    " nmap <buffer> <leader>rn <plug>(lsp-rename)
    " nmap <buffer> [g <plug>(lsp-previous-diagnostic)
    " nmap <buffer> ]g <plug>(lsp-next-diagnostic)
    nmap <buffer> K <plug>(lsp-hover)
    nnoremap <buffer> <expr>d lsp#scroll(+4)
    nnoremap <buffer> <expr>u lsp#scroll(-4)

    let g:lsp_format_sync_timeout = 1000

    " refer to vim-lsp doc to add more commands
endfunction

augroup lsp_install
  au!
  " call s:on_lsp_buffer_enabled only for languages that has the server registered.
  autocmd User lsp_buffer_enabled call s:on_lsp_buffer_enabled()
augroup END
```

## 4/4. Use `vim-lsp` with `dagger`

Given a local copy of https://github.com/dagger/todoapp, open up `dagger.cue` in vim.

Place the cursor over a definition and type `CTRL ]` to jump to the implementation.
Type `CTRL t` to jump back.

You can also type `K` to open the definition description in a hover area.
If there is more content than fits the hover area, press `d` to go scroll down and `u` to scroll up.
This is what that looks like in practice:

![](./vim-dagger-lsp.gif)
