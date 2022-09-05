window.onload = () => {
    const signupForm = document.querySelector("#sign-up-form");

    document.getElementById("sign-up-form").addEventListener("submit", e => signup(e, signupForm))
}

const signup = async (e, form) => {
    e.preventDefault();
    const formData = new FormData(form);

    try {
        const response = await performPostHttpRequest(
            `/user/sign-up`,
            formData
        );

        if (response) {
            sessionStorage.setItem("userID", response.ID)
            window.location = `/scrapbooks`;
        }
    } catch (e) {
        document.getElementById("sign-up-error").innerHTML = e.message;
    }


}
