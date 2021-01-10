# Development

Clone the project and move to the project directory

```bash
git clone https://github.com/jonti09/themer.git
cd themer
```

## Server

To run the server, you will need Golang to be installed, which you can get from [here](https://golang.org/dl/).

Once the Golang compiler is installed, run the following command that starts dev server on port `8000`

```bash
go run main.go --dev
```

> Note: if you do not pass `--dev` server will start in production mode!

## Client

To run the client, you will need node.js to be installed, which can be found [here](https://nodejs.org/en/download/).

Once installed, run the following commands to start the client development server

```bash
cd client
npm start
```

Which will start the development server on `3000` port.
