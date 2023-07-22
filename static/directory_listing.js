const directoryListing = document.getElementById('directoryListing');
const pathOutput = document.getElementById('pathOutput');
const emptyDirectoryWarning = document.getElementById('emptyDirectoryWarning');

function fetchDirectoryListing(hostId, path) {
    if (path.startsWith("//")) {
        path = path.substring(1);
    }

    fetch(`/app/hosts/${hostId}/directory?path=${encodeURIComponent(path)}`)
        .then(response => response.json())
        .then(data => {
            let listingHtml = '';

            if (!data) {
                data = [];
            }

            if (path !== '/') {
                const parentPath = path.substring(0, path.lastIndexOf('/'));
                listingHtml += `
                        <div class="directoryItem">
                            <button type="button" class="btn btn-secondary flex-fill" onclick="fetchDirectoryListing('${hostId}', '${parentPath}')">
                                <i class="fas fa-folder"></i> ..
                            </button>
                        </div>
                    `;
            }

            data.filter(entry => entry.isDir).forEach(directoryEntry => {
                listingHtml += `
                        <div class="directoryItem">
                            <button type="button" class="btn btn-secondary flex-fill" onclick="fetchDirectoryListing('${hostId}', '${path}/${directoryEntry.path}')">
                                <i class="fas fa-folder"></i> ${directoryEntry.path}
                            </button>
                        </div>
                    `;
            });

            directoryListing.innerHTML = listingHtml;
            pathOutput.innerText = path;
            emptyDirectoryWarning.classList.toggle('d-none', !data.length);
            emptyDirectoryWarning.classList.toggle('d-flex', data.length);
        });
}
