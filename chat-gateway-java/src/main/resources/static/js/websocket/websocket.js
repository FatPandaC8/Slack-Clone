let stompClient = null;
let currentSubscription = null;

export function connectWS() {
  const socket = new SockJS("/ws");
  stompClient = Stomp.over(socket);

  stompClient.connect({}, () => {
    console.log("WS connected");
  });
}

export function subscribeConversation(conversationId) {
  if (currentSubscription) {
    currentSubscription.unsubscribe();
  }

  currentSubscription = stompClient.subscribe(
    `/topic/conversations/${conversationId}`,
    (msg) => {
      const message = JSON.parse(msg.body);
      appendMessage(message);
    }
  );
}

function appendMessage(msg) {
  const box = document.getElementById("messages");
  const div = document.createElement("div");
  div.innerText = `${msg.senderId}: ${msg.text}`;
  box.appendChild(div);
}