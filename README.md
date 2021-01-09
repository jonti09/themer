# THEMER

A Golang powered, React rendered terminal theme utility <span style="font-size: 12px">(WIP)</span>

# Development

Clone the project and move to the project directory

```bash
git clone https://github.com/jonti09/themer.git
cd themer
```

## Server

---

To run the server, you will need Golang to be installed, which you can get from [here](https://golang.org/dl/).

Once the Golang compiler is installed, run the following commands to start the server

```bash
go run main.go
```

Which will start the http server on `8000` port.

## Client

---

To run the client, you will need node.js to be installed, which can be found [here](https://nodejs.org/en/download/).

Once installed, run the following commands to start the client development server

```bash
cd client
npm start
```

Which will start the development server on `3000` port.

## Usage

---

Once the server and client are running, you can install the theme using the `download icon` on the left corner in the card.

Once the installation is successful, you can change your terminal theme like following screenshot (check the `Profiles` section)

![alt text](https://github.com/jonti09/themer/blob/main/public/images/theme-selection.png?raw=true)

## Features

---

- [x] Serve and display themes in browser
- [x] Connect API endpoint to install the theme in terminal
- [ ] Add pagination on client side
- [ ] Remove dependency on NPM, serve built content
- [ ] Build everything in a package, ready for distribution
