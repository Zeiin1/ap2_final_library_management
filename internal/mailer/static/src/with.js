
const header = document.querySelector('header');

window.addEventListener("scroll", () => {
    if (window.scrollY >= 450) {
        header.classList.add("scrolled")
    }
    else if (window.scrollY <= 400) {
        header.classList.remove("scrolled")
    }
});


function aituon() {
    document.getElementById("cardimg").src = "/URLToReach/img/asaitu.jpg"
}
function aituoff() {
    document.getElementById("cardimg").src = "/URLToReach/img/mac.jpg"
}

function kaznuion() {
    document.getElementById("cardimg1").src = "/URLToReach/img/kaznuilogo.jpg"
}
function kaznuioff() {
    document.getElementById("cardimg1").src = "/URLToReach/img/kazakhskij-natsionalnyj-universitet-iskusstv.jpg"
}

function sduon() {
    document.getElementById("cardimg2").src = "/URLToReach/img/sdulogo.jpg"
}
function sduoff() {
    document.getElementById("cardimg2").src = "/URLToReach/img/scale_1200.jfif"
}

function kasguuon() {
    document.getElementById("cardimg3").src = "/URLToReach/img/kazguulogo.jpg"
}
function kasguuoff() {
    document.getElementById("cardimg3").src = "/URLToReach/img/PVeuQczq.jpg"
}

function kbtuon() {
    document.getElementById("cardimg4").src = "img/kbtulogo.jpg"
}
function kbtuoff() {
    document.getElementById("cardimg4").src = "img/KBTU-building.jpg"
}

function kstuon() {
    document.getElementById("cardimg5").src = "img/kargtulogo.jpg"
}
function kstuoff() {
    document.getElementById("cardimg5").src = "img/kstu.jpg"
}


/*sound*/
function play() {
    var audio = new Audio("sounds/playsound.wav");
    audio.play();
}




/*side bar*/
function openNav() {
    document.getElementById("mySidenav").style.width = "250px";
}


function closeNav() {
    document.getElementById("mySidenav").style.width = "0";
}


document.onkeypress = function (event) {
    console.log(event.key);
}




var acc = document.getElementsByClassName("accordion");
var i;

for (i = 0; i < acc.length; i++) {
    acc[i].addEventListener("click", function () {

        this.classList.toggle("active");


        var panel = this.nextElementSibling;
        if (panel.style.display === "block") {
            panel.style.display = "none";
        } else {
            panel.style.display = "block";
        }
    });
}
