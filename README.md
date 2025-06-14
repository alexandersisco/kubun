# Kubun  くぶん 

A CLI for slicing delimiter-separated text — Python-style

Kubun draws inspiration from Python's slice syntax: `[start:stop:step]`. It is meant to be a more intuitive solution for slicing fields and segments of delimiter-separated text on the command line than tools like `cut`, `awk`, `sed`, et al.

## Usage
Kubun takes up to two positional arguments: the SLICE pattern followed by the optional INPUT. 

Here are some basic examples:
```
  kubun '[:]' /usr/bin/sort               -> /usr/bin/sort
  kubun '[-1:]' /usr/bin/sort             -> sort
  kubun '[-2:]' /usr/bin/sort             -> bin/sort
```
You can also pipe text into Kubun:
```
  echo "/usr/bin/sort" | kubun '[-3:]'    -> usr/bin/sort
```
To replace the delimiters in a string of text, include the old and new delimiters on either side of the slice pattern like so: `'<old>[:]<new>'`.
```
  kubun '/[:]\'                           -> \usr\bin\sort
  kubun '/[1:], '                         -> usr, bin, sort
  kubun '/[1:]\n'                         -> usr
                                             bin
                                             sort
```
Kubun supports reversing fields by passing in a negative step. However, Python-style behavior for step values greater than |1| is not yet implemented.
```
  kubun '[::-1]' /usr/bin/sort            -> sort/bin/usr/
```
