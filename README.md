# 🏺 nilus

Nilus is a tool for writing ancient Greek with a Latin keyboard. It allows you to use all of those pesky accents and diacritics not usually allowed with a Greek keyboard.

Each Latin character is replaced with a phonetically equivalent Greek character, or a Greek character found in the same place on the keyboard.

Leading symbols can modify the following character, adding accents and diacritics and whatnot.

Those symbols are:

```
' Acute accent
` Grave accent
~ Circumflex accent
| Tonos accent
[ Rough breathing
] Smooth breathing
_ Long vowel
^ Short vowel
: Diaeresis accent
* Iota subscript
```

There is also punctuation conversion

```
? -> ;
; -> ·
```

Usage:

```
nilus [TEXT]... (also accepts input from STDIN)

Usage of nilus:
  -a	Use alternate/archaic characters
```
