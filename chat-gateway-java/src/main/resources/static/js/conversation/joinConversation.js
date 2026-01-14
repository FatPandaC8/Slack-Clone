import { getToken } from "../auth/auth.js";
import { loadMyConversations } from "./loadMyConversation.js";

export function getCurrentUser() {
    return JSON.parse(localStorage.getItem("currentUser"));
}

export async function joinConversation(inviteCode) {
  const invite = inviteCode.trim();
  const user = getCurrentUser();
  const token = getToken();

  if (!user) {
    alert("Please create/login user first.");
    return;
  }

  if (!token) {
    alert("No token found. Please log in again.");
    return;
  }

  if (!invite) {
    alert("Invite code is required");
    return;
  }

  const res = await fetch("/conversations/join", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token
    },
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