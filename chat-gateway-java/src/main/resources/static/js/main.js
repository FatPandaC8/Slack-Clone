import { setUser, createConversation } from "./creation/creation.js";
import { send } from "./message/send.js";

// Grab inputs once
const userIdInput = document.getElementById("userId");
const convIdInput = document.getElementById("conversationId");
const memberIdsInput = document.getElementById("memberIds");
const textInput = document.getElementById("text");

// Buttons
const setUserBtn = document.getElementById("setUserBtn");
const createConvBtn = document.getElementById("createConvBtn");
const sendBtn = document.getElementById("sendBtn");

// Event listeners
setUserBtn.addEventListener("click", () => setUser(userIdInput));
createConvBtn.addEventListener("click", () => createConversation(convIdInput, memberIdsInput));
sendBtn.addEventListener("click", () => send(textInput));

// Enter key for sending
textInput.addEventListener("keydown", (event) => {
    if (event.key === "Enter") send(textInput);
});