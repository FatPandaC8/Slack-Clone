import { createConversation, currentUserId, currentConversationId, createUser } from "./creation/creation.js";
import { initTypingIndicator } from "./message/typing-indicator.js";
import { joinConversation } from "./conversation/joinConversation.js";
import { sendTyping } from "./message/typing-sender.js";
import { send } from "./message/send.js";

// Grab inputs once
const convNameInput         = document.getElementById("createConvInput");
const textInput             = document.getElementById("textInput");
const typingIndicator       = document.getElementById("typingIndicator");
const typingText            = document.getElementById("typingText");
const inviteCodeInput       = document.getElementById("inviteCodeInput");

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
const joinBtn               = document.getElementById("joinConvBtn");

// Event listeners
createUserBtn.addEventListener("click", async () => {
    const user = await createUser(
        nameInput.value,
        emailInput.value,
        passwordInput.value
    );

    alert("Logged in as " + user.name + " with id of " + user.id);
});
createConvBtn.addEventListener("click", () => createConversation(convNameInput));

sendMessageBtn.addEventListener("click",       () => send(textInput));

joinBtn.addEventListener("click", async () => joinConversation(inviteCodeInput));

textInput.addEventListener("input",     () => {
    sendTyping(currentConversationId, currentUserId);
    textInput.style.height = "auto";
    textInput.style.height = textInput.scrollHeight + "px";
});

// Enter key for sending
textInput.addEventListener("keydown", (event) => {
    if (event.key === "Enter") send(textInput);
});