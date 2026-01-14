import { getToken } from "../auth/auth.js";
import { currentConversationId, currentUserId } from "../creation/creation.js"

export async function send(textInput) {
  if (!currentConversationId) return alert("create a conversation first");
  const token = getToken();
  const text = textInput.value.trim();
  if (!text) return;

  await fetch(`/conversations/${currentConversationId}/messages`, {
    method: "POST",
    headers: { 
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token
    },
    body: JSON.stringify({
      senderId: currentUserId,
      text: text
    })
  });

  textInput.value = ""; // clear input
}