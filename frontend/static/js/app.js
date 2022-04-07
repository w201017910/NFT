/* Theme Name: Artlink -
  File Description: Main JS file of the template
*/

/********************* Menu Js **********************/

function windowScroll() {
  const navbar = document.getElementById("navbar");
  if (
    document.body.scrollTop >= 50 ||
    document.documentElement.scrollTop >= 50
  ) {
    navbar.classList.add("nav-sticky");
  } else {
    navbar.classList.remove("nav-sticky");
  }
}

window.addEventListener("scroll", (ev) => {
  ev.preventDefault();
  windowScroll();
});


//
/********************* Swicher js ************************/
//

function toggleSwitcher() {
  var i = document.getElementById("style-switcher");
  if (i.style.left === "-189px") {
    i.style.left = "-0px";
  } else {
    i.style.left = "-189px";
  }
}

function setColor(theme) {
  document.getElementById("color-opt").href = "./css/colors/" + theme + ".css";
  toggleSwitcher(false);
}