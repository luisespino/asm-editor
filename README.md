# asm-editor
Open-source assembly language editor for terminals.

Supports assembly languages supported by GNU Assembler. Uses the "as" and "ld" commands to create executables.

Written in Go using Termbox for the text-based UI. The editor follows a Nano-style interface.

In the preview version, you have options to save, exit, and run.

To run:
```
./asm-editor test/write.s
```

Editor screenshot:
![Alt text](https://github.com/luisespino/asm-editor/blob/main/img/editor.png?raw=true "editor")

Run screenshot:
![Alt text](https://github.com/luisespino/asm-editor/blob/main/img/run.png?raw=true "run")
