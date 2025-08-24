# gencvcpass

## Build a snapshot

To build the solution with dirty repository, use the following command with `--snapshot` parameter.

```bash
goreleaser build --clean --snapshot
```

## Typical release workflow

```bash
git add --update
```

```bash
git commit -m "fix: Change."
```

```bash
git tag "$(svu next --always)"
git push --tags
goreleaser release --clean
```

## Cookiecutter initiation

```bash
cookiecutter \
  ssh://git@github.com/lukasz-lobocki/go-cookiecutter.git \
  package_name="gencvcpass"
```

### was run with following variables

- package_name: **`gencvcpass`**;
package_short_description: `Generate cvc password`

- package_version: `0.1.0`

- author_name: `Lukasz Lobocki`;
open_source_license: `CC0 v1.0 Universal`

- __package_slug: `gencvcpass`

### on

`2025-08-24 07:54:14 +0200`
