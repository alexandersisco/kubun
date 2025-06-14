# Kubun  くぶん 

A CLI for slicing delimiter-separated text

Kubun draws inspiration from the way that Python makes it easy to slice and dice lists. It is meant to be a more intuitive solution for slicing fields and segments of text on the command line than tools like `cut`, `awk`, `sed`, et al.

## Install
```
go install github.com/alexandersisco/kubun@latest
```

## Usage
Kubun takes up to two positional arguments: the SLICE pattern followed by the optional INPUT. 

Python-like syntax:

The syntax for Kubun's slice pattern follows that of Python which includes the following components `[start:stop:step]`. Each is a zero-based index into the delimited fields being sliced. The stop index is exclusive, meaning Kubun will stop before the field that the index points to, excluding it from the output. A negative start or stop index is measured from the end of the input with a -1 representing the last field. For brevity, the step may be omitted: like this: `[start:stop]`.

Here are some basic examples:
```
  kubun '[:]' /usr/bin/sort               -> /usr/bin/sort
  kubun '[:-1]' /usr/bin/sort             -> /usr/bin
  kubun '[-1:]' /usr/bin/sort             -> sort
  kubun '[-2:]' /usr/bin/sort             -> bin/sort
```
You can also pipe text into Kubun:
```
  echo "/usr/bin/sort" | kubun '[-3:]'    -> usr/bin/sort
```
Kubun supports reversing fields by passing in a negative step. However, Python-style behavior for step values greater than |1| is not yet implemented.
```
  kubun '[::-1]' /usr/bin/sort            -> sort/bin/usr/
```
### Replacing delimiters
Kubun makes it easy to replace delimiters.

To replace the delimiters in a string of text, include the old and new delimiters on either side of the slice pattern like so: `'<old>[:]<new>'`.
```
  kubun '/[:]\'                           -> \usr\bin\sort
  kubun '/[1:], '                         -> usr, bin, sort
  kubun '/[1:]\n'                         -> usr
                                             bin
                                             sort
```

