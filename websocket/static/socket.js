const counter = document.getElementById('counter');
const resetButton = document.getElementById('reset');

resetButton.addEventListener('click', sendResetRequest, false);

let Ws = new WebSocket('ws://localhost:3000/counter');

Ws.onopen = (e) => {
    Ws.onmessage = (e) => {
        let msg = JSON.parse(e.data);
        console.log(msg)
        counter.innerText = msg.count;
    }
}

function sendResetRequest() {
Ws.send("reset");
}

