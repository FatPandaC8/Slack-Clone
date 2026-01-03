let currentConversationId = null;
let currentUserId = null;

async function setUser() {
  const uid = document.getElementById("userId").value.trim();
  if (!uid) return alert("userId required");

  currentUserId = uid;
  alert(`You're: ${uid}`)
  loadMyConversations();
}

async function createConversation() {
  const conversationId = document.getElementById("conversationId").value;
  const members = document.getElementById("memberIds").value
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

async function loadConversation() {
  if (!currentConversationId) return;

  const res = await fetch(`/conversations/${currentConversationId}`);
  const data = await res.json();

  const box = document.getElementById("messages");
  box.innerHTML = "";

  data.messages.forEach(m => {
    const div = document.createElement("div");
    div.className = "msg";
    div.innerText = `${m.senderId}: ${m.text}`;
    box.appendChild(div);
  });
}

async function send() {
  if (!currentConversationId) return alert("create a conversation first");

  const text = document.getElementById("text").value;
  if (!text) return;

  await fetch(`/conversations/${currentConversationId}/messages`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      senderId: currentUserId,
      text: text
    })
  });

  document.getElementById("text").value = "";
  loadConversation();
}

async function loadMyConversations() {
  const res = await fetch(`/users/${currentUserId}/conversations`);
  const convs = await res.json();

  const list = document.getElementById("conversationList");
  list.innerHTML = ""

  convs.forEach(id => {
    const btn = document.createElement("button");
    btn.innerText = id;
    btn.onclick = () => {
      currentConversationId = id;
      loadConversation();
    };
    list.appendChild(btn);
  });
}