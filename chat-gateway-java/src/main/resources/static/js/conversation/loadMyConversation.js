import { currentUserId, setCurrentConversationId } from "../creation/creation.js";
import { subscribeConversation, subscribeTyping } from "../websocket/websocket.js";
import { loadConversation } from "./loadConversation.js";

export async function loadMyConversations() {
  const res = await fetch(`/users/${currentUserId}/conversations`);
  const convs = await res.json();

  const list = document.getElementById("conversationList");
  list.innerHTML = ""

  convs.forEach(id => {
    const btn = document.createElement("button");
    btn.innerText = id;
    btn.onclick = () => {
      setCurrentConversationId(id);
      loadConversation(id);
      subscribeTyping(id);
      subscribeConversation(id);
    };
    list.appendChild(btn);
  });
}