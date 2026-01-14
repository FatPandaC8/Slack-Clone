import { getToken } from "../auth/auth.js";
import { loadConversation } from "../conversation/loadConversation.js";
import { loadMyConversations } from "../conversation/loadMyConversation.js";

export let currentUserId = null;
export let currentConversationId = null;

export async function createConversation(convNameInput) {
  const name = convNameInput.value.trim();
  const token = getToken();
  if (!name) {
    alert("Conversation name required");
    return;
  }

  const res = await fetch("/conversations", {
    method: "POST",
    headers: { 
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token
    },
    body: JSON.stringify({ 
      name: name,
      creatorId: currentUserId
    })
  });

  const data = await res.json();

  setCurrentConversationId(data.id);

  alert(`Conversation "${data.name}" created.\nInvite code: ${data.inviteCode}`);

  loadMyConversations();
  loadConversation();
}

export function setCurrentConversationId(id) {
  currentConversationId = id;
}

export function setCurrentUserId(id) {
  currentUserId = id;
}