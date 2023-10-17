const eyes = document.getElementsByClassName('eye');
for (let i = 0; i < eyes.length; i++) {
    eyes[i].addEventListener('click', () => {
        let nodeFor = eyes[i].getAttributeNode('for');
        if (nodeFor !== null) {
            let node = document.getElementById(nodeFor.value);
            if (node.getAttributeNode('type').value === 'password') {
                node.setAttribute('type', 'text');
            } else {
                node.setAttribute('type', 'password');
            }
        }
    })
}


const btnCreateAccount = document.getElementById('btn-go-account');

btnCreateAccount.addEventListener('mouseenter', () => {
    btnCreateAccount.style.borderRadius = '40px';
});

btnCreateAccount.addEventListener('mouseleave', () => {
    btnCreateAccount.style.borderRadius = '0';
});