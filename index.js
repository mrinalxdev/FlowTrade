const express = require("express");
const http = require("http");
const socketIO = require("socket.io");
const path = require("path");
const PORT = 3000;

const app = express();
const server = http.createServer(app);
const io = socketIO(server);

app.use(express.static(path.join(__dirname, "public")));
let userCount = 0;

io.on("connection", (socket) => {
  userCount++;
  socket.userName = `User ${userCount}`;

  io.emit("chat message", {
    user: "system",
    message: `${socket.userName} jas joined chat`,
  });

  console.log(`${socket.userName} is connected`);

  socket.on("chat message", (msg) => {
    io.emit("chat message", { user: socket.userName, message: msg });
  });

  socket.on("typing", () => {
    socket.broadcast.emit("typing", socket.userName);
  });

  socket.on("stop typing", () => {
    socket.broadcast.emit("stop typing");
  });

  socket.on("disconnect", () => {
    io.emit("chat message", {
      user: "system",
      message: `${socket.userName} has left the chat`,
    });
    console.log(`${socket.userName} disconnected`);
  });
});

server.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
