async function sha256(message) {
    // encode as UTF-8
    const msgBuffer = new TextEncoder().encode(message);                    
    // hash the message
    const hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer);
    // convert ArrayBuffer to Array
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    // convert bytes to hex string                  
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
    return hashHex;
}

const masterKey = "0ca05ed7ee2a7757cbb61efdf1c2422c1b2b376ad20fa554635870600a371c9b";

async function LoginSecure() {
    let login = document.login.Username.value;
    let password = document.login.Password.value;
    let hash = await sha256(password);
    if (login == "admin" && hash == masterKey) {
        window.location.replace("success.html");
    } else {
        alert("Login ou mot de passe incorrect");
    }
}