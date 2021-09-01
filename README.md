# 🏺 nilus

Nilus (Νεῖλόν, Ne~il\|on) is a tool for writing ancient Greek with a Latin keyboard. It allows you to use all of those pesky accents and diacritics not usually allowed with a Greek keyboard.

Each Latin character is replaced with a phonetically equivalent Greek character, or a Greek character found in the same place on the keyboard. (When using `-k`)

Leading symbols can modify the following character or digraph, adding accents and diacritics and whatnot.

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

Keyboard mode (non-transliterative characters)

```
h -> Eta
u -> Theta
j -> Xi
x -> Chi
c -> Psi
v -> Omega
f -> Phi
w -> Ultimate Sigma
q -> Hooked Upsilon
```

Normal mode (diagraphs and alternate characters)

```
\e -> Eta
Th -> Theta
\y -> Hooked Upsilon
Ch -> Chi
Ps -> Psi
\o -> Omega
Ph -> Phi
\s -> Ultimate Sigma
Rh -> Rho (rough breathing)
```

Usage:

```
nilus [TEXT]... (also accepts input from STDIN)

Usage of nilus:
  -k	Use keyboard substitution mode
```

Built using Go version 1.16+
