# 🏺 nilus

Nilus (Νεῖλόν, Ne~il\|on) is a command for writing Ancient (polytonic) Greek with a Latin keyboard. It reads command-line arguments or `stdin` and allows you to use all of those pesky accents and diacritics not usually allowed with a Greek keyboard.

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
\e  -> Eta
th  -> Theta
\y  -> Hooked Upsilon
Ch  -> Chi
ps  -> Psi
\o  -> Omega
ph  -> Phi
\s  -> Ultimate Sigma
rh  -> Rho (rough breathing)
\r  -> Rho (with a tail)
w/v -> Digamma
```

Usage:

```
nilus [TEXT]... (also accepts input from STDIN)

Usage of nilus:
  -k	Use keyboard substitution mode
```

Built using Go version 1.16+
