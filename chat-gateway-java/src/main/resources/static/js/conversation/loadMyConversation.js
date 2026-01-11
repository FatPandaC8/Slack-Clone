import { currentUserId, setCurrentConversationId } from "../creation/creation.js";
import { subscribeConversation, subscribeTyping } from "../websocket/websocket.js";
import { loadConversation } from "./loadConversation.js";

export async function loadMyConversations() {
  const res = await fetch(`/users/${currentUserId}/conversations`);
  const convs = await res.json();

  const list = document.getElementById("conversationList");
  list.innerHTML = ""

  console.log(convs);

  convs.forEach(conv => {
    const btn = document.createElement("button");
    btn.innerText = conv.name;
    btn.onclick = () => {
      setCurrentConversationId(conv.id);
      loadConversation(conv.id);
      subscribeTyping(conv.id);
      subscribeConversation(conv.id);
    };
    list.appendChild(btn);
  });
}