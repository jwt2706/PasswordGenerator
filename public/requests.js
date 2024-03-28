// password generator
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

// strength test
document
  .getElementById("strength-test-form")
  .addEventListener("submit", function (e) {
    var password = document.getElementById("test-password").value;
    if (!password) {
      e.preventDefault();
      alert("Please enter a password to test.");
      return;
    }
    document.getElementById("strength-loading").style.display = "block";
    fetch("/api/strength-test.go?_=" + new Date().getTime(), {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ password: password }),
    })
      .then(function (response) {
        return response.json();
      })
      .then(function (data) {
        document.getElementById("strength").innerText =
          "Strength: " + data.strength;
        document.getElementById("strength").style.display = "block";
        document.getElementById("strength-loading").style.display = "none";
      })
      .catch(function (error) {
        console.error("Error:", error);
        document.getElementById("strength-loading").style.display = "none";
      });
  });
