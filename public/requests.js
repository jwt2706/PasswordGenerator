document
  .getElementById("password-form")
  .addEventListener("submit", function (e) {
    e.preventDefault();
    var length = document.getElementById("length").value;
    document.getElementById("password-loading").style.display = "block";
    fetch("/api/generate-password.go?length=" + length, {
      method: "POST",
    })
      .then(function (response) {
        return response.json();
      })
      .then(function (data) {
        document.getElementById("password").innerText = data.password;
        document.getElementById("password-container").style.display = "block";
        document.getElementById("password-loading").style.display = "none";
      })
      .catch(function (error) {
        console.error("Error:", error);
        document.getElementById("password-loading").style.display = "none";
      });
  });

document.getElementById("copy-button").addEventListener("click", function () {
  var password = document.getElementById("password").innerText;
  navigator.clipboard.writeText(password);
});
