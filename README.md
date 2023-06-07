# rn
File renamer: give your files a predictable, descriptive name.

## Format
```
DATE--TITLE__AUTHORS_TAGS.extension
```

## Examples
```sh
rn README.md
# output: README.md -> 20060102T150405--readme.md

rn README.md technical computing manual
# output: README.md -> 20060102T150405--readme__technical_computing_manual.md

rn README.md title="Don't Readme!" technical computing manual
# output: README.md -> 20060102T150405--dont-readme__technical_computing_manual.md

rn README.md title="Don't README!" name="JessebotX" technical computing manual
# output: README.md -> 20060102T150405--dont-readme__jessebotx_technical_computing_manual.md
```
