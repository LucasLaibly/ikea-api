# IKEA API

## The Not-So-Imaginary Problem

IKEA offers it's customers a wide variety of decoration and housing furnature options. Walking through the staged rooms, seeing how things look, touching items, whatever you need to do to feel confident in a particular item BEFORE purchasing is satisfied. 

Except it is not.

There is a particular time when walking through IKEA that the item you would like to purchase is not in stock, and there is absolutely NO WAY to know this until you get to the end of the store (post meatball meal of course), only to use one of their computers to figure out what shelf that item is on. It is here that the crux of their problems starts and end. 

Why does IKEA not offer a QR code to scan that provides immeadiate feedback on in-stock options where the item is being displayed prior to walking the marathon of rooms and learning it is out of stock. 

This API is nothing more than a series of endpoints to categorize and simulate that flow of a customer (you) scanning each item. 

Also, I want to learn golang and making an imaginary problem keeps it interesting. 

## Things needed
1. `docker` 
2. `postgres` 
3. `go`

## To run application locally
Begin by running `cp .env.example .env` , this is where you will need to fill in environment variables.