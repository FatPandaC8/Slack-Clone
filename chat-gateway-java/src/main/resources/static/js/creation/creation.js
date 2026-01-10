import { loadConversation } from "../conversation/loadConversation.js";
import { loadMyConversations } from "../conversation/loadMyConversation.js";
import { connectWS } from "../websocket/websocket.js";

export let currentUserId = null;
export let currentConversationId = null;

function setCurrentUser(user) {
  localStorage.setItem("currentUser", JSON.stringify(user));
}

function getCurrentUser() {
    return JSON.parse(localStorage.getItem("currentUser"));
}

export async function createUser(name, email, password) {
  const payload = {name, email, password};

  const res = await fetch("/users", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload)
  });

  const data = await res.json();

  setCurrentUser(data);
  currentUserId = data.id;
  await connectWS();
  loadMyConversations();
  return data;
}

export async function createConversation(convNameInput) {
  const name = convNameInput.value.trim();

  if (!name) {
    alert("Conversation name required");
    return;
  }

  const res = await fetch("/conversations", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
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