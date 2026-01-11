import { loadMyConversations } from "./loadMyConversation.js";

function getCurrentUser() {
    return JSON.parse(localStorage.getItem("currentUser"));
}

export async function joinConversation(inviteCode) {
  const invite = inviteCode.trim();
  const user = getCurrentUser();

  if (!user) {
    alert("Please create/login user first.");
    return;
  }

  if (!invite) {
    alert("Invite code is required");
    return;
  }

  const res = await fetch("/conversations/join", {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify({
      inviteCode: invite,
      userId: user.id
    })
  });

  console.log(invite);

  if (!res.ok) {
    alert("Failed to join conversation");
    return;
  }

  loadMyConversations();
  alert("Joined conversation!");
}