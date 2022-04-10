import { createServer } from "http";
import {Server} from 'socket.io';
import express from 'express';
import cors from 'cors';

const app = express();

app.use(cors())

const httpServer = createServer(app);
const io = new Server(httpServer,{
    cors :{ 
        origin: "*",
    }
});

io.on("connection", socket => {
    console.log(socket)
})

httpServer.listen(9000, () => {
    console.log("server running");
});