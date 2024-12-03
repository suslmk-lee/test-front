async function fetchData() {
    try {
        const response = await fetch('/data');
        const data = await response.json();
        updateUI(data);
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}

function updateUI(data) {
    // Update hostname
    document.getElementById('hostname').textContent = data.hostname;

    // Update IP addresses
    const ipAddressesDiv = document.getElementById('ip-addresses');
    ipAddressesDiv.innerHTML = data.ip_addresses
        .map(ip => `<div class="ip-address">${ip}</div>`)
        .join('');

    // Update time information
    const timeInfoDiv = document.getElementById('time-info');
    timeInfoDiv.innerHTML = `
        <div>UTC: ${data.time.utc}</div>
        <div>KST: ${data.time.kst}</div>
        <div>Timestamp: ${data.time.timestamp}</div>
    `;
}

// Fetch data immediately when page loads
document.addEventListener('DOMContentLoaded', fetchData);
