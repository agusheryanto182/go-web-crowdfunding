# Crowdfunding Backend API

This project is a backend API for a Crowdfunding website implemented using Go Fiber. It utilizes MySQL for database management, Redis for caching data, Cloudinary for image storage, Midtrans for payment processing, and OpenAI for assistant services.

## Features

- User authentication and authorization
- Project creation, update, deletion
- Campaign management
- Payment processing with Midtrans
- Image storage with Cloudinary
- Caching with Redis
- Assistant service with OpenAI

## Requirements

- Go
- MySQL
- Redis
- Cloudinary account
- Midtrans account
- OpenAI API key

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/agusheryanto182/go-web-crowdfunding.git
   ```

2. Go to folder go-web-crowdfunding

   ```bash
   cd go-web-crowdfunding
   ```

3. Copy file .env.example with name .env

   ```bash
   cp .env.example .env
   ```

4. Run docker compose

   ```bash
   docker-compose up --build
   ```

## API Documentation

   ```bash
   https://documenter.getpostman.com/view/32137512/2sA35BcjWd
   ```
