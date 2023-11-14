Object.values(document.getElementsByClassName('not-selected-language')).forEach((node) => {
    node.addEventListener('click', () => {
       fetch('/setEnLanguage').then(() => {
           window.location.reload();
       })
    });
})