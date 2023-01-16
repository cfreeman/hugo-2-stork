# hugo-2-stork

hugo-2-stork is an application that generates a stork config file from a hugo static website.

[Stork search](https://stork-search.net) is a library for creating search interfaces on static websites. While [HUGO](https://gohugo.io) is a framework for creating static websites.

Installing hugo-2-stork:
```bash
$GOBIN=/bin go install github.com/cfreeman/hugo-2-stork@latest
```

```bash
$hugo-2-stork --src=content/posts --url=https://myweb.com
$stork build --input stork.toml --output index.st
```

## TODO
* Support multiple source folders.
* ~~Need to make sure the YAML frontmatter block is indexed.~~
* ~~More testing of output.~~
* ~~ability to configure URL stem.~~
* ~~ability to configure destination of output stork file.~~