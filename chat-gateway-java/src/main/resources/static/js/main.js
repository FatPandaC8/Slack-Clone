import { createConversation, currentUserId, currentConversationId, createUser } from "./creation/creation.js";
import { send } from "./message/send.js";
import { initTypingIndicator } from "./message/typing-indicator.js";
import { sendTyping } from "./message/typing-sender.js";

// Grab inputs once
const userIdInput           = document.getElementById("userId");
const convIdInput           = document.getElementById("conversationId");
const memberIdsInput        = document.getElementById("memberIds");
const textInput             = document.getElementById("textInput");
const typingIndicator       = document.getElementById("typingIndicator");
const typingText            = document.getElementById("typingText");

// Init
initTypingIndicator({
    indicator: typingIndicator,
    text: typingText
});

// Buttons
const createUserBtn         = document.getElementById("createUserBtn");
const nameInput             = document.getElementById("name");
const emailInput            = document.getElementById("email");
const passwordInput         = document.getElementById("password");
const createConvBtn         = document.getElementById("createConvBtn");
const sendMessageBtn        = document.getElementById("sendMessageBtn");

// Event listeners
createUserBtn.addEventListener("click", async () => {
    const user = await createUser(
        nameInput.value,
        emailInput.value,
        passwordInput.value
    );

    alert("Logged in as " + user.name + "with id of " + user.id);
});
createConvBtn.addEventListener("click", () => createConversation(convIdInput, memberIdsInput));
sendMessageBtn.addEventListener("click",       () => send(textInput));
textInput.addEventListener("input",     () => {
    sendTyping(currentConversationId, currentUserId);
    textInput.style.height = "auto";
    textInput.style.height = textInput.scrollHeight + "px";
});

// Enter key for sending
textInput.addEventListener("keydown", (event) => {
    if (event.key === "Enter") send(textInput);
});