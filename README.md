# Mobile ordering exercise

## Scenario

You are a developer at a company that provides a mobile ordering platform for
restaurants. The platform allows customers to order food from restaurants
through a mobile app. The app is built using React Native and communicates with
a backend API written in Node.js.

The backend forwards orders to a third-party system (called Lightspeed) that is
connected to the Point of Sale (POS) terminals in the restaurants, so that the
orders can be processed and prepared by the restaurant staff.

## Problem statement

Customers at the same table place orders individually, from their own phones.
This means that orders to the kitchen from the same table can arrive at
different times, mixed with orders from other tables.

The restaurant staff would like to have orders from the same table arrive at
(roughly) the same time, so that they can be processed together.

## Setup

Install dependencies:

```
npm install
```

For this exercise, you can run a very lightweight version of the third-party API
locally:

```
npm run lightspeed
```

This API will accept orders on `localhost`, port `4000`. It will also print out
the orders it receives to the console.

To simulate orders being placed, you can run the following command:

```
npm run orderbot
```

This will send orders, on average once per second, to the API running on port
`3000`, which you can start with:

```
npm start
```

The orderbot will terminate with an error if the API is not running.

You shouldn’t need it, but in case the Go binaries don’t work, install Go via
Homebrew and build the binaries yourself.:

```
brew install go
npm run build
```

Or if you don't want to install Go on your machine, you can use the dev container provided and start the Go files with:

```
go run go/lightspeed.go
go run go/orderbot.go
```


You should edit the code in `src/index.ts` to solve the problem. You may install
any additional dependencies you need. You may also use any online resources you
like.
