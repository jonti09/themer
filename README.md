# THEMER

A Golang powered, React rendered terminal theme utility <span style="font-size: 12px">(WIP)</span>

## Usage

Download the [executable binary](https://github.com/jonti09/themer/releases/download/latest/themer.tgz) from the [release page](https://github.com/jonti09/themer/releases/tag/latest) (themer.tgz)

To extract and run the binary, enter the following commands in the terminal:

```bash
cd ~/Downloads/  # typical path, but change if your path is different!
tar -xvzf themer.tgz
chmod +x themer
./themer
```

Which should print out `Starting the server on http://localhost:8000/`

Once the server has started, visit `http://localhost:8000/`. From there, you can install the theme using the `download icon` on the left corner of the card.

Once the installation is successful, you can change your terminal theme like the following screenshot (check the `Profiles` section)

![alt text](https://github.com/jonti09/themer/blob/main/assets/images/theme-selection.png?raw=true)

## Features

- [x] Serve and display themes in the browser
- [x] Connect API endpoint to install the theme in terminal
- [x] Add pagination on client-side
- [x] Remove dependency on NPM, serve built content
- [x] Build everything in a package, ready for distribution

## Credits

### Huge credit to [Mayccoll](https://github.com/Mayccoll) for their project [Gogh](https://github.com/Mayccoll/Gogh) from which the project is using the executable theme files!

### Shout out to [meeDamian](https://github.com/meeDamian/) for their Github action, [github-release](https://github.com/meeDamian/github-release), which is what enables me to auto-release the executable binary as well as built assets for the project!

## Development

To run the project in development mode, please follow the instructions given in [development.md](https://github.com/jonti09/themer/blob/main/development.md)
