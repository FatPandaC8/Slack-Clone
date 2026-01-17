export function getToken() {
  return localStorage.getItem("token");
}

export function logoutFunc() {
  localStorage.removeItem("currentUser");
  localStorage.removeItem("token");
  alert("Logged out");
  window.location.href = "auth.html";
}