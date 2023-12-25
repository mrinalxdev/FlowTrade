document.addEventListener("DOMContentLoader", () => {
  const socket = io();
  let type = false;

  document.querySelector("form").addEventListener("submit", (event) => {
    socket.emit("stop typing");
    socket.emit("typing", document.getElementById("m").value);
    event.preventDefault();
  });
  document.getElementById("m").addEventListener("input", () => {
    if (!typing) {
      typing = true;
      socket.emit("stop typing");
    }
    setTimeout(() => {
      typing = false;
      socket.emit("stop typing");
    }, 1000);
  });
});
