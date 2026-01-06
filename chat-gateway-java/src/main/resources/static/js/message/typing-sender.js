import { sendWS } from "../websocket/websocket.js";
let typingTimeout = null;
let isTyping = false;

export function sendTyping(conversationId, userId) {
    if (!conversationId || !userId) return;

    if (!isTyping) {
        sendWS("/app/typing.start", {
            conversationId,
            userId
        });
        isTyping = true;
    }

    clearTimeout(typingTimeout);
    typingTimeout = setTimeout(() => {
        sendWS("/app/typing.stop", {
        conversationId,
        userId
        });
        isTyping = false;
    }, 2000);
}