import { currentConversationId, currentUserId, setCurrentConversationId } from "../creation/creation.js";

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
      subscribeConversation(id);
    };
    list.appendChild(btn);
  });
}