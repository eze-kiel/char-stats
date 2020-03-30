# char-stats

![Licence](https://img.shields.io/badge/License-GPL-brightgreen)

Characters frequency analyser for mono-alphabetical substitution encryption.

## Usage

```Shell
options :
	-f <folder name>
		file which will be analysed
	-o <graph name>
		name of the output graph (.png)
```
Typical use :
```Shell
$ go run main.go -f text.txt -o output.png
```
This will analyse `text.txt` and display the results in a graph named `output.png`. You have have to specify the extension of the output (default : PNG).

## Example
Example with *La comédie humaine vol. 1* from Honoré de Balzac :
![](graph.png)

## Packages required

* [github.com/namsral/flag](https://github.com/namsral/flag)
* [github.com/wcharczuk/go-chart](https://github.com/wcharczuk/go-chart)

## Author

Written by ezekiel.

## Copyright

License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>. This is free software: you are free to change and redistribute it.