import { getCurrentUser } from "../conversation/joinConversation.js";

export function showUserInfo() {
    const userDiv = document.getElementById("userInfo");
    const user = getCurrentUser();
    userDiv.innerHTML = `<p><b>Name:<b> ${user.name}</p>`;
}

showUserInfo();