"use strict";
const emailInput = document.getElementById("email");
const usernameInput = document.getElementById("username");
const passwordInput = document.getElementById("password");
const ageInput = document.getElementById("age");
const submitButn = document.getElementById("signup-btn");
const bannerText = document.getElementById("banner-text");
const domain = "http://localhost:3000";
submitButn.addEventListener("click", async (e) => {
  e.preventDefault();
  const emailVal = emailInput.value;
  const usernameVal = usernameInput.value;
  const passwordVal = passwordInput.value;
  const ageVal = ageInput.value;
  if (bannerText) {
    bannerText.style.color = "black";
    bannerText.textContent = "Signing up...";
  }
  try {
    const jsonBody = JSON.stringify({
      email: emailVal,
      username: usernameVal,
      password: passwordVal,
      age: ageVal,
    });
    const response = await fetch(domain + "/service/signup", {
      method: "POST",
      body: jsonBody,
    });
    const responseJSON = await response.json();
    if (bannerText) {
      if (responseJSON.redirectURL) {
        window.location.href = domain + responseJSON.redirectURL;
      } else {
        bannerText.style.color = "maroon";
        bannerText.textContent = `
    ${responseJSON.emailError && `email: ${responseJSON.emailError}`} \n
    ${responseJSON.ageError && `age: ${responseJSON.ageError}`} \n
    ${responseJSON.generalMessage && `error: ${responseJSON.generalMessage}`} \n
            `;
      }
    }
  } catch (error) {
    console.log(error);
  }
});
