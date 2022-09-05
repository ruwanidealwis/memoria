async function performPostHttpRequest(fetchLink, body, headers) {
    if (!fetchLink || !body) {
        throw new Error("Missing POST request parameters");
    }
    try {

        let rawResponse;
        let requestBody = body;
        if (headers?.["Content-Type"] == "application/json") {
            requestBody = JSON.stringify(body);
        }
        if (headers) {
            rawResponse = await fetch(fetchLink, {
                method: "POST",
                body: requestBody,
                headers: headers,
            })
        }
        else {
            rawResponse = await fetch(fetchLink, {
                method: "POST",
                body: requestBody,
            })
        }
        if (!rawResponse.ok) {
            res = await rawResponse.json()
            throw new Error(res.error)
        }

        return await rawResponse.json();
    } catch (err) {
        console.error(`Error at fetch POST: ${err}`);
        throw err;
    }
}

function logout() {
    sessionStorage.setItem("userID", null);
    window.location = `/`;
};

async function performHttpGetRequest(fetchLink, params = {}) {
    if (!fetchLink) {
        throw new Error("Missing GET request parameters");
    }
    try {
        let url = fetchLink;
        if (params) {
            url += `?${new URLSearchParams(params).toString()}`
        }
        const rawResponse = await fetch(url);
        if (!rawResponse.ok) {
            res = await rawResponse.json()
            throw new Error(res.error)
        }

        return await rawResponse.json();
    } catch (err) {
        console.error(`Error at fetch GET: ${err} `);
        throw err;
    }
}

function buildHeaders(type = "application/json") {
    const headers = {
        "Content-Type": type,
    };
    return headers
}