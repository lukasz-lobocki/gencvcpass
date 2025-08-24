# gencvcpass ![Static](https://img.shields.io/badge/flota-pani-darkcyan?style=for-the-badge&labelColor=lightsalmon)

Generates consonant-vowel-consonant patterned password.

**zeqBif-xa4sas-Zadkut-xufbim-ra2tel**

## Use

```text
Flags:
  -s, --sets int      number of sets between separators (default 5)
  -u, --upper int     number of uppercase letters (default 2)
  -d, --digits int    number of digits (default 2)
      --sep string    separator character (default "-")
```

- Uppercase letters are generated as first character of a random *cvc* triplets.
- Digits are generated as last character of a random *cvc* triplets.

## Build

See [BUILD.md](BUILD.md) file.

## License

`gencvcpass` was created by Lukasz Lobocki. It is licensed under the terms of the CC0 v1.0 Universal license.

All components used retain their original licenses.

## Credits

`gencvcpass` was created with [cookiecutter](https://cookiecutter.readthedocs.io/en/latest/) and [template](https://github.com/lukasz-lobocki/go-cookiecutter).
