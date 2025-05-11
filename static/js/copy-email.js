function copyEmail(event) {
    event.preventDefault();
    const email = "davnub@yandex.ru";
    navigator.clipboard.writeText(email).then(() => {
        showCopyToast();
    }).catch(() => {});
}

function showCopyToast() {
    const toast = document.getElementById("copy-toast");
    if (!toast) return;

    toast.style.opacity = "1";
    setTimeout(() => {
        toast.style.opacity = "0";
    }, 2000); // show for 2 seconds
}

