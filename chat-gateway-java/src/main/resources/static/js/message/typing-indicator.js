import { currentUserId } from "../creation/creation.js";

let indicatorEl = null;
let textEl = null;

export function initTypingIndicator({ indicator, text }) {
  indicatorEl = indicator;
  textEl = text;
}

export function updateTypingUI(usersTyping) {
    const others = usersTyping.filter(u => u !== currentUserId);

    if (!others || others.length === 0) {
        indicatorEl.classList.add("hidden");
        return;
    }

    indicatorEl.classList.remove("hidden");

    if (others.length === 1) {
        textEl.textContent = `${others[0]} is typing`;
    } else if (others.length === 2) {
        textEl.textContent = `${others[0]} and ${others[1]} are typing`;
    } else {
        textEl.textContent =
        `${others[0]} and ${others.length - 1} others are typing...`;
    }
}