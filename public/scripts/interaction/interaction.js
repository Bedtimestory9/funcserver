"use strict";
const displayDiv1 = document.getElementsByClassName("display1")[0];
const nameDiv = document.getElementsByClassName("d1")[0];
const moodDiv = document.getElementsByClassName("d2")[0];
const button1 = document.getElementsByClassName("btn1")[0];
const button2 = document.getElementsByClassName("btn2")[0];
const baseURL = window.location.pathname;
const getRequestURL = "http://127.0.0.1:3000/service/interaction/get-user-mood";
button1.addEventListener("click", () => {
    displayDiv1.textContent = "Waiting for server response...";
    void (async function () {
        try {
            const res = await fetch(getRequestURL);
            button1.disabled = true;
            button1.textContent = "Disabled";
            if (!res.ok) {
                throw new Error(`HTTP error, status = ${res.status}`);
            }
            await new Promise((resolve) => setTimeout(resolve, 2000));
            const textContent = await res.text();
            displayDiv1.textContent = textContent;
            button1.disabled = false;
            button1.textContent = "Click Me";
        }
        catch (error) {
            button1.disabled = true;
            button1.textContent = "Disabled";
            console.error(error);
        }
    })();
});
button2.addEventListener("click", () => {
    nameDiv.textContent = "Waiting response...";
    moodDiv.textContent = "Waiting response...";
    void (async function () {
        const res = await fetch("http://127.0.0.1:3000/service/interaction/get-user-mood", {});
        if (!res.ok) {
            throw new Error(`HTTP error, status = ${res.status}`);
        }
        await new Promise((resolve) => setTimeout(resolve, 1000));
        const json = await res.json();
        nameDiv.textContent = json.name;
        moodDiv.textContent = json.mood.toString();
    })();
});
