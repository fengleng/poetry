function GetCodeImg() {
    var img = document.getElementById("imgCode");
    var src = img.src;
    img.src = src + "?t=" + new Date().valueOf().toString();
}