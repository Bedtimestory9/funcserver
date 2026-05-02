const usernameInput = document.getElementById("username") as HTMLInputElement
const passwordInput = document.getElementById("password") as HTMLInputElement
const submitButn = document.getElementById("login-btn") as HTMLButtonElement

const domain = "http://localhost:3000"

submitButn.addEventListener("click", async (e) => {
    e.preventDefault()
    const usernameVal = usernameInput.value
    const passwordVal = passwordInput.value

    try {
        const response = await fetch(domain + "/service/validation/login-validation", {
            method: "POST",
            body: JSON.stringify([{
                username: usernameVal,
                password: passwordVal,
            }])
        })
        const result = await response.text()
        console.log("Success", result)
    } catch (error) {
        console.log(error)
    }
})

