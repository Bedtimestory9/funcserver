"use strict";
const usernameInput = document.getElementById("username");
const passwordInput = document.getElementById("password");
const submitButn = document.getElementById("login-btn");
const domain = "http://localhost:3000";
submitButn.addEventListener("click", async (e) => {
    e.preventDefault();
    const usernameVal = usernameInput.value;
    const passwordVal = passwordInput.value;
    try {
        const response = await fetch(domain + "/service/validation/login-validation", {
            method: "POST",
            body: JSON.stringify([{
                    username: usernameVal,
                    password: passwordVal,
                }])
        });
        const result = await response.text();
        console.log("Success", result);
    }
    catch (error) {
        console.log(error);
    }
});
