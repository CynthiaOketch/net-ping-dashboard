const form = document.getElementById('ping-form');
const resultsDiv = document.getElementById('results');

form.addEventListener('submit', async (e) => {
  e.preventDefault();
  const urls = document.getElementById('urls').value
    .split('\n')
    .map(url => url.trim())
    .filter(url => url !== "");

  await fetch('http://localhost:8080/ping', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(urls)
  });
});

// Poll results every 5 seconds
setInterval(async () => {
  const res = await fetch('http://localhost:8080/results');
  const data = await res.json();
  renderResults(data);
}, 5000);

function renderResults(data) {
    if (!Array.isArray(data)) {
      resultsDiv.innerHTML = '<p>No results yet. Submit some URLs first.</p>';
      return;
    }
  
    resultsDiv.innerHTML = '';
    data.forEach(item => {
      const div = document.createElement('div');
      div.classList.add('result');
      const statusClass = item.up ? 'status-up' : 'status-down';
      div.innerHTML = `
        <div><strong>${item.url}</strong></div>
        <div>Latency: ${item.latency.toFixed(2)} ms</div>
        <div>Status: <span class="${statusClass}">${item.up ? 'UP' : 'DOWN'}</span></div>
      `;
      resultsDiv.appendChild(div);
    });
  }
  
