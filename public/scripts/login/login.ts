const usernameInput = document.getElementById("username") as HTMLInputElement
const passwordInput = document.getElementById("password") as HTMLInputElement
const submitButn = document.getElementById("login-btn") as HTMLButtonElement
const bannerText = document.getElementById("banner-text")

const domain = "http://localhost:3000"

submitButn.addEventListener("click", async (e) => {
    e.preventDefault()
    const usernameVal = usernameInput.value
    const passwordVal = passwordInput.value
    if (bannerText) {
        bannerText.style.color = 'black'
        bannerText.textContent = "Logging in..."
    }

    try {
        const response = await fetch(domain + "/service/validation/login-validation", {
            method: "POST",
            body: JSON.stringify([{
                username: usernameVal,
                password: passwordVal,
            }])
        })

        const status = response.status
        const responseText = await response.text()
        if (bannerText) {
            switch (status) {
                case 200:
                    window.location.href = domain + responseText
                    break;
                case 303:
                    window.location.href = domain + responseText
                    break;
                default:
                    bannerText.style.color = 'maroon'
                    bannerText.textContent = responseText
            }
        }
        await new Promise(resolve => setTimeout(resolve, 3000))
    } catch (error) {
        console.log(error)
    }
})

