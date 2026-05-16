"use strict";
const usernameInput = document.getElementById("username");
const passwordInput = document.getElementById("password");
const submitButn = document.getElementById("login-btn");
const bannerText = document.getElementById("banner-text");
const domain = "http://localhost:3000";
submitButn.addEventListener("click", async (e) => {
  e.preventDefault();
  const usernameVal = usernameInput.value;
  const passwordVal = passwordInput.value;
  if (bannerText) {
    bannerText.style.color = "black";
    bannerText.textContent = "Logging in...";
  }
  try {
    const response = await fetch(domain + "/service/login", {
      method: "POST",
      body: JSON.stringify([
        {
          username: usernameVal,
          password: passwordVal,
        },
      ]),
    });
    const responseJSON = await response.json();
    if (bannerText) {
      switch (responseJSON.result) {
        case "success":
          window.location.href = domain + responseJSON.redirectURL;
          break;
        default:
          bannerText.style.color = "maroon";
          bannerText.textContent = responseJSON.message;
      }
    }
    await new Promise((resolve) => setTimeout(resolve, 3000));
  } catch (error) {
    console.log(error);
  }
});
