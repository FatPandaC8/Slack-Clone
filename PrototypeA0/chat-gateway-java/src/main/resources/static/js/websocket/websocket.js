import { renderMessage } from "../conversation/loadConversation.js";
import { updateTypingUI } from "../message/typing-indicator.js";

let stompClient = null;
let currentSubscription = null;
let currentTypingSub = null;

export function connectWS() {
  const socket = new SockJS("/ws");
  stompClient = Stomp.over(socket);

  return new Promise((resolve) => {
    stompClient.connect(
      {}, () => {
        console.log("WS connected");
        resolve();
      }
    );
  });
}

export function sendWS(destination, payload) {
  if (!stompClient) return;

  stompClient.send(
    destination,
    {},
    JSON.stringify(payload)
  );
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

export function subscribeTyping(conversationId) {
  if (currentTypingSub) {
    currentTypingSub.unsubscribe();
  }

  currentTypingSub = stompClient.subscribe(
    `/topic/conversations/${conversationId}/typing`,
    (msg) => {
      const data = JSON.parse(msg.body);
      updateTypingUI([...data.usersTyping]);
    }
  );
}

function appendMessage(msg) {
  const box = document.getElementById("messages");
  box.appendChild(renderMessage(msg.senderId, msg.text));
}