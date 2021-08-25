window.onload = function() {

    const submitButton = document.querySelector("#send");

    submitButton.disabled = true;
    document.getElementById("urlToProccess").innerHTML = "";

    document.getElementById("urlToProccess").addEventListener("keyup", function() {
        var urlToProccess = document.getElementById('urlToProccess').value;
        if (urlToProccess == "") {
            submitButton.setAttribute("disabled");
        } else {
            submitButton.removeAttribute("disabled", null);
        }
    });
}