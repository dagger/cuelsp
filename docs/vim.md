# Dagger CUE LSP in vim

You can use this LSP in vim thanks to vim plugins handling LSP.
It was tested in [vim-lsp](https://github.com/prabirshrestha/vim-lsp)

## Dagger CUE LSP

### Install

Refer to [this documentation](/README.md#install)

## vim-lsp

### Install

Refer to [this documentation](https://github.com/prabirshrestha/vim-lsp#installing)

### Configure

Add this to your `.vimrc`
```
if executable('daggerlsp')
    au User lsp_setup call lsp#register_server({
        \ 'name': 'daggerlsp',
        \ 'cmd': {server_info->['daggerlsp']},
        \ 'allowlist': ['cue'],
        \ })
endif

function! s:on_lsp_buffer_enabled() abort
    nmap <buffer> gd <plug>(lsp-definition)
endfunction

augroup lsp_install
    au!
    " call s:on_lsp_buffer_enabled only for languages that has the server registered.
    autocmd User lsp_buffer_enabled call s:on_lsp_buffer_enabled()
augroup END
```
