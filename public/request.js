document
  .getElementById("password-form")
  .addEventListener("submit", function (e) {
    e.preventDefault();
    var length = document.getElementById("length").value;
    fetch("/api/generate-password.go?length=" + length, { method: "POST" })
      .then(function (response) {
        return response.json();
      })
      .then(function (data) {
        document.getElementById("password").innerText = data.password;
        document.getElementById("password-container").style.display = "block";
      })
      .catch(function (error) {
        console.error("Error:", error);
      });
  });

document.getElementById("copy-button").addEventListener("click", function () {
  var password = document.getElementById("password").innerText;
  navigator.clipboard.writeText(password);
});
