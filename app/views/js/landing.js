window.onload = function() {
  document.getElementById("login-btn").addEventListener("click", function () {
    window.location = `/login`;
  });

  document.getElementById("start-btn").addEventListener("click", function () {
    window.location = `/sign-up`;
  });
  
};