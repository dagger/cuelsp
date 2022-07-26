# How CUELSP loads CUE

## Theory

CUELSP works with all CUE configurations, but it really shines when using CUE to configure [Dagger](https://dagger.io).

One big challenge of CUE is to know how load the current file. There are 2 ways:

- As a directory
- As a simple file

As you would split your Go package for organization purposes, you can write multiple CUE files in the same directory and
then merge everything at evaluation.
**BUT** with Dagger, you can also create multiples plans in a same directory that must be evaluated one by one.  
As an example, see our [docker tests](https://github.com/dagger/dagger/tree/main/pkg/universe.dagger.io/docker/test).

```
tree -P "*.cue" -I "testdata"
.
├── build.cue
├── dockerfile.cue
├── image.cue
└── run.cue
```

:bulb: A Dagger _plan_ is a file that contains a `dagger.#Plan` definition. For more information,
see the [official documentation](https://docs.dagger.io/1202/plan).

At the same time, CUELSP must work when you want to develop new definitions that are not part of a `dagger.#Plan`, for
instance if you want to create your own [reusable Dagger action](https://docs.dagger.io/1239/making-reusable-package).

So we have a problem : how to load the plan?

## Practice

First, it is not possible to just load the content of the opened file: if there are imported definition inside, we would
not be able to load the corresponding CUE value. It is the same for keys/definitions in other files of the directory. So
it's mandatory to **load from user's filesystem**.

Then it's straightforward, if there is something from an external file, trying to load plan as a file will fail.

So what we do is simple :

![load flow chart](../.github/assets/load-flow.png)

## :warning: Limitation

Some plans can work alone even if they are supposed to be merged during usage.  
That should not affect usability of CUELSP for now, but if it does, feel free to open an issue, so we can prioritize this
problem
