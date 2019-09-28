# autocompletion
autocompletion is a go command auto completion demo

# Usage
## Build

macOS
```
make mac
```

Ubuntu
```
make linux
```

## Config

```
$ sudo cp autocompletion /usr/local/bin/
```

~/.bashrc
```
source <(autocompletion completion bash)
```

```
$ autocompletion completion bash > $(brew --prefix)/etc/bash_completion.d/autocompletion
$ source ~/.bashrc
$ source ~/.bash_profile
```

```
$ autocompletion
Autocompletion implements two simple commands, list and cat.

Usage:
  autocompletion [command]

Available Commands:
  cat         concatenate files and print on the standard output
  completion  Output shell completion code for the specified shell (bash or zsh)
  help        Help about any command
  list        list directory contents

Flags:
  -h, --help   help for autocompletion

Use "autocompletion [command] --help" for more information about a command.
```