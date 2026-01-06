import { setUser, createConversation, currentUserId, currentConversationId } from "./creation/creation.js";
import { send } from "./message/send.js";
import { initTypingIndicator } from "./message/typing-indicator.js";
import { sendTyping } from "./message/typing-sender.js";

// Grab inputs once
const userIdInput           = document.getElementById("userId");
const convIdInput           = document.getElementById("conversationId");
const memberIdsInput        = document.getElementById("memberIds");
const textInput             = document.getElementById("text");
const typingIndicator       = document.getElementById("typingIndicator");
const typingText            = document.getElementById("typingText");

// Init
initTypingIndicator({
    indicator: typingIndicator,
    text: typingText
});

// Buttons
const setUserBtn            = document.getElementById("setUserBtn");
const createConvBtn         = document.getElementById("createConvBtn");
const sendBtn               = document.getElementById("sendBtn");

// Event listeners
setUserBtn.addEventListener("click",    () => setUser(userIdInput));
createConvBtn.addEventListener("click", () => createConversation(convIdInput, memberIdsInput));
sendBtn.addEventListener("click",       () => send(textInput));
textInput.addEventListener("input",     () => sendTyping(currentConversationId, currentUserId));

// Enter key for sending
textInput.addEventListener("keydown", (event) => {
    if (event.key === "Enter") send(textInput);
});