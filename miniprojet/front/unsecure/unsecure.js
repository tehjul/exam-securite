async function LoginUnsecure() {
    let message = document.login.message.value;
    fetch("http://localhost:8080/connect", {
        method: "POST",
        body: JSON.stringify({
            user: document.login.Username.value,
            password: document.login.Password.value
        })
    }).then(async res => {
        if (res.ok) {
            document.cookie = "session=" + res['session']
            document.getElementById("main").innerHTML = "<p> Login successfull </p><p>"+message+"</p>";
        } else {
            document.getElementById("main").innerHTML = "<p> Login failed </p><p>"+message+"</p>";
        }
    })
}