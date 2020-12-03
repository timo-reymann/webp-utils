Sample using docker in combination with alias
===

First define a alias e. g. in your `.bash_aliases`
```bash
alias webp-convert="docker run -v \$PWD:/workspace --rm -it timoreymann/webp-utils cwebp --config /etc/webp-utils/default.json --file-glob "
```

Now you can use it like this:
```bash
webp-convert *.png
```
