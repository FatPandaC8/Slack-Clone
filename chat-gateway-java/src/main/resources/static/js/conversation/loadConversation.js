import { getToken } from "../auth/auth.js";
import { setCurrentConversationId } from "../creation/creation.js";

export async function loadConversation(conversationId) {
  if (!conversationId) return;
  const token = getToken();
  const res = await fetch(`/conversations/${conversationId}`, {
    headers: {
      "Authorization": "Bearer " + token
    }
  });
  const data = await res.json();
  
  const box = document.getElementById("messages");
  box.innerHTML = "";
  
  data.messages.forEach(m => {
    box.appendChild(renderMessage(m.name, m.text));
  });
  
  // Update the global
  setCurrentConversationId(conversationId);
}

export function renderMessage(name, text) {
  const div = document.createElement("div");
  div.classList.add("msg");
  div.innerText = `${name}: ${text}`;
  return div;
}