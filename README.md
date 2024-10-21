# `minimsg`

Tiny (incomplete) [minimessage](https://docs.advntr.dev/minimessage/format.html) parser for [Gate](https://gate.minekube.com/).

You can find the API docs here: <https://pkg.go.dev/github.com/thedevminertv/minimsg>.

Based on [Gate's plugin template](https://github.com/minekube/gate-plugin-template/blob/2fee6678fc6782d3560768d00149ad46693db549/util/mini/mini.go).

## Features

If you're looking at this from pkg.go.dev, please open the GitHub page. Go's package registry doesn't render the checkboxes.

- [ ] Reset (`<reset>`)
- [ ] Fonts (`<font:<FONT>>`)
- [x] Colors (`<color:<COLOR>>`, `<colour:<COLOR>>`, `<c:<COLOR>>`, `<COLOR>`)
  - [x] Hex (`#fff`, `#ffffff`)
  - [x] By Name (`gray`, etc.)
    - Black
    - DarkBlue
    - DarkGreen
    - DarkAqua
    - DarkRed
    - DarkPurple
    - Gold
    - Gray
    - DarkGray
    - Blue
    - Green
    - Aqua
    - Red
    - LightPurple
    - Yellow
    - White
  - [ ] Rainbow (`<rainbow[:[!][PHASE]]>`)
  - Gradients (`<gradient>`)
    - [x] Without phase (`<gradient:<FROM_COLOR>:<TO_COLOR>[:COLORS...]>`)
    - [ ] With phase (`<gradient:<FROM_COLOR>:<TO_COLOR>[:COLORS...]:<PHASE>`)
  - [ ] Transitions (`<transition>`)
    - [ ] Without phase (`<transition:<FROM_COLOR>:[COLORS...]>`)
    - [ ] With phase (`<transition:<FROM_COLOR>:[COLORS...]:<PHASE>>`)
- Text
  - [x] Normal text
  - [ ] Translatable text (`<lang:<KEY>:<VALUE_1>:<VALUE_2>`, `<tr:<KEY>:<VALUE_1>:<VALUE_2>`, `<translate:<KEY>:<VALUE_1>:<VALUE_2>`)
  - [ ] New line (`<newline>`, `<br>`)
  - Selector (`<selector:<SELECTOR>[:<SEPARATOR>]`, `<sel:<SELECTOR>[:<SEPARATOR>]`)
    - Won't be implemented, since you need a server that knows the entities
  - Score (`<score:<NAME>:<OBJECTIVE>`)
    - Won't be implemented, since you need a server that knows the scores
  - NBT (`<nbt:<block|entity|storage>:<ID>:<PATH>[:<SEPARATOR>][:<INTERPRET>`)
    - Won't be implemented, since you need a server that knows the NBT data
- Decorations
  - [ ] Inverting (`<!<DECORATION>>`, `<<DECORATION>:false>`)
  - [x] Bold (`<bold>`, `<b>`)
  - [x] Italics (`<italic>`, `<i>`, `<em>`)
  - [x] Underline (`<underlined>`, `<underline>`, `<u>`)
  - [x] Strikethrough (`<strikethrough>`, `<st>`)
  - [x] Obfuscated (`<obfuscated>`, `<obfuscated>`, `<obf>`)
- Behaviour
  - [ ] Insertion (`<insertion:<TEXT:<TEXT>>`)
  - Hover (`<hover:<ACTION>:<VALUE>>`)
    - [ ] Show achievement <!-- TODO: Check if it's actually possible to send this -->
    - [ ] Show entity
    - [ ] Show item
    - [ ] Show text
  - Click (`<click:<ACTION>:<VALUE>>`)
    - [ ] Change page
    - [ ] Copy to clipboard
    - [ ] Open file
    - [ ] Open URL
    - [ ] Run command
    - [ ] Suggest command

## License

This project is licensed under the MIT license.
For more information, please see the included [LICENSE](/License) file.
