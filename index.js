'use strict'

const BASE_URL = "http://localhost:8080"

const generateUserId = () => {
    return Date.now().toString(36) + Math.random().toString(36).substring(2, 5);
};

const onLoad = () => {
    const userId = localStorage.getItem("url_shortener_user_id");

    if (!userId) {
        localStorage.setItem("url_shortener_user_id", generateUserId());
    } else {
        fetch(`${BASE_URL}/user/${userId}/urls`)
            .then((res) => res.json())
            .then((data) => {
                if (data.urls && data.urls.length > 0) {
                    console.log("URLs:", data.urls);
                    displayUserUrls(data.urls);
                } else {
                    console.log(`No URLs found for ${userId}`);
                }
            })
            .catch((err) => {
                console.error("Error retrieving URLs:", err);
            })
    }

    const btn = document.getElementById("generateBtn");
    btn.addEventListener("click", addNewUrl);
};

const deleteUrl = (shortUrl) => {
    const userId = localStorage.getItem("url_shortener_user_id");
    let endpoint = `${BASE_URL}/user/${userId}/urls`
    if (shortUrl) {
        endpoint = `${endpoint}/${shortUrl}`
    }
    fetch(endpoint, { method: "DELETE" })
        .then((res) => res.json())
        .then(() => {
            console.log(`Url [${shortUrl}] deleted successfully`);
            document.location.reload();
        })
        .catch((err) => {
            console.error(`Error deleting Url [${shortUrl}]:`, err);
        })
};

const createUrlCell = (href, content) => {
    const urlCell = document.createElement("td");
    const urlLink = document.createElement("a");
    urlLink.href = href;
    urlLink.target = "_blank";
    urlLink.textContent = content;
    urlCell.appendChild(urlLink);

    return urlCell;
};

const displayUserUrls = (urls) => {
    const table = document.getElementById("urlTableBody");
    const userId = localStorage.getItem("url_shortener_user_id");

    urls.forEach((url) => {
        const row = document.createElement("tr");
        const shortUrlCell = createUrlCell(`http://localhost:8080/${url.shortUrl}?user_id=${userId}`, url.shortUrl);
        row.appendChild(shortUrlCell);

        const fullUrlCell = createUrlCell(url.fullUri, url.fullUrl);
        row.appendChild(fullUrlCell);

        const deleteCell = document.createElement("td");
        const deleteButton = document.createElement("button");
        deleteButton.classList.add("button", "is-danger");
        deleteButton.textContent = "Delete"
        deleteButton.onclick = () => deleteUrl(url.shortUrl);
        deleteCell.appendChild(deleteButton);
        row.appendChild(deleteCell);

        table.appendChild(row);
    });
};

const addNewUrl = () => {
    const fullUrl = document.getElementById("fullUrl").value?.trim();
    const userId = localStorage.getItem("url_shortener_user_id");

    if (!fullUrl) {
        return;
    }

    fetch(`${BASE_URL}/generate`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            full_url: fullUrl,
            user_id: userId
        }),
    })
        .then((res) => res.json())
        .then(() => {
            console.log(`Short Url for ${fullUrl} created successfully`);
            document.location.reload();
        })
        .catch((err) => {
            console.error("Error generating short URL", err);
        });
}