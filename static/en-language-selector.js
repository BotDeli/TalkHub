Object.values(document.getElementsByClassName('not-selected-language')).forEach((node) => {
    node.addEventListener('click', () => {
       fetch('/setRuLanguage').then(() => {
           window.location.reload();
       })
    });
})