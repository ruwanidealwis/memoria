window.onload = () => {

    document.getElementById("login-form").addEventListener("submit", e => signup(e))
}

const signup = async (e) => {
    e.preventDefault();
    const data = {
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
    }
    const headers = buildHeaders("application/json")
    try {
        const response = await performPostHttpRequest(
            `/user/login`,
            data,
            headers
        );

        if (response) {
            sessionStorage.setItem("userID", response.ID)
            window.location = `/scrapbooks`;
        }
    } catch (e) {
        document.getElementById("login-error").innerHTML = e.message;
    }


}