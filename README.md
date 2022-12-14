# Fantiger Community API

[![Build Status](https://travis-ci.org/hagopj13/node-express-boilerplate.svg?branch=master)](https://travis-ci.org/hagopj13/node-express-boilerplate)
[![Coverage Status](https://coveralls.io/repos/github/hagopj13/node-express-boilerplate/badge.svg?branch=master)](https://coveralls.io/github/hagopj13/node-express-boilerplate?branch=master)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

A project for fantiger community RESTful APIs using Node.js, Express, and Mongoose.

## Quick Start

## Installation

If you would still prefer to do the installation manually, follow these steps:

Clone the repo:

```bash
git clone https://github.com/rkgitvinay/gin-mongo-boilerplate.git <folder-name>
cd <folder-name>
```

Install the dependencies:

```bash
yarn
```

Set the environment variables:

```bash
cp env.example .env

## Commands

Running locally:

```bash
yarn dev
```

Running in production:

```bash
yarn start
```

## Environment Variables

The environment variables can be found and modified in the `.env` file. They come with these default values:

```bash
# Port number
PORT=6000

# URL of the Mongo DB
MONGOURI=mongodb://127.0.0.1:27017/gin-mongo-boilderplate
```

## Project Structure

```
src\
 |--config\         # Environment variables and configuration related things
 |--controllers\    # Route controllers (controller layer)
 |--docs\           # Swagger files
 |--middlewares\    # Custom express middlewares
 |--models\         # Mongoose models (data layer)
 |--routes\         # Routes
 |--services\       # Business logic (service layer)
 |--utils\          # Utility classes and functions
 |--validations\    # Request data validation schemas
 |--app.js          # Express app
 |--index.js        # App entry point
```

**User routes**:\
`POST /user` - create a user\
`GET /users` - get all users\
`GET /user/:userId` - get user\
`PATCH /user/:userId` - update user\
`DELETE /user/:userId` - delete user


