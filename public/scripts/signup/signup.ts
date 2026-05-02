const emailInput = document.getElementById("email") as HTMLInputElement
const usernameInput = document.getElementById("username") as HTMLInputElement
const passwordInput = document.getElementById("password") as HTMLInputElement
const ageInput = document.getElementById("age") as HTMLInputElement
const submitButn = document.getElementById("signup-btn") as HTMLButtonElement
const bannerText = document.getElementById("banner-text")

const domain = "http://localhost:3000"

type UserSignupResponse = {
    emailError: string
    usernameError: string
    passwordError: string
    ageError: string
    generalMessage: string
    redirectURL: string
}

submitButn.addEventListener("click", async (e) => {
    e.preventDefault()
    const emailVal = emailInput.value
    const usernameVal = usernameInput.value
    const passwordVal = passwordInput.value
    const ageVal = ageInput.value
    if (bannerText) {
        bannerText.style.color = 'black'
        bannerText.textContent = "Signing up..."
    }

    try {
        const jsonBody = JSON.stringify([{
            email: emailVal,
            username: usernameVal,
            password: passwordVal,
            age: ageVal
        }])

        const response = await fetch(domain + "/service/validation/signup-validation", {
            method: "POST",
            body: jsonBody
        })


        const responseJSON: UserSignupResponse = await response.json()
        if (bannerText) {
            if (responseJSON.redirectURL) {
                window.location.href = domain + responseJSON.redirectURL
            } else {
                bannerText.style.color = 'maroon'
                bannerText.textContent = `
    ${responseJSON.emailError && `email: ${responseJSON.emailError}`} \n
    ${responseJSON.ageError && `age: ${responseJSON.ageError}`} \n
    ${responseJSON.generalMessage && `error: ${responseJSON.generalMessage}`} \n
            `
            }

        }
    } catch (error) {
        console.log(error)
    }
})

