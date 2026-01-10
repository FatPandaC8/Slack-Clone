import { loadConversation } from "../conversation/loadConversation.js";
import { loadMyConversations } from "../conversation/loadMyConversation.js";
import { connectWS } from "../websocket/websocket.js";

export let currentUserId = null;
export let currentConversationId = null;

function getCurrentUser() {
  return JSON.parse(localStorage.getItem("currentUser"));
}

function setCurrentUser(user) {
  localStorage.setItem("currentUser", JSON.stringify(user));
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
  await connectWS();
  loadMyConversations();
  return data;
}

export async function createConversation(convIdInput, memberIdsInput) {
  const conversationId = convIdInput.value.trim();
  const members = memberIdsInput.value
                                .split(",")
                                .map(s => s.trim());

  if (!conversationId) return alert("conversationId required");

  await fetch("/conversations", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      conversationId: conversationId,
      memberIds: members
    })
  });

  setCurrentConversationId(conversationId);
  alert(`conversation: ${conversationId} created successfully`);
  loadMyConversations();
  loadConversation();
}

export function setCurrentConversationId(id) {
  currentConversationId = id;
}

export function setCurrentUserId(id) {
  currentUserId = id;
}