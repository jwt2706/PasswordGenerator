document.getElementById('password-form').addEventListener('submit', function(event) {
  event.preventDefault();
  
  let length = document.getElementById('length').value;
  
  fetch('https://backend-mauve-delta.vercel.app/generate-password', {
    method: 'POST',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ length: length }),
  })
  .then(response => response.json())
  .then(data => {
    document.getElementById('password').textContent = data.password;
  })
  .catch((error) => {
    console.error('Error:', error);
  });
});