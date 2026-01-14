export function getToken() {
  return localStorage.getItem("token");
}

export function logout() {
  localStorage.removeItem("currentUser");
  localStorage.removeItem("token");
  window.location.href = "auth.html";
}