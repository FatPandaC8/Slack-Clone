import { loadConversation } from "../conversation/loadConversation.js";
import {loadMyConversations} from "../conversation/loadMyConversation.js";
import {connectWS} from "../websocket/websocket.js";

export let currentUserId = null;
export let currentConversationId = null;

export async function setUser(userIdInput) {
  const uid = userIdInput.value.trim();
  if (!uid) return alert("userId required");

  currentUserId = uid;
  alert(`You're: ${uid}`)
  connectWS();
  loadMyConversations();
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

  currentConversationId = conversationId;
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