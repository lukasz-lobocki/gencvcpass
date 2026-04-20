# gencvcpass ![Static](https://img.shields.io/badge/flota-pani-darkcyan?style=for-the-badge&labelColor=lightsalmon)

Generates consonant-vowel-consonant-vowel... patterned password. Spiced up with uppercase letters and digits.

**saneFun3-fitolute-sor8fyba-pygoFali**

## Use

```text
Flags:
  -s, --sets int          number of sets between separators (default 4)
  -u, --upper int         number of uppercase letters (default 2)
  -d, --digits int        number of digits (default 2)
      --sep string        separator character (default "-")
  -l, --less-non-polish   do not use non-polish consonants
```

- Uppercase letters are generated as first character of random *cvcv* multilets.
- Digits are generated as last character of random *cvcv* multilets.

## Gemini's verdict

Assessment for the password: `saneFun3-fitolute-sor8fyba-pygoFali`

**Overall Strength: Exceptional**

This password is an excellent example of a passphrase, which balances high security with relative human memorability. It is significantly stronger than the vast majority of passwords currently in use.

**Detailed Breakdown**

* **Length (35 Characters):** Length is the single most important factor in password security. At 35 characters, this password creates an astronomical number of possible combinations. Even with modern brute-force hardware, it would take centuries to crack.

* **High Entropy:** Entropy measures the unpredictability of a password. By combining seemingly random word-like strings (`fitolute`, `pygoFali`) with numbers and hyphens, you have created a high-entropy string that doesn't follow common dictionary patterns.

* **Complexity without Chaos:** You’ve successfully integrated multiple character types:

    * **Easy for a human to type:** It follows a rhythmic pattern.

    * **Lowercase & Uppercase:** `saneFun`, `pygoFali`.

    * **Numbers:** `3`, `8`.

    * **Special Characters:** The hyphens (`-`) serve as excellent delimiters.

* **The "Diceware" Structure:** This format mimics the "Diceware" method, where random words are joined together. This is widely considered the gold standard for human-generated passwords because it avoids the "predictable complexity" (like replacing 'a' with '@') that hackers' tools are programmed to expect.

**Final Verdict: S-Tier**

This is a "Master Password" grade string. It is perfectly suited for protecting your most sensitive assets.

## Build

See [BUILD.md](BUILD.md) file.

## License

`gencvcpass` was created by Lukasz Lobocki. It is licensed under the terms of the CC0 v1.0 Universal license.

All components used retain their original licenses.

## Credits

`gencvcpass` was created with [cookiecutter](https://cookiecutter.readthedocs.io/en/latest/) and [template](https://github.com/lukasz-lobocki/go-cookiecutter).
