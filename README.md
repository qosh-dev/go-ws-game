# Game README

## Overview

Welcome to GameDream, a simple but engaging console game built with Golang, Gin, Socket.IO, and GORM! It features a multi-player environment where players can communicate, track stats, and unleash firebombs upon each other.

## Modules

### Server

The server module is responsible for managing the game logic and communication between players. It uses Gin for the web framework, Socket.IO for real-time communication, and Gorm for interacting with the database.

### Client

The client module represents the players in the game. Players can run multiple instances of the client module to participate in the game. They can communicate with each other, receive stats, and throw firebombs.

## Setup

To start the game, follow these steps:

1. Create a <code>.env</code> file in the server directory with the following keys:

   ```
   PORT=<>
   DB_CONNECTION_STRING=<>
   SECRET=<>
   ```

2. Make sure you have an active PostgreSQL instance running, and specify the connection string in the <code>DB_CONNECTION_STRING</code> key of the <code>.env</code> file.

3. Dependencies Installation:
   <ul>
       <li>Navigate to the server directory and run <code>go mod download</code>.</li>
       <li>Navigate to the client directory and run <code>go mod download</code>.</li>
   </ul>

4. Server Build and Start:
    <ul>
        <li>Navigate to the server directory and run <code>go build -o build</code>.</li>
        <li>Execute the built server binary (./build or similar(base on your os)).</li>
    </ul>

5. Client Build and Start:
    <ul>
        <li>Navigate to the client directory and run <code>go build -o build</code>.</li>
        <li>Execute the built client binary (./build or similar(base on your os)) for each player.</li>
        <li>Than follow game instractions.</li>
    </ul>

## Additional Information

    Consider exploring the server and client code for further details and customization.
    Have fun and unleash those firebombs!
