document.getElementById("exit-account").addEventListener("click", () => {
   fetch("/exitAccount", {
       method: "DELETE"
   })
       .then(() => document.location.reload())
       .catch(() => console.error("Failed exit account!"));
});