import { setCurrentConversationId } from "../creation/creation.js";

export async function loadConversation(conversationId) {
  if (!conversationId) return;
  
  const res = await fetch(`/conversations/${conversationId}`);
  const data = await res.json();
  
  const box = document.getElementById("messages");
  box.innerHTML = "";
  
  data.messages.forEach(m => {
    const div = document.createElement("div");
    div.className = "msg";
    div.innerText = `${m.senderId}: ${m.text}`;
    box.appendChild(div);
  });
  
  // Update the global
  setCurrentConversationId(conversationId);
}