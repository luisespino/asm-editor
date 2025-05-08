# ASM-Editor
An open-source assembly language editor for the terminal.

Supports any assembly languages compatible with the GNU Assembler. Uses the "as" and "ld" commands to assembler and link executables.

Written in Go, featuring a text-based user interface (TUI) built with Termbox, following a Nano-style interaction model.

In the current version, you can:
- Save files
- Exit the editor
- Run your code.

## Build Instructions

Requirements:
- Go
- Git

Clone the repository:
```
git clone https://github.com/luisespino/asm-editor.git
cd asm-editor
```

Check the Go version:
```
go version
```

>[!IMPORTANT]
>Make sure the version matches the one specified in your go.mod file.
If it doesn't, update go.mod to reflect your current Go version.

Build the editor:
```
go build -o asm-editor
```

Move the binary to a system path:
```
sudo mv asm-editor /usr/local/bin/
```

Run the editor (use the appropriate write.s file for your architecture):
```
asm-editor test/write.s
```

Editor screenshot:
![Alt text](https://github.com/luisespino/asm-editor/blob/main/img/editor.png?raw=true "editor")

Run screenshot:
![Alt text](https://github.com/luisespino/asm-editor/blob/main/img/run.png?raw=true "run")
