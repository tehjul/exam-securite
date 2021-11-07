function LoginUnsecure() {
    let login = document.login.Username.value;
    let password = document.login.Password.value;
    console.log(login);
    if (login == "admin" && password == "4dm1n1str4t0r") {
        window.location.replace("success.html");
    } else {
        alert("Login ou mot de passe incorrect");
    }
}