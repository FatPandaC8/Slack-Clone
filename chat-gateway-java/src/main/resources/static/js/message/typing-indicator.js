let indicatorEl = null;
let textEl = null;
let typingTimeout = null;
let getCurrentUserId = null;
const typingUsers = new Set();

export function initTypingIndicator({indicator, text, getUserId}) {
    indicatorEl = indicator;
    textEl = text;
    getCurrentUserId = getUserId;
}

export function updateTypingUI() {
    if (typingUsers.size == 0) {
        indicatorEl.classList.add("hidden");
        return;
    }

    indicatorEl.classList.remove("hidden");
    console.log(typingUsers);

    const names = Array.from(typingUsers);

    if (names.length === 1) {
        textEl.textContent = `${names[0]} is typing`;
    } else if (names.length === 2) {
        textEl.textContent = `${names[0]} and ${names[1]} are typing`;
    } else {
        textEl.textContent = `${names[0]} and ${names.length - 1} others are typing...`
    }
}

export function userStartedTyping(userId) {
    if (!userId) return;
    typingUsers.add(userId);
    updateTypingUI();
    resetAutoHide();
}

export function userStoppedTyping(userId) {
    if (!userId) return;
    typingUsers.delete(userId);
    updateTypingUI();
}

function resetAutoHide() {
    clearTimeout(typingTimeout);
    typingTimeout = setTimeout(() => {
        typingUsers.clear();
        updateTypingUI();
    }, 2000); // hide after 2s
}