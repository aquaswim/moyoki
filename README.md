# ꦩꦺꦴꦪꦺꦴꦏꦶ

Moyoki (javanese) mean mocking, this is simple program to do http server mocking.

## Prerequisite

* go >= 1.24
* devbox

## How to run

1. clone this repo
2. create .env file `cp .env.example .env`
3. run `devbox run setup`
4. run `devbox run start`

## How to run FE development mode

FE development mode will enable vite HMR.

1. make sure .env file is exist
2. start the vite server with: `devbox run dev:front`
3. after vite server ready start the backend: `devbox run dev:be`