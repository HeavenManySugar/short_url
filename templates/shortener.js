document.getElementById('shorten-form')
    .addEventListener('submit', async function(event) {
      event.preventDefault();  // Prevent form from refreshing the page

      const urlInput = document.getElementById('url').value;
      const resultDiv = document.getElementById('result');

      try {
        const response = await fetch('/shorten', {
          method: 'POST',
          headers: {'Content-Type': 'application/json'},
          body: JSON.stringify({url: urlInput})
        });

        if (!response.ok) {
          throw new Error('Failed to shorten URL');
        }

        const data = await response.json();
        const shortUrl = `${window.location.origin}/${data.hash}`;
        resultDiv.innerHTML = `
            <p>Shortened URL: 
                <a href="${shortUrl}" target="_blank">${shortUrl}</a>
            </p>
            <button id="copy-button">Copy URL</button>
        `;

        document.getElementById('copy-button')
            .addEventListener('click', function() {
              navigator.clipboard.writeText(shortUrl)
                  .then(() => {
                    alert('URL copied to clipboard!');
                  })
                  .catch(err => {
                    alert('Failed to copy URL: ' + err);
                  });
            });
      } catch (error) {
        resultDiv.innerHTML =
            `<p style="color: red;">Error: ${error.message}</p>`;
      }
    });