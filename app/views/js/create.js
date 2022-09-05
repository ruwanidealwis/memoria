let tabButtons = [document.querySelectorAll(".drop-down")];

window.onload = function () {
    showCanvasContent("get-started-btn");

    function displayPageTitle() {
        document.getElementById("title").innerHTML = pageTitleEl.value;
    }

    document.getElementById("logout-btn").addEventListener("click", function () {
        sessionStorage.setItem("userID", null);
        window.location = `/`;
    })

    document
        .getElementById("scrapbooks-tab")
        .addEventListener("click", function () {
            window.location = `/scrapbooks`;
        });

    document
        .getElementById("image-3-upload")
        .addEventListener("change", function () {
            displayImageName("image-3-upload");
        });

    document
        .getElementById("image-2-upload")
        .addEventListener("change", function () {
            displayImageName("image-2-upload");
        });

    document
        .getElementById("image-1-upload")
        .addEventListener("change", function () {
            displayImageName("image-1-upload");
        });


    let pageTitleEl = document.getElementById("page-title");
    pageTitleEl.addEventListener("change", function () {
        displayPageTitle();
    });

    let getStartedEl = document.getElementById("get-started-btn");
    getStartedEl.addEventListener("click", function () {
        showCanvasContent("get-started-btn");
    });

    let imagesEl = document.getElementById("images-btn");
    imagesEl.addEventListener("click", function () {
        showCanvasContent("images-btn");
    });

    let textEl = document.getElementById("text-btn");
    textEl.addEventListener("click", function () {
        showCanvasContent("text-btn");
    });

    let musicEl = document.getElementById("music-btn");
    musicEl.addEventListener("click", function () {
        showCanvasContent("music-btn");
    });

    let mapsEl = document.getElementById("maps-btn");
    mapsEl.addEventListener("click", function () {
        showCanvasContent("maps-btn");
    });

    const createPageForm = document.querySelector("#create-page-form");
    let generateButton = document.getElementById("generate-btn");

    generateButton.addEventListener("click", function (e) {
        submitForm(e, createPageForm);
    });
    
    document
        .getElementById("search-song-button")
        .addEventListener("click", function (event) {
            event.preventDefault()
            getSongData();

        });


    document
        .getElementById("search-maps-button")
        .addEventListener("click", function (event) {
            event.preventDefault()
            getMapData();

        });


};

// Toggle content when tabs are clicked
function showCanvasContent(id) {
    for (let i = 0; i < tabButtons[0].length; i++) {
        if (tabButtons[0][i].id == id) {
            document.getElementById(id.replace("btn", "content")).style.display =
                "block";
        } else {
            document.getElementById(
                tabButtons[0][i].id.replace("btn", "content")
            ).style.display = "none";
        }
    }
}

async function getSongData() {
    const params = { name: document.getElementById("search-song-input").value };
    try {
        const res = await performHttpGetRequest("/search/song", params);
        displaySongInfo(res[0]);
    } catch (err) {
        document.getElementById("form-error").innerHTML = err.message;
    }
}


async function getMapData() {
    const params = { location: document.getElementById("search-map-input").value };
    try {
        const res = await performHttpGetRequest("/map", params);
        displayMapInfo(res);
    } catch (err) {
        document.getElementById("form-error").innerHTML = err.message;
    }
}

async function getScrapbookData() {
    const params = { name: document.getElementById("scrapbook-name").value };
    try {
        const res = await performHttpGetRequest("/scrapbook", params);
        sessionStorage.setItem("scrapbookID", res.ID);
    } catch (err) {
        document.getElementById("form-error").innerHTML = err.message;
    }
}

function displaySongInfo(songData) {
    document.getElementById("song-cover").src = songData.albumCover;
    document.getElementById("song-name").innerHTML = songData.name;
    document.getElementById("song-artist").innerHTML = songData.artist;
    document.getElementById("display-song-content").style.display = 'flex';
    sessionStorage.setItem("spotifyID", songData.ID);
}
function displayMapInfo(mapData) {
    document.getElementById("map-image").src = mapData.image;

    document.getElementById("display-map-content").style.display = 'flex';
    sessionStorage.setItem("mapImage", mapData.image);
}

function displayImageName(imageUploadEl) {
    var file = document.getElementById(imageUploadEl);
    document.getElementById(
        imageUploadEl.replace("upload", "name")
    ).innerHTML = file.files.item(0).name;
}

// taken from https://dev.to/amjadmh73/submit-html-forms-to-json-apis-easily-137l
async function submitForm(e, form) {
    e.preventDefault();
    // disable button while submitting
    const formData = new FormData(form);
    let generateButton = document.getElementById("generate-btn");

    generateButton.disabled = true;
    setTimeout(() => (generateButton.disabled = false), 2000);
    // build json

    getScrapbookData();

    formData.set("song", sessionStorage.getItem("spotifyID"));
    formData.set("maps", sessionStorage.getItem("mapImage"));
    formData.set("scrapbook-id", sessionStorage.getItem("scrapbookID")); //TODO: make me dynamic


    const headers = buildHeaders("multipart/form-data");
    try {
        const response = await performPostHttpRequest(
            `/page`, // placeholder
            formData
        );
        if (response) {
            window.location = `/export?id=${response.ID}`;
        }
    } catch (err) {
        alert(`An error occured creating your form.`)
    }
}



