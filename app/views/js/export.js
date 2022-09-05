window.onload = async () => {
    let params = new URLSearchParams(document.location.search);
    const res = await getPageDetails(params.get("id"))
    setText(res);
    setImages(res);
    const song = await getSong({ id: res.song.spotifyID })
    setMusicDetails(song)
    const map = await getMapImage({ id: res.map.image })
    setMapDetails(map);
    const user = await getUserDetails(sessionStorage.getItem("userID"))
    console.log(user)
    setUserDetails(user);

    let createMoreBtn = document.getElementById('create-more-btn');
    createMoreBtn.addEventListener("click", function() {
        window.location = `/create`;
    });

    document.getElementById("logout-btn").addEventListener("click", function () {
        sessionStorage.setItem("userID", null);
        window.location = `/`;
    })
}

getPageDetails = async (params) => {
    return await performHttpGetRequest("/page", params);
}

getPageDetails = async (params) => {
    return await performHttpGetRequest(`/page/${params}`);
}

getUserDetails = async (user) => {
    return await performHttpGetRequest(`/user/${user}`);
}

getMapImage = async (params) => {
    return await performHttpGetRequest("/image", params);
}

getSong = async (id) => {
    return await performHttpGetRequest("/song", id);
}

const setMusicDetails = (song) => {
    document.getElementById("playlist-cover").src = song.albumCover;
    document.getElementById("song-name").innerHTML = song.name;
    document.getElementById("artist-name").innerHTML = song.artist;

}

const setMapDetails = (map) => {
    document.getElementById("map").src = map.file;
}

const setImages = (data) => {
    for (let i = 0; i < 3; i++) {
        document.getElementById(`image-${i + 1}`).src = data.images[i].file;
    }


}

const setText = (data) => {
    document.getElementById("title").innerHTML = data.title;
    document.getElementById("heading-1").innerHTML = data.headingOne;
    document.getElementById("heading-2").innerHTML = data.headingTwo;
    document.getElementById("heading-3").innerHTML = data.headingThree;

}
const setUserDetails = (data) => {
    document.getElementById("user").innerHTML = `By ${data.firstName} ${data.lastName}`
}